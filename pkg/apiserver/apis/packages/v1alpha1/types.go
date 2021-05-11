// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	kcv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Package struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PackageSpec `json:"spec"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PackageVersion struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PackageVersionSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PackageVersionList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []PackageVersion `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PackageList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Package `json:"items"`
}

type PackageVersionSpec struct {
	PackageName                     string   `json:"packageName,omitempty"`
	Version                         string   `json:"version,omitempty"`
	Licenses                        []string `json:"licenses,omitempty"`
	ReleasedAt                      string   `json:"releasedAt,omitempty"`
	CapactiyRequirementsDescription string   `json:"capacityRequirementsDescription,omitempty"`
	ReleaseNotes                    string   `json:"releaseNotes,omitempty"`

	Template AppTemplateSpec `json:"template,omitempty"`
	// TODO ValuesSchema
}

type PackageSpec struct {
	DisplayName        string       `json:"displayName,omitempty"`
	LongDescription    string       `json:"longDescription,omitempty"`
	ShortDescription   string       `json:"shortDescription,omitempty"`
	IconSVGBase64      string       `json:"iconSVGBase64,omitempty"`
	ProviderName       string       `json:"providerName,omitempty"`
	Maintainers        []Maintainer `json:"maintainers,omitempty"`
	Categories         []string     `json:"categories,omitempty"`
	SupportDescription string       `json:"supportDescription,omitempty"`
}

type Maintainer struct {
	Name string `json:"name,omitempty"`
}

type AppTemplateSpec struct {
	Spec *kcv1alpha1.AppSpec `json:"spec"`
}
