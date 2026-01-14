package user

import (
	"TOC/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Module struct {
	r *gin.Engine

	uc domain.UserUsecase
}

func NewModule(r *gin.Engine, uc domain.UserUsecase) *Module {
	m := &Module{
		r:  r,
		uc: uc,
	}

	r.POST("/toc/v1/user/create", m.createUser)

	return m
}

func (m Module) createUser(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
	}

	_, err = m.uc.Create(c, &user)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
