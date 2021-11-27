package products

type Drink struct{
	Name string
	Price float32
}

type IDrink interface {
	GetName() string
	GetPrice() float32
}

func (d *Drink) GetName() string {
	return d.Name
}

func (d *Drink) GetPrice() float32{
	return d.Price
}