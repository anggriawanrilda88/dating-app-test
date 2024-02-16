package response

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dating-app-test/src/infra/constants"
	"github.com/gin-gonic/gin"
)

// ResponseMessage consist of payload details
// Data -> Payload
type responseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type responseMessageOnly struct {
	Message string `json:"message,omitempty"`
}

func ResponseJSON(
	c *gin.Context,
	message string,
	data interface{},
) {
	if data == nil {
		c.JSON(http.StatusOK, responseMessageOnly{
			Message: message,
		})
	} else {
		c.JSON(http.StatusOK, responseMessage{
			Message: message,
			Data:    data,
		})
	}
}

type CustomError struct {
	HttpErrCode int    `json:"status"`
	Message     string `json:"message"`
}

func ErrorHandler(c *gin.Context, err error) {
	// mapping error from string message and split to slice
	errMessage := err.Error()
	errMessage = strings.ReplaceAll(errMessage, "httpStatus:", "")
	errMessage = strings.ReplaceAll(errMessage, "message:", "")
	errSlice := strings.Split(errMessage, ";")

	if strings.Contains(errMessage, "strconv.ParseUint") {
		errMessage = constants.ErrTypeOfData
	}

	if strings.Contains(errMessage, "strconv.ParseBool") {
		errMessage = constants.ErrTypeOfDataBool
	}

	if strings.Contains(errMessage, "request Content-Type isn't multipart/form-data") {
		errMessage = constants.ErrEmptyRequest
	}

	if strings.Contains(errMessage, "record not found") {
		errMessage = constants.ErrNotFoundRequest
	}

	// set error
	customError := new(CustomError)
	if len(errSlice) == 1 {
		customError.HttpErrCode = 400
		customError.Message = errMessage
		c.JSON(customError.HttpErrCode, customError)
	} else {
		statusStr := errSlice[0]
		status, _ := strconv.Atoi(statusStr)

		customError.HttpErrCode = status
		customError.Message = errSlice[1]
		c.JSON(customError.HttpErrCode, customError)
	}
}
