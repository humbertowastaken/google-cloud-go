// Copyright 2022 Google LLC
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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newBackendBucketsClientHook clientHook

// BackendBucketsCallOptions contains the retry settings for each method of BackendBucketsClient.
type BackendBucketsCallOptions struct {
	AddSignedUrlKey       []gax.CallOption
	Delete                []gax.CallOption
	DeleteSignedUrlKey    []gax.CallOption
	Get                   []gax.CallOption
	Insert                []gax.CallOption
	List                  []gax.CallOption
	Patch                 []gax.CallOption
	SetEdgeSecurityPolicy []gax.CallOption
	Update                []gax.CallOption
}

func defaultBackendBucketsRESTCallOptions() *BackendBucketsCallOptions {
	return &BackendBucketsCallOptions{
		AddSignedUrlKey:       []gax.CallOption{},
		Delete:                []gax.CallOption{},
		DeleteSignedUrlKey:    []gax.CallOption{},
		Get:                   []gax.CallOption{},
		Insert:                []gax.CallOption{},
		List:                  []gax.CallOption{},
		Patch:                 []gax.CallOption{},
		SetEdgeSecurityPolicy: []gax.CallOption{},
		Update:                []gax.CallOption{},
	}
}

// internalBackendBucketsClient is an interface that defines the methods available from Google Compute Engine API.
type internalBackendBucketsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddSignedUrlKey(context.Context, *computepb.AddSignedUrlKeyBackendBucketRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeleteBackendBucketRequest, ...gax.CallOption) (*Operation, error)
	DeleteSignedUrlKey(context.Context, *computepb.DeleteSignedUrlKeyBackendBucketRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetBackendBucketRequest, ...gax.CallOption) (*computepb.BackendBucket, error)
	Insert(context.Context, *computepb.InsertBackendBucketRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListBackendBucketsRequest, ...gax.CallOption) *BackendBucketIterator
	Patch(context.Context, *computepb.PatchBackendBucketRequest, ...gax.CallOption) (*Operation, error)
	SetEdgeSecurityPolicy(context.Context, *computepb.SetEdgeSecurityPolicyBackendBucketRequest, ...gax.CallOption) (*Operation, error)
	Update(context.Context, *computepb.UpdateBackendBucketRequest, ...gax.CallOption) (*Operation, error)
}

// BackendBucketsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The BackendBuckets API.
type BackendBucketsClient struct {
	// The internal transport-dependent client.
	internalClient internalBackendBucketsClient

	// The call options for this service.
	CallOptions *BackendBucketsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BackendBucketsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BackendBucketsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BackendBucketsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddSignedUrlKey adds a key for validating requests with signed URLs for this backend bucket.
func (c *BackendBucketsClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AddSignedUrlKey(ctx, req, opts...)
}

// Delete deletes the specified BackendBucket resource.
func (c *BackendBucketsClient) Delete(ctx context.Context, req *computepb.DeleteBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DeleteSignedUrlKey deletes a key for validating requests with signed URLs for this backend bucket.
func (c *BackendBucketsClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.DeleteSignedUrlKey(ctx, req, opts...)
}

// Get returns the specified BackendBucket resource. Gets a list of available backend buckets by making a list() request.
func (c *BackendBucketsClient) Get(ctx context.Context, req *computepb.GetBackendBucketRequest, opts ...gax.CallOption) (*computepb.BackendBucket, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a BackendBucket resource in the specified project using the data included in the request.
func (c *BackendBucketsClient) Insert(ctx context.Context, req *computepb.InsertBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of BackendBucket resources available to the specified project.
func (c *BackendBucketsClient) List(ctx context.Context, req *computepb.ListBackendBucketsRequest, opts ...gax.CallOption) *BackendBucketIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified BackendBucket resource with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *BackendBucketsClient) Patch(ctx context.Context, req *computepb.PatchBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetEdgeSecurityPolicy sets the edge security policy for the specified backend bucket.
func (c *BackendBucketsClient) SetEdgeSecurityPolicy(ctx context.Context, req *computepb.SetEdgeSecurityPolicyBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetEdgeSecurityPolicy(ctx, req, opts...)
}

// Update updates the specified BackendBucket resource with the data included in the request.
func (c *BackendBucketsClient) Update(ctx context.Context, req *computepb.UpdateBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type backendBucketsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// operationClient is used to call the operation-specific management service.
	operationClient *GlobalOperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing BackendBucketsClient
	CallOptions **BackendBucketsCallOptions
}

// NewBackendBucketsRESTClient creates a new backend buckets rest client.
//
// The BackendBuckets API.
func NewBackendBucketsRESTClient(ctx context.Context, opts ...option.ClientOption) (*BackendBucketsClient, error) {
	clientOpts := append(defaultBackendBucketsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultBackendBucketsRESTCallOptions()
	c := &backendBucketsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	o := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opC, err := NewGlobalOperationsRESTClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	c.operationClient = opC

	return &BackendBucketsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultBackendBucketsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *backendBucketsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *backendBucketsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	if err := c.operationClient.Close(); err != nil {
		return err
	}
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *backendBucketsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddSignedUrlKey adds a key for validating requests with signed URLs for this backend bucket.
func (c *backendBucketsRESTClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSignedUrlKeyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v/addSignedUrlKey", req.GetProject(), req.GetBackendBucket())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).AddSignedUrlKey[0:len((*c.CallOptions).AddSignedUrlKey):len((*c.CallOptions).AddSignedUrlKey)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// Delete deletes the specified BackendBucket resource.
func (c *backendBucketsRESTClient) Delete(ctx context.Context, req *computepb.DeleteBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v", req.GetProject(), req.GetBackendBucket())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// DeleteSignedUrlKey deletes a key for validating requests with signed URLs for this backend bucket.
func (c *backendBucketsRESTClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v/deleteSignedUrlKey", req.GetProject(), req.GetBackendBucket())

	params := url.Values{}
	params.Add("keyName", fmt.Sprintf("%v", req.GetKeyName()))
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).DeleteSignedUrlKey[0:len((*c.CallOptions).DeleteSignedUrlKey):len((*c.CallOptions).DeleteSignedUrlKey)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// Get returns the specified BackendBucket resource. Gets a list of available backend buckets by making a list() request.
func (c *backendBucketsRESTClient) Get(ctx context.Context, req *computepb.GetBackendBucketRequest, opts ...gax.CallOption) (*computepb.BackendBucket, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v", req.GetProject(), req.GetBackendBucket())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.BackendBucket{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Insert creates a BackendBucket resource in the specified project using the data included in the request.
func (c *backendBucketsRESTClient) Insert(ctx context.Context, req *computepb.InsertBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetBackendBucketResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// List retrieves the list of BackendBucket resources available to the specified project.
func (c *backendBucketsRESTClient) List(ctx context.Context, req *computepb.ListBackendBucketsRequest, opts ...gax.CallOption) *BackendBucketIterator {
	it := &BackendBucketIterator{}
	req = proto.Clone(req).(*computepb.ListBackendBucketsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.BackendBucket, string, error) {
		resp := &computepb.BackendBucketList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch updates the specified BackendBucket resource with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *backendBucketsRESTClient) Patch(ctx context.Context, req *computepb.PatchBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetBackendBucketResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v", req.GetProject(), req.GetBackendBucket())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// SetEdgeSecurityPolicy sets the edge security policy for the specified backend bucket.
func (c *backendBucketsRESTClient) SetEdgeSecurityPolicy(ctx context.Context, req *computepb.SetEdgeSecurityPolicyBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v/setEdgeSecurityPolicy", req.GetProject(), req.GetBackendBucket())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).SetEdgeSecurityPolicy[0:len((*c.CallOptions).SetEdgeSecurityPolicy):len((*c.CallOptions).SetEdgeSecurityPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// Update updates the specified BackendBucket resource with the data included in the request.
func (c *backendBucketsRESTClient) Update(ctx context.Context, req *computepb.UpdateBackendBucketRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetBackendBucketResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendBuckets/%v", req.GetProject(), req.GetBackendBucket())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Update[0:len((*c.CallOptions).Update):len((*c.CallOptions).Update)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// BackendBucketIterator manages a stream of *computepb.BackendBucket.
type BackendBucketIterator struct {
	items    []*computepb.BackendBucket
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.BackendBucket, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *BackendBucketIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *BackendBucketIterator) Next() (*computepb.BackendBucket, error) {
	var item *computepb.BackendBucket
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *BackendBucketIterator) bufLen() int {
	return len(it.items)
}

func (it *BackendBucketIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
