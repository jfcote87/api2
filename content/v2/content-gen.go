// Package content provides access to the Content API for Shopping.
//
// See https://developers.google.com/shopping-content/v2/
//
// Usage example:
//
//   import "github.com/jfcote87/api2/content/v2"
//   ...
//   contentService, err := content.New(oauthHttpClient)
package content

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

const apiId = "content:v2"
const apiName = "content"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/content/v2/"}

// OAuth2 scopes used by this API.
const (
	// Manage your product listings and accounts for Google Shopping
	ContentScope = "https://www.googleapis.com/auth/content"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Accounts = NewAccountsService(s)
	s.Accountshipping = NewAccountshippingService(s)
	s.Accountstatuses = NewAccountstatusesService(s)
	s.Accounttax = NewAccounttaxService(s)
	s.Datafeeds = NewDatafeedsService(s)
	s.Datafeedstatuses = NewDatafeedstatusesService(s)
	s.Inventory = NewInventoryService(s)
	s.Products = NewProductsService(s)
	s.Productstatuses = NewProductstatusesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Accounts *AccountsService

	Accountshipping *AccountshippingService

	Accountstatuses *AccountstatusesService

	Accounttax *AccounttaxService

	Datafeeds *DatafeedsService

	Datafeedstatuses *DatafeedstatusesService

	Inventory *InventoryService

	Products *ProductsService

	Productstatuses *ProductstatusesService
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	return rs
}

type AccountsService struct {
	s *Service
}

func NewAccountshippingService(s *Service) *AccountshippingService {
	rs := &AccountshippingService{s: s}
	return rs
}

type AccountshippingService struct {
	s *Service
}

func NewAccountstatusesService(s *Service) *AccountstatusesService {
	rs := &AccountstatusesService{s: s}
	return rs
}

type AccountstatusesService struct {
	s *Service
}

func NewAccounttaxService(s *Service) *AccounttaxService {
	rs := &AccounttaxService{s: s}
	return rs
}

type AccounttaxService struct {
	s *Service
}

func NewDatafeedsService(s *Service) *DatafeedsService {
	rs := &DatafeedsService{s: s}
	return rs
}

type DatafeedsService struct {
	s *Service
}

func NewDatafeedstatusesService(s *Service) *DatafeedstatusesService {
	rs := &DatafeedstatusesService{s: s}
	return rs
}

type DatafeedstatusesService struct {
	s *Service
}

func NewInventoryService(s *Service) *InventoryService {
	rs := &InventoryService{s: s}
	return rs
}

type InventoryService struct {
	s *Service
}

func NewProductsService(s *Service) *ProductsService {
	rs := &ProductsService{s: s}
	return rs
}

type ProductsService struct {
	s *Service
}

func NewProductstatusesService(s *Service) *ProductstatusesService {
	rs := &ProductstatusesService{s: s}
	return rs
}

type ProductstatusesService struct {
	s *Service
}

type Account struct {
	// AdultContent: Indicates whether the merchant sells adult content.
	AdultContent bool `json:"adultContent,omitempty"`

	// AdwordsLinks: List of linked AdWords accounts.
	AdwordsLinks []*AccountAdwordsLink `json:"adwordsLinks,omitempty"`

	// Id: Merchant Center account ID.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#account".
	Kind string `json:"kind,omitempty"`

	// Name: Display name for the account.
	Name string `json:"name,omitempty"`

	// ReviewsUrl: URL for individual seller reviews, i.e., reviews for each
	// child account.
	ReviewsUrl string `json:"reviewsUrl,omitempty"`

	// SellerId: Client-specific, locally-unique, internal ID for the child
	// account.
	SellerId string `json:"sellerId,omitempty"`

	// Users: Users with access to the account. Every account (except for
	// subaccounts) must have at least one admin user.
	Users []*AccountUser `json:"users,omitempty"`

	// WebsiteUrl: The merchant's website.
	WebsiteUrl string `json:"websiteUrl,omitempty"`
}

type AccountAdwordsLink struct {
	// AdwordsId: Customer ID of the AdWords account.
	AdwordsId uint64 `json:"adwordsId,omitempty,string"`

	// Status: Status of the link between this Merchant Center account and
	// the AdWords account.
	Status string `json:"status,omitempty"`
}

type AccountIdentifier struct {
	// AggregatorId: The aggregator ID, set for aggregators and subaccounts
	// (in that case, it represents the aggregator of the subaccount).
	AggregatorId uint64 `json:"aggregatorId,omitempty,string"`

	// MerchantId: The merchant account ID, set for individual accounts and
	// subaccounts.
	MerchantId uint64 `json:"merchantId,omitempty,string"`
}

type AccountShipping struct {
	// AccountId: The ID of the account to which these account shipping
	// settings belong.
	AccountId uint64 `json:"accountId,omitempty,string"`

	// CarrierRates: Carrier-based shipping calculations.
	CarrierRates []*AccountShippingCarrierRate `json:"carrierRates,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountShipping".
	Kind string `json:"kind,omitempty"`

	// LocationGroups: Location groups for shipping.
	LocationGroups []*AccountShippingLocationGroup `json:"locationGroups,omitempty"`

	// RateTables: Rate tables definitions.
	RateTables []*AccountShippingRateTable `json:"rateTables,omitempty"`

	// Services: Shipping services describing shipping fees calculation.
	Services []*AccountShippingShippingService `json:"services,omitempty"`
}

type AccountShippingCarrierRate struct {
	// Carrier: The carrier that is responsible for the shipping, such as
	// "UPS", "FedEx", or "USPS".
	Carrier string `json:"carrier,omitempty"`

	// CarrierService: The carrier service, such as "Ground" or "2Day".
	CarrierService string `json:"carrierService,omitempty"`

	// ModifierFlatRate: Additive shipping rate modifier.
	ModifierFlatRate *Price `json:"modifierFlatRate,omitempty"`

	// ModifierPercent: Multiplicative shipping rate modifier in percent.
	// Represented as a floating point number without the percentage
	// character.
	ModifierPercent string `json:"modifierPercent,omitempty"`

	// Name: The name of the carrier rate.
	Name string `json:"name,omitempty"`

	// SaleCountry: Sale country for which this carrier rate is valid,
	// represented as an ISO 3166-1 Alpha-2 code.
	SaleCountry string `json:"saleCountry,omitempty"`

	// ShippingOrigin: Shipping origin represented as a postal code.
	ShippingOrigin string `json:"shippingOrigin,omitempty"`
}

type AccountShippingCondition struct {
	// DeliveryLocationGroup: Delivery location in terms of a location group
	// name. A location group with this name must be specified among
	// location groups.
	DeliveryLocationGroup string `json:"deliveryLocationGroup,omitempty"`

	// DeliveryLocationId: Delivery location in terms of a location ID. Can
	// be used to represent administrative areas, smaller country
	// subdivisions, or cities.
	DeliveryLocationId int64 `json:"deliveryLocationId,omitempty,string"`

	// DeliveryPostalCode: Delivery location in terms of a postal code.
	DeliveryPostalCode string `json:"deliveryPostalCode,omitempty"`

	// DeliveryPostalCodeRange: Delivery location in terms of a postal code
	// range.
	DeliveryPostalCodeRange *AccountShippingPostalCodeRange `json:"deliveryPostalCodeRange,omitempty"`

	// PriceMax: Maximum shipping price. Forms an interval between the
	// maximum of smaller prices (exclusive) and this price (inclusive).
	PriceMax *Price `json:"priceMax,omitempty"`

	// ShippingLabel: Shipping label of the product. The products with the
	// label are matched.
	ShippingLabel string `json:"shippingLabel,omitempty"`

	// WeightMax: Maximum shipping weight. Forms an interval between the
	// maximum of smaller weight (exclusive) and this weight (inclusive).
	WeightMax *Weight `json:"weightMax,omitempty"`
}

type AccountShippingLocationGroup struct {
	// Country: The country in which this location group is, represented as
	// ISO 3166-1 Alpha-2 code.
	Country string `json:"country,omitempty"`

	// LocationIds: A location ID (also called criteria ID) representing
	// administrative areas, smaller country subdivisions (counties), or
	// cities.
	LocationIds googleapi.Int64s `json:"locationIds,omitempty"`

	// Name: The name of the location group.
	Name string `json:"name,omitempty"`

	// PostalCodeRanges: A postal code range representing a city or a set of
	// cities.
	PostalCodeRanges []*AccountShippingPostalCodeRange `json:"postalCodeRanges,omitempty"`

	// PostalCodes: A postal code representing a city or a set of cities.
	//
	// - A single postal code (e.g., 12345)
	// - A postal code prefix followed
	// by a star (e.g., 1234*)
	PostalCodes []string `json:"postalCodes,omitempty"`
}

type AccountShippingPostalCodeRange struct {
	// End: The last (inclusive) postal code or prefix of the range.
	End string `json:"end,omitempty"`

	// Start: The first (inclusive) postal code or prefix of the range.
	Start string `json:"start,omitempty"`
}

type AccountShippingRateTable struct {
	// Content: One-dimensional table cells define one condition along the
	// same dimension. Bi-dimensional table cells use two dimensions with
	// respectively M and N distinct values and must contain exactly M * N
	// cells with distinct conditions (for each possible value pairs).
	Content []*AccountShippingRateTableCell `json:"content,omitempty"`

	// Name: The name of the rate table.
	Name string `json:"name,omitempty"`

	// SaleCountry: Sale country for which this table is valid, represented
	// as an ISO 3166-1 Alpha-2 code.
	SaleCountry string `json:"saleCountry,omitempty"`
}

type AccountShippingRateTableCell struct {
	// Condition: Conditions for which the cell is valid. All cells in a
	// table must use the same dimension or pair of dimensions among price,
	// weight, shipping label or delivery location. If no condition is
	// specified, the cell acts as a catch-all and matches all the elements
	// that are not matched by other cells in this dimension.
	Condition *AccountShippingCondition `json:"condition,omitempty"`

	// Rate: The rate applicable if the cell conditions are matched.
	Rate *Price `json:"rate,omitempty"`
}

type AccountShippingShippingService struct {
	// Active: Whether the shipping service is available.
	Active bool `json:"active,omitempty"`

	// CalculationMethod: Calculation method for the "simple" case that
	// needs no rules.
	CalculationMethod *AccountShippingShippingServiceCalculationMethod `json:"calculationMethod,omitempty"`

	// CostRuleTree: Decision tree for "complicated" shipping cost
	// calculation.
	CostRuleTree *AccountShippingShippingServiceCostRule `json:"costRuleTree,omitempty"`

	// Name: The name of this shipping service.
	Name string `json:"name,omitempty"`

	// SaleCountry: Sale country for which this service can be used,
	// represented as an ISO 3166-1 Alpha-2 code.
	SaleCountry string `json:"saleCountry,omitempty"`
}

type AccountShippingShippingServiceCalculationMethod struct {
	// CarrierRate: Name of the carrier rate to use for the calculation.
	CarrierRate string `json:"carrierRate,omitempty"`

	// Excluded: Delivery is excluded. Valid only within cost rules tree.
	Excluded bool `json:"excluded,omitempty"`

	// FlatRate: Fixed price shipping, represented as a floating point
	// number associated with a currency.
	FlatRate *Price `json:"flatRate,omitempty"`

	// PercentageRate: Percentage of the price, represented as a floating
	// point number without the percentage character.
	PercentageRate string `json:"percentageRate,omitempty"`

	// RateTable: Name of the rate table to use for the calculation.
	RateTable string `json:"rateTable,omitempty"`
}

type AccountShippingShippingServiceCostRule struct {
	// CalculationMethod: Final calculation method to be used only in leaf
	// nodes.
	CalculationMethod *AccountShippingShippingServiceCalculationMethod `json:"calculationMethod,omitempty"`

	// Children: Subsequent rules to be applied, only for inner nodes. The
	// last child must not specify a condition and acts as a catch-all.
	Children []*AccountShippingShippingServiceCostRule `json:"children,omitempty"`

	// Condition: Condition for this rule to be applicable. If no condition
	// is specified, the rule acts as a catch-all.
	Condition *AccountShippingCondition `json:"condition,omitempty"`
}

type AccountStatus struct {
	// AccountId: The ID of the account for which the status is reported.
	AccountId string `json:"accountId,omitempty"`

	// DataQualityIssues: A list of data quality issues.
	DataQualityIssues []*AccountStatusDataQualityIssue `json:"dataQualityIssues,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountStatus".
	Kind string `json:"kind,omitempty"`
}

