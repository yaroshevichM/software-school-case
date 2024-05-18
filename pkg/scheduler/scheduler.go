package scheduler

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/yaroshevichM/software-school-case/pkg/service"
)

type Scheduler struct {
	cron    *cron.Cron
	service *service.Service
}

func NewScheduler(service *service.Service) *Scheduler {
	return &Scheduler{cron: cron.New(cron.WithSeconds()), service: service}
}

func (s *Scheduler) AddProcessEmail(spec string) error {
	_, err := s.cron.AddFunc(spec, s.processEmails)
	return err
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}

func (s *Scheduler) processEmails() {
	subscriptions, err := s.service.Subscription.GetAll()
	if err != nil {
		return
	}

	rate, err := s.service.Rate.GetUSDtoUAHRate()

	if err != nil {
		return
	}

	for _, subscription := range subscriptions {
		go s.service.Mail.SendMail(subscription.Email, "USD rate to UAH", fmt.Sprintf("Поточкний курс доллара до гривні %f", rate))
	}
}
