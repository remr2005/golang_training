package domain

import (
	"math/rand"
	"time"
)

func Add_task(title string) {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	TASKS = append(TASKS, Task{rand.Intn(10000000), title, false})
}
