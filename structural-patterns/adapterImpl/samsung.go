package main

import "fmt"

//The SamsungTV API already implements the generic television interface & therefore its methods can be called directly
type SamsungTV struct {
	vol int
	channel int
	isOn bool
}

func (st *SamsungTV) turnOn(){
	st.isOn = true
	fmt.Println("Samsung TV is now on")
}

func (st *SamsungTV) turnOff() {
	st.isOn = false
	fmt.Println("Samsung TV is now off")
}

func (st *SamsungTV) volumeUp() int{
	st.vol++
	fmt.Println("Increasing SamsungTV volume to ", st.vol)
	return st.vol
}

func (st *SamsungTV) volumeDown() int{
	st.vol--
	fmt.Println("Decreasing SamsungTV volume to ", st.vol)
	return st.vol
}

func (st *SamsungTV) channelUp() int {
	st.channel++
	fmt.Println("Increasing SamsungTV channel to ", st.channel)
	return st.channel
}

func (st *SamsungTV) channelDown() int {
	st.channel--
	fmt.Println("Decreasing SamsungTV channel to ", st.channel)
	return st.channel
}

func (st *SamsungTV) goToChannel(ch int)  {
	st.channel = ch
	fmt.Println("Setting SamsungTV channel to ", st.channel)

}