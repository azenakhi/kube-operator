package services

import (
	"github.com/azenakhi/kube-operator/internal/utils"
	"github.com/rs/zerolog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type MyService struct {
	logger     zerolog.Logger
	kubeClient *kubernetes.Clientset
}

func NewService(kubeconfig *rest.Config) (*MyService, error) {
	logger := utils.Log.With().Str("service", "myService").Logger()

	kubeClient, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		logger.Error().Msgf("Could not create argocd client: %v", err)
		return nil, err
	}

	return &MyService{
		logger:     logger,
		kubeClient: kubeClient,
	}, nil
}
