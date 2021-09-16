package main

import (
	"fmt"
	"time"
)

var week time.Duration

// Format("20060102") ？？？？

func main() {
	t := time.Now()
	fmt.Println(t)
	//fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Weekday())
	//fmt.Printf("%02d.%02d,.%04d\n", t.Day(), t.Month(), t.Year())
	//
	//fmt.Println(t.Format(time.RFC822))
	//
	//fmt.Println(t.Format("02 Jan 2006 15:04"))

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)

	s := t.Format("20060102")
	fmt.Println(t, "===>", s)

}
