package services

import (
	"fmt"

	"github.com/azenakhi/kube-operator/internal/utils"
	"github.com/rs/zerolog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ManagePods struct {
	logger     zerolog.Logger
	kubeClient *kubernetes.Clientset
}

func NewManagePods(kubeconfig *rest.Config) (*ManagePods, error) {
	logger := utils.Log.With().Str("service", "managePods").Logger()

	kubeClient, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		logger.Error().Msgf("Could not create argocd client: %v", err)
		return nil, err
	}

	return &ManagePods{
		logger:     logger,
		kubeClient: kubeClient,
	}, nil
}

func (p *ManagePods) ListPod() {
	api := p.kubeClient.CoreV1()
	pods, err := api.Pods("").List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Printf("%s \n", pod.GetName())
	}
}
