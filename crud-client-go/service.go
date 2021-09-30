package main

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

type ServiceClient struct {
	v1.ServiceInterface
}

func (svcClient *ServiceClient) CreateService()  {
	ctx := context.TODO()

	svc := coreV1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mysvc",
		},
		Spec: coreV1.ServiceSpec{
			Selector: map[string]string{
				"app": "demo",
			},
			Type: coreV1.ServiceTypeNodePort,
			Ports: []coreV1.ServicePort{
				{
					NodePort: int32(30012),
					Port: 2345,
					TargetPort: intstr.IntOrString{
						IntVal: 8080,
					},
				},
			},
		},
	}
	_, err:= svcClient.Create(ctx,&svc, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	}
}

func (svcClient *ServiceClient) GetServiceNodePort()  int32{
	svc, _ := svcClient.Get(context.TODO(), "mysvc", metav1.GetOptions{})

	svcPorts := svc.Spec.Ports

	for _, p := range svcPorts{
		return p.NodePort
	}
	return int32(0)
}

func (svcClient *ServiceClient) DeleteService(){

	deletePolicy := metav1.DeletePropagationForeground
	err := svcClient.Delete(context.TODO(), "mysvc", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("service successfully Deleted")
}
