package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pervcheckv1alpha1 "github.com/daemonfire300/pervcheck/api/v1alpha1"
)

// PerformanceMonitorReconciler reconciles a PerformanceMonitor object
type PerformanceMonitorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=pervcheck.k8s.accountr.eu,resources=performancemonitors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pervcheck.k8s.accountr.eu,resources=performancemonitors/status,verbs=get;update;patch

func (r *PerformanceMonitorReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("performancemonitor", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PerformanceMonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pervcheckv1alpha1.PerformanceMonitor{}).
		Complete(r)
}
