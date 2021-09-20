package dto

type EditBookInPut struct {
	ID         int    `json:"id" valid:"Required;Min(1)"`
	Name       string `json:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `json:"state" valid:"Range(0,1)"`
}
