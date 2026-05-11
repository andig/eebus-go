package ohpcf

import (
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// HandleEvent dispatches SPINE events to the OHPCF handlers.
func (e *OHPCF) HandleEvent(payload spineapi.EventPayload) {
	if !e.IsCompatibleEntityType(payload.Entity) {
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate {
		return
	}

	if _, ok := payload.Data.(*model.SmartEnergyManagementPsDataType); ok {
		e.compressorPowerSequencesUpdate(payload)
	}
}

// compressorPowerSequencesUpdate is called when the compressor's
// SmartEnergyManagementPs data was updated. The CEM forwards a generic
// data-update event so the caller can re-read via PowerConsumptionAlternatives.
func (e *OHPCF) compressorPowerSequencesUpdate(payload spineapi.EventPayload) {
	if e.EventCB == nil {
		return
	}
	e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdatePowerConsumptionAlternatives)
}
