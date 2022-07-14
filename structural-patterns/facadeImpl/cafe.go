package main

import "fmt"

//A high level simplified Coffee API

func makeAmericano(size float32){
	fmt.Println("\nMaking an Americano\n----------------------")
	americano := &CoffeeMachine{}

	beansAmount := (size / 8.0) * 5.0
	americano.startCoffee(beansAmount, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("Americano is great")
}

func makeLatte(size float32, foam bool){
	fmt.Println("\nMaking a Latte\n--------------------------")

	latte := &CoffeeMachine{}

	beansAmount := (size / 8.0) * 5.0
	latte.startCoffee(beansAmount, 4)
	latte.grindBeans()
	latte.useHotWater(size)

	milk := (size /8.0) * 2.0
	latte.useMilk(milk)
	latte.endCoffee()
	fmt.Println("Latte is ready")
}
