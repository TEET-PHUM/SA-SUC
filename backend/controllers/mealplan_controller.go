package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Teeth/app/ent"
	"github.com/Teeth/app/ent/mealplan"
	"github.com/Teeth/app/ent/user"
	"github.com/gin-gonic/gin"
)

// MealplanController defines the struct for the mealplan controller
type MealplanController struct {
	client *ent.Client
	router gin.IRouter
}

// Mealplan defines the struct for the mealplan
type Mealplan struct {
	MEALPLANNAME string
	OWNERID      int
}

// CreateMealplan handles POST requests for adding mealplan entities
// @Summary Create mealplan
// @Description Create mealplan
// @ID create-mealplan
// @Accept   json
// @Produce  json
// @Param mealplan body Mealplan true "Mealplan entity"
// @Success 200 {object} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans [post]
func (ctl *MealplanController) CreateMealplan(c *gin.Context) {
	obj := Mealplan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "mealplan binding failed",
		})
		return
	}

	m, err := ctl.client.Mealplan.
		Create().
		SetMealplanName(obj.MEALPLANNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	u, err := ctl.client.User.
		UpdateOneID(int(obj.OWNERID)).
		AddMealplan(m).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving edge failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetMealplan handles GET requests to retrieve a mealplan entity
// @Summary Get a mealplan entity by ID
// @Description get mealplan by ID
// @ID get-mealplan
// @Produce  json
// @Param id path int true "Mealplan ID"
// @Success 200 {object} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans/{id} [get]
func (ctl *MealplanController) GetMealplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Mealplan.
		Query().
		WithOwner().
		Where(mealplan.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListMealplan handles request to get a list of mealplan entities
// @Summary List mealplan entities
// @Description list mealplan entities
// @ID list-mealplan
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans [get]
func (ctl *MealplanController) ListMealplan(c *gin.Context) {
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

	mealplans, err := ctl.client.Mealplan.
		Query().
		WithOwner().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, mealplans)
}

// DeleteMealplan handles DELETE requests to delete a mealplan entity
// @Summary Delete a mealplan entity by ID
// @Description get mealplan by ID
// @ID delete-mealplan
// @Produce  json
// @Param id path int true "Mealplan ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans/{id} [delete]
func (ctl *MealplanController) DeleteMealplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Mealplan.
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

// UpdateMealplan handles PUT requests to update a mealplan entity
// @Summary Update a mealplan entity by ID
// @Description update mealplan by ID
// @ID update-mealplan
// @Accept   json
// @Produce  json
// @Param id path int true "Mealplan ID"
// @Param mealplan body ent.Mealplan true "Mealplan entity"
// @Success 200 {object} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans/{id} [put]
func (ctl *MealplanController) UpdateMealplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Mealplan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "mealplan binding failed",
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.OWNERID))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Mealplan.
		UpdateOneID(int(id)).
		SetMealplanName(obj.MEALPLANNAME).
		SetOwner(u).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update mealplan failed",
		})
		return
	}

	c.JSON(200, m)
}

// NewMealplanController creates and registers handles for the mealplan controller
func NewMealplanController(router gin.IRouter, client *ent.Client) *MealplanController {
	mc := &MealplanController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMealplanController registers routes to the main engine
func (ctl *MealplanController) register() {
	mealplans := ctl.router.Group("/mealplans")

	mealplans.GET("", ctl.ListMealplan)

	// CRUD
	mealplans.POST("", ctl.CreateMealplan)
	mealplans.GET(":id", ctl.GetMealplan)
	mealplans.PUT(":id", ctl.UpdateMealplan)
	mealplans.DELETE(":id", ctl.DeleteMealplan)
}
