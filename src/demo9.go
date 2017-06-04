package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}
type TV struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("connect:", pc.name)
}
func (tv TV) Connect() {
	fmt.Println("connect:", tv.name)
}
func main() {
	var a USB
	a = PhoneConnecter{"phoneconnect"}
	a.Connect()
	Disconnect(a)
	// tv := TV{"TVCONNECT"}
	// var a USB
	// a = USB(tv)
	// a.Connect()
	// pc := PhoneConnecter{"phone"}
	// var a USB
	// a = USB(pc)
	// a.Connect()
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("disconnected:", v.name)
	default:
		fmt.Println("unknow.")
	}
}
