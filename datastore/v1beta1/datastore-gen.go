// Package datastore provides access to the Google Cloud Datastore API.
//
// See https://developers.google.com/datastore/
//
// Usage example:
//
//   import "github.com/jfcote87/api2/datastore/v1beta1"
//   ...
//   datastoreService, err := datastore.New(oauthHttpClient)
package datastore

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

const apiId = "datastore:v1beta1"
const apiName = "datastore"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/datastore/v1beta1/datasets/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Cloud Datastore data
	DatastoreScope = "https://www.googleapis.com/auth/datastore"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Datasets = NewDatasetsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Datasets *DatasetsService
}

func NewDatasetsService(s *Service) *DatasetsService {
	rs := &DatasetsService{s: s}
	return rs
}

type DatasetsService struct {
	s *Service
}

type AllocateIdsRequest struct {
	// Keys: A list of keys with incomplete key paths to allocate IDs for.
	// No key may be reserved/read-only.
	Keys []*Key `json:"keys,omitempty"`
}

type AllocateIdsResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`

	// Keys: The keys specified in the request (in the same order), each
	// with its key path completed with a newly allocated ID.
	Keys []*Key `json:"keys,omitempty"`
}

type BeginTransactionRequest struct {
	// IsolationLevel: The transaction isolation level. Either snapshot or
	// serializable. The default isolation level is snapshot isolation,
	// which means that another transaction may not concurrently modify the
	// data that is modified by this transaction. Optionally, a transaction
	// can request to be made serializable which means that another
	// transaction cannot concurrently modify the data that is read or
	// modified by this transaction.
	IsolationLevel string `json:"isolationLevel,omitempty"`
}

type BeginTransactionResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`

	// Transaction: The transaction identifier (always present).
	Transaction string `json:"transaction,omitempty"`
}

type BlindWriteRequest struct {
	// Mutation: The mutation to perform.
	Mutation *Mutation `json:"mutation,omitempty"`
}

type BlindWriteResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`

	// MutationResult: The result of performing the mutation (always
	// present).
	MutationResult *MutationResult `json:"mutationResult,omitempty"`
}

type CommitRequest struct {
	IgnoreReadOnly bool `json:"ignoreReadOnly,omitempty"`

	// Mutation: The mutation to perform. Optional.
	Mutation *Mutation `json:"mutation,omitempty"`

	// Transaction: The transaction identifier, returned by a call to
	// beginTransaction. Must be set when mode is TRANSACTIONAL.
	Transaction string `json:"transaction,omitempty"`
}

type CommitResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`

	// MutationResult: The result of performing the mutation (if any).
	MutationResult *MutationResult `json:"mutationResult,omitempty"`
}

type CompositeFilter struct {
	// Filters: The list of filters to combine. Must contain at least one
	// filter.
	Filters []*Filter `json:"filters,omitempty"`

	// Operator: The operator for combining multiple filters. Only "and" is
	// currently supported.
	Operator string `json:"operator,omitempty"`
}

type Entity struct {
	// Key: The entity's key.
	//
	// An entity must have a key, unless otherwise
	// documented (for example, an entity in Value.entityValue may have no
	// key). An entity's kind is its key's path's last element's kind, or
	// null if it has no key.
	Key *Key `json:"key,omitempty"`

	// Properties: The entity's properties.
	Properties map[string]Property `json:"properties,omitempty"`
}

type EntityResult struct {
	// Entity: The resulting entity.
	Entity *Entity `json:"entity,omitempty"`
}

type Filter struct {
	// CompositeFilter: A composite filter.
	CompositeFilter *CompositeFilter `json:"compositeFilter,omitempty"`

	// PropertyFilter: A filter on a property.
	PropertyFilter *PropertyFilter `json:"propertyFilter,omitempty"`
}

