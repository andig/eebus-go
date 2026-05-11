package ohpcf

import (
	"github.com/enbility/eebus-go/api"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	"github.com/enbility/eebus-go/usecases/usecase"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

// OHPCF implements the CEM-side of the EEBus use case
// "Optimization of Self-Consumption by Heat Pump Compressor Flexibility" (V1.0.0).
//
// The CEM is a client of the compressor's SmartEnergyManagementPs feature,
// reading announced power-consumption alternatives (Scenario 1) and writing
// the chosen schedule back to control the compressor (Scenario 2).
type OHPCF struct {
	*usecase.UseCaseBase
}

var _ ucapi.CemOHPCFInterface = (*OHPCF)(nil)

// NewOHPCF adds CEM-side support for the OHPCF use case to localEntity.
//
// Parameters:
//   - localEntity: the local entity which should support the use case
//   - eventCB: callback invoked when an event is triggered (optional, may be nil)
func NewOHPCF(localEntity spineapi.EntityLocalInterface, eventCB api.EntityEventCallback) *OHPCF {
	validActorTypes := []model.UseCaseActorType{
		model.UseCaseActorTypeCompressor,
	}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeCompressor,
	}

	useCaseScenarios := []api.UseCaseScenario{
		{
			Scenario:  model.UseCaseScenarioSupportType(1),
			Mandatory: true,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeSmartEnergyManagementPs,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(2),
			Mandatory: true,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeSmartEnergyManagementPs,
			},
		},
	}

	usecase := usecase.NewUseCaseBase(
		localEntity,
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeOptimizationOfSelfConsumptionByHeatPumpCompressorFlexibility,
		"1.0.0",
		"release",
		useCaseScenarios,
		eventCB,
		UseCaseSupportUpdate,
		validActorTypes,
		validEntityTypes,
	)

	uc := &OHPCF{
		UseCaseBase: usecase,
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

// AddFeatures registers the features needed by the CEM actor on the local entity.
//
// The CEM is a SmartEnergyManagementPs client (reads compressor data and writes
// the chosen schedule) and runs the standard DeviceDiagnosis server.
func (e *OHPCF) AddFeatures() {
	// client features
	clientFeatures := []model.FeatureTypeType{
		model.FeatureTypeTypeSmartEnergyManagementPs,
	}
	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}

	// server features
	f := e.LocalEntity.GetOrAddFeature(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
	f.AddFunctionType(model.FunctionTypeDeviceDiagnosisStateData, true, false)
	f.AddFunctionType(model.FunctionTypeDeviceDiagnosisHeartbeatData, true, false)
}
