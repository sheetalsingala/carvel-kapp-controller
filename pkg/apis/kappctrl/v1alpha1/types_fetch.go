// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// +k8s:openapi-gen=true
type AppFetch struct {
	Inline       *AppFetchInline       `json:"inline,omitempty"`
	Image        *AppFetchImage        `json:"image,omitempty"`
	HTTP         *AppFetchHTTP         `json:"http,omitempty"`
	Git          *AppFetchGit          `json:"git,omitempty"`
	HelmChart    *AppFetchHelmChart    `json:"helmChart,omitempty"`
	ImgpkgBundle *AppFetchImgpkgBundle `json:"imgpkgBundle,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchInline struct {
	Paths     map[string]string      `json:"paths,omitempty"`
	PathsFrom []AppFetchInlineSource `json:"pathsFrom,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchInlineSource struct {
	SecretRef    *AppFetchInlineSourceRef `json:"secretRef,omitempty"`
	ConfigMapRef *AppFetchInlineSourceRef `json:"configMapRef,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchInlineSourceRef struct {
	DirectoryPath               string `json:"directoryPath,omitempty"`
	corev1.LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
}

// +k8s:openapi-gen=true
type AppFetchImage struct {
	// Example: username/app1-config:v0.1.0
	URL string `json:"url,omitempty"`
	// Secret may include one or more keys: username, password, token.
	// By default anonymous access is used for authentication.
	// TODO support docker config formated secret
	// +optional
	SecretRef *AppFetchLocalRef `json:"secretRef,omitempty"`
	// +optional
	SubPath string `json:"subPath,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchHTTP struct {
	// URL can point to one of following formats: text, tgz, zip
	URL string `json:"url,omitempty"`
	// +optional
	SHA256 string `json:"sha256,omitempty"`
	// Secret may include one or more keys: username, password
	// +optional
	SecretRef *AppFetchLocalRef `json:"secretRef,omitempty"`
	// +optional
	SubPath string `json:"subPath,omitempty"`
}

// TODO implement git
// +k8s:openapi-gen=true
type AppFetchGit struct {
	URL string `json:"url,omitempty"`
	// +optional
	Ref string `json:"ref,omitempty"`
	// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts
	// +optional
	SecretRef *AppFetchLocalRef `json:"secretRef,omitempty"`
	// +optional
	SubPath string `json:"subPath,omitempty"`
	// +optional
	LFSSkipSmudge bool `json:"lfsSkipSmudge,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchHelmChart struct {
	// Example: stable/redis
	Name string `json:"name,omitempty"`
	// +optional
	Version    string                 `json:"version,omitempty"`
	Repository *AppFetchHelmChartRepo `json:"repository,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchHelmChartRepo struct {
	URL string `json:"url,omitempty"`
	// +optional
	SecretRef *AppFetchLocalRef `json:"secretRef,omitempty"`
}

// +k8s:openapi-gen=true
type AppFetchLocalRef struct {
	// Object is expected to be within same namespace
	corev1.LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
}

// +k8s:openapi-gen=true
type AppFetchImgpkgBundle struct {
	Image string `json:"image,omitempty"`
	// Secret may include one or more keys: username, password, token.
	// By default anonymous access is used for authentication.
	// TODO support docker config formated secret
	// +optional
	SecretRef *AppFetchLocalRef `json:"secretRef,omitempty"`
}
