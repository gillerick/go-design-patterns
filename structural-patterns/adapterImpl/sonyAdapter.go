package main

type sonyAdapter struct {
	stv *SonyTV
}

func (sa *sonyAdapter) turnOn() {
	sa.stv.setOnState(true)
}

func (sa *sonyAdapter) turnOff() {
	sa.stv.setOnState(false)

}

func (sa *sonyAdapter) volumeUp() int {
	vol := sa.stv.getVolume() + 1
	sa.stv.setVolume(vol)
	return vol

}

func (sa *sonyAdapter) volumeDown() int {
	vol := sa.stv.getVolume() - 1
	sa.stv.setVolume(vol)
	return vol
}

func (sa *sonyAdapter) channelUp() int {
	ch := sa.stv.getChannel() + 1
	sa.stv.setChannel(ch)
	return ch
}

func (sa *sonyAdapter) channelDown() int {
	ch := sa.stv.getChannel() - 1
	sa.stv.setChannel(ch)
	return ch
}