type AccountStatusDataQualityIssue struct {
	// Country: Country for which this issue is reported.
	Country string `json:"country,omitempty"`

	// DisplayedValue: Actual value displayed on the landing page.
	DisplayedValue string `json:"displayedValue,omitempty"`

	// ExampleItems: Example items featuring the issue.
	ExampleItems []*AccountStatusExampleItem `json:"exampleItems,omitempty"`

	// Id: Issue identifier.
	Id string `json:"id,omitempty"`

	// LastChecked: Last time the account was checked for this issue.
	LastChecked string `json:"lastChecked,omitempty"`

	// NumItems: Number of items in the account found to have the said
	// issue.
	NumItems int64 `json:"numItems,omitempty"`

	// Severity: Severity of the problem.
	Severity string `json:"severity,omitempty"`

	// SubmittedValue: Submitted value that causes the issue.
	SubmittedValue string `json:"submittedValue,omitempty"`
}

type AccountStatusExampleItem struct {
	// ItemId: Unique item ID as specified in the uploaded product data.
	ItemId string `json:"itemId,omitempty"`

	// Link: Landing page of the item.
	Link string `json:"link,omitempty"`

	// SubmittedValue: The item value that was submitted.
	SubmittedValue string `json:"submittedValue,omitempty"`

	// Title: Title of the item.
	Title string `json:"title,omitempty"`

	// ValueOnLandingPage: The actual value on the landing page.
	ValueOnLandingPage string `json:"valueOnLandingPage,omitempty"`
}

type AccountTax struct {
	// AccountId: The ID of the account to which these account tax settings
	// belong.
	AccountId uint64 `json:"accountId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountTax".
	Kind string `json:"kind,omitempty"`

	// Rules: Tax rules. Updating the tax rules will enable US taxes (not
	// reversible). Defining no rules is equivalent to not charging tax at
	// all.
	Rules []*AccountTaxTaxRule `json:"rules,omitempty"`
}

type AccountTaxTaxRule struct {
	// Country: Country code in which tax is applicable.
	Country string `json:"country,omitempty"`

	// LocationId: State (or province) is which the tax is applicable,
	// described by its location id (also called criteria id).
	LocationId uint64 `json:"locationId,omitempty,string"`

	// RatePercent: Explicit tax rate in percent, represented as a floating
	// point number without the percentage character. Must not be negative.
	RatePercent string `json:"ratePercent,omitempty"`

	// ShippingTaxed: If true, shipping charges are also taxed.
	ShippingTaxed bool `json:"shippingTaxed,omitempty"`

	// UseGlobalRate: Whether the tax rate is taken from a global tax table
	// or specified explicitly.
	UseGlobalRate bool `json:"useGlobalRate,omitempty"`
}

type AccountUser struct {
	// Admin: Whether user is an admin.
	Admin bool `json:"admin,omitempty"`

	// EmailAddress: User's email address.
	EmailAddress string `json:"emailAddress,omitempty"`
}

type AccountsAuthInfoResponse struct {
	// AccountIdentifiers: The account identifiers corresponding to the
	// authenticated user.
	// - For an individual account: only the merchant ID
	// is defined
	// - For an aggregator: only the aggregator ID is defined
	// -
	// For a subaccount of an MCA: both the merchant ID and the aggregator
	// ID are defined.
	AccountIdentifiers []*AccountIdentifier `json:"accountIdentifiers,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountsAuthInfoResponse".
	Kind string `json:"kind,omitempty"`
}

type AccountsCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*AccountsCustomBatchRequestEntry `json:"entries,omitempty"`
}

type AccountsCustomBatchRequestEntry struct {
	// Account: The account to create or update. Only defined if the method
	// is insert or update.
	Account *Account `json:"account,omitempty"`

	// AccountId: The ID of the account to get or delete. Only defined if
	// the method is get or delete.
	AccountId uint64 `json:"accountId,omitempty,string"`

	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`
}

type AccountsCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*AccountsCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountsCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type AccountsCustomBatchResponseEntry struct {
	// Account: The retrieved, created, or updated account. Not defined if
	// the method was delete.
	Account *Account `json:"account,omitempty"`

	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountsCustomBatchResponseEntry".
	Kind string `json:"kind,omitempty"`
}

type AccountsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// accounts.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*Account `json:"resources,omitempty"`
}

type AccountshippingCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*AccountshippingCustomBatchRequestEntry `json:"entries,omitempty"`
}

type AccountshippingCustomBatchRequestEntry struct {
	// AccountId: The ID of the account for which to get/update account
	// shipping settings.
	AccountId uint64 `json:"accountId,omitempty,string"`

	// AccountShipping: The account shipping settings to update. Only
	// defined if the method is update.
	AccountShipping *AccountShipping `json:"accountShipping,omitempty"`

	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`
}

type AccountshippingCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*AccountshippingCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountshippingCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type AccountshippingCustomBatchResponseEntry struct {
	// AccountShipping: The retrieved or updated account shipping settings.
	AccountShipping *AccountShipping `json:"accountShipping,omitempty"`

	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountshippingCustomBatchResponseEntry".
	Kind string `json:"kind,omitempty"`
}

type AccountshippingListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountshippingListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// account shipping settings.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*AccountShipping `json:"resources,omitempty"`
}

type AccountstatusesCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*AccountstatusesCustomBatchRequestEntry `json:"entries,omitempty"`
}

type AccountstatusesCustomBatchRequestEntry struct {
	// AccountId: The ID of the (sub-)account whose status to get.
	AccountId uint64 `json:"accountId,omitempty,string"`

	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	// Method: The method (get).
	Method string `json:"method,omitempty"`
}

type AccountstatusesCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*AccountstatusesCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountstatusesCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type AccountstatusesCustomBatchResponseEntry struct {
	// AccountStatus: The requested account status. Defined if and only if
	// the request was successful.
	AccountStatus *AccountStatus `json:"accountStatus,omitempty"`

	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`
}

type AccountstatusesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accountstatusesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// account statuses.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*AccountStatus `json:"resources,omitempty"`
}

type AccounttaxCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*AccounttaxCustomBatchRequestEntry `json:"entries,omitempty"`
}

type AccounttaxCustomBatchRequestEntry struct {
	// AccountId: The ID of the account for which to get/update account tax
	// settings.
	AccountId uint64 `json:"accountId,omitempty,string"`

	// AccountTax: The account tax settings to update. Only defined if the
	// method is update.
	AccountTax *AccountTax `json:"accountTax,omitempty"`

	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`
}

type AccounttaxCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*AccounttaxCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accounttaxCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type AccounttaxCustomBatchResponseEntry struct {
	// AccountTax: The retrieved or updated account tax settings.
	AccountTax *AccountTax `json:"accountTax,omitempty"`

	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accounttaxCustomBatchResponseEntry".
	Kind string `json:"kind,omitempty"`
}

type AccounttaxListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#accounttaxListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// account tax settings.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*AccountTax `json:"resources,omitempty"`
}

type Datafeed struct {
	// AttributeLanguage: The two-letter ISO 639-1 language in which the
	// attributes are defined in the data feed.
	AttributeLanguage string `json:"attributeLanguage,omitempty"`

	// ContentLanguage: The two-letter ISO 639-1 language of the items in
	// the feed.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// ContentType: The type of data feed.
	ContentType string `json:"contentType,omitempty"`

	// FetchSchedule: Fetch schedule for the feed file.
	FetchSchedule *DatafeedFetchSchedule `json:"fetchSchedule,omitempty"`

	// FileName: The filename of the feed. All feeds must have a unique file
	// name.
	FileName string `json:"fileName,omitempty"`

	// Format: Format of the feed file.
	Format *DatafeedFormat `json:"format,omitempty"`

	// Id: The ID of the data feed.
	Id int64 `json:"id,omitempty,string"`

	// IntendedDestinations: The list of intended destinations (corresponds
	// to checked check boxes in Merchant Center).
	IntendedDestinations []string `json:"intendedDestinations,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#datafeed".
	Kind string `json:"kind,omitempty"`

	// Name: A descriptive name of the data feed.
	Name string `json:"name,omitempty"`

	// TargetCountry: The two-letter ISO 3166 country where the items in the
	// feed will be included in the search index.
	TargetCountry string `json:"targetCountry,omitempty"`
}

type DatafeedFetchSchedule struct {
	// DayOfMonth: The day of the month the feed file should be fetched
	// (1-31).
	DayOfMonth int64 `json:"dayOfMonth,omitempty"`

	// FetchUrl: The URL where the feed file can be fetched. Google Merchant
	// Center will support automatic scheduled uploads using the HTTP,
	// HTTPS, FTP, or SFTP protocols, so the value will need to be a valid
	// link using one of those four protocols.
	FetchUrl string `json:"fetchUrl,omitempty"`

	// Hour: The hour of the day the feed file should be fetched (0-24).
	Hour int64 `json:"hour,omitempty"`

	// Password: An optional password for fetch_url.
	Password string `json:"password,omitempty"`

	// TimeZone: Time zone used for schedule. UTC by default. E.g.,
	// "America/Los_Angeles".
	TimeZone string `json:"timeZone,omitempty"`

	// Username: An optional user name for fetch_url.
	Username string `json:"username,omitempty"`

	// Weekday: The day of the week the feed file should be fetched.
	Weekday string `json:"weekday,omitempty"`
}

type DatafeedFormat struct {
	// ColumnDelimiter: Delimiter for the separation of values in a
	// delimiter-separated values feed. If not specified, the delimiter will
	// be auto-detected. Ignored for non-DSV data feeds.
	ColumnDelimiter string `json:"columnDelimiter,omitempty"`

	// FileEncoding: Character encoding scheme of the data feed. If not
	// specified, the encoding will be auto-detected.
	FileEncoding string `json:"fileEncoding,omitempty"`

	// QuotingMode: Specifies how double quotes are interpreted. If not
	// specified, the mode will be auto-detected. Ignored for non-DSV data
	// feeds.
	QuotingMode string `json:"quotingMode,omitempty"`
}

type DatafeedStatus struct {
	// DatafeedId: The ID of the feed for which the status is reported.
	DatafeedId uint64 `json:"datafeedId,omitempty,string"`

	// Errors: The list of errors occurring in the feed.
	Errors []*DatafeedStatusError `json:"errors,omitempty"`

	// ItemsTotal: The number of items in the feed that were processed.
	ItemsTotal uint64 `json:"itemsTotal,omitempty,string"`

	// ItemsValid: The number of items in the feed that were valid.
	ItemsValid uint64 `json:"itemsValid,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#datafeedStatus".
	Kind string `json:"kind,omitempty"`

	// LastUploadDate: The last date at which the feed was uploaded.
	LastUploadDate string `json:"lastUploadDate,omitempty"`

	// ProcessingStatus: The processing status of the feed.
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// Warnings: The list of errors occurring in the feed.
	Warnings []*DatafeedStatusError `json:"warnings,omitempty"`
}

