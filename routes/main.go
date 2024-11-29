package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/reallease/gollege/data/questions/enem/year2023"
)

// List ENEM questions from one year.
//
//	@Summary		List questions
//	@Description	Get question
//	@Tags			ENEM
//	@Produces		json
//	@Param			year	path	string	true	"Year"
//	@Router			/enem/{year} [get]

func GetByYear(ctx *gin.Context) {
	year := ctx.Param("year")
	lang := ctx.Param("lang")

	//  var question questions.Question

	// root := "./data/question2/enem"
	// path := fmt.Sprintf("%s/%s", root, year)
	// question_path := fmt.Sprintf("%s/question.json", path)

	// questions_file, _ := os.ReadFile(question_path)

	// ctx.Data(200, gin.MIMEJSON, questions_file)

	switch year {
	case "2023":
		switch lang {
		case "languages":
			ctx.JSON(200, year2023.Languages)
		case "natural_ciencies":
			ctx.JSON(200, year2023.Natural_sciences)
		}
		// case "2022":
		// 	ctx.JSON(200, year2022.Questions)

		// case "2021":
		// 	ctx.JSON(200, year2021.Questions)

		// case "2020":
		// 	ctx.JSON(200, year2020.Humans_sciences)

		// case "2019":
		// 	ctx.JSON(200, year2019.Questions)

		// default:
		// 	ctx.JSON(200, gin.H{})
		//
	}
}

//
// * EnemQuestions[26] -> Question{}
//

// http://localhost:8080/swagger/index.html

func AllRouters(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/v1")

	v1.GET("/enem/:year/:lang", GetByYear)

	return v1
}
