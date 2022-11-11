package db

import (
	. "background_location_server/helper"
	. "background_location_server/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserDBContext struct {
	ctx context.Context
	mongoclient *mongo.Client
}

type UserMongoService interface{
	InsertUserDataToDB(newuser *UserModel) (*UserModel,error) 
	UpdateUserDataOnDB(userLocation *UserLocationModel) (*UserLocationModel,*string,error)

}

func UserMongoServiceInit(ctx context.Context,mongoclient *mongo.Client)  UserMongoService{
	return &UserDBContext{
		ctx: ctx,
		mongoclient: mongoclient,
	}
}

func (userMongoContext *UserDBContext)InsertUserDataToDB(newUser *UserModel) (*UserModel,error)  {
	dbref:= userMongoContext.mongoclient.Database(string(DatabasePath.DATABASE)).Collection(string(DatabasePath.USERS))

	// myOption:= options.FindOne()
	// myOption.SetSort(bson.M{"$natural":-1})
	
	// var product UserModel
	// dbref.FindOne(userMongoContext.ctx,bson.M{},myOption).Decode(&product)
	// lastElementId,err:= strconv.Atoi(product.UId)
	// if err!=nil {
	// 	//set the lastElementid to 0, if there is error in find last element, mean the collection is empty
	// 	lastElementId=0
	// }
	// newUser.Id=strconv.Itoa(lastElementId+1)

	result,err:=dbref.InsertOne(userMongoContext.ctx,newUser)
	if err!=nil {
		return nil,err
	}else if (newUser.UId!=result.InsertedID){
		return nil,errors.New("Unable to add the user to database")
	}else{
		return newUser,nil
	}

}

func (userMongoContext *UserDBContext)UpdateUserDataOnDB(userLocation *UserLocationModel) (*UserLocationModel,*string,error)  {
	dbref:= userMongoContext.mongoclient.Database(string(DatabasePath.DATABASE)).Collection(string(DatabasePath.USERS))
	
	filter:=bson.M{"_id":userLocation.UId}

	update:=bson.M{"$set":userLocation}
	
	result,err:=dbref.UpdateOne(userMongoContext.ctx,filter,update)
	if err!=nil {
		return nil,nil,err
	}else if result.MatchedCount !=1 {
		return nil,nil,errors.New("No matched user found for update")
	}else{
		var user UserModel
		err=dbref.FindOne(userMongoContext.ctx,filter).Decode(&user)
		if err!=nil {
			return nil,nil,err
		}
		return userLocation,&user.UserName,nil
	}

}