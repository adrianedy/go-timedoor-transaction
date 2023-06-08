package transactions

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/backend-timedoor/go-transaction-module/database"
	transactionModel "github.com/backend-timedoor/go-transaction-module/models/transaction"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Meta struct {
	Page PageInfo `json:"page"`
}

type PageInfo struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type TransactionResponse struct {
	ID          primitive.ObjectID        `bson:"_id,omitempty" json:"_id,omitempty"`
	Code        int                       `bson:"code,omitempty" json:"code,omitempty"`
	User_id     int                       `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Total_price float64                   `bson:"total_price,omitempty" json:"total_price,omitempty"`
	Status      string                    `bson:"status,omitempty" json:"status,omitempty"`
	Products    transactionModel.Products `bson:"products,omitempty" json:"products,omitempty"`
}

type PaginatedResponse struct {
	Data  interface{} `json:"data"`
	Meta  Meta        `json:"meta"`
	Links Links       `json:"links"`
}

func GetTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		collection := database.Collection(transactionModel.CollectionName)
		defer collection.Database().Client().Disconnect(context.Background())

		pageStr := c.QueryParam("page")
		limitStr := c.QueryParam("limit")

		pageNum, err := strconv.ParseInt(pageStr, 10, 64)
		limitNum, err := strconv.ParseInt(limitStr, 10, 64)

		// Set default value
		if pageNum <= 0 {
			pageNum = 1
		}
		if limitNum <= 0 {
			limitNum = 10
		}

		skip := (pageNum - 1) * limitNum

		cursor, err := collection.Find(context.Background(), bson.M{}, &options.FindOptions{
			Skip:  &skip,
			Limit: &limitNum,
		})
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "Failed to retrieve transactions")
		}
		defer cursor.Close(context.Background())

		transactions := []TransactionResponse{}
		for cursor.Next(context.Background()) {
			var transaction transactionModel.Transaction
			err := cursor.Decode(&transaction)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, "Failed to decode transactions")
			}

			TransactionResponse := ToResponse(transaction)
			transactions = append(transactions, TransactionResponse)
		}

		totalCount, err := collection.CountDocuments(context.Background(), bson.M{})
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "Failed to count transactions")
		}

		lastPage := int(math.Ceil(float64(totalCount) / float64(limitNum)))

		from := int(skip + 1)
		to := int(skip + limitNum)
		if int64(to) > totalCount {
			to = int(totalCount)
		}

		pageInfo := PageInfo{
			CurrentPage: int(pageNum),
			From:        from,
			LastPage:    lastPage,
			PerPage:     int(limitNum),
			To:          to,
			Total:       int(totalCount),
		}

		links := Links{
			First: generatePageLink(1, int(limitNum)),
			Last:  generatePageLink(lastPage, int(limitNum)),
			Next:  generateNextPageLink(int(pageNum), lastPage, int(limitNum)),
			Prev:  generatePrevPageLink(int(pageNum), int(limitNum)),
		}

		response := PaginatedResponse{
			Data:  transactions,
			Meta:  Meta{Page: pageInfo},
			Links: links,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func ToResponse(t transactionModel.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:          t.ID,
		Code:        t.Code,
		User_id:     t.User_id,
		Total_price: t.Total_price,
		Status:      t.Status.String(),
		Products:    t.Products,
	}
}

func generatePageLink(pageNum, limitNum int) string {
	baseURL := viper.GetString("url")
	return fmt.Sprintf("%s?page=%d&limit=%d", baseURL, pageNum, limitNum)
}

func generateNextPageLink(currentPage, lastPage, limitNum int) string {
	if currentPage >= lastPage {
		return ""
	}
	nextPage := currentPage + 1
	return generatePageLink(nextPage, limitNum)
}

func generatePrevPageLink(currentPage, limitNum int) string {
	if currentPage <= 1 {
		return ""
	}
	prevPage := currentPage - 1
	return generatePageLink(prevPage, limitNum)
}
