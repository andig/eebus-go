package service

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/mocks"
	shipapi "github.com/enbility/ship-go/api"
	"github.com/enbility/ship-go/cert"
	"github.com/enbility/ship-go/logging"
	shipmocks "github.com/enbility/ship-go/mocks"
	spinemocks "github.com/enbility/spine-go/mocks"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite

	config *api.Configuration

	sut *Service

	serviceReader *mocks.ServiceReaderInterface
	conHub        *shipmocks.HubInterface
	mdns          *shipmocks.MdnsInterface
	logging       *shipmocks.LoggingInterface
	localDevice   *spinemocks.DeviceLocalInterface
}

func (s *ServiceSuite) WriteShipMessageWithPayload(message []byte) {}

func (s *ServiceSuite) BeforeTest(suiteName, testName string) {
	s.serviceReader = mocks.NewServiceReaderInterface(s.T())

	s.conHub = shipmocks.NewHubInterface(s.T())

	s.mdns = shipmocks.NewMdnsInterface(s.T())

	s.logging = shipmocks.NewLoggingInterface(s.T())

	s.localDevice = spinemocks.NewDeviceLocalInterface(s.T())

	certificate := tls.Certificate{}
	var err error
	s.config, err = api.NewConfiguration(
		"vendor", "brand", "model", "serial",
		[]shipapi.DeviceCategoryType{shipapi.DeviceCategoryTypeEnergyManagementSystem},
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeCEM}, 4729, certificate, time.Second*4, nil, nil)
	assert.Nil(s.T(), nil, err)

	s.sut = NewService(s.config, s.serviceReader)
}

func (s *ServiceSuite) Test_AddUseCase() {
	ucMock := mocks.NewUseCaseInterface(s.T())
	ucMock.EXPECT().AddFeatures().Return().Once()
	ucMock.EXPECT().AddUseCase().Return().Once()

	s.sut.AddUseCase(ucMock)
}

func (s *ServiceSuite) Test_EEBUSHandler() {
	testSki := "test"

	s.sut.spineLocalDevice = s.localDevice

	testIdentity := shipapi.NewServiceIdentity(testSki, "", "")

	entry := shipapi.RemoteMdnsService{
		Ski: testSki,
	}

	entries := []shipapi.RemoteMdnsService{entry}
	s.serviceReader.EXPECT().VisibleRemoteMdnsServicesUpdated(mock.Anything, mock.Anything).Return()
	s.sut.VisibleRemoteMdnsServicesUpdated(entries)

	s.serviceReader.EXPECT().RemoteServiceConnected(mock.Anything, mock.Anything).Return()
	s.sut.RemoteServiceConnected(testIdentity)

	s.serviceReader.EXPECT().RemoteServiceDisconnected(mock.Anything, mock.Anything).Return()
	s.localDevice.EXPECT().RemoveRemoteDeviceConnection(testSki).Return()
	s.sut.RemoteServiceDisconnected(testIdentity)

	s.serviceReader.EXPECT().ServiceUpdated(mock.Anything).Return()
	s.sut.ServiceUpdated(testIdentity)

	s.serviceReader.EXPECT().ServicePairingDetailUpdate(mock.Anything, mock.Anything).Return()
	detail := &shipapi.ConnectionStateDetail{}
	s.sut.ServicePairingDetailUpdate(testIdentity, detail)

	s.sut.UserIsAbleToApproveOrCancelPairingRequests(true)
	result := s.sut.AllowWaitingForTrust(testIdentity)
	assert.Equal(s.T(), true, result)

	conf := s.sut.Configuration()
	assert.Equal(s.T(), s.sut.configuration, conf)

	lService := s.sut.LocalService()
	assert.Equal(s.T(), s.sut.localService, lService)
}

