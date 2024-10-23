package service

import (
	"github.com/robfig/cron/v3"
)

type CronService struct {
	cron *cron.Cron
}

func NewCronService() *CronService {
	return &CronService{
		cron: cron.New(),
	}
}

func (c *CronService) Start() {
	c.cron.Start()
}

func (c *CronService) Stop() {
	c.cron.Stop()
}

func (c *CronService) AddJob(spec string, cmd func()) (cron.EntryID, error) {
	return c.cron.AddFunc(spec, cmd)
}
