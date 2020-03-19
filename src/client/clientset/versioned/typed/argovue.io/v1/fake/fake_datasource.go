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

package fake

import (
	argovueiov1 "argovue/apis/argovue.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatasources implements DatasourceInterface
type FakeDatasources struct {
	Fake *FakeArgovueV1
	ns   string
}

var datasourcesResource = schema.GroupVersionResource{Group: "argovue.io", Version: "v1", Resource: "datasources"}

var datasourcesKind = schema.GroupVersionKind{Group: "argovue.io", Version: "v1", Kind: "Datasource"}

// Get takes name of the datasource, and returns the corresponding datasource object, and an error if there is any.
func (c *FakeDatasources) Get(name string, options v1.GetOptions) (result *argovueiov1.Datasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasourcesResource, c.ns, name), &argovueiov1.Datasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*argovueiov1.Datasource), err
}

// List takes label and field selectors, and returns the list of Datasources that match those selectors.
func (c *FakeDatasources) List(opts v1.ListOptions) (result *argovueiov1.DatasourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasourcesResource, datasourcesKind, c.ns, opts), &argovueiov1.DatasourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &argovueiov1.DatasourceList{ListMeta: obj.(*argovueiov1.DatasourceList).ListMeta}
	for _, item := range obj.(*argovueiov1.DatasourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasources.
func (c *FakeDatasources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasourcesResource, c.ns, opts))

}

// Create takes the representation of a datasource and creates it.  Returns the server's representation of the datasource, and an error, if there is any.
func (c *FakeDatasources) Create(datasource *argovueiov1.Datasource) (result *argovueiov1.Datasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasourcesResource, c.ns, datasource), &argovueiov1.Datasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*argovueiov1.Datasource), err
}

// Update takes the representation of a datasource and updates it. Returns the server's representation of the datasource, and an error, if there is any.
func (c *FakeDatasources) Update(datasource *argovueiov1.Datasource) (result *argovueiov1.Datasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasourcesResource, c.ns, datasource), &argovueiov1.Datasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*argovueiov1.Datasource), err
}

// Delete takes name of the datasource and deletes it. Returns an error if one occurs.
func (c *FakeDatasources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasourcesResource, c.ns, name), &argovueiov1.Datasource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &argovueiov1.DatasourceList{})
	return err
}

// Patch applies the patch and returns the patched datasource.
func (c *FakeDatasources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *argovueiov1.Datasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasourcesResource, c.ns, name, pt, data, subresources...), &argovueiov1.Datasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*argovueiov1.Datasource), err
}