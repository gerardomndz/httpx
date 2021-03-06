package httpx

import (
	"time"
)

// Options contains configuration options for the client
type Options struct {
	Threads int
	// Timeout is the maximum time to wait for the request
	Timeout time.Duration
	// RetryMax is the maximum number of retries
	RetryMax int

	CustomHeaders    map[string]string
	FollowRedirects  bool
	DefaultUserAgent string

	HttpProxy  string
	SocksProxy string

	// VHOSTs options
	VHostIgnoreStatusCode    bool
	VHostIgnoreContentLength bool
	VHostIgnoreNumberOfWords bool
	VHostIgnoreNumberOfLines bool
	VHostStripHTML           bool

	// VHostimilarityRatio 1 - 100
	VHostSimilarityRatio int
}

// DefaultOptions contains the default options
var DefaultOptions = Options{
	Threads:  25,
	Timeout:  30 * time.Second,
	RetryMax: 5,
	// VHOSTs options
	VHostIgnoreStatusCode:    false,
	VHostIgnoreContentLength: true,
	VHostIgnoreNumberOfWords: false,
	VHostIgnoreNumberOfLines: false,
	VHostStripHTML:           false,
	VHostSimilarityRatio:     85,
	DefaultUserAgent:         "httpx - Open-source project (github.com/projectdiscovery/httpx)",
	// Smuggling Options
}
