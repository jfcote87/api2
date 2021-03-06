// Package resourceviews provides access to the Resource Views API.
//
// See https://developers.google.com/compute/
//
// Usage example:
//
//   import "github.com/jfcote87/api2/resourceviews/v1beta1"
//   ...
//   resourceviewsService, err := resourceviews.New(oauthHttpClient)
package resourceviews

import (
	"errors"
	"fmt"
	"github.com/jfcote87/api2/googleapi"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "resourceviews:v1beta1"
const apiName = "resourceviews"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/resourceviews/v1beta1/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Compute Engine resources
	ComputeScope = "https://www.googleapis.com/auth/compute"

	// View your Google Compute Engine resources
	ComputeReadonlyScope = "https://www.googleapis.com/auth/compute.readonly"

	// View and manage your Google Cloud Platform management resources and
	// deployment status information
	NdevCloudmanScope = "https://www.googleapis.com/auth/ndev.cloudman"

	// View your Google Cloud Platform management resources and deployment
	// status information
	NdevCloudmanReadonlyScope = "https://www.googleapis.com/auth/ndev.cloudman.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.RegionViews = NewRegionViewsService(s)
	s.ZoneViews = NewZoneViewsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	RegionViews *RegionViewsService

	ZoneViews *ZoneViewsService
}

func NewRegionViewsService(s *Service) *RegionViewsService {
	rs := &RegionViewsService{s: s}
	return rs
}

type RegionViewsService struct {
	s *Service
}

func NewZoneViewsService(s *Service) *ZoneViewsService {
	rs := &ZoneViewsService{s: s}
	return rs
}

type ZoneViewsService struct {
	s *Service
}

type Label struct {
	// Key: Key of the label.
	Key string `json:"key,omitempty"`

	// Value: Value of the label.
	Value string `json:"value,omitempty"`
}

type RegionViewsAddResourcesRequest struct {
	// Resources: The list of resources to be added.
	Resources []string `json:"resources,omitempty"`
}

type RegionViewsInsertResponse struct {
	// Resource: The resource view object inserted.
	Resource *ResourceView `json:"resource,omitempty"`
}

type RegionViewsListResourcesResponse struct {
	// Members: The resources in the view.
	Members []string `json:"members,omitempty"`

	// NextPageToken: A token used for pagination.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type RegionViewsListResponse struct {
	// NextPageToken: A token used for pagination.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResourceViews: The list of resource views that meet the criteria.
	ResourceViews []*ResourceView `json:"resourceViews,omitempty"`
}

type RegionViewsRemoveResourcesRequest struct {
	// Resources: The list of resources to be removed.
	Resources []string `json:"resources,omitempty"`
}

