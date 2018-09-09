/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package kubernetes

import (
	discovery "k8s.io/client-go/discovery"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	auditregistrationv1alpha1 "k8s.io/client-go/kubernetes/typed/auditregistration/v1alpha1"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	batchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface
	AppsV1() appsv1.AppsV1Interface
	AppsV1beta1() appsv1beta1.AppsV1beta1Interface
	AppsV1beta2() appsv1beta2.AppsV1beta2Interface
	AuditregistrationV1alpha1() auditregistrationv1alpha1.AuditregistrationV1alpha1Interface
	AuthenticationV1() authenticationv1.AuthenticationV1Interface
	AuthorizationV1() authorizationv1.AuthorizationV1Interface
	AutoscalingV1() autoscalingv1.AutoscalingV1Interface
	AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface
	AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface
	BatchV1() batchv1.BatchV1Interface
	BatchV1beta1() batchv1beta1.BatchV1beta1Interface
	BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface
	CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface
	CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface
	CoordinationV1() coordinationv1.CoordinationV1Interface
	CoreV1() corev1.CoreV1Interface
	ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
	NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface
	NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface
	NodeV1beta1() nodev1beta1.NodeV1beta1Interface
	PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface
	RbacV1() rbacv1.RbacV1Interface
	RbacV1beta1() rbacv1beta1.RbacV1beta1Interface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface
	SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface
	SchedulingV1() schedulingv1.SchedulingV1Interface
	StorageV1beta1() storagev1beta1.StorageV1beta1Interface
	StorageV1() storagev1.StorageV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	admissionregistrationV1beta1 *admissionregistrationv1beta1.AdmissionregistrationV1beta1Client
	appsV1                       *appsv1.AppsV1Client
	appsV1beta1                  *appsv1beta1.AppsV1beta1Client
	appsV1beta2                  *appsv1beta2.AppsV1beta2Client
	auditregistrationV1alpha1    *auditregistrationv1alpha1.AuditregistrationV1alpha1Client
	authenticationV1             *authenticationv1.AuthenticationV1Client
	authorizationV1              *authorizationv1.AuthorizationV1Client
	autoscalingV1                *autoscalingv1.AutoscalingV1Client
	autoscalingV2beta1           *autoscalingv2beta1.AutoscalingV2beta1Client
	autoscalingV2beta2           *autoscalingv2beta2.AutoscalingV2beta2Client
	batchV1                      *batchv1.BatchV1Client
	batchV1beta1                 *batchv1beta1.BatchV1beta1Client
	batchV2alpha1                *batchv2alpha1.BatchV2alpha1Client
	certificatesV1beta1          *certificatesv1beta1.CertificatesV1beta1Client
	coordinationV1beta1          *coordinationv1beta1.CoordinationV1beta1Client
	coordinationV1               *coordinationv1.CoordinationV1Client
	coreV1                       *corev1.CoreV1Client
	extensionsV1beta1            *extensionsv1beta1.ExtensionsV1beta1Client
	networkingV1                 *networkingv1.NetworkingV1Client
	networkingV1beta1            *networkingv1beta1.NetworkingV1beta1Client
	nodeV1alpha1                 *nodev1alpha1.NodeV1alpha1Client
	nodeV1beta1                  *nodev1beta1.NodeV1beta1Client
	policyV1beta1                *policyv1beta1.PolicyV1beta1Client
	rbacV1                       *rbacv1.RbacV1Client
	rbacV1beta1                  *rbacv1beta1.RbacV1beta1Client
	schedulingV1alpha1           *schedulingv1alpha1.SchedulingV1alpha1Client
	schedulingV1beta1            *schedulingv1beta1.SchedulingV1beta1Client
	schedulingV1                 *schedulingv1.SchedulingV1Client
	storageV1beta1               *storagev1beta1.StorageV1beta1Client
	storageV1                    *storagev1.StorageV1Client
}

// AdmissionregistrationV1beta1 retrieves the AdmissionregistrationV1beta1Client
func (c *Clientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return c.admissionregistrationV1beta1
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return c.appsV1beta1
}

// AppsV1beta2 retrieves the AppsV1beta2Client
func (c *Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return c.appsV1beta2
}

