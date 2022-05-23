package spine

import (
	"fmt"
	"time"

	"github.com/DerAndereAndi/eebus-go/spine/model"
)

const defaultMaxResponseDelayInSeconds = 10

type FeatureRemoteImpl struct {
	*FeatureImpl
	entity           *EntityRemoteImpl
	functionDataMap  map[model.FunctionType]FunctionData
	maxResponseDelay *model.MaxResponseDelayType
}

func NewFeatureRemoteImpl(id uint, entity *EntityRemoteImpl, ftype model.FeatureTypeType, role model.RoleType) *FeatureRemoteImpl {
	res := &FeatureRemoteImpl{
		FeatureImpl: NewFeatureImpl(
			featureAddressType(id, entity.Address()),
			ftype,
			role),
		entity:          entity,
		functionDataMap: make(map[model.FunctionType]FunctionData),
	}
	for _, fd := range CreateFunctionData[FunctionData](ftype) {
		res.functionDataMap[fd.Function()] = fd
	}

	res.operations = make(map[model.FunctionType]*Operations)

	return res
}

func (r *FeatureRemoteImpl) Data(function model.FunctionType) any {
	return r.functionData(function).DataAny()
}

func (r *FeatureRemoteImpl) SetData(function model.FunctionType, data any) {
	r.functionData(function).SetDataAny(data)

	// TODO: fire event
}

func (r *FeatureRemoteImpl) Sender() Sender {
	return r.Device().Sender()
}

func (r *FeatureRemoteImpl) Device() *DeviceRemoteImpl {
	return r.entity.Device()
}

func (r *FeatureRemoteImpl) SetOperations(functions []model.FunctionPropertyType) {
	r.operations = make(map[model.FunctionType]*Operations)
	for _, sf := range functions {
		r.operations[*sf.Function] = NewOperations(sf.PossibleOperations.Read != nil, sf.PossibleOperations.Write != nil)
	}
}

func (r *FeatureRemoteImpl) SetMaxResponseDelay(delay *model.MaxResponseDelayType) {
	r.maxResponseDelay = delay
}

func (r *FeatureRemoteImpl) MaxResponseDelayDuration() time.Duration {
	//	if r.maxResponseDelay != nil {
	// TODO: parse the ISO8601 string of MaxResponseDelay()
	//	}
	return time.Duration(time.Second * defaultMaxResponseDelayInSeconds)
}

func (r *FeatureRemoteImpl) functionData(function model.FunctionType) FunctionData {
	fd, found := r.functionDataMap[function]
	if !found {
		panic(fmt.Errorf("Data was not found for function '%s'", function))
	}
	return fd
}
