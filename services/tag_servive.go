package services

import (
	. "Teach/dto"
	"Teach/models"
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