// AuditregistrationV1alpha1 retrieves the AuditregistrationV1alpha1Client
func (c *Clientset) AuditregistrationV1alpha1() auditregistrationv1alpha1.AuditregistrationV1alpha1Interface {
	return c.auditregistrationV1alpha1
}

// AuthenticationV1 retrieves the AuthenticationV1Client
func (c *Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return c.authenticationV1
}

// AuthorizationV1 retrieves the AuthorizationV1Client
func (c *Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return c.authorizationV1
}

// AutoscalingV1 retrieves the AutoscalingV1Client
func (c *Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return c.autoscalingV1
}

// AutoscalingV2beta1 retrieves the AutoscalingV2beta1Client
func (c *Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return c.autoscalingV2beta1
}

// AutoscalingV2beta2 retrieves the AutoscalingV2beta2Client
func (c *Clientset) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	return c.autoscalingV2beta2
}

// BatchV1 retrieves the BatchV1Client
func (c *Clientset) BatchV1() batchv1.BatchV1Interface {
	return c.batchV1
}

// BatchV1beta1 retrieves the BatchV1beta1Client
func (c *Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return c.batchV1beta1
}

// BatchV2alpha1 retrieves the BatchV2alpha1Client
func (c *Clientset) BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface {
	return c.batchV2alpha1
}

// CertificatesV1beta1 retrieves the CertificatesV1beta1Client
func (c *Clientset) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	return c.certificatesV1beta1
}

// CoordinationV1beta1 retrieves the CoordinationV1beta1Client
func (c *Clientset) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	return c.coordinationV1beta1
}

// CoordinationV1 retrieves the CoordinationV1Client
func (c *Clientset) CoordinationV1() coordinationv1.CoordinationV1Interface {
	return c.coordinationV1
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}

// ExtensionsV1beta1 retrieves the ExtensionsV1beta1Client
func (c *Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return c.extensionsV1beta1
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return c.networkingV1
}

// NetworkingV1beta1 retrieves the NetworkingV1beta1Client
func (c *Clientset) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	return c.networkingV1beta1
}

// NodeV1alpha1 retrieves the NodeV1alpha1Client
func (c *Clientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return c.nodeV1alpha1
}

// NodeV1beta1 retrieves the NodeV1beta1Client
func (c *Clientset) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	return c.nodeV1beta1
}

// PolicyV1beta1 retrieves the PolicyV1beta1Client
func (c *Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return c.policyV1beta1
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return c.rbacV1
}

