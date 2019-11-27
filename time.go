package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p()

	then := time.Date(2019, 11, 27, 12, 29, 40, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Date())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(then.Nanosecond())
	p(then.Add(diff))
	p(then.Add(-diff))
}
