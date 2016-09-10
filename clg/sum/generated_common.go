package sum

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the clg package. There is the go generate statement placed to invoke clggen.

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

func injectValues(payload spec.NetworkPayload, values []reflect.Value) (spec.NetworkPayload, error) {
	ctx, err := payload.GetContext()
	if err != nil {
		return nil, maskAny(err)
	}

	err = payload.SetArgs(append([]reflect.Value{reflect.ValueOf(ctx)}, values...))
	if err != nil {
		return nil, maskAny(err)
	}

	return payload, nil
}
