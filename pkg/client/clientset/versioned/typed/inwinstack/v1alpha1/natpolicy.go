/*
Copyright © 2018 Kyle Bai(kyle.b@inwinstack.com)

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

package v1alpha1

import (
	v1alpha1 "github.com/inwinstack/pan-operator/pkg/apis/inwinstack/v1alpha1"
	scheme "github.com/inwinstack/pan-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NATPoliciesGetter has a method to return a NATPolicyInterface.
// A group's client should implement this interface.
type NATPoliciesGetter interface {
	NATPolicies(namespace string) NATPolicyInterface
}

// NATPolicyInterface has methods to work with NATPolicy resources.
type NATPolicyInterface interface {
	Create(*v1alpha1.NATPolicy) (*v1alpha1.NATPolicy, error)
	Update(*v1alpha1.NATPolicy) (*v1alpha1.NATPolicy, error)
	UpdateStatus(*v1alpha1.NATPolicy) (*v1alpha1.NATPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NATPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.NATPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NATPolicy, err error)
	NATPolicyExpansion
}

// nATPolicies implements NATPolicyInterface
type nATPolicies struct {
	client rest.Interface
	ns     string
}

// newNATPolicies returns a NATPolicies
func newNATPolicies(c *InwinstackV1alpha1Client, namespace string) *nATPolicies {
	return &nATPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nATPolicy, and returns the corresponding nATPolicy object, and an error if there is any.
func (c *nATPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.NATPolicy, err error) {
	result = &v1alpha1.NATPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NATPolicies that match those selectors.
func (c *nATPolicies) List(opts v1.ListOptions) (result *v1alpha1.NATPolicyList, err error) {
	result = &v1alpha1.NATPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nATPolicies.
func (c *nATPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("natpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nATPolicy and creates it.  Returns the server's representation of the nATPolicy, and an error, if there is any.
func (c *nATPolicies) Create(nATPolicy *v1alpha1.NATPolicy) (result *v1alpha1.NATPolicy, err error) {
	result = &v1alpha1.NATPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("natpolicies").
		Body(nATPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nATPolicy and updates it. Returns the server's representation of the nATPolicy, and an error, if there is any.
func (c *nATPolicies) Update(nATPolicy *v1alpha1.NATPolicy) (result *v1alpha1.NATPolicy, err error) {
	result = &v1alpha1.NATPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natpolicies").
		Name(nATPolicy.Name).
		Body(nATPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nATPolicies) UpdateStatus(nATPolicy *v1alpha1.NATPolicy) (result *v1alpha1.NATPolicy, err error) {
	result = &v1alpha1.NATPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natpolicies").
		Name(nATPolicy.Name).
		SubResource("status").
		Body(nATPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the nATPolicy and deletes it. Returns an error if one occurs.
func (c *nATPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nATPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nATPolicy.
func (c *nATPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NATPolicy, err error) {
	result = &v1alpha1.NATPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("natpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
