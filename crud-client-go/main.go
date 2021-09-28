package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var clientset *kubernetes.Clientset

func main()  {
	CreateDeployment(clientset)
	//ListPods(clientset)
	ListDeployments(clientset)
	//DeletePod(clientset)
	CreateService(clientset)
	fmt.Println("Server is listening on ", GetNodeIp(clientset),":", GetServiceNodePort(clientset))

}
func init()  {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		fmt.Println(home)
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	fmt.Println(*kubeconfig)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("init() completed.")

}