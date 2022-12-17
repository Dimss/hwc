package main

import (
	hw1alpha1 "github.com/Dimss/hwc/apis/hwc.dev/v1alpha1"
	"github.com/Dimss/hwc/pkg/controller"
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	//rc, err := config.GetConfig()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//clientset, err := hwcClient.NewForConfig(rc)
	//if err != nil {
	//	log.Printf("error getting clientset %s\n", err)
	//}
	//
	//messages, err := clientset.HwcV1alpha1().HelloWorlds("").List(context.Background(), metav1.ListOptions{})
	//if err != nil {
	//	log.Printf("error: %s\n", err.Error())
	//}
	//for _, hw := range messages.Items {
	//	log.Printf("The message: %s\n", hw.Spec.Message)
	//}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil {
		log.Fatal(err)
	}

	if err := hw1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatal(err)
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(&hw1alpha1.HelloWorld{}).
		Complete(&controller.HelloWorldReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		})
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {

		log.Fatal(err)
	}
}
