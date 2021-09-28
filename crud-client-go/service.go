package main

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

func CreateService(clientset *kubernetes.Clientset)  {
	ctx := context.TODO()
	svcClient := clientset.CoreV1().Services(metav1.NamespaceDefault)

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
					NodePort: int32(30033),
					Port: 2345,
					TargetPort: intstr.IntOrString{
						IntVal: 80,
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

func GetServiceNodePort(clientset *kubernetes.Clientset)  int32{
	svcClient := clientset.CoreV1().Services(metav1.NamespaceDefault)

	svc, _ := svcClient.Get(context.TODO(), "mysvc", metav1.GetOptions{})

	svcPorts := svc.Spec.Ports

	for _, p := range svcPorts{
		return p.NodePort
	}
	return int32(0)
}
