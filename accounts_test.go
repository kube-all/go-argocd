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
	"github.com/kube-all/go-argocd/models"
	"testing"
	"time"
)

func TestListAccount(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	accounts, _, err := client.Accounts.ListAccounts()
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range accounts.Items {
		t.Logf("%v", a)
	}
}
func TestCanI(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	yes, _, err := client.Accounts.CanI(models.CanIRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if yes.Value == "yes" {
		t.Logf("you can do ")
	} else {
		t.Logf("you can't do")
	}
}

func TestUpdatePassword(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	yes, _, err := client.Accounts.UpdatePassword(models.UpdatePasswordRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if yes {
		t.Logf("you can do")
	} else {
		t.Logf("you can't do")
	}
}

func TestGetAccount(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	account, _, err := client.Accounts.GetAccount("admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("account info: %v", account)
}

func TestCreateToken(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	token, _, err := client.Accounts.CreateToken("alice", "api-create", time.Now().Add(10*time.Minute).Unix())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token info: %v", token)
}

func TestDeleteToken(t *testing.T) {
	client, err := NewClient(TestAddress, TestUsername, TestPwd, "")
	if err != nil {
		t.Fatal(err)
	}
	client.Init()
	ok, _, err := client.Accounts.DeleteToken("alice", "api-create")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("delete token: %v", ok)
}
