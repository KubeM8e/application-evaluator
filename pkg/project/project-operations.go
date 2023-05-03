package project

import (
	"application-evaluator/models"
	"application-evaluator/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const (
	EvaluatorDB    = "Evaluator"
	ProjectDetails = "ProjectDetails"
)

func CreateProject(projectData models.ProjectData) models.ProjectData {
	ctx, client := utils.ConnectMongoDB(utils.MongoURI)

	projectData.ProjectId = utils.GenerateUUID()

	collection := utils.CreateDatabase(client, EvaluatorDB, ProjectDetails)
	_, err := collection.InsertOne(ctx, projectData)
	if err != nil {
		log.Printf("Could not insert document: %s", err)
	}

	return projectData
}

func GetAllProjects() []models.ProjectData {
	ctx, client := utils.ConnectMongoDB(utils.MongoURI)
	collection := utils.CreateDatabase(client, EvaluatorDB, ProjectDetails)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Could not find all documents: %s", err)
	}
	var allProjectDetails []models.ProjectData
	if err = cursor.All(ctx, &allProjectDetails); err != nil {
		log.Printf("Could not iterate all project details: %s", err)
	}

	return allProjectDetails
}

func GetProject(projectId string) models.ProjectData {
	ctx, client := utils.ConnectMongoDB(utils.MongoURI)
	collection := client.Database(EvaluatorDB).Collection(ProjectDetails)

	projectData := models.ProjectData{}

	err := collection.FindOne(ctx, bson.M{"projectId": projectId}).Decode(&projectData)
	if err != nil {
		log.Printf("Could not find collection: %s", err)
	}

	return projectData
}

func UpdateProject(projectId string, data models.ProjectData) bool {
	var isError bool
	ctx, client := utils.ConnectMongoDB(utils.MongoURI)
	collection := client.Database(EvaluatorDB).Collection(ProjectDetails)

	//dataBson, errM := bson.Marshal(data)
	//if errM != nil {
	//	log.Printf("Could not marshal to bson: %s", errM)
	//}

	//result, err := collection.UpdateOne(
	//	ctx,
	//	bson.M{"projectId": projectId},
	//	bson.D{
	//		{"$set", bson.D{dataBson}},
	//	},
	//)

	filter := bson.M{"projectId": projectId}

	response := collection.FindOneAndReplace(ctx, filter, data)

	if response.Err() != nil {
		isError = true
		log.Printf("Could not update the document: %s", response.Err())
	} else {
		isError = false
	}

	return isError

}
