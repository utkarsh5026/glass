package handlers

import (
	"net/http"
	"server/app/models"
	"server/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	serv *services.EnrollmentService
}

func NewEnrollmentHandler(serv *services.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{serv: serv}
}

func (h *EnrollmentHandler) EnrollToCourse(c *gin.Context) {
	adminId := GetUserID(c)
	enrollId, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.serv.ApproveEnrolment(adminId, uint(enrollId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment approved"})
}

func (h *EnrollmentHandler) RejectEnrollment(c *gin.Context) {
	adminId := GetUserID(c)
	enrollID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.serv.RejectEnrolment(adminId, uint(enrollID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment rejected"})
}

func (h *EnrollmentHandler) JoinCourseByCode(c *gin.Context) {
	userId := GetUserID(c)
	var input struct {
		InvitationCode string `json:"code" binding:"required"`
		Role           string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.serv.IsValidRole(input.Role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	code := input.InvitationCode
	role := input.Role
	if err := h.serv.JoinCourseByCode(userId, code, models.Role(role)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined course"})
}

func (h *EnrollmentHandler) GetPendingEnrollments(c *gin.Context) {
	adminId := GetUserID(c)
	courseId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.serv.IsAdmin(adminId, uint(courseId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only course creator can get pending enrollments"})
		return
	}
	enrollments, err := h.serv.GetPendingEnrollments(uint(courseId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"enrollments": enrollments})
}
