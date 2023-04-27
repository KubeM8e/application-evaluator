package main

import (
	"application-evaluator/api/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/upload", handlers.UploadSourceCodeHandler)
	//e.POST("/files", handlers.SourceFileEvalHandler)
	//e.POST("/techs", handlers.TechnologyEvalHandler)

	e.POST("/data/tech", handlers.CreateTechDataHandler)
	e.POST("/data/db", handlers.CreateDBDataHandler)
	e.POST("/data/service", handlers.CreateServiceDataHandler)

	e.Logger.Fatal(e.Start(":8082"))

}
