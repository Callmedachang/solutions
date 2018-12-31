package main

import (
	"log"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	lenI := len(intervals)
	if lenI < 1 {
		return intervals
	}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	res := make([]Interval, 0)
	for i := 0; i < lenI; i++ {
		if i != lenI-1 {
			for canMerge(intervals[i],intervals[i+1]) && i < lenI-1 {
				newStart := 0
				newEnd := 0
				if intervals[i].Start > intervals[i+1].Start {
					newStart = intervals[i+1].Start
				} else {
					newStart = intervals[i].Start
				}
				if intervals[i].End > intervals[i+1].End {
					newEnd = intervals[i].End
				} else {
					newEnd = intervals[i+1].End
				}
				intervals[i+1] = Interval{Start: newStart, End: newEnd}
				i++
				if i >= lenI-1 {
					break
				}
			}
		}
		res = append(res, intervals[i])
	}
	return res
}

func canMerge(intervals1, intervals2 Interval) bool {
	if intervals1.Start <= intervals2.Start && intervals1.End >= intervals2.End {
		return true
	}
	if intervals2.Start <= intervals1.Start && intervals2.End >= intervals1.End {
		return true
	}
	if intervals2.Start <= intervals1.End && intervals2.Start >= intervals1.Start {
		return true
	}
	if intervals2.End <= intervals1.End && intervals2.End >= intervals1.Start {
		return true
	}
	if intervals2.End==intervals1.Start-1&&intervals1.End==intervals2.Start-1{
		return true
	}
	return false
}

func main() {
	res := make([]Interval, 0)
	res = append(res, Interval{1, 3})
	res = append(res, Interval{2, 6})
	//res = append(res, Interval{8, 10})
	//res = append(res, Interval{15, 18})
	log.Println(merge(res))
}
