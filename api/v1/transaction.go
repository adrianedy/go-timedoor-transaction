package v1

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func PostTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
		// coll := database.Collection("transaction")
		// title := "Back to the Future"
		// var result bson.M
		// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
		// if err == mongo.ErrNoDocuments {
		// 	fmt.Printf("No document was found with the title %s\n", title)
		// 	return
		// }
		// if err != nil {
		// 	panic(err)
		// }
		// jsonData, err := json.MarshalIndent(result, "", "    ")
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("%s\n", jsonData)
	}
}