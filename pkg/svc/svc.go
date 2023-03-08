package svc

import (
	repo "github.com/MungaiVic/inventory/pkg/repository"
)

type SVC struct{
	dao	repo.ItemRepository
}

func New (dao repo.ItemRepository) Service {
	return &SVC{dao}
}

func (s *SVC) GetAll() ([]ItemResponse, error){
	items, err := s.dao.GetAll()
	return ConvertItemModelToItemService(items), err
}
