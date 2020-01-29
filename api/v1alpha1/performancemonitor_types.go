package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PerformanceMonitorSpec defines the desired state of PerformanceMonitor
type PerformanceMonitorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Simple collection of selectors (aka labels)
	ServiceSelector map[string]string `json:"serviceSelector,omitempty"`
}

// PerformanceMonitorStatus defines the observed state of PerformanceMonitor
type PerformanceMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerformanceMonitor is the Schema for the performancemonitors API
type PerformanceMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerformanceMonitorSpec   `json:"spec,omitempty"`
	Status PerformanceMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerformanceMonitorList contains a list of PerformanceMonitor
type PerformanceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerformanceMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerformanceMonitor{}, &PerformanceMonitorList{})
}
