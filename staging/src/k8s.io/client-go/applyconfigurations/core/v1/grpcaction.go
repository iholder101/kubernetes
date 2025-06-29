/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GRPCActionApplyConfiguration represents a declarative configuration of the GRPCAction type for use
// with apply.
type GRPCActionApplyConfiguration struct {
	Port    *int32  `json:"port,omitempty"`
	Service *string `json:"service,omitempty"`
}

// GRPCActionApplyConfiguration constructs a declarative configuration of the GRPCAction type for use with
// apply.
func GRPCAction() *GRPCActionApplyConfiguration {
	return &GRPCActionApplyConfiguration{}
}
func (b GRPCActionApplyConfiguration) IsApplyConfiguration() {}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *GRPCActionApplyConfiguration) WithPort(value int32) *GRPCActionApplyConfiguration {
	b.Port = &value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *GRPCActionApplyConfiguration) WithService(value string) *GRPCActionApplyConfiguration {
	b.Service = &value
	return b
}
