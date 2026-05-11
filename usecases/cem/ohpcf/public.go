package ohpcf

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/client"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/util"
)

// PowerConsumptionAlternatives returns the SmartEnergyManagementPs data
// announced by the compressor (Scenario 1, OHPCF-001 / OHPCF-002).
//
// The returned struct describes the optional or scheduled power-consumption
// alternatives — slot power, schedule constraints, interrupt options, etc.
// Callers should consult the EEBus OHPCF UC TS V1.0.0 (§3.2.1.3) for the
// semantics of each field.
func (e *OHPCF) PowerConsumptionAlternatives(entity spineapi.EntityRemoteInterface) (*model.SmartEnergyManagementPsDataType, error) {
	if !e.IsCompatibleEntityType(entity) {
		return nil, api.ErrNoCompatibleEntity
	}

	semp, err := client.NewSmartEnergyManagementPs(e.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	data, err := semp.GetData()
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, api.ErrDataNotAvailable
	}
	return data, nil
}

// ActivateAlternative writes a scheduled start time for the announced
// powerSequence (Scenario 2, OHPCF-004 / OHPCF-021/1). startTime SHALL be a
// SPINE AbsoluteOrRelativeTime value; per §3.1.8.2 OHPCF uses relative times
// (ISO-8601 duration, e.g. "PT0S" for "now").
func (e *OHPCF) ActivateAlternative(
	entity spineapi.EntityRemoteInterface,
	alternativesID uint,
	sequenceID uint,
	startTime string,
	resultCB func(result model.ResultDataType),
) (*model.MsgCounterType, error) {
	if !e.IsCompatibleEntityType(entity) {
		return nil, api.ErrNoCompatibleEntity
	}

	semp, err := client.NewSmartEnergyManagementPs(e.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	data := alternativeWriteData(alternativesID, sequenceID, model.SmartEnergyManagementPsPowerSequenceType{
		Description: &model.PowerSequenceDescriptionDataType{
			SequenceId: util.Ptr(model.PowerSequenceIdType(sequenceID)),
		},
		Schedule: &model.PowerSequenceScheduleDataType{
			SequenceId: util.Ptr(model.PowerSequenceIdType(sequenceID)),
			StartTime:  model.NewAbsoluteOrRelativeTimeType(startTime),
		},
	})

	msg, err := semp.WriteData(data)
	if err != nil {
		return nil, err
	}
	// TODO: wire resultCB through the SPINE result callback registry once a
	// helper for that lands. Until then the caller can poll via events.
	_ = resultCB
	return msg, nil
}

// StopAlternative requests the compressor to invalidate the running
// powerSequence (OHPCF-005). The compressor decides between stop and pause
// based on its announced isStoppable / isPausable flags.
func (e *OHPCF) StopAlternative(
	entity spineapi.EntityRemoteInterface,
	alternativesID uint,
	sequenceID uint,
	resultCB func(result model.ResultDataType),
) (*model.MsgCounterType, error) {
	if !e.IsCompatibleEntityType(entity) {
		return nil, api.ErrNoCompatibleEntity
	}

	semp, err := client.NewSmartEnergyManagementPs(e.LocalEntity, entity)
	if err != nil {
		return nil, err
	}

	data := alternativeWriteData(alternativesID, sequenceID, model.SmartEnergyManagementPsPowerSequenceType{
		Description: &model.PowerSequenceDescriptionDataType{
			SequenceId: util.Ptr(model.PowerSequenceIdType(sequenceID)),
		},
		State: &model.PowerSequenceStateDataType{
			SequenceId: util.Ptr(model.PowerSequenceIdType(sequenceID)),
			State:      util.Ptr(model.PowerSequenceStateTypeInvalid),
		},
	})

	msg, err := semp.WriteData(data)
	if err != nil {
		return nil, err
	}
	_ = resultCB
	return msg, nil
}

// alternativeWriteData wraps a single PowerSequence in the
// SmartEnergyManagementPsData envelope expected for a partial write.
func alternativeWriteData(alternativesID, sequenceID uint, ps model.SmartEnergyManagementPsPowerSequenceType) *model.SmartEnergyManagementPsDataType {
	rel := model.SmartEnergyManagementPsAlternativesRelationType{
		AlternativesId: util.Ptr(model.AlternativesIdType(alternativesID)),
		SequenceId:     []model.PowerSequenceIdType{model.PowerSequenceIdType(sequenceID)},
	}
	return &model.SmartEnergyManagementPsDataType{
		Alternatives: []model.SmartEnergyManagementPsAlternativesType{{
			Relation:      &rel,
			PowerSequence: []model.SmartEnergyManagementPsPowerSequenceType{ps},
		}},
	}
}
