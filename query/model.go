package query

type Product struct {
	ID      int
	SKU     string
	Stock   int
	Version uint64
}

type Order struct {
	ID        int
	ProductID int
	Quantity  int
}
