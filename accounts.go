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
	"fmt"
	"net/http"

	"github.com/kube-all/go-argocd/models"
	"github.com/parnurzeal/gorequest"
)

type AccountsService struct {
	client *Client
}

//ListAccounts returns the list of accounts
func (s *AccountsService) ListAccounts() (accountList models.AccountsList, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"account").
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &accountList)
	} else {
		_ = json.Unmarshal([]byte(data), &err)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//CanI checks if the current account has permission to perform an action
func (s *AccountsService) CanI(request models.CanIRequest) (response models.CanIResponse, resp gorequest.Response, err error) {
	var p string
	if len(request.Subresource) > 0 {
		p = apiV1Prefix + fmt.Sprintf("account/can-i/%s/%s/%s", request.Resource, request.Action, request.Subresource)
	} else {
		p = apiV1Prefix + fmt.Sprintf("account/can-i/%s/%s", request.Resource, request.Action)
	}
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, p).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &response)
	}
	err = s.client.ErrsWrapper(errs)

	return
}

//UpdatePassword updates an account's password to a new value
func (s *AccountsService) UpdatePassword(request models.UpdatePasswordRequest) (success bool, resp gorequest.Response, err error) {
	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.PUT, apiV1Prefix+"account/password").
		SendMap(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//GetAccount returns an account
func (s *AccountsService) GetAccount(name string) (response models.Account, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"account/"+name).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &response)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//CreateToken creates a token
func (s *AccountsService) CreateToken(name, id string, expiresIn int64) (token models.CreateTokenResponse, resp gorequest.Response, err error) {
	var (
		errs []error
	)
	sendMap := make(map[string]interface{})
	sendMap["name"] = name
	sendMap["id"] = id
	sendMap["expiresIn"] = expiresIn
	resp, _, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+fmt.Sprintf("account/%s/token", name)).
		SendMap(sendMap).
		EndStruct(&token)
	err = s.client.ErrsWrapper(errs)
	return
}

//DeleteToken deletes a token
func (s *AccountsService) DeleteToken(name, id string) (success bool, resp gorequest.Response, err error) {
	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.DELETE, apiV1Prefix+fmt.Sprintf("account/%s/token/%s", name, id)).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}
