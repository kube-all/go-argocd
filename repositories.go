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
	repositorypkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/repository"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/reposerver/apiclient"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

type RepositoriesService struct {
	client *Client
}

//ListRepositories gets a list of all configured repositories
func (s *RepositoriesService) ListRepositories(request repositorypkg.RepoQuery) (repoList v1alpha1.RepositoryList, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"repositories").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &repoList)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//CreateRepository creates a new repository configuration
func (s *RepositoriesService) CreateRepository(request repositorypkg.RepoCreateRequest) (result v1alpha1.Repository, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"repositories").
		SendStruct(request.Repo).
		Query(fmt.Sprintf("upsert=%t&credsOnly=%t", request.Upsert, request.CredsOnly)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//UpdateRepository updates a repository configuration
func (s *RepositoriesService) UpdateRepository(request repositorypkg.RepoUpdateRequest) (result v1alpha1.Repository, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.PUT, apiV1Prefix+"repositories/"+request.Repo.Repo).
		SendStruct(request.Repo).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//GetRepository returns a repository or its credentials
func (s *RepositoriesService) GetRepository(request repositorypkg.RepoQuery) (result v1alpha1.Repository, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"repositories/"+request.Repo).
		Query(fmt.Sprintf("forceRefresh=%t", request.ForceRefresh)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//DeleteRepository deletes a repository from the configuration
func (s *RepositoriesService) DeleteRepository(request repositorypkg.RepoQuery) (result v1alpha1.Repository, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.DELETE, apiV1Prefix+"repositories/"+request.Repo).
		Query(fmt.Sprintf("forceRefresh=%t", request.ForceRefresh)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//ListApps returns list of apps in the repo
func (s *RepositoriesService) ListApps(query repositorypkg.RepoAppsQuery) (result repositorypkg.RepoAppsResponse, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"repositories/"+query.Repo+"/apps").
		Query(&query).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//GetHelmCharts returns list of helm charts in the specified repository
func (s *RepositoriesService) GetHelmCharts(request repositorypkg.RepoQuery) (result apiclient.HelmChartsResponse, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"repositories/"+request.Repo+"/helmcharts").
		Query(fmt.Sprintf("forceRefresh=%t", request.ForceRefresh)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}
func (s *RepositoriesService) ListRefs(request repositorypkg.RepoQuery) (result apiclient.Refs, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"repositories/"+request.Repo+"/refs").
		Query(fmt.Sprintf("forceRefresh=%t", request.ForceRefresh)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//ValidateAccess validates access to a repository with given parameters
func (s *RepositoriesService) ValidateAccess(request repositorypkg.RepoQuery) (result apiclient.Refs, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"repositories/"+request.Repo+"/validate").
		Query(fmt.Sprintf("forceRefresh=%t", request.ForceRefresh)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}
