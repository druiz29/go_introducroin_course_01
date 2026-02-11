package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	p("now:", now)
	then :=  time.Date(2019, 11 , 16, 12, 41, 12, 896518, time.UTC)
	p("then:", then)
	then =  then.Add(time.Hour * 1)
	p("then:", then)
	p("then year:", then.Year())
	p("then month:", then.Month())
	p("then day:", then.Day())
	p("then hour:", then.Hour())
	p("then minute:", then.Minute())
	p("then location:", then.Location())
	p("then weekday:", then.Weekday())
	p("then second:", then.Second())

	p("then before now?", then.Before(now))
	p("then after now?", then.After(now))
	p("then equal now?", then.Equal(now))

	diff :=  now.Sub(then)
	p(diff)


}



