package factory

import (
	"GoDesignPatterns/abstractfactory/internal/products"
	"fmt"
)

const (
	italianCuisine  = "italian"
	americanCuisine = "american"
)

type mealCreator interface {
	MakeDish() products.IDish
	MakeDrink() products.IDrink
}

func GetMeal(cuisineType string) (mealCreator, error) {
	switch cuisineType {
	case italianCuisine:
		return &italianMeal{}, nil
	case americanCuisine:
		return &americanMeal{}, nil
	default:
		return nil, fmt.Errorf("this restaurant doesn't have this type of food")
	}
}
