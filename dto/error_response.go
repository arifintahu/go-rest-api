package dto

func DefaultErrorResponse() BaseResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) BaseResponse {
	return BaseResponse{
		Success:      false,
		MessageTitle: "Oops, something went wrong.",
		Message:      msg,
	}
}

func DefaultErrorInvalidDataWithMessage(msg string) BaseResponse {
	return BaseResponse{
		Success:      false,
		MessageTitle: "Oops, something went wrong.",
		Message:      msg,
	}
}

func DefaultBadRequestResponse() BaseResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
