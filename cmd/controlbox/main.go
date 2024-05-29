package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/client"
	"github.com/enbility/eebus-go/service"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/eg/lpc"
	shipapi "github.com/enbility/ship-go/api"
	"github.com/enbility/ship-go/cert"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

var remoteSki string

type controlbox struct {
	myService *service.Service

	uclpc *lpc.EgLPC

	isConnected bool
}

func (h *controlbox) run() {
	var err error
	var certificate tls.Certificate

	if len(os.Args) == 5 {
		remoteSki = os.Args[2]

		certificate, err = tls.LoadX509KeyPair(os.Args[3], os.Args[4])
		if err != nil {
			usage()
			log.Fatal(err)
		}
	} else {
		certificate, err = cert.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-01")
		if err != nil {
			log.Fatal(err)
		}

		pemdata := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certificate.Certificate[0],
		})
		fmt.Println(string(pemdata))

		b, err := x509.MarshalECPrivateKey(certificate.PrivateKey.(*ecdsa.PrivateKey))
		if err != nil {
			log.Fatal(err)
		}
		pemdata = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
		fmt.Println(string(pemdata))
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		log.Fatal(err)
	}

	configuration, err := api.NewConfiguration(
		"Demo", "Demo", "ControlBox", "123456789",
		model.DeviceTypeTypeElectricitySupplySystem,
		[]model.EntityTypeType{model.EntityTypeTypeGridGuard},
		port, certificate, 230, time.Second*4)
	if err != nil {
		log.Fatal(err)
	}
	configuration.SetAlternateIdentifier("Demo-ControlBox-123456789")

	h.myService = service.NewService(configuration, h)
	h.myService.SetLogging(h)

	if err = h.myService.Setup(); err != nil {
		fmt.Println(err)
		return
	}

	localEntity := h.myService.LocalDevice().EntityForType(model.EntityTypeTypeGridGuard)
	h.uclpc = lpc.NewEgLPC(localEntity, h.OnLPCEvent)
	h.myService.AddUseCase(h.uclpc)

	if len(remoteSki) == 0 {
		os.Exit(0)
	}

	h.myService.RegisterRemoteSKI(remoteSki)

	h.myService.Start()
	// defer h.myService.Shutdown()
}

// EEBUSServiceHandler

func (h *controlbox) RemoteSKIConnected(service api.ServiceInterface, ski string) {
	h.isConnected = true

	fmt.Println("Sending a limit in 5s...")
	time.AfterFunc(time.Second*5, func() {
		remoteDevice := service.LocalDevice().RemoteDeviceForSki(ski)

		// search a compatible entity on the remote device
		for _, entity := range remoteDevice.Entities() {
			if ok, err := h.uclpc.IsUseCaseSupported(entity); err != nil || !ok {
				continue
			}

			limit := ucapi.LoadLimit{
				Duration: time.Minute * 2,
				IsActive: true,
				Value:    7000,
			}
			msgCounter, err := h.uclpc.WriteConsumptionLimit(entity, limit)
			if err != nil {
				fmt.Println("Failed to send limit", err)
				continue
			}

			// add a callback to the result
			if lc, err := client.NewLoadControl(h.uclpc.LocalEntity, entity); err == nil {
				err := lc.AddResponseCallback(*msgCounter, func(msg spineapi.ResponseMessage) {
					// the response has to be of model.ResultDataType
					response, ok := msg.Data.(*model.ResultDataType)
					if !ok || response.ErrorNumber == nil {
						return
					}

					if *response.ErrorNumber == model.ErrorNumberTypeNoError {
						fmt.Println("Limit accepted.")
					} else {
						fmt.Println("Limit rejected. Code", response.ErrorNumber, "Description", response.Description)
					}
				})
				if err != nil {
					fmt.Println("Failed to add callback", err)
				}
			}

			fmt.Println("Sent limit to", ski, "with msgCounter", msgCounter)
		}
	})
}

func (h *controlbox) RemoteSKIDisconnected(service api.ServiceInterface, ski string) {
	h.isConnected = false
}

func (h *controlbox) VisibleRemoteServicesUpdated(service api.ServiceInterface, entries []shipapi.RemoteService) {
}

func (h *controlbox) ServiceShipIDUpdate(ski string, shipdID string) {}

func (h *controlbox) ServicePairingDetailUpdate(ski string, detail *shipapi.ConnectionStateDetail) {
	if ski == remoteSki && detail.State() == shipapi.ConnectionStateRemoteDeniedTrust {
		fmt.Println("The remote service denied trust. Exiting.")
		h.myService.CancelPairingWithSKI(ski)
		h.myService.UnregisterRemoteSKI(ski)
		h.myService.Shutdown()
		os.Exit(0)
	}
}

func (h *controlbox) AllowWaitingForTrust(ski string) bool {
	return ski == remoteSki
}

// LPC Event Handler

func (h *controlbox) OnLPCEvent(ski string, device spineapi.DeviceRemoteInterface, entity spineapi.EntityRemoteInterface, event api.EventType) {
	if !h.isConnected {
		return
	}

	switch event {
	case lpc.DataUpdateLimit:
		if currentLimit, err := h.uclpc.ConsumptionLimit(entity); err == nil {
			fmt.Println("New Limit received", currentLimit.Value, "W")
		}
	default:
		return
	}
}

// main app
func usage() {
	fmt.Println("First Run:")
	fmt.Println("  go run /cmd/controlbox/main.go <serverport>")
	fmt.Println()
	fmt.Println("General Usage:")
	fmt.Println("  go run /cmd/controlbox/main.go <serverport> <evse-ski> <crtfile> <keyfile>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	h := controlbox{}
	h.run()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit
}

// Logging interface

func (h *controlbox) Trace(args ...interface{}) {
	h.print("TRACE", args...)
}

func (h *controlbox) Tracef(format string, args ...interface{}) {
	h.printFormat("TRACE", format, args...)
}

func (h *controlbox) Debug(args ...interface{}) {
	h.print("DEBUG", args...)
}

func (h *controlbox) Debugf(format string, args ...interface{}) {
	h.printFormat("DEBUG", format, args...)
}

func (h *controlbox) Info(args ...interface{}) {
	h.print("INFO ", args...)
}

func (h *controlbox) Infof(format string, args ...interface{}) {
	h.printFormat("INFO ", format, args...)
}

func (h *controlbox) Error(args ...interface{}) {
	h.print("ERROR", args...)
}

func (h *controlbox) Errorf(format string, args ...interface{}) {
	h.printFormat("ERROR", format, args...)
}

func (h *controlbox) currentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (h *controlbox) print(msgType string, args ...interface{}) {
	value := fmt.Sprintln(args...)
	fmt.Printf("%s %s %s", h.currentTimestamp(), msgType, value)
}

func (h *controlbox) printFormat(msgType, format string, args ...interface{}) {
	value := fmt.Sprintf(format, args...)
	fmt.Println(h.currentTimestamp(), msgType, value)
}
