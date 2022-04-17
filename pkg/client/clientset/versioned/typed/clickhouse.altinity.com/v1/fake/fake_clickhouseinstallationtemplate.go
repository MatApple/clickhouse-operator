/*
Copyright 2021.

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

	clickhousealtinitycomv1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClickHouseInstallationTemplates implements ClickHouseInstallationTemplateInterface
type FakeClickHouseInstallationTemplates struct {
	Fake *FakeClickhouseV1
	ns   string
}

var clickhouseinstallationtemplatesResource = schema.GroupVersionResource{Group: "clickhouse.altinity.com", Version: "v1", Resource: "clickhouseinstallationtemplates"}

var clickhouseinstallationtemplatesKind = schema.GroupVersionKind{Group: "clickhouse.altinity.com", Version: "v1", Kind: "ClickHouseInstallationTemplate"}

// Get takes name of the clickHouseInstallationTemplate, and returns the corresponding clickHouseInstallationTemplate object, and an error if there is any.
func (c *FakeClickHouseInstallationTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *clickhousealtinitycomv1.ClickHouseInstallationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clickhouseinstallationtemplatesResource, c.ns, name), &clickhousealtinitycomv1.ClickHouseInstallationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousealtinitycomv1.ClickHouseInstallationTemplate), err
}

// List takes label and field selectors, and returns the list of ClickHouseInstallationTemplates that match those selectors.
func (c *FakeClickHouseInstallationTemplates) List(ctx context.Context, opts v1.ListOptions) (result *clickhousealtinitycomv1.ClickHouseInstallationTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clickhouseinstallationtemplatesResource, clickhouseinstallationtemplatesKind, c.ns, opts), &clickhousealtinitycomv1.ClickHouseInstallationTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clickhousealtinitycomv1.ClickHouseInstallationTemplateList{ListMeta: obj.(*clickhousealtinitycomv1.ClickHouseInstallationTemplateList).ListMeta}
	for _, item := range obj.(*clickhousealtinitycomv1.ClickHouseInstallationTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clickHouseInstallationTemplates.
func (c *FakeClickHouseInstallationTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clickhouseinstallationtemplatesResource, c.ns, opts))

}

// Create takes the representation of a clickHouseInstallationTemplate and creates it.  Returns the server's representation of the clickHouseInstallationTemplate, and an error, if there is any.
func (c *FakeClickHouseInstallationTemplates) Create(ctx context.Context, clickHouseInstallationTemplate *clickhousealtinitycomv1.ClickHouseInstallationTemplate, opts v1.CreateOptions) (result *clickhousealtinitycomv1.ClickHouseInstallationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clickhouseinstallationtemplatesResource, c.ns, clickHouseInstallationTemplate), &clickhousealtinitycomv1.ClickHouseInstallationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousealtinitycomv1.ClickHouseInstallationTemplate), err
}

// Update takes the representation of a clickHouseInstallationTemplate and updates it. Returns the server's representation of the clickHouseInstallationTemplate, and an error, if there is any.
func (c *FakeClickHouseInstallationTemplates) Update(ctx context.Context, clickHouseInstallationTemplate *clickhousealtinitycomv1.ClickHouseInstallationTemplate, opts v1.UpdateOptions) (result *clickhousealtinitycomv1.ClickHouseInstallationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clickhouseinstallationtemplatesResource, c.ns, clickHouseInstallationTemplate), &clickhousealtinitycomv1.ClickHouseInstallationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousealtinitycomv1.ClickHouseInstallationTemplate), err
}

// Delete takes name of the clickHouseInstallationTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClickHouseInstallationTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clickhouseinstallationtemplatesResource, c.ns, name), &clickhousealtinitycomv1.ClickHouseInstallationTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClickHouseInstallationTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clickhouseinstallationtemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &clickhousealtinitycomv1.ClickHouseInstallationTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clickHouseInstallationTemplate.
func (c *FakeClickHouseInstallationTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clickhousealtinitycomv1.ClickHouseInstallationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clickhouseinstallationtemplatesResource, c.ns, name, pt, data, subresources...), &clickhousealtinitycomv1.ClickHouseInstallationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousealtinitycomv1.ClickHouseInstallationTemplate), err
}
