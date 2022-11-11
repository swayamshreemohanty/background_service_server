package model

type UserModel struct {
	UId       string  `json:"id" bson:"_id" form:"id"`
	UserName  string  `json:"user_name" bson:"user_name" form:"user_name" binding:"required"`
}


type UserLocationModel struct {
	UId       string  `json:"id" bson:"_id" form:"id"`
	Longitude string `json:"longitude" bson:"longitude" form:"longitude" binding:"required"`
	Latitude  string `json:"latitude" bson:"latitude" form:"latitude" binding:"required"`
	Altitude  string `json:"altitude" bson:"altitude" form:"altitude" binding:"required"`
	Heading   string `json:"heading" bson:"heading" form:"heading" binding:"required"`
	Speed     string `json:"speed" bson:"speed" form:"speed" binding:"required"`
}