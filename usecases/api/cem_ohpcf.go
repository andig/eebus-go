package api

import (
	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// Actor: Customer Energy Management
// UseCase: Optimization of Self-Consumption by Heat Pump Compressor Flexibility
//
// Reference: EEBus UC TS "OptimizationOfSelfConsumptionByHeatPumpCompressorFlexibility" V1.0.0
type CemOHPCFInterface interface {
	api.UseCaseInterface

	// Scenario 1 — Monitor heat pump compressor's power consumption flexibility

	// PowerConsumptionAlternatives returns the SmartEnergyManagementPs data
	// announced by the compressor, describing the optional / scheduled power
	// consumption alternatives.
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such data is (yet) available
	//   - and others
	PowerConsumptionAlternatives(entity spineapi.EntityRemoteInterface) (*model.SmartEnergyManagementPsDataType, error)

	// Scenario 2 — Control heat pump compressor's power consumption flexibility

	// ActivateAlternative selects an announced alternative / powerSequence and
	// writes its requested start time to the compressor.
	//
	// parameters:
	//   - entity: the compressor entity
	//   - alternativesID, sequenceID: identifiers from a prior PowerConsumptionAlternatives reply
	//   - startTime: ISO-8601 relative start time (see spec §3.1.8.2 — relative times only)
	//   - resultCB: callback for the SPINE result response
	ActivateAlternative(
		entity spineapi.EntityRemoteInterface,
		alternativesID uint,
		sequenceID uint,
		startTime string,
		resultCB func(result model.ResultDataType),
	) (*model.MsgCounterType, error)

	// StopAlternative requests the compressor to stop or pause the currently
	// running powerSequence. Whether stop or pause is honoured depends on the
	// alternative's `isStoppable` / `isPausable` flags (spec §[OHPCF-011/5..7]).
	StopAlternative(
		entity spineapi.EntityRemoteInterface,
		alternativesID uint,
		sequenceID uint,
		resultCB func(result model.ResultDataType),
	) (*model.MsgCounterType, error)
}
