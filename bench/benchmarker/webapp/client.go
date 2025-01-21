package webapp

import (
	"context"
	"io"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/isucon/isucandar/agent"
)

type Client struct {
	agent            *agent.Agent
	requestModifiers []func(*http.Request)
}

type ClientConfig struct {
	TargetBaseURL         string
	TargetAddr            string
	ClientIdleConnTimeout time.Duration
}

type roundRobin[T any] struct {
	ips     []T
	counter uint64
}

func (i *roundRobin[T]) getNext() T {
	idx := atomic.AddUint64(&i.counter, 1) % uint64(len(i.ips))
	return i.ips[idx]
}

var local = roundRobin[net.IP]{
	ips: []net.IP{
		// "192.168.0.101",
		net.ParseIP("192.168.0.102"),
		// net.ParseIP("192.168.0.103"),
		// net.ParseIP("192.168.0.104"),
		// net.ParseIP("192.168.0.105"),
		// net.ParseIP("192.168.0.106"),
		// net.ParseIP("192.168.0.107"),
		// net.ParseIP("192.168.0.108"),
		// net.ParseIP("192.168.0.109"),
		// net.ParseIP("192.168.0.110"),
	},
	counter: 0,
}

var dest = roundRobin[string]{
	ips: []string{
		"192.168.0.249:8080",
		// "192.168.0.250:8080",
		// "192.168.0.251:8080",
		// "192.168.0.252:8080",
	},
	counter: 0,
}

func NewClient(config ClientConfig, agentOptions ...agent.AgentOption) (*Client, error) {
	trs := agent.DefaultTransport.Clone()
	trs.IdleConnTimeout = config.ClientIdleConnTimeout

	trs.DialContext = func(ctx context.Context, network, _ string) (net.Conn, error) {
		dialer := net.Dialer{
			LocalAddr: &net.TCPAddr{
				IP:   local.getNext(),
				Port: 0,
			},
		}
		return dialer.DialContext(ctx, network, dest.getNext())
	}
	ag, err := agent.NewAgent(
		append([]agent.AgentOption{
			agent.WithBaseURL(config.TargetBaseURL),
			agent.WithTimeout(10 * time.Second),
			agent.WithTransport(trs),
		}, agentOptions...)...,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		agent: ag,
		requestModifiers: []func(*http.Request){func(req *http.Request) {
			if req.Method == http.MethodPost && req.Header.Get("Content-Type") == "" {
				req.Header.Add("Content-Type", "application/json; charset=utf-8")
			}
		}},
	}, nil
}

func (c *Client) AddRequestModifier(modifier func(*http.Request)) {
	c.requestModifiers = append(c.requestModifiers, modifier)
}

func (c *Client) SetCookie(cookie *http.Cookie) {
	c.agent.HttpClient.Jar.SetCookies(c.agent.BaseURL, []*http.Cookie{cookie})
}

func closeBody(resp *http.Response) {
	if resp.Body != nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}