type ResourceView struct {
	// CreationTime: The creation time of the resource view.
	CreationTime string `json:"creationTime,omitempty"`

	// Description: The detailed description of the resource view.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] The ID of the resource view.
	Id string `json:"id,omitempty"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Labels: The labels for events.
	Labels []*Label `json:"labels,omitempty"`

	// LastModified: The last modified time of the view. Not supported yet.
	LastModified string `json:"lastModified,omitempty"`

	// Members: A list of all resources in the resource view.
	Members []string `json:"members,omitempty"`

	// Name: The name of the resource view.
	Name string `json:"name,omitempty"`

	// NumMembers: The total number of resources in the resource view.
	NumMembers int64 `json:"numMembers,omitempty"`

	// SelfLink: [Output Only] A self-link to the resource view.
	SelfLink string `json:"selfLink,omitempty"`
}

type ZoneViewsAddResourcesRequest struct {
	// Resources: The list of resources to be added.
	Resources []string `json:"resources,omitempty"`
}

type ZoneViewsInsertResponse struct {
	// Resource: The resource view object that has been inserted.
	Resource *ResourceView `json:"resource,omitempty"`
}

type ZoneViewsListResourcesResponse struct {
	// Members: The full URL of resources in the view.
	Members []string `json:"members,omitempty"`

	// NextPageToken: A token used for pagination.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ZoneViewsListResponse struct {
	// NextPageToken: A token used for pagination.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResourceViews: The result that contains all resource views that meet
	// the criteria.
	ResourceViews []*ResourceView `json:"resourceViews,omitempty"`
}

type ZoneViewsRemoveResourcesRequest struct {
	// Resources: The list of resources to be removed.
	Resources []string `json:"resources,omitempty"`
}

// method id "resourceviews.regionViews.addresources":

type RegionViewsAddresourcesCall struct {
	s                              *Service
	projectName                    string
	region                         string
	resourceViewName               string
	regionviewsaddresourcesrequest *RegionViewsAddResourcesRequest
	caller_                        googleapi.Caller
	params_                        url.Values
	pathTemplate_                  string
	context_                       context.Context
}

// Addresources: Add resources to the view.

func (r *RegionViewsService) Addresources(projectName string, region string, resourceViewName string, regionviewsaddresourcesrequest *RegionViewsAddResourcesRequest) *RegionViewsAddresourcesCall {
	return &RegionViewsAddresourcesCall{
		s:                              r.s,
		projectName:                    projectName,
		region:                         region,
		resourceViewName:               resourceViewName,
		regionviewsaddresourcesrequest: regionviewsaddresourcesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/regions/{region}/resourceViews/{resourceViewName}/addResources",
		context_:      googleapi.NoContext,
	}
}
func (c *RegionViewsAddresourcesCall) Context(ctx context.Context) *RegionViewsAddresourcesCall {
	c.context_ = ctx
	return c
}

func (c *RegionViewsAddresourcesCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"region":           c.region,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.regionviewsaddresourcesrequest,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Add resources to the view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.regionViews.addresources",
	//   "parameterOrder": [
	//     "projectName",
	//     "region",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews/{resourceViewName}/addResources",
	//   "request": {
	//     "$ref": "RegionViewsAddResourcesRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.regionViews.delete":

type RegionViewsDeleteCall struct {
	s                *Service
	projectName      string
	region           string
	resourceViewName string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Delete: Delete a resource view.

func (r *RegionViewsService) Delete(projectName string, region string, resourceViewName string) *RegionViewsDeleteCall {
	return &RegionViewsDeleteCall{
		s:                r.s,
		projectName:      projectName,
		region:           region,
		resourceViewName: resourceViewName,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{projectName}/regions/{region}/resourceViews/{resourceViewName}",
		context_:         googleapi.NoContext,
	}
}
func (c *RegionViewsDeleteCall) Context(ctx context.Context) *RegionViewsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *RegionViewsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"region":           c.region,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Delete a resource view.",
	//   "httpMethod": "DELETE",
	//   "id": "resourceviews.regionViews.delete",
	//   "parameterOrder": [
	//     "projectName",
	//     "region",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews/{resourceViewName}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.regionViews.get":

type RegionViewsGetCall struct {
	s                *Service
	projectName      string
	region           string
	resourceViewName string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Get: Get the information of a resource view.

func (r *RegionViewsService) Get(projectName string, region string, resourceViewName string) *RegionViewsGetCall {
	return &RegionViewsGetCall{
		s:                r.s,
		projectName:      projectName,
		region:           region,
		resourceViewName: resourceViewName,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{projectName}/regions/{region}/resourceViews/{resourceViewName}",
		context_:         googleapi.NoContext,
	}
}
func (c *RegionViewsGetCall) Context(ctx context.Context) *RegionViewsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionViewsGetCall) Fields(s ...googleapi.Field) *RegionViewsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionViewsGetCall) Do() (*ResourceView, error) {
	var returnValue *ResourceView
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"region":           c.region,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Get the information of a resource view.",
	//   "httpMethod": "GET",
	//   "id": "resourceviews.regionViews.get",
	//   "parameterOrder": [
	//     "projectName",
	//     "region",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews/{resourceViewName}",
	//   "response": {
	//     "$ref": "ResourceView"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "resourceviews.regionViews.insert":

type RegionViewsInsertCall struct {
	s             *Service
	projectName   string
	region        string
	resourceview  *ResourceView
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Create a resource view.

func (r *RegionViewsService) Insert(projectName string, region string, resourceview *ResourceView) *RegionViewsInsertCall {
	return &RegionViewsInsertCall{
		s:             r.s,
		projectName:   projectName,
		region:        region,
		resourceview:  resourceview,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/regions/{region}/resourceViews",
		context_:      googleapi.NoContext,
	}
}
func (c *RegionViewsInsertCall) Context(ctx context.Context) *RegionViewsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionViewsInsertCall) Fields(s ...googleapi.Field) *RegionViewsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionViewsInsertCall) Do() (*RegionViewsInsertResponse, error) {
	var returnValue *RegionViewsInsertResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName": c.projectName,
		"region":      c.region,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.resourceview,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Create a resource view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.regionViews.insert",
	//   "parameterOrder": [
	//     "projectName",
	//     "region"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews",
	//   "request": {
	//     "$ref": "ResourceView"
	//   },
	//   "response": {
	//     "$ref": "RegionViewsInsertResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.regionViews.list":

type RegionViewsListCall struct {
	s             *Service
	projectName   string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: List resource views.

func (r *RegionViewsService) List(projectName string, region string) *RegionViewsListCall {
	return &RegionViewsListCall{
		s:             r.s,
		projectName:   projectName,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/regions/{region}/resourceViews",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Acceptable values are 0 to 5000, inclusive.
// (Default: 5000)
func (c *RegionViewsListCall) MaxResults(maxResults int64) *RegionViewsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// nextPageToken returned by a previous list request. This token can be
// used to request the next page of results from a previous list
// request.
func (c *RegionViewsListCall) PageToken(pageToken string) *RegionViewsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *RegionViewsListCall) Context(ctx context.Context) *RegionViewsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionViewsListCall) Fields(s ...googleapi.Field) *RegionViewsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionViewsListCall) Do() (*RegionViewsListResponse, error) {
	var returnValue *RegionViewsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName": c.projectName,
		"region":      c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "List resource views.",
	//   "httpMethod": "GET",
	//   "id": "resourceviews.regionViews.list",
	//   "parameterOrder": [
	//     "projectName",
	//     "region"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "5000",
	//       "description": "Maximum count of results to be returned. Acceptable values are 0 to 5000, inclusive. (Default: 5000)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a nextPageToken returned by a previous list request. This token can be used to request the next page of results from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews",
	//   "response": {
	//     "$ref": "RegionViewsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "resourceviews.regionViews.listresources":

type RegionViewsListresourcesCall struct {
	s                *Service
	projectName      string
	region           string
	resourceViewName string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Listresources: List the resources in the view.

func (r *RegionViewsService) Listresources(projectName string, region string, resourceViewName string) *RegionViewsListresourcesCall {
	return &RegionViewsListresourcesCall{
		s:                r.s,
		projectName:      projectName,
		region:           region,
		resourceViewName: resourceViewName,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{projectName}/regions/{region}/resourceViews/{resourceViewName}/resources",
		context_:         googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Acceptable values are 0 to 5000, inclusive.
// (Default: 5000)
func (c *RegionViewsListresourcesCall) MaxResults(maxResults int64) *RegionViewsListresourcesCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// nextPageToken returned by a previous list request. This token can be
// used to request the next page of results from a previous list
// request.
func (c *RegionViewsListresourcesCall) PageToken(pageToken string) *RegionViewsListresourcesCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *RegionViewsListresourcesCall) Context(ctx context.Context) *RegionViewsListresourcesCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionViewsListresourcesCall) Fields(s ...googleapi.Field) *RegionViewsListresourcesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionViewsListresourcesCall) Do() (*RegionViewsListResourcesResponse, error) {
	var returnValue *RegionViewsListResourcesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"region":           c.region,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "List the resources in the view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.regionViews.listresources",
	//   "parameterOrder": [
	//     "projectName",
	//     "region",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "5000",
	//       "description": "Maximum count of results to be returned. Acceptable values are 0 to 5000, inclusive. (Default: 5000)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a nextPageToken returned by a previous list request. This token can be used to request the next page of results from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews/{resourceViewName}/resources",
	//   "response": {
	//     "$ref": "RegionViewsListResourcesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "resourceviews.regionViews.removeresources":

type RegionViewsRemoveresourcesCall struct {
	s                                 *Service
	projectName                       string
	region                            string
	resourceViewName                  string
	regionviewsremoveresourcesrequest *RegionViewsRemoveResourcesRequest
	caller_                           googleapi.Caller
	params_                           url.Values
	pathTemplate_                     string
	context_                          context.Context
}

// Removeresources: Remove resources from the view.

func (r *RegionViewsService) Removeresources(projectName string, region string, resourceViewName string, regionviewsremoveresourcesrequest *RegionViewsRemoveResourcesRequest) *RegionViewsRemoveresourcesCall {
	return &RegionViewsRemoveresourcesCall{
		s:                                 r.s,
		projectName:                       projectName,
		region:                            region,
		resourceViewName:                  resourceViewName,
		regionviewsremoveresourcesrequest: regionviewsremoveresourcesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/regions/{region}/resourceViews/{resourceViewName}/removeResources",
		context_:      googleapi.NoContext,
	}
}
func (c *RegionViewsRemoveresourcesCall) Context(ctx context.Context) *RegionViewsRemoveresourcesCall {
	c.context_ = ctx
	return c
}

func (c *RegionViewsRemoveresourcesCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"region":           c.region,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.regionviewsremoveresourcesrequest,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Remove resources from the view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.regionViews.removeresources",
	//   "parameterOrder": [
	//     "projectName",
	//     "region",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The region name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/regions/{region}/resourceViews/{resourceViewName}/removeResources",
	//   "request": {
	//     "$ref": "RegionViewsRemoveResourcesRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.addresources":

type ZoneViewsAddresourcesCall struct {
	s                            *Service
	projectName                  string
	zone                         string
	resourceViewName             string
	zoneviewsaddresourcesrequest *ZoneViewsAddResourcesRequest
	caller_                      googleapi.Caller
	params_                      url.Values
	pathTemplate_                string
	context_                     context.Context
}

// Addresources: Add resources to the view.

func (r *ZoneViewsService) Addresources(projectName string, zone string, resourceViewName string, zoneviewsaddresourcesrequest *ZoneViewsAddResourcesRequest) *ZoneViewsAddresourcesCall {
	return &ZoneViewsAddresourcesCall{
		s:                            r.s,
		projectName:                  projectName,
		zone:                         zone,
		resourceViewName:             resourceViewName,
		zoneviewsaddresourcesrequest: zoneviewsaddresourcesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/zones/{zone}/resourceViews/{resourceViewName}/addResources",
		context_:      googleapi.NoContext,
	}
}
func (c *ZoneViewsAddresourcesCall) Context(ctx context.Context) *ZoneViewsAddresourcesCall {
	c.context_ = ctx
	return c
}

func (c *ZoneViewsAddresourcesCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"zone":             c.zone,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.zoneviewsaddresourcesrequest,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Add resources to the view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.zoneViews.addresources",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews/{resourceViewName}/addResources",
	//   "request": {
	//     "$ref": "ZoneViewsAddResourcesRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.delete":

type ZoneViewsDeleteCall struct {
	s                *Service
	projectName      string
	zone             string
	resourceViewName string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Delete: Delete a resource view.

func (r *ZoneViewsService) Delete(projectName string, zone string, resourceViewName string) *ZoneViewsDeleteCall {
	return &ZoneViewsDeleteCall{
		s:                r.s,
		projectName:      projectName,
		zone:             zone,
		resourceViewName: resourceViewName,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{projectName}/zones/{zone}/resourceViews/{resourceViewName}",
		context_:         googleapi.NoContext,
	}
}
func (c *ZoneViewsDeleteCall) Context(ctx context.Context) *ZoneViewsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *ZoneViewsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"zone":             c.zone,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Delete a resource view.",
	//   "httpMethod": "DELETE",
	//   "id": "resourceviews.zoneViews.delete",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews/{resourceViewName}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.get":

type ZoneViewsGetCall struct {
	s                *Service
	projectName      string
	zone             string
	resourceViewName string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Get: Get the information of a zonal resource view.

func (r *ZoneViewsService) Get(projectName string, zone string, resourceViewName string) *ZoneViewsGetCall {
	return &ZoneViewsGetCall{
		s:                r.s,
		projectName:      projectName,
		zone:             zone,
		resourceViewName: resourceViewName,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{projectName}/zones/{zone}/resourceViews/{resourceViewName}",
		context_:         googleapi.NoContext,
	}
}
func (c *ZoneViewsGetCall) Context(ctx context.Context) *ZoneViewsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneViewsGetCall) Fields(s ...googleapi.Field) *ZoneViewsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneViewsGetCall) Do() (*ResourceView, error) {
	var returnValue *ResourceView
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"zone":             c.zone,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Get the information of a zonal resource view.",
	//   "httpMethod": "GET",
	//   "id": "resourceviews.zoneViews.get",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews/{resourceViewName}",
	//   "response": {
	//     "$ref": "ResourceView"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.insert":

type ZoneViewsInsertCall struct {
	s             *Service
	projectName   string
	zone          string
	resourceview  *ResourceView
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Create a resource view.

func (r *ZoneViewsService) Insert(projectName string, zone string, resourceview *ResourceView) *ZoneViewsInsertCall {
	return &ZoneViewsInsertCall{
		s:             r.s,
		projectName:   projectName,
		zone:          zone,
		resourceview:  resourceview,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/zones/{zone}/resourceViews",
		context_:      googleapi.NoContext,
	}
}
func (c *ZoneViewsInsertCall) Context(ctx context.Context) *ZoneViewsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneViewsInsertCall) Fields(s ...googleapi.Field) *ZoneViewsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneViewsInsertCall) Do() (*ZoneViewsInsertResponse, error) {
	var returnValue *ZoneViewsInsertResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName": c.projectName,
		"zone":        c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.resourceview,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Create a resource view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.zoneViews.insert",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews",
	//   "request": {
	//     "$ref": "ResourceView"
	//   },
	//   "response": {
	//     "$ref": "ZoneViewsInsertResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.list":

type ZoneViewsListCall struct {
	s             *Service
	projectName   string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: List resource views.

func (r *ZoneViewsService) List(projectName string, zone string) *ZoneViewsListCall {
	return &ZoneViewsListCall{
		s:             r.s,
		projectName:   projectName,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/zones/{zone}/resourceViews",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Acceptable values are 0 to 5000, inclusive.
// (Default: 5000)
func (c *ZoneViewsListCall) MaxResults(maxResults int64) *ZoneViewsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// nextPageToken returned by a previous list request. This token can be
// used to request the next page of results from a previous list
// request.
func (c *ZoneViewsListCall) PageToken(pageToken string) *ZoneViewsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ZoneViewsListCall) Context(ctx context.Context) *ZoneViewsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneViewsListCall) Fields(s ...googleapi.Field) *ZoneViewsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneViewsListCall) Do() (*ZoneViewsListResponse, error) {
	var returnValue *ZoneViewsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName": c.projectName,
		"zone":        c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "List resource views.",
	//   "httpMethod": "GET",
	//   "id": "resourceviews.zoneViews.list",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "5000",
	//       "description": "Maximum count of results to be returned. Acceptable values are 0 to 5000, inclusive. (Default: 5000)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a nextPageToken returned by a previous list request. This token can be used to request the next page of results from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews",
	//   "response": {
	//     "$ref": "ZoneViewsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.listresources":

type ZoneViewsListresourcesCall struct {
	s                *Service
	projectName      string
	zone             string
	resourceViewName string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Listresources: List the resources of the resource view.

func (r *ZoneViewsService) Listresources(projectName string, zone string, resourceViewName string) *ZoneViewsListresourcesCall {
	return &ZoneViewsListresourcesCall{
		s:                r.s,
		projectName:      projectName,
		zone:             zone,
		resourceViewName: resourceViewName,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{projectName}/zones/{zone}/resourceViews/{resourceViewName}/resources",
		context_:         googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Acceptable values are 0 to 5000, inclusive.
// (Default: 5000)
func (c *ZoneViewsListresourcesCall) MaxResults(maxResults int64) *ZoneViewsListresourcesCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// nextPageToken returned by a previous list request. This token can be
// used to request the next page of results from a previous list
// request.
func (c *ZoneViewsListresourcesCall) PageToken(pageToken string) *ZoneViewsListresourcesCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ZoneViewsListresourcesCall) Context(ctx context.Context) *ZoneViewsListresourcesCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneViewsListresourcesCall) Fields(s ...googleapi.Field) *ZoneViewsListresourcesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneViewsListresourcesCall) Do() (*ZoneViewsListResourcesResponse, error) {
	var returnValue *ZoneViewsListResourcesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"zone":             c.zone,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "List the resources of the resource view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.zoneViews.listresources",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "5000",
	//       "description": "Maximum count of results to be returned. Acceptable values are 0 to 5000, inclusive. (Default: 5000)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a nextPageToken returned by a previous list request. This token can be used to request the next page of results from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews/{resourceViewName}/resources",
	//   "response": {
	//     "$ref": "ZoneViewsListResourcesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "resourceviews.zoneViews.removeresources":

type ZoneViewsRemoveresourcesCall struct {
	s                               *Service
	projectName                     string
	zone                            string
	resourceViewName                string
	zoneviewsremoveresourcesrequest *ZoneViewsRemoveResourcesRequest
	caller_                         googleapi.Caller
	params_                         url.Values
	pathTemplate_                   string
	context_                        context.Context
}

// Removeresources: Remove resources from the view.

func (r *ZoneViewsService) Removeresources(projectName string, zone string, resourceViewName string, zoneviewsremoveresourcesrequest *ZoneViewsRemoveResourcesRequest) *ZoneViewsRemoveresourcesCall {
	return &ZoneViewsRemoveresourcesCall{
		s:                               r.s,
		projectName:                     projectName,
		zone:                            zone,
		resourceViewName:                resourceViewName,
		zoneviewsremoveresourcesrequest: zoneviewsremoveresourcesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectName}/zones/{zone}/resourceViews/{resourceViewName}/removeResources",
		context_:      googleapi.NoContext,
	}
}
func (c *ZoneViewsRemoveresourcesCall) Context(ctx context.Context) *ZoneViewsRemoveresourcesCall {
	c.context_ = ctx
	return c
}

func (c *ZoneViewsRemoveresourcesCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectName":      c.projectName,
		"zone":             c.zone,
		"resourceViewName": c.resourceViewName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.zoneviewsremoveresourcesrequest,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Remove resources from the view.",
	//   "httpMethod": "POST",
	//   "id": "resourceviews.zoneViews.removeresources",
	//   "parameterOrder": [
	//     "projectName",
	//     "zone",
	//     "resourceViewName"
	//   ],
	//   "parameters": {
	//     "projectName": {
	//       "description": "The project name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceViewName": {
	//       "description": "The name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone name of the resource view.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectName}/zones/{zone}/resourceViews/{resourceViewName}/removeResources",
	//   "request": {
	//     "$ref": "ZoneViewsRemoveResourcesRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}
