package poller

import (
	"discord-poller/db"
	"discord-poller/util"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

type Poller struct {
	ApiKey     string
	Interval   time.Duration
	DiscordAPI *DiscordAPI
	Queries    *db.Queries
}

func New(conn *pgx.Conn) (*Poller, error) {
	apiKey, err := util.GetEnv("API_KEY", "", true)
	if err != nil {
		return nil, err
	}

	intervalStr, err := util.GetEnv("POLL_INTERVAL", "10", false)
	if err != nil {
		return nil, err
	}

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		return nil, err
	}

	newDiscordApi := &DiscordAPI{
		BaseUrl:                        "https://discord.com/api/v10",
		ListMessageFromChannelEndpoint: "/channels/1245256033530548268/messages",
		ApiKey:                         apiKey,
	}

	return &Poller{
		ApiKey:     apiKey,
		Interval:   time.Duration(interval) * time.Minute,
		DiscordAPI: newDiscordApi,
		Queries:    db.New(conn),
	}, nil
}

// Poll starts polling from the channel that we want
func (p *Poller) Poll() {
	for {
		fmt.Println("polling")
		messages, err := p.DiscordAPI.FetchDiscordMessagesFromChannel()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Messages :", messages)
		time.Sleep(p.Interval)
	}
}
