package usecase

import (
	"fmt"
	"toko/customerror"
	"toko/query"

	log "github.com/sirupsen/logrus"
)

type Usecase interface {
	Buy(sku string, quantity int) error
}

type usecase struct {
	dbQuery query.DBQuery
}

func NewUsecase(productQuery query.DBQuery) Usecase {
	return &usecase{dbQuery: productQuery}
}

func (u *usecase) Buy(sku string, quantity int) error {
	if quantity <= 0 {
		return customerror.NewUserError("quantity must be > 0")
	}

	prod, err := u.dbQuery.GetProduct(sku)
	if err != nil {
		return err
	}
	if prod == nil {
		return customerror.NewUserError(fmt.Sprintf(`product with sku = %s does not exist`, sku))
	}

	if prod.Stock-quantity < 0 {
		return customerror.NewUserError(fmt.Sprintf("insufficient stock. (sku = %s, quantity = %d, version %d)", sku, quantity, prod.Version))
	}

	err = u.dbQuery.Buy(quantity, sku, prod.Version, &query.Order{
		ProductID: prod.ID,
		Quantity:  quantity,
	})
	if err != nil {
		return err
	}

	log.Info("Successfully ordered. (Quantity: %d, SKU %s, Version: %d)\n", quantity, sku, prod.Version)

	return nil
}
