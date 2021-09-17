package services

import (
	"Teach/models"
)

type EditTag struct {
	ID         int    `json:"id" valid:"Required;Min(1)"`
	Name       string `json:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `json:"state" valid:"Range(0,1)"`
}

func (t *EditTag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}
	return models.EditTag(t.ID, data)
}
