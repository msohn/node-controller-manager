// This file was automatically generated by informer-gen

package v1alpha1

import (
	machine_v1alpha1 "github.com/gardener/node-controller-manager/pkg/apis/machine/v1alpha1"
	clientset "github.com/gardener/node-controller-manager/pkg/client/clientset"
	internalinterfaces "github.com/gardener/node-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/client/listers/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// MachineDeploymentInformer provides access to a shared informer and lister for
// MachineDeployments.
type MachineDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineDeploymentLister
}

type machineDeploymentInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewMachineDeploymentInformer constructs a new informer for MachineDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineDeploymentInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.MachineV1alpha1().MachineDeployments().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.MachineV1alpha1().MachineDeployments().Watch(options)
			},
		},
		&machine_v1alpha1.MachineDeployment{},
		resyncPeriod,
		indexers,
	)
}

func defaultMachineDeploymentInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewMachineDeploymentInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *machineDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machine_v1alpha1.MachineDeployment{}, defaultMachineDeploymentInformer)
}

func (f *machineDeploymentInformer) Lister() v1alpha1.MachineDeploymentLister {
	return v1alpha1.NewMachineDeploymentLister(f.Informer().GetIndexer())
}
