package dto

type EmployeeCreateDto struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"` // for new employees only
	MobileNo      string `json:"mobile_no"`
	DesignationId uint   `json:"designation_id"`
	Role          int    `json:"role"`
	CompanyId     uint   `json:"company_id"`
}

type EmployeeUpdateDto struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	MobileNo      string `json:"mobile_no"`
	DesignationId uint   `json:"designation_id"`
	Role          int    `json:"role"`
	CompanyId     uint   `json:"company_id"`
	Status        string `json:"status"`
}
