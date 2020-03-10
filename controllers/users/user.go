package users

import (
	"net/http"
	"strconv"

	"github.com/gaurav1der/bookstore_user-api/domain/users"
	"github.com/gaurav1der/bookstore_user-api/services"
	"github.com/gaurav1der/bookstore_user-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO error handler
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO Json error hanlder
	// 	fmt.Println(err.Error())
	// 	return
	// }

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		// restErr := errors.RestErr{
		// 	Message: "Invalid Json Body",
		// 	Status:  http.StatusBadRequest,
		// 	Error:   "bad_request",
		// }
		//fmt.Println(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO create user error
		return
	}
	// fmt.Println(user)
	//fmt.Println(string(bytes))
	//fmt.Println(err)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User ID should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Implement me")

// }
