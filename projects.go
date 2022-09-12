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

type ProjectService struct {
	client *Client
}

//List returns list of projects
func (s *ProjectService) List(name string) (result v1alpha1.AppProjectList, resp gorequest.Response, err error) {
	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"projects").
		Query(fmt.Sprintf("name=%s", name)).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Create creates an project
func (s *ProjectService) Create(project v1alpha1.AppProject, upsert bool) (result v1alpha1.AppProject, resp gorequest.Response, err error) {
	sendMap := make(map[string]interface{})
	sendMap["upsert"] = upsert
	sendMap["project"] = project

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.POST, apiV1Prefix+"projects").
		SendStruct(&sendMap).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Get returns a project by server address
func (s *ProjectService) Get(name string) (result v1alpha1.AppProject, resp gorequest.Response, err error) {

	var (
		data string
		errs []error
	)
	resp, data, errs = s.client.
		newRequest(gorequest.GET, apiV1Prefix+"project/"+name).
		End()
	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal([]byte(data), &result)
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//Delete deletes a project
func (s *ProjectService) Delete(name string) (success bool, resp gorequest.Response, err error) {

	var (
		errs []error
	)
	resp, _, errs = s.client.
		newRequest(gorequest.DELETE, apiV1Prefix+"projects/"+name).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	err = s.client.ErrsWrapper(errs)
	return
}

//GetDetailedProject returns a project that include project, global project and scoped resources by name
func (s *ProjectService) GetDetailedProject(name string) (success bool, resp gorequest.Response, err error) {
	return
}
