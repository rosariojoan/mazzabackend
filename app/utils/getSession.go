package utils

import (
	"context"
	"mazza/ent"
)

func GetSession(ctx *context.Context) (user *ent.User, company *ent.Company, employeeID *int) {
	ginCtx, err := GetGinContext(*ctx)
	if err != nil {
		return nil, nil, nil
	}
	_user, _ := ginCtx.Get("user")
	user = _user.(*ent.User)

	_c, _ := ginCtx.Get("company")
	company = _c.(*ent.Company)

	_emp, _ := ginCtx.Get("employeeID")
	_employeeID := _emp.(int)

	return user, company, &_employeeID
}
