package main

import (
	"fmt"
	"strconv"
	"time"
	//"//fmt"
)

const count = 5

func gorutin() {
	l := make([][]string, count)
	c := make(chan []string)

	for i := 0; i < count; i++ {
		tmp := []string{"a"}
		fmt.Printf("#%d: len= %d, cap= %d, addr= %p, tmp= %#v\n", i, len(tmp), cap(tmp), tmp, tmp)
		var a func(i, j int, c chan []string)
		a = func(i, j int, c chan []string) {
			if j == count {
				c <- tmp
				fmt.Printf("ret - l: #%d:   len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
				return
			}

			tmp = append(tmp, strconv.Itoa(i)+"-"+strconv.Itoa(j))
			fmt.Printf("#%d - %d:   len= %d, cap= %d, addr= %p, tmp= %#v\n", i, j, len(tmp), cap(tmp), tmp, tmp)
			a(i, j+1, c)
		}
		go a(i, 0, c)
		fmt.Printf("befor chan   len= %d, cap= %d, addr= %p, tmp= %#v\n", len(tmp), cap(tmp), tmp, tmp)
		fmt.Printf("befor chan %d-l:    len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])

	}
	for i := 0; i < count; i++ {
		l[i] = <-c
		fmt.Printf("after chan %d-l:    len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
	}

	time.Sleep(4 * time.Second)
	for i := 0; i < count; i++ {
		fmt.Printf("ready l: #%d:   len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
	}
}

func gorutinPoint() {
	l := make([][]string, count)
	c := make(chan []string)

	for i := 0; i < count; i++ {
		tmp := []string{"a"}
		fmt.Printf("#%d: len= %d, cap= %d, addr= %p, tmp= %#v\n", i, len(tmp), cap(tmp), tmp, tmp)
		var a func(i int, j *int, c chan []string)
		a = func(i int, j *int, c chan []string) {
			if *j == count {
				c <- tmp
				fmt.Printf("ret - l: #%d:   len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
				return
			}

			tmp = append(tmp, strconv.Itoa(i)+"-"+strconv.Itoa(*j))
			fmt.Printf("#%d - %d:   len= %d, cap= %d, addr= %p, tmp= %#v\n", i, *j, len(tmp), cap(tmp), tmp, tmp)
			(*j)++
			a(i, j, c)
		}
		jj:=0
		go a(i, &jj, c)
		fmt.Printf("befor chan   len= %d, cap= %d, addr= %p, tmp= %#v\n", len(tmp), cap(tmp), tmp, tmp)
		fmt.Printf("befor chan %d-l:    len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])

	}
	for i := 0; i < count; i++ {
		l[i] = <-c
		fmt.Printf("after chan %d-l:    len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
	}

	time.Sleep(4 * time.Second)
	for i := 0; i < count; i++ {
		fmt.Printf("ready l: #%d:   len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
	}
}
func nogorutin() {
	l := make([][]string, count)

	for i := 0; i < count; i++ {
		tmp := []string{"a"}
		// fmt.Printf("#%d: len= %d, cap= %d, addr= %p, tmp= %#v\n", i, len(tmp), cap(tmp), tmp, tmp)
		var a func(i, j int)
		a = func(i, j int) {
			if j == count {
				l[i] = tmp
				// fmt.Printf("ret - l: #%d:   len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
				return
			}

			tmp = append(tmp, strconv.Itoa(i)+"-"+strconv.Itoa(j))
			// fmt.Printf("#%d - %d:   len= %d, cap= %d, addr= %p, tmp= %#v\n", i, j, len(tmp), cap(tmp), tmp, tmp)
			a(i, j+1)
		}
		a(i, 0)
		// fmt.Printf("after func a   len= %d, cap= %d, addr= %p, tmp= %#v\n", len(tmp), cap(tmp), tmp, tmp)
		// fmt.Printf("after func a   %d-l:    len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])

	}

	time.Sleep(count * time.Second)
	for i := 0; i < count; i++ {
		// fmt.Printf("ready l: #%d:   len= %d, cap= %d, addr= %p, l= %#v\n", i, len(l[i]), cap(l[i]), l[i], l[i])
	}
}
