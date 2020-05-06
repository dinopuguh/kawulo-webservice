package response

import "github.com/gemcook/pagination-go"

type PaginationResponse struct {
	Pages      pagination.Pages `json:"pages"`
	TotalCount int              `json:"total_count"`
	TotalPages int              `json:"total_pages"`
}
