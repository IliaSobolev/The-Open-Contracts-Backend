package codeBlock

import (
	"TOC/pkg/domain"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type Module struct {
	r *gin.Engine

	cb domain.CodeBlockUsecase
}

func NewModule(r *gin.Engine, cb domain.CodeBlockUsecase) *Module {
	m := &Module{
		r:  r,
		cb: cb,
	}

	r.POST("/toc/v1/codeblock/create", m.createCodeBlock)
	r.GET("/toc/v1/codeblock/list", m.getCodeBlocks)
	r.GET("/toc/v1/codeblock/:id", m.getCodeBlock)
	return m
}

func (m *Module) createCodeBlock(c *gin.Context) {
	var codeBlock domain.CodeBlockDTO
	err := c.ShouldBindJSON(&codeBlock)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
		return
	}

	err = m.cb.Create(c, codeBlock)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create codeBlock"})
		return
	}

	c.JSON(200, gin.H{
		"message": "create codeBlock",
	})
}
func (m *Module) getCodeBlock(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
		return
	}
	res, err := m.cb.Get(context.Background(), id)
	if err != nil {
		if errors.Is(err, domain.ErrCodeBlockNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "code block not found"})
			return
		}
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get codeBlock"})
		return
	}
	c.JSON(200, res)
}

func (m *Module) getCodeBlocks(c *gin.Context) {
	res, err := m.cb.List(context.Background())
	if err != nil {
		if errors.Is(err, domain.ErrCodeBlockNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "code blocks not found"})
			return
		}
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get codeBlocks"})
		return
	}
	c.JSON(200, res)
}
