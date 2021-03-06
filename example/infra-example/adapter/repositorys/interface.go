package repositorys

import (
	"github.com/8treenet/freedom/example/infra-example/application/objects"
)

type GoodsInterface interface {
	Get(id int) (objects.Goods, error)
	GetAll() ([]objects.Goods, error)
	Save(*objects.Goods) error
}

type OrderInterface interface {
	Get(id int, userID int) (objects.Order, error)
	GetAll(userID int) ([]objects.Order, error)
	Create(goodsID, num, userID int) error
}
