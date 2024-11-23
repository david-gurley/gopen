package gopen

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	PSM_CLIENT_KIND_ENT = "ent"
	PSM_CLIENT_KIND_DSS = "dss"
)

var (
	ErrLoginMissingInfo      = fmt.Errorf("login missing username, password, or hostname")
	ErrLoginKindNotDetected  = fmt.Errorf("login client kind not detected")
	ErrRequestMissingCookies = fmt.Errorf("reqeust missing cookies")
	ErrNotAuthorized         = fmt.Errorf("not authorized: %s", http.StatusText(http.StatusUnauthorized))
)

type RequestParameters struct {
	URL     *url.URL
	Method  string
	Payload interface{}
	Out     interface{}
}

type ListResponse[T any] struct {
	ApiResponseList
	Items Items[T] `json:"items"`
}

type Items[T any] []T

type Client interface {
	PsmEntClient | PsmDssClient
}

type Clients[T Client] []T

type PsmClientConfig struct {
	Kind     string `json:"kind"`
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Tenant   string `json:"tenant"`
}

type PsmClient struct {
	client   *http.Client
	Kind     string
	username string
	password string
	hostname string
	tenant   string
	cookies  []*http.Cookie
}

func NewPsmClient(config *PsmClientConfig) (*PsmClient, error) {
	psmClient := PsmClient{
		newHttpClient(),
		config.Kind,
		config.Username,
		config.Password,
		config.Hostname,
		config.Tenant,
		make([]*http.Cookie, 0),
	}
	if psmClient.tenant == "" {
		psmClient.tenant = "default"
	}
	return &psmClient, nil
}

func (client *PsmClient) Hostname() string {
	return client.hostname
}
func (client *PsmClient) Login() error {
	if client.username != "" && client.password != "" && client.hostname != "" {
		u := fmt.Sprintf("https://%s/v1/login", client.hostname)
		login := LoginPayload{
			Username: client.username,
			Password: client.password,
			Tenant:   client.tenant,
		}
		loginJson, err := json.Marshal(login)
		if err != nil {
			return err
		}
		request, err := http.NewRequest("POST", u, bytes.NewBuffer(loginJson))
		if err != nil {
			return err
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		request = request.WithContext(ctx)
		response, err := client.client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		switch statusCode := response.StatusCode; statusCode {
		case 200:
			client.cookies = response.Cookies()
			return nil
		case 401:
			return ErrNotAuthorized
		default:
			return fmt.Errorf("unknown response code: %s", http.StatusText(response.StatusCode))
		}
	}
	return ErrLoginMissingInfo

}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tenant   string `json:"tenant"`
}

// basic auth
func Login(config *PsmClientConfig) (interface{}, error) {
	if config.Username != "" && config.Password != "" && config.Hostname != "" {
		u := fmt.Sprintf("https://%s/v1/login", config.Hostname)
		login := LoginPayload{
			Username: config.Username,
			Password: config.Password,
			Tenant:   config.Tenant,
		}
		loginJson, err := json.Marshal(login)
		if err != nil {
			return nil, err
		}
		httpClient := newHttpClient()
		request, err := http.NewRequest("POST", u, bytes.NewBuffer(loginJson))
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		request = request.WithContext(ctx)
		response, err := httpClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		switch statusCode := response.StatusCode; statusCode {
		case 200:
			cookies := response.Cookies()
			u := url.URL{
				Scheme: "https",
				Host:   fmt.Sprintf("%s", config.Hostname),
				Path:   "/configs/cluster/v1/version",
			}
			out := make(map[string]interface{})
			rp := RequestParameters{
				Method: "GET",
				URL:    &u,
				Out:    &out,
			}
			err := Request(httpClient, rp, cookies)
			if err != nil {
				return nil, err
			}
			if out["status"] != nil {
				build, err := NewPsmBuild()
				if err != nil {
					return nil, err
				}
				version := build.Kind(out["status"].(map[string]interface{})["build-version"].(string))
				err = Request(httpClient, rp, cookies)
				if err != nil {
					return nil, err
				}
				client, err := NewPsmClient(config)
				if err != nil {
					return nil, err
				}
				client.cookies = cookies
				switch version {
				case PSM_CLIENT_KIND_ENT:
					client.Kind = PSM_CLIENT_KIND_ENT
					return &PsmEntClient{client}, nil
				case PSM_CLIENT_KIND_DSS:
					client.Kind = PSM_CLIENT_KIND_DSS
					return &PsmDssClient{client}, nil
				default:
					return nil, ErrLoginKindNotDetected
				}
			} else {
			}
		case 401:
			return nil, ErrNotAuthorized
		default:
			return nil, fmt.Errorf("unknown response code: %s", http.StatusText(response.StatusCode))
		}

	} else {
		return nil, ErrLoginMissingInfo
	}
	return nil, nil
}

