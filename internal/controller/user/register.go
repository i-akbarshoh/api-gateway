package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/i-akbarshoh/api-gateway/internal/entity"
	"github.com/i-akbarshoh/api-gateway/internal/pkg/jwt"
	"context"
)

type Controller struct {
	useCase Usecase
}

func NewController(us Usecase) *Controller {
	return &Controller{useCase: us}
}

func (con *Controller) SignUp(c *gin.Context) {
	var (
		body entity.User
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}
	body.ID = uuid.NewString()
	ctx := context.Background()
	err := con.useCase.SignUp(ctx, body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot register user, " + err.Error(),
		})
		return
	}
	tokens, err := jwt.GenerateNewTokens(body.ID, map[string]string{"role": body.Role})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot generate tokens, " + err.Error(),
		})
		return
	}

	c.JSON(200, tokens)
}

func (con *Controller) Login(c *gin.Context) {
	var (
		body entity.User
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}
	ctx := context.Background()
	if err := con.useCase.Login(ctx, entity.Login{
		Email: body.Email,
		Password: body.Password,
	}); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot login, " + err.Error(),
		})
		return
	}

	tokens, err := jwt.GenerateNewTokens(body.ID, map[string]string{"role": body.Role})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot generate tokens, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access": tokens.Access,
		"expire": tokens.AccExpire,
		"refresh": tokens.Refresh,
	})
}