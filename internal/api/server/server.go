package server

import (
	"net/http"
	"parqueadero-back/internal/api/model"
	"parqueadero-back/internal/api/service"
	"parqueadero-back/internal/simulation"

	"github.com/gin-gonic/gin"
)

type Server struct {
	// Router  *mux.Router
	Router    *gin.Engine
	service   *service.Service
	socketURL string
}

func New(s *service.Service, socketURL string) Server {
	server := &Server{service: s, socketURL: socketURL}

	engine := gin.Default()
	apiRouter := engine.Group("/api")
	apiRouter.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	apiRouter.GET("/municipios", server.getMunicipios)

	apiRouter.GET("/parqueaderos", server.GetParqueaderos)
	apiRouter.POST("/parqueaderos", server.CreateParqueaderos)

	apiRouter.GET("/buses", server.GetBuses)
	apiRouter.POST("/buses", server.CreateBuses)

	apiRouter.GET("/programas", server.GetProgramaciones)
	apiRouter.POST("/programas", server.CreateProgramacion)

	apiRouter.POST("/simulate", server.Simulate)

	server.Router = engine
	return *server
}

func (s *Server) getMunicipios(c *gin.Context) {
	municipios, err := s.service.GetMunicipios()
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, municipios)
}

func (s *Server) GetParqueaderos(c *gin.Context) {
	parqueaderos, err := s.service.GetParqueaderos()
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, parqueaderos)
}

func (s *Server) CreateParqueaderos(c *gin.Context) {
	var parqueadero model.Parqueadero
	if err := c.ShouldBindJSON(&parqueadero); err != nil {
		responseBadRequest(c, err)
		return
	}

	parqueaderos, err := s.service.CreateParqueadero(parqueadero)
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, parqueaderos)
}

func (s *Server) GetBuses(c *gin.Context) {
	buses, err := s.service.GetBuses()
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, buses)
}

func (s *Server) CreateBuses(c *gin.Context) {
	var bus model.Bus
	if err := c.ShouldBindJSON(&bus); err != nil {
		responseBadRequest(c, err)
		return
	}

	buses, err := s.service.CreateBus(bus)
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, buses)
}

func (s *Server) GetProgramaciones(c *gin.Context) {
	programas, err := s.service.GetProgramaciones()
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, programas)
}

func (s *Server) CreateProgramacion(c *gin.Context) {
	var programa model.CreateProgramacion
	if err := c.ShouldBindJSON(&programa); err != nil {
		responseBadRequest(c, err)
		return
	}

	buses, err := s.service.CreateProgramacion(programa)
	if err != nil {
		responseError(c, err)
		return
	}

	responseJson(c, buses)
}

func (s *Server) Simulate(c *gin.Context) {
	go simulation.Simulate(s.socketURL)
	c.Status(http.StatusOK)
}

func responseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func responseError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func responseJson(c *gin.Context, obj any) {
	c.JSON(http.StatusAccepted, obj)
}
