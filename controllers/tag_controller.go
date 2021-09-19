package controllers

import (
	. "Teach/dto"
	"Teach/models"
	"Teach/pkg/response"
	tagservice "Teach/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add article tag
// @Produce  json
// @param tag body AddTagInPut true "AddTagInPut"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/add_tag [post]
func AddTag(c *gin.Context) {
	appG := response.Gin{C: c}
	var addTagInPut AddTagInPut
	if err := c.ShouldBindJSON(&addTagInPut); err != nil {
		appG.Response(http.StatusBadRequest, response.INVALID_PARAMS, nil)
		return
	}
	exists, err := models.ExistTagByName(addTagInPut.Name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, response.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, response.ERROR_EXIST_TAG, nil)
		return
	}
	err = tagservice.AddTag(&addTagInPut)
	if err != nil {
		appG.Response(http.StatusInternalServerError, response.ERROR_ADD_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, response.SUCCESS, nil)
}

// @Summary Edit tag
// @Produce  json
// @param tag body EditTagInPut true "EditTagInPut"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/edit_tag [post]
func EditTag(c *gin.Context) {
	appG := response.Gin{C: c}
	var editTagInPut EditTagInPut
	if err := c.ShouldBindJSON(&editTagInPut); err != nil {
		appG.Response(http.StatusBadRequest, response.INVALID_PARAMS, nil)
		return
	}
	err := tagservice.EditTag(&editTagInPut)
	if err != nil {
		appG.Response(http.StatusInternalServerError, response.ERROR_EDIT_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, response.SUCCESS, nil)
}
