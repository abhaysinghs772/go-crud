package router

import (
	"net/http"

	db "github.com/abhaysinghs772/go-crud/db"
	"github.com/gin-gonic/gin"
)

func postMovie(ctx *gin.Context) {
	var movie db.Movie
	err := ctx.Bind(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"movie": res,
	})
}

func rootMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "api is up and running",
	})
}

func Initrouter() *gin.Engine {
	r := gin.Default()
	// r.GET("/movies", getMovies)
	// r.GET("/movies/:id", getMovie)
	r.GET("/", rootMethod)
	r.POST("/movies", postMovie)
	// r.PUT("/movies/:id", updateMovie)
	// r.DELETE("/movies/:id", deleteMovie)
	return r
}
