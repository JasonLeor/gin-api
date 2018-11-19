package apis

import (
	m "gindemo/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
func NewPersonHandler(c *gin.Context) {
	person := m.Person{}
	err := c.Bind(&person)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := person.NewPerson()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.String(http.StatusOK, "that's %d", id)
}