func (client *PsmClient) Close() error {
	return nil
}

func (client *PsmClient) Get(path string, qp map[string]string, out interface{}) error {

	values := url.Values{}

	for k, v := range qp {
		values.Set(k, v)
	}
	u := url.URL{
		Scheme:   "https",
		Host:     fmt.Sprintf("%s", client.hostname),
		Path:     path,
		RawQuery: values.Encode(),
	}
	rp := RequestParameters{
		Method: "GET",
		URL:    &u,
		Out:    out,
	}
	err := client.Request(rp)
	if err != nil {
		return err
	}
	return nil
}

func (client *PsmClient) Post(path string, p interface{}, out interface{}) error {
	u := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s", client.hostname),
		Path:   path,
	}
	rp := RequestParameters{
		Method:  "POST",
		URL:     &u,
		Payload: p,
		Out:     out,
	}
	err := client.Request(rp)
	if err != nil {
		return err
	}

	return nil

}

func (client *PsmClient) Put(path string, p interface{}, out interface{}) error {
	u := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s", client.hostname),
		Path:   path,
	}
	rp := RequestParameters{
		Method:  "PUT",
		URL:     &u,
		Payload: p,
		Out:     out,
	}
	err := client.Request(rp)
	if err != nil {
		return err
	}
	return nil

}

func (client *PsmClient) Delete() (*http.Response, error) {
	return &http.Response{}, nil

}

func (client *PsmClient) Noop() bool {
	return false
}

func (client *PsmClient) Request(rp RequestParameters) error {
	return Request(client.client, rp, client.cookies)
}

func Request(cl *http.Client, rp RequestParameters, cookies []*http.Cookie) error {
	var payload io.Reader
	if rp.Method == "POST" || rp.Method == "PUT" {
		if rp.Payload == nil {
			var p string
			p = "{}"
			payload = strings.NewReader(p)
		} else {
			b := bytes.Buffer{}
			encoder := json.NewEncoder(&b)
			encoder.Encode(rp.Payload)
			payload = &b
		}

	} else if rp.Method == "GET" {
		payload = nil
	}
	var httpClient *http.Client
	if cl == nil {
		httpClient = newHttpClient()
	} else {
		httpClient = cl
	}
	request, err := http.NewRequest(rp.Method, rp.URL.String(), payload)
	if err != nil {
		return fmt.Errorf("error creating new request: %s", err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	request = request.WithContext(ctx)
	buildRequestHeaders(cookies, request)
	resp, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error: %s executing request: %#v url: %s", err.Error(), request, request.URL.String())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %s", err.Error())
	}
	err = json.Unmarshal(body, rp.Out)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %s", err.Error())
	}
	switch statusCode := resp.StatusCode; statusCode {
	case 400:
		return fmt.Errorf("Bad request parameters: %s : %v", http.StatusText(resp.StatusCode), string(body))
	case 401:
		return fmt.Errorf("Unauthorized request: %s : %v", http.StatusText(resp.StatusCode), string(body))
	case 403:
		return fmt.Errorf("Authentication domain does not currently allow log-ins: %s : %v", http.StatusText(resp.StatusCode), string(body))
	case 404:
		return fmt.Errorf("Page not found: %s: %v", http.StatusText(resp.StatusCode), string(body))
	default:
		return nil
	}
	return nil
}
func (client *PsmClient) buildRequestHeaders(request *http.Request) error {
	return buildRequestHeaders(client.cookies, request)
}
func buildRequestHeaders(cookies []*http.Cookie, request *http.Request) error {
	if cookies != nil {
		for _, cookie := range cookies {
			request.AddCookie(cookie)
		}
	} else {
		return ErrRequestMissingCookies
	}
	return nil
}

func newHttpClient() *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	tr.MaxIdleConns = 100
	tr.MaxConnsPerHost = 100
	tr.MaxIdleConnsPerHost = 100
	return &http.Client{Transport: tr}
}
