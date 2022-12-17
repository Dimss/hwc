package controller

import (
	"context"
	"github.com/Dimss/hwc/apis/hwc.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type HelloWorldReconciler struct {
	client.Client
	*runtime.Scheme
}

func (r *HelloWorldReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	hw := v1alpha1.HelloWorld{}
	name := types.NamespacedName{Namespace: req.Namespace, Name: req.Name}
	if err := r.Get(context.Background(), name, &hw); err != nil {
		log.Printf("error: %v\n", err.Error())
		return ctrl.Result{}, err
	}
	log.Printf("Message: %s", hw.Spec.Message)
	return ctrl.Result{}, nil
}
