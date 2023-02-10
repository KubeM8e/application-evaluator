package utils

import (
	"application-evaluator/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	databaseName = "Technologies"
)

// TODO: move this to an env?
var mongoURI = "mongodb://localhost:27017"

func ConnectMongoDB(techData []models.CreateTechDataRequest) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range techData {
		specificTechData := models.Technology{}
		specificTechData.Keywords = data.Keywords
		CreateDatabase(ctx, client, data.Technology, specificTechData)
	}

	defer client.Disconnect(ctx)
}

func CreateDatabase(ctx context.Context, mongoClient *mongo.Client, collectionName string, DBRequest models.Technology) {
	baseCollection := mongoClient.Database(databaseName).Collection(collectionName)
	_, err := baseCollection.InsertOne(ctx, DBRequest)
	if err != nil {
		log.Fatalf("Could not create collection: %v", err)
	}

}

func ReadFromDB() map[string][]string {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	technologies, err := client.Database(databaseName).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Could not retrieve collections in the database: %v", err)
	}

	techMap := make(map[string][]string)

	for _, tech := range technologies {
		collectionHandler := client.Database(databaseName).Collection(tech)

		projection := bson.D{
			{"_id", 0},
		}
		cursor, errColl := collectionHandler.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
		if errColl != nil {
			log.Fatalf("Could not access the collection: %v", err)
		}

		for cursor.Next(ctx) {
			var result bson.M
			err = cursor.Decode(&result)
			if err != nil {
				return nil
			}

			var keywordsSlice []string
			keywords := result["keywords"].(primitive.A)
			for _, keyword := range keywords {
				keywordStr := keyword.(string)
				keywordsSlice = append(keywordsSlice, keywordStr)
			}

			techMap[tech] = keywordsSlice

		}

	}

	return techMap
}
