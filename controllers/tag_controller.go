package controllers

import (
	tagservice "Teach/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Edit tag
// @Produce  json
// @param tag body tagservice.EditTag true "Tag"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":500,"data":{},"msg":"ok"}"
// @Router /api/v1/edit_tags [post]
func EditTag(c *gin.Context) {
	var tagService tagservice.EditTag
	if err := c.ShouldBindJSON(&tagService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tagService.Edit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": true})
}
