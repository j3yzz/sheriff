package response

type SuccessResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type ErrorResponse struct {
	Success  bool `json:"success"`
	Message  any  `json:"message"`
	MetaData any  `json:"meta_data"`
}
