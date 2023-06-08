package transaction

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CollectionName string = "transactions"

type Status int

const (
	Created        Status = 1
	WaitingPayment Status = 2
	Paid           Status = 3
	Finish         Status = 4
	Cancel         Status = 4
)

func (s Status) String() string {
	switch s {
	case Created:
		return "created"
	case WaitingPayment:
		return "waiting_payment"
	case Paid:
		return "paid"
	case Finish:
		return "finish"
	default:
		return "cancel"
	}
}

type Products struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64            `bson:"price,omitempty" json:"price,omitempty"`
	Amount      int                `bson:"amount,omitempty" json:"amount,omitempty"`
	Images      []Images           `bson:"images,omitempty" json:"images,omitempty"`
}

type Images struct {
	Url string `bson:"url,omitempty" json:"url,omitempty"`
}

type StatusLogs struct {
	Status     Status             `bson:"status,omitempty" json:"status,omitempty"`
	Created_at primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Code        int                `bson:"code,omitempty" json:"code,omitempty"`
	User_id     int                `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Total_price float64            `bson:"total_price,omitempty" json:"total_price,omitempty"`
	Status      Status             `bson:"status,omitempty" json:"status,omitempty"`
	Products    Products           `bson:"products,omitempty" json:"products,omitempty"`
	Status_logs StatusLogs         `bson:"status_logs,omitempty" json:"status_logs,omitempty"`
	Created_at  primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Updated_at  primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
