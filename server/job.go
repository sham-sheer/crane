package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/crane/models"
)


func (s *Server) getJobs(c *gin.Context) {
	var jobs []model.Job
	if err := s.db.Find(&jobs).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, jobs)
}

func (s *Server) createJob(c *gin.Context) {
	var job model.Job
	if err := BindJSON(c, &job); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&job).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, job)

}

func (s *Server) getJob(c *gin.Context) {
	id := c.Param("jobID")
	var job model.Job
	if err := s.db.Find(&job, id).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, job)
	}
}

func (s *Server) updateJob(c *gin.Context) {
	var job model.Job
	if err := BindJSON(c, &job); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(job).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, job)
	}
}

func (s *Server) deleteJob(c *gin.Context) {
	jobID := c.Param("jobID")
	req := s.db.Delete(model.Job{}, "ID = ?", jobID)
	if err := req.Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, gin.H{"jobID": jobID})
	}
}