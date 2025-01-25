package dto

type DesignationCreateDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DesgFor     int    `json:"desg_for" `
	Description string `json:"description" `
	ParentId    int    `json:"parent_id"`
}

type DesignationUpdateDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DesgFor     int    `json:"desg_for" `
	Description string `json:"description" `
	ParentId    int    `json:"parent_id"`
	Status      int    `json:"status"`
}
