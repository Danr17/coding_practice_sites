package summultiples

//SumMultiples sums of all the unique multiples of particular numbers up to but not including that number
func SumMultiples(limit int, divisors ...int) int {
	seen := map[int]bool{}
	for _, factor := range divisors {
		if factor > 0 {
			iter := factor
			for iter < limit {
				seen[iter] = true
				iter += factor
			}
		}
	}
	var sum int = 0
	for k := range seen {
		sum += k
	}
	return sum
}

/*
$ go test -v --bench . --benchmem
=== RUN   TestSumMultiples
--- PASS: TestSumMultiples (0.00s)
goos: linux
goarch: amd64
pkg: summultiples
BenchmarkSumMultiples-8   	    1111	   1030360 ns/op	  487849 B/op	     467 allocs/op
*/

/*
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for num := 1; num < limit; num++ {
		set := make(map[int]bool)
		for _, div := range divisors {
			if div == 0 {
				continue
			}
			if num%div == 0 {
				set[num] = true
			}
		}
		for key := range set {
			sum = sum + key
		}
	}
	return sum
}
*/

/*
$ go test -v --bench . --benchmem
=== RUN   TestSumMultiples
--- PASS: TestSumMultiples (0.01s)
goos: linux
goarch: amd64
pkg: summultiples
BenchmarkSumMultiples-8   	     456	   2298477 ns/op	  579716 B/op	   18116 allocs/op
*/
