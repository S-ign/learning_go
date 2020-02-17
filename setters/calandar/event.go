package calandar

import (
	"errors"
)

type Event struct {
	title string
	Date
}

//EVENT SETTERS/////////////////////////////////////////////
func (e *Event) SetTitle(title string) error {
	if len(title) > 30 {
		return errors.New("title is too long!")
	}
	e.title = title
	return nil
}
/*//////////////////////////////////////////////////////////
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//EVENT GETTERS////////////////////////////////////////////*/
func (e *Event) Title() string {
	return e.title
}
////////////////////////////////////////////////////////////
