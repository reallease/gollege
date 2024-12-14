package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/reallease/gollege/data/questions/enem/year2023"
)

// List question by year and language
//
//	@Summary		List questions
//	@Description	List question from an specific year and language
//
//	@Tags			ENEM
//	@Produce		application/json
//
//	@Param			year	path	string	true	"Year"
//	@Param			lang	path	string	true	"Language"
//
//	@Router			/enem/{year}/{lang} [get]
func GetByYear(ctx *gin.Context) {
	year := ctx.Param("year")
	lang := ctx.Param("lang")

	switch year {
	case "2023":
		switch lang {
		case "languages":
			ctx.JSON(200, year2023.Languages)
		case "natural_ciencies":
			ctx.JSON(200, year2023.Natural_sciences)
		}
	}
}

// Get file by year and id
//
//	@Summary		Get file
//	@Description	Get file from an specific year and id
//
//	@Tags			ENEM
//	@Produce		*/*
//
//	@Param			year	path	string	true	"Year"
//	@Param			id		path	string	true	"ID"
//
//	@Router			/enem/files/{year}/{id} [get]
func GetFile(ctx *gin.Context) {
	year := ctx.Param("year")
	id := ctx.Param("id")

	root := "data/files/enem"
	path := fmt.Sprintf("%s/%s", root, year)
	question := fmt.Sprintf("%s/%s", path, id)

	ctx.File(question)
}

//
// * EnemQuestions[26] -> Question{}
//

// http://localhost:8080/swagger/index.html

func AllRouters(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/v1")

	v1.GET("/enem/:year/:lang", GetByYear)
	v1.GET("/enem/files/:year/:id", GetFile)

	return v1
}
