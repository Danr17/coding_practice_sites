package clock

// The option with the clock define as an int

import "fmt"

//Clock defines the Clock
type Clock int

//New instantiate a Clock
func New(hour, minute int) Clock {

	min := (hour*60 + minute) % (24 * 60)

	if min < 0 {
		min += (24 * 60)
	}

	return Clock(min)

}

func (c Clock) String() string {

	hour := int(c) / 60
	minute := int(c) % 60

	return fmt.Sprintf("%02d:%02d", hour, minute)
}

//Add minutes to the Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)

}

//Subtract minutes from the Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
