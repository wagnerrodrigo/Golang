package monitor

import (
	"log"
	"time"
	"uptime-monitor/internal/db"
)

func StartMonitoring(urls []string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, url := range urls {
				status, err := CheckWebSite(url)
				if err != nil {
					log.Printf("Error checking %s: %v", url, err)
					continue
				}

				log.Printf("URL: %s, Status: %t", url, status)
				db.SaveResult(url, status)
			}
		}
	}
}
