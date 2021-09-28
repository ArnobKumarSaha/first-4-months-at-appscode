package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPods(clientset *kubernetes.Clientset)  {
	podClient := clientset.CoreV1().Pods("")
	pods, _ := podClient.List(context.TODO(), metav1.ListOptions{})

	for _, p := range pods.Items{
		fmt.Println(p.Name, p.Kind)
	}
}

func DeletePod(clientset *kubernetes.Clientset)  {
	podClient := clientset.CoreV1().Pods(metav1.NamespaceDefault)
	err := podClient.Delete(context.TODO(), "mypod", metav1.DeleteOptions{})

	if err != nil {
		fmt.Println(err)
	}
}

func GetNodeIp(clientset *kubernetes.Clientset) string {
	nodeClient := clientset.CoreV1().Nodes()
	p , _ := nodeClient.Get(context.TODO(), "first-control-plane", metav1.GetOptions{})

	for _, j := range p.Status.Addresses{
		if j.Type == "InternalIP" {
			return j.Address
		}
	}
	return ""
}