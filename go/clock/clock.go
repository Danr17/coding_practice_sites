package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func (c Clock) SetTime() Clock {
	for c.m < 0 {
		c.m += 60
		c.h--
	}
	for c.h < 0 {
		c.h += 24
	}

	c.h += int(c.m / 60)
	c.m = c.m % 60
	c.h = c.h % 24

	return c
}

func New(hour, minute int) Clock {

	c := Clock{
		h: hour,
		m: minute,
	}

	return c.SetTime()
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	return c.SetTime()

}

func (c Clock) Subtract(minutes int) Clock {
	c.m -= minutes
	return c.SetTime()

}
