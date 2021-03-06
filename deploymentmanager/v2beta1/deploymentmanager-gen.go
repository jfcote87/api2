// Package deploymentmanager provides access to the Google Cloud Deployment Manager API V2.
//
// See https://developers.google.com/deployment-manager/
//
// Usage example:
//
//   import "github.com/jfcote87/api2/deploymentmanager/v2beta1"
//   ...
//   deploymentmanagerService, err := deploymentmanager.New(oauthHttpClient)
package deploymentmanager

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

const apiId = "deploymentmanager:v2beta1"
const apiName = "deploymentmanager"
const apiVersion = "v2beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/deploymentmanager/v2beta1/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

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
	s.Deployments = NewDeploymentsService(s)
	s.Manifests = NewManifestsService(s)
	s.Operations = NewOperationsService(s)
	s.Resources = NewResourcesService(s)
	s.Types = NewTypesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Deployments *DeploymentsService

	Manifests *ManifestsService

	Operations *OperationsService

	Resources *ResourcesService

	Types *TypesService
}

func NewDeploymentsService(s *Service) *DeploymentsService {
	rs := &DeploymentsService{s: s}
	return rs
}

type DeploymentsService struct {
	s *Service
}

func NewManifestsService(s *Service) *ManifestsService {
	rs := &ManifestsService{s: s}
	return rs
}

type ManifestsService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewResourcesService(s *Service) *ResourcesService {
	rs := &ResourcesService{s: s}
	return rs
}

type ResourcesService struct {
	s *Service
}

func NewTypesService(s *Service) *TypesService {
	rs := &TypesService{s: s}
	return rs
}

type TypesService struct {
	s *Service
}

type Deployment struct {
	// Description: ! An optional user-provided description of the
	// deployment.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Manifest: ! [Output Only] URL of the manifest representing the full
	// configuration ! of this deployment.
	Manifest string `json:"manifest,omitempty"`

	// Name: ! The name of the deployment, which must be unique within the
	// project.
	Name string `json:"name,omitempty"`

	// TargetConfig: ! [Input Only] The YAML configuration to use in
	// processing this deployment. ! ! When you create a deployment, the
	// server creates a new manifest with the ! given YAML configuration and
	// sets the `manifest` property to the URL of ! the manifest resource.
	TargetConfig string `json:"targetConfig,omitempty"`
}

