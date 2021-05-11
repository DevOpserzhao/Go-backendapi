package cron

import (
	"backend/pkg/cache"
	"github.com/robfig/cron/v3"
)

type Task struct {
	cache cache.RedisFace
	cron  *cron.Cron
}

func New(cache cache.RedisFace, cron *cron.Cron) *Task {
	return &Task{
		cache: cache,
		cron:  cron,
	}
}

func (c *Task) ClearBlackLogoutList() {
	c.cache.SClear("BlackLogoutList")
}

func (c *Task) Build() {
	_, err := c.cron.AddFunc("0 0 4 * * ?", c.ClearBlackLogoutList)
	if err != nil {
		return
	}
	c.cron.Start()
}