type GqlQuery struct {
	// AllowLiteral: When false, the query string must not contain a
	// literal.
	AllowLiteral bool `json:"allowLiteral,omitempty"`

	// NameArgs: A named argument must set field GqlQueryArg.name. No two
	// named arguments may have the same name. For each non-reserved named
	// binding site in the query string, there must be a named argument with
	// that name, but not necessarily the inverse.
	NameArgs []*GqlQueryArg `json:"nameArgs,omitempty"`

	// NumberArgs: Numbered binding site @1 references the first numbered
	// argument, effectively using 1-based indexing, rather than the usual
	// 0. A numbered argument must NOT set field GqlQueryArg.name. For each
	// binding site numbered i in query_string, there must be an ith
	// numbered argument. The inverse must also be true.
	NumberArgs []*GqlQueryArg `json:"numberArgs,omitempty"`

	// QueryString: The query string.
	QueryString string `json:"queryString,omitempty"`
}

type GqlQueryArg struct {
	Cursor string `json:"cursor,omitempty"`

	// Name: Must match regex "[A-Za-z_$][A-Za-z_$0-9]*". Must not match
	// regex "__.*__". Must not be "".
	Name string `json:"name,omitempty"`

	Value *Value `json:"value,omitempty"`
}

type Key struct {
	// PartitionId: Entities are partitioned into subsets, currently
	// identified by a dataset (usually implicitly specified by the project)
	// and namespace ID. Queries are scoped to a single partition.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Path: The entity path. An entity path consists of one or more
	// elements composed of a kind and a string or numerical identifier,
	// which identify entities. The first element identifies a root entity,
	// the second element identifies a child of the root entity, the third
	// element a child of the second entity, and so forth. The entities
	// identified by all prefixes of the path are called the element's
	// ancestors. An entity path is always fully complete: ALL of the
	// entity's ancestors are required to be in the path along with the
	// entity identifier itself. The only exception is that in some
	// documented cases, the identifier in the last path element (for the
	// entity) itself may be omitted. A path can never be empty. The path
	// can have at most 100 elements.
	Path []*KeyPathElement `json:"path,omitempty"`
}

type KeyPathElement struct {
	// Id: The ID of the entity. Never equal to zero. Values less than zero
	// are discouraged and will not be supported in the future.
	Id int64 `json:"id,omitempty,string"`

	// Kind: The kind of the entity. A kind matching regex "__.*__" is
	// reserved/read-only. A kind must not contain more than 500 characters.
	// Cannot be "".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the entity. A name matching regex "__.*__" is
	// reserved/read-only. A name must not be more than 500 characters.
	// Cannot be "".
	Name string `json:"name,omitempty"`
}

type KindExpression struct {
	// Name: The name of the kind.
	Name string `json:"name,omitempty"`
}

