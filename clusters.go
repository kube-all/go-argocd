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
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

type ClusterService struct {
	client *Client
}

//List returns list of clusters
func (s *ClusterService) List(request cluster.ClusterQuery) (result v1alpha1.ClusterList, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"clusters").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Create create project
func (s *ClusterService) Create(cluster v1alpha1.Cluster, upsert bool) (result v1alpha1.Cluster, resp gorequest.Response, err error) {
	queryMap := make(map[string]bool)
	queryMap["upsert"] = upsert
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"clusters").
		SendStruct(&cluster).
		Query(&queryMap).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Get returns a cluster by server address
func (s *ClusterService) Get(idValue string, option cluster.ClusterQuery) (result v1alpha1.Cluster, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"clusters/"+idValue).
		SendStruct(&option).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Update updates a cluster
func (s *ClusterService) Update(idValue, idType string, updatedFields []string, cluster v1alpha1.Cluster) (result v1alpha1.Cluster, resp gorequest.Response, err error) {
	sendMap := make(map[string]interface{})
	sendMap["id.type"] = idType
	sendMap["updatedFields"] = updatedFields
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.PUT, apiV1Prefix+"clusters/"+idValue).
		SendStruct(&cluster).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Delete deletes a cluster
func (s *ClusterService) Delete(idValue string, option cluster.ClusterQuery) (success bool, resp gorequest.Response, err error) {

	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.DELETE, apiV1Prefix+"clusters/"+idValue).
		SendStruct(&option).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//InvalidateCache invalidates cluster cache
func (s *ClusterService) InvalidateCache(idValue string) (success bool, resp gorequest.Response, err error) {

	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"clusters/"+idValue+"/invalidate-cache").
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//RotateAuth rotates the bearer token used for a cluster
func (s *ClusterService) RotateAuth(idValue string) (success bool, resp gorequest.Response, err error) {

	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"clusters/"+idValue+"/rotate-auth").
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}
