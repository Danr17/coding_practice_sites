package clock

import "fmt"

//Clock defines the clock
type Clock struct {
	h int
	m int
}

//New instantiate a Clock
func New(hour, minute int) Clock {

	for minute < 0 {
		minute += 60
		hour--
	}
	for hour < 0 {
		hour += 24
	}

	hour += int(minute / 60)
	minute = minute % 60
	hour = hour % 24

	return Clock{
		h: hour,
		m: minute,
	}

}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

//Add minutes to the Clock
func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	return New(c.h, c.m)

}

//Subtract minutes from the Clock
func (c Clock) Subtract(minutes int) Clock {
	c.m -= minutes
	return New(c.h, c.m)

}