type DeploymentsListResponse struct {
	// Deployments: ! The deployments contained in this response.
	Deployments []*Deployment `json:"deployments,omitempty"`

	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Manifest struct {
	// Config: v2beta1: YAML with config - described above v2beta2: YAML +
	// templates. ! The YAML configuration for this manifest.
	Config string `json:"config,omitempty"`

	// EvaluatedConfig: ! [Output Only] The fully-expanded configuration
	// file, including any ! templates and references.
	EvaluatedConfig string `json:"evaluatedConfig,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Name: ! [Output Only] The name of the manifest.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] Self link for the manifest.
	SelfLink string `json:"selfLink,omitempty"`
}

type ManifestsListResponse struct {
	// Manifests: ! Manifests contained in this list response.
	Manifests []*Manifest `json:"manifests,omitempty"`

	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Operation struct {
	// CreationTimestamp: ! [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// EndTime: ! [Output Only] The time that this operation was completed.
	// This is in ! RFC3339 format.
	EndTime string `json:"endTime,omitempty"`

	// Error: ! [Output Only] If errors occurred during processing of this
	// operation, ! this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	// HttpErrorMessage: ! [Output Only] If operation fails, the HTTP error
	// message returned, ! e.g. NOT FOUND.
	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	// HttpErrorStatusCode: ! [Output Only] If operation fails, the HTTP
	// error status code returned, ! e.g. 404.
	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: ! [Output Only] Unique identifier for the resource; defined by
	// the server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: ! [Output Only] The time that this operation was
	// requested. ! This is in RFC 3339 format.
	InsertTime string `json:"insertTime,omitempty"`

	// Name: ! [Output Only] Name of the operation.
	Name string `json:"name,omitempty"`

	// OperationType: ! [Output Only] Type of the operation. Examples
	// include "insert", or ! "delete"
	OperationType string `json:"operationType,omitempty"`

	// Progress: ! [Output Only] An optional progress indicator that ranges
	// from 0 to 100. ! There is no requirement that this be linear or
	// support any granularity ! of operations. This should not be used to
	// guess at when the operation will ! be complete. This number should be
	// monotonically increasing as the ! operation progresses.
	Progress int64 `json:"progress,omitempty"`

	// SelfLink: [Output Only] Self link for the manifest.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: ! [Output Only] The time that this operation was started
	// by the server. ! This is in RFC 3339 format.
	StartTime string `json:"startTime,omitempty"`

	// Status: ! [Output Only] Status of the operation. Can be one of the
	// following: ! "PENDING", "RUNNING", or "DONE".
	Status string `json:"status,omitempty"`

	// StatusMessage: ! [Output Only] An optional textual description of the
	// current status of ! the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: ! [Output Only] Unique target id which identifies a
	// particular ! incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: ! [Output Only] URL of the resource the operation is
	// mutating.
	TargetLink string `json:"targetLink,omitempty"`

	// User: ! [Output Only] User who requested the operation, for example !
	// "user@example.com"
	User string `json:"user,omitempty"`

	// Warnings: ! [Output Only] If warning messages generated during
	// processing of this ! operation, this field will be populated.
	Warnings []*OperationWarnings `json:"warnings,omitempty"`
}

type OperationError struct {
	// Errors: ! The array of errors encountered while processing this
	// operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	// Code: ! The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: ! Indicates the field in the request which caused the
	// error. ! This property is optional.
	Location string `json:"location,omitempty"`

	// Message: ! An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	// Code: ! The warning type identifier for this warning.
	Code interface{} `json:"code,omitempty"`

	// Data: ! Metadata for this warning in 'key: value' format.
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: ! Optional human-readable details for this warning.
	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	// Key: ! A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: ! A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type OperationsListResponse struct {
	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: ! Operations contained in this list response.
	Operations []*Operation `json:"operations,omitempty"`
}

type Resource struct {
	// Errors: ! [Output Only] A list of any errors that occurred during
	// deployment.
	Errors []string `json:"errors,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Intent: ! [Output Only] The intended state of the resource.
	Intent string `json:"intent,omitempty"`

	// Manifest: ! [Output Only] URL of the manifest representing the
	// current configuration ! of this resource.
	Manifest string `json:"manifest,omitempty"`

	// Name: ! [Output Only] The name of the resource as it appears in the
	// YAML config.
	Name string `json:"name,omitempty"`

	// State: ! [Output Only] The state of the resource.
	State string `json:"state,omitempty"`

	// Type: ! [Output Only] The type of the resource, for example !
	// ?compute.v1.instance?, or ?replicaPools.v1beta2.instanceGroupManager?
	Type string `json:"type,omitempty"`

	// Url: ! [Output Only] The URL of the actual resource.
	Url string `json:"url,omitempty"`
}

type ResourcesListResponse struct {
	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Resources: ! Resources contained in this list response.
	Resources []*Resource `json:"resources,omitempty"`
}

type Type struct {
	// Name: ! Name of the type.
	Name string `json:"name,omitempty"`
}

type TypesListResponse struct {
	// Types: ! Types supported by Deployment Manager
	Types []*Type `json:"types,omitempty"`
}

// method id "deploymentmanager.deployments.delete":

type DeploymentsDeleteCall struct {
	s             *Service
	project       string
	deployment    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: ! Deletes a deployment and all of the resources in the
// deployment.

func (r *DeploymentsService) Delete(project string, deployment string) *DeploymentsDeleteCall {
	return &DeploymentsDeleteCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments/{deployment}",
		context_:      googleapi.NoContext,
	}
}
func (c *DeploymentsDeleteCall) Context(ctx context.Context) *DeploymentsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsDeleteCall) Fields(s ...googleapi.Field) *DeploymentsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DeploymentsDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Deletes a deployment and all of the resources in the deployment.",
	//   "httpMethod": "DELETE",
	//   "id": "deploymentmanager.deployments.delete",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.get":

