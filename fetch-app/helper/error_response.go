package helper

import "github.com/kataras/iris/v12"

// ErrResponse represent error response
type ErrResponse struct {
	reponse Response
}

// CreateErrorResponse create error response
func CreateErrorResponse(ctx iris.Context, err error) ErrResponse {
	return ErrResponse{
		CreateResponse(ctx).SetMessage(err.Error()),
	}
}

// BadRequest 400
func (err ErrResponse) BadRequest() Response {
	return err.reponse.SetStatus(400)
}

// InternalServer 500
func (err ErrResponse) InternalServer() Response {
	return err.reponse.SetStatus(500)
}

// Unauthorized 401
func (err ErrResponse) Unauthorized() Response {
	return err.reponse.SetStatus(401)
}

// Forbidden 403
func (err ErrResponse) Forbidden() Response {
	return err.reponse.SetStatus(403)
}

// NotFound 404
func (err ErrResponse) NotFound() Response {
	return err.reponse.SetStatus(404)
}
