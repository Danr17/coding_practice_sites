package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//Robot holds the name
type Robot string

//Name generate robot's name
func (r *Robot) Name() (string, error) {
	if *r != "" {
		return r.String(), nil
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	name := arr[0] + arr[1] + strconv.Itoa(rand.Intn(999-100)+100)
	*r = Robot(name)
	return r.String(), nil
}

// Reset the robot name
func (r *Robot) Reset() {
	*r = ""
	_, _ = r.Name()
}

// Reset the robot name
func (r *Robot) String() string {
	return string(*r)
}
