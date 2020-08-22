package robotname

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

const maxGenNames int = 26 * 26 * 10 * 10 * 10

//Robot holds the name
type Robot string

var exist = map[int]bool{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Name generate robot's name
func (r *Robot) Name() (string, error) {
	if len(*r) != 0 {
		return r.String(), nil
	}
	r.Reset()
	return r.String(), nil
}

// Reset the robot name
func (r *Robot) Reset() {
	for {
		if len(exist) == maxGenNames {
			log.Fatalln("you reached the maximum numer of names")
		}

		num := generateName()
		if exist[num] {
			continue
		}
		exist[num] = true

		stringify := strconv.Itoa(num)
		lit1, _ := strconv.Atoi(stringify[0:2])
		lit2, _ := strconv.Atoi(stringify[2:4])
		name := string(lit1) + string(lit2) + stringify[4:]
		*r = Robot(name)
		break
	}
}

func generateName() int {
	num := rand.Intn(999)
	return ('A'+rand.Intn(26))*100000 + ('A'+rand.Intn(26))*1000 + num

	//number := fmt.Sprintf("%03d", rand.Intn(999))
	//name := string('A'+rand.Intn(26)) + string('A'+rand.Intn(26)) + number
	//copy(buf, name)
}

// string method
func (r *Robot) String() string {
	return string(*r)
}
