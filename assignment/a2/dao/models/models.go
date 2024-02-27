package models

type Item struct {
	Id string
	Code string
	Quantity int
	Description string
}

type Order struct {
	Id string
	CostumerName string
	OrderedAt string
	Items item[]
}


