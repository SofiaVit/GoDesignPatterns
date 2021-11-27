package factory

import "GoDesignPatterns/abstractfactory/internal/products"

type americanMeal struct {}

func (a *americanMeal) MakeDish() products.IDish{
	return &products.AmericanDish{
		Dish: products.Dish{
			Name: "Hamburger",
			Price: 7,
		},
	}
}

func (i *americanMeal) MakeDrink() products.IDrink{
	return &products.AmericanDrink{
		Drink: products.Drink{
			Name:  "Cola",
			Price: 1.5,
		},
	}
}


