package utils

import (
	"context"
	"fmt"
	ent "mazza/ent/generated"

	"github.com/gin-gonic/gin"
)

func GetSession(ctx *context.Context) (user *ent.User, company *ent.Company) {
	/*
		This function should be used inside resolvers:
		gc, err := GinContextFromContext(ctx)
	*/
	getGinContext := func(ctx context.Context) (*gin.Context, error) {
		ginContext := ctx.Value("GinContextKey")
		if ginContext == nil {
			err := fmt.Errorf("could not retrieve gin.Context")
			return nil, err
		}

		gc, ok := ginContext.(*gin.Context)
		if !ok {
			err := fmt.Errorf("gin.Context has wrong type")
			return nil, err
		}
		return gc, nil
	}

	ginCtx, err := getGinContext(*ctx)
	if err != nil {
		return nil, nil
	}
	_user, _ := ginCtx.Get("user")
	user = _user.(*ent.User)

	_c, _ := ginCtx.Get("company")
	company = _c.(*ent.Company)

	return user, company
}
