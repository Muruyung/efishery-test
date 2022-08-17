package fish

import (
	"MyAPI/controller"
	"MyAPI/helper"
	"errors"

	"github.com/kataras/iris/v12"
)

// ReadList controller using get method
func ReadList(ctx iris.Context) {
	payload, err := helper.BearerTokenAuth(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return
	}

	user, err := controller.Service.GetUserById(payload["sub"].(string))
	if err != nil || user.Id == "" {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return
	}

	fish, err := controller.Service.GetFishList()
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(fish).JSON()
}

func Aggregate(ctx iris.Context) {
	payload, err := helper.BearerTokenAuth(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return
	}

	user, err := controller.Service.GetUserById(payload["sub"].(string))
	if err != nil || user.Id == "" {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return
	}

	if user.Role != "admin" {
		ctx.StopWithStatus(iris.StatusForbidden)
		helper.CreateErrorResponse(ctx, errors.New("forbidden access")).Forbidden().JSON()
		return
	}

	fish, err := controller.Service.GroupFish()

	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(fish).JSON()
}
