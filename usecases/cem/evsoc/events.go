package evsoc

import (
	"github.com/enbility/eebus-go/features/client"
	internal "github.com/enbility/eebus-go/usecases/internal"
	"github.com/enbility/ship-go/logging"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/util"
)

// handle SPINE events
func (e *EVSOC) HandleEvent(payload spineapi.EventPayload) {
	// only about events from an EV entity or device changes for this remote device

	if !e.IsCompatibleEntityType(payload.Entity) {
		return
	}

	if internal.IsEntityConnected(payload) {
		e.evConnected(payload.Entity)
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate {
		return
	}

	switch payload.Data.(type) {
	case *model.MeasurementListDataType:
		e.evMeasurementDataUpdate(payload)
	default:
		return
	}
}

// an EV was connected
func (e *EVSOC) evConnected(entity spineapi.EntityRemoteInterface) {
	// initialise features, e.g. subscriptions, descriptions
	if evMeasurement, err := client.NewMeasurement(e.LocalEntity, entity); err == nil {
		if _, err := evMeasurement.Subscribe(); err != nil {
			logging.Log().Debug(err)
		}

		// get measurement descriptions
		if _, err := evMeasurement.RequestDescriptions(); err != nil {
			logging.Log().Debug(err)
		}

		// get measurement constraints
		if _, err := evMeasurement.RequestConstraints(); err != nil {
			logging.Log().Debug(err)
		}
	}
}

// the measurement data of an EV was updated
func (e *EVSOC) evMeasurementDataUpdate(payload spineapi.EventPayload) {
	// Scenario 1
	if evMeasurement, err := client.NewMeasurement(e.LocalEntity, payload.Entity); err == nil {
		filter := model.MeasurementDescriptionDataType{
			ScopeType: util.Ptr(model.ScopeTypeTypeStateOfCharge),
		}
		if evMeasurement.CheckEventPayloadDataForFilter(payload.Data, filter) {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateStateOfCharge)
		}
	}
}
