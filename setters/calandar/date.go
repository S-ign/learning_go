package calandar

import (
	"errors"
)

//combines three defined underlined types to create date structure
type Date struct {
	year    int
	month   int
	day     int
}

////////////DATE SETTERS////////////////////////////////////
//setter methods that validates arguments.
//marking the reciever (*Date) a pointer is important
//since the passed in argument will be lost since
//we are not returning it.
func (d *Date) SetYear(year int) error {
	if year < 2000 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
////////////////////////////////////////////////////////////
