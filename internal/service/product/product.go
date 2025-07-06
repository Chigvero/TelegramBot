package product

import (
	"github.com/chigvero/TelegramBot/internal/model"
)

type Service struct {
}

var allProduct = []model.Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

func NewService() *Service {
	return &Service{}
}

func (*Service) List() []model.Product {
	return allProduct
}
