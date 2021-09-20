package services

import (
	. "Teach/dto"
	"Teach/models"
	"github.com/unknwon/com"
)

func EditTag(t *EditTagInPut) error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}
	return models.EditTag(t.ID, data)
}
func AddTag(t *AddTagInPut) error {
	t.CreatedBy = "admin"
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}
func GetTage(t *GetTagInPut) (int, []models.BlogTag) {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if t.Name != "" {
		maps["name"] = t.Name
	}
	state := -1
	if t.State != "" {
		state = com.StrTo(t.State).MustInt()
	}
	if state >= 0 {
		maps["state"] = state
	}
	count, _ := models.GetTagTotal(maps)
	t.PageNum = (t.PageNum - 1) * t.PageSize
	list, _ := models.GetTags(t.PageNum, t.PageSize, maps)
	return count, list
}
