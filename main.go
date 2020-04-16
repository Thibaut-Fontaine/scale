package main

import (
	"fmt"
	"sync"
)

type scale struct {
	s uint16
	n string
}

func scaleList() [4]scale {
	return [4]scale{
		{0xb5a, "Eolien"},
		{0xad5, "Ionien"},
		{0xad6, "Mixolydien"},
		{0xab5, "Lydien"},
	}
}

// C 0, C# 1, ..., A# 10, B 11.
func shiftScale(x uint16, i int) uint16 {
	const nn = 12
	const mask = 0xfff
	var u uint
	if i < 0 {
		i = -i
		u = nn - (uint(i % nn))
	} else {
		u = uint(i % nn)
	}
	return ((x >> u) | (x << (nn - u))) & mask
}

func noteName(i int) string {
	switch i {
	case 1:
		return "C#"
	case 2:
		return "D "
	case 3:
		return "D#"
	case 4:
		return "E "
	case 5:
		return "F "
	case 6:
		return "F#"
	case 7:
		return "G "
	case 8:
		return "G#"
	case 9:
		return "A "
	case 10:
		return "A#"
	case 11:
		return "B "
	default:
		return "C "
	}
}

func findScale(x uint16) []string {
	var r []string
	s := scaleList()
	var wg sync.WaitGroup

	for i := 0; i < len(s); i += 1 {
		wg.Add(1)
		go func(i int, r *[]string) {
			for k := 0; k < 12; k += 1 {
				a := shiftScale(x, k)
				if (s[i].s & a) == a {
					*r = append(*r, noteName(k)+" "+s[i].n+"\n")
				}
			}
			wg.Done()
		}(i, &r)
	}
	wg.Wait()
	return r
}

func main() {
	fmt.Println(findScale(0xad6))
}
