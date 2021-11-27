package factory

import "GoDesignPatterns/abstractfactory/internal/products"

type italianMeal struct {}

func (i *italianMeal) MakeDish() products.IDish{
	return &products.ItalianDish{
		Dish: products.Dish{
			Name: "Pasta",
			Price: 10.5,
		},
	}
}

func (i *italianMeal) MakeDrink() products.IDrink{
	return &products.ItalianDrink{
		Drink: products.Drink{
			Name:  "Apertivo",
			Price: 3,
		},
	}
}