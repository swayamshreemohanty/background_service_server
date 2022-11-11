package main

import (
	. "background_location_server/controller"
	. "background_location_server/db"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var 
(
	ctx context.Context
 	mongoClient *mongo.Client
	userController UserController
	err error
)

func init() {
	ctx=context.TODO()
	mongoDbConnection:=options.Client().ApplyURI("mongodb://0.0.0.0:27017")
	mongoClient,err=mongo.Connect(ctx,mongoDbConnection)
	if err!=nil {
		log.Fatal("error while connecting with mongo",err)
	}
	err=mongoClient.Ping(ctx,readpref.Primary())
	if err!=nil {
		log.Fatal("error while tring to ping mongo",err)
	}
	productMongoService:=UserMongoServiceInit(ctx,mongoClient)
	userController=	UserNewController(productMongoService)
}

func main() {
	defer mongoClient.Disconnect(ctx)
	router:=gin.Default()
	versionOne:=router.Group("/backgroundservice/v1/")
	userController.RegisterUserRoutes(versionOne)
	router.Run("0.0.0.0:8000")
}