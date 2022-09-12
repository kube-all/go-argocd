/*
Copyright 2022 The kubeall.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"k8s.io/klog/v2"
	"net/url"
	"strings"
)

const (
	userAgent   = "kubeall-go-argocd-client"
	apiV1Prefix = "api/v1/"
)

type ClientOptionFunc func(*Client) error
type Client struct {
	client     *gorequest.SuperAgent
	baseURL    *url.URL
	apiVersion string
	UserAgent  string
	token      string
	username   string
	password   string
	// services
	Accounts     *AccountsService
	Sessions     *SessionsService
	Applications *ApplicationService
	Clusters     *ClusterService
	Projects     *ProjectService
	Repositories *RepositoriesService
	RepoCreds    *RepoCredsService
}

func (c *Client) ErrsWrapper(errs []error) error {
	s := ""
	if len(errs) == 0 {
		return nil
	}
	for _, e := range errs {
		s += e.Error() + "\\n"
	}
	return errors.New(s)
}
func (c *Client) Init() (err error) {
	if len(c.token) == 0 {
		if token, _, err := c.Sessions.CreateUserJWT(); err == nil {
			c.token = token.Token
		} else {
			return
		}
	}
	if len(c.token) == 0 {
		err = errors.New("client token is empty")
	}
	return
}
func NewClient(baseUrl, username, password, token string, options ...ClientOptionFunc) (client *Client, err error) {
	client, err = newClient(baseUrl, username, password, token, options...)
	if err != nil {
		klog.Error(err)
		return
	}
	return
}
func newClient(baseUrl, username, password, token string, options ...ClientOptionFunc) (client *Client, err error) {
	client = &Client{}
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}

	baseURL, err := url.Parse(baseUrl)
	if err != nil {
		return
	}
	client.baseURL = baseURL
	var harborClient *gorequest.SuperAgent
	if harborClient == nil {
		harborClient = gorequest.New()
	}
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(client); err != nil {
			return nil, err
		}
	}
	client.client = harborClient
	client.UserAgent = userAgent
	client.username = username
	client.password = password
	client.token = token
	if strings.HasPrefix(client.baseURL.String(), "https") {
		client.client.TLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}
	client.Accounts = &AccountsService{client: client}
	client.Sessions = &SessionsService{client: client}
	client.Applications = &ApplicationService{client: client}
	client.Clusters = &ClusterService{client: client}
	client.Projects = &ProjectService{client: client}
	client.Repositories = &RepositoriesService{client: client}
	client.RepoCreds = &RepoCredsService{client: client}
	return
}
func (c *Client) newRequest(method, subPath string) *gorequest.SuperAgent {
	var u string
	h := c.client.Set("Accept", "application/json")
	if c.UserAgent != "" {
		h.Set("User-Agent", c.UserAgent)
	}
	if len(c.token) > 0 {
		if !strings.HasPrefix(c.token, "Bearer ") {
			c.token = fmt.Sprintf("Bearer %s", c.token)
		}
		c.client.Set("Authorization", c.token)
	}
	u = c.baseURL.String() + subPath
	switch method {
	case gorequest.PUT:
		return c.client.Put(u).Set("Content-Type", "application/json")
	case gorequest.POST:
		return c.client.Post(u).Set("Content-Type", "application/json")
	case gorequest.GET:
		return c.client.Get(u)
	case gorequest.HEAD:
		return c.client.Head(u)
	case gorequest.DELETE:
		return c.client.Delete(u)
	case gorequest.PATCH:
		return c.client.Patch(u)
	case gorequest.OPTIONS:
		return c.client.Options(u)
	default:
		return c.client.Get(u)
	}
}
