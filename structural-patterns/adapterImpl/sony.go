package main

import "fmt"

type SonyTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *SonyTV) getVolume() int {
	fmt.Println("SonyTV volume is", tv.currentVolume)
	return tv.currentVolume
}

func (tv *SonyTV) setVolume(vol int) {
	fmt.Println("Setting SonyTV volume to", vol)
	tv.currentVolume = vol
}

func (tv *SonyTV) getChannel() int {
	fmt.Println("SonyTV channel is", tv.currentChan)
	return tv.currentChan
}

func (tv *SonyTV) setChannel(ch int) {
	fmt.Println("Setting SonyTV channel to", ch)
	tv.currentChan = ch
}

func (tv *SonyTV) setOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("SonyTV is on")
	} else {
		fmt.Println("SonyTV is off")
	}
	tv.tvOn = tvOn
}