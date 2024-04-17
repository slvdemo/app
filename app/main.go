package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	router.GET("/secret/:name", getSecret(clientset))

	err = router.Run(":" + fmt.Sprintf("%d", 8888))
	if err != nil {
		panic(err)
	}
}
