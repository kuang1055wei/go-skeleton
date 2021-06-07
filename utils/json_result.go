package utils

import "go-skeleton/pkg/common"

type JsonResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

//func JsonPageData(results interface{}, page *Paging) *JsonResult {
//	return JsonData(&PageResult{
//		Results: results,
//		Page:    page,
//	})
//}
//
//func JsonCursorData(results interface{}, cursor string) *JsonResult {
//	return JsonData(&CursorResult{
//		Results: results,
//		Cursor:  cursor,
//	})
//}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      nil,
		Success:   true,
	}
}

func JsonCodeError(err *common.CodeError) *JsonResult {
	return &JsonResult{
		ErrorCode: err.Code,
		Message:   err.Message,
		Data:      err.Data,
		Success:   false,
	}
}

func JsonErrorMsg(message string) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorCode(code int, message string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorData(code int, message string, data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   false,
	}
}

func JsonError(err error) *JsonResult {
	return JsonCodeError(common.FromError(err))
}
