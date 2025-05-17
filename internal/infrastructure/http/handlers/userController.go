package handlers

import (
	"net/http"

	"github.com/DevAthhh/auth-service/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// TODO: create delete_user route

func (c *Controller) NewRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Username string `json:"username"`
		}
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := models.NewUser(request.Username, request.Email, request.Password)
		result, err := c.userService.CreateUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"uuid": result.GetID(),
		})
	}
}

func (c *Controller) NewLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := models.NewUser("", request.Email, request.Password)
		resUser, err := c.userService.FindUserByEmail(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.userService.ComparePassword(resUser, user); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid email or password",
			})
			return
		}

		tokenString, err := c.authService.GenerateToken(resUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}

func (c *Controller) RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := models.NewUser("", request.Email, request.Password)
		resUser, err := c.userService.FindUserByEmail(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.userService.ComparePassword(resUser, user); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid email or password",
			})
			return
		}

		tokenString, err := c.authService.GenerateToken(resUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"fresh-token": tokenString,
		})
	}
}

func (c *Controller) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Token    string `json:"token"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		claims, err := c.authService.ValidateToken(request.Token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if claims.Email != request.Email || claims.Username != request.Username {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "token is active",
		})
	}
}
