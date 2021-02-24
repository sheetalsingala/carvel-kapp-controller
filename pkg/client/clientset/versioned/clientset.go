// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	installv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/typed/installpackage/v1alpha1"
	kappctrlv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/typed/kappctrl/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	InstallV1alpha1() installv1alpha1.InstallV1alpha1Interface
	KappctrlV1alpha1() kappctrlv1alpha1.KappctrlV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	installV1alpha1  *installv1alpha1.InstallV1alpha1Client
	kappctrlV1alpha1 *kappctrlv1alpha1.KappctrlV1alpha1Client
}

// InstallV1alpha1 retrieves the InstallV1alpha1Client
func (c *Clientset) InstallV1alpha1() installv1alpha1.InstallV1alpha1Interface {
	return c.installV1alpha1
}

// KappctrlV1alpha1 retrieves the KappctrlV1alpha1Client
func (c *Clientset) KappctrlV1alpha1() kappctrlv1alpha1.KappctrlV1alpha1Interface {
	return c.kappctrlV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.installV1alpha1, err = installv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kappctrlV1alpha1, err = kappctrlv1alpha1.NewForConfig(&configShallowCopy)
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
	cs.installV1alpha1 = installv1alpha1.NewForConfigOrDie(c)
	cs.kappctrlV1alpha1 = kappctrlv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.installV1alpha1 = installv1alpha1.New(c)
	cs.kappctrlV1alpha1 = kappctrlv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
