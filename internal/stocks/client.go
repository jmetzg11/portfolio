package stocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Stock struct {
	Symbol        string
	MarketValue   float64
	CostBasis     float64
	UnrealizedPL  float64
	UnrealizedPPC float64
	CurrentPrice  float64
	LastdayPrice  float64
	ChangeToday   float64
	Reserve       float64
}

type Coin struct {
	Currency     string
	Qty          float64
	DollarAmount float64
	Reserve      float64
}

type Summary struct {
	Alpaca []Stock `json:"alpaca"`
	Crypto []Coin  `json:"crypto"`
}

type Client struct {
	url  string
	http *http.Client
	ttl  time.Duration

	mu      sync.Mutex
	cached  *Summary
	fetched time.Time
}

func New(url string) *Client {
	return &Client{
		url:  url,
		http: &http.Client{Timeout: 5 * time.Second},
		ttl:  5 * time.Minute,
	}
}

func (c *Client) Summary() (*Summary, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.cached != nil && time.Since(c.fetched) < c.ttl {
		return c.cached, nil
	}

	summary, err := c.fetch()
	if err != nil {
		if c.cached != nil {
			return c.cached, nil
		}
		return nil, err
	}

	c.cached = summary
	c.fetched = time.Now()
	return summary, nil
}

func (c *Client) fetch() (*Summary, error) {
	resp, err := c.http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("stocks api returned %s", resp.Status)
	}

	var summary Summary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, err
	}
	return &summary, nil
}
