/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package internalversion

import (
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	scheme "github.com/openshift/origin/pkg/oauth/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OAuthAccessTokensGetter has a method to return a OAuthAccessTokenInterface.
// A group's client should implement this interface.
type OAuthAccessTokensGetter interface {
	OAuthAccessTokens() OAuthAccessTokenInterface
}

// OAuthAccessTokenInterface has methods to work with OAuthAccessToken resources.
type OAuthAccessTokenInterface interface {
	Create(*oauth.OAuthAccessToken) (*oauth.OAuthAccessToken, error)
	Update(*oauth.OAuthAccessToken) (*oauth.OAuthAccessToken, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*oauth.OAuthAccessToken, error)
	List(opts v1.ListOptions) (*oauth.OAuthAccessTokenList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAccessToken, err error)
	OAuthAccessTokenExpansion
}

// oAuthAccessTokens implements OAuthAccessTokenInterface
type oAuthAccessTokens struct {
	client rest.Interface
}

// newOAuthAccessTokens returns a OAuthAccessTokens
func newOAuthAccessTokens(c *OauthClient) *oAuthAccessTokens {
	return &oAuthAccessTokens{
		client: c.RESTClient(),
	}
}

// Get takes name of the oAuthAccessToken, and returns the corresponding oAuthAccessToken object, and an error if there is any.
func (c *oAuthAccessTokens) Get(name string, options v1.GetOptions) (result *oauth.OAuthAccessToken, err error) {
	result = &oauth.OAuthAccessToken{}
	err = c.client.Get().
		Resource("oauthaccesstokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OAuthAccessTokens that match those selectors.
func (c *oAuthAccessTokens) List(opts v1.ListOptions) (result *oauth.OAuthAccessTokenList, err error) {
	result = &oauth.OAuthAccessTokenList{}
	err = c.client.Get().
		Resource("oauthaccesstokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oAuthAccessTokens.
func (c *oAuthAccessTokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("oauthaccesstokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a oAuthAccessToken and creates it.  Returns the server's representation of the oAuthAccessToken, and an error, if there is any.
func (c *oAuthAccessTokens) Create(oAuthAccessToken *oauth.OAuthAccessToken) (result *oauth.OAuthAccessToken, err error) {
	result = &oauth.OAuthAccessToken{}
	err = c.client.Post().
		Resource("oauthaccesstokens").
		Body(oAuthAccessToken).
		Do().
		Into(result)
	return
}

// Update takes the representation of a oAuthAccessToken and updates it. Returns the server's representation of the oAuthAccessToken, and an error, if there is any.
func (c *oAuthAccessTokens) Update(oAuthAccessToken *oauth.OAuthAccessToken) (result *oauth.OAuthAccessToken, err error) {
	result = &oauth.OAuthAccessToken{}
	err = c.client.Put().
		Resource("oauthaccesstokens").
		Name(oAuthAccessToken.Name).
		Body(oAuthAccessToken).
		Do().
		Into(result)
	return
}

// Delete takes name of the oAuthAccessToken and deletes it. Returns an error if one occurs.
func (c *oAuthAccessTokens) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("oauthaccesstokens").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oAuthAccessTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("oauthaccesstokens").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched oAuthAccessToken.
func (c *oAuthAccessTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *oauth.OAuthAccessToken, err error) {
	result = &oauth.OAuthAccessToken{}
	err = c.client.Patch(pt).
		Resource("oauthaccesstokens").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}