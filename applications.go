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
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/reposerver/apiclient"
	"github.com/parnurzeal/gorequest"
	v1 "k8s.io/api/core/v1"
	"net/http"
	"strings"
)

type ApplicationResourceRequest struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	ResourceName string `json:"resourceName"`
	Version      string `json:"version"`
	Group        string `json:"group"`
	Kind         string `json:"kind"`
}

type ApplicationService struct {
	client *Client
}

//List returns list of applications
func (s *ApplicationService) List(request application.ApplicationQuery) (result v1alpha1.ApplicationList, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Create creates an application
func (s *ApplicationService) Create(request application.ApplicationCreateRequest) (result v1alpha1.Application, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	queryMap := make(map[string]bool)
	if request.Upsert != nil {
		queryMap["upsert"] = *request.Upsert
	}
	if request.Validate != nil {
		queryMap["validate"] = *request.Validate
	}
	resp, data, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"applications").
		SendStruct(request.Application).
		Query(&queryMap).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//ManagedResources returns list of managed resources
func (s *ApplicationService) ManagedResources(request application.ResourcesQuery) (results []*v1alpha1.ResourceDiff, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+*request.ApplicationName+"/managed-resources").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &results)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//ResourceTree returns resource tree
func (s *ApplicationService) ResourceTree(request application.ResourcesQuery) (result v1alpha1.ApplicationTree, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+*request.ApplicationName+"/resource-tree").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Get returns an application by name
func (s *ApplicationService) Get(request application.ApplicationQuery) (result v1alpha1.Application, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+*request.Name+"/resource-tree").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Update updates an application
func (s *ApplicationService) Update(request application.ApplicationUpdateRequest) (result v1alpha1.Application, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.PUT, apiV1Prefix+"applications/"+request.Application.Name).
		SendStruct(request.Application).
		Query(fmt.Sprintf("validate=%t", *request.Validate)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Patch patch an application
func (s *ApplicationService) Patch(request application.ApplicationPatchRequest) (results []*v1alpha1.ResourceDiff, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.PATCH, apiV1Prefix+"applications/"+*request.Name).
		SendStruct(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &results)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//ListResourceEvents returns a list of event resources
func (s *ApplicationService) ListResourceEvents(request application.ApplicationResourceEventsQuery) (result v1.EventList, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+*request.Name+"/events").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//ApplicationPodLogs returns stream of log entries for the specified pod. Pod
func (s *ApplicationService) ApplicationPodLogs(request application.ApplicationPodLogsQuery) (result application.LogEntry, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+*request.Name+"/logs").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//GetManifests returns application manifests
func (s *ApplicationService) GetManifests(name, revision string) (result apiclient.ManifestResponse, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+name+"/manifests").
		Query(fmt.Sprintf("revision=%s", revision)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//TerminateOperation terminates the currently running operation
func (s *ApplicationService) TerminateOperation(name string) (success bool, resp gorequest.Response, err error) {

	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.DELETE, apiV1Prefix+"applications/"+name+"/operation").
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//PodLogs returns stream of log entries for the specified pod. Pod
func (s *ApplicationService) PodLogs(request application.ApplicationPodLogsQuery) (result application.LogEntry, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+*request.Name+"/pods/"+*request.PodName+"/logs").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//GetResource returns single application resource
func (s *ApplicationService) GetResource(request ApplicationResourceRequest) (result application.ApplicationResourceResponse, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+request.Name+"/resource").
		Query(&request).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//
////PatchResource patch single application resource
//func (s *ApplicationService) PatchResource(request application.ApplicationResourcePatchRequest) (result application.ApplicationResourceResponse, resp gorequest.Response, err error) {
//	var (
//		data string
//		errs []error
//	)
//	resp, data, errs = s.client.
//		newRequest(gorequest.POST, apiV1Prefix+"applications/"+*request.Name+"/resource").
//		SendString(request.String()).
//		Query(&request).
//		End()
//	if resp.StatusCode == http.StatusOK {
//		_ = json.Unmarshal([]byte(data), &result)
//	} else {
//		_ = json.Unmarshal([]byte(data), &err)
//	}
//	if len(errs) > 0 {
//		err.Code = -1
//		err.UpExpectErrs = errs
//	}
//	return
//}

//ListResourceActions returns list of resource actions
func (s *ApplicationService) ListResourceActions(request ApplicationResourceRequest) (result application.ResourceActionsListResponse, resp gorequest.Response, err error) {
	var (
		data    string
		errs    []error
		queries []string
	)
	queryMap := make(map[string]string)
	queryMap["name"] = request.Name
	queryMap["namespace"] = request.Namespace
	queryMap["resourceName"] = request.ResourceName
	queryMap["version"] = request.Version
	queryMap["group"] = request.Group
	queryMap["kind"] = request.Kind
	for k, v := range queryMap {
		queries = append(queries, fmt.Sprintf("%s=%s", k, v))
	}
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"applications/"+request.Name+"/resource/actions").
		Query(strings.Join(queries, "&")).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}
