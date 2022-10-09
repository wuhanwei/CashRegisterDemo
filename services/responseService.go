package services

import "CashRegisterDemo/dtos"

func Response_success(response_order dtos.StructResponseOrderData) dtos.StructResponse {
	response_data := dtos.StructResponse{
		Status:     "OK",
		Message:    "",
		Data:       &response_order,
		StatusCode: "200",
	}
	return response_data
}

func Response_error(err error, code string) *dtos.StructResponse {
	response_data := dtos.StructResponse{
		Status:     "error",
		Message:    err.Error(),
		StatusCode: code,
	}
	return &response_data
}
