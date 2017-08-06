package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

const minutesInDay = 24 * 60

func New(hour, minute int) Clock {
	// calculate minutes from midnight (`(x+minutesInDay)%minutesInDay` ensures value >=0)
	totalMinutes := (((hour*60 + minute) % minutesInDay) + minutesInDay) % minutesInDay
	return Clock{totalMinutes / 60, totalMinutes % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}