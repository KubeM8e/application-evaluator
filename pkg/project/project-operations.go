package project

import (
	"application-evaluator/models"
	"application-evaluator/pkg/utils"
	"log"
)

var mongoURI = "mongodb://localhost:27017"

const DBName = "AppDetails"

func CreateProject(appData models.AppData) models.AppData {
	ctx, client := utils.ConnectMongoDB(mongoURI)

	appData.Id = utils.GenerateUUID()

	collection := utils.CreateDatabase(client, DBName, appData.AppName)
	_, err := collection.InsertOne(ctx, appData)
	if err != nil {
		log.Printf("Could not create collection: %s", err)
	}

	return appData
}
