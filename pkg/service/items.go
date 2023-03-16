package service

import (
	"github.com/MungaiVic/inventory/pkg/models"
	"github.com/MungaiVic/inventory/pkg/repository"
)

type SVC struct {
	dao	repository.ItemRepository
}

func New (dao repository.ItemRepository) ItemService {
	return &SVC{dao}
}

func (s *SVC) GetItems()([]models.Item, error){
	items := s.dao.GetAll()
	return items, nil
}
func (s *SVC) GetItemByID(){}
func (s *SVC) CreateItem(){}
func (s *SVC) UpdateItem(){}
func (s *SVC) DeleteItem(){}
