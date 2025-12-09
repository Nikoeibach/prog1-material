package main

type Duration int

func FromSeconds(s int) Duration {
	return Duration(s)
}

func FromMinutes(m int) Duration {
	return Duration(m * 60)
}

func FromHours(o int) Duration {
	return Duration(o * 60 * 60)
}

func (d Duration) Seconds() int {
	return int(d)
}

func (d Duration) Minutes() int {
	return int(d / 60)
}

func (d Duration) Hours() int {
	return int(d / 60 / 60)
}
