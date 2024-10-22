package movies

type MoviePagination struct {
	Page     int64   `json:"page"`
	PerPage  int64   `json:"per_page"`
	TotalRow int64   `json:"totoal_row"`
	Data     []Movie `json:"data"`
}