// RbacV1beta1 retrieves the RbacV1beta1Client
func (c *Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return c.rbacV1beta1
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// SchedulingV1beta1 retrieves the SchedulingV1beta1Client
func (c *Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return c.schedulingV1beta1
}

// SchedulingV1 retrieves the SchedulingV1Client
func (c *Clientset) SchedulingV1() schedulingv1.SchedulingV1Interface {
	return c.schedulingV1
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return c.storageV1beta1
}

// StorageV1 retrieves the StorageV1Client
func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
	return c.storageV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.admissionregistrationV1beta1, err = admissionregistrationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1, err = appsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta1, err = appsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta2, err = appsv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.auditregistrationV1alpha1, err = auditregistrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1, err = authenticationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1, err = authorizationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1, err = autoscalingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta1, err = autoscalingv2beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta2, err = autoscalingv2beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV1, err = batchv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV1beta1, err = batchv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV2alpha1, err = batchv2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1beta1, err = certificatesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1beta1, err = coordinationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1, err = coordinationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.extensionsV1beta1, err = extensionsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1beta1, err = networkingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.nodeV1alpha1, err = nodev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.nodeV1beta1, err = nodev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.policyV1beta1, err = policyv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1, err = rbacv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1beta1, err = rbacv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1beta1, err = schedulingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1, err = schedulingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1beta1, err = storagev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1, err = storagev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.admissionregistrationV1beta1 = admissionregistrationv1beta1.NewForConfigOrDie(c)
	cs.appsV1 = appsv1.NewForConfigOrDie(c)
	cs.appsV1beta1 = appsv1beta1.NewForConfigOrDie(c)
	cs.appsV1beta2 = appsv1beta2.NewForConfigOrDie(c)
	cs.auditregistrationV1alpha1 = auditregistrationv1alpha1.NewForConfigOrDie(c)
	cs.authenticationV1 = authenticationv1.NewForConfigOrDie(c)
	cs.authorizationV1 = authorizationv1.NewForConfigOrDie(c)
	cs.autoscalingV1 = autoscalingv1.NewForConfigOrDie(c)
	cs.autoscalingV2beta1 = autoscalingv2beta1.NewForConfigOrDie(c)
	cs.autoscalingV2beta2 = autoscalingv2beta2.NewForConfigOrDie(c)
	cs.batchV1 = batchv1.NewForConfigOrDie(c)
	cs.batchV1beta1 = batchv1beta1.NewForConfigOrDie(c)
	cs.batchV2alpha1 = batchv2alpha1.NewForConfigOrDie(c)
	cs.certificatesV1beta1 = certificatesv1beta1.NewForConfigOrDie(c)
	cs.coordinationV1beta1 = coordinationv1beta1.NewForConfigOrDie(c)
	cs.coordinationV1 = coordinationv1.NewForConfigOrDie(c)
	cs.coreV1 = corev1.NewForConfigOrDie(c)
	cs.extensionsV1beta1 = extensionsv1beta1.NewForConfigOrDie(c)
	cs.networkingV1 = networkingv1.NewForConfigOrDie(c)
	cs.networkingV1beta1 = networkingv1beta1.NewForConfigOrDie(c)
	cs.nodeV1alpha1 = nodev1alpha1.NewForConfigOrDie(c)
	cs.nodeV1beta1 = nodev1beta1.NewForConfigOrDie(c)
	cs.policyV1beta1 = policyv1beta1.NewForConfigOrDie(c)
	cs.rbacV1 = rbacv1.NewForConfigOrDie(c)
	cs.rbacV1beta1 = rbacv1beta1.NewForConfigOrDie(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.NewForConfigOrDie(c)
	cs.schedulingV1beta1 = schedulingv1beta1.NewForConfigOrDie(c)
	cs.schedulingV1 = schedulingv1.NewForConfigOrDie(c)
	cs.storageV1beta1 = storagev1beta1.NewForConfigOrDie(c)
	cs.storageV1 = storagev1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.admissionregistrationV1beta1 = admissionregistrationv1beta1.New(c)
	cs.appsV1 = appsv1.New(c)
	cs.appsV1beta1 = appsv1beta1.New(c)
	cs.appsV1beta2 = appsv1beta2.New(c)
	cs.auditregistrationV1alpha1 = auditregistrationv1alpha1.New(c)
	cs.authenticationV1 = authenticationv1.New(c)
	cs.authorizationV1 = authorizationv1.New(c)
	cs.autoscalingV1 = autoscalingv1.New(c)
	cs.autoscalingV2beta1 = autoscalingv2beta1.New(c)
	cs.autoscalingV2beta2 = autoscalingv2beta2.New(c)
	cs.batchV1 = batchv1.New(c)
	cs.batchV1beta1 = batchv1beta1.New(c)
	cs.batchV2alpha1 = batchv2alpha1.New(c)
	cs.certificatesV1beta1 = certificatesv1beta1.New(c)
	cs.coordinationV1beta1 = coordinationv1beta1.New(c)
	cs.coordinationV1 = coordinationv1.New(c)
	cs.coreV1 = corev1.New(c)
	cs.extensionsV1beta1 = extensionsv1beta1.New(c)
	cs.networkingV1 = networkingv1.New(c)
	cs.networkingV1beta1 = networkingv1beta1.New(c)
	cs.nodeV1alpha1 = nodev1alpha1.New(c)
	cs.nodeV1beta1 = nodev1beta1.New(c)
	cs.policyV1beta1 = policyv1beta1.New(c)
	cs.rbacV1 = rbacv1.New(c)
	cs.rbacV1beta1 = rbacv1beta1.New(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.New(c)
	cs.schedulingV1beta1 = schedulingv1beta1.New(c)
	cs.schedulingV1 = schedulingv1.New(c)
	cs.storageV1beta1 = storagev1beta1.New(c)
	cs.storageV1 = storagev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