type LookupRequest struct {
	// Keys: Keys of entities to look up from the datastore.
	Keys []*Key `json:"keys,omitempty"`

	// ReadOptions: Options for this lookup request. Optional.
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`
}

type LookupResponse struct {
	// Deferred: A list of keys that were not looked up due to resource
	// constraints.
	Deferred []*Key `json:"deferred,omitempty"`

	// Found: Entities found.
	Found []*EntityResult `json:"found,omitempty"`

	Header *ResponseHeader `json:"header,omitempty"`

	// Missing: Entities not found, with only the key populated.
	Missing []*EntityResult `json:"missing,omitempty"`
}

type Mutation struct {
	// Delete: Keys of entities to delete. Each key must have a complete key
	// path and must not be reserved/read-only.
	Delete []*Key `json:"delete,omitempty"`

	// Force: Ignore a user specified read-only period. Optional.
	Force bool `json:"force,omitempty"`

	// Insert: Entities to insert. Each inserted entity's key must have a
	// complete path and must not be reserved/read-only.
	Insert []*Entity `json:"insert,omitempty"`

	// InsertAutoId: Insert entities with a newly allocated ID. Each
	// inserted entity's key must omit the final identifier in its path and
	// must not be reserved/read-only.
	InsertAutoId []*Entity `json:"insertAutoId,omitempty"`

	// Update: Entities to update. Each updated entity's key must have a
	// complete path and must not be reserved/read-only.
	Update []*Entity `json:"update,omitempty"`

	// Upsert: Entities to upsert. Each upserted entity's key must have a
	// complete path and must not be reserved/read-only.
	Upsert []*Entity `json:"upsert,omitempty"`
}

type MutationResult struct {
	// IndexUpdates: Number of index writes.
	IndexUpdates int64 `json:"indexUpdates,omitempty"`

	// InsertAutoIdKeys: Keys for insertAutoId entities. One per entity from
	// the request, in the same order.
	InsertAutoIdKeys []*Key `json:"insertAutoIdKeys,omitempty"`
}

type PartitionId struct {
	// DatasetId: The dataset ID.
	DatasetId string `json:"datasetId,omitempty"`

	// Namespace: The namespace.
	Namespace string `json:"namespace,omitempty"`
}

type Property struct {
	// Multi: If this property contains a list of values. Input values may
	// explicitly set multi to false, but otherwise false is always
	// represented by omitting multi.
	Multi bool `json:"multi,omitempty"`

	// Values: The value(s) of the property. When multi is false there is
	// always exactly one value. When multi is true there are always one or
	// more values. Each value can have only one value property populated.
	// For example, you cannot have a values list of { values: [ {
	// integerValue: 22, stringValue: "a" } ] }, but you can have { multi:
	// true, values: [ { integerValue: 22 }, { stringValue: "a" } ] }.
	Values []*Value `json:"values,omitempty"`
}

type PropertyExpression struct {
	// AggregationFunction: The aggregation function to apply to the
	// property. Optional. Can only be used when grouping by at least one
	// property. Must then be set on all properties in the projection that
	// are not being grouped by. Aggregation functions: first selects the
	// first result as determined by the query's order.
	AggregationFunction string `json:"aggregationFunction,omitempty"`

	// Property: The property to project.
	Property *PropertyReference `json:"property,omitempty"`
}

type PropertyFilter struct {
	// Operator: The operator to filter by. One of lessThan,
	// lessThanOrEqual, greaterThan, greaterThanOrEqual, equal, or
	// hasAncestor.
	Operator string `json:"operator,omitempty"`

	// Property: The property to filter by.
	Property *PropertyReference `json:"property,omitempty"`

	// Value: The value to compare the property to.
	Value *Value `json:"value,omitempty"`
}

type PropertyOrder struct {
	// Direction: The direction to order by. One of ascending or descending.
	// Optional, defaults to ascending.
	Direction string `json:"direction,omitempty"`

	// Property: The property to order by.
	Property *PropertyReference `json:"property,omitempty"`
}

type PropertyReference struct {
	// Name: The name of the property.
	Name string `json:"name,omitempty"`
}

type Query struct {
	// EndCursor: An ending point for the query results. Optional. Query
	// cursors are returned in query result batches.
	EndCursor string `json:"endCursor,omitempty"`

	// Filter: The filter to apply (optional).
	Filter *Filter `json:"filter,omitempty"`

	// GroupBy: The properties to group by (if empty, no grouping is applied
	// to the result set).
	GroupBy []*PropertyReference `json:"groupBy,omitempty"`

	// Kinds: The kinds to query (if empty, returns entities from all
	// kinds).
	Kinds []*KindExpression `json:"kinds,omitempty"`

	// Limit: The maximum number of results to return. Applies after all
	// other constraints. Optional.
	Limit int64 `json:"limit,omitempty"`

	// Offset: The number of results to skip. Applies before limit, but
	// after all other constraints (optional, defaults to 0).
	Offset int64 `json:"offset,omitempty"`

	// Order: The order to apply to the query results (if empty, order is
	// unspecified).
	Order []*PropertyOrder `json:"order,omitempty"`

	// Projection: The projection to return. If not set the entire entity is
	// returned.
	Projection []*PropertyExpression `json:"projection,omitempty"`

	// StartCursor: A starting point for the query results. Optional. Query
	// cursors are returned in query result batches.
	StartCursor string `json:"startCursor,omitempty"`
}

type QueryResultBatch struct {
	// EndCursor: A cursor that points to the position after the last result
	// in the batch. May be absent. TODO(arfuller): Once all plans produce
	// cursors update documentation here.
	EndCursor string `json:"endCursor,omitempty"`

	// EntityResultType: The result type for every entity in entityResults.
	// full for full entities, projection for entities with only projected
	// properties, keyOnly for entities with only a key.
	EntityResultType string `json:"entityResultType,omitempty"`

	// EntityResults: The results for this batch.
	EntityResults []*EntityResult `json:"entityResults,omitempty"`

	// MoreResults: The state of the query after the current batch. One of
	// notFinished, moreResultsAfterLimit, noMoreResults.
	MoreResults string `json:"moreResults,omitempty"`

	// SkippedResults: The number of results skipped because of
	// Query.offset.
	SkippedResults int64 `json:"skippedResults,omitempty"`
}

type ReadOptions struct {
	// ReadConsistency: The read consistency to use. One of default, strong,
	// or eventual. Cannot be set when transaction is set. Lookup and
	// ancestor queries default to strong, global queries default to
	// eventual and cannot be set to strong. Optional. Default is default.
	ReadConsistency string `json:"readConsistency,omitempty"`

	// Transaction: The transaction to use. Optional.
	Transaction string `json:"transaction,omitempty"`
}

type ResponseHeader struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "datastore#responseHeader".
	Kind string `json:"kind,omitempty"`
}

type RollbackRequest struct {
	// Transaction: The transaction identifier, returned by a call to
	// beginTransaction.
	Transaction string `json:"transaction,omitempty"`
}

type RollbackResponse struct {
	Header *ResponseHeader `json:"header,omitempty"`
}

type RunQueryRequest struct {
	// GqlQuery: The GQL query to run. Either this field or field query must
	// be set, but not both.
	GqlQuery *GqlQuery `json:"gqlQuery,omitempty"`

	// PartitionId: Entities are partitioned into subsets, identified by a
	// dataset (usually implicitly specified by the project) and namespace
	// ID. Queries are scoped to a single partition. This partition ID is
	// normalized with the standard default context partition ID, but all
	// other partition IDs in RunQueryRequest are normalized with this
	// partition ID as the context partition ID.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Query: The query to run. Either this field or field gql_query must be
	// set, but not both.
	Query *Query `json:"query,omitempty"`

	// ReadOptions: The options for this query.
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`
}

