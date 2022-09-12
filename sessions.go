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
	"encoding/json"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/session"
	"github.com/parnurzeal/gorequest"

	"net/http"
)

type SessionsService struct {
	client *Client
}

//CreateUserJWT Create a new JWT for authentication and set a cookie if using HTTP
func (s *SessionsService) CreateUserJWT() (token session.SessionResponse, resp gorequest.Response, err error) {
	au := make(map[string]string)
	au["username"] = s.client.username
	au["password"] = s.client.password
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"session").
		SendMap(au).End()

	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &token)
	}
	err = s.client.ErrsWrapper(errs)
	return
}
