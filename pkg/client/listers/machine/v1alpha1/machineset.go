// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineSetLister helps list MachineSets.
type MachineSetLister interface {
	// List lists all MachineSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MachineSet, err error)
	// Get retrieves the MachineSet from the index for a given name.
	Get(name string) (*v1alpha1.MachineSet, error)
	MachineSetListerExpansion
}

// machineSetLister implements the MachineSetLister interface.
type machineSetLister struct {
	indexer cache.Indexer
}

// NewMachineSetLister returns a new MachineSetLister.
func NewMachineSetLister(indexer cache.Indexer) MachineSetLister {
	return &machineSetLister{indexer: indexer}
}

// List lists all MachineSets in the indexer.
func (s *machineSetLister) List(selector labels.Selector) (ret []*v1alpha1.MachineSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineSet))
	})
	return ret, err
}

// Get retrieves the MachineSet from the index for a given name.
func (s *machineSetLister) Get(name string) (*v1alpha1.MachineSet, error) {
	key := &v1alpha1.MachineSet{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("machineset"), name)
	}
	return obj.(*v1alpha1.MachineSet), nil
}
