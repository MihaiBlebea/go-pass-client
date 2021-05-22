package catalog

import (
	"fmt"
	"time"

	"github.com/MihaiBlebea/go-pass-client/caller"
)

type service struct {
	caller caller.Service
}

type Service interface {
	GetCatalog(ID int) (Catalog, error)
	CreateCatalog(req CreateCatalogRequest) (int, error)
	DeleteCatalog(ID int) (bool, error)
	UpdateCatalog(ID int, req UpdateCatalogRequest) (bool, error)
}

type CatalogResponse struct {
	Catalog *Catalog `json:"catalog,omitempty"`
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
}

type Catalog struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Category string    `json:"category"`
	Created  time.Time `json:"created"`
}

type CreateCatalogRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type CreateCatalogResponse struct {
	CatalogID int    `json:"id"`
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
}

type RemoveCatalogResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type UpdateCatalogRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type UpdateCatalogResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func New(caller caller.Service) Service {
	return &service{caller}
}

func (s *service) GetCatalog(ID int) (Catalog, error) {
	resp := CatalogResponse{}
	endpoint := fmt.Sprintf("/catalog/%d", ID)

	err := s.caller.Get(endpoint, &resp)
	if err != nil {
		return Catalog{}, err
	}

	return *resp.Catalog, nil
}

func (s *service) CreateCatalog(req CreateCatalogRequest) (int, error) {
	resp := CreateCatalogResponse{}

	err := s.caller.Post("/catalog", &req, &resp)
	if err != nil {
		return 0, err
	}

	fmt.Println(resp)

	return resp.CatalogID, nil
}

func (s *service) DeleteCatalog(ID int) (bool, error) {
	resp := RemoveCatalogResponse{}
	endpoint := fmt.Sprintf("/catalog/%d", ID)

	err := s.caller.Delete(endpoint, &resp)
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (s *service) UpdateCatalog(ID int, req UpdateCatalogRequest) (bool, error) {
	resp := UpdateCatalogResponse{}
	endpoint := fmt.Sprintf("/catalog/%d", ID)

	err := s.caller.Update(endpoint, &req, &resp)
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}
