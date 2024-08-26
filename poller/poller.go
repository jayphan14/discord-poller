package poller

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

type Poller struct {
	ApiKey    string
	Interval  time.Duration
	DataStore *map[string]bool
}

func New(newDataStore *map[string]bool) (*Poller, error) {
	apiKey, err := getEnv("API_KEY", "", true)
	if err != nil {
		return nil, err
	}

	intervalStr, err := getEnv("POLL_INTERVAL", "10", false)
	if err != nil {
		return nil, err
	}

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		return nil, err
	}

	return &Poller{
		ApiKey:    apiKey,
		Interval:  time.Duration(interval) * time.Minute,
		DataStore: newDataStore,
	}, nil
}

// getEnv retrieves the environment variable, or returns the default value if it's not set
func getEnv(key, defaultValue string, required bool) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	if required {
		return "", fmt.Errorf("Can't find %v in env var", key)
	}
	return defaultValue, nil
}

// Poll starts polling from the channel that we want
func (p *Poller) Poll() {
	for {
		fmt.Println("polling")
		time.Sleep(p.Interval)
	}
}