type RunQueryResponse struct {
	// Batch: A batch of query results (always present).
	Batch *QueryResultBatch `json:"batch,omitempty"`

	Header *ResponseHeader `json:"header,omitempty"`
}

type Value struct {
	// BlobKeyValue: A blob key value.
	BlobKeyValue string `json:"blobKeyValue,omitempty"`

	// BlobValue: A blob value. May be a maximum of 1,000,000 bytes. When
	// indexed is true, may have at most 500 bytes.
	BlobValue string `json:"blobValue,omitempty"`

	// BooleanValue: A boolean value.
	BooleanValue bool `json:"booleanValue,omitempty"`

	// DateTimeValue: A timestamp value.
	DateTimeValue string `json:"dateTimeValue,omitempty"`

	// DoubleValue: A double value.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// EntityValue: An entity value. May have no key. May have a key with an
	// incomplete key path. May have a reserved/read-only key.
	EntityValue *Entity `json:"entityValue,omitempty"`

	// Indexed: If the value should be indexed.
	//
	// The indexed property may be
	// set for a null value. When indexed is true, stringValue is limited to
	// 500 characters and the blob value is limited to 500 bytes. Input
	// values by default have indexed set to true; however, you can
	// explicitly set indexed to true if you want. (An output value never
	// has indexed explicitly set to true.) If a value is itself an entity,
	// it cannot have indexed set to true.
	Indexed bool `json:"indexed,omitempty"`

	// IntegerValue: An integer value.
	IntegerValue int64 `json:"integerValue,omitempty,string"`

	// KeyValue: A key value.
	KeyValue *Key `json:"keyValue,omitempty"`

	// Meaning: The meaning field is reserved and should not be used.
	Meaning int64 `json:"meaning,omitempty"`

	// StringValue: A UTF-8 encoded string value. When indexed is true, may
	// have at most 500 characters.
	StringValue string `json:"stringValue,omitempty"`
}

