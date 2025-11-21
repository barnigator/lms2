package main

import (
	"fmt"
	"sort"
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	name       string
	position   string
	salary     uint
	experience uint
}

type Company struct {
	workers []Worker
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	worker := Worker{
		name,
		position,
		salary,
		experience,
	}
	c.workers = append(c.workers, worker)
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	sort.Slice(c.workers, func(i, j int) bool {
		return c.workers[i].salary*c.workers[i].experience*12 > c.workers[j].salary*c.workers[j].experience*12
	})

	res := make([]string, 0, len(c.workers))
	for _, w := range c.workers {
		str := fmt.Sprintf("%s — %d — %s", w.name, w.salary*w.experience, w.position)
		res = append(res, str)
	}

	return res, nil
}