type DatafeedStatusError struct {
	// Code: The code of the error, e.g., "validation/invalid_value".
	Code string `json:"code,omitempty"`

	// Count: The number of occurrences of the error in the feed.
	Count uint64 `json:"count,omitempty,string"`

	// Examples: A list of example occurrences of the error, grouped by
	// product.
	Examples []*DatafeedStatusExample `json:"examples,omitempty"`

	// Message: The error message, e.g., "Invalid price".
	Message string `json:"message,omitempty"`
}

type DatafeedStatusExample struct {
	// ItemId: The ID of the example item.
	ItemId string `json:"itemId,omitempty"`

	// LineNumber: Line number in the data feed where the example is found.
	LineNumber uint64 `json:"lineNumber,omitempty,string"`

	// Value: The problematic value.
	Value string `json:"value,omitempty"`
}

type DatafeedsCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*DatafeedsCustomBatchRequestEntry `json:"entries,omitempty"`
}

type DatafeedsCustomBatchRequestEntry struct {
	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// Datafeed: The data feed to insert.
	Datafeed *Datafeed `json:"datafeed,omitempty"`

	// DatafeedId: The ID of the data feed to get or delete.
	DatafeedId uint64 `json:"datafeedId,omitempty,string"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`
}

type DatafeedsCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*DatafeedsCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#datafeedsCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type DatafeedsCustomBatchResponseEntry struct {
	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Datafeed: The requested data feed. Defined if and only if the request
	// was successful.
	Datafeed *Datafeed `json:"datafeed,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`
}

type DatafeedsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#datafeedsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// datafeeds.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*Datafeed `json:"resources,omitempty"`
}

type DatafeedstatusesCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*DatafeedstatusesCustomBatchRequestEntry `json:"entries,omitempty"`
}

type DatafeedstatusesCustomBatchRequestEntry struct {
	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// DatafeedId: The ID of the data feed to get or delete.
	DatafeedId uint64 `json:"datafeedId,omitempty,string"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`
}

type DatafeedstatusesCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*DatafeedstatusesCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#datafeedstatusesCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type DatafeedstatusesCustomBatchResponseEntry struct {
	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// DatafeedStatus: The requested data feed status. Defined if and only
	// if the request was successful.
	DatafeedStatus *DatafeedStatus `json:"datafeedStatus,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`
}

type DatafeedstatusesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#datafeedstatusesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// datafeed statuses.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*DatafeedStatus `json:"resources,omitempty"`
}

type Error struct {
	// Domain: The domain of the error.
	Domain string `json:"domain,omitempty"`

	// Message: A description of the error.
	Message string `json:"message,omitempty"`

	// Reason: The error code.
	Reason string `json:"reason,omitempty"`
}

type Errors struct {
	// Code: The HTTP status of the first error in errors.
	Code int64 `json:"code,omitempty"`

	// Errors: A list of errors.
	Errors []*Error `json:"errors,omitempty"`

	// Message: The message of the first error in errors.
	Message string `json:"message,omitempty"`
}

type Inventory struct {
	// Availability: The availability of the product.
	Availability string `json:"availability,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#inventory".
	Kind string `json:"kind,omitempty"`

	// Price: The price of the product.
	Price *Price `json:"price,omitempty"`

	// Quantity: The quantity of the product. Must be equal to or greater
	// than zero. Supported only for local products.
	Quantity int64 `json:"quantity,omitempty"`

	// SalePrice: The sale price of the product. Mandatory if
	// sale_price_effective_date is defined.
	SalePrice *Price `json:"salePrice,omitempty"`

	// SalePriceEffectiveDate: A date range represented by a pair of ISO
	// 8601 dates separated by a space, comma, or slash. Both dates might be
	// specified as 'null' if undecided.
	SalePriceEffectiveDate string `json:"salePriceEffectiveDate,omitempty"`
}

type InventoryCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*InventoryCustomBatchRequestEntry `json:"entries,omitempty"`
}

type InventoryCustomBatchRequestEntry struct {
	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// Inventory: Price and availability of the product.
	Inventory *Inventory `json:"inventory,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	// ProductId: The ID of the product for which to update price and
	// availability.
	ProductId string `json:"productId,omitempty"`

	// StoreCode: The code of the store for which to update price and
	// availability. Use online to update price and availability of an
	// online product.
	StoreCode string `json:"storeCode,omitempty"`
}

type InventoryCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*InventoryCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#inventoryCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type InventoryCustomBatchResponseEntry struct {
	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#inventoryCustomBatchResponseEntry".
	Kind string `json:"kind,omitempty"`
}

type InventorySetRequest struct {
	// Availability: The availability of the product.
	Availability string `json:"availability,omitempty"`

	// Price: The price of the product.
	Price *Price `json:"price,omitempty"`

	// Quantity: The quantity of the product. Must be equal to or greater
	// than zero. Supported only for local products.
	Quantity int64 `json:"quantity,omitempty"`

	// SalePrice: The sale price of the product. Mandatory if
	// sale_price_effective_date is defined.
	SalePrice *Price `json:"salePrice,omitempty"`

	// SalePriceEffectiveDate: A date range represented by a pair of ISO
	// 8601 dates separated by a space, comma, or slash. Both dates might be
	// specified as 'null' if undecided.
	SalePriceEffectiveDate string `json:"salePriceEffectiveDate,omitempty"`
}

type InventorySetResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#inventorySetResponse".
	Kind string `json:"kind,omitempty"`
}

type LoyaltyPoints struct {
	// Name: Name of loyalty points program. It is recommended to limit the
	// name to 12 full-width characters or 24 Roman characters.
	Name string `json:"name,omitempty"`

	// PointsValue: The retailer's loyalty points in absolute value.
	PointsValue int64 `json:"pointsValue,omitempty,string"`

	// Ratio: The ratio of a point when converted to currency. Google
	// assumes currency based on Merchant Center settings. If ratio is left
	// out, it defaults to 1.0.
	Ratio float64 `json:"ratio,omitempty"`
}

type Price struct {
	// Currency: The currency of the price.
	Currency string `json:"currency,omitempty"`

	// Value: The price represented as a number.
	Value string `json:"value,omitempty"`
}

