package helpers

import "errors"

type PaginationResult struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
}

func GetPaginationInfoData(totalCount int64, CurrentPage, limitPerPage int) (*PaginationResult, error) {
	// Calculate the total number of pages
	if limitPerPage == 0 {
		return nil, errors.New("limitPerPage value must be grater than zero")
	}
	totalPages := totalCount / int64(limitPerPage)
	if totalCount%int64(limitPerPage) > 0 {
		totalPages++
	}

	hasNextPage := int64(CurrentPage) < totalPages
	pagination := &PaginationResult{
		Total:       totalCount,
		PerPage:     limitPerPage,
		CurrentPage: CurrentPage,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
	}

	return pagination, nil
}
