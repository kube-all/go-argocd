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
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/ghodss/yaml"
	"testing"
)

func TestApplicationList(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	apps, _, err := client.Applications.List(application.ApplicationQuery{})
	if err != nil {
		t.Fatal(err)
	}
	for _, app := range apps.Items {
		t.Logf("project:%s app name: %s", app.Spec.Project, app.Name)

	}
}

func TestApplicationCreate(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	yamlData := `
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: guestbook-api-test-2
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://gitee.com/itcloudy/argocd-example-apps.git
    targetRevision: HEAD
    path: guestbook
  destination:
    server: https://kubernetes.default.svc
    namespace: guestbook`
	client.Init()
	var app v1alpha1.Application
	yaml.Unmarshal([]byte(yamlData), &app)
	appRe, _, err := client.Applications.Create(application.ApplicationCreateRequest{
		Application: &app,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(appRe)

}

func TestApplicationUpdate(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	yamlData := `
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: guestbook-api-test
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/argoproj/argocd-example-apps.git
    targetRevision: HEAD
    path: guestbook
  destination:
    server: https://kubernetes.default.svc
    namespace: guestbook-update
`
	client.Init()
	var app v1alpha1.Application
	yaml.Unmarshal([]byte(yamlData), &app)
	_, _, err = client.Applications.Update(application.ApplicationUpdateRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)

}
func TestApplicationManagedResources(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}

	client.Init()
	app, _, err := client.Applications.ManagedResources(application.ResourcesQuery{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}
func TestListResourceActions(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	app, _, err := client.Applications.ListResourceActions(ApplicationResourceRequest{
		Name:         "guestbook-api-test",
		Group:        "apps",
		Kind:         "Deployment",
		Version:      "v1",
		Namespace:    "ingress-nginx",
		ResourceName: "guestbook-ui",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}