func (s *ServiceSuite) Test_ConnectionsHub() {
	testSki := "test"
	testIdentity := shipapi.NewServiceIdentity(testSki, "", "")

	s.sut.connectionsHub = s.conHub
	s.sut.mdns = s.mdns
	s.sut.spineLocalDevice = s.localDevice
	s.sut.localService = shipapi.NewServiceDetails(testSki, "", "")

	s.conHub.EXPECT().PairingDetailFor(mock.Anything).Return(nil)
	s.sut.PairingDetailFor(testIdentity)

	s.conHub.EXPECT().ServiceFor(mock.Anything).Return(nil)
	details := s.sut.RemoteServiceFor(testIdentity)
	assert.Nil(s.T(), details)

	s.localDevice.EXPECT().SetupRemoteDevice(mock.Anything, s).Return(nil)
	s.sut.SetupRemoteService(testIdentity, s)

	s.conHub.EXPECT().SetAutoAccept(mock.Anything).Return()
	s.sut.SetAutoAccept(true)
	assert.True(s.T(), s.sut.IsAutoAcceptEnabled())

	s.conHub.EXPECT().GeneratePairingQR().Return("text", nil)
	qrCode, err := s.sut.QRCodeText()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "text", qrCode)
	s.conHub.EXPECT().RegisterRemoteService(mock.Anything).Return()
	s.sut.RegisterRemoteService(testIdentity)

	s.conHub.EXPECT().UnregisterRemoteService(mock.Anything).Return()
	s.sut.UnregisterRemoteService(testIdentity)

	s.conHub.EXPECT().CancelPairing(mock.Anything).Return()
	s.sut.CancelPairing(testIdentity)

	s.conHub.EXPECT().DisconnectService(mock.Anything, mock.Anything).Return()
	s.sut.DisconnectService(testIdentity, "reason")
}

func (s *ServiceSuite) Test_SetLogging() {
	s.sut.SetLogging(nil)
	assert.Equal(s.T(), &logging.NoLogging{}, logging.Log())

	s.sut.SetLogging(s.logging)
	assert.Equal(s.T(), s.logging, logging.Log())

	s.sut.SetLogging(&logging.NoLogging{})
	assert.Equal(s.T(), &logging.NoLogging{}, logging.Log())
}

func (s *ServiceSuite) Test_Setup() {
	err := s.sut.Setup()
	assert.NotNil(s.T(), err)

	certificate, err := cert.CreateCertificate("unit", "org", "de", "cn")
	assert.Nil(s.T(), err)
	s.config.SetCertificate(certificate)

	err = s.sut.Setup()
	assert.Nil(s.T(), err)

	address := s.sut.LocalDevice().Address()
	assert.Equal(s.T(), "d:_n:vendor_model-serial", string(*address))

	s.sut.connectionsHub = s.conHub
	s.conHub.EXPECT().Start().Return(nil).Once()
	_ = s.sut.Start()

	time.Sleep(time.Millisecond * 200)

	isRunning := s.sut.IsRunning()
	assert.True(s.T(), isRunning)

	// nothing should happen
	_ = s.sut.Start()

	s.conHub.EXPECT().Shutdown().Once()
	s.sut.Shutdown()

	// nothing should happen
	s.sut.Shutdown()

	device := s.sut.LocalDevice()
	assert.NotNil(s.T(), device)
}

func (s *ServiceSuite) Test_Setup_IANA() {
	var err error
	certificate := tls.Certificate{}
	s.config, err = api.NewConfiguration(
		"12345", "brand", "model", "serial",
		[]shipapi.DeviceCategoryType{shipapi.DeviceCategoryTypeEnergyManagementSystem},
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeCEM}, 4729, certificate, time.Second*4, nil, nil)
	assert.Nil(s.T(), nil, err)

	s.sut = NewService(s.config, s.serviceReader)

	err = s.sut.Setup()
	assert.NotNil(s.T(), err)

	certificate, err = cert.CreateCertificate("unit", "org", "de", "cn")
	assert.Nil(s.T(), err)
	s.config.SetCertificate(certificate)

	err = s.sut.Setup()
	assert.Nil(s.T(), err)

	address := s.sut.LocalDevice().Address()
	assert.Equal(s.T(), "d:_i:12345_model-serial", string(*address))

	s.sut.connectionsHub = s.conHub
	s.conHub.EXPECT().Start().Return(nil)
	_ = s.sut.Start()

	time.Sleep(time.Millisecond * 200)

	s.conHub.EXPECT().Shutdown()
	s.sut.Shutdown()

	device := s.sut.LocalDevice()
	assert.NotNil(s.T(), device)
}

func (s *ServiceSuite) Test_Setup_Error_DeviceName() {
	var err error
	certificate := tls.Certificate{}
	s.config, err = api.NewConfiguration(
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
		"brand",
		"modelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodelmodel",
		"serialserialserialserialserialserialserialserialserialserialserialserialserialserialserialserialserial",
		[]shipapi.DeviceCategoryType{shipapi.DeviceCategoryTypeEnergyManagementSystem},
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeCEM}, 4729, certificate, time.Second*4, nil, nil)
	assert.Nil(s.T(), nil, err)

	s.sut = NewService(s.config, s.serviceReader)

	err = s.sut.Setup()
	assert.NotNil(s.T(), err)

	certificate, err = cert.CreateCertificate("unit", "org", "de", "cn")
	assert.Nil(s.T(), err)
	s.config.SetCertificate(certificate)

	err = s.sut.Setup()
	assert.NotNil(s.T(), err)
}
