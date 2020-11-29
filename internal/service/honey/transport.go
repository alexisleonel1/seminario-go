package honey

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

//NewHTTPTransport ...
func NewHTTPTransport(s HoneyService) HTTPService {
	enpoints := makeEndpoints(s)
	return httpService{enpoints}
}

func makeEndpoints(s HoneyService) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/seasons",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/seasons/:id",
		function: getID(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/seasons",
		function: addSeason(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/seasons/:id",
		function: updateSeason(s),
	})

	return list
}

func updateSeason(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Season
		ID := c.Param("id")
		id, err := strconv.Atoi(ID)
		if err == nil {
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(422, gin.H{
					"mensaje": "peticion invalida",
				})
				return
			}
			err := s.UpdateSeason(body, id)
			if err != nil {
				c.JSON(501, gin.H{
					"mensaje": "a ocurrido un error",
				})
				return
			}
			c.JSON(200, gin.H{
				"mensaje": "peticion recibida",
			})
		} else {
			c.JSON(http.StatusNotImplemented, gin.H{
				"season": "Ingrese un valor valido",
			})
		}

	}
}

func addSeason(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Season
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(422, gin.H{
				"mensaje": "peticion invalida",
			})
			return
		}
		s.AddSeason(body)
		c.JSON(200, gin.H{
			"mensaje": "peticion recibida",
		})
	}
}

func getID(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")
		id, err := strconv.Atoi(ID)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"season": s.FindByID(id),
			})
		} else {
			c.JSON(http.StatusNotImplemented, gin.H{
				"season": "Ingrese un valor valido",
			})
		}
	}
}

func getAll(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"seasons": s.FindAll(),
		})
	}
}

//Resgister ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