// method id "datastore.datasets.allocateIds":

type DatasetsAllocateIdsCall struct {
	s                  *Service
	datasetId          string
	allocateidsrequest *AllocateIdsRequest
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
	context_           context.Context
}

// AllocateIds: Allocate IDs for incomplete keys (useful for referencing
// an entity before it is inserted).

func (r *DatasetsService) AllocateIds(datasetId string, allocateidsrequest *AllocateIdsRequest) *DatasetsAllocateIdsCall {
	return &DatasetsAllocateIdsCall{
		s:                  r.s,
		datasetId:          datasetId,
		allocateidsrequest: allocateidsrequest,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "{datasetId}/allocateIds",
		context_:           googleapi.NoContext,
	}
}
func (c *DatasetsAllocateIdsCall) Context(ctx context.Context) *DatasetsAllocateIdsCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsAllocateIdsCall) Fields(s ...googleapi.Field) *DatasetsAllocateIdsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsAllocateIdsCall) Do() (*AllocateIdsResponse, error) {
	var returnValue *AllocateIdsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.allocateidsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Allocate IDs for incomplete keys (useful for referencing an entity before it is inserted).",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.allocateIds",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/allocateIds",
	//   "request": {
	//     "$ref": "AllocateIdsRequest"
	//   },
	//   "response": {
	//     "$ref": "AllocateIdsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.beginTransaction":

type DatasetsBeginTransactionCall struct {
	s                       *Service
	datasetId               string
	begintransactionrequest *BeginTransactionRequest
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
	context_                context.Context
}

// BeginTransaction: Begin a new transaction.

func (r *DatasetsService) BeginTransaction(datasetId string, begintransactionrequest *BeginTransactionRequest) *DatasetsBeginTransactionCall {
	return &DatasetsBeginTransactionCall{
		s:                       r.s,
		datasetId:               datasetId,
		begintransactionrequest: begintransactionrequest,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "{datasetId}/beginTransaction",
		context_:                googleapi.NoContext,
	}
}
func (c *DatasetsBeginTransactionCall) Context(ctx context.Context) *DatasetsBeginTransactionCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsBeginTransactionCall) Fields(s ...googleapi.Field) *DatasetsBeginTransactionCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsBeginTransactionCall) Do() (*BeginTransactionResponse, error) {
	var returnValue *BeginTransactionResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.begintransactionrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Begin a new transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.beginTransaction",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/beginTransaction",
	//   "request": {
	//     "$ref": "BeginTransactionRequest"
	//   },
	//   "response": {
	//     "$ref": "BeginTransactionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.blindWrite":

type DatasetsBlindWriteCall struct {
	s                 *Service
	datasetId         string
	blindwriterequest *BlindWriteRequest
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
	context_          context.Context
}

// BlindWrite: Create, delete or modify some entities outside a
// transaction.

func (r *DatasetsService) BlindWrite(datasetId string, blindwriterequest *BlindWriteRequest) *DatasetsBlindWriteCall {
	return &DatasetsBlindWriteCall{
		s:                 r.s,
		datasetId:         datasetId,
		blindwriterequest: blindwriterequest,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "{datasetId}/blindWrite",
		context_:          googleapi.NoContext,
	}
}
func (c *DatasetsBlindWriteCall) Context(ctx context.Context) *DatasetsBlindWriteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsBlindWriteCall) Fields(s ...googleapi.Field) *DatasetsBlindWriteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsBlindWriteCall) Do() (*BlindWriteResponse, error) {
	var returnValue *BlindWriteResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.blindwriterequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Create, delete or modify some entities outside a transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.blindWrite",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/blindWrite",
	//   "request": {
	//     "$ref": "BlindWriteRequest"
	//   },
	//   "response": {
	//     "$ref": "BlindWriteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.commit":

type DatasetsCommitCall struct {
	s             *Service
	datasetId     string
	commitrequest *CommitRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Commit: Commit a transaction, optionally creating, deleting or
// modifying some entities.

func (r *DatasetsService) Commit(datasetId string, commitrequest *CommitRequest) *DatasetsCommitCall {
	return &DatasetsCommitCall{
		s:             r.s,
		datasetId:     datasetId,
		commitrequest: commitrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{datasetId}/commit",
		context_:      googleapi.NoContext,
	}
}
func (c *DatasetsCommitCall) Context(ctx context.Context) *DatasetsCommitCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsCommitCall) Fields(s ...googleapi.Field) *DatasetsCommitCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsCommitCall) Do() (*CommitResponse, error) {
	var returnValue *CommitResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.commitrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Commit a transaction, optionally creating, deleting or modifying some entities.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.commit",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/commit",
	//   "request": {
	//     "$ref": "CommitRequest"
	//   },
	//   "response": {
	//     "$ref": "CommitResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.lookup":

type DatasetsLookupCall struct {
	s             *Service
	datasetId     string
	lookuprequest *LookupRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Lookup: Look up some entities by key.

func (r *DatasetsService) Lookup(datasetId string, lookuprequest *LookupRequest) *DatasetsLookupCall {
	return &DatasetsLookupCall{
		s:             r.s,
		datasetId:     datasetId,
		lookuprequest: lookuprequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{datasetId}/lookup",
		context_:      googleapi.NoContext,
	}
}
func (c *DatasetsLookupCall) Context(ctx context.Context) *DatasetsLookupCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsLookupCall) Fields(s ...googleapi.Field) *DatasetsLookupCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsLookupCall) Do() (*LookupResponse, error) {
	var returnValue *LookupResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.lookuprequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Look up some entities by key.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.lookup",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/lookup",
	//   "request": {
	//     "$ref": "LookupRequest"
	//   },
	//   "response": {
	//     "$ref": "LookupResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.rollback":

type DatasetsRollbackCall struct {
	s               *Service
	datasetId       string
	rollbackrequest *RollbackRequest
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Rollback: Roll back a transaction.

func (r *DatasetsService) Rollback(datasetId string, rollbackrequest *RollbackRequest) *DatasetsRollbackCall {
	return &DatasetsRollbackCall{
		s:               r.s,
		datasetId:       datasetId,
		rollbackrequest: rollbackrequest,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{datasetId}/rollback",
		context_:        googleapi.NoContext,
	}
}
func (c *DatasetsRollbackCall) Context(ctx context.Context) *DatasetsRollbackCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsRollbackCall) Fields(s ...googleapi.Field) *DatasetsRollbackCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsRollbackCall) Do() (*RollbackResponse, error) {
	var returnValue *RollbackResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.rollbackrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Roll back a transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.rollback",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/rollback",
	//   "request": {
	//     "$ref": "RollbackRequest"
	//   },
	//   "response": {
	//     "$ref": "RollbackResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.runQuery":

type DatasetsRunQueryCall struct {
	s               *Service
	datasetId       string
	runqueryrequest *RunQueryRequest
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// RunQuery: Query for entities.

func (r *DatasetsService) RunQuery(datasetId string, runqueryrequest *RunQueryRequest) *DatasetsRunQueryCall {
	return &DatasetsRunQueryCall{
		s:               r.s,
		datasetId:       datasetId,
		runqueryrequest: runqueryrequest,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{datasetId}/runQuery",
		context_:        googleapi.NoContext,
	}
}
func (c *DatasetsRunQueryCall) Context(ctx context.Context) *DatasetsRunQueryCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsRunQueryCall) Fields(s ...googleapi.Field) *DatasetsRunQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsRunQueryCall) Do() (*RunQueryResponse, error) {
	var returnValue *RunQueryResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.runqueryrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Query for entities.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.runQuery",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/runQuery",
	//   "request": {
	//     "$ref": "RunQueryRequest"
	//   },
	//   "response": {
	//     "$ref": "RunQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
