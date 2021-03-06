package fake

import (
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/apis/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineDeployments implements MachineDeploymentInterface
type FakeMachineDeployments struct {
	Fake *FakeMachineV1alpha1
}

var machinedeploymentsResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "v1alpha1", Resource: "machinedeployments"}

var machinedeploymentsKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "v1alpha1", Kind: "MachineDeployment"}

// Get takes name of the machineDeployment, and returns the corresponding machineDeployment object, and an error if there is any.
func (c *FakeMachineDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machinedeploymentsResource, name), &v1alpha1.MachineDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDeployment), err
}

// List takes label and field selectors, and returns the list of MachineDeployments that match those selectors.
func (c *FakeMachineDeployments) List(opts v1.ListOptions) (result *v1alpha1.MachineDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(machinedeploymentsResource, machinedeploymentsKind, opts), &v1alpha1.MachineDeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineDeploymentList{}
	for _, item := range obj.(*v1alpha1.MachineDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineDeployments.
func (c *FakeMachineDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(machinedeploymentsResource, opts))
}

// Create takes the representation of a machineDeployment and creates it.  Returns the server's representation of the machineDeployment, and an error, if there is any.
func (c *FakeMachineDeployments) Create(machineDeployment *v1alpha1.MachineDeployment) (result *v1alpha1.MachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(machinedeploymentsResource, machineDeployment), &v1alpha1.MachineDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDeployment), err
}

// Update takes the representation of a machineDeployment and updates it. Returns the server's representation of the machineDeployment, and an error, if there is any.
func (c *FakeMachineDeployments) Update(machineDeployment *v1alpha1.MachineDeployment) (result *v1alpha1.MachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machinedeploymentsResource, machineDeployment), &v1alpha1.MachineDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineDeployments) UpdateStatus(machineDeployment *v1alpha1.MachineDeployment) (*v1alpha1.MachineDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(machinedeploymentsResource, "status", machineDeployment), &v1alpha1.MachineDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDeployment), err
}

// Delete takes name of the machineDeployment and deletes it. Returns an error if one occurs.
func (c *FakeMachineDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(machinedeploymentsResource, name), &v1alpha1.MachineDeployment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(machinedeploymentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched machineDeployment.
func (c *FakeMachineDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(machinedeploymentsResource, name, data, subresources...), &v1alpha1.MachineDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDeployment), err
}

// GetScale takes name of the machineDeployment, and returns the corresponding scale object, and an error if there is any.
func (c *FakeMachineDeployments) GetScale(machineDeploymentName string, options v1.GetOptions) (result *v1alpha1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machinedeploymentsResource, machineDeploymentName), &v1alpha1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeMachineDeployments) UpdateScale(machineDeploymentName string, scale *v1alpha1.Scale) (result *v1alpha1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machinedeploymentsResource, machineDeployment), &v1alpha1.MachineDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Scale), err
}
