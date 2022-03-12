package main

import (
	"database/sql"
	"net/http"

	swaggerfiles "github.com/swaggo/files"

	_ "drinks-api/docs"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type cocktail struct {
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
	Ingredients string `json:"Ingredients"`
	Method      string `json:"Method"`
	Garnish     string `json:"Garnish"`
}

// @BasePath /

func main() {
	router := gin.Default()
	router.GET("/cocktails", getCocktails)
	router.GET("/cocktail/:id", getCocktailById)
	router.POST("/cocktail", addCocktail)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

// @Summary List Cocktails
// @Schemes
// @Description List IBA Cocktails
// @Accept json
// @Produce json
// @Router /cocktails [get]
func getCocktails(c *gin.Context) {

	hostname := os.Getenv("DB_HOSTNAME")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+hostname+":3306)/"+database)

	// if there is an error opening the connection, handle it
	if err != nil {
		print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name, ingredients, method, garnish FROM cocktails")
	if err != nil {
		panic(err.Error())
	}

	var list []cocktail

	for results.Next() {
		item := cocktail{0, "ooo", "ooo", "ooo", "oo"}
		err = results.Scan(&item.ID, &item.Name, &item.Ingredients, &item.Method, &item.Garnish)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, item)
	}
	c.IndentedJSON(http.StatusOK, list)
}

// @Summary get Cocktail by Id
// @Schemes
// @Description get one IBA Cocktail
// @Accept json
// @Produce json
// @Router /cocktail [get]
func getCocktailById(c *gin.Context) {

}

// @Summary Add Cocktail
// @Schemes
// @Description Add IBA Cocktail
// @Param cocktail formData cocktail true "Object Cocktail"
// @Accept application/json
// @Produce application/json
// @Router /cocktail [post]
func addCocktail(c *gin.Context) {
	var ck cocktail

	if err := c.BindJSON(&ck); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		addCocktailDB(ck)
		c.IndentedJSON(http.StatusCreated, ck)
	}
}

func addCocktailDB(cocktail cocktail) {

	hostname := os.Getenv("DB_HOSTNAME")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+hostname+":3306)/"+database)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO cocktails (name,ingredients,method,garnish) VALUES (?,?,?,?)",
		cocktail.Name, cocktail.Ingredients, cocktail.Method, cocktail.Garnish)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
