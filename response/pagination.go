package response

import "github.com/gemcook/pagination-go"

type PaginationResponse struct {
	Pages      pagination.Pages `json:"pages"`
	Previous   string           `json:"previous"`
	Next       string           `json:"next"`
	TotalCount int              `json:"total_count"`
	TotalPages int              `json:"total_pages"`
}
