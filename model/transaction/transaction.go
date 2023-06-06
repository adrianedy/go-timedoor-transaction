package transaction

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	Category   	string	`bson:"category,omitempty" json:"category,omitempty"`
	Name   		string	`bson:"name,omitempty" json:"name,omitempty"`
	Description	string	`bson:"description,omitempty" json:"description,omitempty"`
	Price 		float64 `bson:"price,omitempty" json:"price,omitempty"`
}

type StatusLogs struct {
	Status         	string             	`bson:"status,omitempty" json:"status,omitempty"`
	Created_at     	primitive.DateTime 	`bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Transaction struct {
	Id  			primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	Code			int             	`bson:"code,omitempty" json:"code,omitempty"`
	User_id 		int                	`bson:"user_id,omitempty" json:"user_id,omitempty"`
	Total_price 	float64             `bson:"total_price,omitempty" json:"total_price,omitempty"`
	Status         	string             	`bson:"status,omitempty" json:"status,omitempty"`
	Products        Products			`bson:"products,omitempty" json:"products,omitempty"`
	Status_logs     StatusLogs       	`bson:"status_logs,omitempty" json:"status_logs,omitempty"`
	Created_at     	primitive.DateTime 	`bson:"created_at,omitempty" json:"created_at,omitempty"`
	Updated_at     	primitive.DateTime 	`bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}