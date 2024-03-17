package hotel

import (
	"fmt"
	"sort"
)

// [] -> 0
// [(1,2), (0,1)] -> 1
// [(0,2), (0,1)] -> 2
// [(0,99999999999)] -> 1

type guest struct {
	in  int
	out int
}

type event struct {
	time int
	in   bool
}

func maxCapasity(guests []guest) int {
	maxGuests := 0
	currentGuests := 0

	events := make([]event, 0)

	for _, g := range guests {
		events = append(events, event{g.in, true})
		events = append(events, event{g.out, false})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return !events[i].in
		}
		return events[i].time < events[j].time
	})

	for _, e := range events {
		if e.in {
			currentGuests++
		} else {
			currentGuests--
		}

		if currentGuests > maxGuests {
			maxGuests = currentGuests
		}
	}

	return maxGuests
}

type Case struct {
	task     []guest
	expected int
}

func Hotel() {
	cases := []Case{
		{[]guest{}, 0},
		{[]guest{{0, 1}, {1, 2}}, 1},
		{[]guest{{1, 2}, {0, 1}}, 1},
		{[]guest{{0, 1}, {0, 2}}, 2},
		{[]guest{{0, 9999999999999999}}, 1},
	}

	for _, c := range cases {
		got := maxCapasity(c.task)
		fmt.Printf("%v получили %d, должно быть %d\n", c.task, got, c.expected)
		if got != c.expected {
			panic("ты облажался, бро. поробуй ещё раз")
		} else {
			fmt.Println("щас повезло, молодец")
		}
	}
}
