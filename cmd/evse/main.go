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

	"github.com/DerAndereAndi/eebus-go/service"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

type evse struct {
	myService *service.EEBUSService
}

func (h *evse) run() {

	serviceDescription := &service.ServiceDescription{
		DeviceBrand:        "Demo",
		DeviceModel:        "EVSE",
		DeviceSerialNumber: "234567890",
		DeviceIdentifier:   "Demo-EVSE-234567890",
		DeviceType:         model.DeviceTypeTypeChargingStation,
		RemoteDeviceTypes: []model.DeviceTypeType{
			model.DeviceTypeTypeEnergyManagementSystem,
		},
	}

	h.myService = service.NewEEBUSService(serviceDescription)
	h.myService.ServiceImpl = h

	var err error
	var certificate tls.Certificate

	serviceDescription.Port, err = strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		log.Fatal(err)
	}

	remoteSki := os.Args[2]

	fmt.Println(os.Args)
	if len(os.Args) == 5 {
		certificate, err = tls.LoadX509KeyPair(os.Args[3], os.Args[4])
		if err != nil {
			usage()
			log.Fatal(err)
		}
	} else {
		certificate, err = service.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-02")
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

	serviceDescription.Certificate = certificate
	h.myService.Start()
	// defer h.myService.Shutdown()

	remoteService := service.ServiceDetails{
		SKI: remoteSki,
	}
	h.myService.RegisterRemoteService(remoteService)
}

// handle a request to trust a remote service
func (h *evse) RemoteServiceTrustRequested(ski string) {
	// we directly trust it in this example
	h.myService.UpdateRemoteServiceTrust(ski, true)
}

// report the Ship ID of a newly trusted connection
func (h *evse) RemoteServiceShipIDReported(ski string, shipID string) {
	// we should associated the Ship ID with the SKI and store it
	// so the next connection can start trusted
	fmt.Println("SKI", ski, "has Ship ID:", shipID)
}

func usage() {
	fmt.Println("Usage: go run /cmd/evse/main.go <serverport> <hems-ski> <crtfile> <keyfile>")
}

func main() {
	if len(os.Args) < 4 {
		usage()
		return
	}

	h := evse{}
	h.run()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sig:
		// User exit
	}
}
