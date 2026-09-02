package pki

import (
	"github.com/openshift/hypershift/support/config"

	corev1 "k8s.io/api/core/v1"
)

// ReconcileGCPLBWebhookServingCert creates/updates the TLS serving certificate
// for the GCP load-balancer labels admission webhook. The webhook listens on
// 127.0.0.1 (loopback, co-located with KAS), so only the IP SAN is required.
func ReconcileGCPLBWebhookServingCert(secret, ca *corev1.Secret, ownerRef config.OwnerRef) error {
	return reconcileSignedCertWithAddresses(secret, ca, ownerRef, "127.0.0.1", nil, X509UsageClientServerAuth, nil, []string{"127.0.0.1"})
}
