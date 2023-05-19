/*
Copyright 2022 The OpenFunction Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1beta2 "github.com/openfunction/apis/core/v1beta2"
)

// FakeServings implements ServingInterface
type FakeServings struct {
	Fake *FakeCoreV1beta2
	ns   string
}

var servingsResource = schema.GroupVersionResource{Group: "core.openfunction.io", Version: "v1beta2", Resource: "servings"}

var servingsKind = schema.GroupVersionKind{Group: "core.openfunction.io", Version: "v1beta2", Kind: "Serving"}

// Get takes name of the serving, and returns the corresponding serving object, and an error if there is any.
func (c *FakeServings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.Serving, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servingsResource, c.ns, name), &v1beta2.Serving{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Serving), err
}

// List takes label and field selectors, and returns the list of Servings that match those selectors.
func (c *FakeServings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ServingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servingsResource, servingsKind, c.ns, opts), &v1beta2.ServingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ServingList{ListMeta: obj.(*v1beta2.ServingList).ListMeta}
	for _, item := range obj.(*v1beta2.ServingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servings.
func (c *FakeServings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servingsResource, c.ns, opts))

}

// Create takes the representation of a serving and creates it.  Returns the server's representation of the serving, and an error, if there is any.
func (c *FakeServings) Create(ctx context.Context, serving *v1beta2.Serving, opts v1.CreateOptions) (result *v1beta2.Serving, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servingsResource, c.ns, serving), &v1beta2.Serving{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Serving), err
}

// Update takes the representation of a serving and updates it. Returns the server's representation of the serving, and an error, if there is any.
func (c *FakeServings) Update(ctx context.Context, serving *v1beta2.Serving, opts v1.UpdateOptions) (result *v1beta2.Serving, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servingsResource, c.ns, serving), &v1beta2.Serving{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Serving), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServings) UpdateStatus(ctx context.Context, serving *v1beta2.Serving, opts v1.UpdateOptions) (*v1beta2.Serving, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servingsResource, "status", c.ns, serving), &v1beta2.Serving{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Serving), err
}

// Delete takes name of the serving and deletes it. Returns an error if one occurs.
func (c *FakeServings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servingsResource, c.ns, name), &v1beta2.Serving{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.ServingList{})
	return err
}

// Patch applies the patch and returns the patched serving.
func (c *FakeServings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.Serving, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servingsResource, c.ns, name, pt, data, subresources...), &v1beta2.Serving{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Serving), err
}
