package domain

import "errors"

func Update_task(id int, done bool) error {
	for _, val := range TASKS {
		if val.ID == id {
			val.Done = done
			return nil
		}
	}
	return errors.New("can not find element")
}
