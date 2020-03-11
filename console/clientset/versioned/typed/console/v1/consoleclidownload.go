// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openshift/api/console/v1"
	scheme "github.com/openshift/client-go/console/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsoleCLIDownloadsGetter has a method to return a ConsoleCLIDownloadInterface.
// A group's client should implement this interface.
type ConsoleCLIDownloadsGetter interface {
	ConsoleCLIDownloads() ConsoleCLIDownloadInterface
}

// ConsoleCLIDownloadInterface has methods to work with ConsoleCLIDownload resources.
type ConsoleCLIDownloadInterface interface {
	Create(ctx context.Context, consoleCLIDownload *v1.ConsoleCLIDownload, opts metav1.CreateOptions) (*v1.ConsoleCLIDownload, error)
	Update(ctx context.Context, consoleCLIDownload *v1.ConsoleCLIDownload, opts metav1.UpdateOptions) (*v1.ConsoleCLIDownload, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ConsoleCLIDownload, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ConsoleCLIDownloadList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConsoleCLIDownload, err error)
	ConsoleCLIDownloadExpansion
}

// consoleCLIDownloads implements ConsoleCLIDownloadInterface
type consoleCLIDownloads struct {
	client rest.Interface
}

// newConsoleCLIDownloads returns a ConsoleCLIDownloads
func newConsoleCLIDownloads(c *ConsoleV1Client) *consoleCLIDownloads {
	return &consoleCLIDownloads{
		client: c.RESTClient(),
	}
}

// Get takes name of the consoleCLIDownload, and returns the corresponding consoleCLIDownload object, and an error if there is any.
func (c *consoleCLIDownloads) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConsoleCLIDownload, err error) {
	result = &v1.ConsoleCLIDownload{}
	err = c.client.Get().
		Resource("consoleclidownloads").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConsoleCLIDownloads that match those selectors.
func (c *consoleCLIDownloads) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsoleCLIDownloadList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ConsoleCLIDownloadList{}
	err = c.client.Get().
		Resource("consoleclidownloads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoleCLIDownloads.
func (c *consoleCLIDownloads) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("consoleclidownloads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a consoleCLIDownload and creates it.  Returns the server's representation of the consoleCLIDownload, and an error, if there is any.
func (c *consoleCLIDownloads) Create(ctx context.Context, consoleCLIDownload *v1.ConsoleCLIDownload, opts metav1.CreateOptions) (result *v1.ConsoleCLIDownload, err error) {
	result = &v1.ConsoleCLIDownload{}
	err = c.client.Post().
		Resource("consoleclidownloads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleCLIDownload).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a consoleCLIDownload and updates it. Returns the server's representation of the consoleCLIDownload, and an error, if there is any.
func (c *consoleCLIDownloads) Update(ctx context.Context, consoleCLIDownload *v1.ConsoleCLIDownload, opts metav1.UpdateOptions) (result *v1.ConsoleCLIDownload, err error) {
	result = &v1.ConsoleCLIDownload{}
	err = c.client.Put().
		Resource("consoleclidownloads").
		Name(consoleCLIDownload.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleCLIDownload).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the consoleCLIDownload and deletes it. Returns an error if one occurs.
func (c *consoleCLIDownloads) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("consoleclidownloads").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoleCLIDownloads) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("consoleclidownloads").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched consoleCLIDownload.
func (c *consoleCLIDownloads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConsoleCLIDownload, err error) {
	result = &v1.ConsoleCLIDownload{}
	err = c.client.Patch(pt).
		Resource("consoleclidownloads").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
