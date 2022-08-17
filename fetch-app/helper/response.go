package helper

import (
	"github.com/kataras/iris/v12"
)

type Response struct {
	ctx     iris.Context
	content struct {
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
		Message *string     `json:"message"`
	}
}

// CreateResponse create new response
func CreateResponse(ctx iris.Context) (response Response) {
	response.ctx = ctx
	return
}

// Ok http 200
func (response Response) Ok() Response {
	response.content.Status = 200
	return response
}

// SetData set data
func (response Response) SetData(data interface{}) Response {
	response.content.Data = data
	return response
}

// SetStatus set status
func (response Response) SetStatus(status int) Response {
	response.content.Status = status
	return response
}

// SetMessage set message
func (response Response) SetMessage(message string) Response {
	response.content.Message = &message
	return response
}

// JSON send response as JSON
func (response Response) JSON() {
	response.ctx.JSON(response.content)
}
