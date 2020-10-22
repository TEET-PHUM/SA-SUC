package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Teeth/app/ent"
	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/user"
	"github.com/gin-gonic/gin"
)

// FoodmenuController defines the struct for the foodmenu controller
type FoodmenuController struct {
	client *ent.Client
	router gin.IRouter
}

// Foodmenu defines the struct for the foodmenu
type Foodmenu struct {
	FOODMENUNAME string
	FOODMENUTYPE string
	OWNERID      int
}

// CreateFoodmenu handles POST requests for adding foodmenu entities
// @Summary Create foodmenu
// @Description Create foodmenu
// @ID create-foodmenu
// @Accept   json
// @Produce  json
// @Param foodmenu body Foodmenu true "Foodmenu entity"
// @Success 200 {object} ent.Foodmenu
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foodmenus [post]
func (ctl *FoodmenuController) CreateFoodmenu(c *gin.Context) {
	obj := Foodmenu{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "foodmenu binding failed",
		})
		return
	}

	f, err := ctl.client.Foodmenu.
		Create().
		SetFoodmenuName(obj.FOODMENUNAME).
		SetFoodmenuType(obj.FOODMENUTYPE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	u, err := ctl.client.User.
		UpdateOneID(int(obj.OWNERID)).
		AddFoodmenu(f).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving edge failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetFoodmenu handles GET requests to retrieve a foodmenu entity
// @Summary Get a foodmenu entity by ID
// @Description get foodmenu by ID
// @ID get-foodmenu
// @Produce  json
// @Param id path int true "Foodmenu ID"
// @Success 200 {object} ent.Foodmenu
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foodmenus/{id} [get]
func (ctl *FoodmenuController) GetFoodmenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Foodmenu.
		Query().
		WithOwner().
		Where(foodmenu.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListFoodmenu handles request to get a list of foodmenu entities
// @Summary List foodmenu entities
// @Description list foodmenu entities
// @ID list-foodmenu
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Foodmenu
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foodmenus [get]
func (ctl *FoodmenuController) ListFoodmenu(c *gin.Context) {
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

	foodmenus, err := ctl.client.Foodmenu.
		Query().
		WithOwner().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, foodmenus)
}

// DeleteFoodmenu handles DELETE requests to delete a foodmenu entity
// @Summary Delete a foodmenu entity by ID
// @Description get foodmenu by ID
// @ID delete-foodmenu
// @Produce  json
// @Param id path int true "Foodmenu ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foodmenus/{id} [delete]
func (ctl *FoodmenuController) DeleteFoodmenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Foodmenu.
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

// UpdateFoodmenu handles PUT requests to update a foodmenu entity
// @Summary Update a foodmenu entity by ID
// @Description update foodmenu by ID
// @ID update-foodmenu
// @Accept   json
// @Produce  json
// @Param id path int true "Foodmenu ID"
// @Param foodmenu body ent.Foodmenu true "Foodmenu entity"
// @Success 200 {object} ent.Foodmenu
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foodmenus/{id} [put]
func (ctl *FoodmenuController) UpdateFoodmenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Foodmenu{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "foodmenu binding failed",
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

	f, err := ctl.client.Foodmenu.
		UpdateOneID(int(id)).
		SetFoodmenuName(obj.FOODMENUNAME).
		SetOwner(u).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update foodmenu failed",
		})
		return
	}

	c.JSON(200, f)
}

// NewFoodmenuController creates and registers handles for the foodmenu controller
func NewFoodmenuController(router gin.IRouter, client *ent.Client) *FoodmenuController {
	fc := &FoodmenuController{
		client: client,
		router: router,
	}
	fc.register()
	return fc
}

// InitFoodmenuController registers routes to the main engine
func (ctl *FoodmenuController) register() {
	foodmenus := ctl.router.Group("/foodmenus")

	foodmenus.GET("", ctl.ListFoodmenu)

	// CRUD
	foodmenus.POST("", ctl.CreateFoodmenu)
	foodmenus.GET(":id", ctl.GetFoodmenu)
	foodmenus.PUT(":id", ctl.UpdateFoodmenu)
	foodmenus.DELETE(":id", ctl.DeleteFoodmenu)
}
