package manifests

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GCPLBWebhookKubeconfig returns the Secret manifest for the kubeconfig used by
// the GCP load-balancer labels webhook sidecar to authenticate against the
// hosted cluster's kube-apiserver.
func GCPLBWebhookKubeconfig(ns string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "gcp-lb-webhook-kubeconfig",
			Namespace: ns,
		},
	}
}
