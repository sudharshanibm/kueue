/*
Copyright 2022 The Kubernetes Authors.

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
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/kueue/apis/visibility/v1alpha1"
	visibilityv1alpha1 "sigs.k8s.io/kueue/client-go/applyconfiguration/visibility/v1alpha1"
	scheme "sigs.k8s.io/kueue/client-go/clientset/versioned/scheme"
)

// LocalQueuesGetter has a method to return a LocalQueueInterface.
// A group's client should implement this interface.
type LocalQueuesGetter interface {
	LocalQueues(namespace string) LocalQueueInterface
}

// LocalQueueInterface has methods to work with LocalQueue resources.
type LocalQueueInterface interface {
	Create(ctx context.Context, localQueue *v1alpha1.LocalQueue, opts v1.CreateOptions) (*v1alpha1.LocalQueue, error)
	Update(ctx context.Context, localQueue *v1alpha1.LocalQueue, opts v1.UpdateOptions) (*v1alpha1.LocalQueue, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.LocalQueue, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LocalQueueList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalQueue, err error)
	Apply(ctx context.Context, localQueue *visibilityv1alpha1.LocalQueueApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LocalQueue, err error)
	GetPendingWorkloadsSummary(ctx context.Context, localQueueName string, options v1.GetOptions) (*v1alpha1.PendingWorkloadsSummary, error)

	LocalQueueExpansion
}

// localQueues implements LocalQueueInterface
type localQueues struct {
	client rest.Interface
	ns     string
}

// newLocalQueues returns a LocalQueues
func newLocalQueues(c *VisibilityV1alpha1Client, namespace string) *localQueues {
	return &localQueues{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the localQueue, and returns the corresponding localQueue object, and an error if there is any.
func (c *localQueues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalQueue, err error) {
	result = &v1alpha1.LocalQueue{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localqueues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalQueues that match those selectors.
func (c *localQueues) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalQueueList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LocalQueueList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localQueues.
func (c *localQueues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("localqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a localQueue and creates it.  Returns the server's representation of the localQueue, and an error, if there is any.
func (c *localQueues) Create(ctx context.Context, localQueue *v1alpha1.LocalQueue, opts v1.CreateOptions) (result *v1alpha1.LocalQueue, err error) {
	result = &v1alpha1.LocalQueue{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("localqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localQueue).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a localQueue and updates it. Returns the server's representation of the localQueue, and an error, if there is any.
func (c *localQueues) Update(ctx context.Context, localQueue *v1alpha1.LocalQueue, opts v1.UpdateOptions) (result *v1alpha1.LocalQueue, err error) {
	result = &v1alpha1.LocalQueue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localqueues").
		Name(localQueue.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localQueue).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the localQueue and deletes it. Returns an error if one occurs.
func (c *localQueues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localqueues").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localQueues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localqueues").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched localQueue.
func (c *localQueues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalQueue, err error) {
	result = &v1alpha1.LocalQueue{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("localqueues").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied localQueue.
func (c *localQueues) Apply(ctx context.Context, localQueue *visibilityv1alpha1.LocalQueueApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LocalQueue, err error) {
	if localQueue == nil {
		return nil, fmt.Errorf("localQueue provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(localQueue)
	if err != nil {
		return nil, err
	}
	name := localQueue.Name
	if name == nil {
		return nil, fmt.Errorf("localQueue.Name must be provided to Apply")
	}
	result = &v1alpha1.LocalQueue{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("localqueues").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetPendingWorkloadsSummary takes name of the localQueue, and returns the corresponding v1alpha1.PendingWorkloadsSummary object, and an error if there is any.
func (c *localQueues) GetPendingWorkloadsSummary(ctx context.Context, localQueueName string, options v1.GetOptions) (result *v1alpha1.PendingWorkloadsSummary, err error) {
	result = &v1alpha1.PendingWorkloadsSummary{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localqueues").
		Name(localQueueName).
		SubResource("pendingworkloads").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}