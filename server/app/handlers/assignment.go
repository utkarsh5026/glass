package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/models"
	"server/app/services"
	"strconv"
)

type AssignmentHandler struct {
	serv *services.AssignmentService
}

func NewAssignmentHandler(assignmentService *services.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{
		serv: assignmentService,
	}
}

func (h *AssignmentHandler) Create(c *gin.Context) {
	var assignment models.Assignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.serv.Create(&assignment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assignment"})
		return
	}

	c.JSON(http.StatusCreated, assignment)
}

func (h *AssignmentHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	assignment, err := h.serv.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	var assignment models.Assignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment.ID = uint(id)
	if err := h.serv.Update(&assignment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update assignment"})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

func (h *AssignmentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	if err := h.serv.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete assignment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}

func (h *AssignmentHandler) GetAssignmentsForCourse(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	assignments, err := h.serv.GetAssignmentsForCourse(uint(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assignments"})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func (h *AssignmentHandler) PublishAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	if err := h.serv.Publish(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish assignment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment published successfully"})
}

func (h *AssignmentHandler) UnpublishAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	if err := h.serv.Unpublish(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unpublish assignment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment unpublished successfully"})
}

func (h *AssignmentHandler) GetUpcomingAssignments(c *gin.Context) {
	userID := c.GetUint("userID") // Assuming you have middleware to extract user ID from token
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	assignments, err := h.serv.GetUpcomingAssignments(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve upcoming assignments"})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func (h *AssignmentHandler) GetOverdueAssignments(c *gin.Context) {
	userID := c.GetUint("userID") // Assuming you have middleware to extract user ID from token

	assignments, err := h.serv.GetOverdueAssignments(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve overdue assignments"})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func (h *AssignmentHandler) GetAssignmentCompletion(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	completion, err := h.serv.GetAssignmentCompletion(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assignment completion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"completion": completion})
}
