package zdesk

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/ajg/form"
	"github.com/hashicorp/go-cleanhttp"
)

// APIUsernameEnvVar is the environment variable where the Zendesk username /
// email should be read from.
const APIEmailEnvVar = "ZDESK_EMAIL"

// APIKeyEnvVar is the name of the environment variable where the Zendesk API
// key should be read from.
const APITokenEnvVar = "ZDESK_TOKEN"

// APIDomainURLEnvVar is the name of the environment variable where the Zendesk
// domain, such as https://example.zendesk.com should be read from.
const APIDomainEnvVar = "ZDESK_DOMAIN"

// ProjectURL is the url for this library.
var ProjectURL = "github.com/fprimex/zdesk-go"

// ProjectVersion is the version of this library.
var ProjectVersion = "0.1"

// UserAgent is the user agent for this particular client.
var UserAgent = fmt.Sprintf("zdesk-go/%s (+%s; %s)",
	ProjectVersion, ProjectURL, runtime.Version())

// Client is the main entrypoint to the API library.
type Client struct {
	// domain is the address of Zendesk's API endpoint.
	domain string

	// email is the Zendesk login / email.
	email string

	// apiKey is the Zendesk API token to authenticate requests.
	token string

	// url is the parsed URL from Address
	url *url.URL

	// HTTPClient is the HTTP client to use. If one is not provided, a default
	// client will be used.
	HTTPClient *http.Client
}

// DefaultClient instantiates a new Zendesk API client. This function requires
// the environment variables to be set.
func DefaultClient() (*Client, error) {
	return NewClient(os.Getenv(APIEmailEnvVar),
		os.Getenv(APITokenEnvVar),
		os.Getenv(APIDomainEnvVar))
}

// NewClient creates a new API client with the given key and the domain
// endpoint.
func NewClient(email string, token string, domain string) (*Client, error) {
	client := &Client{email: fmt.Sprintf("%s/token", email), token: token, domain: domain}
	return client.init()
}

func (c *Client) init() (*Client, error) {
	u, err := url.Parse(c.domain)
	if err != nil {
		return nil, err
	}
	c.url = u

	if c.HTTPClient == nil {
		c.HTTPClient = cleanhttp.DefaultClient()
	}

	return c, nil
}

// Get issues an HTTP GET request.
func (c *Client) Get(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("GET", p, ro)
}

// Head issues an HTTP HEAD request.
func (c *Client) Head(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("HEAD", p, ro)
}

// Post issues an HTTP POST request.
func (c *Client) Post(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("POST", p, ro)
}

// PostForm issues an HTTP POST request with the given interface form-encoded.
func (c *Client) PostForm(p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	return c.RequestForm("POST", p, i, ro)
}

// Put issues an HTTP PUT request.
func (c *Client) Put(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("PUT", p, ro)
}

// PutForm issues an HTTP PUT request with the given interface form-encoded.
func (c *Client) PutForm(p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	return c.RequestForm("PUT", p, i, ro)
}

// Delete issues an HTTP DELETE request.
func (c *Client) Delete(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("DELETE", p, ro)
}

// Request makes an HTTP request against the HTTPClient using the given verb,
// Path, and request options.
func (c *Client) Request(verb, p string, ro *RequestOptions) (*http.Response, error) {
	req, err := c.RawRequest(verb, p, ro)
	if err != nil {
		return nil, err
	}

	resp, err := checkResp(c.HTTPClient.Do(req))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RequestForm makes an HTTP request with the given interface being encoded as
// form data.
func (c *Client) RequestForm(verb, p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	if ro == nil {
		ro = new(RequestOptions)
	}

	if ro.Headers == nil {
		ro.Headers = make(map[string]string)
	}
	ro.Headers["Content-Type"] = "application/x-www-form-urlencoded"

	buf := new(bytes.Buffer)
	if err := form.NewEncoder(buf).KeepZeros(true).DelimitWith('|').Encode(i); err != nil {
		return nil, err
	}
	body := buf.String()

	ro.Body = strings.NewReader(body)
	ro.BodyLength = int64(len(body))

	return c.Request(verb, p, ro)
}

// checkResp wraps an HTTP request from the default client and verifies that the
// request was successful. A non-200 request returns an error formatted to
// included any validation problems or otherwise.
func checkResp(resp *http.Response, err error) (*http.Response, error) {
	// If the err is already there, there was an error higher up the chain, so
	// just return that.
	if err != nil {
		return resp, err
	}

	switch resp.StatusCode {
	case 200, 201, 202, 204, 205, 206:
		return resp, nil
	default:
		return resp, NewHTTPError(resp)
	}
}
