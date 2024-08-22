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

// ChronosphereClassicDashboardSpec should be the same as models.Configv1GrafanaDashboard
type ChronosphereClassicDashboardSpec models.Configv1GrafanaDashboard

// ChronosphereClassicDashboardStatus defines the observed state of ChronosphereClassicDashboard
type ChronosphereClassicDashboardStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State                   string `json:"state,omitempty"`
	Message                 string `json:"message,omitempty"`
	LastProcessedGeneration int64  `json:"lastProcessedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ChronosphereClassicDashboard is the Schema for the chronosphereclassicdashboards API
type ChronosphereClassicDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//Spec   ChronosphereClassicDashboardSpec   `json:"spec,omitempty"`
	Spec   ChronosphereClassicDashboardSpec   `json:"spec,omitempty"`
	Status ChronosphereClassicDashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChronosphereClassicDashboardList contains a list of ChronosphereClassicDashboard
type ChronosphereClassicDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronosphereClassicDashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronosphereClassicDashboard{}, &ChronosphereClassicDashboardList{})
}
