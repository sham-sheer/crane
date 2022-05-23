package server

import (
	"net/http"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	model "github.com/crane/models"
)

type Server struct {
	db *gorm.DB
}

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	if err := db.AutoMigrate(&model.User{}, &model.Job{}, &model.Recruiter{}, &model.Application{}); err != nil {
		panic(err)
	}

	return db
}

// NewServer creates a new instance of a Server.
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

// RegisterRouter registers a router onto the Server.
func (s *Server) RegisterRouter(router *gin.Engine) {
	router.GET("/ping", s.ping)

	router.GET("/user", s.getUsers)
	router.POST("/user", s.createUser)
	router.GET("/user/:userID", s.getUser)
	router.PUT("/user/:userID", s.updateUser)
	router.DELETE("/user/:userID", s.deleteUser)

	router.GET("/job", s.getJobs)
	router.POST("/job", s.createJob)
	router.GET("/job/:jobID", s.getJob)
	router.PUT("/job/:jobID", s.updateJob)
	router.DELETE("/job/:jobID", s.deleteJob)

	router.GET("/application", s.getApplications)
	router.POST("/application", s.createApplication)
	router.GET("/application/:applicationID", s.getApplication)
	router.PUT("/application/:applicationID", s.updateApplication)
	router.DELETE("/application/:applicationID", s.deleteApplication)
}

func (s *Server) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "server ready to go!"})
}

func BindJSON(c *gin.Context, obj interface{}) (err error) {
	if err = c.ShouldBindWith(obj, binding.JSON); err != nil {
		return err
	}
	return
}