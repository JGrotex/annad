package state

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var networkNotFoundError = errgo.New("network not found")

// IsNetworkNotFound checks for the given error to be networkNotFoundError.
// This error is returned in case there is no network as required.
func IsNetworkNotFound(err error) bool {
	return errgo.Cause(err) == networkNotFoundError
}

var neuronNotFoundError = errgo.New("neuron not found")

// IsNeuronNotFound checks for the given error to be neuronNotFoundError.
// This error is returned in case there is no neuron as required.
func IsNeuronNotFound(err error) bool {
	return errgo.Cause(err) == neuronNotFoundError
}

var bytesNotFoundError = errgo.New("bytes not found")

// IsBytesNotFound checks for the given error to be bytesNotFoundError.
// This error is returned in case there is no bytes as required.
func IsBytesNotFound(err error) bool {
	return errgo.Cause(err) == bytesNotFoundError
}

var coreNotFoundError = errgo.New("core not found")

// IsCoreNotFound checks for the given error to be coreNotFoundError.
// This error is returned in case there is no core as required.
func IsCoreNotFound(err error) bool {
	return errgo.Cause(err) == coreNotFoundError
}

var impulseNotFoundError = errgo.New("impulse not found")

// IsImpulseNotFound checks for the given error to be impulseNotFoundError.
// This error is returned in case there is no impulse as required.
func IsImpulseNotFound(err error) bool {
	return errgo.Cause(err) == impulseNotFoundError
}

var invalidStateReaderError = errgo.New("invalid state reader")

// IsInvalidStateReader checks for the given error to be
// invalidStateReaderError. This error is returned in case there is no valid
// state reader as required.
func IsInvalidStateReader(err error) bool {
	return errgo.Cause(err) == invalidStateReaderError
}

var invalidStateWriterError = errgo.New("invalid state writer")

// IsInvalidStateWriter checks for the given error to be
// invalidStateWriterError. This error is returned in case there is no valid
// state writer as required.
func IsInvalidStateWriter(err error) bool {
	return errgo.Cause(err) == invalidStateWriterError
}

var invalidObjectTypeError = errgo.New("invalid object type")

// IsInvalidObjectType checks for the given error to be invalidObjectTypeError.
// This error is returned in case there is no valid object type as required.
func IsInvalidObjectType(err error) bool {
	return errgo.Cause(err) == invalidObjectTypeError
}
