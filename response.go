package daliclient

// Status information
type Status int

// Enum of available status
const (
	StatusSuccess Status = iota
	StatusResponse
	StatusBroadcast
	StatusError
	StatusOther
)

// Response represents a parsed response
type Response struct {
	ResponseStatus Status
	// ResponseData is only set if ResponseStatus is StatusResponse
	ResponseData byte
	// Address is only set if ResponseStatus is StatusBroadcast
	Address byte
	// Command is only set if ResponseStatus is StatusBroadcast
	Command byte
	// Data0 is only set if an unknown ResponseStatus is StatusOther
	Data0 byte
	// Data1 is only set if an unknown ResponseStatus is StatusOther
	Data1 byte
}

// ParseResponse creates a response object based on the raw response
func ParseResponse(response []byte) (Response, error) {
	result := Response{}
	if len(response) < 4 {
		return result, ErrResponseLength
	}

	switch response[1] {
	case 0:
		result.ResponseStatus = StatusSuccess
	case 1:
		result.ResponseStatus = StatusResponse
		result.ResponseData = response[2]
	case 2:
		result.ResponseStatus = StatusBroadcast
		result.Address = response[2]
		result.Command = response[3]
	case 255:
		result.ResponseStatus = StatusError
	default:
		result.ResponseStatus = StatusOther
		result.Data0 = response[2]
		result.Data1 = response[3]
	}

	return result, nil
}
