package httpd

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunwenlei/golang/go_webapi/lib"
)

//Ping for root of /api return http 200
func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}

//Getusers call lib.GetPersons() and return all persons json type
func Getusers(c *gin.Context) {
	users := lib.GetPersons()
	c.JSON(http.StatusOK, *users)
}

//Getuser get person data by personCd in parameter[personcd]
func Getuser(c *gin.Context) {
	user, err := lib.GetPerson(c.Param("personcd"))

	switch {
	case err == sql.ErrNoRows:
		c.Status(http.StatusNoContent)
		return
	case err != nil:
		c.Status(http.StatusInternalServerError)
		return

	}
	c.JSON(http.StatusOK, *user)
}

//Createuser create a new person with request json and return new person code
func Createuser(c *gin.Context) {
	var newuser lib.MstPersonInfo

	c.BindJSON(&newuser)

	newPersonCD, err := lib.CreatePerson(&newuser)

	if err == nil {
		c.JSON(http.StatusOK, map[string]string{"personCD": newPersonCD})
	} else {
		c.Status(http.StatusInternalServerError)
	}

	return
}

//Updateuser update person data by request json and parameter [personcd]
func Updateuser(c *gin.Context) {
	var user lib.MstPersonInfo
	c.BindJSON(&user)

	err := lib.UpdatePerson(&user, c.Param("personcd"))

	if err == nil {
		result, _ := lib.GetPerson(c.Param("personcd"))

		c.JSON(http.StatusOK, *result)
	} else {
		if err == sql.ErrNoRows {
			c.Status(http.StatusNoContent)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

}

//Deleteuser delete person data with parameter [personcd]
func Deleteuser(c *gin.Context) {
	cnt, err := lib.DeletePerson(c.Param("personcd"))

	if err == nil {
		if cnt > 0 {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusNoContent)
		}

	} else {
		c.Status(http.StatusInternalServerError)
	}
}
