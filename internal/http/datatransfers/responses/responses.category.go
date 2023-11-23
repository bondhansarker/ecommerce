package responses

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type CategoryResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ParentID  *int64    `json:"parent_id"`
	Sequence  uint      `json:"sequence"`
	StatusID  int8      `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryTreeNode struct {
	ID       int64               `json:"id"`
	Name     string              `json:"name"`
	Children []*CategoryTreeNode `json:"children"`
}

func (categoryResponse *CategoryResponse) ToV1Domain() V1Domains.CategoryDomain {
	return V1Domains.CategoryDomain{
		ID:        categoryResponse.ID,
		Name:      categoryResponse.Name,
		ParentID:  categoryResponse.ParentID,
		Sequence:  categoryResponse.Sequence,
		StatusID:  helpers.IntegerToBoolean[categoryResponse.StatusID],
		CreatedAt: categoryResponse.CreatedAt,
	}
}

func FromV1DomainToCategoryResponse(categoryDomain V1Domains.CategoryDomain) CategoryResponse {
	return CategoryResponse{
		ID:        categoryDomain.ID,
		Name:      categoryDomain.Name,
		ParentID:  categoryDomain.ParentID,
		Sequence:  categoryDomain.Sequence,
		StatusID:  helpers.BooleanToInteger[*categoryDomain.StatusID],
		CreatedAt: categoryDomain.CreatedAt,
	}
}

func FromV1DomainToTreeNode(categoryDomain V1Domains.CategoryDomain) *CategoryTreeNode {
	return &CategoryTreeNode{
		ID:       categoryDomain.ID,
		Name:     categoryDomain.Name,
		Children: []*CategoryTreeNode{},
	}
}

func ToTreeListOfCategories(domains []V1Domains.CategoryDomain) []*CategoryTreeNode {
	// Create a map to store categories by their ID for quick lookup
	categoryMap := make(map[int64]*CategoryTreeNode)

	// Initialize the categoryTree
	var categoryTree = make([]*CategoryTreeNode, 0)

	for _, category := range domains {
		categoryMap[category.ID] = FromV1DomainToTreeNode(category)
	}

	// Iterate through the categories to build the tree
	for _, category := range domains {
		// If the category has a parent, add it to the parent's children list
		if category.ParentID != nil {
			parent, exists := categoryMap[*category.ParentID]
			if exists {
				parent.Children = append(parent.Children, categoryMap[category.ID])
			}
		} else {
			// If the category has no parent, it's a top-level category
			categoryTree = append(categoryTree, categoryMap[category.ID])
		}
	}
	return categoryTree
}

func ToResponseListOfCategories(domains []V1Domains.CategoryDomain) []CategoryResponse {
	var result = make([]CategoryResponse, 0)

	for _, val := range domains {
		result = append(result, FromV1DomainToCategoryResponse(val))
	}

	return result
}
