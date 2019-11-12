package frutils

import (
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) int {
	idUser, _ := c.Get("ID")
	return ToInt(idUser.(string))
}

func GetToken(c *gin.Context) string {
	return strings.Trim(strings.TrimSuffix(c.Request.Header.Get("Authorization"), "\n"), `"`)
}

func GetRequestIP(c *gin.Context) string {
	return c.Request.RemoteAddr
}

func GetRequestBody(c *gin.Context) string {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return "error reading request body"
	}
	return string(body)
}
