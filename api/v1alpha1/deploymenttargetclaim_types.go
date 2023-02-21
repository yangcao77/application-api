/*
Copyright 2021-2022 Red Hat, Inc.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeploymentTargetClaimSpec defines the desired state of DeploymentTargetClaim
type DeploymentTargetClaimSpec struct {
	DeploymentTargetClassName DeploymentTargetClassName `json:"deploymentTargetClassName"`
}

type DeploymentTargetClassName string

const (
	DTCTName_IsolationLevelNamespace DeploymentTargetClassName = "isolation-level-namespace"
)

// DeploymentTargetClaimStatus defines the observed state of DeploymentTargetClaim
type DeploymentTargetClaimStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DeploymentTargetClaim is the Schema for the deploymenttargetclaims API.
// It represents a request for a DeploymentTarget.
type DeploymentTargetClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentTargetClaimSpec   `json:"spec,omitempty"`
	Status DeploymentTargetClaimStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeploymentTargetClaimList contains a list of DeploymentTargetClaim
type DeploymentTargetClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentTargetClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentTargetClaim{}, &DeploymentTargetClaimList{})
}