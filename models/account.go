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

package models

type AccountsList struct {
	Items []*Account `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

type Account struct {
	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled      bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Capabilities []string `protobuf:"bytes,3,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	Tokens       []*Token `protobuf:"bytes,4,rep,name=tokens,proto3" json:"tokens,omitempty"`
}
type Token struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IssuedAt  int64  `protobuf:"varint,2,opt,name=issuedAt,proto3" json:"issuedAt,omitempty"`
	ExpiresAt int64  `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

type CanIRequest struct {
	Resource    string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Action      string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Subresource string `protobuf:"bytes,3,opt,name=subresource,proto3" json:"subresource,omitempty"`
}

type CanIResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}
type UpdatePasswordRequest struct {
	NewPassword     string `protobuf:"bytes,1,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	CurrentPassword string `protobuf:"bytes,2,opt,name=currentPassword,proto3" json:"currentPassword,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}
type CreateTokenResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}
type DeleteTokenRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}
