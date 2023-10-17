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
	"context"

	v1alpha1 "github.com/kyverno/kyverno-json/pkg/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeValidationPolicies implements ValidationPolicyInterface
type FakeValidationPolicies struct {
	Fake *FakeJsonV1alpha1
}

var validationpoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("validationpolicies")

var validationpoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("ValidationPolicy")

// Get takes name of the validationPolicy, and returns the corresponding validationPolicy object, and an error if there is any.
func (c *FakeValidationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ValidationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(validationpoliciesResource, name), &v1alpha1.ValidationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidationPolicy), err
}

// List takes label and field selectors, and returns the list of ValidationPolicies that match those selectors.
func (c *FakeValidationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ValidationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(validationpoliciesResource, validationpoliciesKind, opts), &v1alpha1.ValidationPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ValidationPolicyList{ListMeta: obj.(*v1alpha1.ValidationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ValidationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested validationPolicies.
func (c *FakeValidationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(validationpoliciesResource, opts))
}

// Create takes the representation of a validationPolicy and creates it.  Returns the server's representation of the validationPolicy, and an error, if there is any.
func (c *FakeValidationPolicies) Create(ctx context.Context, validationPolicy *v1alpha1.ValidationPolicy, opts v1.CreateOptions) (result *v1alpha1.ValidationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(validationpoliciesResource, validationPolicy), &v1alpha1.ValidationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidationPolicy), err
}

// Update takes the representation of a validationPolicy and updates it. Returns the server's representation of the validationPolicy, and an error, if there is any.
func (c *FakeValidationPolicies) Update(ctx context.Context, validationPolicy *v1alpha1.ValidationPolicy, opts v1.UpdateOptions) (result *v1alpha1.ValidationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(validationpoliciesResource, validationPolicy), &v1alpha1.ValidationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidationPolicy), err
}

// Delete takes name of the validationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeValidationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(validationpoliciesResource, name, opts), &v1alpha1.ValidationPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeValidationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(validationpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ValidationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched validationPolicy.
func (c *FakeValidationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ValidationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(validationpoliciesResource, name, pt, data, subresources...), &v1alpha1.ValidationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidationPolicy), err
}
