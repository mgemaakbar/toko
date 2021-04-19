package controller

type BuyRequest struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
