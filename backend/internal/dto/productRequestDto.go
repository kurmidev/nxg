package dto

type ProductCreateDto struct {
	ID               uint                        `json:"id"`
	Name             string                      `json:"name"`
	ProjectAtrribute []ProductAttributeCreateDto `json:"project_attribute"`
	Description      string                      `json:"description"`
}

type ProductUpdateDto struct {
	ID               uint                        `json:"id"`
	Name             string                      `json:"name"`
	ProjectAtrribute []ProductAttributeCreateDto `json:"project_attribute"`
	Description      string                      `json:"description"`
}

type ProductAttributeCreateDto struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Status int    `json:"status"`
}
