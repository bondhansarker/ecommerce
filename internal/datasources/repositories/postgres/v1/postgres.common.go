package v1

import (
	"errors"
)

var NoDataFound = errors.New("no data found")
var RowsPerPageZeroError = errors.New("per_page value must be grater than zero")

type Pagination struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
}

func GetPaginationInfoData(totalCount int64, CurrentPage, limitPerPage int) (*Pagination, error) {
	// Calculate the total number of pages
	if limitPerPage == 0 {
		return nil, RowsPerPageZeroError
	}
	totalPages := totalCount / int64(limitPerPage)
	if totalCount%int64(limitPerPage) > 0 {
		totalPages++
	}

	hasNextPage := int64(CurrentPage) < totalPages
	pagination := &Pagination{
		Total:       totalCount,
		PerPage:     limitPerPage,
		CurrentPage: CurrentPage,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
	}

	return pagination, nil
}
