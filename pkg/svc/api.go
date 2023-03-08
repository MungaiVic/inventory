package svc

type Service interface {
	GetAll() ([]ItemResponse, error)
}

type ItemResponse struct {
	ID         uint
	Name       string
	Price      uint32
	Quantity   uint8
	Reorderlvl uint8
}

