package main

import (
	"os"

	"github.com/azenakhi/kube-operator/internal/services"
	"github.com/azenakhi/kube-operator/internal/utils"
)

func main() {
	kubeconfig, err := utils.GetClientConfig()
	if err != nil {
		utils.Log.Error().Msgf("Couldn't load K8S client config:", err)
		os.Exit(1)
	}

	managePods, err := services.NewManagePods(kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	managePods.ListPod()
}
