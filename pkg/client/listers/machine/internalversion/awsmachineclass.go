// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSMachineClassLister helps list AWSMachineClasses.
type AWSMachineClassLister interface {
	// List lists all AWSMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*machine.AWSMachineClass, err error)
	// Get retrieves the AWSMachineClass from the index for a given name.
	Get(name string) (*machine.AWSMachineClass, error)
	AWSMachineClassListerExpansion
}

// aWSMachineClassLister implements the AWSMachineClassLister interface.
type aWSMachineClassLister struct {
	indexer cache.Indexer
}

// NewAWSMachineClassLister returns a new AWSMachineClassLister.
func NewAWSMachineClassLister(indexer cache.Indexer) AWSMachineClassLister {
	return &aWSMachineClassLister{indexer: indexer}
}

// List lists all AWSMachineClasses in the indexer.
func (s *aWSMachineClassLister) List(selector labels.Selector) (ret []*machine.AWSMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.AWSMachineClass))
	})
	return ret, err
}

// Get retrieves the AWSMachineClass from the index for a given name.
func (s *aWSMachineClassLister) Get(name string) (*machine.AWSMachineClass, error) {
	key := &machine.AWSMachineClass{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("awsmachineclass"), name)
	}
	return obj.(*machine.AWSMachineClass), nil
}
