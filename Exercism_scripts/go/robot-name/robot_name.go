package robotname

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const maxGenNames int = 26 * 26 * 10 * 10 * 10

//Robot holds the name
type Robot string

var exist = map[string]bool{}

//Name generate robot's name
func (r *Robot) Name() (string, error) {
	if *r != "" {
		return r.String(), nil
	}
	r.Reset()
	return r.String(), nil
}

// Reset the robot name
func (r *Robot) Reset() {
	for {
		if len(exist) == maxGenNames-1 {
			log.Fatalln("you reached the maximum numer of names")
		}

		r.generateName()
		if exist[r.String()] {
			continue
		}
		exist[r.String()] = true
		break
	}
}

func (r *Robot) generateName() {
	rand.Seed(time.Now().UnixNano())
	number := fmt.Sprintf("%03d", rand.Intn(999))
	name := string('A'+rand.Intn(26)) + string('A'+rand.Intn(26)) + number
	*r = Robot(name)
}

// string method
func (r *Robot) String() string {
	return string(*r)
}
