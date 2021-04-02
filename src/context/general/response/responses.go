package response

import (
	"os"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

// SuccessResponse - Estructure representing success
type SuccessResponse struct {
	Meta    meta        `json:"meta"`
	Records interface{} `json:"records"`
}

// ErrorReponse - Estructure representing error
type ErrorReponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
	MoreInfo  string `json:"moreInfo"`
}

// Meta - Default return success
type meta struct {
	Server      string `json:"server"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	RecordCount int    `json:"recordCount"`
}

// GenerateSuccessResponse - Made a success response
func GenerateSuccessResponse(obj interface{}, limit, offset, recordCount int) SuccessResponse {
	var successResponse SuccessResponse

	hostName, _ := os.Hostname()
	successResponse.Meta.Server = hostName

	successResponse.Meta.Limit = recordCount
	if limit > 0 {
		successResponse.Meta.Limit = limit
	}

	successResponse.Meta.Offset = offset
	successResponse.Meta.RecordCount = recordCount

	if reflect.TypeOf(obj).Kind() != reflect.Slice {
		records := make([]interface{}, 1)
		records[0] = obj
		successResponse.Records = records
		return successResponse
	}

	successResponse.Records = obj

	return successResponse
}

// GenerateErrorResponse - Made a error response
func GenerateErrorResponse(message, moreInfo string, errorCode int) ErrorReponse {
	errorResponse := ErrorReponse{
		Message:   message,
		ErrorCode: errorCode,
		MoreInfo:  moreInfo,
	}

	return errorResponse
}

// GetPagingParameters - Find query params limit and offset
func GetPagingParameters(r echo.Context) (int, int, string) {
	var limit, offset int

	limitParam := r.QueryParam("limit")
	offsetParam := r.QueryParam("offset")

	if (limitParam != "" && offsetParam == "") || (limitParam == "" && offsetParam != "") {
		return 0, 0, "The params: 'limit' and 'offset' are mandatories"
	}

	if limitParam != "" && offsetParam != "" {
		limit, err := strconv.Atoi(r.QueryParam("limit"))
		if err != nil {
			return 0, 0, "Error trying convert the 'limit' param"
		}

		offset, err = strconv.Atoi(r.QueryParam("offset"))
		if err != nil {
			return 0, 0, "Error trying convert the 'offset' param"
		}

		return limit, offset, ""
	}

	return limit, offset, ""
}
