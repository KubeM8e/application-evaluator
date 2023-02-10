package main

import (
	"application-evaluator/api/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/upload", handlers.UploadSourceCodeHandler)
	e.POST("/data/techs", handlers.CreateTechDataHandler)
	e.POST("files", handlers.SourceFileEvalHandler)
	e.POST("techs", handlers.TechnologyEvalHandler)

	e.Logger.Fatal(e.Start(":8082"))

}
