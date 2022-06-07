package main

import "fmt"

func main() {
	var builder = newNotificationBuilder()

	builder.setTitle("New Notification")
	builder.setIcon("icon.png")
	builder.setImage("image.png")
	builder.setPriority(10)
	builder.setMessage("This is a basic notification")
	builder.setNotType("alert")

	notification, err := builder.Build()
	if err != nil{
		fmt.Println("Error creating the notification")
	} else{
		fmt.Printf("Notification: %+v ", notification)
	}
}
