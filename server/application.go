package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/crane/models"

)

func (s *Server) getApplications(c *gin.Context) {
	var applications []model.Application
	if err := s.db.Find(&applications).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, applications)
}

func (s *Server) createApplication(c *gin.Context) {
	var application model.Application
	if err := BindJSON(c, &application); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&application).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, application)

}

func (s *Server) getApplication(c *gin.Context) {
	id := c.Param("applicationID")
	var application model.Application
	if err := s.db.Find(&application, id).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, application)
	}
}

func (s *Server) updateApplication(c *gin.Context) {
	var application model.Application
	if err := BindJSON(c, &application); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(application).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, application)
	}
}

func (s *Server) deleteApplication(c *gin.Context) {
	applicationID := c.Param("applicationID")
	req := s.db.Delete(model.Application{}, "ID = ?", applicationID)
	if err := req.Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, gin.H{"applicationID": applicationID})
	}
}