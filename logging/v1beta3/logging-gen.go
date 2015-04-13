// Package logging provides access to the Google Cloud Logging API.
//
// Usage example:
//
//   import "github.com/jfcote87/api2/logging/v1beta3"
//   ...
//   loggingService, err := logging.New(oauthHttpClient)
package logging

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

const apiId = "logging:v1beta3"
const apiName = "logging"
const apiVersion = "v1beta3"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "logging.googleapis.com", Path: "/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.LogServices = NewProjectsLogServicesService(s)
	rs.Logs = NewProjectsLogsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	LogServices *ProjectsLogServicesService

	Logs *ProjectsLogsService
}

func NewProjectsLogServicesService(s *Service) *ProjectsLogServicesService {
	rs := &ProjectsLogServicesService{s: s}
	rs.Indexes = NewProjectsLogServicesIndexesService(s)
	rs.Sinks = NewProjectsLogServicesSinksService(s)
	return rs
}

type ProjectsLogServicesService struct {
	s *Service

	Indexes *ProjectsLogServicesIndexesService

	Sinks *ProjectsLogServicesSinksService
}

func NewProjectsLogServicesIndexesService(s *Service) *ProjectsLogServicesIndexesService {
	rs := &ProjectsLogServicesIndexesService{s: s}
	return rs
}

type ProjectsLogServicesIndexesService struct {
	s *Service
}

func NewProjectsLogServicesSinksService(s *Service) *ProjectsLogServicesSinksService {
	rs := &ProjectsLogServicesSinksService{s: s}
	return rs
}

type ProjectsLogServicesSinksService struct {
	s *Service
}

func NewProjectsLogsService(s *Service) *ProjectsLogsService {
	rs := &ProjectsLogsService{s: s}
	rs.Entries = NewProjectsLogsEntriesService(s)
	rs.Sinks = NewProjectsLogsSinksService(s)
	return rs
}

type ProjectsLogsService struct {
	s *Service

	Entries *ProjectsLogsEntriesService

	Sinks *ProjectsLogsSinksService
}

func NewProjectsLogsEntriesService(s *Service) *ProjectsLogsEntriesService {
	rs := &ProjectsLogsEntriesService{s: s}
	return rs
}

type ProjectsLogsEntriesService struct {
	s *Service
}

func NewProjectsLogsSinksService(s *Service) *ProjectsLogsSinksService {
	rs := &ProjectsLogsSinksService{s: s}
	return rs
}

type ProjectsLogsSinksService struct {
	s *Service
}

type Empty struct {
}

type ListLogServiceIndexesResponse struct {
	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the response. To get the next batch of indexes, use the
	// value of `nextPageToken` as `pageToken` in the next call of
	// `ListLogServiceIndexess`. If `nextPageToken` is empty, then there are
	// no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServiceIndexPrefixes: A list of log service index prefixes.
	ServiceIndexPrefixes []string `json:"serviceIndexPrefixes,omitempty"`
}

