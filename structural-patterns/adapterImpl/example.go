package main

import "fmt"

func main() {
	//Create instances of the two TV types with some default values
	//The SamsungTV API already implements the generic television interface & therefore its methods can be called directly
	tv1 := &SamsungTV{
		vol:     13,
		channel: 34,
		isOn:    true,
	}

	tv2 := &SonyTV{
		currentChan:   89,
		currentVolume: 12,
		tvOn:          false,
	}

	tv1.turnOn()
	tv1.volumeUp()
	tv1.goToChannel(90)
	tv1.volumeDown()
	tv1.channelUp()
	tv1.channelDown()
	tv1.turnOff()

	fmt.Printf("SamsungTV %+t\n", tv1)
	fmt.Printf("SonyTV %+t\n", tv2)

	sonyTVAdapter := &sonyAdapter{stv: tv2}

	sonyTVAdapter.turnOn()
	sonyTVAdapter.volumeUp()
	sonyTVAdapter.channelUp()
	sonyTVAdapter.channelDown()
	sonyTVAdapter.volumeDown()
	sonyTVAdapter.channelUp()
	sonyTVAdapter.channelDown()
	sonyTVAdapter.turnOff()

	fmt.Printf("SONY Adapter %+t\n", sonyTVAdapter)

}
