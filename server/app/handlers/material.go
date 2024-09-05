package handlers

import (
	"net/http"
	"server/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	serv *services.MaterialService
}

func NewMaterialHandler(serv *services.MaterialService) *MaterialHandler {
	return &MaterialHandler{serv: serv}
}

func (h *MaterialHandler) GetMaterial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param(MaterialIDKey), 10, 64)
	if err != nil {
		HandleBadRequest(c, InvalidMaterialID)
		return
	}

	material, err := h.serv.GetMaterial(uint(id))
	if err != nil {
		SendError(err, c)
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *MaterialHandler) DeleteMaterial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param(MaterialIDKey), 10, 64)
	if err != nil {
		HandleBadRequest(c, InvalidMaterialID)
		return
	}

	if err := h.serv.DeleteMaterial(uint(id)); err != nil {
		SendError(err, c)
		return
	}

	HandleDeleted(c, "Material deleted successfully")
}

func (h *MaterialHandler) CreateMaterial(c *gin.Context) {
}

func (h *MaterialHandler) UpdateMaterial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param(MaterialIDKey), 10, 64)
	if err != nil {
		HandleBadRequest(c, InvalidMaterialID)
		return
	}

	var updateData struct {
		Title       string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
		RemoveFiles []uint `form:"removeFiles[]" binding:"required"`
	}

	if err := c.ShouldBind(&updateData); err != nil {
		HandleBadRequest(c, err.Error())
		return
	}

	material, err := h.serv.GetMaterial(uint(id))
	if err != nil {
		SendError(err, c)
		return
	}

	material.Title = updateData.Title
	material.Description = updateData.Description

	form, err := c.MultipartForm()
	if err != nil {
		HandleBadRequest(c, FailedToParseMultipartForm)
		return
	}

	filesToAdd := form.File["files"]
	filesToRemove := updateData.RemoveFiles

	if err := h.serv.UpdateMaterial(material, filesToAdd, filesToRemove); err != nil {
		SendError(err, c)
		return
	}

	c.JSON(http.StatusOK, material)

}
