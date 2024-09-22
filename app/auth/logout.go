package auth

import (
	"context"
	"fmt"

	// "mazza/app/utils"
	"mazza/app/utils"
	u "mazza/ent/utils"
	"net/http"
)

/*
Log out user. It's only worth using this function if the user is
directly logged in from the browser where cookie is used to store the access token.
Otherwise, there is no need to use this function.
*/
func Logout(ctx *context.Context) error {
	// client := ent.FromContext(*ctx)
	ginCtx, err := utils.GetGinContext(*ctx)
	if err == nil {
		return fmt.Errorf("unauthorized")
	}
	currentUser, _, _ := u.GetSession(ctx)
	if currentUser == nil {
		return fmt.Errorf("unauthorized")
	}

	ginCtx.SetSameSite(http.SameSiteLaxMode)
	ginCtx.SetCookie("Authorization", "", -1, "", "localhost", false, true)

	return nil
}
