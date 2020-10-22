package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Teeth/app/controllers"
	"github.com/Teeth/app/ent"
	"github.com/Teeth/app/ent/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	USERNAME  string
	USEREMAIL string
}

type Mealplans struct {
	Mealplan []Mealplan
}

type Mealplan struct {
	MEALPLANNAME  string
	MEALPLANOWNER int
}

type Foodmenus struct {
	Foodmenu []Foodmenu
}

type Foodmenu struct {
	FOODMENUNAME  string
	FOODMENUTYPE  string
	FOODMENUOWNER int
}

type Tastes struct {
	Taste []Taste
}

type Taste struct {
	TASTENAME string
}

// @title SUT SA Example API Eatinghistory
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewFoodmenuController(v1, client)
	controllers.NewTasteController(v1, client)
	controllers.NewMealplanController(v1, client)
	controllers.NewEatinghistoryController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"Chao Pramong", "salt@gmail.com"},
			User{"Moung Najae", "rainbow@hotmail.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.USEREMAIL).
			SetName(u.USERNAME).
			Save(context.Background())
	}

	// Set Taste Data
	tastes := Tastes{
		Taste: []Taste{
			Taste{"Delicious"},
			Taste{"Good"},
			Taste{"Normal"},
			Taste{"Bad"},
			Taste{"Can not eat"},
		},
	}

	for _, t := range tastes.Taste {
		client.Taste.
			Create().
			SetTasteName(t.TASTENAME).
			Save(context.Background())
	}

	// Set Mealplan Data
	mealplans := Mealplans{
		Mealplan: []Mealplan{
			Mealplan{"Set A", 1},
			Mealplan{"Set B", 1},
			Mealplan{"Set Clean", 2},
			Mealplan{"Set Meat", 1},
			Mealplan{"Set Favorite", 2},
		},
	}

	for _, m := range mealplans.Mealplan {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(m.MEALPLANOWNER))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Mealplan.
			Create().
			SetMealplanName(m.MEALPLANNAME).
			SetOwner(u).
			Save(context.Background())
	}

	// Set Foodmenu Data
	foodmenus := Foodmenus{
		Foodmenu: []Foodmenu{
			Foodmenu{"Boil Egg", "Dish", 1},
			Foodmenu{"Fried Chicken", "Junk food", 1},
			Foodmenu{"Cup Noodle", "junk food", 2},
			Foodmenu{"Juice", "Drink", 1},
			Foodmenu{"Sushi", "Dish", 2},
			Foodmenu{"Cola", "Drink", 2},
		},
	}

	for _, f := range foodmenus.Foodmenu {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(f.FOODMENUOWNER))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Foodmenu.
			Create().
			SetFoodmenuName(f.FOODMENUNAME).
			SetFoodmenuType(f.FOODMENUTYPE).
			SetOwner(u).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
