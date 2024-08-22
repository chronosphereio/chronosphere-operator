/*
Copyright 2024.

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

package v1alpha1

import (
	"chronosphere.io/chronosphere-operator/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ChronosphereDashboardSpec models.Configv1Dashboard

// ChronosphereDashboardStatus defines the observed state of ChronosphereDashboard
type ChronosphereDashboardStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State                   string `json:"state,omitempty"`
	Message                 string `json:"message,omitempty"`
	LastProcessedGeneration int64  `json:"lastProcessedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ChronosphereDashboard is the Schema for the chronospheredashboards API
type ChronosphereDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChronosphereDashboardSpec   `json:"spec,omitempty"`
	Status ChronosphereDashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChronosphereDashboardList contains a list of ChronosphereDashboard
type ChronosphereDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronosphereDashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronosphereDashboard{}, &ChronosphereDashboardList{})
}
