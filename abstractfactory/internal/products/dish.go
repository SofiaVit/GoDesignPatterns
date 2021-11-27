package products

type Dish struct {
	Name  string
	Price float32
}

type IDish interface {
	GetName() string
	GetPrice() float32
}

func (d *Dish) GetName() string {
	return d.Name
}

func (d *Dish) GetPrice() float32{
	return d.Price
}