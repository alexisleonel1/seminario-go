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

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/seasons/:id",
		function: deleteSeason(s),
	})

	return list
}

func deleteSeason(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {

		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"mensaje": "Valor invalido",
			})
			return
		}

		err = s.DeleteSeason(ID)
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"mensaje": "A ocurrido un error en el servidor",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje": "se elimino",
		})
	}
}

func updateSeason(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Season
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"season": "Ingrese un valor valido",
			})
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(422, gin.H{
				"mensaje": "peticion invalida",
			})
			return
		}

		err = s.UpdateSeason(body, ID)
		if err != nil {
			c.JSON(501, gin.H{
				"mensaje": "a ocurrido un error",
			})
			return
		}
		c.JSON(200, gin.H{
			"mensaje": "peticion recibida",
		})

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
		c.JSON(201, gin.H{
			"mensaje": "peticion recibida",
		})
	}
}

func getID(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {

		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"season": "Ingrese un valor valido",
			})
			return
		}

		sn, err := s.FindByID(ID)
		if err != nil {
			c.JSON(204, gin.H{
				"mensaje": "no existe",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"season": sn,
		})
	}
}

func getAll(s HoneyService) gin.HandlerFunc {
	return func(c *gin.Context) {

		sn, err := s.FindAll()
		if err != nil {
			c.JSON(204, gin.H{
				"mensaje": "no existen archivos",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"seasons": sn,
		})
	}
}

//Resgister ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
