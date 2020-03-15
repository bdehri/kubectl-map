package common

import (
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

const (
	kubeconfigEnvVar = "KUBECONFIG"
)

type ClientOptions struct {
	Namespace  string
	Kubeconfig string
	*kubernetes.Clientset
}

func AddClientFlags(command *cobra.Command, options ClientOptions) {
	command.Flags().StringVar(&options.Kubeconfig, "kubeconfig", "", "path to kubeconfig")
	command.Flags().StringVarP(&options.Namespace, "namespace", "n", "default", "namespace to be searched")
}

func buildConfig(kubeconf string) (*rest.Config, error) {
	kubeconfigEnv := os.Getenv(kubeconfigEnvVar)
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	if kubeconfigEnv != "" {
		kubeconfig = kubeconfigEnv
	}

	if kubeconf != "" {
		kubeconfig = kubeconf
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		return nil, err
	}

	return config, nil

}

func (clientOptions *ClientOptions) InitClient(kubeconf string) error {
	config, err := buildConfig(kubeconf)

	if err != nil {
		return err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	clientOptions.Clientset = clientSet
	return nil
}
