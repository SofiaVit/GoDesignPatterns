package main

import (
	"GoDesignPatterns/abstractfactory/internal/factory"
	"GoDesignPatterns/abstractfactory/internal/products"
	"fmt"
)

const (
	americanMeal = "american"
	italianMeal  = "italian"
)

func main(){
	americanMealFactory, _ := factory.GetMeal(americanMeal)
	italianMealFactory, _ := factory.GetMeal(italianMeal)

	americanDish := americanMealFactory.MakeDish()
	americanDrink := americanMealFactory.MakeDrink()

	italianDish := italianMealFactory.MakeDish()
	italianDrink := italianMealFactory.MakeDrink()

	printMealDetails(americanDish, americanDrink)
	printMealDetails(italianDish, italianDrink)

}

func printMealDetails(dish products.IDish, drink products.IDrink){
	fmt.Printf("Meal containts %s and %s. The price is %.2f\n", dish.GetName(), drink.GetName(), dish.GetPrice()+drink.GetPrice())
}

