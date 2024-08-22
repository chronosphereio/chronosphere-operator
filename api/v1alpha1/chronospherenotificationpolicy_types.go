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

// ChronosphereNotificationPolicySpec defines the desired state of ChronosphereNotificationPolicy
type ChronosphereNotificationPolicySpec models.Configv1NotificationPolicy

// ChronosphereNotificationPolicyStatus defines the observed state of ChronosphereNotificationPolicy
type ChronosphereNotificationPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State                   string `json:"state,omitempty"`
	Message                 string `json:"message,omitempty"`
	LastProcessedGeneration int64  `json:"lastProcessedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ChronosphereNotificationPolicy is the Schema for the chronospherenotificationpolicies API
type ChronosphereNotificationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChronosphereNotificationPolicySpec   `json:"spec,omitempty"`
	Status ChronosphereNotificationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChronosphereNotificationPolicyList contains a list of ChronosphereNotificationPolicy
type ChronosphereNotificationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronosphereNotificationPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronosphereNotificationPolicy{}, &ChronosphereNotificationPolicyList{})
}
