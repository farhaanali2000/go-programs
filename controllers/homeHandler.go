package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/farhaan/go-programs/models"
)

var movie models.Film

func HomeHandler(c *gin.Context) {

	movs := models.GetMovies()
	// codeNo := c.Param("code")

	// c.JSON(http.StatusOK, gin.H{
	// 	"name":     movie.Name,
	// 	"genre":    movie.Genre,
	// 	"director": movie.Director,
	// 	"code No":  codeNo,
	// })
	c.IndentedJSON(http.StatusAccepted, movs)
}

func AddMovie(c *gin.Context) {

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error agaya",
		})
	}

	models.AddMovietoDB(movie)
	c.IndentedJSON(http.StatusCreated, movie)

}
