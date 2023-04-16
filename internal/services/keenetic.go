package services

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"time"
)

type KeeneticService struct {
	KeenDnsDomain string
}


func NewKeeneticService(keenDnsDomain string) *KeeneticService {
	return &KeeneticService{
		KeenDnsDomain: keenDnsDomain,
	}
}

func (s *KeeneticService) Check() error {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	log.Printf("[Checking] Domain: %s", s.KeenDnsDomain)
	req, _ := http.NewRequest("GET", s.KeenDnsDomain, bytes.NewReader([]byte{}))

	res, err := client.Do(req)

	if err != nil {
		log.Print(err.Error())
		return err
	}

        log.Printf("[Checked] Status code: %d; Domain: %s", res.StatusCode, s.KeenDnsDomain)

	if res.StatusCode != http.StatusOK {
		log.Printf("Error status code: %d", res.StatusCode)

		return errors.New("Keenetic API request bad satatus")
	}

	return nil
}
