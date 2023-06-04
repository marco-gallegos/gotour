// NOTE: part of a quick challenge
// Write a Tracker struct - job is to track integers
// - track(int) -> keeps track of the given integer
// - min() -> Returns the smallest integer tracked so far
// - max() -> largest
// - mean() -> average of all integers tracked so far
// - mode() -> return the integer that has been tracked the most
//
// All functions must run in O(1) time

// Final State: potentially solved

package challenges

import (
    "fmt"
)

type Tracker struct {
	values []int

	min int
	max int
	sumary int

	modes map[int]int // the idea is store how many times a integer is tracked()
	currentModeCount int // the number in the map
	mode int // the key of the map
}


// kind of contructor
func newTracker() Tracker {
	newTracker := Tracker{}
	newTracker.min = 10000 //TODO: potential bug. not so elegant but should work for now
	newTracker.modes = make(map[int]int) // fill modes
	newTracker.currentModeCount = 0
	newTracker.mode = 0
	return newTracker
}


func (t *Tracker) track(value int){
	t.values = append(t.values, value)
	
	// add logic to check if this is the minimum so far
	if (value < t.min ){
		t.min = value
	}

	// check if is the max so far
	if (value > t.max ){
        t.max = value
    }


	t.sumary = t.sumary + value // segregate to mean funtion the division

	t.modes[value] += 1

	if (t.modes[value] > t.currentModeCount){
		t.mode = value
		t.currentModeCount = t.modes[value]
	}
}

func (t *Tracker) Min() int {
    return t.min
}


func (t *Tracker) Max() int {
    return t.max
}


func (t *Tracker) Mean() float64 {
	var test float64
	test = float64(t.sumary) / float64(len(t.values)) //TODO: not sure about O value of len
	return test
}

func (t *Tracker) Mode() int {
  return t.mode
}


func Main() {
  tracker := newTracker()
	tracker.track(2)
	tracker.track(2)
	tracker.track(5)
	tracker.track(6)
	tracker.track(1)
	tracker.track(1)
	tracker.track(2)

	fmt.Println(tracker.values)
	fmt.Println(tracker.Min())
	fmt.Println(tracker.Max())
	fmt.Println(tracker.Mean())
	fmt.Println(tracker.Mode())
}
