package dto

type EditTagInPut struct {
	ID         int    `json:"id" valid:"Required;Min(1)"`
	Name       string `json:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `json:"state" valid:"Range(0,1)"`
}
type AddTagInPut struct {
	Name       string `json:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `json:"state" valid:"Range(0,1)"`
	CreatedBy  string
}
type GetTagInPut struct {
	Name     string `form:"name" valid:"Required;MaxSize(100)"`
	State    string `form:"state" valid:"Range(0,1)"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
type DeleteTagInPut struct {
	ID int `json:"id" valid:"Required;Min(1)"`
}
