package ingress

import (
	"fmt"
	"github.com/bdehri/kubectl-map/cmd/common"
	"github.com/spf13/cobra"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Options struct {
	common.NetworkPolicyClient
	Pod []string
}

func NewIngressCommand() *cobra.Command {
	options := &Options{}

	cmd := &cobra.Command{
		Use:   "ingress",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			kubeconfig, err := cmd.Flags().GetString("kubeconfig")
			if err != nil {

			}

			err := options.InitClient(kubeconfig)
			if err != nil {

			}
			netPolList, err := options.NetworkPolicyClient.List(options.Clientset)
			if err != nil {

			}
			processedList := make(networkv1.NetworkPolicyList.Items)
			for _,i := range netPolList.Items {
				if
			}
		},
	}
	cmd.Flags().StringArrayVarP(&options.Pod, "pod", "p", make([]string,0), "Pod Name to be selected")
	cmd.MarkFlagRequired("pod")
	common.AddClientFlags(cmd, options.ClientOptions)
	return cmd
}

func parseStringIntoMap(str []string)(map[string]string,error){
	labelSelector := make(map[string]string)
	for _,i := range str{
		keyVal := strings.Split(i,"=")
		labelSelector[keyVal[0]]=keyVal[1]
	}
}
