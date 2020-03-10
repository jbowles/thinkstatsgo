package thinkstats

import (
	"math"
	"testing"
)

var (
	percscores = []float64{55, 66, 77, 88, 99}
)

func TestOdds(t *testing.T) {
	oddstest := []struct {
		odds   float64
		expect float64
	}{
		{0.75, 3.0},        //0.75 == (75 for, 25 against) == 3:1 odds
		{1.0, math.Inf(1)}, //1.0 1:1 odds, infinity!!
	}
	for i, o := range oddstest {
		r := Odds(o.odds)
		if o.expect != r {
			t.Errorf("for %d expect %f but got %f", i, o.expect, r)
		}
	}
}

func TestPercentileRank(t *testing.T) {
	prank := []struct {
		scores []float64
		score  float64
		expect int
	}{
		{percscores, 77, 60},
		{percscores, 88, 80},
	}
	for id, p := range prank {
		res := PercentileRank(p.scores, p.score)
		if p.expect != res {
			t.Errorf("for %d expect %d but got %d", id, p.expect, res)
		}
	}
}
func TestPercentile(t *testing.T) {
	perci := []struct {
		scores []float64
		rank   int
		expect float64
	}{
		{percscores, 50, 77},
		{percscores, 80, 88},
	}
	for id, p := range perci {
		res := Percentile(p.scores, p.rank)
		if p.expect != res {
			t.Errorf("for %d expect %f but got %f", id, p.expect, res)
		}
	}
}