type Product struct {
	// AdditionalImageLinks: Additional URLs of images of the item.
	AdditionalImageLinks []string `json:"additionalImageLinks,omitempty"`

	// Adult: Set to true if the item is targeted towards adults.
	Adult bool `json:"adult,omitempty"`

	// AdwordsGrouping: Used to group items in an arbitrary way. Only for
	// CPA%, discouraged otherwise.
	AdwordsGrouping string `json:"adwordsGrouping,omitempty"`

	// AdwordsLabels: Similar to adwords_grouping, but only works on CPC.
	AdwordsLabels []string `json:"adwordsLabels,omitempty"`

	// AdwordsRedirect: Allows advertisers to override the item URL when the
	// product is shown within the context of Product Ads.
	AdwordsRedirect string `json:"adwordsRedirect,omitempty"`

	// AgeGroup: Target age group of the item.
	AgeGroup string `json:"ageGroup,omitempty"`

	// Aspects: Specifies the intended aspects for the product.
	Aspects []*ProductAspect `json:"aspects,omitempty"`

	// Availability: Availability status of the item.
	Availability string `json:"availability,omitempty"`

	// AvailabilityDate: The day a pre-ordered product becomes available for
	// delivery, in ISO 8601 format.
	AvailabilityDate string `json:"availabilityDate,omitempty"`

	// Brand: Brand of the item.
	Brand string `json:"brand,omitempty"`

	// Channel: The item's channel (online or local).
	Channel string `json:"channel,omitempty"`

	// Color: Color of the item.
	Color string `json:"color,omitempty"`

	// Condition: Condition or state of the item.
	Condition string `json:"condition,omitempty"`

	// ContentLanguage: The two-letter ISO 639-1 language code for the item.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// CustomAttributes: A list of custom (merchant-provided) attributes. It
	// can also be used for submitting any attribute of the feed
	// specification in its generic form (e.g., { "name": "size type",
	// "type": "text", "value": "regular" }). This is useful for submitting
	// attributes not explicitly exposed by the API.
	CustomAttributes []*ProductCustomAttribute `json:"customAttributes,omitempty"`

	// CustomGroups: A list of custom (merchant-provided) custom attribute
	// groups.
	CustomGroups []*ProductCustomGroup `json:"customGroups,omitempty"`

	// CustomLabel0: Custom label 0 for custom grouping of items in a
	// Shopping campaign.
	CustomLabel0 string `json:"customLabel0,omitempty"`

	// CustomLabel1: Custom label 1 for custom grouping of items in a
	// Shopping campaign.
	CustomLabel1 string `json:"customLabel1,omitempty"`

	// CustomLabel2: Custom label 2 for custom grouping of items in a
	// Shopping campaign.
	CustomLabel2 string `json:"customLabel2,omitempty"`

	// CustomLabel3: Custom label 3 for custom grouping of items in a
	// Shopping campaign.
	CustomLabel3 string `json:"customLabel3,omitempty"`

	// CustomLabel4: Custom label 4 for custom grouping of items in a
	// Shopping campaign.
	CustomLabel4 string `json:"customLabel4,omitempty"`

	// Description: Description of the item.
	Description string `json:"description,omitempty"`

	// Destinations: Specifies the intended destinations for the product.
	Destinations []*ProductDestination `json:"destinations,omitempty"`

	// DisplayAdsId: An identifier for an item for dynamic remarketing
	// campaigns.
	DisplayAdsId string `json:"displayAdsId,omitempty"`

	// DisplayAdsLink: URL directly to your item's landing page for dynamic
	// remarketing campaigns.
	DisplayAdsLink string `json:"displayAdsLink,omitempty"`

	// DisplayAdsSimilarIds: Advertiser-specified recommendations.
	DisplayAdsSimilarIds []string `json:"displayAdsSimilarIds,omitempty"`

	// DisplayAdsTitle: Title of an item for dynamic remarketing campaigns.
	DisplayAdsTitle string `json:"displayAdsTitle,omitempty"`

	// DisplayAdsValue: Offer margin for dynamic remarketing campaigns.
	DisplayAdsValue float64 `json:"displayAdsValue,omitempty"`

	// EnergyEfficiencyClass: The energy efficiency class as defined in EU
	// directive 2010/30/EU.
	EnergyEfficiencyClass string `json:"energyEfficiencyClass,omitempty"`

	// ExpirationDate: Date on which the item should expire, as specified
	// upon insertion, in ISO 8601 format. The actual expiration date in
	// Google Shopping is exposed in productstatuses as googleExpirationDate
	// and might be earlier if expirationDate is too far in the future.
	ExpirationDate string `json:"expirationDate,omitempty"`

	// Gender: Target gender of the item.
	Gender string `json:"gender,omitempty"`

	// GoogleProductCategory: Google's category of the item (see Google
	// product taxonomy).
	GoogleProductCategory string `json:"googleProductCategory,omitempty"`

	// Gtin: Global Trade Item Number (GTIN) of the item.
	Gtin string `json:"gtin,omitempty"`

	// Id: The REST id of the product.
	Id string `json:"id,omitempty"`

	// IdentifierExists: False when the item does not have unique product
	// identifiers appropriate to its category, such as GTIN, MPN, and
	// brand. Required according to the Unique Product Identifier Rules for
	// all target countries except for Canada.
	IdentifierExists bool `json:"identifierExists,omitempty"`

	// ImageLink: URL of an image of the item.
	ImageLink string `json:"imageLink,omitempty"`

	// Installment: Number and amount of installments to pay for an item.
	// Brazil only.
	Installment *ProductInstallment `json:"installment,omitempty"`

	// IsBundle: Whether the item is a merchant-defined bundle. A bundle is
	// a custom grouping of different products sold by a merchant for a
	// single price.
	IsBundle bool `json:"isBundle,omitempty"`

	// ItemGroupId: Shared identifier for all variants of the same product.
	ItemGroupId string `json:"itemGroupId,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#product".
	Kind string `json:"kind,omitempty"`

	// Link: URL directly linking to your item's page on your website.
	Link string `json:"link,omitempty"`

	// LoyaltyPoints: Loyalty points that users receive after purchasing the
	// item. Japan only.
	LoyaltyPoints *LoyaltyPoints `json:"loyaltyPoints,omitempty"`

	// Material: The material of which the item is made.
	Material string `json:"material,omitempty"`

	// MobileLink: Link to a mobile-optimized version of the landing page.
	MobileLink string `json:"mobileLink,omitempty"`

	// Mpn: Manufacturer Part Number (MPN) of the item.
	Mpn string `json:"mpn,omitempty"`

	// Multipack: The number of identical products in a merchant-defined
	// multipack.
	Multipack int64 `json:"multipack,omitempty,string"`

	// OfferId: An identifier of the item.
	OfferId string `json:"offerId,omitempty"`

	// OnlineOnly: Whether an item is available for purchase only online.
	OnlineOnly bool `json:"onlineOnly,omitempty"`

	// Pattern: The item's pattern (e.g. polka dots).
	Pattern string `json:"pattern,omitempty"`

	// Price: Price of the item.
	Price *Price `json:"price,omitempty"`

	// ProductType: Your category of the item (formatted as in product feeds
	// specification).
	ProductType string `json:"productType,omitempty"`

	// SalePrice: Advertised sale price of the item.
	SalePrice *Price `json:"salePrice,omitempty"`

	// SalePriceEffectiveDate: Date range during which the item is on sale
	// (see product feed specifications).
	SalePriceEffectiveDate string `json:"salePriceEffectiveDate,omitempty"`

	// Shipping: Shipping rules.
	Shipping []*ProductShipping `json:"shipping,omitempty"`

	// ShippingHeight: Height of the item for shipping.
	ShippingHeight *ProductShippingDimension `json:"shippingHeight,omitempty"`

	// ShippingLabel: The shipping label of the product, used to group
	// product in account-level shipping rules.
	ShippingLabel string `json:"shippingLabel,omitempty"`

	// ShippingLength: Length of the item for shipping.
	ShippingLength *ProductShippingDimension `json:"shippingLength,omitempty"`

	// ShippingWeight: Weight of the item for shipping.
	ShippingWeight *ProductShippingWeight `json:"shippingWeight,omitempty"`

	// ShippingWidth: Width of the item for shipping.
	ShippingWidth *ProductShippingDimension `json:"shippingWidth,omitempty"`

	// SizeSystem: System in which the size is specified. Recommended for
	// apparel items.
	SizeSystem string `json:"sizeSystem,omitempty"`

	// SizeType: The cut of the item. Recommended for apparel items.
	SizeType string `json:"sizeType,omitempty"`

	// Sizes: Size of the item.
	Sizes []string `json:"sizes,omitempty"`

	// TargetCountry: The two-letter ISO 3166 country code for the item.
	TargetCountry string `json:"targetCountry,omitempty"`

	// Taxes: Tax information.
	Taxes []*ProductTax `json:"taxes,omitempty"`

	// Title: Title of the item.
	Title string `json:"title,omitempty"`

	// UnitPricingBaseMeasure: The preference of the denominator of the unit
	// price.
	UnitPricingBaseMeasure *ProductUnitPricingBaseMeasure `json:"unitPricingBaseMeasure,omitempty"`

	// UnitPricingMeasure: The measure and dimension of an item.
	UnitPricingMeasure *ProductUnitPricingMeasure `json:"unitPricingMeasure,omitempty"`

	// ValidatedDestinations: The read-only list of intended destinations
	// which passed validation.
	ValidatedDestinations []string `json:"validatedDestinations,omitempty"`

	// Warnings: Read-only warnings.
	Warnings []*Error `json:"warnings,omitempty"`
}

type ProductAspect struct {
	// AspectName: The name of the aspect.
	AspectName string `json:"aspectName,omitempty"`

	// DestinationName: The name of the destination. Leave out to apply to
	// all destinations.
	DestinationName string `json:"destinationName,omitempty"`

	// Intention: Whether the aspect is required, excluded or should be
	// validated.
	Intention string `json:"intention,omitempty"`
}

type ProductCustomAttribute struct {
	// Name: The name of the attribute. Underscores will be replaced by
	// spaces upon insertion.
	Name string `json:"name,omitempty"`

	// Type: The type of the attribute.
	Type string `json:"type,omitempty"`

	// Unit: Free-form unit of the attribute. Unit can only be used for
	// values of type INT or FLOAT.
	Unit string `json:"unit,omitempty"`

	// Value: The value of the attribute.
	Value string `json:"value,omitempty"`
}

type ProductCustomGroup struct {
	// Attributes: The sub-attributes.
	Attributes []*ProductCustomAttribute `json:"attributes,omitempty"`

	// Name: The name of the group. Underscores will be replaced by spaces
	// upon insertion.
	Name string `json:"name,omitempty"`
}

type ProductDestination struct {
	// DestinationName: The name of the destination.
	DestinationName string `json:"destinationName,omitempty"`

	// Intention: Whether the destination is required, excluded or should be
	// validated.
	Intention string `json:"intention,omitempty"`
}

type ProductInstallment struct {
	// Amount: The amount the buyer has to pay per month.
	Amount *Price `json:"amount,omitempty"`

	// Months: The number of installments the buyer has to pay.
	Months int64 `json:"months,omitempty,string"`
}

type ProductShipping struct {
	// Country: The two-letter ISO 3166 country code for the country to
	// which an item will ship.
	Country string `json:"country,omitempty"`

	// LocationGroupName: The location where the shipping is applicable,
	// represented by a location group name.
	LocationGroupName string `json:"locationGroupName,omitempty"`

	// LocationId: The numeric id of a location that the shipping rate
	// applies to as defined in the AdWords API.
	LocationId int64 `json:"locationId,omitempty,string"`

	// PostalCode: The postal code range that the shipping rate applies to,
	// represented by a postal code, a postal code prefix followed by a *
	// wildcard, a range between two postal codes or two postal code
	// prefixes of equal length.
	PostalCode string `json:"postalCode,omitempty"`

	// Price: Fixed shipping price, represented as a number.
	Price *Price `json:"price,omitempty"`

	// Region: The geographic region to which a shipping rate applies (e.g.
	// zip code).
	Region string `json:"region,omitempty"`

	// Service: A free-form description of the service class or delivery
	// speed.
	Service string `json:"service,omitempty"`
}

type ProductShippingDimension struct {
	// Unit: The unit of value.
	//
	// Acceptable values are:
	// - "cm"
	// - "in"
	Unit string `json:"unit,omitempty"`

	// Value: The dimension of the product used to calculate the shipping
	// cost of the item.
	Value float64 `json:"value,omitempty"`
}

type ProductShippingWeight struct {
	// Unit: The unit of value.
	Unit string `json:"unit,omitempty"`

	// Value: The weight of the product used to calculate the shipping cost
	// of the item.
	Value float64 `json:"value,omitempty"`
}

type ProductStatus struct {
	// CreationDate: Date on which the item has been created, in ISO 8601
	// format.
	CreationDate string `json:"creationDate,omitempty"`

	// DataQualityIssues: A list of data quality issues associated with the
	// product.
	DataQualityIssues []*ProductStatusDataQualityIssue `json:"dataQualityIssues,omitempty"`

	// DestinationStatuses: The intended destinations for the product.
	DestinationStatuses []*ProductStatusDestinationStatus `json:"destinationStatuses,omitempty"`

	// GoogleExpirationDate: Date on which the item expires in Google
	// Shopping, in ISO 8601 format.
	GoogleExpirationDate string `json:"googleExpirationDate,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productStatus".
	Kind string `json:"kind,omitempty"`

	// LastUpdateDate: Date on which the item has been last updated, in ISO
	// 8601 format.
	LastUpdateDate string `json:"lastUpdateDate,omitempty"`

	// Link: The link to the product.
	Link string `json:"link,omitempty"`

	// ProductId: The id of the product for which status is reported.
	ProductId string `json:"productId,omitempty"`

	// Title: The title of the product.
	Title string `json:"title,omitempty"`
}

type ProductStatusDataQualityIssue struct {
	// Detail: A more detailed error string.
	Detail string `json:"detail,omitempty"`

	// FetchStatus: The fetch status for landing_page_errors.
	FetchStatus string `json:"fetchStatus,omitempty"`

	// Id: The id of the data quality issue.
	Id string `json:"id,omitempty"`

	// Location: The attribute name that is relevant for the issue.
	Location string `json:"location,omitempty"`

	// Severity: The severity of the data quality issue.
	Severity string `json:"severity,omitempty"`

	// Timestamp: The time stamp of the data quality issue.
	Timestamp string `json:"timestamp,omitempty"`

	// ValueOnLandingPage: The value of that attribute that was found on the
	// landing page
	ValueOnLandingPage string `json:"valueOnLandingPage,omitempty"`

	// ValueProvided: The value the attribute had at time of evaluation.
	ValueProvided string `json:"valueProvided,omitempty"`
}

type ProductStatusDestinationStatus struct {
	// ApprovalStatus: The destination's approval status.
	ApprovalStatus string `json:"approvalStatus,omitempty"`

	// Destination: The name of the destination
	Destination string `json:"destination,omitempty"`

	// Intention: Whether the destination is required, excluded, selected by
	// default or should be validated.
	Intention string `json:"intention,omitempty"`
}

