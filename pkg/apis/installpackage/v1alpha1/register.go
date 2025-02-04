// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: "install.package.carvel.dev", Version: "v1alpha1"}

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(SchemeGroupVersion, &PackageRepository{}, &PackageRepositoryList{})
		scheme.AddKnownTypes(SchemeGroupVersion, &InstalledPackage{}, &InstalledPackageList{})
		scheme.AddKnownTypes(SchemeGroupVersion, &InternalPackage{}, &InternalPackageList{})
		scheme.AddKnownTypes(SchemeGroupVersion, &metav1.Status{})
		metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
		return nil
	})
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
