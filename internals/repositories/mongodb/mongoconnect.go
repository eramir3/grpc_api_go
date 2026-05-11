package mongodb

import (
	"context"
	"grpcapi/pkg/utils"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateMongoClient() (*mongo.Client, error) {
	ctx := context.Background()
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://estebanramirezcc_db_user:79kopShZilmmn0MA@cluster0.scjbqo6.mongodb.net/?appName=Cluster0").
		SetBSONOptions(&options.BSONOptions{ObjectIDAsHexString: true})

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to connect to database")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to ping database")
	}

	log.Println("Connected to MongoDB")
	return client, nil
}
