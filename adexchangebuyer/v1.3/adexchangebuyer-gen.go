// Package adexchangebuyer provides access to the Ad Exchange Buyer API.
//
// See https://developers.google.com/ad-exchange/buyer-rest
//
// Usage example:
//
//   import "github.com/jfcote87/api2/adexchangebuyer/v1.3"
//   ...
//   adexchangebuyerService, err := adexchangebuyer.New(oauthHttpClient)
package adexchangebuyer

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

const apiId = "adexchangebuyer:v1.3"
const apiName = "adexchangebuyer"
const apiVersion = "v1.3"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/adexchangebuyer/v1.3/"}

// OAuth2 scopes used by this API.
const (
	// Manage your Ad Exchange buyer account configuration
	AdexchangeBuyerScope = "https://www.googleapis.com/auth/adexchange.buyer"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Accounts = NewAccountsService(s)
	s.BillingInfo = NewBillingInfoService(s)
	s.Budget = NewBudgetService(s)
	s.Creatives = NewCreativesService(s)
	s.DirectDeals = NewDirectDealsService(s)
	s.PerformanceReport = NewPerformanceReportService(s)
	s.PretargetingConfig = NewPretargetingConfigService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Accounts *AccountsService

	BillingInfo *BillingInfoService

	Budget *BudgetService

	Creatives *CreativesService

	DirectDeals *DirectDealsService

	PerformanceReport *PerformanceReportService

	PretargetingConfig *PretargetingConfigService
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	return rs
}

type AccountsService struct {
	s *Service
}

func NewBillingInfoService(s *Service) *BillingInfoService {
	rs := &BillingInfoService{s: s}
	return rs
}

type BillingInfoService struct {
	s *Service
}

func NewBudgetService(s *Service) *BudgetService {
	rs := &BudgetService{s: s}
	return rs
}

type BudgetService struct {
	s *Service
}

func NewCreativesService(s *Service) *CreativesService {
	rs := &CreativesService{s: s}
	return rs
}

type CreativesService struct {
	s *Service
}

func NewDirectDealsService(s *Service) *DirectDealsService {
	rs := &DirectDealsService{s: s}
	return rs
}

type DirectDealsService struct {
	s *Service
}

func NewPerformanceReportService(s *Service) *PerformanceReportService {
	rs := &PerformanceReportService{s: s}
	return rs
}

type PerformanceReportService struct {
	s *Service
}

func NewPretargetingConfigService(s *Service) *PretargetingConfigService {
	rs := &PretargetingConfigService{s: s}
	return rs
}

type PretargetingConfigService struct {
	s *Service
}

