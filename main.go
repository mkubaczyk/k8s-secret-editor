package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mkubaczyk/k8s-secret-editor/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	settings.Init()
	settings.Logger.Info("Starting...")

	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secrets, err := clientset.CoreV1().Secrets("kube-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"secrets": secrets.Items,
		})
	})
	router.Run(":8080")
}
