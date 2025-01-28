package domain

import (
	"errors"
	"slices"
)

func Remove_task(id int) error {
	for i, val := range TASKS {
		if val.ID == id {
			slices.Delete(TASKS, i, i+1)
			return nil
		}
	}
	return errors.New("can not find element")
}