type Account struct {
	// BidderLocation: Your bidder locations that have distinct URLs.
	BidderLocation []*AccountBidderLocation `json:"bidderLocation,omitempty"`

	// CookieMatchingNid: The nid parameter value used in cookie match
	// requests. Please contact your technical account manager if you need
	// to change this.
	CookieMatchingNid string `json:"cookieMatchingNid,omitempty"`

	// CookieMatchingUrl: The base URL used in cookie match requests.
	CookieMatchingUrl string `json:"cookieMatchingUrl,omitempty"`

	// Id: Account id.
	Id int64 `json:"id,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// MaximumActiveCreatives: The maximum number of active creatives that
	// an account can have, where a creative is active if it was inserted or
	// bid with in the last 30 days. Please contact your technical account
	// manager if you need to change this.
	MaximumActiveCreatives int64 `json:"maximumActiveCreatives,omitempty"`

	// MaximumTotalQps: The sum of all bidderLocation.maximumQps values
	// cannot exceed this. Please contact your technical account manager if
	// you need to change this.
	MaximumTotalQps int64 `json:"maximumTotalQps,omitempty"`

	// NumberActiveCreatives: The number of creatives that this account
	// inserted or bid with in the last 30 days.
	NumberActiveCreatives int64 `json:"numberActiveCreatives,omitempty"`
}

type AccountBidderLocation struct {
	// MaximumQps: The maximum queries per second the Ad Exchange will send.
	MaximumQps int64 `json:"maximumQps,omitempty"`

	// Region: The geographical region the Ad Exchange should send requests
	// from. Only used by some quota systems, but always setting the value
	// is recommended. Allowed values:
	// - ASIA
	// - EUROPE
	// - US_EAST
	// -
	// US_WEST
	Region string `json:"region,omitempty"`

	// Url: The URL to which the Ad Exchange will send bid requests.
	Url string `json:"url,omitempty"`
}

type AccountsList struct {
	// Items: A list of accounts.
	Items []*Account `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type BillingInfo struct {
	// AccountId: Account id.
	AccountId int64 `json:"accountId,omitempty"`

	// AccountName: Account name.
	AccountName string `json:"accountName,omitempty"`

	// BillingId: A list of adgroup IDs associated with this particular
	// account. These IDs may show up as part of a realtime bidding
	// BidRequest, which indicates a bid request for this account.
	BillingId []string `json:"billingId,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type BillingInfoList struct {
	// Items: A list of billing info relevant for your account.
	Items []*BillingInfo `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type Budget struct {
	// AccountId: The id of the account. This is required for get and update
	// requests.
	AccountId int64 `json:"accountId,omitempty,string"`

	// BillingId: The billing id to determine which adgroup to provide
	// budget information for. This is required for get and update requests.
	BillingId int64 `json:"billingId,omitempty,string"`

	// BudgetAmount: The budget amount to apply for the billingId provided.
	// This is required for update requests.
	BudgetAmount int64 `json:"budgetAmount,omitempty,string"`

	// CurrencyCode: The currency code for the buyer. This cannot be altered
	// here.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Id: The unique id that describes this item.
	Id string `json:"id,omitempty"`

	// Kind: The kind of the resource, i.e. "adexchangebuyer#budget".
	Kind string `json:"kind,omitempty"`
}

type Creative struct {
	// HTMLSnippet: The HTML snippet that displays the ad when inserted in
	// the web page. If set, videoURL should not be set.
	HTMLSnippet string `json:"HTMLSnippet,omitempty"`

	// AccountId: Account id.
	AccountId int64 `json:"accountId,omitempty"`

	// AdvertiserId: Detected advertiser id, if any. Read-only. This field
	// should not be set in requests.
	AdvertiserId googleapi.Int64s `json:"advertiserId,omitempty"`

	// AdvertiserName: The name of the company being advertised in the
	// creative.
	AdvertiserName string `json:"advertiserName,omitempty"`

	// AgencyId: The agency id for this creative.
	AgencyId int64 `json:"agencyId,omitempty,string"`

	// Attribute: All attributes for the ads that may be shown from this
	// snippet.
	Attribute []int64 `json:"attribute,omitempty"`

	// BuyerCreativeId: A buyer-specific id identifying the creative in this
	// ad.
	BuyerCreativeId string `json:"buyerCreativeId,omitempty"`

	// ClickThroughUrl: The set of destination urls for the snippet.
	ClickThroughUrl []string `json:"clickThroughUrl,omitempty"`

	// Corrections: Shows any corrections that were applied to this
	// creative. Read-only. This field should not be set in requests.
	Corrections []*CreativeCorrections `json:"corrections,omitempty"`

	// DisapprovalReasons: The reasons for disapproval, if any. Note that
	// not all disapproval reasons may be categorized, so it is possible for
	// the creative to have a status of DISAPPROVED with an empty list for
	// disapproval_reasons. In this case, please reach out to your TAM to
	// help debug the issue. Read-only. This field should not be set in
	// requests.
	DisapprovalReasons []*CreativeDisapprovalReasons `json:"disapprovalReasons,omitempty"`

	// FilteringReasons: The filtering reasons for the creative. Read-only.
	// This field should not be set in requests.
	FilteringReasons *CreativeFilteringReasons `json:"filteringReasons,omitempty"`

	// Height: Ad height.
	Height int64 `json:"height,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// ProductCategories: Detected product categories, if any. Read-only.
	// This field should not be set in requests.
	ProductCategories []int64 `json:"productCategories,omitempty"`

	// RestrictedCategories: All restricted categories for the ads that may
	// be shown from this snippet.
	RestrictedCategories []int64 `json:"restrictedCategories,omitempty"`

	// SensitiveCategories: Detected sensitive categories, if any.
	// Read-only. This field should not be set in requests.
	SensitiveCategories []int64 `json:"sensitiveCategories,omitempty"`

	// Status: Creative serving status. Read-only. This field should not be
	// set in requests.
	Status string `json:"status,omitempty"`

	// VendorType: All vendor types for the ads that may be shown from this
	// snippet.
	VendorType []int64 `json:"vendorType,omitempty"`

	// VideoURL: The url to fetch a video ad. If set, HTMLSnippet should not
	// be set.
	VideoURL string `json:"videoURL,omitempty"`

	// Width: Ad width.
	Width int64 `json:"width,omitempty"`
}

type CreativeCorrections struct {
	// Details: Additional details about the correction.
	Details []string `json:"details,omitempty"`

	// Reason: The type of correction that was applied to the creative.
	Reason string `json:"reason,omitempty"`
}

type CreativeDisapprovalReasons struct {
	// Details: Additional details about the reason for disapproval.
	Details []string `json:"details,omitempty"`

	// Reason: The categorized reason for disapproval.
	Reason string `json:"reason,omitempty"`
}

type CreativeFilteringReasons struct {
	// Date: The date in ISO 8601 format for the data. The data is collected
	// from 00:00:00 to 23:59:59 in PST.
	Date string `json:"date,omitempty"`

	// Reasons: The filtering reasons.
	Reasons []*CreativeFilteringReasonsReasons `json:"reasons,omitempty"`
}

type CreativeFilteringReasonsReasons struct {
	// FilteringCount: The number of times the creative was filtered for the
	// status. The count is aggregated across all publishers on the
	// exchange.
	FilteringCount int64 `json:"filteringCount,omitempty,string"`

	// FilteringStatus: The filtering status code. Please refer to the
	// creative-status-codes.txt file for different statuses.
	FilteringStatus int64 `json:"filteringStatus,omitempty"`
}

type CreativesList struct {
	// Items: A list of creatives.
	Items []*Creative `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through creatives. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type DirectDeal struct {
	// AccountId: The account id of the buyer this deal is for.
	AccountId int64 `json:"accountId,omitempty"`

	// Advertiser: The name of the advertiser this deal is for.
	Advertiser string `json:"advertiser,omitempty"`

	// CurrencyCode: The currency code that applies to the fixed_cpm value.
	// If not set then assumed to be USD.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// EndTime: End time for when this deal stops being active. If not set
	// then this deal is valid until manually disabled by the publisher. In
	// seconds since the epoch.
	EndTime int64 `json:"endTime,omitempty,string"`

	// FixedCpm: The fixed price for this direct deal. In cpm micros of
	// currency according to currency_code. If set, then this deal is
	// eligible for the fixed price tier of buying (highest priority, pay
	// exactly the configured fixed price).
	FixedCpm int64 `json:"fixedCpm,omitempty,string"`

	// Id: Deal id.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// Name: Deal name.
	Name string `json:"name,omitempty"`

	// PrivateExchangeMinCpm: The minimum price for this direct deal. In cpm
	// micros of currency according to currency_code. If set, then this deal
	// is eligible for the private exchange tier of buying (below fixed
	// price priority, run as a second price auction).
	PrivateExchangeMinCpm int64 `json:"privateExchangeMinCpm,omitempty,string"`

	// PublisherBlocksOverriden: If true, the publisher has opted to have
	// their blocks ignored when a creative is bid with for this deal.
	PublisherBlocksOverriden bool `json:"publisherBlocksOverriden,omitempty"`

	// SellerNetwork: The name of the publisher offering this direct deal.
	SellerNetwork string `json:"sellerNetwork,omitempty"`

	// StartTime: Start time for when this deal becomes active. If not set
	// then this deal is active immediately upon creation. In seconds since
	// the epoch.
	StartTime int64 `json:"startTime,omitempty,string"`
}

type DirectDealsList struct {
	// DirectDeals: A list of direct deals relevant for your account.
	DirectDeals []*DirectDeal `json:"directDeals,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type PerformanceReport struct {
	// CalloutStatusRate: Rate of various prefiltering statuses per match.
	// Please refer to the callout-status-codes.txt file for different
	// statuses.
	CalloutStatusRate []interface{} `json:"calloutStatusRate,omitempty"`

	// CookieMatcherStatusRate: Average QPS for cookie matcher operations.
	CookieMatcherStatusRate []interface{} `json:"cookieMatcherStatusRate,omitempty"`

	// CreativeStatusRate: Rate of ads with a given status. Please refer to
	// the creative-status-codes.txt file for different statuses.
	CreativeStatusRate []interface{} `json:"creativeStatusRate,omitempty"`

	// HostedMatchStatusRate: Average QPS for hosted match operations.
	HostedMatchStatusRate []interface{} `json:"hostedMatchStatusRate,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// Latency50thPercentile: The 50th percentile round trip latency(ms) as
	// perceived from Google servers for the duration period covered by the
	// report.
	Latency50thPercentile float64 `json:"latency50thPercentile,omitempty"`

	// Latency85thPercentile: The 85th percentile round trip latency(ms) as
	// perceived from Google servers for the duration period covered by the
	// report.
	Latency85thPercentile float64 `json:"latency85thPercentile,omitempty"`

	// Latency95thPercentile: The 95th percentile round trip latency(ms) as
	// perceived from Google servers for the duration period covered by the
	// report.
	Latency95thPercentile float64 `json:"latency95thPercentile,omitempty"`

	// NoQuotaInRegion: Rate of various quota account statuses per quota
	// check.
	NoQuotaInRegion float64 `json:"noQuotaInRegion,omitempty"`

	// OutOfQuota: Rate of various quota account statuses per quota check.
	OutOfQuota float64 `json:"outOfQuota,omitempty"`

	// PixelMatchRequests: Average QPS for pixel match requests from
	// clients.
	PixelMatchRequests float64 `json:"pixelMatchRequests,omitempty"`

	// PixelMatchResponses: Average QPS for pixel match responses from
	// clients.
	PixelMatchResponses float64 `json:"pixelMatchResponses,omitempty"`

	// QuotaConfiguredLimit: The configured quota limits for this account.
	QuotaConfiguredLimit float64 `json:"quotaConfiguredLimit,omitempty"`

	// QuotaThrottledLimit: The throttled quota limits for this account.
	QuotaThrottledLimit float64 `json:"quotaThrottledLimit,omitempty"`

	// Region: The trading location of this data.
	Region string `json:"region,omitempty"`

	// Timestamp: The unix timestamp of the starting time of this
	// performance data.
	Timestamp int64 `json:"timestamp,omitempty,string"`
}

type PerformanceReportList struct {
	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// PerformanceReport: A list of performance reports relevant for the
	// account.
	PerformanceReport []*PerformanceReport `json:"performanceReport,omitempty"`
}

type PretargetingConfig struct {
	// BillingId: The id for billing purposes, provided for reference. Leave
	// this field blank for insert requests; the id will be generated
	// automatically.
	BillingId int64 `json:"billingId,omitempty,string"`

	// ConfigId: The config id; generated automatically. Leave this field
	// blank for insert requests.
	ConfigId int64 `json:"configId,omitempty,string"`

	// ConfigName: The name of the config. Must be unique. Required for all
	// requests.
	ConfigName string `json:"configName,omitempty"`

	// CreativeType: List must contain exactly one of
	// PRETARGETING_CREATIVE_TYPE_HTML or PRETARGETING_CREATIVE_TYPE_VIDEO.
	CreativeType []string `json:"creativeType,omitempty"`

	// Dimensions: Requests which allow one of these (width, height) pairs
	// will match. All pairs must be supported ad dimensions.
	Dimensions []*PretargetingConfigDimensions `json:"dimensions,omitempty"`

	// ExcludedContentLabels: Requests with any of these content labels will
	// not match. Values are from content-labels.txt in the downloadable
	// files section.
	ExcludedContentLabels googleapi.Int64s `json:"excludedContentLabels,omitempty"`

	// ExcludedGeoCriteriaIds: Requests containing any of these geo criteria
	// ids will not match.
	ExcludedGeoCriteriaIds googleapi.Int64s `json:"excludedGeoCriteriaIds,omitempty"`

	// ExcludedPlacements: Requests containing any of these placements will
	// not match.
	ExcludedPlacements []*PretargetingConfigExcludedPlacements `json:"excludedPlacements,omitempty"`

	// ExcludedUserLists: Requests containing any of these users list ids
	// will not match.
	ExcludedUserLists googleapi.Int64s `json:"excludedUserLists,omitempty"`

	// ExcludedVerticals: Requests containing any of these vertical ids will
	// not match. Values are from the publisher-verticals.txt file in the
	// downloadable files section.
	ExcludedVerticals googleapi.Int64s `json:"excludedVerticals,omitempty"`

	// GeoCriteriaIds: Requests containing any of these geo criteria ids
	// will match.
	GeoCriteriaIds googleapi.Int64s `json:"geoCriteriaIds,omitempty"`

	// IsActive: Whether this config is active. Required for all requests.
	IsActive bool `json:"isActive,omitempty"`

	// Kind: The kind of the resource, i.e.
	// "adexchangebuyer#pretargetingConfig".
	Kind string `json:"kind,omitempty"`

	// Languages: Request containing any of these language codes will match.
	Languages []string `json:"languages,omitempty"`

	// MobileCarriers: Requests containing any of these mobile carrier ids
	// will match. Values are from mobile-carriers.csv in the downloadable
	// files section.
	MobileCarriers googleapi.Int64s `json:"mobileCarriers,omitempty"`

	// MobileDevices: Requests containing any of these mobile device ids
	// will match. Values are from mobile-devices.csv in the downloadable
	// files section.
	MobileDevices googleapi.Int64s `json:"mobileDevices,omitempty"`

	// MobileOperatingSystemVersions: Requests containing any of these
	// mobile operating system version ids will match. Values are from
	// mobile-os.csv in the downloadable files section.
	MobileOperatingSystemVersions googleapi.Int64s `json:"mobileOperatingSystemVersions,omitempty"`

	// Placements: Requests containing any of these placements will match.
	Placements []*PretargetingConfigPlacements `json:"placements,omitempty"`

	// Platforms: Requests matching any of these platforms will match.
	// Possible values are PRETARGETING_PLATFORM_MOBILE,
	// PRETARGETING_PLATFORM_DESKTOP, and PRETARGETING_PLATFORM_TABLET.
	Platforms []string `json:"platforms,omitempty"`

	// SupportedCreativeAttributes: Creative attributes should be declared
	// here if all creatives corresponding to this pretargeting
	// configuration have that creative attribute. Values are from
	// pretargetable-creative-attributes.txt in the downloadable files
	// section.
	SupportedCreativeAttributes googleapi.Int64s `json:"supportedCreativeAttributes,omitempty"`

	// UserLists: Requests containing any of these user list ids will match.
	UserLists googleapi.Int64s `json:"userLists,omitempty"`

	// VendorTypes: Requests that allow any of these vendor ids will match.
	// Values are from vendors.txt in the downloadable files section.
	VendorTypes googleapi.Int64s `json:"vendorTypes,omitempty"`

	// Verticals: Requests containing any of these vertical ids will match.
	Verticals googleapi.Int64s `json:"verticals,omitempty"`
}

type PretargetingConfigDimensions struct {
	// Height: Height in pixels.
	Height int64 `json:"height,omitempty,string"`

	// Width: Width in pixels.
	Width int64 `json:"width,omitempty,string"`
}

type PretargetingConfigExcludedPlacements struct {
	// Token: The value of the placement. Interpretation depends on the
	// placement type, e.g. URL for a site placement, channel name for a
	// channel placement, app id for a mobile app placement.
	Token string `json:"token,omitempty"`

	// Type: The type of the placement.
	Type string `json:"type,omitempty"`
}

type PretargetingConfigPlacements struct {
	// Token: The value of the placement. Interpretation depends on the
	// placement type, e.g. URL for a site placement, channel name for a
	// channel placement, app id for a mobile app placement.
	Token string `json:"token,omitempty"`

	// Type: The type of the placement.
	Type string `json:"type,omitempty"`
}

type PretargetingConfigList struct {
	// Items: A list of pretargeting configs
	Items []*PretargetingConfig `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

// method id "adexchangebuyer.accounts.get":

type AccountsGetCall struct {
	s             *Service
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Gets one account by ID.

func (r *AccountsService) Get(id int64) *AccountsGetCall {
	return &AccountsGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{id}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsGetCall) Context(ctx context.Context) *AccountsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsGetCall) Fields(s ...googleapi.Field) *AccountsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsGetCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets one account by ID.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.accounts.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The account id",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{id}",
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.accounts.list":

type AccountsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the authenticated user's list of accounts.

func (r *AccountsService) List() *AccountsListCall {
	return &AccountsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsListCall) Context(ctx context.Context) *AccountsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsListCall) Fields(s ...googleapi.Field) *AccountsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsListCall) Do() (*AccountsList, error) {
	var returnValue *AccountsList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the authenticated user's list of accounts.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.accounts.list",
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "AccountsList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.accounts.patch":

type AccountsPatchCall struct {
	s             *Service
	id            int64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates an existing account. This method supports patch
// semantics.

func (r *AccountsService) Patch(id int64, account *Account) *AccountsPatchCall {
	return &AccountsPatchCall{
		s:             r.s,
		id:            id,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{id}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsPatchCall) Context(ctx context.Context) *AccountsPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPatchCall) Fields(s ...googleapi.Field) *AccountsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPatchCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.account,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an existing account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "adexchangebuyer.accounts.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The account id",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{id}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.accounts.update":

type AccountsUpdateCall struct {
	s             *Service
	id            int64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates an existing account.

func (r *AccountsService) Update(id int64, account *Account) *AccountsUpdateCall {
	return &AccountsUpdateCall{
		s:             r.s,
		id:            id,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{id}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsUpdateCall) Context(ctx context.Context) *AccountsUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsUpdateCall) Fields(s ...googleapi.Field) *AccountsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsUpdateCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.account,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an existing account.",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer.accounts.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The account id",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{id}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.billingInfo.get":

type BillingInfoGetCall struct {
	s             *Service
	accountId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the billing information for one account specified by
// account ID.

func (r *BillingInfoService) Get(accountId int64) *BillingInfoGetCall {
	return &BillingInfoGetCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "billinginfo/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *BillingInfoGetCall) Context(ctx context.Context) *BillingInfoGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingInfoGetCall) Fields(s ...googleapi.Field) *BillingInfoGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BillingInfoGetCall) Do() (*BillingInfo, error) {
	var returnValue *BillingInfo
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the billing information for one account specified by account ID.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.billingInfo.get",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "billinginfo/{accountId}",
	//   "response": {
	//     "$ref": "BillingInfo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.billingInfo.list":

type BillingInfoListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves a list of billing information for all accounts of the
// authenticated user.

func (r *BillingInfoService) List() *BillingInfoListCall {
	return &BillingInfoListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "billinginfo",
		context_:      googleapi.NoContext,
	}
}
func (c *BillingInfoListCall) Context(ctx context.Context) *BillingInfoListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingInfoListCall) Fields(s ...googleapi.Field) *BillingInfoListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BillingInfoListCall) Do() (*BillingInfoList, error) {
	var returnValue *BillingInfoList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves a list of billing information for all accounts of the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.billingInfo.list",
	//   "path": "billinginfo",
	//   "response": {
	//     "$ref": "BillingInfoList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.budget.get":

type BudgetGetCall struct {
	s             *Service
	accountId     int64
	billingId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the budget information for the adgroup specified by the
// accountId and billingId.

func (r *BudgetService) Get(accountId int64, billingId int64) *BudgetGetCall {
	return &BudgetGetCall{
		s:             r.s,
		accountId:     accountId,
		billingId:     billingId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "billinginfo/{accountId}/{billingId}",
		context_:      googleapi.NoContext,
	}
}
func (c *BudgetGetCall) Context(ctx context.Context) *BudgetGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BudgetGetCall) Fields(s ...googleapi.Field) *BudgetGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BudgetGetCall) Do() (*Budget, error) {
	var returnValue *Budget
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"billingId": strconv.FormatInt(c.billingId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the budget information for the adgroup specified by the accountId and billingId.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.budget.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "billingId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to get the budget information for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "billingId": {
	//       "description": "The billing id to get the budget information for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "billinginfo/{accountId}/{billingId}",
	//   "response": {
	//     "$ref": "Budget"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.budget.patch":

type BudgetPatchCall struct {
	s             *Service
	accountId     int64
	billingId     int64
	budget        *Budget
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates the budget amount for the budget of the adgroup
// specified by the accountId and billingId, with the budget amount in
// the request. This method supports patch semantics.

func (r *BudgetService) Patch(accountId int64, billingId int64, budget *Budget) *BudgetPatchCall {
	return &BudgetPatchCall{
		s:             r.s,
		accountId:     accountId,
		billingId:     billingId,
		budget:        budget,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "billinginfo/{accountId}/{billingId}",
		context_:      googleapi.NoContext,
	}
}
func (c *BudgetPatchCall) Context(ctx context.Context) *BudgetPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BudgetPatchCall) Fields(s ...googleapi.Field) *BudgetPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BudgetPatchCall) Do() (*Budget, error) {
	var returnValue *Budget
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"billingId": strconv.FormatInt(c.billingId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.budget,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the budget amount for the budget of the adgroup specified by the accountId and billingId, with the budget amount in the request. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "adexchangebuyer.budget.patch",
	//   "parameterOrder": [
	//     "accountId",
	//     "billingId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id associated with the budget being updated.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "billingId": {
	//       "description": "The billing id associated with the budget being updated.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "billinginfo/{accountId}/{billingId}",
	//   "request": {
	//     "$ref": "Budget"
	//   },
	//   "response": {
	//     "$ref": "Budget"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.budget.update":

type BudgetUpdateCall struct {
	s             *Service
	accountId     int64
	billingId     int64
	budget        *Budget
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates the budget amount for the budget of the adgroup
// specified by the accountId and billingId, with the budget amount in
// the request.

func (r *BudgetService) Update(accountId int64, billingId int64, budget *Budget) *BudgetUpdateCall {
	return &BudgetUpdateCall{
		s:             r.s,
		accountId:     accountId,
		billingId:     billingId,
		budget:        budget,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "billinginfo/{accountId}/{billingId}",
		context_:      googleapi.NoContext,
	}
}
func (c *BudgetUpdateCall) Context(ctx context.Context) *BudgetUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BudgetUpdateCall) Fields(s ...googleapi.Field) *BudgetUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BudgetUpdateCall) Do() (*Budget, error) {
	var returnValue *Budget
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"billingId": strconv.FormatInt(c.billingId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.budget,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the budget amount for the budget of the adgroup specified by the accountId and billingId, with the budget amount in the request.",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer.budget.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "billingId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id associated with the budget being updated.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "billingId": {
	//       "description": "The billing id associated with the budget being updated.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "billinginfo/{accountId}/{billingId}",
	//   "request": {
	//     "$ref": "Budget"
	//   },
	//   "response": {
	//     "$ref": "Budget"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.creatives.get":

type CreativesGetCall struct {
	s               *Service
	accountId       int64
	buyerCreativeId string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Get: Gets the status for a single creative. A creative will be
// available 30-40 minutes after submission.

func (r *CreativesService) Get(accountId int64, buyerCreativeId string) *CreativesGetCall {
	return &CreativesGetCall{
		s:               r.s,
		accountId:       accountId,
		buyerCreativeId: buyerCreativeId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "creatives/{accountId}/{buyerCreativeId}",
		context_:        googleapi.NoContext,
	}
}
func (c *CreativesGetCall) Context(ctx context.Context) *CreativesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesGetCall) Fields(s ...googleapi.Field) *CreativesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesGetCall) Do() (*Creative, error) {
	var returnValue *Creative
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"buyerCreativeId": c.buyerCreativeId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the status for a single creative. A creative will be available 30-40 minutes after submission.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.creatives.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "buyerCreativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The id for the account that will serve this creative.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "buyerCreativeId": {
	//       "description": "The buyer-specific id for this creative.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "creatives/{accountId}/{buyerCreativeId}",
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.creatives.insert":

type CreativesInsertCall struct {
	s             *Service
	creative      *Creative
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Submit a new creative.

func (r *CreativesService) Insert(creative *Creative) *CreativesInsertCall {
	return &CreativesInsertCall{
		s:             r.s,
		creative:      creative,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "creatives",
		context_:      googleapi.NoContext,
	}
}
func (c *CreativesInsertCall) Context(ctx context.Context) *CreativesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesInsertCall) Fields(s ...googleapi.Field) *CreativesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesInsertCall) Do() (*Creative, error) {
	var returnValue *Creative
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.creative,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Submit a new creative.",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer.creatives.insert",
	//   "path": "creatives",
	//   "request": {
	//     "$ref": "Creative"
	//   },
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.creatives.list":

type CreativesListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves a list of the authenticated user's active creatives.
// A creative will be available 30-40 minutes after submission.

func (r *CreativesService) List() *CreativesListCall {
	return &CreativesListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "creatives",
		context_:      googleapi.NoContext,
	}
}

// AccountId sets the optional parameter "accountId": When specified,
// only creatives for the given account ids are returned.
func (c *CreativesListCall) AccountId(accountId ...int64) *CreativesListCall {
	for _, v_ := range accountId {
		c.params_.Add("accountId", fmt.Sprintf("%v", v_))
	}
	return c
}

// BuyerCreativeId sets the optional parameter "buyerCreativeId": When
// specified, only creatives for the given buyer creative ids are
// returned.
func (c *CreativesListCall) BuyerCreativeId(buyerCreativeId ...string) *CreativesListCall {
	c.params_["buyerCreativeId"] = buyerCreativeId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. If not set, the default is
// 100.
func (c *CreativesListCall) MaxResults(maxResults int64) *CreativesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad clients. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *CreativesListCall) PageToken(pageToken string) *CreativesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// StatusFilter sets the optional parameter "statusFilter": When
// specified, only creatives having the given status are returned.
func (c *CreativesListCall) StatusFilter(statusFilter string) *CreativesListCall {
	c.params_.Set("statusFilter", fmt.Sprintf("%v", statusFilter))
	return c
}
func (c *CreativesListCall) Context(ctx context.Context) *CreativesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesListCall) Fields(s ...googleapi.Field) *CreativesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesListCall) Do() (*CreativesList, error) {
	var returnValue *CreativesList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves a list of the authenticated user's active creatives. A creative will be available 30-40 minutes after submission.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.creatives.list",
	//   "parameters": {
	//     "accountId": {
	//       "description": "When specified, only creatives for the given account ids are returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "integer"
	//     },
	//     "buyerCreativeId": {
	//       "description": "When specified, only creatives for the given buyer creative ids are returned.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. If not set, the default is 100. Optional.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad clients. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "statusFilter": {
	//       "description": "When specified, only creatives having the given status are returned.",
	//       "enum": [
	//         "approved",
	//         "disapproved",
	//         "not_checked"
	//       ],
	//       "enumDescriptions": [
	//         "Creatives which have been approved.",
	//         "Creatives which have been disapproved.",
	//         "Creatives whose status is not yet checked."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "creatives",
	//   "response": {
	//     "$ref": "CreativesList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.directDeals.get":

type DirectDealsGetCall struct {
	s             *Service
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Gets one direct deal by ID.

func (r *DirectDealsService) Get(id int64) *DirectDealsGetCall {
	return &DirectDealsGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "directdeals/{id}",
		context_:      googleapi.NoContext,
	}
}
func (c *DirectDealsGetCall) Context(ctx context.Context) *DirectDealsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DirectDealsGetCall) Fields(s ...googleapi.Field) *DirectDealsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DirectDealsGetCall) Do() (*DirectDeal, error) {
	var returnValue *DirectDeal
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets one direct deal by ID.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.directDeals.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The direct deal id",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "directdeals/{id}",
	//   "response": {
	//     "$ref": "DirectDeal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.directDeals.list":

type DirectDealsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the authenticated user's list of direct deals.

func (r *DirectDealsService) List() *DirectDealsListCall {
	return &DirectDealsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "directdeals",
		context_:      googleapi.NoContext,
	}
}
func (c *DirectDealsListCall) Context(ctx context.Context) *DirectDealsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DirectDealsListCall) Fields(s ...googleapi.Field) *DirectDealsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DirectDealsListCall) Do() (*DirectDealsList, error) {
	var returnValue *DirectDealsList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the authenticated user's list of direct deals.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.directDeals.list",
	//   "path": "directdeals",
	//   "response": {
	//     "$ref": "DirectDealsList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.performanceReport.list":

type PerformanceReportListCall struct {
	s             *Service
	accountId     int64
	endDateTime   string
	startDateTime string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the authenticated user's list of performance metrics.

func (r *PerformanceReportService) List(accountId int64, endDateTime string, startDateTime string) *PerformanceReportListCall {
	return &PerformanceReportListCall{
		s:             r.s,
		accountId:     accountId,
		endDateTime:   endDateTime,
		startDateTime: startDateTime,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "performancereport",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. If not set, the default is
// 100.
func (c *PerformanceReportListCall) MaxResults(maxResults int64) *PerformanceReportListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through performance reports. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *PerformanceReportListCall) PageToken(pageToken string) *PerformanceReportListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *PerformanceReportListCall) Context(ctx context.Context) *PerformanceReportListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PerformanceReportListCall) Fields(s ...googleapi.Field) *PerformanceReportListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PerformanceReportListCall) Do() (*PerformanceReportList, error) {
	var returnValue *PerformanceReportList
	c.params_.Set("accountId", fmt.Sprintf("%v", c.accountId))
	c.params_.Set("endDateTime", fmt.Sprintf("%v", c.endDateTime))
	c.params_.Set("startDateTime", fmt.Sprintf("%v", c.startDateTime))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the authenticated user's list of performance metrics.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.performanceReport.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "endDateTime",
	//     "startDateTime"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to get the reports.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDateTime": {
	//       "description": "The end time of the report in ISO 8601 timestamp format using UTC.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. If not set, the default is 100. Optional.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through performance reports. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startDateTime": {
	//       "description": "The start time of the report in ISO 8601 timestamp format using UTC.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "performancereport",
	//   "response": {
	//     "$ref": "PerformanceReportList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.pretargetingConfig.delete":

type PretargetingConfigDeleteCall struct {
	s             *Service
	accountId     int64
	configId      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes an existing pretargeting config.

func (r *PretargetingConfigService) Delete(accountId int64, configId int64) *PretargetingConfigDeleteCall {
	return &PretargetingConfigDeleteCall{
		s:             r.s,
		accountId:     accountId,
		configId:      configId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "pretargetingconfigs/{accountId}/{configId}",
		context_:      googleapi.NoContext,
	}
}
func (c *PretargetingConfigDeleteCall) Context(ctx context.Context) *PretargetingConfigDeleteCall {
	c.context_ = ctx
	return c
}

func (c *PretargetingConfigDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"configId":  strconv.FormatInt(c.configId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes an existing pretargeting config.",
	//   "httpMethod": "DELETE",
	//   "id": "adexchangebuyer.pretargetingConfig.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "configId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to delete the pretargeting config for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "configId": {
	//       "description": "The specific id of the configuration to delete.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "pretargetingconfigs/{accountId}/{configId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.pretargetingConfig.get":

type PretargetingConfigGetCall struct {
	s             *Service
	accountId     int64
	configId      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Gets a specific pretargeting configuration

func (r *PretargetingConfigService) Get(accountId int64, configId int64) *PretargetingConfigGetCall {
	return &PretargetingConfigGetCall{
		s:             r.s,
		accountId:     accountId,
		configId:      configId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "pretargetingconfigs/{accountId}/{configId}",
		context_:      googleapi.NoContext,
	}
}
func (c *PretargetingConfigGetCall) Context(ctx context.Context) *PretargetingConfigGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PretargetingConfigGetCall) Fields(s ...googleapi.Field) *PretargetingConfigGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PretargetingConfigGetCall) Do() (*PretargetingConfig, error) {
	var returnValue *PretargetingConfig
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"configId":  strconv.FormatInt(c.configId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets a specific pretargeting configuration",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.pretargetingConfig.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "configId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to get the pretargeting config for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "configId": {
	//       "description": "The specific id of the configuration to retrieve.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "pretargetingconfigs/{accountId}/{configId}",
	//   "response": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.pretargetingConfig.insert":

type PretargetingConfigInsertCall struct {
	s                  *Service
	accountId          int64
	pretargetingconfig *PretargetingConfig
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
	context_           context.Context
}

// Insert: Inserts a new pretargeting configuration.

func (r *PretargetingConfigService) Insert(accountId int64, pretargetingconfig *PretargetingConfig) *PretargetingConfigInsertCall {
	return &PretargetingConfigInsertCall{
		s:                  r.s,
		accountId:          accountId,
		pretargetingconfig: pretargetingconfig,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "pretargetingconfigs/{accountId}",
		context_:           googleapi.NoContext,
	}
}
func (c *PretargetingConfigInsertCall) Context(ctx context.Context) *PretargetingConfigInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PretargetingConfigInsertCall) Fields(s ...googleapi.Field) *PretargetingConfigInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PretargetingConfigInsertCall) Do() (*PretargetingConfig, error) {
	var returnValue *PretargetingConfig
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.pretargetingconfig,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Inserts a new pretargeting configuration.",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer.pretargetingConfig.insert",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to insert the pretargeting config for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "pretargetingconfigs/{accountId}",
	//   "request": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "response": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.pretargetingConfig.list":

type PretargetingConfigListCall struct {
	s             *Service
	accountId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves a list of the authenticated user's pretargeting
// configurations.

func (r *PretargetingConfigService) List(accountId int64) *PretargetingConfigListCall {
	return &PretargetingConfigListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "pretargetingconfigs/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *PretargetingConfigListCall) Context(ctx context.Context) *PretargetingConfigListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PretargetingConfigListCall) Fields(s ...googleapi.Field) *PretargetingConfigListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PretargetingConfigListCall) Do() (*PretargetingConfigList, error) {
	var returnValue *PretargetingConfigList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves a list of the authenticated user's pretargeting configurations.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.pretargetingConfig.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to get the pretargeting configs for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "pretargetingconfigs/{accountId}",
	//   "response": {
	//     "$ref": "PretargetingConfigList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.pretargetingConfig.patch":

type PretargetingConfigPatchCall struct {
	s                  *Service
	accountId          int64
	configId           int64
	pretargetingconfig *PretargetingConfig
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
	context_           context.Context
}

// Patch: Updates an existing pretargeting config. This method supports
// patch semantics.

func (r *PretargetingConfigService) Patch(accountId int64, configId int64, pretargetingconfig *PretargetingConfig) *PretargetingConfigPatchCall {
	return &PretargetingConfigPatchCall{
		s:                  r.s,
		accountId:          accountId,
		configId:           configId,
		pretargetingconfig: pretargetingconfig,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "pretargetingconfigs/{accountId}/{configId}",
		context_:           googleapi.NoContext,
	}
}
func (c *PretargetingConfigPatchCall) Context(ctx context.Context) *PretargetingConfigPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PretargetingConfigPatchCall) Fields(s ...googleapi.Field) *PretargetingConfigPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PretargetingConfigPatchCall) Do() (*PretargetingConfig, error) {
	var returnValue *PretargetingConfig
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"configId":  strconv.FormatInt(c.configId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.pretargetingconfig,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an existing pretargeting config. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "adexchangebuyer.pretargetingConfig.patch",
	//   "parameterOrder": [
	//     "accountId",
	//     "configId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to update the pretargeting config for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "configId": {
	//       "description": "The specific id of the configuration to update.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "pretargetingconfigs/{accountId}/{configId}",
	//   "request": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "response": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer.pretargetingConfig.update":

type PretargetingConfigUpdateCall struct {
	s                  *Service
	accountId          int64
	configId           int64
	pretargetingconfig *PretargetingConfig
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
	context_           context.Context
}

// Update: Updates an existing pretargeting config.

func (r *PretargetingConfigService) Update(accountId int64, configId int64, pretargetingconfig *PretargetingConfig) *PretargetingConfigUpdateCall {
	return &PretargetingConfigUpdateCall{
		s:                  r.s,
		accountId:          accountId,
		configId:           configId,
		pretargetingconfig: pretargetingconfig,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "pretargetingconfigs/{accountId}/{configId}",
		context_:           googleapi.NoContext,
	}
}
func (c *PretargetingConfigUpdateCall) Context(ctx context.Context) *PretargetingConfigUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PretargetingConfigUpdateCall) Fields(s ...googleapi.Field) *PretargetingConfigUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PretargetingConfigUpdateCall) Do() (*PretargetingConfig, error) {
	var returnValue *PretargetingConfig
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
		"configId":  strconv.FormatInt(c.configId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.pretargetingconfig,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an existing pretargeting config.",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer.pretargetingConfig.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "configId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account id to update the pretargeting config for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "configId": {
	//       "description": "The specific id of the configuration to update.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "pretargetingconfigs/{accountId}/{configId}",
	//   "request": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "response": {
	//     "$ref": "PretargetingConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}
