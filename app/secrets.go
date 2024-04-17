package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var namespace *string

func getNamespace() string {
	if namespace == nil {
		ns := os.Getenv("NAMESPACE")
		if ns == "" {
			namespaceBytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
			if err != nil {
				panic(err)
			}
			ns = string(namespaceBytes)
		}
		namespace = &ns
	}
	return *namespace
}

func getSecret(clientset *kubernetes.Clientset) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		secret, err := clientset.CoreV1().Secrets(getNamespace()).Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"secret": secret.Data})
	}
}
