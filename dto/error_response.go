package dto

func BaseErrorResponse(err error) BaseResponse {
	return DefaultErrorResponseWithMessage(err.Error())
}

func DefaultErrorResponseWithMessage(msg string) BaseResponse {
	return BaseResponse{
		Success:      false,
		MessageTitle: "Oops, something went wrong.",
		Message:      msg,
	}
}

func DefaultBadRequestResponse() BaseResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
