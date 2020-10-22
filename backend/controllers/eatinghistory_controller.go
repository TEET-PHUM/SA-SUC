package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Teeth/app/ent"
	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/mealplan"
	"github.com/Teeth/app/ent/taste"
	"github.com/Teeth/app/ent/user"
	"github.com/gin-gonic/gin"
)

// EatinghistoryController defines the struct for the eatinghistory controller
type EatinghistoryController struct {
	client *ent.Client
	router gin.IRouter
}

// Eatinghistory defines the struct for the eatinghistory
type Eatinghistory struct {
	MealplanID int
	FoodmenuID int
	TasteID    int
	UserID     int
	AddedTime  string
}

// CreateEatinghistory handles POST requests for adding eatinghistory entities
// @Summary Create eatinghistory
// @Description Create eatinghistory
// @ID create-eatinghistory
// @Accept   json
// @Produce  json
// @Param eatinghistory body Eatinghistory true "Eatinghistory entity"
// @Success 200 {object} ent.Eatinghistory
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /eatinghistorys [post]
func (ctl *EatinghistoryController) CreateEatinghistory(c *gin.Context) {
	obj := Eatinghistory{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "eatinghistory binding failed",
		})
		return
	}

	f, err := ctl.client.Foodmenu.
		Query().
		Where(foodmenu.IDEQ(int(obj.FoodmenuID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "foodmenu not found",
		})
		return
	}

	m, err := ctl.client.Mealplan.
		Query().
		Where(mealplan.IDEQ(int(obj.MealplanID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "mealplan not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	t, err := ctl.client.Taste.
		Query().
		Where(taste.IDEQ(int(obj.TasteID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "taste not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.AddedTime)

	et, err := ctl.client.Eatinghistory.
		Create().
		SetUser(u).
		SetMealplan(m).
		SetFoodmenu(f).
		SetTaste(t).
		SetAddedTime(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, et)
}

// ListEatinghistory handles request to get a list of eatinghistory entities
// @Summary List eatinghistory entities
// @Description list eatinghistory entities
// @ID list-eatinghistory
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Eatinghistory
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /eatinghistorys [get]
func (ctl *EatinghistoryController) ListEatinghistory(c *gin.Context) {
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

	eatinghistorys, err := ctl.client.Eatinghistory.
		Query().
		WithUser().
		WithFoodmenu().
		WithMealplan().
		WithTaste().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, eatinghistorys)
}

// DeleteEatinghistory handles DELETE requests to delete a eatinghistory entity
// @Summary Delete a eatinghistory entity by ID
// @Description get eatinghistory by ID
// @ID delete-eatinghistory
// @Produce  json
// @Param id path int true "Eatinghistory ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /eatinghistorys/{id} [delete]
func (ctl *EatinghistoryController) DeleteEatinghistory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Eatinghistory.
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

// NewEatinghistoryController creates and registers handles for the eatinghistory controller
func NewEatinghistoryController(router gin.IRouter, client *ent.Client) *EatinghistoryController {
	etc := &EatinghistoryController{
		client: client,
		router: router,
	}
	etc.register()
	return etc
}

// InitEatinghistoryController registers routes to the main engine
func (ctl *EatinghistoryController) register() {
	eatinghistorys := ctl.router.Group("/eatinghistorys")

	eatinghistorys.GET("", ctl.ListEatinghistory)

	// CRUD
	eatinghistorys.POST("", ctl.CreateEatinghistory)
	eatinghistorys.DELETE(":id", ctl.DeleteEatinghistory)
}
