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
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

type RepoCredsService struct {
	client *Client
}

//ListRepositoryCredentials gets a list of all configured repository credential sets
func (s *RepoCredsService) ListRepositoryCredentials(url string) (result v1alpha1.RepoCredsList, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"projects").
		Query(fmt.Sprintf("url=%s", url)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}
