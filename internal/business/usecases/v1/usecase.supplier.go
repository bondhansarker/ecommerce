package v1

import (
	"context"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"net/http"
)

type supplierUseCase struct {
	repo V1Domains.SupplierRepository
}

// NewSupplierUseCase creates a new instance of SupplierUseCase
func NewSupplierUseCase(repo V1Domains.SupplierRepository) V1Domains.SupplierUseCase {
	return &supplierUseCase{
		repo: repo,
	}
}

func (s *supplierUseCase) Create(ctx context.Context, inputDomain V1Domains.SupplierDomain) (outputDomain V1Domains.SupplierDomain, statusCode int, err error) {
	// Additional business logic/validation can be added here before calling the repository

	// Example: You might want to ensure that the email is unique before creating a new supplier
	// if exists, _ := s.repo.IsEmailUnique(ctx, inputDomain.Email); !exists {
	// 	return V1Domains.SupplierDomain{}, http.StatusConflict, fmt.Errorf("email is not unique")
	// }

	// Call the repository to create the supplier record
	outputDomain, err = s.repo.CreateRecord(ctx, inputDomain)
	if err != nil {
		return V1Domains.SupplierDomain{}, http.StatusInternalServerError, err
	}

	return outputDomain, http.StatusCreated, nil
}

func (s *supplierUseCase) GetByID(ctx context.Context, id int64) (outputDomain V1Domains.SupplierDomain, statusCode int, err error) {
	// Call the repository to get the supplier record by ID
	outputDomain, err = s.repo.GetRecordByID(ctx, id)
	if err != nil {
		status, err := helpers.HandleCommonRepositoryError(err)
		return V1Domains.SupplierDomain{}, status, err
	}
	return outputDomain, http.StatusOK, nil
}

func (s *supplierUseCase) GetList(ctx context.Context) (outputDomains []V1Domains.SupplierDomain, statusCode int, err error) {
	// Call the repository to get the list of suppliers
	outputDomains, err = s.repo.GetRecords(ctx)
	if err != nil {
		return []V1Domains.SupplierDomain{}, http.StatusInternalServerError, err
	}
	return outputDomains, http.StatusOK, nil
}

func (s *supplierUseCase) Update(ctx context.Context, inputDomain V1Domains.SupplierDomain) (statusCode int, err error) {
	// Call the repository to update the supplier record
	err = s.repo.UpdateRecord(ctx, inputDomain)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}

func (s *supplierUseCase) Delete(ctx context.Context, id int64) (statusCode int, err error) {
	// Call the repository to delete the supplier record
	err = s.repo.DeleteRecordByID(ctx, id)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}