type ProductTax struct {
	// Country: The country within which the item is taxed, specified with a
	// two-letter ISO 3166 country code.
	Country string `json:"country,omitempty"`

	// LocationId: The numeric id of a location that the tax rate applies to
	// as defined in the AdWords API.
	LocationId int64 `json:"locationId,omitempty,string"`

	// PostalCode: The postal code range that the tax rate applies to,
	// represented by a ZIP code, a ZIP code prefix using * wildcard, a
	// range between two ZIP codes or two ZIP code prefixes of equal length.
	// Examples: 94114, 94*, 94002-95460, 94*-95*.
	PostalCode string `json:"postalCode,omitempty"`

	// Rate: The percentage of tax rate that applies to the item price.
	Rate float64 `json:"rate,omitempty"`

	// Region: The geographic region to which the tax rate applies.
	Region string `json:"region,omitempty"`

	// TaxShip: Set to true if tax is charged on shipping.
	TaxShip bool `json:"taxShip,omitempty"`
}

type ProductUnitPricingBaseMeasure struct {
	// Unit: The unit of the denominator.
	Unit string `json:"unit,omitempty"`

	// Value: The denominator of the unit price.
	Value int64 `json:"value,omitempty,string"`
}

type ProductUnitPricingMeasure struct {
	// Unit: The unit of the measure.
	Unit string `json:"unit,omitempty"`

	// Value: The measure of an item.
	Value float64 `json:"value,omitempty"`
}

type ProductsCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*ProductsCustomBatchRequestEntry `json:"entries,omitempty"`
}

type ProductsCustomBatchRequestEntry struct {
	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`

	// Product: The product to insert. Only required if the method is
	// insert.
	Product *Product `json:"product,omitempty"`

	// ProductId: The ID of the product to get or delete. Only defined if
	// the method is get or delete.
	ProductId string `json:"productId,omitempty"`
}

type ProductsCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*ProductsCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productsCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type ProductsCustomBatchResponseEntry struct {
	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors defined if and only if the request failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productsCustomBatchResponseEntry".
	Kind string `json:"kind,omitempty"`

	// Product: The inserted product. Only defined if the method is insert
	// and if the request was successful.
	Product *Product `json:"product,omitempty"`
}

type ProductsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// products.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*Product `json:"resources,omitempty"`
}

type ProductstatusesCustomBatchRequest struct {
	// Entries: The request entries to be processed in the batch.
	Entries []*ProductstatusesCustomBatchRequestEntry `json:"entries,omitempty"`
}

type ProductstatusesCustomBatchRequestEntry struct {
	// BatchId: An entry ID, unique within the batch request.
	BatchId int64 `json:"batchId,omitempty"`

	// MerchantId: The ID of the managing account.
	MerchantId uint64 `json:"merchantId,omitempty,string"`

	Method string `json:"method,omitempty"`

	// ProductId: The ID of the product whose status to get.
	ProductId string `json:"productId,omitempty"`
}

type ProductstatusesCustomBatchResponse struct {
	// Entries: The result of the execution of the batch requests.
	Entries []*ProductstatusesCustomBatchResponseEntry `json:"entries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productstatusesCustomBatchResponse".
	Kind string `json:"kind,omitempty"`
}

type ProductstatusesCustomBatchResponseEntry struct {
	// BatchId: The ID of the request entry this entry responds to.
	BatchId int64 `json:"batchId,omitempty"`

	// Errors: A list of errors, if the request failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productstatusesCustomBatchResponseEntry".
	Kind string `json:"kind,omitempty"`

	// ProductStatus: The requested product status. Only defined if the
	// request was successful.
	ProductStatus *ProductStatus `json:"productStatus,omitempty"`
}

type ProductstatusesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "content#productstatusesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token for the retrieval of the next page of
	// products statuses.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Resources []*ProductStatus `json:"resources,omitempty"`
}

type Weight struct {
	// Unit: The weight unit.
	Unit string `json:"unit,omitempty"`

	// Value: The weight represented as a number.
	Value string `json:"value,omitempty"`
}

// method id "content.accounts.authinfo":

type AccountsAuthinfoCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Authinfo: Returns information about the authenticated user.

