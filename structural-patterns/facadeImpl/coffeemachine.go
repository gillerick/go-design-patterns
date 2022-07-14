package main

import (
	"fmt"
)
//A complex Coffee API

type CoffeeMachine struct {
	beanAmount   float32 //amount in ounces of beans to use
	grinderLevel int     //the granularity of the bean grinder
	waterTemp    int     //temperature of the water to use
	waterAmt     float32 //amount of water to use
	milkAmount   float32 //amount of milk to use
	addFoam      bool    //whether to add foam or not
}

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans: ", beanAmount, "and grind level: ", grind)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans: ", c.beanAmount, "beans at ", c.grinderLevel)
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk: ", amount)
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temperature: ", temp)
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water: ", amount)
	c.waterAmt = amount
	return amount
}
