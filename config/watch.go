package config

import (
	"io/ioutil"
	"log"
	"time"
)

// Service reload config dynamic
type Service struct {
	Config *Config
	Path   string
}

//Watch reload config every time.duration
func (s *Service) Watch(d time.Duration) {
	for {
		err := s.Reload()
		if err != nil {
			log.Println(err)
		}

		log.Println("watcing....")

		time.Sleep(d)
	}
}

// Reload read the config and apply the change
func (s *Service) Reload() error {
	data, err := ioutil.ReadFile(s.Path)
	if err != nil {
		return err
	}

	err = s.Config.SetConfig(data)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
