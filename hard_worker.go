package main

import (
	"fmt"
	"time"
)

type Hardworker struct {
	*Job
	name string
	age float64
}

func (worker *Hardworker) Perform() {
	fmt.Printf("His name is %s age %g\n", worker.name, worker.age)
	time.Sleep(time.Duration(worker.age) * time.Second)
	fmt.Printf("Finished: %s\n", worker.Jid)
}

func NewHardworker(job Job) Worker {
	name := job.Args[0].(string)
	age := job.Args[1].(float64)
	return &Hardworker{&job, name, age}
}
