package responses

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type BrandResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	StatusID  int8      `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (brandResponse *BrandResponse) ToV1Domain() V1Domains.BrandDomain {
	return V1Domains.BrandDomain{
		ID:        brandResponse.ID,
		Name:      brandResponse.Name,
		StatusID:  helpers.IntegerToBoolean[brandResponse.StatusID],
		CreatedAt: brandResponse.CreatedAt,
	}
}

func FromV1DomainToBrandResponse(brandDomain V1Domains.BrandDomain) BrandResponse {
	return BrandResponse{
		ID:        brandDomain.ID,
		Name:      brandDomain.Name,
		StatusID:  helpers.BooleanToInteger[*brandDomain.StatusID],
		CreatedAt: brandDomain.CreatedAt,
	}
}

func ToResponseListOfBrands(domains []V1Domains.BrandDomain) []BrandResponse {
	var result = make([]BrandResponse, 0)

	for _, val := range domains {
		result = append(result, FromV1DomainToBrandResponse(val))
	}

	return result
}
