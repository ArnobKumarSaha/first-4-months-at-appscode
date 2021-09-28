package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

func CreateDeployment(clientset *kubernetes.Clientset)  {
	deplClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)
	depObj := appsv1.Deployment{
		ObjectMeta : metav1.ObjectMeta{
			Name: "mydepl",
		},
		Spec : appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "mypod",
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: "myvol",
							Image: "nginx:latest",
							Ports: []corev1.ContainerPort{
								{
									Protocol: corev1.ProtocolTCP,
									// There is nothing to do with ContainerPort , as nginx by default listen on port 80.
									// But you have to explicitly specify it.
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	result, err := deplClient.Create(context.TODO(), &depObj, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func ListDeployments(clientset *kubernetes.Clientset)  {
	deplClient := clientset.AppsV1().Deployments("")
	deps, _ := deplClient.List(context.TODO(), metav1.ListOptions{})

	for _, d := range deps.Items{
		fmt.Println(d.Name)
	}
}

func UpdateDeployment(clientset *kubernetes.Clientset)  {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deplClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)
		result, getErr := deplClient.Get(context.TODO(), "mydepl", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		_, updateErr := deplClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
}

func int32Ptr(i int32) *int32 { return &i }