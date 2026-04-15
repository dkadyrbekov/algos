package intervals

import "sort"

type TimeEvent struct {
	Time       int
	StartCount int
	EndCount   int
}

func findSets(intervals [][]int) int {
	timeEvents := make(map[int]TimeEvent)

	for i := 0; i < len(intervals); i++ {
		event := timeEvents[intervals[i][0]]

		event.StartCount++

		timeEvents[intervals[i][0]] = event

		event = timeEvents[intervals[i][1]]

		event.EndCount++

		timeEvents[intervals[i][1]] = event
	}

	events := make([]TimeEvent, 0, len(timeEvents))
	for t, e := range timeEvents {

		events = append(events, TimeEvent{
			Time:       t,
			StartCount: e.StartCount,
			EndCount:   e.EndCount,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Time < events[j].Time
	})

	maxCounter := 0
	counter := 0
	for _, e := range events {
		counter -= e.EndCount
		counter += e.StartCount

		if counter > maxCounter {
			maxCounter = counter
		}
	}

	return maxCounter
}
