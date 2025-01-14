// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"

	"github.com/go-logr/logr"
	pkgclient "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/client/clientset/versioned"
	kcclient "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/installedpkg"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type InstalledPkgReconciler struct {
	kcClient  kcclient.Interface
	pkgClient pkgclient.Interface
	log       logr.Logger
}

var _ reconcile.Reconciler = &InstalledPkgReconciler{}

func (r *InstalledPkgReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.log.WithValues("request", request)

	existingInstalledPkg, err := r.kcClient.InstallV1alpha1().InstalledPackages(request.Namespace).Get(ctx, request.Name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Could not find InstalledPkg", "name", request.Name)
			return reconcile.Result{}, nil // No requeue
		}

		log.Error(err, "Could not fetch InstalledPkg")
		return reconcile.Result{}, err
	}

	return installedpkg.NewInstalledPkgCR(existingInstalledPkg, log, r.kcClient, r.pkgClient).Reconcile()
}
