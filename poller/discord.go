package poller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DiscordAPI struct {
	BaseUrl                        string
	ListMessageFromChannelEndpoint string
	ApiKey                         string
}

// FetchDiscordMessages sends a GET request to the Discord API and returns the response body
func (d *DiscordAPI) FetchDiscordMessagesFromChannel() ([]Message, error) {
	url := d.BaseUrl + d.ListMessageFromChannelEndpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", d.ApiKey)

	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code indicates success
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON response into a slice of any type
	var result []Message
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return result, nil
}

type Message struct {
	Content   string    `json:"content"`
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}
