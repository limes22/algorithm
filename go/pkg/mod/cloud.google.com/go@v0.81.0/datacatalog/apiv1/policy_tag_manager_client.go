// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package datacatalog

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newPolicyTagManagerClientHook clientHook

// PolicyTagManagerCallOptions contains the retry settings for each method of PolicyTagManagerClient.
type PolicyTagManagerCallOptions struct {
	CreateTaxonomy     []gax.CallOption
	DeleteTaxonomy     []gax.CallOption
	UpdateTaxonomy     []gax.CallOption
	ListTaxonomies     []gax.CallOption
	GetTaxonomy        []gax.CallOption
	CreatePolicyTag    []gax.CallOption
	DeletePolicyTag    []gax.CallOption
	UpdatePolicyTag    []gax.CallOption
	ListPolicyTags     []gax.CallOption
	GetPolicyTag       []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultPolicyTagManagerClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("datacatalog.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("datacatalog.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://datacatalog.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPolicyTagManagerCallOptions() *PolicyTagManagerCallOptions {
	return &PolicyTagManagerCallOptions{
		CreateTaxonomy:     []gax.CallOption{},
		DeleteTaxonomy:     []gax.CallOption{},
		UpdateTaxonomy:     []gax.CallOption{},
		ListTaxonomies:     []gax.CallOption{},
		GetTaxonomy:        []gax.CallOption{},
		CreatePolicyTag:    []gax.CallOption{},
		DeletePolicyTag:    []gax.CallOption{},
		UpdatePolicyTag:    []gax.CallOption{},
		ListPolicyTags:     []gax.CallOption{},
		GetPolicyTag:       []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
	}
}

// PolicyTagManagerClient is a client for interacting with Google Cloud Data Catalog API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type PolicyTagManagerClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	policyTagManagerClient datacatalogpb.PolicyTagManagerClient

	// The call options for this service.
	CallOptions *PolicyTagManagerCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPolicyTagManagerClient creates a new policy tag manager client.
//
// Policy Tag Manager API service allows clients to manage their policy tags and
// taxonomies.
//
// Policy tags are used to tag BigQuery columns and apply additional access
// control policies. A taxonomy is a hierarchical grouping of policy tags that
// classify data along a common axis.
func NewPolicyTagManagerClient(ctx context.Context, opts ...option.ClientOption) (*PolicyTagManagerClient, error) {
	clientOpts := defaultPolicyTagManagerClientOptions()

	if newPolicyTagManagerClientHook != nil {
		hookOpts, err := newPolicyTagManagerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &PolicyTagManagerClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultPolicyTagManagerCallOptions(),

		policyTagManagerClient: datacatalogpb.NewPolicyTagManagerClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PolicyTagManagerClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PolicyTagManagerClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PolicyTagManagerClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateTaxonomy creates a taxonomy in a specified project. The taxonomy is initially empty,
// i.e., does not contain policy tags.
func (c *PolicyTagManagerClient) CreateTaxonomy(ctx context.Context, req *datacatalogpb.CreateTaxonomyRequest, opts ...gax.CallOption) (*datacatalogpb.Taxonomy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateTaxonomy[0:len(c.CallOptions.CreateTaxonomy):len(c.CallOptions.CreateTaxonomy)], opts...)
	var resp *datacatalogpb.Taxonomy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.CreateTaxonomy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteTaxonomy deletes a taxonomy. This method will also delete all policy tags in this
// taxonomy, their associated policies, and the policy tags references from
// BigQuery columns.
func (c *PolicyTagManagerClient) DeleteTaxonomy(ctx context.Context, req *datacatalogpb.DeleteTaxonomyRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteTaxonomy[0:len(c.CallOptions.DeleteTaxonomy):len(c.CallOptions.DeleteTaxonomy)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.policyTagManagerClient.DeleteTaxonomy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// UpdateTaxonomy updates a taxonomy. This method can update the taxonomy’s display name,
// description, and activated policy types.
func (c *PolicyTagManagerClient) UpdateTaxonomy(ctx context.Context, req *datacatalogpb.UpdateTaxonomyRequest, opts ...gax.CallOption) (*datacatalogpb.Taxonomy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "taxonomy.name", url.QueryEscape(req.GetTaxonomy().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateTaxonomy[0:len(c.CallOptions.UpdateTaxonomy):len(c.CallOptions.UpdateTaxonomy)], opts...)
	var resp *datacatalogpb.Taxonomy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.UpdateTaxonomy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListTaxonomies lists all taxonomies in a project in a particular location that the caller
// has permission to view.
func (c *PolicyTagManagerClient) ListTaxonomies(ctx context.Context, req *datacatalogpb.ListTaxonomiesRequest, opts ...gax.CallOption) *TaxonomyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListTaxonomies[0:len(c.CallOptions.ListTaxonomies):len(c.CallOptions.ListTaxonomies)], opts...)
	it := &TaxonomyIterator{}
	req = proto.Clone(req).(*datacatalogpb.ListTaxonomiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*datacatalogpb.Taxonomy, string, error) {
		var resp *datacatalogpb.ListTaxonomiesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.policyTagManagerClient.ListTaxonomies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTaxonomies(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// GetTaxonomy gets a taxonomy.
func (c *PolicyTagManagerClient) GetTaxonomy(ctx context.Context, req *datacatalogpb.GetTaxonomyRequest, opts ...gax.CallOption) (*datacatalogpb.Taxonomy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetTaxonomy[0:len(c.CallOptions.GetTaxonomy):len(c.CallOptions.GetTaxonomy)], opts...)
	var resp *datacatalogpb.Taxonomy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.GetTaxonomy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreatePolicyTag creates a policy tag in a taxonomy.
func (c *PolicyTagManagerClient) CreatePolicyTag(ctx context.Context, req *datacatalogpb.CreatePolicyTagRequest, opts ...gax.CallOption) (*datacatalogpb.PolicyTag, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreatePolicyTag[0:len(c.CallOptions.CreatePolicyTag):len(c.CallOptions.CreatePolicyTag)], opts...)
	var resp *datacatalogpb.PolicyTag
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.CreatePolicyTag(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeletePolicyTag deletes a policy tag. This method also deletes
//
//   all of its descendant policy tags, if any,
//
//   the policies associated with the policy tag and its descendants, and
//
//   references from BigQuery table schema of the policy tag and its
//   descendants.
func (c *PolicyTagManagerClient) DeletePolicyTag(ctx context.Context, req *datacatalogpb.DeletePolicyTagRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeletePolicyTag[0:len(c.CallOptions.DeletePolicyTag):len(c.CallOptions.DeletePolicyTag)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.policyTagManagerClient.DeletePolicyTag(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// UpdatePolicyTag updates a policy tag. This method can update the policy tag’s display
// name, description, and parent policy tag.
func (c *PolicyTagManagerClient) UpdatePolicyTag(ctx context.Context, req *datacatalogpb.UpdatePolicyTagRequest, opts ...gax.CallOption) (*datacatalogpb.PolicyTag, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "policy_tag.name", url.QueryEscape(req.GetPolicyTag().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdatePolicyTag[0:len(c.CallOptions.UpdatePolicyTag):len(c.CallOptions.UpdatePolicyTag)], opts...)
	var resp *datacatalogpb.PolicyTag
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.UpdatePolicyTag(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListPolicyTags lists all policy tags in a taxonomy.
func (c *PolicyTagManagerClient) ListPolicyTags(ctx context.Context, req *datacatalogpb.ListPolicyTagsRequest, opts ...gax.CallOption) *PolicyTagIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListPolicyTags[0:len(c.CallOptions.ListPolicyTags):len(c.CallOptions.ListPolicyTags)], opts...)
	it := &PolicyTagIterator{}
	req = proto.Clone(req).(*datacatalogpb.ListPolicyTagsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*datacatalogpb.PolicyTag, string, error) {
		var resp *datacatalogpb.ListPolicyTagsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.policyTagManagerClient.ListPolicyTags(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPolicyTags(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// GetPolicyTag gets a policy tag.
func (c *PolicyTagManagerClient) GetPolicyTag(ctx context.Context, req *datacatalogpb.GetPolicyTagRequest, opts ...gax.CallOption) (*datacatalogpb.PolicyTag, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetPolicyTag[0:len(c.CallOptions.GetPolicyTag):len(c.CallOptions.GetPolicyTag)], opts...)
	var resp *datacatalogpb.PolicyTag
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.GetPolicyTag(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetIamPolicy gets the IAM policy for a policy tag or a taxonomy.
func (c *PolicyTagManagerClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetIamPolicy[0:len(c.CallOptions.GetIamPolicy):len(c.CallOptions.GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SetIamPolicy sets the IAM policy for a policy tag or a taxonomy.
func (c *PolicyTagManagerClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SetIamPolicy[0:len(c.CallOptions.SetIamPolicy):len(c.CallOptions.SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TestIamPermissions returns the permissions that a caller has on a specified policy tag or
// taxonomy.
func (c *PolicyTagManagerClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.TestIamPermissions[0:len(c.CallOptions.TestIamPermissions):len(c.CallOptions.TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PolicyTagIterator manages a stream of *datacatalogpb.PolicyTag.
type PolicyTagIterator struct {
	items    []*datacatalogpb.PolicyTag
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.PolicyTag, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PolicyTagIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PolicyTagIterator) Next() (*datacatalogpb.PolicyTag, error) {
	var item *datacatalogpb.PolicyTag
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PolicyTagIterator) bufLen() int {
	return len(it.items)
}

func (it *PolicyTagIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TaxonomyIterator manages a stream of *datacatalogpb.Taxonomy.
type TaxonomyIterator struct {
	items    []*datacatalogpb.Taxonomy
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.Taxonomy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TaxonomyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TaxonomyIterator) Next() (*datacatalogpb.Taxonomy, error) {
	var item *datacatalogpb.Taxonomy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TaxonomyIterator) bufLen() int {
	return len(it.items)
}

func (it *TaxonomyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
