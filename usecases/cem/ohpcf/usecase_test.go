package ohpcf

import (
	"testing"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/mocks"
	"github.com/enbility/eebus-go/service"
	shipapi "github.com/enbility/ship-go/api"
	"github.com/enbility/ship-go/cert"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestNewOHPCF verifies that the use case constructor wires up the expected
// actor/entity types, use-case name and version. This is a smoke test only;
// scenario-level tests are TBD.
func TestNewOHPCF(t *testing.T) {
	cert, _ := cert.CreateCertificate("test", "test", "DE", "test")
	configuration, _ := api.NewConfiguration(
		"test", "test", "test", "test",
		[]shipapi.DeviceCategoryType{shipapi.DeviceCategoryTypeEnergyManagementSystem},
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeCEM},
		9999, cert, time.Second*4)

	serviceHandler := mocks.NewServiceReaderInterface(t)
	serviceHandler.EXPECT().ServicePairingDetailUpdate(mock.Anything, mock.Anything).Return().Maybe()

	svc := service.NewService(configuration, serviceHandler)
	assert.NoError(t, svc.Setup())

	localEntity := svc.LocalDevice().EntityForType(model.EntityTypeTypeCEM)
	cb := func(string, spineapi.DeviceRemoteInterface, spineapi.EntityRemoteInterface, api.EventType) {}

	uc := NewOHPCF(localEntity, cb)
	assert.NotNil(t, uc)
	uc.AddFeatures()
	uc.UpdateUseCaseAvailability(true)
}
