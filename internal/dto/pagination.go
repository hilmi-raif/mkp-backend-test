package dto

type Pagination struct {
	Page  int `json:"page" validate:"min=1,max=21474836"`
	Limit int `json:"limit" validate:"min=1,max=100"`
}

type PageMeta struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	HasNext bool `json:"has_next"`
}
