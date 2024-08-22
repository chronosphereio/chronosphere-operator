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

// ChronosphereRecordingRuleSpec defines the desired state of ChronosphereRecordingRule
type ChronosphereRecordingRuleSpec models.Configv1RecordingRule

// ChronosphereRecordingRuleStatus defines the observed state of ChronosphereRecordingRule
type ChronosphereRecordingRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State                   string `json:"state,omitempty"`
	Message                 string `json:"message,omitempty"`
	LastProcessedGeneration int64  `json:"lastProcessedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ChronosphereRecordingRule is the Schema for the chronosphererecordingrules API
type ChronosphereRecordingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChronosphereRecordingRuleSpec   `json:"spec,omitempty"`
	Status ChronosphereRecordingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChronosphereRecordingRuleList contains a list of ChronosphereRecordingRule
type ChronosphereRecordingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronosphereRecordingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronosphereRecordingRule{}, &ChronosphereRecordingRuleList{})
}
