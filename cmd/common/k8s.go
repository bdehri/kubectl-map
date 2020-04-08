package common

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubernetesClient struct {
	ClientOptions
	*networkv1.NetworkPolicyList
	*corev1.PodList
}

func (kubernetesClient *KubernetesClient) ListNetPols(matchLabels map[string]string) (*networkv1.NetworkPolicyList, error) {
	netPolList := &networkv1.NetworkPolicyList{}
	listOptions := &client.ListOptions{}

	client.InNamespace(kubernetesClient.Namespace).ApplyToList(listOptions)
	client.MatchingLabels(matchLabels).ApplyToList(listOptions)

	err := kubernetesClient.Client.List(context.TODO(), netPolList, listOptions)
	if err != nil {
		return nil, err
	}

	return netPolList, nil
}

func (kubernetesClient *KubernetesClient) ListPods(matchLabels map[string]string) (*corev1.PodList, error) {
	podList := &corev1.PodList{}
	listOptions := &client.ListOptions{}

	client.InNamespace(kubernetesClient.Namespace).ApplyToList(listOptions)
	client.MatchingLabels(matchLabels).ApplyToList(listOptions)

	err := kubernetesClient.Client.List(context.TODO(), podList, listOptions)

	if err != nil {
		return nil, err
	}

	return podList, nil
}

func (kubernetesClient *KubernetesClient) GetPod(podName string) (*corev1.Pod, error) {
	pod := &corev1.Pod{}

	err := kubernetesClient.Client.Get(context.TODO(), client.ObjectKey{
		Namespace: kubernetesClient.Namespace,
		Name:      podName,
	}, pod)

	if err != nil {
		return nil, err
	}

	return pod, nil
}

func (kubernetesClient *KubernetesClient) GetNetPol(podName string) (*networkv1.NetworkPolicy, error) {
	netPol := &networkv1.NetworkPolicy{}

	err := kubernetesClient.Client.Get(context.TODO(), client.ObjectKey{
		Namespace: kubernetesClient.Namespace,
		Name:      podName,
	}, netPol)

	if err != nil {
		return nil, err
	}

	return netPol, nil
}
