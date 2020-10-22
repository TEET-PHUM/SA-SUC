package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Teeth/app/ent"
	"github.com/Teeth/app/ent/taste"
	"github.com/gin-gonic/gin"
)

// TasteController defines the struct for the taste controller
type TasteController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTaste handles POST requests for adding taste entities
// @Summary Create taste
// @Description Create taste
// @ID create-taste
// @Accept   json
// @Produce  json
// @Param taste body ent.Taste true "Taste entity"
// @Success 200 {object} ent.Taste
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tastes [post]
func (ctl *TasteController) CreateTaste(c *gin.Context) {
	obj := ent.Taste{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "taste binding failed",
		})
		return
	}

	t, err := ctl.client.Taste.
		Create().
		SetTasteName(obj.TasteName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTaste handles GET requests to retrieve a taste entity
// @Summary Get a taste entity by ID
// @Description get taste by ID
// @ID get-taste
// @Produce  json
// @Param id path int true "Taste ID"
// @Success 200 {object} ent.Taste
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tastes/{id} [get]
func (ctl *TasteController) GetTaste(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := ctl.client.Taste.
		Query().
		Where(taste.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListTaste handles request to get a list of taste entities
// @Summary List taste entities
// @Description list taste entities
// @ID list-taste
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Taste
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tastes [get]
func (ctl *TasteController) ListTaste(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	tastes, err := ctl.client.Taste.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tastes)
}

// DeleteTaste handles DELETE requests to delete a taste entity
// @Summary Delete a taste entity by ID
// @Description get taste by ID
// @ID delete-taste
// @Produce  json
// @Param id path int true "Taste ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tastes/{id} [delete]
func (ctl *TasteController) DeleteTaste(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Taste.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTaste handles PUT requests to update a taste entity
// @Summary Update a taste entity by ID
// @Description update taste by ID
// @ID update-taste
// @Accept   json
// @Produce  json
// @Param id path int true "Taste ID"
// @Param taste body ent.Taste true "Taste entity"
// @Success 200 {object} ent.Taste
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tastes/{id} [put]
func (ctl *TasteController) UpdateTaste(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Taste{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "taste binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	t, err := ctl.client.Taste.
		UpdateOneID(int(id)).
		SetTasteName(obj.TasteName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, t)
}

// NewTasteController creates and registers handles for the taste controller
func NewTasteController(router gin.IRouter, client *ent.Client) *TasteController {
	tc := &TasteController{
		client: client,
		router: router,
	}
	tc.register()
	return tc
}

// InitTasteController registers routes to the main engine
func (ctl *TasteController) register() {
	tastes := ctl.router.Group("/tastes")

	tastes.GET("", ctl.ListTaste)

	// CRUD
	tastes.POST("", ctl.CreateTaste)
	tastes.GET(":id", ctl.GetTaste)
	tastes.PUT(":id", ctl.UpdateTaste)
	tastes.DELETE(":id", ctl.DeleteTaste)
}
