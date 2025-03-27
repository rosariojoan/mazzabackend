package auth

import (
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/membersignuptoken"
	"mazza/ent/generated/predicate"
	"mazza/inits"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func VerifyMemberInvitationToken(ctx *gin.Context) {
	var input struct {
		Token string
		Email string
	}

	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	filter := []predicate.MemberSignupToken{
		membersignuptoken.Token(input.Token),
		membersignuptoken.AlreadyUsed(false),
		membersignuptoken.ExpiresAtGTE(time.Now()),
		membersignuptoken.Or(
			membersignuptoken.EmailIsNil(),
			membersignuptoken.EmailEQ(input.Email),
		),
	}

	currentToken, err := inits.Client.MemberSignupToken.Query().
		Where(filter...).
		First(ctx)
	if err != nil {
		fmt.Println("VerifyMemberSignupToken update err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}
	_, _ = currentToken.Update().AddNumberAccessed(1).Save(ctx)

	var companies []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	err = currentToken.QueryCompany().Select(company.FieldID, company.FieldName).Scan(ctx, &companies)
	if err != nil {
		fmt.Println("VerifyMemberSignupToken get company err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	token := generated.MemberSignupToken{
		ID:     currentToken.ID,
		Name:   currentToken.Name,
		Token:  currentToken.Token,
		Avatar: currentToken.Avatar,
		Role:   currentToken.Role,
		Note:   currentToken.Note,
	}

	currentCompany := companies[0]
	ctx.JSON(http.StatusCreated, gin.H{"token": token, "company": currentCompany})
}