func (r *AccountsService) Authinfo() *AccountsAuthinfoCall {
	return &AccountsAuthinfoCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/authinfo",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsAuthinfoCall) Context(ctx context.Context) *AccountsAuthinfoCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAuthinfoCall) Fields(s ...googleapi.Field) *AccountsAuthinfoCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAuthinfoCall) Do() (*AccountsAuthInfoResponse, error) {
	var returnValue *AccountsAuthInfoResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns information about the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "content.accounts.authinfo",
	//   "path": "accounts/authinfo",
	//   "response": {
	//     "$ref": "AccountsAuthInfoResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.custombatch":

type AccountsCustombatchCall struct {
	s                          *Service
	accountscustombatchrequest *AccountsCustomBatchRequest
	caller_                    googleapi.Caller
	params_                    url.Values
	pathTemplate_              string
	context_                   context.Context
}

// Custombatch: Retrieves, inserts, updates, and deletes multiple
// Merchant Center (sub-)accounts in a single request.

func (r *AccountsService) Custombatch(accountscustombatchrequest *AccountsCustomBatchRequest) *AccountsCustombatchCall {
	return &AccountsCustombatchCall{
		s: r.s,
		accountscustombatchrequest: accountscustombatchrequest,
		caller_:                    googleapi.JSONCall{},
		params_:                    make(map[string][]string),
		pathTemplate_:              "accounts/batch",
		context_:                   googleapi.NoContext,
	}
}
func (c *AccountsCustombatchCall) Context(ctx context.Context) *AccountsCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustombatchCall) Fields(s ...googleapi.Field) *AccountsCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsCustombatchCall) Do() (*AccountsCustomBatchResponse, error) {
	var returnValue *AccountsCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountscustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves, inserts, updates, and deletes multiple Merchant Center (sub-)accounts in a single request.",
	//   "httpMethod": "POST",
	//   "id": "content.accounts.custombatch",
	//   "path": "accounts/batch",
	//   "request": {
	//     "$ref": "AccountsCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "AccountsCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.delete":

type AccountsDeleteCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes a Merchant Center sub-account.

func (r *AccountsService) Delete(merchantId uint64, accountId uint64) *AccountsDeleteCall {
	return &AccountsDeleteCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounts/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsDeleteCall) Context(ctx context.Context) *AccountsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *AccountsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes a Merchant Center sub-account.",
	//   "httpMethod": "DELETE",
	//   "id": "content.accounts.delete",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounts/{accountId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.get":

type AccountsGetCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves a Merchant Center account.

func (r *AccountsService) Get(merchantId uint64, accountId uint64) *AccountsGetCall {
	return &AccountsGetCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounts/{accountId}",
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
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves a Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.accounts.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounts/{accountId}",
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.insert":

type AccountsInsertCall struct {
	s             *Service
	merchantId    uint64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a Merchant Center sub-account.

func (r *AccountsService) Insert(merchantId uint64, account *Account) *AccountsInsertCall {
	return &AccountsInsertCall{
		s:             r.s,
		merchantId:    merchantId,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounts",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountsInsertCall) Context(ctx context.Context) *AccountsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsInsertCall) Fields(s ...googleapi.Field) *AccountsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsInsertCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.account,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a Merchant Center sub-account.",
	//   "httpMethod": "POST",
	//   "id": "content.accounts.insert",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounts",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.list":

type AccountsListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the sub-accounts in your Merchant Center account.

func (r *AccountsService) List(merchantId uint64) *AccountsListCall {
	return &AccountsListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounts",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of accounts to return in the response, used for paging.
func (c *AccountsListCall) MaxResults(maxResults int64) *AccountsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AccountsListCall) PageToken(pageToken string) *AccountsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
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

func (c *AccountsListCall) Do() (*AccountsListResponse, error) {
	var returnValue *AccountsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the sub-accounts in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.accounts.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of accounts to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounts",
	//   "response": {
	//     "$ref": "AccountsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.patch":

type AccountsPatchCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates a Merchant Center account. This method supports patch
// semantics.

func (r *AccountsService) Patch(merchantId uint64, accountId uint64, account *Account) *AccountsPatchCall {
	return &AccountsPatchCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounts/{accountId}",
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
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
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
	//   "description": "Updates a Merchant Center account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "content.accounts.patch",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounts/{accountId}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounts.update":

type AccountsUpdateCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates a Merchant Center account.

func (r *AccountsService) Update(merchantId uint64, accountId uint64, account *Account) *AccountsUpdateCall {
	return &AccountsUpdateCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounts/{accountId}",
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
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
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
	//   "description": "Updates a Merchant Center account.",
	//   "httpMethod": "PUT",
	//   "id": "content.accounts.update",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounts/{accountId}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountshipping.custombatch":

type AccountshippingCustombatchCall struct {
	s                                 *Service
	accountshippingcustombatchrequest *AccountshippingCustomBatchRequest
	caller_                           googleapi.Caller
	params_                           url.Values
	pathTemplate_                     string
	context_                          context.Context
}

// Custombatch: Retrieves and updates the shipping settings of multiple
// accounts in a single request.

func (r *AccountshippingService) Custombatch(accountshippingcustombatchrequest *AccountshippingCustomBatchRequest) *AccountshippingCustombatchCall {
	return &AccountshippingCustombatchCall{
		s: r.s,
		accountshippingcustombatchrequest: accountshippingcustombatchrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accountshipping/batch",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountshippingCustombatchCall) Context(ctx context.Context) *AccountshippingCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountshippingCustombatchCall) Fields(s ...googleapi.Field) *AccountshippingCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountshippingCustombatchCall) Do() (*AccountshippingCustomBatchResponse, error) {
	var returnValue *AccountshippingCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountshippingcustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves and updates the shipping settings of multiple accounts in a single request.",
	//   "httpMethod": "POST",
	//   "id": "content.accountshipping.custombatch",
	//   "path": "accountshipping/batch",
	//   "request": {
	//     "$ref": "AccountshippingCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "AccountshippingCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountshipping.get":

type AccountshippingGetCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the shipping settings of the account.

func (r *AccountshippingService) Get(merchantId uint64, accountId uint64) *AccountshippingGetCall {
	return &AccountshippingGetCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accountshipping/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountshippingGetCall) Context(ctx context.Context) *AccountshippingGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountshippingGetCall) Fields(s ...googleapi.Field) *AccountshippingGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountshippingGetCall) Do() (*AccountShipping, error) {
	var returnValue *AccountShipping
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the shipping settings of the account.",
	//   "httpMethod": "GET",
	//   "id": "content.accountshipping.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account for which to get/update account shipping settings.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accountshipping/{accountId}",
	//   "response": {
	//     "$ref": "AccountShipping"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountshipping.list":

type AccountshippingListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the shipping settings of the sub-accounts in your
// Merchant Center account.

func (r *AccountshippingService) List(merchantId uint64) *AccountshippingListCall {
	return &AccountshippingListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accountshipping",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of shipping settings to return in the response, used for
// paging.
func (c *AccountshippingListCall) MaxResults(maxResults int64) *AccountshippingListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AccountshippingListCall) PageToken(pageToken string) *AccountshippingListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *AccountshippingListCall) Context(ctx context.Context) *AccountshippingListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountshippingListCall) Fields(s ...googleapi.Field) *AccountshippingListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountshippingListCall) Do() (*AccountshippingListResponse, error) {
	var returnValue *AccountshippingListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the shipping settings of the sub-accounts in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.accountshipping.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of shipping settings to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accountshipping",
	//   "response": {
	//     "$ref": "AccountshippingListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountshipping.patch":

type AccountshippingPatchCall struct {
	s               *Service
	merchantId      uint64
	accountId       uint64
	accountshipping *AccountShipping
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Patch: Updates the shipping settings of the account. This method
// supports patch semantics.

func (r *AccountshippingService) Patch(merchantId uint64, accountId uint64, accountshipping *AccountShipping) *AccountshippingPatchCall {
	return &AccountshippingPatchCall{
		s:               r.s,
		merchantId:      merchantId,
		accountId:       accountId,
		accountshipping: accountshipping,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{merchantId}/accountshipping/{accountId}",
		context_:        googleapi.NoContext,
	}
}
func (c *AccountshippingPatchCall) Context(ctx context.Context) *AccountshippingPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountshippingPatchCall) Fields(s ...googleapi.Field) *AccountshippingPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountshippingPatchCall) Do() (*AccountShipping, error) {
	var returnValue *AccountShipping
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountshipping,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the shipping settings of the account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "content.accountshipping.patch",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account for which to get/update account shipping settings.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accountshipping/{accountId}",
	//   "request": {
	//     "$ref": "AccountShipping"
	//   },
	//   "response": {
	//     "$ref": "AccountShipping"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountshipping.update":

type AccountshippingUpdateCall struct {
	s               *Service
	merchantId      uint64
	accountId       uint64
	accountshipping *AccountShipping
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Update: Updates the shipping settings of the account.

func (r *AccountshippingService) Update(merchantId uint64, accountId uint64, accountshipping *AccountShipping) *AccountshippingUpdateCall {
	return &AccountshippingUpdateCall{
		s:               r.s,
		merchantId:      merchantId,
		accountId:       accountId,
		accountshipping: accountshipping,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{merchantId}/accountshipping/{accountId}",
		context_:        googleapi.NoContext,
	}
}
func (c *AccountshippingUpdateCall) Context(ctx context.Context) *AccountshippingUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountshippingUpdateCall) Fields(s ...googleapi.Field) *AccountshippingUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountshippingUpdateCall) Do() (*AccountShipping, error) {
	var returnValue *AccountShipping
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountshipping,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the shipping settings of the account.",
	//   "httpMethod": "PUT",
	//   "id": "content.accountshipping.update",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account for which to get/update account shipping settings.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accountshipping/{accountId}",
	//   "request": {
	//     "$ref": "AccountShipping"
	//   },
	//   "response": {
	//     "$ref": "AccountShipping"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountstatuses.custombatch":

type AccountstatusesCustombatchCall struct {
	s                                 *Service
	accountstatusescustombatchrequest *AccountstatusesCustomBatchRequest
	caller_                           googleapi.Caller
	params_                           url.Values
	pathTemplate_                     string
	context_                          context.Context
}

// Custombatch:

func (r *AccountstatusesService) Custombatch(accountstatusescustombatchrequest *AccountstatusesCustomBatchRequest) *AccountstatusesCustombatchCall {
	return &AccountstatusesCustombatchCall{
		s: r.s,
		accountstatusescustombatchrequest: accountstatusescustombatchrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accountstatuses/batch",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountstatusesCustombatchCall) Context(ctx context.Context) *AccountstatusesCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountstatusesCustombatchCall) Fields(s ...googleapi.Field) *AccountstatusesCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountstatusesCustombatchCall) Do() (*AccountstatusesCustomBatchResponse, error) {
	var returnValue *AccountstatusesCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountstatusescustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "httpMethod": "POST",
	//   "id": "content.accountstatuses.custombatch",
	//   "path": "accountstatuses/batch",
	//   "request": {
	//     "$ref": "AccountstatusesCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "AccountstatusesCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountstatuses.get":

type AccountstatusesGetCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the status of a Merchant Center account.

func (r *AccountstatusesService) Get(merchantId uint64, accountId uint64) *AccountstatusesGetCall {
	return &AccountstatusesGetCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accountstatuses/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccountstatusesGetCall) Context(ctx context.Context) *AccountstatusesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountstatusesGetCall) Fields(s ...googleapi.Field) *AccountstatusesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountstatusesGetCall) Do() (*AccountStatus, error) {
	var returnValue *AccountStatus
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the status of a Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.accountstatuses.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accountstatuses/{accountId}",
	//   "response": {
	//     "$ref": "AccountStatus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accountstatuses.list":

type AccountstatusesListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the statuses of the sub-accounts in your Merchant Center
// account.

func (r *AccountstatusesService) List(merchantId uint64) *AccountstatusesListCall {
	return &AccountstatusesListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accountstatuses",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of account statuses to return in the response, used for
// paging.
func (c *AccountstatusesListCall) MaxResults(maxResults int64) *AccountstatusesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AccountstatusesListCall) PageToken(pageToken string) *AccountstatusesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *AccountstatusesListCall) Context(ctx context.Context) *AccountstatusesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountstatusesListCall) Fields(s ...googleapi.Field) *AccountstatusesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountstatusesListCall) Do() (*AccountstatusesListResponse, error) {
	var returnValue *AccountstatusesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the statuses of the sub-accounts in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.accountstatuses.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of account statuses to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accountstatuses",
	//   "response": {
	//     "$ref": "AccountstatusesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounttax.custombatch":

type AccounttaxCustombatchCall struct {
	s                            *Service
	accounttaxcustombatchrequest *AccounttaxCustomBatchRequest
	caller_                      googleapi.Caller
	params_                      url.Values
	pathTemplate_                string
	context_                     context.Context
}

// Custombatch: Retrieves and updates tax settings of multiple accounts
// in a single request.

func (r *AccounttaxService) Custombatch(accounttaxcustombatchrequest *AccounttaxCustomBatchRequest) *AccounttaxCustombatchCall {
	return &AccounttaxCustombatchCall{
		s: r.s,
		accounttaxcustombatchrequest: accounttaxcustombatchrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounttax/batch",
		context_:      googleapi.NoContext,
	}
}
func (c *AccounttaxCustombatchCall) Context(ctx context.Context) *AccounttaxCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccounttaxCustombatchCall) Fields(s ...googleapi.Field) *AccounttaxCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccounttaxCustombatchCall) Do() (*AccounttaxCustomBatchResponse, error) {
	var returnValue *AccounttaxCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.accounttaxcustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves and updates tax settings of multiple accounts in a single request.",
	//   "httpMethod": "POST",
	//   "id": "content.accounttax.custombatch",
	//   "path": "accounttax/batch",
	//   "request": {
	//     "$ref": "AccounttaxCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "AccounttaxCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounttax.get":

type AccounttaxGetCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the tax settings of the account.

func (r *AccounttaxService) Get(merchantId uint64, accountId uint64) *AccounttaxGetCall {
	return &AccounttaxGetCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounttax/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccounttaxGetCall) Context(ctx context.Context) *AccounttaxGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccounttaxGetCall) Fields(s ...googleapi.Field) *AccounttaxGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccounttaxGetCall) Do() (*AccountTax, error) {
	var returnValue *AccountTax
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the tax settings of the account.",
	//   "httpMethod": "GET",
	//   "id": "content.accounttax.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account for which to get/update account tax settings.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounttax/{accountId}",
	//   "response": {
	//     "$ref": "AccountTax"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounttax.list":

type AccounttaxListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the tax settings of the sub-accounts in your Merchant
// Center account.

func (r *AccounttaxService) List(merchantId uint64) *AccounttaxListCall {
	return &AccounttaxListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounttax",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of tax settings to return in the response, used for paging.
func (c *AccounttaxListCall) MaxResults(maxResults int64) *AccounttaxListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AccounttaxListCall) PageToken(pageToken string) *AccounttaxListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *AccounttaxListCall) Context(ctx context.Context) *AccounttaxListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccounttaxListCall) Fields(s ...googleapi.Field) *AccounttaxListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccounttaxListCall) Do() (*AccounttaxListResponse, error) {
	var returnValue *AccounttaxListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the tax settings of the sub-accounts in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.accounttax.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of tax settings to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounttax",
	//   "response": {
	//     "$ref": "AccounttaxListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounttax.patch":

type AccounttaxPatchCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	accounttax    *AccountTax
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates the tax settings of the account. This method supports
// patch semantics.

func (r *AccounttaxService) Patch(merchantId uint64, accountId uint64, accounttax *AccountTax) *AccounttaxPatchCall {
	return &AccounttaxPatchCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		accounttax:    accounttax,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounttax/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccounttaxPatchCall) Context(ctx context.Context) *AccounttaxPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccounttaxPatchCall) Fields(s ...googleapi.Field) *AccounttaxPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccounttaxPatchCall) Do() (*AccountTax, error) {
	var returnValue *AccountTax
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.accounttax,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the tax settings of the account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "content.accounttax.patch",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account for which to get/update account tax settings.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounttax/{accountId}",
	//   "request": {
	//     "$ref": "AccountTax"
	//   },
	//   "response": {
	//     "$ref": "AccountTax"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.accounttax.update":

type AccounttaxUpdateCall struct {
	s             *Service
	merchantId    uint64
	accountId     uint64
	accounttax    *AccountTax
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates the tax settings of the account.

func (r *AccounttaxService) Update(merchantId uint64, accountId uint64, accounttax *AccountTax) *AccounttaxUpdateCall {
	return &AccounttaxUpdateCall{
		s:             r.s,
		merchantId:    merchantId,
		accountId:     accountId,
		accounttax:    accounttax,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/accounttax/{accountId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AccounttaxUpdateCall) Context(ctx context.Context) *AccounttaxUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccounttaxUpdateCall) Fields(s ...googleapi.Field) *AccounttaxUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccounttaxUpdateCall) Do() (*AccountTax, error) {
	var returnValue *AccountTax
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"accountId":  strconv.FormatUint(c.accountId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.accounttax,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the tax settings of the account.",
	//   "httpMethod": "PUT",
	//   "id": "content.accounttax.update",
	//   "parameterOrder": [
	//     "merchantId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account for which to get/update account tax settings.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/accounttax/{accountId}",
	//   "request": {
	//     "$ref": "AccountTax"
	//   },
	//   "response": {
	//     "$ref": "AccountTax"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.custombatch":

type DatafeedsCustombatchCall struct {
	s                           *Service
	datafeedscustombatchrequest *DatafeedsCustomBatchRequest
	caller_                     googleapi.Caller
	params_                     url.Values
	pathTemplate_               string
	context_                    context.Context
}

// Custombatch:

func (r *DatafeedsService) Custombatch(datafeedscustombatchrequest *DatafeedsCustomBatchRequest) *DatafeedsCustombatchCall {
	return &DatafeedsCustombatchCall{
		s: r.s,
		datafeedscustombatchrequest: datafeedscustombatchrequest,
		caller_:                     googleapi.JSONCall{},
		params_:                     make(map[string][]string),
		pathTemplate_:               "datafeeds/batch",
		context_:                    googleapi.NoContext,
	}
}
func (c *DatafeedsCustombatchCall) Context(ctx context.Context) *DatafeedsCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedsCustombatchCall) Fields(s ...googleapi.Field) *DatafeedsCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedsCustombatchCall) Do() (*DatafeedsCustomBatchResponse, error) {
	var returnValue *DatafeedsCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.datafeedscustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "httpMethod": "POST",
	//   "id": "content.datafeeds.custombatch",
	//   "path": "datafeeds/batch",
	//   "request": {
	//     "$ref": "DatafeedsCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "DatafeedsCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.delete":

type DatafeedsDeleteCall struct {
	s             *Service
	merchantId    uint64
	datafeedId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes a datafeed from your Merchant Center account.

func (r *DatafeedsService) Delete(merchantId uint64, datafeedId uint64) *DatafeedsDeleteCall {
	return &DatafeedsDeleteCall{
		s:             r.s,
		merchantId:    merchantId,
		datafeedId:    datafeedId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeeds/{datafeedId}",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedsDeleteCall) Context(ctx context.Context) *DatafeedsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *DatafeedsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"datafeedId": strconv.FormatUint(c.datafeedId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes a datafeed from your Merchant Center account.",
	//   "httpMethod": "DELETE",
	//   "id": "content.datafeeds.delete",
	//   "parameterOrder": [
	//     "merchantId",
	//     "datafeedId"
	//   ],
	//   "parameters": {
	//     "datafeedId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeeds/{datafeedId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.get":

type DatafeedsGetCall struct {
	s             *Service
	merchantId    uint64
	datafeedId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves a datafeed from your Merchant Center account.

func (r *DatafeedsService) Get(merchantId uint64, datafeedId uint64) *DatafeedsGetCall {
	return &DatafeedsGetCall{
		s:             r.s,
		merchantId:    merchantId,
		datafeedId:    datafeedId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeeds/{datafeedId}",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedsGetCall) Context(ctx context.Context) *DatafeedsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedsGetCall) Fields(s ...googleapi.Field) *DatafeedsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedsGetCall) Do() (*Datafeed, error) {
	var returnValue *Datafeed
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"datafeedId": strconv.FormatUint(c.datafeedId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves a datafeed from your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.datafeeds.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "datafeedId"
	//   ],
	//   "parameters": {
	//     "datafeedId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeeds/{datafeedId}",
	//   "response": {
	//     "$ref": "Datafeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.insert":

type DatafeedsInsertCall struct {
	s             *Service
	merchantId    uint64
	datafeed      *Datafeed
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Registers a datafeed with your Merchant Center account.

func (r *DatafeedsService) Insert(merchantId uint64, datafeed *Datafeed) *DatafeedsInsertCall {
	return &DatafeedsInsertCall{
		s:             r.s,
		merchantId:    merchantId,
		datafeed:      datafeed,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeeds",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedsInsertCall) Context(ctx context.Context) *DatafeedsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedsInsertCall) Fields(s ...googleapi.Field) *DatafeedsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedsInsertCall) Do() (*Datafeed, error) {
	var returnValue *Datafeed
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.datafeed,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Registers a datafeed with your Merchant Center account.",
	//   "httpMethod": "POST",
	//   "id": "content.datafeeds.insert",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "merchantId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeeds",
	//   "request": {
	//     "$ref": "Datafeed"
	//   },
	//   "response": {
	//     "$ref": "Datafeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.list":

type DatafeedsListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the datafeeds in your Merchant Center account.

func (r *DatafeedsService) List(merchantId uint64) *DatafeedsListCall {
	return &DatafeedsListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeeds",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of products to return in the response, used for paging.
func (c *DatafeedsListCall) MaxResults(maxResults int64) *DatafeedsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *DatafeedsListCall) PageToken(pageToken string) *DatafeedsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DatafeedsListCall) Context(ctx context.Context) *DatafeedsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedsListCall) Fields(s ...googleapi.Field) *DatafeedsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedsListCall) Do() (*DatafeedsListResponse, error) {
	var returnValue *DatafeedsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the datafeeds in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.datafeeds.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of products to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeeds",
	//   "response": {
	//     "$ref": "DatafeedsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.patch":

type DatafeedsPatchCall struct {
	s             *Service
	merchantId    uint64
	datafeedId    uint64
	datafeed      *Datafeed
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates a datafeed of your Merchant Center account. This
// method supports patch semantics.

func (r *DatafeedsService) Patch(merchantId uint64, datafeedId uint64, datafeed *Datafeed) *DatafeedsPatchCall {
	return &DatafeedsPatchCall{
		s:             r.s,
		merchantId:    merchantId,
		datafeedId:    datafeedId,
		datafeed:      datafeed,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeeds/{datafeedId}",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedsPatchCall) Context(ctx context.Context) *DatafeedsPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedsPatchCall) Fields(s ...googleapi.Field) *DatafeedsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedsPatchCall) Do() (*Datafeed, error) {
	var returnValue *Datafeed
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"datafeedId": strconv.FormatUint(c.datafeedId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.datafeed,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates a datafeed of your Merchant Center account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "content.datafeeds.patch",
	//   "parameterOrder": [
	//     "merchantId",
	//     "datafeedId"
	//   ],
	//   "parameters": {
	//     "datafeedId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeeds/{datafeedId}",
	//   "request": {
	//     "$ref": "Datafeed"
	//   },
	//   "response": {
	//     "$ref": "Datafeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeeds.update":

type DatafeedsUpdateCall struct {
	s             *Service
	merchantId    uint64
	datafeedId    uint64
	datafeed      *Datafeed
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates a datafeed of your Merchant Center account.

func (r *DatafeedsService) Update(merchantId uint64, datafeedId uint64, datafeed *Datafeed) *DatafeedsUpdateCall {
	return &DatafeedsUpdateCall{
		s:             r.s,
		merchantId:    merchantId,
		datafeedId:    datafeedId,
		datafeed:      datafeed,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeeds/{datafeedId}",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedsUpdateCall) Context(ctx context.Context) *DatafeedsUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedsUpdateCall) Fields(s ...googleapi.Field) *DatafeedsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedsUpdateCall) Do() (*Datafeed, error) {
	var returnValue *Datafeed
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"datafeedId": strconv.FormatUint(c.datafeedId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.datafeed,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates a datafeed of your Merchant Center account.",
	//   "httpMethod": "PUT",
	//   "id": "content.datafeeds.update",
	//   "parameterOrder": [
	//     "merchantId",
	//     "datafeedId"
	//   ],
	//   "parameters": {
	//     "datafeedId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeeds/{datafeedId}",
	//   "request": {
	//     "$ref": "Datafeed"
	//   },
	//   "response": {
	//     "$ref": "Datafeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeedstatuses.custombatch":

type DatafeedstatusesCustombatchCall struct {
	s                                  *Service
	datafeedstatusescustombatchrequest *DatafeedstatusesCustomBatchRequest
	caller_                            googleapi.Caller
	params_                            url.Values
	pathTemplate_                      string
	context_                           context.Context
}

// Custombatch:

func (r *DatafeedstatusesService) Custombatch(datafeedstatusescustombatchrequest *DatafeedstatusesCustomBatchRequest) *DatafeedstatusesCustombatchCall {
	return &DatafeedstatusesCustombatchCall{
		s: r.s,
		datafeedstatusescustombatchrequest: datafeedstatusescustombatchrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datafeedstatuses/batch",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedstatusesCustombatchCall) Context(ctx context.Context) *DatafeedstatusesCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedstatusesCustombatchCall) Fields(s ...googleapi.Field) *DatafeedstatusesCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedstatusesCustombatchCall) Do() (*DatafeedstatusesCustomBatchResponse, error) {
	var returnValue *DatafeedstatusesCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.datafeedstatusescustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "httpMethod": "POST",
	//   "id": "content.datafeedstatuses.custombatch",
	//   "path": "datafeedstatuses/batch",
	//   "request": {
	//     "$ref": "DatafeedstatusesCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "DatafeedstatusesCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeedstatuses.get":

type DatafeedstatusesGetCall struct {
	s             *Service
	merchantId    uint64
	datafeedId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the status of a datafeed from your Merchant Center
// account.

func (r *DatafeedstatusesService) Get(merchantId uint64, datafeedId uint64) *DatafeedstatusesGetCall {
	return &DatafeedstatusesGetCall{
		s:             r.s,
		merchantId:    merchantId,
		datafeedId:    datafeedId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeedstatuses/{datafeedId}",
		context_:      googleapi.NoContext,
	}
}
func (c *DatafeedstatusesGetCall) Context(ctx context.Context) *DatafeedstatusesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedstatusesGetCall) Fields(s ...googleapi.Field) *DatafeedstatusesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedstatusesGetCall) Do() (*DatafeedStatus, error) {
	var returnValue *DatafeedStatus
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"datafeedId": strconv.FormatUint(c.datafeedId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the status of a datafeed from your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.datafeedstatuses.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "datafeedId"
	//   ],
	//   "parameters": {
	//     "datafeedId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "merchantId": {
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeedstatuses/{datafeedId}",
	//   "response": {
	//     "$ref": "DatafeedStatus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.datafeedstatuses.list":

type DatafeedstatusesListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the statuses of the datafeeds in your Merchant Center
// account.

func (r *DatafeedstatusesService) List(merchantId uint64) *DatafeedstatusesListCall {
	return &DatafeedstatusesListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/datafeedstatuses",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of products to return in the response, used for paging.
func (c *DatafeedstatusesListCall) MaxResults(maxResults int64) *DatafeedstatusesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *DatafeedstatusesListCall) PageToken(pageToken string) *DatafeedstatusesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DatafeedstatusesListCall) Context(ctx context.Context) *DatafeedstatusesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatafeedstatusesListCall) Fields(s ...googleapi.Field) *DatafeedstatusesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatafeedstatusesListCall) Do() (*DatafeedstatusesListResponse, error) {
	var returnValue *DatafeedstatusesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the statuses of the datafeeds in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.datafeedstatuses.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of products to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/datafeedstatuses",
	//   "response": {
	//     "$ref": "DatafeedstatusesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.inventory.custombatch":

type InventoryCustombatchCall struct {
	s                           *Service
	inventorycustombatchrequest *InventoryCustomBatchRequest
	caller_                     googleapi.Caller
	params_                     url.Values
	pathTemplate_               string
	context_                    context.Context
}

// Custombatch: Updates price and availability for multiple products or
// stores in a single request.

func (r *InventoryService) Custombatch(inventorycustombatchrequest *InventoryCustomBatchRequest) *InventoryCustombatchCall {
	return &InventoryCustombatchCall{
		s: r.s,
		inventorycustombatchrequest: inventorycustombatchrequest,
		caller_:                     googleapi.JSONCall{},
		params_:                     make(map[string][]string),
		pathTemplate_:               "inventory/batch",
		context_:                    googleapi.NoContext,
	}
}
func (c *InventoryCustombatchCall) Context(ctx context.Context) *InventoryCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InventoryCustombatchCall) Fields(s ...googleapi.Field) *InventoryCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InventoryCustombatchCall) Do() (*InventoryCustomBatchResponse, error) {
	var returnValue *InventoryCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.inventorycustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates price and availability for multiple products or stores in a single request.",
	//   "httpMethod": "POST",
	//   "id": "content.inventory.custombatch",
	//   "path": "inventory/batch",
	//   "request": {
	//     "$ref": "InventoryCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "InventoryCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.inventory.set":

type InventorySetCall struct {
	s                   *Service
	merchantId          uint64
	storeCode           string
	productId           string
	inventorysetrequest *InventorySetRequest
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
	context_            context.Context
}

// Set: Updates price and availability of a product in your Merchant
// Center account.

func (r *InventoryService) Set(merchantId uint64, storeCode string, productId string, inventorysetrequest *InventorySetRequest) *InventorySetCall {
	return &InventorySetCall{
		s:                   r.s,
		merchantId:          merchantId,
		storeCode:           storeCode,
		productId:           productId,
		inventorysetrequest: inventorysetrequest,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "{merchantId}/inventory/{storeCode}/products/{productId}",
		context_:            googleapi.NoContext,
	}
}
func (c *InventorySetCall) Context(ctx context.Context) *InventorySetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InventorySetCall) Fields(s ...googleapi.Field) *InventorySetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InventorySetCall) Do() (*InventorySetResponse, error) {
	var returnValue *InventorySetResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"storeCode":  c.storeCode,
		"productId":  c.productId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.inventorysetrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates price and availability of a product in your Merchant Center account.",
	//   "httpMethod": "POST",
	//   "id": "content.inventory.set",
	//   "parameterOrder": [
	//     "merchantId",
	//     "storeCode",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product for which to update price and availability.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "storeCode": {
	//       "description": "The code of the store for which to update price and availability. Use online to update price and availability of an online product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/inventory/{storeCode}/products/{productId}",
	//   "request": {
	//     "$ref": "InventorySetRequest"
	//   },
	//   "response": {
	//     "$ref": "InventorySetResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.products.custombatch":

type ProductsCustombatchCall struct {
	s                          *Service
	productscustombatchrequest *ProductsCustomBatchRequest
	caller_                    googleapi.Caller
	params_                    url.Values
	pathTemplate_              string
	context_                   context.Context
}

// Custombatch: Retrieves, inserts, and deletes multiple products in a
// single request.

func (r *ProductsService) Custombatch(productscustombatchrequest *ProductsCustomBatchRequest) *ProductsCustombatchCall {
	return &ProductsCustombatchCall{
		s: r.s,
		productscustombatchrequest: productscustombatchrequest,
		caller_:                    googleapi.JSONCall{},
		params_:                    make(map[string][]string),
		pathTemplate_:              "products/batch",
		context_:                   googleapi.NoContext,
	}
}

// DryRun sets the optional parameter "dryRun": Flag to run the request
// in dry-run mode.
func (c *ProductsCustombatchCall) DryRun(dryRun bool) *ProductsCustombatchCall {
	c.params_.Set("dryRun", fmt.Sprintf("%v", dryRun))
	return c
}
func (c *ProductsCustombatchCall) Context(ctx context.Context) *ProductsCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsCustombatchCall) Fields(s ...googleapi.Field) *ProductsCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsCustombatchCall) Do() (*ProductsCustomBatchResponse, error) {
	var returnValue *ProductsCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.productscustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves, inserts, and deletes multiple products in a single request.",
	//   "httpMethod": "POST",
	//   "id": "content.products.custombatch",
	//   "parameters": {
	//     "dryRun": {
	//       "description": "Flag to run the request in dry-run mode.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "products/batch",
	//   "request": {
	//     "$ref": "ProductsCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "ProductsCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.products.delete":

type ProductsDeleteCall struct {
	s             *Service
	merchantId    uint64
	productId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes a product from your Merchant Center account.

func (r *ProductsService) Delete(merchantId uint64, productId string) *ProductsDeleteCall {
	return &ProductsDeleteCall{
		s:             r.s,
		merchantId:    merchantId,
		productId:     productId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/products/{productId}",
		context_:      googleapi.NoContext,
	}
}

// DryRun sets the optional parameter "dryRun": Flag to run the request
// in dry-run mode.
func (c *ProductsDeleteCall) DryRun(dryRun bool) *ProductsDeleteCall {
	c.params_.Set("dryRun", fmt.Sprintf("%v", dryRun))
	return c
}
func (c *ProductsDeleteCall) Context(ctx context.Context) *ProductsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *ProductsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"productId":  c.productId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes a product from your Merchant Center account.",
	//   "httpMethod": "DELETE",
	//   "id": "content.products.delete",
	//   "parameterOrder": [
	//     "merchantId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "dryRun": {
	//       "description": "Flag to run the request in dry-run mode.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/products/{productId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.products.get":

type ProductsGetCall struct {
	s             *Service
	merchantId    uint64
	productId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves a product from your Merchant Center account.

func (r *ProductsService) Get(merchantId uint64, productId string) *ProductsGetCall {
	return &ProductsGetCall{
		s:             r.s,
		merchantId:    merchantId,
		productId:     productId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/products/{productId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProductsGetCall) Context(ctx context.Context) *ProductsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetCall) Fields(s ...googleapi.Field) *ProductsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsGetCall) Do() (*Product, error) {
	var returnValue *Product
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"productId":  c.productId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves a product from your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.products.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/products/{productId}",
	//   "response": {
	//     "$ref": "Product"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.products.insert":

type ProductsInsertCall struct {
	s             *Service
	merchantId    uint64
	product       *Product
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Uploads a product to your Merchant Center account.

func (r *ProductsService) Insert(merchantId uint64, product *Product) *ProductsInsertCall {
	return &ProductsInsertCall{
		s:             r.s,
		merchantId:    merchantId,
		product:       product,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/products",
		context_:      googleapi.NoContext,
	}
}

// DryRun sets the optional parameter "dryRun": Flag to run the request
// in dry-run mode.
func (c *ProductsInsertCall) DryRun(dryRun bool) *ProductsInsertCall {
	c.params_.Set("dryRun", fmt.Sprintf("%v", dryRun))
	return c
}
func (c *ProductsInsertCall) Context(ctx context.Context) *ProductsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsInsertCall) Fields(s ...googleapi.Field) *ProductsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsInsertCall) Do() (*Product, error) {
	var returnValue *Product
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.product,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Uploads a product to your Merchant Center account.",
	//   "httpMethod": "POST",
	//   "id": "content.products.insert",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "dryRun": {
	//       "description": "Flag to run the request in dry-run mode.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/products",
	//   "request": {
	//     "$ref": "Product"
	//   },
	//   "response": {
	//     "$ref": "Product"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.products.list":

type ProductsListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the products in your Merchant Center account.

func (r *ProductsService) List(merchantId uint64) *ProductsListCall {
	return &ProductsListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/products",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of products to return in the response, used for paging.
func (c *ProductsListCall) MaxResults(maxResults int64) *ProductsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *ProductsListCall) PageToken(pageToken string) *ProductsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ProductsListCall) Context(ctx context.Context) *ProductsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsListCall) Fields(s ...googleapi.Field) *ProductsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsListCall) Do() (*ProductsListResponse, error) {
	var returnValue *ProductsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the products in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.products.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of products to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/products",
	//   "response": {
	//     "$ref": "ProductsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.productstatuses.custombatch":

type ProductstatusesCustombatchCall struct {
	s                                 *Service
	productstatusescustombatchrequest *ProductstatusesCustomBatchRequest
	caller_                           googleapi.Caller
	params_                           url.Values
	pathTemplate_                     string
	context_                          context.Context
}

// Custombatch: Gets the statuses of multiple products in a single
// request.

func (r *ProductstatusesService) Custombatch(productstatusescustombatchrequest *ProductstatusesCustomBatchRequest) *ProductstatusesCustombatchCall {
	return &ProductstatusesCustombatchCall{
		s: r.s,
		productstatusescustombatchrequest: productstatusescustombatchrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "productstatuses/batch",
		context_:      googleapi.NoContext,
	}
}
func (c *ProductstatusesCustombatchCall) Context(ctx context.Context) *ProductstatusesCustombatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductstatusesCustombatchCall) Fields(s ...googleapi.Field) *ProductstatusesCustombatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductstatusesCustombatchCall) Do() (*ProductstatusesCustomBatchResponse, error) {
	var returnValue *ProductstatusesCustomBatchResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.productstatusescustombatchrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the statuses of multiple products in a single request.",
	//   "httpMethod": "POST",
	//   "id": "content.productstatuses.custombatch",
	//   "path": "productstatuses/batch",
	//   "request": {
	//     "$ref": "ProductstatusesCustomBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "ProductstatusesCustomBatchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.productstatuses.get":

type ProductstatusesGetCall struct {
	s             *Service
	merchantId    uint64
	productId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Gets the status of a product from your Merchant Center account.

func (r *ProductstatusesService) Get(merchantId uint64, productId string) *ProductstatusesGetCall {
	return &ProductstatusesGetCall{
		s:             r.s,
		merchantId:    merchantId,
		productId:     productId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/productstatuses/{productId}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProductstatusesGetCall) Context(ctx context.Context) *ProductstatusesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductstatusesGetCall) Fields(s ...googleapi.Field) *ProductstatusesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductstatusesGetCall) Do() (*ProductStatus, error) {
	var returnValue *ProductStatus
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
		"productId":  c.productId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the status of a product from your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.productstatuses.get",
	//   "parameterOrder": [
	//     "merchantId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/productstatuses/{productId}",
	//   "response": {
	//     "$ref": "ProductStatus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}

// method id "content.productstatuses.list":

type ProductstatusesListCall struct {
	s             *Service
	merchantId    uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Lists the statuses of the products in your Merchant Center
// account.

func (r *ProductstatusesService) List(merchantId uint64) *ProductstatusesListCall {
	return &ProductstatusesListCall{
		s:             r.s,
		merchantId:    merchantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{merchantId}/productstatuses",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of product statuses to return in the response, used for
// paging.
func (c *ProductstatusesListCall) MaxResults(maxResults int64) *ProductstatusesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *ProductstatusesListCall) PageToken(pageToken string) *ProductstatusesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ProductstatusesListCall) Context(ctx context.Context) *ProductstatusesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductstatusesListCall) Fields(s ...googleapi.Field) *ProductstatusesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductstatusesListCall) Do() (*ProductstatusesListResponse, error) {
	var returnValue *ProductstatusesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"merchantId": strconv.FormatUint(c.merchantId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Lists the statuses of the products in your Merchant Center account.",
	//   "httpMethod": "GET",
	//   "id": "content.productstatuses.list",
	//   "parameterOrder": [
	//     "merchantId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of product statuses to return in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "merchantId": {
	//       "description": "The ID of the managing account.",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{merchantId}/productstatuses",
	//   "response": {
	//     "$ref": "ProductstatusesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/content"
	//   ]
	// }

}
