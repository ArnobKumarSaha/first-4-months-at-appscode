package main

import (
	"context"
	"fmt"
	apiappsv1 "k8s.io/api/apps/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
)


type DeploymentClient struct {
	v1.DeploymentInterface
}


func (deploymentClient * DeploymentClient) CreateDeployment()  {
	depObj := apiappsv1.Deployment{
		ObjectMeta : metav1.ObjectMeta{
			Name: "mydepl",
		},
		Spec : apiappsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apicorev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "mypod",
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apicorev1.PodSpec{
					Containers: []apicorev1.Container{
						{
							Name: "mycontainer",
							Image: "arnobkumarsaha/learning_go:latest",
							Ports: []apicorev1.ContainerPort{
								{
									Protocol: apicorev1.ProtocolTCP,
									// There is nothing to do with ContainerPort , as nginx by default listen on port 80.
									// But you have to explicitly specify it.
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
	result, err := deploymentClient.Create(context.TODO(), &depObj, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}


func (deploymentClient * DeploymentClient) ListDeployments()  {
	deps, _ := deploymentClient.List(context.TODO(), metav1.ListOptions{})

	for _, d := range deps.Items{
		fmt.Println(d.Name)
	}
}


func  (deploymentClient * DeploymentClient) UpdateDeployment()  {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deplClient := clientset.AppsV1().Deployments(apicorev1.NamespaceDefault)
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


func  (deploymentClient * DeploymentClient) DeleteDeployment(){
	deletePolicy := metav1.DeletePropagationForeground
	err := deploymentClient.Delete(context.TODO(), "mydepl", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deployment successfully Deleted")
}


func (deploymentClient * DeploymentClient) IsPresent(name string)  bool{
	depl := deploymentClient.GetDeployment("depl")
	if depl.Name != ""{
		return false
	}
	return true
}


func (deploymentClient * DeploymentClient) GetDeployment(name string )  *apiappsv1.Deployment{

	deploy, err := deploymentClient.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		fmt.Println(err)
	}

	return deploy
}


func int32Ptr(i int32) *int32 { return &i }