package auth

import (
	"context"
	"fmt"
	"mazza/app/utils"
	"mazza/app/utils/email"
	ent "mazza/ent/generated"
	"mazza/ent/generated/token"
	"mazza/ent/generated/user"
	"time"
)

/*
Check if the username is valid
Create password reset token, store it in the DB and send it to the user email
*/
func ForgotPassword(ctx *context.Context, username string) error {
	client := ent.FromContext(*ctx)
	currentUser, err := client.User.Query().Where(user.UsernameEQ(username)).First(*ctx)

	if err != nil {
		return err
	}
	_ = currentUser

	// Permanently remove old pwd reset tokens of this user
	_, _ = client.Token.Delete().Where(token.And(
		token.CategoryEQ(token.CategoryPasswordReset),
		token.HasUserWith(user.UsernameEQ(username)),
	)).Exec(*ctx)

	// Generate new pwd reset token
	newToken := "frokfo"
	expiry := time.Now().Add(time.Hour * 24)

	// Save the new token
	_, err = client.Token.Create().
		SetExpiry(expiry).
		SetToken(newToken).
		SetCategory(token.CategoryPasswordReset).
		SetUserID(currentUser.ID).
		Save(*ctx)
	if err != nil {
		return err
	}

	// Send email to the user
	go email.SendPasswordResetToken(currentUser.Name, currentUser.Username, newToken)
	return nil
}

/* Verify password reset token. Returns error if the token is not valid, otherwise, return nil */
func VerifyPasswordResetToken(ctx *context.Context, client *ent.Client, tokenInput string) error {
	_, err := client.Token.Query().Where(token.And(
		token.TokenEQ(tokenInput),
		token.ExpiryGT(time.Now()),
	)).First(*ctx)

	if err != nil {
		return err
	}

	return nil
}

func ResetPassword(ctx *context.Context, tokenInput string, passwordInput string) error {
	// err := VerifyPasswordResetToken(ctx, r.client, input.Token)
	client := ent.FromContext(*ctx)
	currentUser, err := client.Token.Query().Where(token.And(
		token.TokenEQ(tokenInput),
		token.ExpiryGTE(time.Now()),
	)).QueryUser().First(*ctx)

	if err != nil {
		return err
	}

	// Hash password
	newPwd, err := utils.HashPwd(&passwordInput)
	if err != nil {
		return fmt.Errorf("error generating password")
	}

	// Save changes
	_, err = client.User.UpdateOne(currentUser).SetPassword(newPwd).Save(*ctx)
	if err != nil {
		return err
	}

	// Delete the token
	client.Token.Delete().Where(token.TokenEQ(tokenInput)).Exec(*ctx)

	return nil
}
