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
	repositorypkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/repository"
	"testing"
)

func TestRepositoryList(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	pros, _, err := client.Repositories.ListRepositories(repositorypkg.RepoQuery{})
	if err != nil {
		t.Fatal(err)
	}
	for _, app := range pros.Items {
		t.Logf("project: %s", app.Name)
	}
}
