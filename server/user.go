package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/crane/models"

)

func (s *Server) getUsers(c *gin.Context) {
	var users []User
	if err := s.db.Find(&users).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) createUser(c *gin.Context) {
	var user User
	if err := BindJSON(c, &user); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, user)

}

func (s *Server) getUser(c *gin.Context) {
	id := c.Param("userID")
	var user User
	if err := s.db.Find(&user, id).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (s *Server) updateUser(c *gin.Context) {
	var user User
	if err := BindJSON(c, &user); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(user).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (s *Server) deleteUser(c *gin.Context) {
	userID := c.Param("userID")
	req := s.db.Delete(User{}, "ID = ?", userID)
	if err := req.Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, gin.H{"userId": userID})
	}
}
