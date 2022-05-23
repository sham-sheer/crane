package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/crane/models"
)


func (s *Server) getRecruiters(c *gin.Context) {
	var recruiters []model.Recruiter
	if err := s.db.Find(&recruiters).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, recruiters)
}

func (s *Server) createRecruiter(c *gin.Context) {
	var recruiter model.Recruiter
	if err := BindJSON(c, &recruiter); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&recruiter).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, recruiter)

}

func (s *Server) getRecruiter(c *gin.Context) {
	id := c.Param("recruiterID")
	var recruiter model.Recruiter
	if err := s.db.Find(&recruiter, id).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, recruiter)
	}
}

func (s *Server) updateRecruiter(c *gin.Context) {
	var recruiter model.Recruiter
	if err := BindJSON(c, &recruiter); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(recruiter).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, recruiter)
	}
}

func (s *Server) deleteRecruiter(c *gin.Context) {
	recruiterID := c.Param("recruiterID")
	req := s.db.Delete(model.Recruiter{}, "ID = ?", recruiterID)
	if err := req.Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, gin.H{"recruiterID": recruiterID})
	}
}