type DeploymentsGetCall struct {
	s             *Service
	project       string
	deployment    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: ! Gets information about a specific deployment.

func (r *DeploymentsService) Get(project string, deployment string) *DeploymentsGetCall {
	return &DeploymentsGetCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments/{deployment}",
		context_:      googleapi.NoContext,
	}
}
func (c *DeploymentsGetCall) Context(ctx context.Context) *DeploymentsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsGetCall) Fields(s ...googleapi.Field) *DeploymentsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DeploymentsGetCall) Do() (*Deployment, error) {
	var returnValue *Deployment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Gets information about a specific deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "response": {
	//     "$ref": "Deployment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.insert":

type DeploymentsInsertCall struct {
	s             *Service
	project       string
	deployment    *Deployment
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: ! Creates a deployment and all of the resources described by
// the ! deployment manifest.

func (r *DeploymentsService) Insert(project string, deployment *Deployment) *DeploymentsInsertCall {
	return &DeploymentsInsertCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments",
		context_:      googleapi.NoContext,
	}
}
func (c *DeploymentsInsertCall) Context(ctx context.Context) *DeploymentsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsInsertCall) Fields(s ...googleapi.Field) *DeploymentsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DeploymentsInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.deployment,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Creates a deployment and all of the resources described by the ! deployment manifest.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments",
	//   "request": {
	//     "$ref": "Deployment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.list":

type DeploymentsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: ! Lists all deployments for a given project.

func (r *DeploymentsService) List(project string) *DeploymentsListCall {
	return &DeploymentsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *DeploymentsListCall) MaxResults(maxResults int64) *DeploymentsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *DeploymentsListCall) PageToken(pageToken string) *DeploymentsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DeploymentsListCall) Context(ctx context.Context) *DeploymentsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsListCall) Fields(s ...googleapi.Field) *DeploymentsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DeploymentsListCall) Do() (*DeploymentsListResponse, error) {
	var returnValue *DeploymentsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Lists all deployments for a given project.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments",
	//   "response": {
	//     "$ref": "DeploymentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.manifests.get":

type ManifestsGetCall struct {
	s             *Service
	project       string
	deployment    string
	manifest      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: ! Gets information about a specific manifest.

func (r *ManifestsService) Get(project string, deployment string, manifest string) *ManifestsGetCall {
	return &ManifestsGetCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		manifest:      manifest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments/{deployment}/manifests/{manifest}",
		context_:      googleapi.NoContext,
	}
}
func (c *ManifestsGetCall) Context(ctx context.Context) *ManifestsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManifestsGetCall) Fields(s ...googleapi.Field) *ManifestsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ManifestsGetCall) Do() (*Manifest, error) {
	var returnValue *Manifest
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
		"manifest":   c.manifest,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Gets information about a specific manifest.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.manifests.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment",
	//     "manifest"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "manifest": {
	//       "description": "! The name of the manifest for this request.",
	//       "location": "path",
	//       "pattern": "[-a-z0-9]{1,61}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/manifests/{manifest}",
	//   "response": {
	//     "$ref": "Manifest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.manifests.list":

type ManifestsListCall struct {
	s             *Service
	project       string
	deployment    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: ! Lists all manifests for a given deployment.

func (r *ManifestsService) List(project string, deployment string) *ManifestsListCall {
	return &ManifestsListCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments/{deployment}/manifests",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *ManifestsListCall) MaxResults(maxResults int64) *ManifestsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *ManifestsListCall) PageToken(pageToken string) *ManifestsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ManifestsListCall) Context(ctx context.Context) *ManifestsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManifestsListCall) Fields(s ...googleapi.Field) *ManifestsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ManifestsListCall) Do() (*ManifestsListResponse, error) {
	var returnValue *ManifestsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Lists all manifests for a given deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.manifests.list",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/manifests",
	//   "response": {
	//     "$ref": "ManifestsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.operations.get":

type OperationsGetCall struct {
	s             *Service
	project       string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: ! Gets information about a specific Operation.

func (r *OperationsService) Get(project string, operation string) *OperationsGetCall {
	return &OperationsGetCall{
		s:             r.s,
		project:       project,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OperationsGetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"operation": c.operation,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Gets information about a specific Operation.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.operations.get",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "! The name of the operation for this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.operations.list":

type OperationsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: ! Lists all Operations for a project.

func (r *OperationsService) List(project string) *OperationsListCall {
	return &OperationsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/operations",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *OperationsListCall) MaxResults(maxResults int64) *OperationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *OperationsListCall) Context(ctx context.Context) *OperationsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OperationsListCall) Do() (*OperationsListResponse, error) {
	var returnValue *OperationsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Lists all Operations for a project.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.operations.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations",
	//   "response": {
	//     "$ref": "OperationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.resources.get":

type ResourcesGetCall struct {
	s             *Service
	project       string
	deployment    string
	resource      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: ! Gets information about a single resource.

func (r *ResourcesService) Get(project string, deployment string, resource string) *ResourcesGetCall {
	return &ResourcesGetCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		resource:      resource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments/{deployment}/resources/{resource}",
		context_:      googleapi.NoContext,
	}
}
func (c *ResourcesGetCall) Context(ctx context.Context) *ResourcesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesGetCall) Fields(s ...googleapi.Field) *ResourcesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ResourcesGetCall) Do() (*Resource, error) {
	var returnValue *Resource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
		"resource":   c.resource,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Gets information about a single resource.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.resources.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment",
	//     "resource"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resource": {
	//       "description": "! The name of the resource for this request.",
	//       "location": "path",
	//       "pattern": "[-a-z0-9]{1,61}",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/resources/{resource}",
	//   "response": {
	//     "$ref": "Resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.resources.list":

type ResourcesListCall struct {
	s             *Service
	project       string
	deployment    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: ! Lists all resources in a given deployment.

func (r *ResourcesService) List(project string, deployment string) *ResourcesListCall {
	return &ResourcesListCall{
		s:             r.s,
		project:       project,
		deployment:    deployment,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/deployments/{deployment}/resources",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *ResourcesListCall) MaxResults(maxResults int64) *ResourcesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *ResourcesListCall) PageToken(pageToken string) *ResourcesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ResourcesListCall) Context(ctx context.Context) *ResourcesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesListCall) Fields(s ...googleapi.Field) *ResourcesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ResourcesListCall) Do() (*ResourcesListResponse, error) {
	var returnValue *ResourcesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Lists all resources in a given deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.resources.list",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/resources",
	//   "response": {
	//     "$ref": "ResourcesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.types.list":

type TypesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: ! Lists all Types for Deployment Manager.

func (r *TypesService) List(project string) *TypesListCall {
	return &TypesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/types",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *TypesListCall) MaxResults(maxResults int64) *TypesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *TypesListCall) PageToken(pageToken string) *TypesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TypesListCall) Context(ctx context.Context) *TypesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypesListCall) Fields(s ...googleapi.Field) *TypesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TypesListCall) Do() (*TypesListResponse, error) {
	var returnValue *TypesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "! Lists all Types for Deployment Manager.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.types.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/types",
	//   "response": {
	//     "$ref": "TypesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}
