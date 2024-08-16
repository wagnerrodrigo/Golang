package monitor

import (
	"net/http"
	"time"
)

func CheckWebSite(url string) (bool, error) {
	cliente := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := cliente.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil

}
