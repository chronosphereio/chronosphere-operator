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

package controller

import (
	"chronosphere.io/chronosphere-operator/chronosphere_client/grafana_dashboard"
	"chronosphere.io/chronosphere-operator/models"
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strings"

	"github.com/jinzhu/copier"

	monitoringv1alpha1 "chronosphere.io/chronosphere-operator/api/v1alpha1"
)

// ChronosphereClassicDashboardReconciler reconciles a ChronosphereClassicDashboard object
type ChronosphereClassicDashboardReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=monitoring.chronosphere.io,resources=chronosphereclassicdashboards,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitoring.chronosphere.io,resources=chronosphereclassicdashboards/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=monitoring.chronosphere.io,resources=chronosphereclassicdashboards/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ChronosphereClassicDashboard object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.2/pkg/reconcile
func (r *ChronosphereClassicDashboardReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var obj monitoringv1alpha1.ChronosphereClassicDashboard
	if err := r.Get(ctx, req.NamespacedName, &obj); err != nil {
		obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
			State:   "Failed",
			Message: "Failed to get ChronosphereClassicDashboard with name: " + req.Name + " in namespace: " + req.Namespace,
		}
		_ = r.Status().Update(ctx, &obj)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	myFinalizerName := "finalizer.monitoring.chronosphere.io"
	if obj.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// to registering our finalizer.
		if !controllerutil.ContainsFinalizer(&obj, myFinalizerName) {
			controllerutil.AddFinalizer(&obj, myFinalizerName)
			if err := r.Update(ctx, &obj); err != nil {
				return ctrl.Result{}, err
			}

			if err := r.createExternalDependency(&obj, ctx); err != nil {
				return ctrl.Result{}, err
			}
		} else {
			// The object is being updated
			if err := r.createExternalDependency(&obj, ctx); err != nil {
				return ctrl.Result{}, err
			}
		}

	} else {
		// The object is being deleted
		if controllerutil.ContainsFinalizer(&obj, myFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.deleteExternalDependency(&obj, ctx); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried.
				return ctrl.Result{}, err
			}
			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(&obj, myFinalizerName)
			if err := r.Update(ctx, &obj); err != nil {
				return ctrl.Result{}, err
			}
		}
		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ChronosphereClassicDashboardReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringv1alpha1.ChronosphereClassicDashboard{}).
		Complete(r)
}

func (r *ChronosphereClassicDashboardReconciler) deleteExternalDependency(obj *monitoringv1alpha1.ChronosphereClassicDashboard, ctx context.Context) error {
	logger := log.FromContext(ctx)

	params := grafana_dashboard.DeleteGrafanaDashboardParams{
		Slug: obj.Spec.Slug,
	}

	chronosphereClient, err := GetChronosphereClient()
	if err != nil {
		logger.Info("Failed to get ChronosphereClient", "error", err)
		obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
			State:   "Failed",
			Message: "Failed to get ChronosphereClient with slug: " + obj.Spec.Slug,
		}
		_ = r.Status().Update(ctx, obj)
		return err
	}
	resp, err := chronosphereClient.GrafanaDashboard.DeleteGrafanaDashboard(&params)
	if err != nil {
		if strings.Contains(err.Error(), "code=NOT_FOUND") {
			logger.Info("Already deleted ChronosphereClassicDashboard", "error", err, "response", resp)
			obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
				State:   "Success",
				Message: "Already deleted ChronosphereClassicDashboard with slug: " + obj.Spec.Slug,
			}
			_ = r.Status().Update(ctx, obj)
			return nil
		}
		logger.Info("Failed to delete ChronosphereClassicDashboard", "error", err, "response", resp)
		obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
			State:   "Failed",
			Message: "Failed to delete ChronosphereClassicDashboard with slug: " + obj.Spec.Slug + " with error: " + err.Error(),
		}
		_ = r.Status().Update(ctx, obj)
		return err
	}

	logger.Info("Successfully deleted ChronosphereClassicDashboard", "response", resp)
	obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
		State:   "Success",
		Message: "Successfully deleted ChronosphereClassicDashboard with slug: " + obj.Spec.Slug,
	}
	_ = r.Status().Update(ctx, obj)

	return nil
}

func (r *ChronosphereClassicDashboardReconciler) createExternalDependency(obj *monitoringv1alpha1.ChronosphereClassicDashboard, ctx context.Context) error {
	logger := log.FromContext(ctx)

	if obj.Status.LastProcessedGeneration == obj.ObjectMeta.Generation {
		// Resource has not changed, no need to update external dependency
		logger.Info("ChronosphereClassicDashboard Resource with slug: " + obj.Spec.Slug + " has not changed, skipping update")
		return nil
	}

	newDashboard := &models.Configv1GrafanaDashboard{}
	err := copier.Copy(&newDashboard, &obj.Spec)
	if err != nil {
		logger.Error(err, "Failed to copy ChronosphereClassicDashboard spec")
		obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
			State:   "Failed",
			Message: "Failed to copy ChronosphereClassicDashboard spec with slug: " + obj.Spec.Slug,
		}
		_ = r.Status().Update(ctx, obj)
		return err
	}

	params := grafana_dashboard.UpdateGrafanaDashboardParams{
		Body: &models.ConfigV1UpdateGrafanaDashboardBody{
			CreateIfMissing:  true,
			DryRun:           false,
			GrafanaDashboard: newDashboard,
		},
		Slug: obj.Spec.Slug,
	}

	chronosphereClient, err := GetChronosphereClient()
	if err != nil {
		logger.Info("Failed to get ChronosphereClient", "error", err)
		obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
			State:   "Failed",
			Message: "Failed to get ChronosphereClient with slug: " + obj.Spec.Slug,
		}
		_ = r.Status().Update(ctx, obj)
		return err
	}
	resp, err := chronosphereClient.GrafanaDashboard.UpdateGrafanaDashboard(&params)
	if err != nil {
		logger.Info("Failed to create ChronosphereClassicDashboard", "error", err, "response", resp)
		obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
			State:   "Failed",
			Message: "Failed to create ChronosphereClassicDashboard with slug: " + obj.Spec.Slug + " with error: " + err.Error(),
		}
		_ = r.Status().Update(ctx, obj)
		return err
	}
	logger.Info("Successfully created ChronosphereClassicDashboard", "response", resp)

	obj.Status = monitoringv1alpha1.ChronosphereClassicDashboardStatus{
		State:                   "Success",
		Message:                 "Successfully created ChronosphereClassicDashboard with slug: " + obj.Spec.Slug,
		LastProcessedGeneration: obj.ObjectMeta.Generation,
	}
	_ = r.Status().Update(ctx, obj)

	return nil
}
