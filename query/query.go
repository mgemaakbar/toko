package query

import (
	"database/sql"
	"fmt"
	"toko/customerror"
)

type DBQuery interface {
	GetProduct(SKU string) (*Product, error)
	Buy(deltaStock int, SKU string, version uint64, o *Order) error
}

type dbQuery struct {
	conn *sql.DB
}

func NewDBQuery(conn *sql.DB) DBQuery {
	return &dbQuery{conn: conn}
}

func (p *dbQuery) Buy(quantity int, SKU string, version uint64, o *Order) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return customerror.NewInternalError(err.Error())
	}
	err = p.updateProductStock(tx, quantity, SKU, version)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = p.createOrder(tx, o)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return customerror.NewInternalError(err.Error())
	}

	return nil
}

func (p *dbQuery) GetProduct(SKU string) (*Product, error) {
	ret := Product{}
	err := p.conn.QueryRow(`SELECT id, sku, stock, version FROM product WHERE sku = $1`, SKU).Scan(&ret.ID, &ret.SKU, &ret.Stock, &ret.Version)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil { // unexpected error
		return nil, customerror.NewInternalError(err.Error())
	}
	return &ret, nil
}

func (p *dbQuery) updateProductStock(txConn *sql.Tx, deltaStock int, SKU string, version uint64) error {
	q := `UPDATE product SET stock = stock - $1, version = version + 1 WHERE sku = $2 AND version = $3`
	res, err := txConn.Exec(q, deltaStock, SKU, version)
	if err != nil {
		return customerror.NewInternalError(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return customerror.NewInternalError(err.Error())
	}

	if rows == 0 {
		return fmt.Errorf(`race condition happened while updating the stock. version not found. (Quantity: %d, SKU %s, Version: %d)`, deltaStock, SKU, version) // assuming the SKU exists, but the version doesn't
	}

	return nil
}

func (p *dbQuery) createOrder(txConn *sql.Tx, o *Order) error {
	_, err := txConn.Exec(`INSERT INTO orders (product_id, quantity) VALUES ($1, $2)`, o.ProductID, o.Quantity)
	if err != nil {
		return customerror.NewInternalError(err.Error())
	}
	return nil
}
