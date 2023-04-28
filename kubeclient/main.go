package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	config, err := clientcmd.BuildConfigFromFlags("https://api.crc.testing:6443", filepath.Join(dir, ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, namespace := range namespaces.Items {
		log.Println(namespace.Name)
	}
}
