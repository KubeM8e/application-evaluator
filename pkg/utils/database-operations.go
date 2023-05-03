package utils

import (
	"application-evaluator/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	techDB     = "Technologies"
	databaseDB = "Databases"
	serviceDB  = "ServiceDependencies"
)

// TODO: move this to an env?
var mongoURI = "mongodb://localhost:27017"

func CreateTechDataDB(techData []models.TechData) {
	ctx, client := ConnectMongoDB(mongoURI)
	for _, data := range techData {
		specificTechData := models.Technology{}
		specificTechData.Keywords = data.Keywords

		// creates database and collection
		collection := CreateDatabase(client, techDB, data.Technology)

		// inserts data into the created collection
		_, err := collection.InsertOne(ctx, specificTechData)
		if err != nil {
			log.Printf("Could not create tech data collection: %v", err)
		}
	}
}

func CreateDBDataDB(DBData []models.DBData) {
	ctx, client := ConnectMongoDB(mongoURI)
	for _, data := range DBData {
		specificDBData := models.Database{}
		specificDBData.Keywords = data.Keywords
		specificDBData.Language = data.Language

		// creates database and collection
		collection := CreateDatabase(client, databaseDB, data.Database)

		// inserts data into the created collection
		_, err := collection.InsertOne(ctx, specificDBData)
		if err != nil {
			log.Printf("Could not create DB data collection: %v", err)
		}
	}
}

func CreateServiceDataDB(serviceData []models.ServiceData) {
	ctx, client := ConnectMongoDB(mongoURI)
	for _, data := range serviceData {

		// creates database and collection
		collection := CreateDatabase(client, serviceDB, data.Language)

		// inserts data into the created collection
		_, err := collection.InsertOne(ctx, data)
		if err != nil {
			log.Printf("Could not create DB data collection: %v", err)
		}
	}
}

func ConnectMongoDB(databaseURI string) (context.Context, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURI))
	if err != nil {
		log.Printf("Could not get new mongo client: %s", err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Could not connect to MongoDB: %s", err)
	}

	//defer client.Disconnect(ctx)

	return ctx, client
}

func CreateDatabase(mongoClient *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	return mongoClient.Database(databaseName).Collection(collectionName)
}

//func ReadFromTechDB(databaseName string) map[string][]string {
//	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ctx := context.Background()
//	err = client.Connect(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Disconnect(ctx)
//
//	technologies, err := client.Database(databaseName).ListCollectionNames(ctx, bson.M{})
//	if err != nil {
//		log.Fatalf("Could not retrieve collections in the database: %v", err)
//	}
//
//	dataMap := make(map[string][]string)
//
//	for _, tech := range technologies {
//		collectionHandler := client.Database(databaseName).Collection(tech)
//
//		projection := bson.D{
//			{"_id", 0},
//		}
//		cursor, errColl := collectionHandler.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
//		if errColl != nil {
//			log.Fatalf("Could not access the collection: %v", err)
//		}
//
//		for cursor.Next(ctx) {
//			var result bson.M
//			err = cursor.Decode(&result)
//			if err != nil {
//				return nil
//			}
//
//			var keywordsSlice []string
//			keywords := result["keywords"].(primitive.A)
//			for _, keyword := range keywords {
//				keywordStr := keyword.(string)
//				keywordsSlice = append(keywordsSlice, keywordStr)
//			}
//
//			dataMap[tech] = keywordsSlice
//
//		}
//
//	}
//
//	return dataMap
//}

func ReadFromTechDB(databaseName string) []models.TechData {
	var techData models.TechData
	var techDataAll []models.TechData

	technologies, ctx, client := listCollections(databaseName)

	for _, tech := range technologies {
		collectionHandler := client.Database(databaseName).Collection(tech)

		projection := bson.D{
			{"_id", 0},
		}
		cursor, errColl := collectionHandler.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
		if errColl != nil {
			log.Fatalf("Could not access the technologies collection: %v", errColl)
		}

		for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Printf("TechDB - error in the cursor %v", err)
			}

			var keywordsSlice []string
			keywords := result["keywords"].(primitive.A)
			for _, keyword := range keywords {
				keywordStr := keyword.(string)
				keywordsSlice = append(keywordsSlice, keywordStr)
			}

			techData.Technology = tech
			techData.Keywords = keywordsSlice
			techDataAll = append(techDataAll, techData)

		}

	}

	return techDataAll
}

func ReadFromDatabaseDB(databaseName string) []models.DBData {
	var DBData models.DBData
	var DBDataAll []models.DBData

	databases, ctx, client := listCollections(databaseName)

	for _, database := range databases {
		collectionHandler := client.Database(databaseName).Collection(database)

		projection := bson.D{
			{"_id", 0},
		}
		cursor, errColl := collectionHandler.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
		if errColl != nil {
			log.Printf("Could not access the database collection: %v", errColl)
		}

		for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Printf("DatabaseDB - error in the cursor %v", err)
			}

			var keywordsSlice []string
			keywords := result["keywords"].(primitive.A)
			for _, keyword := range keywords {
				keywordStr := keyword.(string)
				keywordsSlice = append(keywordsSlice, keywordStr)
			}

			DBData.Database = database
			DBData.Language = result["language"].(string)
			DBData.Keywords = keywordsSlice
			DBDataAll = append(DBDataAll, DBData)

		}

	}

	return DBDataAll
}

func ReadFromServiceDB(databaseName string) []models.ServiceData {
	var serviceData models.ServiceData
	var serviceDataAll []models.ServiceData

	databases, ctx, client := listCollections(databaseName)

	for _, database := range databases {
		collectionHandler := client.Database(databaseName).Collection(database)

		projection := bson.D{
			{"_id", 0},
		}
		cursor, errColl := collectionHandler.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
		if errColl != nil {
			log.Printf("Could not access the database collection: %v", errColl)
		}

		for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Printf("DatabaseDB - error in the cursor %v", err)
			}

			var keywordsSlice []string
			keywords := result["keywords"].(primitive.A)
			for _, keyword := range keywords {
				keywordStr := keyword.(string)
				keywordsSlice = append(keywordsSlice, keywordStr)
			}

			serviceData.Language = result["language"].(string)
			serviceData.Keywords = keywordsSlice
			serviceDataAll = append(serviceDataAll, serviceData)

		}

	}

	return serviceDataAll
}

func listCollections(databaseName string) ([]string, context.Context, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)

	data, err := client.Database(databaseName).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Printf("Could not retrieve collections in the database: %v", err)
	}

	return data, ctx, client
}
