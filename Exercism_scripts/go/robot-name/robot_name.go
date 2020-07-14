package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

//Robot holds the name
type Robot string

var exist = map[string]bool{}

//Name generate robot's name
func (r *Robot) Name() (string, error) {
	if *r != "" {
		return r.String(), nil
	}
	for {

		r.Reset()
		if _, ok := exist[r.String()]; ok {
			continue
		}
		exist[r.String()] = true
		break
	}

	return r.String(), nil
}

// Reset the robot name
func (r *Robot) Reset() {

	rand.Seed(time.Now().UnixNano())
	name := string('A'+rand.Intn(26)) + string('A'+rand.Intn(26)) + strconv.Itoa(rand.Intn(999-100)+100)

	*r = Robot(name)
}

// Reset the robot name
func (r *Robot) String() string {
	return string(*r)
}
