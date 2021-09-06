package controllers

import (
	"ContactStore/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func NewContactController(r *gin.RouterGroup) *repositories.ContactRepository {
	cr := &repositories.ContactRepository{}
	cr.Preload()
	controller := ContactController{r: r, contactRepository: cr}
	controller.registerRoutes()

	return cr
}

type ContactController struct {
	r *gin.RouterGroup
	contactRepository *repositories.ContactRepository
}

func (c *ContactController) registerRoutes() {
	c.r.GET("", c.list)
	c.r.POST("", c.store)
	c.r.DELETE(":id", c.remove)
}

func (c *ContactController) list(ctx *gin.Context) {
	contacts := c.contactRepository.List()
	ctx.JSON(http.StatusOK, contacts)
}

func (c *ContactController) store(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.PostForm("age")
	_age, _ := strconv.ParseInt(age, 10, 0)
	phone := ctx.PostForm("phone")
	_uuid, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"message": "Failed to generate ID"})
	}
	contact := repositories.Contact{
		Id:    _uuid.String(),
		Name:  name,
		Age: int(_age),
		Phone: phone,
	}
	c.contactRepository.Store(contact)
	ctx.JSON(http.StatusOK, gin.H{"message": "Stored!"})
}

func (c *ContactController) remove(ctx *gin.Context) {
	id := ctx.Param("id")

	c.contactRepository.Remove(id)

	ctx.JSON(http.StatusOK, gin.H{"message": "Removed!"})
}
