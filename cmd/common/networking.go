package common

import (
	"context"

	networkv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NetworkPolicyClient struct {
	ClientOptions
	*networkv1.NetworkPolicyList
}

func (networkPolicyClient *NetworkPolicyClient) List(client *kubernetes.Clientset) (*networkv1.NetworkPolicyList, error) {

	networkPolicyList, err := client.NetworkingV1().NetworkPolicies(networkPolicyClient.Namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return networkPolicyList, nil
}
