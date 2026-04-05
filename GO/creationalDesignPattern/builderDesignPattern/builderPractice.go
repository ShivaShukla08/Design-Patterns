package main

import (
	"fmt"
	"strings"
)

type BurgerMeal struct {
	BunType   string
	Patty     string
	HasCheese bool
	Toppings  []string
	Side      *string
	Drink     *string
}

type BurgerBuilder struct{
	bunType string
	patty   string
	
	// optional
	hasCheese bool
	toppings  []string
	side      *string
	drink     *string
}

func NewBurgerBuilder(bunType string, patty string) *BurgerBuilder {
	return &BurgerBuilder{
		bunType: bunType,
		patty:   patty,
		toppings: []string{},
	}
}

func (b *BurgerBuilder) WithCheese(hasCheese bool) *BurgerBuilder {
	b.hasCheese = hasCheese
	return b
}

func (b *BurgerBuilder) WithToppings(toppings []string) *BurgerBuilder {
	b.toppings = toppings
	return b
}

func (b *BurgerBuilder) WithSide(side string) *BurgerBuilder {
	b.side = &side
	return b
}

func (b *BurgerBuilder) WithDrink(drink string) *BurgerBuilder {
	b.drink = &drink
	return b
}

func (b *BurgerBuilder) Build() (*BurgerMeal, error) {
	if b.bunType == "" || b.patty == "" {
		return nil, fmt.Errorf("bunType and patty are mandatory")
	}

	// Return the immutable object
	return &BurgerMeal{
		BunType:   b.bunType,
		Patty:     b.patty,
		HasCheese: b.hasCheese,
		Toppings:  b.toppings,
		Side:      b.side,
		Drink:     b.drink,
	}, nil
}

func (b *BurgerMeal) Describe() string {
	description := fmt.Sprintf("🍔 Burger with %s bun and %s patty", b.BunType, b.Patty)

	if b.HasCheese {
		description += ", with cheese"
	} else {
		description += ", no cheese"
	}

	if len(b.Toppings) > 0 {
		description += fmt.Sprintf(", toppings: %s", strings.Join(b.Toppings, ", "))
	}

	if b.Side != nil {
		description += fmt.Sprintf(", side: %s", *b.Side)
	}

	if b.Drink != nil {
		description += fmt.Sprintf(", drink: %s", *b.Drink)
	}

	return description
}

func main() {
	// Plain burger
	// plainBurger, _ := NewBurgerBuilder("wheat", "veg").Build()
	// fmt.Println(plainBurger)

	// Burger with cheese
	// cheeseBurger, _ := NewBurgerBuilder("wheat", "veg").WithCheese(true).Build()
	// fmt.Println(cheeseBurger)

	// Fully loaded burger
	loadedBurger, _ := NewBurgerBuilder("multigrain", "chicken").
		WithToppings([]string{"lettuce", "onion", "jalapeno"}).
		WithSide("fries").
		WithDrink("coke").
		WithCheese(true).
		Build()

	// fmt.Println(loadedBurger)
	fmt.Println(loadedBurger.Describe())
}