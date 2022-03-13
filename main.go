package main

import (
	"drinks-api/config"
	"drinks-api/middleware"
	"drinks-api/models"
	"drinks-api/utils"
	"net/http"

	swaggerfiles "github.com/swaggo/files"

	_ "drinks-api/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securitydefinitions.apikey AuthToken
// @in header
// @name Authorization
// @BasePath /
func main() {
	router := gin.Default()

	// configure firebase
	firebaseAuth := config.SetupFirebase()
	router.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	//router.Use(middleware.AuthMiddleware)

	router.GET("/cocktails", middleware.AuthMiddleware, getCocktails)
	router.GET("/cocktail/:id", middleware.AuthMiddleware, getCocktailById)
	router.DELETE("/cocktail/:id", middleware.AuthMiddleware, deleteCocktailById)
	router.PUT("/cocktail", middleware.AuthMiddleware, addCocktail)
	router.POST("/login", getLogin)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

// @Summary Login
// @Schemes
// @Description Request Login
// @Param  user body models.User true "User auth data"
// @Accept  json
// @Produce  json
// @Router /login [post]
func getLogin(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		user.ReturnSecureToken = "true"
		result := utils.ValidateUser(user)
		c.IndentedJSON(http.StatusOK, result)
	}

}

// @Summary List Cocktails
// @Schemes
// @Description List IBA Cocktails
// @Accept json
// @Produce json
// @security AuthToken
// @Router /cocktails [get]
func getCocktails(c *gin.Context) {
	cocktails := models.GetAllCocktails()

	if cocktails == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, cocktails)
	}
}

// @Summary get Cocktail by Id
// @Schemes
// @Description get one IBA Cocktail
// @Accept json
// @Produce json
// @Param id path int true "Id cocktail"
// @security AuthToken
// @Router /cocktail/{id} [get]
func getCocktailById(c *gin.Context) {

	id := c.Param("id")

	ck := models.GetCocktail(id)

	if ck == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, ck)
	}
}

// @Summary Add Cocktail
// @Schemes
// @Description Add IBA Cocktail
// @Param cocktail formData models.Cocktail true "Object Cocktail"
// @security AuthToken
// @Accept application/json
// @Produce application/json
// @Router /cocktail [put]
func addCocktail(c *gin.Context) {
	var ck models.Cocktail

	if err := c.BindJSON(&ck); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddCocktailDB(ck)
		c.IndentedJSON(http.StatusCreated, ck)
	}
}

// @Summary Delete Cocktail by Id
// @Schemes
// @Description Delete One IBA Cocktail
// @Accept json
// @Produce json
// @Param id path int true "Id cocktail"
// @security AuthToken
// @Router /cocktail/{id} [delete]
func deleteCocktailById(c *gin.Context) {

	id := c.Param("id")

	ck := models.DeleteCocktail(id)

	if ck == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, ck)
	}
}
