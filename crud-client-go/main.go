package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"time"
)

var clientset *kubernetes.Clientset

func main()  {
	deploymentClient := DeploymentClient{DeploymentInterface: clientset.AppsV1().Deployments("default")}
	svcClient := ServiceClient{ServiceInterface: clientset.CoreV1().Services("default")}

	if deploymentClient.IsPresent("depl"){
		deploymentClient.DeleteDeployment()
	}
	svcClient.DeleteService()
	time.Sleep(time.Second * 5)

	deploymentClient.CreateDeployment()
	//ListPods(clientset)
	deploymentClient.ListDeployments()
	//DeletePod(clientset)
	svcClient.CreateService()
	fmt.Println("Server is listening on ", GetNodeIp(clientset),":", svcClient.GetServiceNodePort())

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