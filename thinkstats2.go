package thinkstats

//https://github.com/gonum/stat/blob/master/stat.go
import (
	"math"
	"math/rand"
	"sort"
)

// RandomSeed initialize random number with seed.
func RandomSeed(x int64) {
	rand.Seed(x)
}

// Odds for a given probability.
// Example: p=0.75 means 75 for and 25 against, or 3:1 odds in favor.
func Odds(p float64) float64 {
	if p == 1 {
		return math.Inf(1)
	}
	return p / (1. - p)
}

// Probability corresponding to given odds.
// Example: o=2 means 2:1 odds in favor, or 2/3 probability.
func Probability(o float64) float64 {
	return o / (o + 1)
}

// Probability2 corresponding to given odds.
// Example: yes=2, no=1 means 2:1 odds in favor, or 2/3 probability.
func Probability2(yes, no float64) float64 {
	return yes / (yes + no)
}

// Interpolator struct for mapping between sorted sequences; does linear interpolation.
// xs,ys are sorted slices
type Interpolator struct {
	xs []float64
	ys []float64
}

// InterpolatorNew creates struct with sorted slices
func InterpolatorNew(xs, ys []float64) Interpolator {
	sort.Float64s(xs)
	sort.Float64s(ys)
	return Interpolator{
		xs: xs,
		ys: ys,
	}
}

// Lookup x and return corresponding value of y.
func (i Interpolator) Lookup(x float64) float64 {
	return i.bisect(x)
}

// Reverse search y and return corresponding value of x.
func (i Interpolator) Reverse(y float64) float64 {
	return i.bisect(y)
}

// binarySearch
func binarySearch(t int, a []int) int {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] < t {
			lo = mid + 1
		} else if a[mid] == t {
			return mid
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// https://github.com/python/cpython/blob/3c88199e0be352c0813f145d7c4c83af044268aa/Lib/bisect.py
// bbisect Return the index where to insert item x in list a, assuming a is sorted
// similar to binary search
func bbisect(xs []float64, x float64) int {
	lo, hi := 0, len(xs)
	for lo < hi {
		mid := (lo + hi) / 2
		if x < xs[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func (in Interpolator) bisect(term float64) float64 {
	if term <= in.xs[0] {
		return in.ys[0]
	}
	if term >= in.xs[len(in.xs)-1] {
		return in.ys[len(in.ys)-1]
	}
	idx := bbisect(in.xs, term)
	frac := 1.0 * (term - in.xs[idx-1]) / (in.xs[idx] - in.xs[idx-1])
	return in.ys[idx-1] + frac*1.0*(in.ys[idx]-in.ys[idx-1])
}