type ListLogServiceSinksResponse struct {
	// Sinks: The requested log service sinks. If any of the returned
	// `LogSink` objects have an empty `destination` field, then call
	// `logServices.sinks.get` to retrieve the complete `LogSink` object.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogServicesResponse struct {
	// LogServices: A list of log services.
	LogServices []*LogService `json:"logServices,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the response. To get the next batch of services, use the
	// value of `nextPageToken` as `pageToken` in the next call of
	// `ListLogServices`. If `nextPageToken` is empty, then there are no
	// more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogSinksResponse struct {
	// Sinks: The requested log sinks. If any of the returned `LogSink`
	// objects have an empty `destination` field, then call
	// `logServices.sinks.get` to retrieve the complete `LogSink` object.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogsResponse struct {
	// Logs: A list of log resources.
	Logs []*Log `json:"logs,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the response. To get the next batch of logs, use the
	// value of `nextPageToken` as `pageToken` in the next call of
	// `ListLogs`. If `nextPageToken` is empty, then there are no more
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Log struct {
	// DisplayName: Name used when displaying the log to the user (for
	// example, in a UI). Example: `"activity_log"`
	DisplayName string `json:"displayName,omitempty"`

	// Name: REQUIRED: The log's name name. Example:
	// `"compute.googleapis.com/activity_log"`.
	Name string `json:"name,omitempty"`

	// PayloadType: Type URL describing the expected payload type for the
	// log.
	PayloadType string `json:"payloadType,omitempty"`
}

type LogEntry struct {
	// InsertId: A unique ID for the log entry. If you provide this field,
	// the logging service considers other log entries in the same log with
	// the same ID as duplicates which can be removed.
	InsertId string `json:"insertId,omitempty"`

	// Log: The log to which this entry belongs. When a log entry is
	// ingested, the value of this field is set by the logging system.
	Log string `json:"log,omitempty"`

	// Metadata: Information about the log entry.
	Metadata *LogEntryMetadata `json:"metadata,omitempty"`

	// ProtoPayload: The log entry payload, represented as a protocol buffer
	// that is expressed as a JSON object. You can only pass `protoPayload`
	// values that belong to a set of approved types.
	ProtoPayload *LogEntryProtoPayload `json:"protoPayload,omitempty"`

	// StructPayload: The log entry payload, represented as a structure that
	// is expressed as a JSON object.
	StructPayload *LogEntryStructPayload `json:"structPayload,omitempty"`

	// TextPayload: The log entry payload, represented as a text string.
	TextPayload string `json:"textPayload,omitempty"`
}

type LogEntryProtoPayload struct {
}

type LogEntryStructPayload struct {
}

type LogEntryMetadata struct {
	// Labels: A set of (key, value) data that provides additional
	// information about the log entry. If the log entry is from one of the
	// Google Cloud Platform sources listed below, the indicated (key,
	// value) information must be provided: Google App Engine, service_name
	// `appengine.googleapis.com`: "appengine.googleapis.com/module_id",
	// "appengine.googleapis.com/version_id",  and one of:
	// "appengine.googleapis.com/replica_index",
	// "appengine.googleapis.com/clone_id",  or else provide the following
	// Compute Engine labels: Google Compute Engine, service_name
	// `compute.googleapis.com`: "compute.googleapis.com/resource_type",
	// "instance" "compute.googleapis.com/resource_id",
	Labels map[string]string `json:"labels,omitempty"`

	// ProjectId: The project ID of the Google Cloud Platform service that
	// created the log entry.
	ProjectId string `json:"projectId,omitempty"`

	// Region: The region name of the Google Cloud Platform service that
	// created the log entry. For example, `"us-central1"`.
	Region string `json:"region,omitempty"`

	// ServiceName: The API name of the Google Cloud Platform service that
	// created the log entry. For example, `"compute.googleapis.com"`.
	ServiceName string `json:"serviceName,omitempty"`

	// Severity: The severity of the log entry.
	Severity string `json:"severity,omitempty"`

	// Timestamp: The time the event described by the log entry occurred.
	// Timestamps must be later than January 1, 1970.
	Timestamp string `json:"timestamp,omitempty"`

	// UserId: The fully-qualified email address of the authenticated user
	// that performed or requested the action represented by the log entry.
	// If the log entry does not apply to an action taken by an
	// authenticated user, then the field should be empty.
	UserId string `json:"userId,omitempty"`

	// Zone: The zone of the Google Cloud Platform service that created the
	// log entry. For example, `"us-central1-a"`.
	Zone string `json:"zone,omitempty"`
}

type LogError struct {
	// Resource: The resource associated with the error. It may be different
	// from the sink destination. For example, the sink may point to a
	// BigQuery dataset, but the error may refer to a table resource inside
	// the dataset.
	Resource string `json:"resource,omitempty"`

	// Status: The description of the last error observed.
	Status *Status `json:"status,omitempty"`

	// TimeNanos: The last time the error was observed, in nanoseconds since
	// the Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`
}

type LogService struct {
	// IndexKeys: Label keys used when labeling log entries for this
	// service. The order of the keys is significant, with higher priority
	// keys coming earlier in the list.
	IndexKeys []string `json:"indexKeys,omitempty"`

	// Name: The service's name.
	Name string `json:"name,omitempty"`
}

type LogSink struct {
	// Destination: The resource to send log entries to. The supported sink
	// resource types are: + Google Cloud Storage:
	// `storage.googleapis.com/BUCKET` or `BUCKET.storage.googleapis.com/` +
	// Google BigQuery:
	// `bigquery.googleapis.com/projects/PROJECT/datasets/DATASET` Currently
	// the Cloud Logging API supports at most one sink for each resource
	// type per log or log service resource.
	Destination string `json:"destination,omitempty"`

	// Errors: _Output only._ All active errors found for this sink.
	Errors []*LogError `json:"errors,omitempty"`

	// Name: The name of this sink. This is a client-assigned identifier for
	// the resource. This is ignored by UpdateLogSink and
	// UpdateLogServicesSink.
	Name string `json:"name,omitempty"`
}

type Status struct {
	// Code: The status code, which should be an enum value of
	// [google.rpc.Code][].
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There will
	// be a common set of message types for APIs to use.
	Details []*StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. The user-facing error message should be localized and stored
	// in the [google.rpc.Status.details][google.rpc.Status.details] field.
	Message string `json:"message,omitempty"`
}

type StatusDetails struct {
}

type WriteLogEntriesRequest struct {
	// CommonLabels: Metadata labels that apply to all entries in this
	// request. If one of the log entries contains a (key, value) with the
	// same key that is in `commonLabels`, then the entry's (key, value)
	// overrides the one in `commonLabels`.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// Entries: Log entries to insert.
	Entries []*LogEntry `json:"entries,omitempty"`
}

type WriteLogEntriesResponse struct {
}

// method id "logging.projects.logServices.list":

type ProjectsLogServicesListCall struct {
	s             *Service
	projectsId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists log services associated with log entries ingested for a
// project.

func (r *ProjectsLogServicesService) List(projectsId string) *ProjectsLogServicesListCall {
	return &ProjectsLogServicesListCall{
		s:             r.s,
		projectsId:    projectsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices",
		context_:      googleapi.NoContext,
	}
}

// Log sets the optional parameter "log": The name of the log resource
// whose services are to be listed. log for which to list services. When
// empty, all services are listed.
func (c *ProjectsLogServicesListCall) Log(log string) *ProjectsLogServicesListCall {
	c.params_.Set("log", fmt.Sprintf("%v", log))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of `LogService` objects to return in one operation.
func (c *ProjectsLogServicesListCall) PageSize(pageSize int64) *ProjectsLogServicesListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogServices` operation.
// If `pageToken` is supplied, then the other fields of this request are
// ignored, and instead the previous `ListLogServices` operation is
// continued.
func (c *ProjectsLogServicesListCall) PageToken(pageToken string) *ProjectsLogServicesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ProjectsLogServicesListCall) Context(ctx context.Context) *ProjectsLogServicesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesListCall) Do() (*ListLogServicesResponse, error) {
	var returnValue *ListLogServicesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists log services associated with log entries ingested for a project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "description": "The name of the log resource whose services are to be listed. log for which to list services. When empty, all services are listed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of `LogService` objects to return in one operation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogServices` operation. If `pageToken` is supplied, then the other fields of this request are ignored, and instead the previous `ListLogServices` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project resource whose services are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices",
	//   "response": {
	//     "$ref": "ListLogServicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.indexes.list":

type ProjectsLogServicesIndexesListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists log service indexes associated with a log service.

func (r *ProjectsLogServicesIndexesService) List(projectsId string, logServicesId string) *ProjectsLogServicesIndexesListCall {
	return &ProjectsLogServicesIndexesListCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
		context_:      googleapi.NoContext,
	}
}

// Depth sets the optional parameter "depth": A limit to the number of
// levels of the index hierarchy that are expanded. If `depth` is 0, it
// defaults to the level specified by the prefix field (the number of
// slash separators). The default empty prefix implies a `depth` of 1.
// It is an error for `depth` to be any non-zero value less than the
// number of components in `indexPrefix`.
func (c *ProjectsLogServicesIndexesListCall) Depth(depth int64) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("depth", fmt.Sprintf("%v", depth))
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": Restricts the
// indexes returned to be those with a specified prefix. The prefix has
// the form `"/label_value/label_value/..."`, in order corresponding to
// the [`LogService
// indexKeys`][google.logging.v1.LogService.index_keys]. Non-empty
// prefixes must begin with `/` . Example prefixes: + `"/myModule/"`
// retrieves App Engine versions associated with `myModule`. The
// trailing slash terminates the value. + `"/myModule"` retrieves App
// Engine modules with names beginning with `myModule`. + `""` retrieves
// all indexes.
func (c *ProjectsLogServicesIndexesListCall) IndexPrefix(indexPrefix string) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("indexPrefix", fmt.Sprintf("%v", indexPrefix))
	return c
}

// Log sets the optional parameter "log": A log resource like
// `/projects/project_id/logs/log_name`, identifying the log for which
// to list service indexes.
func (c *ProjectsLogServicesIndexesListCall) Log(log string) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("log", fmt.Sprintf("%v", log))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of log service index resources to return in one operation.
func (c *ProjectsLogServicesIndexesListCall) PageSize(pageSize int64) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogServiceIndexes`
// operation. If `pageToken` is supplied, then the other fields of this
// request are ignored, and instead the previous `ListLogServiceIndexes`
// operation is continued.
func (c *ProjectsLogServicesIndexesListCall) PageToken(pageToken string) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ProjectsLogServicesIndexesListCall) Context(ctx context.Context) *ProjectsLogServicesIndexesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesIndexesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesIndexesListCall) Do() (*ListLogServiceIndexesResponse, error) {
	var returnValue *ListLogServiceIndexesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists log service indexes associated with a log service.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.indexes.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "description": "A limit to the number of levels of the index hierarchy that are expanded. If `depth` is 0, it defaults to the level specified by the prefix field (the number of slash separators). The default empty prefix implies a `depth` of 1. It is an error for `depth` to be any non-zero value less than the number of components in `indexPrefix`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "Restricts the indexes returned to be those with a specified prefix. The prefix has the form `\"/label_value/label_value/...\"`, in order corresponding to the [`LogService indexKeys`][google.logging.v1.LogService.index_keys]. Non-empty prefixes must begin with `/` . Example prefixes: + `\"/myModule/\"` retrieves App Engine versions associated with `myModule`. The trailing slash terminates the value. + `\"/myModule\"` retrieves App Engine modules with names beginning with `myModule`. + `\"\"` retrieves all indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "log": {
	//       "description": "A log resource like `/projects/project_id/logs/log_name`, identifying the log for which to list service indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of log service index resources to return in one operation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogServiceIndexes` operation. If `pageToken` is supplied, then the other fields of this request are ignored, and instead the previous `ListLogServiceIndexes` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. A log service resource of the form `/projects/*/logServices/*`. The service indexes of the log service are returned. Example: `\"/projects/myProj/logServices/appengine.googleapis.com\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
	//   "response": {
	//     "$ref": "ListLogServiceIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.create":

type ProjectsLogServicesSinksCreateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Create: Creates the specified log service sink resource.

func (r *ProjectsLogServicesSinksService) Create(projectsId string, logServicesId string, logsink *LogSink) *ProjectsLogServicesSinksCreateCall {
	return &ProjectsLogServicesSinksCreateCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogServicesSinksCreateCall) Context(ctx context.Context) *ProjectsLogServicesSinksCreateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksCreateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates the specified log service sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logServices.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The name of the service in which to create a sink.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.delete":

type ProjectsLogServicesSinksDeleteCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified log service sink.

func (r *ProjectsLogServicesSinksService) Delete(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksDeleteCall {
	return &ProjectsLogServicesSinksDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogServicesSinksDeleteCall) Context(ctx context.Context) *ProjectsLogServicesSinksDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified log service sink.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logServices.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.get":

type ProjectsLogServicesSinksGetCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Gets the specified log service sink resource.

func (r *ProjectsLogServicesSinksService) Get(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksGetCall {
	return &ProjectsLogServicesSinksGetCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogServicesSinksGetCall) Context(ctx context.Context) *ProjectsLogServicesSinksGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksGetCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the specified log service sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.list":

type ProjectsLogServicesSinksListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists log service sinks associated with the specified service.

func (r *ProjectsLogServicesSinksService) List(projectsId string, logServicesId string) *ProjectsLogServicesSinksListCall {
	return &ProjectsLogServicesSinksListCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogServicesSinksListCall) Context(ctx context.Context) *ProjectsLogServicesSinksListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksListCall) Do() (*ListLogServiceSinksResponse, error) {
	var returnValue *ListLogServiceSinksResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists log service sinks associated with the specified service.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The name of the service for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "response": {
	//     "$ref": "ListLogServiceSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.update":

type ProjectsLogServicesSinksUpdateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Creates or update the specified log service sink resource.

func (r *ProjectsLogServicesSinksService) Update(projectsId string, logServicesId string, sinksId string, logsink *LogSink) *ProjectsLogServicesSinksUpdateCall {
	return &ProjectsLogServicesSinksUpdateCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		sinksId:       sinksId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogServicesSinksUpdateCall) Context(ctx context.Context) *ProjectsLogServicesSinksUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksUpdateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates or update the specified log service sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logServices.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.delete":

type ProjectsLogsDeleteCall struct {
	s             *Service
	projectsId    string
	logsId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified log resource and all log entries
// contained in it.

func (r *ProjectsLogsService) Delete(projectsId string, logsId string) *ProjectsLogsDeleteCall {
	return &ProjectsLogsDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogsDeleteCall) Context(ctx context.Context) *ProjectsLogsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified log resource and all log entries contained in it.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log resource to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.list":

type ProjectsLogsListCall struct {
	s             *Service
	projectsId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists log resources belonging to the specified project.

func (r *ProjectsLogsService) List(projectsId string) *ProjectsLogsListCall {
	return &ProjectsLogsListCall{
		s:             r.s,
		projectsId:    projectsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs",
		context_:      googleapi.NoContext,
	}
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return.
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogs` operation. If
// `pageToken` is supplied, then the other fields of this request are
// ignored, and instead the previous `ListLogs` operation is continued.
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ServiceIndexPrefix sets the optional parameter "serviceIndexPrefix":
// A log service index prefix for which to list logs. Only logs
// containing entries whose metadata that includes these label values
// (associated with index keys) are returned. The prefix is a slash
// separated list of values, and need not specify all index labels. An
// empty index (or a single slash) matches all log service indexes.
func (c *ProjectsLogsListCall) ServiceIndexPrefix(serviceIndexPrefix string) *ProjectsLogsListCall {
	c.params_.Set("serviceIndexPrefix", fmt.Sprintf("%v", serviceIndexPrefix))
	return c
}

// ServiceName sets the optional parameter "serviceName": A service name
// for which to list logs. Only logs containing entries whose metadata
// includes this service name are returned. If `serviceName` and
// `serviceIndexPrefix` are both empty, then all log names are returned.
// To list all log names, regardless of service, leave both the
// `serviceName` and `serviceIndexPrefix` empty. To list log names
// containing entries with a particular service name (or explicitly
// empty service name) set `serviceName` to the desired value and
// `serviceIndexPrefix` to `"/"`.
func (c *ProjectsLogsListCall) ServiceName(serviceName string) *ProjectsLogsListCall {
	c.params_.Set("serviceName", fmt.Sprintf("%v", serviceName))
	return c
}
func (c *ProjectsLogsListCall) Context(ctx context.Context) *ProjectsLogsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsListCall) Fields(s ...googleapi.Field) *ProjectsLogsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsListCall) Do() (*ListLogsResponse, error) {
	var returnValue *ListLogsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists log resources belonging to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogs` operation. If `pageToken` is supplied, then the other fields of this request are ignored, and instead the previous `ListLogs` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project name for which to list the log resources.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "serviceIndexPrefix": {
	//       "description": "A log service index prefix for which to list logs. Only logs containing entries whose metadata that includes these label values (associated with index keys) are returned. The prefix is a slash separated list of values, and need not specify all index labels. An empty index (or a single slash) matches all log service indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceName": {
	//       "description": "A service name for which to list logs. Only logs containing entries whose metadata includes this service name are returned. If `serviceName` and `serviceIndexPrefix` are both empty, then all log names are returned. To list all log names, regardless of service, leave both the `serviceName` and `serviceIndexPrefix` empty. To list log names containing entries with a particular service name (or explicitly empty service name) set `serviceName` to the desired value and `serviceIndexPrefix` to `\"/\"`.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.entries.write":

type ProjectsLogsEntriesWriteCall struct {
	s                      *Service
	projectsId             string
	logsId                 string
	writelogentriesrequest *WriteLogEntriesRequest
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
	context_               context.Context
}

// Write: Creates one or more log entries in a log. You must supply a
// list of `LogEntry` objects, named `entries`. Each `LogEntry` object
// must contain a payload object and a `LogEntryMetadata` object that
// describes the entry. You must fill in all the fields of the entry,
// metadata, and payload. You can also supply a map, `commonLabels`,
// that supplies default (key, value) data for the
// `entries[].metadata.labels` maps, saving you the trouble of creating
// identical copies for each entry.

func (r *ProjectsLogsEntriesService) Write(projectsId string, logsId string, writelogentriesrequest *WriteLogEntriesRequest) *ProjectsLogsEntriesWriteCall {
	return &ProjectsLogsEntriesWriteCall{
		s:                      r.s,
		projectsId:             projectsId,
		logsId:                 logsId,
		writelogentriesrequest: writelogentriesrequest,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
		context_:               googleapi.NoContext,
	}
}
func (c *ProjectsLogsEntriesWriteCall) Context(ctx context.Context) *ProjectsLogsEntriesWriteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsEntriesWriteCall) Fields(s ...googleapi.Field) *ProjectsLogsEntriesWriteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsEntriesWriteCall) Do() (*WriteLogEntriesResponse, error) {
	var returnValue *WriteLogEntriesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.writelogentriesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates one or more log entries in a log. You must supply a list of `LogEntry` objects, named `entries`. Each `LogEntry` object must contain a payload object and a `LogEntryMetadata` object that describes the entry. You must fill in all the fields of the entry, metadata, and payload. You can also supply a map, `commonLabels`, that supplies default (key, value) data for the `entries[].metadata.labels` maps, saving you the trouble of creating identical copies for each entry.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.entries.write",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The name of the log resource into which to insert the log entries.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
	//   "request": {
	//     "$ref": "WriteLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.create":

type ProjectsLogsSinksCreateCall struct {
	s             *Service
	projectsId    string
	logsId        string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Create: Creates the specified log sink resource.

func (r *ProjectsLogsSinksService) Create(projectsId string, logsId string, logsink *LogSink) *ProjectsLogsSinksCreateCall {
	return &ProjectsLogsSinksCreateCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogsSinksCreateCall) Context(ctx context.Context) *ProjectsLogsSinksCreateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksCreateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates the specified log sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log in which to create a sink resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.delete":

type ProjectsLogsSinksDeleteCall struct {
	s             *Service
	projectsId    string
	logsId        string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified log sink resource.

func (r *ProjectsLogsSinksService) Delete(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksDeleteCall {
	return &ProjectsLogsSinksDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogsSinksDeleteCall) Context(ctx context.Context) *ProjectsLogsSinksDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified log sink resource.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.get":

type ProjectsLogsSinksGetCall struct {
	s             *Service
	projectsId    string
	logsId        string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Gets the specified log sink resource.

func (r *ProjectsLogsSinksService) Get(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksGetCall {
	return &ProjectsLogsSinksGetCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogsSinksGetCall) Context(ctx context.Context) *ProjectsLogsSinksGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksGetCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the specified log sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink resource to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.list":

type ProjectsLogsSinksListCall struct {
	s             *Service
	projectsId    string
	logsId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists log sinks associated with the specified log.

func (r *ProjectsLogsSinksService) List(projectsId string, logsId string) *ProjectsLogsSinksListCall {
	return &ProjectsLogsSinksListCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogsSinksListCall) Context(ctx context.Context) *ProjectsLogsSinksListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksListCall) Do() (*ListLogSinksResponse, error) {
	var returnValue *ListLogSinksResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists log sinks associated with the specified log.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "response": {
	//     "$ref": "ListLogSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.update":

type ProjectsLogsSinksUpdateCall struct {
	s             *Service
	projectsId    string
	logsId        string
	sinksId       string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Creates or updates the specified log sink resource.

func (r *ProjectsLogsSinksService) Update(projectsId string, logsId string, sinksId string, logsink *LogSink) *ProjectsLogsSinksUpdateCall {
	return &ProjectsLogsSinksUpdateCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		sinksId:       sinksId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsLogsSinksUpdateCall) Context(ctx context.Context) *ProjectsLogsSinksUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksUpdateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates or updates the specified log sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logs.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
