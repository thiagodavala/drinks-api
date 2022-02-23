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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Method      string `json:"method"`
	Garnish     string `json:"garnish"`
}

// @BasePath /

func main() {
	router := gin.Default()
	router.GET("/cocktails", getCocktails)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}

// IBA Cocktails API List
// @Summary cocktails API List
// @Schemes
// @Description List IBA Cocktails
// @Tags example
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
