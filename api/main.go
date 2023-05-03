package main

import (
	"application-evaluator/api/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// creates a channel to signal when the /upload API is handled
	ch := make(chan bool)

	// middleware function
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			// Signal that the /file-upload API is handled
			if c.Path() == "/file-upload" {
				ch <- true
			}
			return err
		}
	})

	e.POST("/file-upload", handlers.UploadSourceCodeHandler)
	e.POST("/projects", handlers.CreateProjectHandler)
	e.GET("projects", handlers.GetAllProjectsHandler)
	e.GET("/projects/:id", handlers.GetProjectHandler)
	e.PUT("/projects/:id", handlers.UpdateProjectHandler)

	// Utility APIs
	e.POST("/data/tech", handlers.CreateTechDataHandler)
	e.POST("/data/db", handlers.CreateDBDataHandler)
	e.POST("/data/service", handlers.CreateServiceDataHandler)

	// waits for the /upload API to be handled, then call the Evaluate function
	go func() {
		<-ch
		Evaluate()
	}()

	e.Logger.Fatal(e.Start(":8082"))

}
