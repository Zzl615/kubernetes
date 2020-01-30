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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "k8s.io/code-generator/_examples/MixedCase/apis/example/v1"
	scheme "k8s.io/code-generator/_examples/MixedCase/clientset/versioned/scheme"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

// ClusterTestTypesGetter has a method to return a ClusterTestTypeInterface.
// A group's client should implement this interface.
type ClusterTestTypesGetter interface {
	ClusterTestTypes() ClusterTestTypeInterface
}

// ClusterTestTypeInterface has methods to work with ClusterTestType resources.
type ClusterTestTypeInterface interface {
	Create(*v1.ClusterTestType) (*v1.ClusterTestType, error)
	Update(*v1.ClusterTestType) (*v1.ClusterTestType, error)
	UpdateStatus(*v1.ClusterTestType) (*v1.ClusterTestType, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterTestType, error)
	List(opts metav1.ListOptions) (*v1.ClusterTestTypeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterTestType, err error)
	GetScale(clusterTestTypeName string, options metav1.GetOptions) (*autoscaling.Scale, error)
	UpdateScale(clusterTestTypeName string, scale *autoscaling.Scale) (*autoscaling.Scale, error)

	ClusterTestTypeExpansion
}

// clusterTestTypes implements ClusterTestTypeInterface
type clusterTestTypes struct {
	client rest.Interface
}

// newClusterTestTypes returns a ClusterTestTypes
func newClusterTestTypes(c *ExampleV1Client) *clusterTestTypes {
	return &clusterTestTypes{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterTestType, and returns the corresponding clusterTestType object, and an error if there is any.
func (c *clusterTestTypes) Get(name string, options metav1.GetOptions) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Get().
		Resource("clustertesttypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *clusterTestTypes) List(opts metav1.ListOptions) (result *v1.ClusterTestTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterTestTypeList{}
	err = c.client.Get().
		Resource("clustertesttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterTestTypes.
func (c *clusterTestTypes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustertesttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a clusterTestType and creates it.  Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *clusterTestTypes) Create(clusterTestType *v1.ClusterTestType) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Post().
		Resource("clustertesttypes").
		Body(clusterTestType).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a clusterTestType and updates it. Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *clusterTestTypes) Update(clusterTestType *v1.ClusterTestType) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestType.Name).
		Body(clusterTestType).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterTestTypes) UpdateStatus(clusterTestType *v1.ClusterTestType) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestType.Name).
		SubResource("status").
		Body(clusterTestType).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the clusterTestType and deletes it. Returns an error if one occurs.
func (c *clusterTestTypes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustertesttypes").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTestTypes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustertesttypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched clusterTestType.
func (c *clusterTestTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Patch(pt).
		Resource("clustertesttypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}

// GetScale takes name of the clusterTestType, and returns the corresponding autoscaling.Scale object, and an error if there is any.
func (c *clusterTestTypes) GetScale(clusterTestTypeName string, options metav1.GetOptions) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Get().
		Resource("clustertesttypes").
		Name(clusterTestTypeName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *clusterTestTypes) UpdateScale(clusterTestTypeName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestTypeName).
		SubResource("scale").
		Body(scale).
		Do(context.TODO()).
		Into(result)
	return
}
