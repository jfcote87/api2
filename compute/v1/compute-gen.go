// Package compute provides access to the Compute Engine API.
//
// See https://developers.google.com/compute/docs/reference/latest/
//
// Usage example:
//
//   import "github.com/jfcote87/api2/compute/v1"
//   ...
//   computeService, err := compute.New(oauthHttpClient)
package compute

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

const apiId = "compute:v1"
const apiName = "compute"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/compute/v1/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Compute Engine resources
	ComputeScope = "https://www.googleapis.com/auth/compute"

	// View your Google Compute Engine resources
	ComputeReadonlyScope = "https://www.googleapis.com/auth/compute.readonly"

	// Manage your data and permissions in Google Cloud Storage
	DevstorageFull_controlScope = "https://www.googleapis.com/auth/devstorage.full_control"

	// View your data in Google Cloud Storage
	DevstorageRead_onlyScope = "https://www.googleapis.com/auth/devstorage.read_only"

	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Addresses = NewAddressesService(s)
	s.BackendServices = NewBackendServicesService(s)
	s.DiskTypes = NewDiskTypesService(s)
	s.Disks = NewDisksService(s)
	s.Firewalls = NewFirewallsService(s)
	s.ForwardingRules = NewForwardingRulesService(s)
	s.GlobalAddresses = NewGlobalAddressesService(s)
	s.GlobalForwardingRules = NewGlobalForwardingRulesService(s)
	s.GlobalOperations = NewGlobalOperationsService(s)
	s.HttpHealthChecks = NewHttpHealthChecksService(s)
	s.Images = NewImagesService(s)
	s.InstanceTemplates = NewInstanceTemplatesService(s)
	s.Instances = NewInstancesService(s)
	s.Licenses = NewLicensesService(s)
	s.MachineTypes = NewMachineTypesService(s)
	s.Networks = NewNetworksService(s)
	s.Projects = NewProjectsService(s)
	s.RegionOperations = NewRegionOperationsService(s)
	s.Regions = NewRegionsService(s)
	s.Routes = NewRoutesService(s)
	s.Snapshots = NewSnapshotsService(s)
	s.TargetHttpProxies = NewTargetHttpProxiesService(s)
	s.TargetInstances = NewTargetInstancesService(s)
	s.TargetPools = NewTargetPoolsService(s)
	s.TargetVpnGateways = NewTargetVpnGatewaysService(s)
	s.UrlMaps = NewUrlMapsService(s)
	s.VpnTunnels = NewVpnTunnelsService(s)
	s.ZoneOperations = NewZoneOperationsService(s)
	s.Zones = NewZonesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Addresses *AddressesService

	BackendServices *BackendServicesService

	DiskTypes *DiskTypesService

	Disks *DisksService

	Firewalls *FirewallsService

	ForwardingRules *ForwardingRulesService

	GlobalAddresses *GlobalAddressesService

	GlobalForwardingRules *GlobalForwardingRulesService

	GlobalOperations *GlobalOperationsService

	HttpHealthChecks *HttpHealthChecksService

	Images *ImagesService

	InstanceTemplates *InstanceTemplatesService

	Instances *InstancesService

	Licenses *LicensesService

	MachineTypes *MachineTypesService

	Networks *NetworksService

	Projects *ProjectsService

	RegionOperations *RegionOperationsService

	Regions *RegionsService

	Routes *RoutesService

	Snapshots *SnapshotsService

	TargetHttpProxies *TargetHttpProxiesService

	TargetInstances *TargetInstancesService

	TargetPools *TargetPoolsService

	TargetVpnGateways *TargetVpnGatewaysService

	UrlMaps *UrlMapsService

	VpnTunnels *VpnTunnelsService

	ZoneOperations *ZoneOperationsService

	Zones *ZonesService
}

func NewAddressesService(s *Service) *AddressesService {
	rs := &AddressesService{s: s}
	return rs
}

type AddressesService struct {
	s *Service
}

func NewBackendServicesService(s *Service) *BackendServicesService {
	rs := &BackendServicesService{s: s}
	return rs
}

type BackendServicesService struct {
	s *Service
}

func NewDiskTypesService(s *Service) *DiskTypesService {
	rs := &DiskTypesService{s: s}
	return rs
}

type DiskTypesService struct {
	s *Service
}

func NewDisksService(s *Service) *DisksService {
	rs := &DisksService{s: s}
	return rs
}

type DisksService struct {
	s *Service
}

func NewFirewallsService(s *Service) *FirewallsService {
	rs := &FirewallsService{s: s}
	return rs
}

type FirewallsService struct {
	s *Service
}

func NewForwardingRulesService(s *Service) *ForwardingRulesService {
	rs := &ForwardingRulesService{s: s}
	return rs
}

type ForwardingRulesService struct {
	s *Service
}

func NewGlobalAddressesService(s *Service) *GlobalAddressesService {
	rs := &GlobalAddressesService{s: s}
	return rs
}

type GlobalAddressesService struct {
	s *Service
}

func NewGlobalForwardingRulesService(s *Service) *GlobalForwardingRulesService {
	rs := &GlobalForwardingRulesService{s: s}
	return rs
}

type GlobalForwardingRulesService struct {
	s *Service
}

func NewGlobalOperationsService(s *Service) *GlobalOperationsService {
	rs := &GlobalOperationsService{s: s}
	return rs
}

type GlobalOperationsService struct {
	s *Service
}

func NewHttpHealthChecksService(s *Service) *HttpHealthChecksService {
	rs := &HttpHealthChecksService{s: s}
	return rs
}

type HttpHealthChecksService struct {
	s *Service
}

func NewImagesService(s *Service) *ImagesService {
	rs := &ImagesService{s: s}
	return rs
}

type ImagesService struct {
	s *Service
}

func NewInstanceTemplatesService(s *Service) *InstanceTemplatesService {
	rs := &InstanceTemplatesService{s: s}
	return rs
}

type InstanceTemplatesService struct {
	s *Service
}

func NewInstancesService(s *Service) *InstancesService {
	rs := &InstancesService{s: s}
	return rs
}

type InstancesService struct {
	s *Service
}

func NewLicensesService(s *Service) *LicensesService {
	rs := &LicensesService{s: s}
	return rs
}

type LicensesService struct {
	s *Service
}

func NewMachineTypesService(s *Service) *MachineTypesService {
	rs := &MachineTypesService{s: s}
	return rs
}

type MachineTypesService struct {
	s *Service
}

func NewNetworksService(s *Service) *NetworksService {
	rs := &NetworksService{s: s}
	return rs
}

type NetworksService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

func NewRegionOperationsService(s *Service) *RegionOperationsService {
	rs := &RegionOperationsService{s: s}
	return rs
}

type RegionOperationsService struct {
	s *Service
}

func NewRegionsService(s *Service) *RegionsService {
	rs := &RegionsService{s: s}
	return rs
}

type RegionsService struct {
	s *Service
}

func NewRoutesService(s *Service) *RoutesService {
	rs := &RoutesService{s: s}
	return rs
}

type RoutesService struct {
	s *Service
}

func NewSnapshotsService(s *Service) *SnapshotsService {
	rs := &SnapshotsService{s: s}
	return rs
}

type SnapshotsService struct {
	s *Service
}

func NewTargetHttpProxiesService(s *Service) *TargetHttpProxiesService {
	rs := &TargetHttpProxiesService{s: s}
	return rs
}

type TargetHttpProxiesService struct {
	s *Service
}

func NewTargetInstancesService(s *Service) *TargetInstancesService {
	rs := &TargetInstancesService{s: s}
	return rs
}

type TargetInstancesService struct {
	s *Service
}

func NewTargetPoolsService(s *Service) *TargetPoolsService {
	rs := &TargetPoolsService{s: s}
	return rs
}

type TargetPoolsService struct {
	s *Service
}

func NewTargetVpnGatewaysService(s *Service) *TargetVpnGatewaysService {
	rs := &TargetVpnGatewaysService{s: s}
	return rs
}

type TargetVpnGatewaysService struct {
	s *Service
}

func NewUrlMapsService(s *Service) *UrlMapsService {
	rs := &UrlMapsService{s: s}
	return rs
}

type UrlMapsService struct {
	s *Service
}

func NewVpnTunnelsService(s *Service) *VpnTunnelsService {
	rs := &VpnTunnelsService{s: s}
	return rs
}

type VpnTunnelsService struct {
	s *Service
}

func NewZoneOperationsService(s *Service) *ZoneOperationsService {
	rs := &ZoneOperationsService{s: s}
	return rs
}

type ZoneOperationsService struct {
	s *Service
}

func NewZonesService(s *Service) *ZonesService {
	rs := &ZonesService{s: s}
	return rs
}

type ZonesService struct {
	s *Service
}

type AccessConfig struct {
	// Kind: [Output Only] Type of the resource. Always compute#accessConfig
	// for access configs.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this access configuration.
	Name string `json:"name,omitempty"`

	// NatIP: An external IP address associated with this instance. Specify
	// an unused static external IP address available to the project or
	// leave this field undefined to use an IP from a shared ephemeral IP
	// address pool. If you specify a static external IP address, it must
	// live in the same region as the zone of the instance.
	NatIP string `json:"natIP,omitempty"`

	// Type: The type of configuration. The default and only option is
	// ONE_TO_ONE_NAT.
	Type string `json:"type,omitempty"`
}

type Address struct {
	// Address: The static external IP address represented by this resource.
	Address string `json:"address,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#address for
	// addresses.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// Region: [Output Only] URL of the region where the regional address
	// resides. This field is not applicable to global addresses.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output Only] The status of the address, which can be either
	// IN_USE or RESERVED. An address that is RESERVED is currently reserved
	// and available to use. An IN_USE address is currently being used by
	// another resource and is not available.
	Status string `json:"status,omitempty"`

	// Users: [Output Only] The URLs of the resources that are using this
	// address.
	Users []string `json:"users,omitempty"`
}

type AddressAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped address lists.
	Items map[string]AddressesScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// compute#addressAggregatedList for aggregated lists of addresses.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type AddressList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Address resources.
	Items []*Address `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#addressList for
	// lists of addresses.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type AddressesScopedList struct {
	// Addresses: [Output Only] List of addresses contained in this scope.
	Addresses []*Address `json:"addresses,omitempty"`

	// Warning: [Output Only] Informational warning which replaces the list
	// of addresses when the list is empty.
	Warning *AddressesScopedListWarning `json:"warning,omitempty"`
}

type AddressesScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*AddressesScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type AddressesScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type AttachedDisk struct {
	// AutoDelete: Specifies whether the disk will be auto-deleted when the
	// instance is deleted (but not when the disk is detached from the
	// instance).
	AutoDelete bool `json:"autoDelete,omitempty"`

	// Boot: Indicates that this is a boot disk. The virtual machine will
	// use the first partition of the disk for its root filesystem.
	Boot bool `json:"boot,omitempty"`

	// DeviceName: Specifies a unique device name of your choice that is
	// reflected into the /dev/ tree of a Linux operating system running
	// within the instance. This name can be used to reference the device
	// for mounting, resizing, and so on, from within the instance.
	//
	// If not
	// specified, the server chooses a default device name to apply to this
	// disk, in the form persistent-disks-x, where x is a number assigned by
	// Google Compute Engine. This field is only applicable for persistent
	// disks.
	DeviceName string `json:"deviceName,omitempty"`

	// Index: Assigns a zero-based index to this disk, where 0 is reserved
	// for the boot disk. For example, if you have many disks attached to an
	// instance, each disk would have a unique index number. If not
	// specified, the server will choose an appropriate value.
	Index int64 `json:"index,omitempty"`

	// InitializeParams: [Input Only] Specifies the parameters for a new
	// disk that will be created alongside the new instance. Use
	// initialization parameters to create boot disks or local SSDs attached
	// to the new instance.
	//
	// This property is mutually exclusive with the
	// source property; you can only define one or the other, but not both.
	InitializeParams *AttachedDiskInitializeParams `json:"initializeParams,omitempty"`

	Interface string `json:"interface,omitempty"`

	// Kind: [Output Only] Type of the resource. Always compute#attachedDisk
	// for attached disks.
	Kind string `json:"kind,omitempty"`

	// Licenses: [Output Only] Any valid publicly visible licenses.
	Licenses []string `json:"licenses,omitempty"`

	// Mode: The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode string `json:"mode,omitempty"`

	// Source: Specifies a valid partial or full URL to an existing
	// Persistent Disk resource. This field is only applicable for
	// persistent disks.
	Source string `json:"source,omitempty"`

	// Type: Specifies the type of the disk, either SCRATCH or PERSISTENT.
	// If not specified, the default is PERSISTENT.
	Type string `json:"type,omitempty"`
}

type AttachedDiskInitializeParams struct {
	// DiskName: Specifies the disk name. If not specified, the default is
	// to use the name of the instance.
	DiskName string `json:"diskName,omitempty"`

	// DiskSizeGb: Specifies the size of the disk in base-2 GB.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty,string"`

	// DiskType: Specifies the disk type to use to create the instance. If
	// not specified, the default is pd-standard, specified using the full
	// URL. For
	// example:
	//
	// https://www.googleapis.com/compute/v1/projects/project/zones
	// /zone/diskTypes/pd-standard
	//
	// Other values include pd-ssd and
	// local-ssd. If you define this field, you can provide either the full
	// or partial URL. For example, the following are valid values:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/disk
	// Types/diskType
	// - projects/project/zones/zone/diskTypes/diskType
	// -
	// zones/zone/diskTypes/diskType
	DiskType string `json:"diskType,omitempty"`

	// SourceImage: A source image used to create the disk. You can provide
	// a private (custom) image, and Compute Engine will use the
	// corresponding image from your project. For
	// example:
	//
	// global/images/my-private-image
	//
	// Or you can provide an
	// image from a publicly-available project. For example, to use a Debian
	// image from the debian-cloud project, make sure to include the project
	// in the
	// URL:
	//
	// projects/debian-cloud/global/images/debian-7-wheezy-vYYYYMMDD
	//
	//
	// where vYYYYMMDD is the image version. The fully-qualified URL will
	// also work in both cases.
	SourceImage string `json:"sourceImage,omitempty"`
}

type Backend struct {
	// BalancingMode: The balancing mode of this backend, default is
	// UTILIZATION.
	BalancingMode string `json:"balancingMode,omitempty"`

	// CapacityScaler: The multiplier (a value between 0 and 1e6) of the max
	// capacity (CPU or RPS, depending on 'balancingMode') the group should
	// serve up to. 0 means the group is totally drained. Default value is
	// 1. Valid range is [0, 1e6].
	CapacityScaler float64 `json:"capacityScaler,omitempty"`

	// Description: An optional textual description of the resource, which
	// is provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Group: URL of a zonal Cloud Resource View resource. This resource
	// view defines the list of instances that serve traffic. Member virtual
	// machine instances from each resource view must live in the same zone
	// as the resource view itself. No two backends in a backend service are
	// allowed to use same Resource View resource.
	Group string `json:"group,omitempty"`

	// MaxRate: The max RPS of the group. Can be used with either balancing
	// mode, but required if RATE mode. For RATE mode, either maxRate or
	// maxRatePerInstance must be set.
	MaxRate int64 `json:"maxRate,omitempty"`

	// MaxRatePerInstance: The max RPS that a single backed instance can
	// handle. This is used to calculate the capacity of the group. Can be
	// used in either balancing mode. For RATE mode, either maxRate or
	// maxRatePerInstance must be set.
	MaxRatePerInstance float64 `json:"maxRatePerInstance,omitempty"`

	// MaxUtilization: Used when 'balancingMode' is UTILIZATION. This ratio
	// defines the CPU utilization target for the group. The default is 0.8.
	// Valid range is [0, 1].
	MaxUtilization float64 `json:"maxUtilization,omitempty"`
}

type BackendService struct {
	// Backends: The list of backends that serve this BackendService.
	Backends []*Backend `json:"backends,omitempty"`

	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Fingerprint: Fingerprint of this resource. A hash of the contents
	// stored in this object. This field is used in optimistic locking. This
	// field will be ignored when inserting a BackendService. An up-to-date
	// fingerprint must be provided in order to update the BackendService.
	Fingerprint string `json:"fingerprint,omitempty"`

	// HealthChecks: The list of URLs to the HttpHealthCheck resource for
	// health checking this BackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks []string `json:"healthChecks,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// Port: Deprecated in favor of port_name. The TCP port to connect on
	// the backend. The default value is 80.
	Port int64 `json:"port,omitempty"`

	// PortName: Name of backend port. The same name should appear in the
	// resource views referenced by this service. Required.
	PortName string `json:"portName,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// TimeoutSec: How many seconds to wait for the backend before
	// considering it a failed request. Default is 30 seconds.
	TimeoutSec int64 `json:"timeoutSec,omitempty"`
}

type BackendServiceGroupHealth struct {
	HealthStatus []*HealthStatus `json:"healthStatus,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`
}

type BackendServiceList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of BackendService resources.
	Items []*BackendService `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type DeprecationStatus struct {
	// Deleted: An optional RFC3339 timestamp on or after which the
	// deprecation state of this resource will be changed to DELETED.
	Deleted string `json:"deleted,omitempty"`

	// Deprecated: An optional RFC3339 timestamp on or after which the
	// deprecation state of this resource will be changed to DEPRECATED.
	Deprecated string `json:"deprecated,omitempty"`

	// Obsolete: An optional RFC3339 timestamp on or after which the
	// deprecation state of this resource will be changed to OBSOLETE.
	Obsolete string `json:"obsolete,omitempty"`

	// Replacement: The URL of the suggested replacement for a deprecated
	// resource. The suggested replacement resource must be the same kind of
	// resource as the deprecated resource.
	Replacement string `json:"replacement,omitempty"`

	// State: The deprecation state of this resource. This can be
	// DEPRECATED, OBSOLETE, or DELETED. Operations which create a new
	// resource using a DEPRECATED resource will return successfully, but
	// with a warning indicating the deprecated resource and recommending
	// its replacement. Operations which use OBSOLETE or DELETED resources
	// will be rejected and result in an error.
	State string `json:"state,omitempty"`
}

type Disk struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#disk for
	// disks.
	Kind string `json:"kind,omitempty"`

	// Licenses: Any applicable publicly visible licenses.
	Licenses []string `json:"licenses,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// Options: Internal use only.
	Options string `json:"options,omitempty"`

	// SelfLink: [Output Only] Server-defined fully-qualified URL for this
	// resource.
	SelfLink string `json:"selfLink,omitempty"`

	// SizeGb: Size of the persistent disk, specified in GB. You can specify
	// this field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	//
	// If you specify this field along with sourceImage or
	// sourceSnapshot, the value of sizeGb must not be less than the size of
	// the sourceImage or the size of the snapshot.
	SizeGb int64 `json:"sizeGb,omitempty,string"`

	// SourceImage: The source image used to create this disk. If the source
	// image is deleted from the system, this field will not be set, even if
	// an image with the same name has been re-created.
	//
	// When creating a
	// disk, you can provide a private (custom) image using the following
	// input, and Compute Engine will use the corresponding image from your
	// project. For example:
	//
	// global/images/my-private-image
	//
	// Or you can
	// provide an image from a publicly-available project. For example, to
	// use a Debian image from the debian-cloud project, make sure to
	// include the project in the
	// URL:
	//
	// projects/debian-cloud/global/images/debian-7-wheezy-vYYYYMMDD
	//
	//
	// where vYYYYMMDD is the image version. The fully-qualified URL will
	// also work in both cases.
	SourceImage string `json:"sourceImage,omitempty"`

	// SourceImageId: The ID value of the image used to create this disk.
	// This value identifies the exact image that was used to create this
	// persistent disk. For example, if you created the persistent disk from
	// an image that was later deleted and recreated under the same name,
	// the source image ID would identify the exact version of the image
	// that was used.
	SourceImageId string `json:"sourceImageId,omitempty"`

	// SourceSnapshot: The source snapshot used to create this disk. You can
	// provide this as a partial or full URL to the resource. For example,
	// the following are valid values:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/global/snapshot
	// s/snapshot
	// - projects/project/global/snapshots/snapshot
	// -
	// global/snapshots/snapshot
	SourceSnapshot string `json:"sourceSnapshot,omitempty"`

	// SourceSnapshotId: [Output Only] The unique ID of the snapshot used to
	// create this disk. This value identifies the exact snapshot that was
	// used to create this persistent disk. For example, if you created the
	// persistent disk from a snapshot that was later deleted and recreated
	// under the same name, the source snapshot ID would identify the exact
	// version of the snapshot that was used.
	SourceSnapshotId string `json:"sourceSnapshotId,omitempty"`

	// Status: [Output Only] The status of disk creation. Applicable
	// statuses includes: CREATING, FAILED, READY, RESTORING.
	Status string `json:"status,omitempty"`

	// Type: URL of the disk type resource describing which disk type to use
	// to create the disk; provided by the client when the disk is created.
	Type string `json:"type,omitempty"`

	// Zone: [Output Only] URL of the zone where the disk resides.
	Zone string `json:"zone,omitempty"`
}

type DiskAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped disk lists.
	Items map[string]DisksScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// compute#diskAggregatedList for aggregated lists of persistent disks.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type DiskList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of persistent disks.
	Items []*Disk `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#diskList for
	// lists of disks.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type DiskMoveRequest struct {
	// DestinationZone: The URL of the destination zone to move the disk to.
	// This can be a full or partial URL. For example, the following are all
	// valid URLs to a zone:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone
	// -
	// projects/project/zones/zone
	// - zones/zone
	DestinationZone string `json:"destinationZone,omitempty"`

	// TargetDisk: The URL of the target disk to move. This can be a full or
	// partial URL. For example, the following are all valid URLs to a disk:
	//
	// -
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/disk
	// s/disk
	// - projects/project/zones/zone/disks/disk
	// -
	// zones/zone/disks/disk
	TargetDisk string `json:"targetDisk,omitempty"`
}

type DiskType struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// DefaultDiskSizeGb: [Output Only] Server defined default disk size in
	// GB.
	DefaultDiskSizeGb int64 `json:"defaultDiskSizeGb,omitempty,string"`

	// Deprecated: [Output Only] The deprecation status associated with this
	// disk type.
	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	// Description: [Output Only] An optional textual description of the
	// resource.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#diskType for
	// disk types.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ValidDiskSize: [Output Only] An optional textual description of the
	// valid disk size, such as "10GB-10TB".
	ValidDiskSize string `json:"validDiskSize,omitempty"`

	// Zone: [Output Only] URL of the zone where the disk type resides.
	Zone string `json:"zone,omitempty"`
}

type DiskTypeAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped disk type lists.
	Items map[string]DiskTypesScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// compute#diskTypeAggregatedList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type DiskTypeList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Disk Type resources.
	Items []*DiskType `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#diskTypeList for
	// disk types.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type DiskTypesScopedList struct {
	// DiskTypes: [Output Only] List of disk types contained in this scope.
	DiskTypes []*DiskType `json:"diskTypes,omitempty"`

	// Warning: [Output Only] Informational warning which replaces the list
	// of disk types when the list is empty.
	Warning *DiskTypesScopedListWarning `json:"warning,omitempty"`
}

type DiskTypesScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*DiskTypesScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type DiskTypesScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type DisksScopedList struct {
	// Disks: [Output Only] List of disks contained in this scope.
	Disks []*Disk `json:"disks,omitempty"`

	// Warning: [Output Only] Informational warning which replaces the list
	// of disks when the list is empty.
	Warning *DisksScopedListWarning `json:"warning,omitempty"`
}

type DisksScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*DisksScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type DisksScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type Firewall struct {
	// Allowed: The list of rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	Allowed []*FirewallAllowed `json:"allowed,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Ony] Type of the resource. Always compute#firewall for
	// firewall rules.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// Network: URL of the network resource for this firewall rule. This
	// field is required for creating an instance but optional when creating
	// a firewall rule. If not specified when creating a firewall rule, the
	// default network is used:
	// global/networks/default
	// If you choose to
	// specify this property, you can specify the network as a full or
	// partial URL. For example, the following are all valid URLs:
	// -
	// https://www.googleapis.com/compute/v1/projects/myproject/global/networ
	// ks/my-network
	// - projects/myproject/global/networks/my-network
	// -
	// global/networks/default
	Network string `json:"network,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// SourceRanges: The IP address blocks that this rule applies to,
	// expressed in CIDR format. One or both of sourceRanges and sourceTags
	// may be set.
	//
	// If both properties are set, an inbound connection is
	// allowed if the range or the tag of the source matches the
	// sourceRanges OR matches the sourceTags property; the connection does
	// not need to match both properties.
	SourceRanges []string `json:"sourceRanges,omitempty"`

	// SourceTags: A list of instance tags which this rule applies to. One
	// or both of sourceRanges and sourceTags may be set.
	//
	// If both
	// properties are set, an inbound connection is allowed if the range or
	// the tag of the source matches the sourceRanges OR matches the
	// sourceTags property; the connection does not need to match both
	// properties.
	SourceTags []string `json:"sourceTags,omitempty"`

	// TargetTags: A list of instance tags indicating sets of instances
	// located on network which may make network connections as specified in
	// allowed[]. If no targetTags are specified, the firewall rule applies
	// to all instances on the specified network.
	TargetTags []string `json:"targetTags,omitempty"`
}

type FirewallAllowed struct {
	// IPProtocol: The IP protocol that is allowed for this rule. The
	// protocol type is required when creating a firewall. This value can
	// either be one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp), or the IP protocol number.
	IPProtocol string `json:"IPProtocol,omitempty"`

	// Ports: An optional list of ports which are allowed. This field is
	// only applicable for UDP or TCP protocol. Each entry must be either an
	// integer or a range. If not specified, connections through any port
	// are allowed
	//
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	Ports []string `json:"ports,omitempty"`
}

type FirewallList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Firewall resources.
	Items []*Firewall `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#firewallList for
	// lists of firewalls.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type ForwardingRule struct {
	// IPAddress: Value of the reserved IP address that this forwarding rule
	// is serving on behalf of. For global forwarding rules, the address
	// must be a global IP; for regional forwarding rules, the address must
	// live in the same region as the forwarding rule. If left empty
	// (default value), an ephemeral IP from the same scope (global or
	// regional) will be assigned.
	IPAddress string `json:"IPAddress,omitempty"`

	// IPProtocol: The IP protocol to which this rule applies, valid options
	// are 'TCP', 'UDP', 'ESP', 'AH' or 'SCTP'.
	IPProtocol string `json:"IPProtocol,omitempty"`

	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// PortRange: Applicable only when 'IPProtocol' is 'TCP', 'UDP' or
	// 'SCTP', only packets addressed to ports in the specified range will
	// be forwarded to 'target'. If 'portRange' is left empty (default
	// value), all ports are forwarded. Forwarding rules with the same
	// [IPAddress, IPProtocol] pair must have disjoint port ranges.
	PortRange string `json:"portRange,omitempty"`

	// Region: URL of the region where the regional forwarding rule resides
	// (output only). This field is not applicable to global forwarding
	// rules.
	Region string `json:"region,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// Target: The URL of the target resource to receive the matched
	// traffic. For regional forwarding rules, this target must live in the
	// same region as the forwarding rule. For global forwarding rules, this
	// target must be a global TargetHttpProxy resource.
	Target string `json:"target,omitempty"`
}

type ForwardingRuleAggregatedList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A map of scoped forwarding rule lists.
	Items map[string]ForwardingRulesScopedList `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type ForwardingRuleList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of ForwardingRule resources.
	Items []*ForwardingRule `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type ForwardingRulesScopedList struct {
	// ForwardingRules: List of forwarding rules contained in this scope.
	ForwardingRules []*ForwardingRule `json:"forwardingRules,omitempty"`

	// Warning: Informational warning which replaces the list of forwarding
	// rules when the list is empty.
	Warning *ForwardingRulesScopedListWarning `json:"warning,omitempty"`
}

type ForwardingRulesScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*ForwardingRulesScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type ForwardingRulesScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type HealthCheckReference struct {
	HealthCheck string `json:"healthCheck,omitempty"`
}

type HealthStatus struct {
	// HealthState: Health state of the instance.
	HealthState string `json:"healthState,omitempty"`

	// Instance: URL of the instance resource.
	Instance string `json:"instance,omitempty"`

	// IpAddress: The IP address represented by this resource.
	IpAddress string `json:"ipAddress,omitempty"`

	// Port: The port on the instance.
	Port int64 `json:"port,omitempty"`
}

type HostRule struct {
	Description string `json:"description,omitempty"`

	// Hosts: The list of host patterns to match. They must be valid
	// hostnames except that they may start with *. or *-. The * acts like a
	// glob and will match any string of atoms (separated by .s and -s) to
	// the left.
	Hosts []string `json:"hosts,omitempty"`

	// PathMatcher: The name of the PathMatcher to match the path portion of
	// the URL, if the this HostRule matches the URL's host portion.
	PathMatcher string `json:"pathMatcher,omitempty"`
}

type HttpHealthCheck struct {
	// CheckIntervalSec: How often (in seconds) to send a health check. The
	// default value is 5 seconds.
	CheckIntervalSec int64 `json:"checkIntervalSec,omitempty"`

	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// HealthyThreshold: A so-far unhealthy VM will be marked healthy after
	// this many consecutive successes. The default value is 2.
	HealthyThreshold int64 `json:"healthyThreshold,omitempty"`

	// Host: The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host string `json:"host,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// Port: The TCP port number for the HTTP health check request. The
	// default value is 80.
	Port int64 `json:"port,omitempty"`

	// RequestPath: The request path of the HTTP health check request. The
	// default value is "/".
	RequestPath string `json:"requestPath,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// TimeoutSec: How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds. It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec int64 `json:"timeoutSec,omitempty"`

	// UnhealthyThreshold: A so-far healthy VM will be marked unhealthy
	// after this many consecutive failures. The default value is 2.
	UnhealthyThreshold int64 `json:"unhealthyThreshold,omitempty"`
}

type HttpHealthCheckList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of HttpHealthCheck resources.
	Items []*HttpHealthCheck `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type Image struct {
	// ArchiveSizeBytes: Size of the image tar.gz archive stored in Google
	// Cloud Storage (in bytes).
	ArchiveSizeBytes int64 `json:"archiveSizeBytes,omitempty,string"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Deprecated: The deprecation status associated with this image.
	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	// Description: Textual description of the resource; provided by the
	// client when the resource is created.
	Description string `json:"description,omitempty"`

	// DiskSizeGb: Size of the image when restored onto a persistent disk
	// (in GB).
	DiskSizeGb int64 `json:"diskSizeGb,omitempty,string"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#image for
	// images.
	Kind string `json:"kind,omitempty"`

	// Licenses: Any applicable publicly visible licenses.
	Licenses []string `json:"licenses,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// RawDisk: The parameters of the raw disk image.
	RawDisk *ImageRawDisk `json:"rawDisk,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// SourceDisk: URL of the The source disk used to create this image.
	// This can be a full or valid partial URL. You must provide either this
	// property or the rawDisk.source property but not both to create an
	// image. For example, the following are valid values:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/disk
	// /disk
	// - projects/project/zones/zone/disk/disk
	// -
	// zones/zone/disks/disk
	SourceDisk string `json:"sourceDisk,omitempty"`

	// SourceDiskId: The ID value of the disk used to create this image.
	// This value may be used to determine whether the image was taken from
	// the current or a previous instance of a given disk name.
	SourceDiskId string `json:"sourceDiskId,omitempty"`

	// SourceType: The type of the image used to create this disk. The
	// default and only value is RAW
	SourceType string `json:"sourceType,omitempty"`

	// Status: [Output Only] The status of the image. An image can be used
	// to create other resources, such as instances, only after the image
	// has been successfully created and the status is set to READY.
	// Possible values are FAILED, PENDING, or READY.
	Status string `json:"status,omitempty"`
}

type ImageRawDisk struct {
	// ContainerType: The format used to encode and transmit the block
	// device, which should be TAR. This is just a container and
	// transmission format and not a runtime format. Provided by the client
	// when the disk image is created.
	ContainerType string `json:"containerType,omitempty"`

	// Sha1Checksum: An optional SHA1 checksum of the disk image before
	// unpackaging; provided by the client when the disk image is created.
	Sha1Checksum string `json:"sha1Checksum,omitempty"`

	// Source: The full Google Cloud Storage URL where the disk image is
	// stored. You must provide either this property or the sourceDisk
	// property but not both.
	Source string `json:"source,omitempty"`
}

type ImageList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of Image resources.
	Items []*Image `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type Instance struct {
	// CanIpForward: Allows this instance to send and receive packets with
	// non-matching destination or source IPs. This is required if you plan
	// to use this instance to forward routes. For more information, see
	// Enabling IP Forwarding.
	CanIpForward bool `json:"canIpForward,omitempty"`

	// CpuPlatform: [Output Only] The CPU platform used by this instance.
	CpuPlatform string `json:"cpuPlatform,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Disks: Array of disks associated with this instance. Persistent disks
	// must be created before you can assign them.
	Disks []*AttachedDisk `json:"disks,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#instance for
	// instances.
	Kind string `json:"kind,omitempty"`

	// MachineType: Full or partial URL of the machine type resource to use
	// for this instance. This is provided by the client when the instance
	// is created. For example, the following is a valid partial
	// url:
	//
	// zones/zone/machineTypes/machine-type
	MachineType string `json:"machineType,omitempty"`

	// Metadata: The metadata key/value pairs assigned to this instance.
	// This includes custom metadata and predefined keys.
	Metadata *Metadata `json:"metadata,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// NetworkInterfaces: An array of configurations for this interface.
	// This specifies how this interface is configured to interact with
	// other network services, such as connecting to the internet.
	NetworkInterfaces []*NetworkInterface `json:"networkInterfaces,omitempty"`

	// Scheduling: Scheduling options for this instance.
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServiceAccounts: A list of service accounts, with their specified
	// scopes, authorized for this instance. Service accounts generate
	// access tokens that can be accessed through the metadata server and
	// used to authenticate applications on the instance. See Authenticating
	// from Google Compute Engine for more information.
	ServiceAccounts []*ServiceAccount `json:"serviceAccounts,omitempty"`

	// Status: [Output Only] The status of the instance. One of the
	// following values: PROVISIONING, STAGING, RUNNING, STOPPING, STOPPED,
	// TERMINATED.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional, human-readable explanation
	// of the status.
	StatusMessage string `json:"statusMessage,omitempty"`

	// Tags: A list of tags to appy to this instance. Tags are used to
	// identify valid sources or targets for network firewalls and are
	// specified by the client during instance creation. The tags can be
	// later modified by the setTags method. Each tag within the list must
	// comply with RFC1035.
	Tags *Tags `json:"tags,omitempty"`

	// Zone: [Output Only] URL of the zone where the instance resides.
	Zone string `json:"zone,omitempty"`
}

type InstanceAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped instance lists.
	Items map[string]InstancesScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// compute#instanceAggregatedList for aggregated lists of Instance
	// resources.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type InstanceList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Instance resources.
	Items []*Instance `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#instanceList for
	// lists of Instance resources.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type InstanceMoveRequest struct {
	// DestinationZone: The URL of the destination zone to move the instance
	// to. This can be a full or partial URL. For example, the following are
	// all valid URLs to a zone:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone
	// -
	// projects/project/zones/zone
	// - zones/zone
	DestinationZone string `json:"destinationZone,omitempty"`

	// TargetInstance: The URL of the target instance to move. This can be a
	// full or partial URL. For example, the following are all valid URLs to
	// an instance:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/zones/zone/inst
	// ances/instance
	// - projects/project/zones/zone/instances/instance
	// -
	// zones/zone/instances/instance
	TargetInstance string `json:"targetInstance,omitempty"`
}

type InstanceProperties struct {
	// CanIpForward: Allows instances created based on this template to send
	// packets with source IP addresses other than their own and receive
	// packets with destination IP addresses other than their own. If these
	// instances will be used as an IP gateway or it will be set as the
	// next-hop in a Route resource, say true. If unsure, leave this set to
	// false.
	CanIpForward bool `json:"canIpForward,omitempty"`

	// Description: An optional textual description for the instances
	// created based on the instance template resource; provided by the
	// client when the template is created.
	Description string `json:"description,omitempty"`

	// Disks: Array of disks associated with instance created based on this
	// template.
	Disks []*AttachedDisk `json:"disks,omitempty"`

	// MachineType: Name of the machine type resource describing which
	// machine type to use to host the instances created based on this
	// template; provided by the client when the instance template is
	// created.
	MachineType string `json:"machineType,omitempty"`

	// Metadata: Metadata key/value pairs assigned to instances created
	// based on this template. Consists of custom metadata or predefined
	// keys; see Instance documentation for more information.
	Metadata *Metadata `json:"metadata,omitempty"`

	// NetworkInterfaces: Array of configurations for this interface. This
	// specifies how this interface is configured to interact with other
	// network services, such as connecting to the internet. Currently,
	// ONE_TO_ONE_NAT is the only access config supported. If there are no
	// accessConfigs specified, then this instances created based based on
	// this template will have no external internet access.
	NetworkInterfaces []*NetworkInterface `json:"networkInterfaces,omitempty"`

	// Scheduling: Scheduling options for the instances created based on
	// this template.
	Scheduling *Scheduling `json:"scheduling,omitempty"`

	// ServiceAccounts: A list of service accounts each with specified
	// scopes, for which access tokens are to be made available to the
	// instances created based on this template, through metadata queries.
	ServiceAccounts []*ServiceAccount `json:"serviceAccounts,omitempty"`

	// Tags: A list of tags to be applied to the instances created based on
	// this template used to identify valid sources or targets for network
	// firewalls. Provided by the client on instance creation. The tags can
	// be later modified by the setTags method. Each tag within the list
	// must comply with RFC1035.
	Tags *Tags `json:"tags,omitempty"`
}

type InstanceReference struct {
	Instance string `json:"instance,omitempty"`
}

type InstanceTemplate struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the instance template
	// resource; provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the instance template resource; provided by the client
	// when the resource is created. The name must be 1-63 characters long,
	// and comply with RFC1035
	Name string `json:"name,omitempty"`

	// Properties: The instance properties portion of this instance template
	// resource.
	Properties *InstanceProperties `json:"properties,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type InstanceTemplateList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of InstanceTemplate resources.
	Items []*InstanceTemplate `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type InstancesScopedList struct {
	// Instances: [Output Only] List of instances contained in this scope.
	Instances []*Instance `json:"instances,omitempty"`

	// Warning: [Output Only] Informational warning which replaces the list
	// of instances when the list is empty.
	Warning *InstancesScopedListWarning `json:"warning,omitempty"`
}

type InstancesScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*InstancesScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type InstancesScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type License struct {
	// ChargesUseFee: If true, the customer will be charged license fee for
	// running software that contains this license on an instance.
	ChargesUseFee bool `json:"chargesUseFee,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#license for
	// licenses.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource. The name must be 1-63 characters long,
	// and comply with RCF1035.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type MachineType struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Deprecated: [Output Only] The deprecation status associated with this
	// machine type.
	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	// Description: [Output Only] An optional textual description of the
	// resource.
	Description string `json:"description,omitempty"`

	// GuestCpus: [Output Only] The tumber of CPUs exposed to the instance.
	GuestCpus int64 `json:"guestCpus,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// ImageSpaceGb: [Deprecated] This property is deprecated and will never
	// be populated with any relevant values.
	ImageSpaceGb int64 `json:"imageSpaceGb,omitempty"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// MaximumPersistentDisks: [Output Only] Maximum persistent disks
	// allowed.
	MaximumPersistentDisks int64 `json:"maximumPersistentDisks,omitempty"`

	// MaximumPersistentDisksSizeGb: [Output Only] Maximum total persistent
	// disks size (GB) allowed.
	MaximumPersistentDisksSizeGb int64 `json:"maximumPersistentDisksSizeGb,omitempty,string"`

	// MemoryMb: [Output Only] The amount of physical memory available to
	// the instance, defined in MB.
	MemoryMb int64 `json:"memoryMb,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// ScratchDisks: [Output Only] List of extended scratch disks assigned
	// to the instance.
	ScratchDisks []*MachineTypeScratchDisks `json:"scratchDisks,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Zone: [Output Only] The name of the zone where the machine type
	// resides, such as us-central1-a.
	Zone string `json:"zone,omitempty"`
}

type MachineTypeScratchDisks struct {
	// DiskGb: Size of the scratch disk, defined in GB.
	DiskGb int64 `json:"diskGb,omitempty"`
}

type MachineTypeAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped machine type lists.
	Items map[string]MachineTypesScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// compute#machineTypeAggregatedList for aggregated lists of machine
	// types.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type MachineTypeList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Machine Type resources.
	Items []*MachineType `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#machineTypeList
	// for lists of machine types.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type MachineTypesScopedList struct {
	// MachineTypes: [Output Only] List of machine types contained in this
	// scope.
	MachineTypes []*MachineType `json:"machineTypes,omitempty"`

	// Warning: [Output Only] An informational warning that appears when the
	// machine types list is empty.
	Warning *MachineTypesScopedListWarning `json:"warning,omitempty"`
}

type MachineTypesScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*MachineTypesScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type MachineTypesScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type Metadata struct {
	// Fingerprint: Specifies a fingerprint for this request, which is
	// essentially a hash of the metadata's contents and used for optimistic
	// locking. The fingerprint is initially generated by Compute Engine and
	// changes after every request to modify or update metadata. You must
	// always provide an up-to-date fingerprint hash in order to update or
	// change metadata.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Items: Array of key/value pairs. The total size of all keys and
	// values must be less than 512 KB.
	Items []*MetadataItems `json:"items,omitempty"`

	// Kind: [Output Only] Type of the resource. Always compute#metadata for
	// metadata.
	Kind string `json:"kind,omitempty"`
}

type MetadataItems struct {
	// Key: Key for the metadata entry. Keys must conform to the following
	// regexp: [a-zA-Z0-9-_]+, and be less than 128 bytes in length. This is
	// reflected as part of a URL in the metadata server. Additionally, to
	// avoid ambiguity, keys must not conflict with any other metadata keys
	// for the project.
	Key string `json:"key,omitempty"`

	// Value: Value for the metadata entry. These are free-form strings, and
	// only have meaning as interpreted by the image running in the
	// instance. The only restriction placed on values is that their size
	// must be less than or equal to 32768 bytes.
	Value string `json:"value,omitempty"`
}

type Network struct {
	// IPv4Range: The range of internal addresses that are legal on this
	// network. This range is a CIDR specification, for example:
	// 192.168.0.0/16. Provided by the client when the network is created.
	IPv4Range string `json:"IPv4Range,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// GatewayIPv4: A gateway address for default routing to other networks.
	// This value is read only and is selected by the Google Compute Engine,
	// typically as the first usable address in the IPv4Range.
	GatewayIPv4 string `json:"gatewayIPv4,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#network for
	// networks.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type NetworkInterface struct {
	// AccessConfigs: An array of configurations for this interface.
	// Currently, <codeONE_TO_ONE_NAT is the only access config supported.
	// If there are no accessConfigs specified, then this instance will have
	// no external internet access.
	AccessConfigs []*AccessConfig `json:"accessConfigs,omitempty"`

	// Name: [Output Only] The name of the network interface, generated by
	// the server. For network devices, these are eth0, eth1, etc.
	Name string `json:"name,omitempty"`

	// Network: URL of the network resource for this instance. This is
	// required for creating an instance but optional when creating a
	// firewall rule. If not specified when creating a firewall rule, the
	// default network is used:
	//
	// global/networks/default
	//
	// If you specify
	// this property, you can specify the network as a full or partial URL.
	// For example, the following are all valid URLs:
	// -
	// https://www.googleapis.com/compute/v1/projects/project/global/networks
	// /network
	// - projects/project/global/networks/network
	// -
	// global/networks/default
	Network string `json:"network,omitempty"`

	// NetworkIP: [Output Only] An optional IPV4 internal network address
	// assigned to the instance for this network interface.
	NetworkIP string `json:"networkIP,omitempty"`
}

type NetworkList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Network resources.
	Items []*Network `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#networkList for
	// lists of networks.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource .
	SelfLink string `json:"selfLink,omitempty"`
}

type Operation struct {
	// ClientOperationId: [Output Only] An optional identifier specified by
	// the client when the mutation was initiated. Must be unique for all
	// operation resources in the project
	ClientOperationId string `json:"clientOperationId,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// EndTime: [Output Only] The time that this operation was completed.
	// This is in RFC3339 text format.
	EndTime string `json:"endTime,omitempty"`

	// Error: [Output Only] If errors are generated during processing of the
	// operation, this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	// HttpErrorMessage: [Output Only] If the operation fails, this field
	// contains the HTTP error message that was returned, such as NOT FOUND.
	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	// HttpErrorStatusCode: [Output Only] If the operation fails, this field
	// contains the HTTP error message that was returned, such as 404.
	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: [Output Only] The time that this operation was requested.
	// This is in RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always compute#Operation
	// for Operation resources.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// OperationType: [Output Only] Type of the operation, such as insert,
	// update, and delete.
	OperationType string `json:"operationType,omitempty"`

	// Progress: [Output Only] An optional progress indicator that ranges
	// from 0 to 100. There is no requirement that this be linear or support
	// any granularity of operations. This should not be used to guess at
	// when the operation will be complete. This number should be
	// monotonically increasing as the operation progresses.
	Progress int64 `json:"progress,omitempty"`

	// Region: [Output Only] URL of the region where the operation resides.
	// Only applicable for regional resources.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: [Output Only] The time that this operation was started by
	// the server. This is in RFC3339 text format.
	StartTime string `json:"startTime,omitempty"`

	// Status: [Output Only] Status of the operation. Can be one of the
	// following: PENDING, RUNNING, or DONE.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: [Output Only] Unique target ID which identifies a
	// particular incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: [Output Only] URL of the resource the operation is
	// mutating.
	TargetLink string `json:"targetLink,omitempty"`

	// User: [Output Only] User who requested the operation, for example:
	// user@example.com.
	User string `json:"user,omitempty"`

	// Warnings: [Output Only] If warning messages are generated during
	// processing of the operation, this field will be populated.
	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	// Zone: [Output Only] URL of the zone where the operation resides.
	Zone string `json:"zone,omitempty"`
}

type OperationError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request which
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type OperationAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped operation lists.
	Items map[string]OperationsScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// compute#operationAggregatedList for aggregated lists of operations.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type OperationList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] The operation resources.
	Items []*Operation `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#operations for
	// Operations resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncate.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type OperationsScopedList struct {
	// Operations: [Output Only] List of operations contained in this scope.
	Operations []*Operation `json:"operations,omitempty"`

	// Warning: [Output Only] Informational warning which replaces the list
	// of operations when the list is empty.
	Warning *OperationsScopedListWarning `json:"warning,omitempty"`
}

type OperationsScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*OperationsScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type OperationsScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type PathMatcher struct {
	// DefaultService: The URL to the BackendService resource. This will be
	// used if none of the 'pathRules' defined by this PathMatcher is met by
	// the URL's path portion.
	DefaultService string `json:"defaultService,omitempty"`

	Description string `json:"description,omitempty"`

	// Name: The name to which this PathMatcher is referred by the HostRule.
	Name string `json:"name,omitempty"`

	// PathRules: The list of path rules.
	PathRules []*PathRule `json:"pathRules,omitempty"`
}

type PathRule struct {
	// Paths: The list of path patterns to match. Each must start with / and
	// the only place a * is allowed is at the end following a /. The string
	// fed to the path matcher does not include any text after the first ?
	// or #, and those chars are not allowed here.
	Paths []string `json:"paths,omitempty"`

	// Service: The URL of the BackendService resource if this rule is
	// matched.
	Service string `json:"service,omitempty"`
}

type Project struct {
	// CommonInstanceMetadata: Metadata key/value pairs available to all
	// instances contained in this project. See Custom metadata for more
	// information.
	CommonInstanceMetadata *Metadata `json:"commonInstanceMetadata,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#project for
	// projects.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource.
	Name string `json:"name,omitempty"`

	// Quotas: [Output Only] Quotas assigned to this project.
	Quotas []*Quota `json:"quotas,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// UsageExportLocation: The location in Cloud Storage and naming method
	// of the daily usage report.
	UsageExportLocation *UsageExportLocation `json:"usageExportLocation,omitempty"`
}

type Quota struct {
	// Limit: [Output Only] Quota limit for this metric.
	Limit float64 `json:"limit,omitempty"`

	// Metric: [Output Only] Name of the quota metric.
	Metric string `json:"metric,omitempty"`

	// Usage: [Output Only] Current usage of this metric.
	Usage float64 `json:"usage,omitempty"`
}

type Region struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Deprecated: [Output Only] The deprecation status associated with this
	// region.
	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	// Description: [Output Only] Textual description of the resource.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server .
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always compute#region for
	// regions.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// Quotas: [Output Only] Quotas assigned to this region.
	Quotas []*Quota `json:"quotas,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output Only] Status of the region, either UP or DOWN.
	Status string `json:"status,omitempty"`

	// Zones: [Output Only] A list of zones available in this region, in the
	// form of resource URLs.
	Zones []string `json:"zones,omitempty"`
}

type RegionList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Region resources.
	Items []*Region `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#regionList for
	// lists of regions.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type ResourceGroupReference struct {
	// Group: A URI referencing one of the resource views listed in the
	// backend service.
	Group string `json:"group,omitempty"`
}

type Route struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// DestRange: Which packets does this route apply to?
	DestRange string `json:"destRange,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// Network: URL of the network to which this route is applied; provided
	// by the client when the route is created.
	Network string `json:"network,omitempty"`

	// NextHopGateway: The URL to a gateway that should handle matching
	// packets.
	NextHopGateway string `json:"nextHopGateway,omitempty"`

	// NextHopInstance: The URL to an instance that should handle matching
	// packets.
	NextHopInstance string `json:"nextHopInstance,omitempty"`

	// NextHopIp: The network IP address of an instance that should handle
	// matching packets.
	NextHopIp string `json:"nextHopIp,omitempty"`

	// NextHopNetwork: The URL of the local network if it should handle
	// matching packets.
	NextHopNetwork string `json:"nextHopNetwork,omitempty"`

	// NextHopVpnTunnel: The URL to a VpnTunnel that should handle matching
	// packets.
	NextHopVpnTunnel string `json:"nextHopVpnTunnel,omitempty"`

	// Priority: Breaks ties between Routes of equal specificity. Routes
	// with smaller values win when tied with routes with larger values.
	Priority int64 `json:"priority,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// Tags: A list of instance tags to which this route applies.
	Tags []string `json:"tags,omitempty"`

	// Warnings: If potential misconfigurations are detected for this route,
	// this field will be populated with warning messages.
	Warnings []*RouteWarnings `json:"warnings,omitempty"`
}

type RouteWarnings struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*RouteWarningsData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type RouteWarningsData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type RouteList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of Route resources.
	Items []*Route `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type Scheduling struct {
	// AutomaticRestart: Specifies whether the instance should be
	// automatically restarted if it is terminated by Compute Engine (not
	// terminated by a user).
	AutomaticRestart bool `json:"automaticRestart,omitempty"`

	// OnHostMaintenance: Defines the maintenance behavior for this
	// instance. The default behavior is MIGRATE. For more information, see
	// Setting maintenance behavior.
	OnHostMaintenance string `json:"onHostMaintenance,omitempty"`
}

type SerialPortOutput struct {
	// Contents: [Output Only] The contents of the console output.
	Contents string `json:"contents,omitempty"`

	// Kind: [Output Only] Type of the resource. Always
	// compute#serialPortOutput for serial port output.
	Kind string `json:"kind,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type ServiceAccount struct {
	// Email: Email address of the service account.
	Email string `json:"email,omitempty"`

	// Scopes: The list of scopes to be made available for this service
	// account.
	Scopes []string `json:"scopes,omitempty"`
}

type Snapshot struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// DiskSizeGb: Size of the persistent disk snapshot, specified in GB
	// (output only).
	DiskSizeGb int64 `json:"diskSizeGb,omitempty,string"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Licenses: Public visible licenses.
	Licenses []string `json:"licenses,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// SourceDisk: The source disk used to create this snapshot.
	SourceDisk string `json:"sourceDisk,omitempty"`

	// SourceDiskId: The 'id' value of the disk used to create this
	// snapshot. This value may be used to determine whether the snapshot
	// was taken from the current or a previous instance of a given disk
	// name.
	SourceDiskId string `json:"sourceDiskId,omitempty"`

	// Status: The status of the persistent disk snapshot (output only).
	Status string `json:"status,omitempty"`

	// StorageBytes: A size of the the storage used by the snapshot. As
	// snapshots share storage this number is expected to change with
	// snapshot creation/deletion.
	StorageBytes int64 `json:"storageBytes,omitempty,string"`

	// StorageBytesStatus: An indicator whether storageBytes is in a stable
	// state, or it is being adjusted as a result of shared storage
	// reallocation.
	StorageBytesStatus string `json:"storageBytesStatus,omitempty"`
}

type SnapshotList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of Snapshot resources.
	Items []*Snapshot `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type Tags struct {
	// Fingerprint: Specifies a fingerprint for this request, which is
	// essentially a hash of the metadata's contents and used for optimistic
	// locking. The fingerprint is initially generated by Compute Engine and
	// changes after every request to modify or update metadata. You must
	// always provide an up-to-date fingerprint hash in order to update or
	// change metadata.
	//
	// To see the latest fingerprint, make get() request
	// to the instance.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Items: An array of tags. Each tag must be 1-63 characters long, and
	// comply with RFC1035.
	Items []string `json:"items,omitempty"`
}

type TargetHttpProxy struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// UrlMap: URL to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap string `json:"urlMap,omitempty"`
}

type TargetHttpProxyList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of TargetHttpProxy resources.
	Items []*TargetHttpProxy `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetInstance struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Instance: The URL to the instance that terminates the relevant
	// traffic.
	Instance string `json:"instance,omitempty"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// NatPolicy: NAT option controlling how IPs are NAT'ed to the VM.
	// Currently only NO_NAT (default value) is supported.
	NatPolicy string `json:"natPolicy,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// Zone: URL of the zone where the target instance resides (output
	// only).
	Zone string `json:"zone,omitempty"`
}

type TargetInstanceAggregatedList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A map of scoped target instance lists.
	Items map[string]TargetInstancesScopedList `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetInstanceList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of TargetInstance resources.
	Items []*TargetInstance `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetInstancesScopedList struct {
	// TargetInstances: List of target instances contained in this scope.
	TargetInstances []*TargetInstance `json:"targetInstances,omitempty"`

	// Warning: Informational warning which replaces the list of addresses
	// when the list is empty.
	Warning *TargetInstancesScopedListWarning `json:"warning,omitempty"`
}

type TargetInstancesScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*TargetInstancesScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type TargetInstancesScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type TargetPool struct {
	// BackupPool: This field is applicable only when the containing target
	// pool is serving a forwarding rule as the primary pool, and its
	// 'failoverRatio' field is properly set to a value between [0,
	// 1].
	//
	// 'backupPool' and 'failoverRatio' together define the fallback
	// behavior of the primary target pool: if the ratio of the healthy VMs
	// in the primary pool is at or below 'failoverRatio', traffic arriving
	// at the load-balanced IP will be directed to the backup pool.
	//
	// In case
	// where 'failoverRatio' and 'backupPool' are not set, or all the VMs in
	// the backup pool are unhealthy, the traffic will be directed back to
	// the primary pool in the "force" mode, where traffic will be spread to
	// the healthy VMs with the best effort, or to all VMs when no VM is
	// healthy.
	BackupPool string `json:"backupPool,omitempty"`

	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// FailoverRatio: This field is applicable only when the containing
	// target pool is serving a forwarding rule as the primary pool (i.e.,
	// not as a backup pool to some other target pool). The value of the
	// field must be in [0, 1].
	//
	// If set, 'backupPool' must also be set. They
	// together define the fallback behavior of the primary target pool: if
	// the ratio of the healthy VMs in the primary pool is at or below this
	// number, traffic arriving at the load-balanced IP will be directed to
	// the backup pool.
	//
	// In case where 'failoverRatio' is not set or all the
	// VMs in the backup pool are unhealthy, the traffic will be directed
	// back to the primary pool in the "force" mode, where traffic will be
	// spread to the healthy VMs with the best effort, or to all VMs when no
	// VM is healthy.
	FailoverRatio float64 `json:"failoverRatio,omitempty"`

	// HealthChecks: A list of URLs to the HttpHealthCheck resource. A
	// member VM in this pool is considered healthy if and only if all
	// specified health checks pass. An empty list means all member VMs will
	// be considered healthy at all times.
	HealthChecks []string `json:"healthChecks,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Instances: A list of resource URLs to the member VMs serving this
	// pool. They must live in zones contained in the same region as this
	// pool.
	Instances []string `json:"instances,omitempty"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// Region: URL of the region where the target pool resides (output
	// only).
	Region string `json:"region,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// SessionAffinity: Sesssion affinity option, must be one of the
	// following values: 'NONE': Connections from the same client IP may go
	// to any VM in the pool; 'CLIENT_IP': Connections from the same client
	// IP will go to the same VM in the pool while that VM remains healthy.
	// 'CLIENT_IP_PROTO': Connections from the same client IP with the same
	// IP protocol will go to the same VM in the pool while that VM remains
	// healthy.
	SessionAffinity string `json:"sessionAffinity,omitempty"`
}

type TargetPoolAggregatedList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A map of scoped target pool lists.
	Items map[string]TargetPoolsScopedList `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetPoolInstanceHealth struct {
	HealthStatus []*HealthStatus `json:"healthStatus,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`
}

type TargetPoolList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of TargetPool resources.
	Items []*TargetPool `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetPoolsAddHealthCheckRequest struct {
	// HealthChecks: Health check URLs to be added to targetPool.
	HealthChecks []*HealthCheckReference `json:"healthChecks,omitempty"`
}

type TargetPoolsAddInstanceRequest struct {
	// Instances: URLs of the instances to be added to targetPool.
	Instances []*InstanceReference `json:"instances,omitempty"`
}

type TargetPoolsRemoveHealthCheckRequest struct {
	// HealthChecks: Health check URLs to be removed from targetPool.
	HealthChecks []*HealthCheckReference `json:"healthChecks,omitempty"`
}

type TargetPoolsRemoveInstanceRequest struct {
	// Instances: URLs of the instances to be removed from targetPool.
	Instances []*InstanceReference `json:"instances,omitempty"`
}

type TargetPoolsScopedList struct {
	// TargetPools: List of target pools contained in this scope.
	TargetPools []*TargetPool `json:"targetPools,omitempty"`

	// Warning: Informational warning which replaces the list of addresses
	// when the list is empty.
	Warning *TargetPoolsScopedListWarning `json:"warning,omitempty"`
}

type TargetPoolsScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*TargetPoolsScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type TargetPoolsScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type TargetReference struct {
	Target string `json:"target,omitempty"`
}

type TargetVpnGateway struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource.
	// Provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// ForwardingRules: [Output Only] A list of URLs to the ForwardingRule
	// resources. ForwardingRules are created using
	// compute.forwardingRules.insert and associated to a VPN gateway.
	ForwardingRules []string `json:"forwardingRules,omitempty"`

	// Id: [Output Only] Unique identifier for the resource. Defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of resource. Always compute#targetVpnGateway
	// for target VPN gateways.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// Network: URL of the network to which this VPN gateway is attached.
	// Provided by the client when the VPN gateway is created.
	Network string `json:"network,omitempty"`

	// Region: [Output Only] URL of the region where the target VPN gateway
	// resides.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output Only] The status of the VPN gateway.
	Status string `json:"status,omitempty"`

	// Tunnels: [Output Only] A list of URLs to VpnTunnel resources.
	// VpnTunnels are created using compute.vpntunnels.insert and associated
	// to a VPN gateway.
	Tunnels []string `json:"tunnels,omitempty"`
}

type TargetVpnGatewayAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource. Defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: A map of scoped target vpn gateway lists.
	Items map[string]TargetVpnGatewaysScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#targetVpnGateway
	// for target VPN gateways.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetVpnGatewayList struct {
	// Id: [Output Only] Unique identifier for the resource. Defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of TargetVpnGateway resources.
	Items []*TargetVpnGateway `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#targetVpnGateway
	// for target VPN gateways.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type TargetVpnGatewaysScopedList struct {
	// TargetVpnGateways: [Output Only] List of target vpn gateways
	// contained in this scope.
	TargetVpnGateways []*TargetVpnGateway `json:"targetVpnGateways,omitempty"`

	// Warning: [Output Only] Informational warning which replaces the list
	// of addresses when the list is empty.
	Warning *TargetVpnGatewaysScopedListWarning `json:"warning,omitempty"`
}

type TargetVpnGatewaysScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*TargetVpnGatewaysScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type TargetVpnGatewaysScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type TestFailure struct {
	ActualService string `json:"actualService,omitempty"`

	ExpectedService string `json:"expectedService,omitempty"`

	Host string `json:"host,omitempty"`

	Path string `json:"path,omitempty"`
}

type UrlMap struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text format (output
	// only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// DefaultService: The URL of the BackendService resource if none of the
	// hostRules match.
	DefaultService string `json:"defaultService,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Fingerprint: Fingerprint of this resource. A hash of the contents
	// stored in this object. This field is used in optimistic locking. This
	// field will be ignored when inserting a UrlMap. An up-to-date
	// fingerprint must be provided in order to update the UrlMap.
	Fingerprint string `json:"fingerprint,omitempty"`

	// HostRules: The list of HostRules to use against the URL.
	HostRules []*HostRule `json:"hostRules,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// PathMatchers: The list of named PathMatchers to use against the URL.
	PathMatchers []*PathMatcher `json:"pathMatchers,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	// Tests: The list of expected URL mappings. Request to update this
	// UrlMap will succeed only all of the test cases pass.
	Tests []*UrlMapTest `json:"tests,omitempty"`
}

type UrlMapList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of UrlMap resources.
	Items []*UrlMap `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type UrlMapReference struct {
	UrlMap string `json:"urlMap,omitempty"`
}

type UrlMapTest struct {
	// Description: Description of this test case.
	Description string `json:"description,omitempty"`

	// Host: Host portion of the URL.
	Host string `json:"host,omitempty"`

	// Path: Path portion of the URL.
	Path string `json:"path,omitempty"`

	// Service: Expected BackendService resource the given URL should be
	// mapped to.
	Service string `json:"service,omitempty"`
}

type UrlMapValidationResult struct {
	LoadErrors []string `json:"loadErrors,omitempty"`

	// LoadSucceeded: Whether the given UrlMap can be successfully loaded.
	// If false, 'loadErrors' indicates the reasons.
	LoadSucceeded bool `json:"loadSucceeded,omitempty"`

	TestFailures []*TestFailure `json:"testFailures,omitempty"`

	// TestPassed: If successfully loaded, this field indicates whether the
	// test passed. If false, 'testFailures's indicate the reason of
	// failure.
	TestPassed bool `json:"testPassed,omitempty"`
}

type UrlMapsValidateRequest struct {
	// Resource: Content of the UrlMap to be validated.
	Resource *UrlMap `json:"resource,omitempty"`
}

type UrlMapsValidateResponse struct {
	Result *UrlMapValidationResult `json:"result,omitempty"`
}

type UsageExportLocation struct {
	// BucketName: The name of an existing bucket in Cloud Storage where the
	// usage report object is stored. The Google Service Account is granted
	// write access to this bucket. This is just the bucket name, with no
	// gs:// or https://storage.googleapis.com/ in front of it.
	BucketName string `json:"bucketName,omitempty"`

	// ReportNamePrefix: An optional prefix for the name of the usage report
	// object stored in bucketName. If not supplied, defaults to usage. The
	// report is stored as a CSV file named
	// report_name_prefix_gce_YYYYMMDD.csv where YYYYMMDD is the day of the
	// usage according to Pacific Time. If you supply a prefix, it should
	// conform to Cloud Storage object naming conventions.
	ReportNamePrefix string `json:"reportNamePrefix,omitempty"`
}

type VpnTunnel struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource.
	// Provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// DetailedStatus: [Output Only] Detailed status message for the VPN
	// tunnel.
	DetailedStatus string `json:"detailedStatus,omitempty"`

	// Id: [Output Only] Unique identifier for the resource. Defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// IkeNetworks: IKE networks to use when establishing the VPN tunnel
	// with peer VPN gateway. The value should be a CIDR formatted string,
	// for example: 192.168.0.0/16. The ranges should be disjoint.
	IkeNetworks []string `json:"ikeNetworks,omitempty"`

	// IkeVersion: IKE protocol version to use when establishing the VPN
	// tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2.
	// Default version is 2.
	IkeVersion int64 `json:"ikeVersion,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#vpnTunnel for
	// VPN tunnels.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long and comply with
	// RFC1035.
	Name string `json:"name,omitempty"`

	// PeerIp: IP address of the peer VPN gateway.
	PeerIp string `json:"peerIp,omitempty"`

	// Region: [Output Only] URL of the region where the VPN tunnel resides.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// SharedSecret: Shared secret used to set the secure session between
	// the GCE VPN gateway and the peer VPN gateway.
	SharedSecret string `json:"sharedSecret,omitempty"`

	// SharedSecretHash: Hash of the shared secret.
	SharedSecretHash string `json:"sharedSecretHash,omitempty"`

	// Status: [Output Only] The status of the VPN tunnel.
	Status string `json:"status,omitempty"`

	// TargetVpnGateway: URL of the VPN gateway to which this VPN tunnel is
	// associated. Provided by the client when the VPN tunnel is created.
	TargetVpnGateway string `json:"targetVpnGateway,omitempty"`
}

type VpnTunnelAggregatedList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A map of scoped vpn tunnel lists.
	Items map[string]VpnTunnelsScopedList `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#vpnTunnel for
	// VPN tunnels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type VpnTunnelList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of VpnTunnel resources.
	Items []*VpnTunnel `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#vpnTunnel for
	// VPN tunnels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type VpnTunnelsScopedList struct {
	// VpnTunnels: List of vpn tunnels contained in this scope.
	VpnTunnels []*VpnTunnel `json:"vpnTunnels,omitempty"`

	// Warning: Informational warning which replaces the list of addresses
	// when the list is empty.
	Warning *VpnTunnelsScopedListWarning `json:"warning,omitempty"`
}

type VpnTunnelsScopedListWarning struct {
	// Code: [Output Only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*VpnTunnelsScopedListWarningData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type VpnTunnelsScopedListWarningData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

type Zone struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Deprecated: [Output Only] The deprecation status associated with this
	// zone.
	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	// Description: [Output Only] Textual description of the resource.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always kind#zone for zones.
	Kind string `json:"kind,omitempty"`

	// MaintenanceWindows: [Output Only] Any scheduled maintenance windows
	// for this zone. When the zone is in a maintenance window, all
	// resources which reside in the zone will be unavailable. For more
	// information, see Maintenance Windows
	MaintenanceWindows []*ZoneMaintenanceWindows `json:"maintenanceWindows,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// Region: [Output Only] Full URL reference to the region which hosts
	// the zone.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output Only] Status of the zone, either UP or DOWN.
	Status string `json:"status,omitempty"`
}

type ZoneMaintenanceWindows struct {
	// BeginTime: [Output Only] Starting time of the maintenance window, in
	// RFC3339 format.
	BeginTime string `json:"beginTime,omitempty"`

	// Description: [Output Only] Textual description of the maintenance
	// window.
	Description string `json:"description,omitempty"`

	// EndTime: [Output Only] Ending time of the maintenance window, in
	// RFC3339 format.
	EndTime string `json:"endTime,omitempty"`

	// Name: [Output Only] Name of the maintenance window.
	Name string `json:"name,omitempty"`
}

type ZoneList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Zone resources.
	Items []*Zone `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

// method id "compute.addresses.aggregatedList":

type AddressesAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of addresses grouped by scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/addresses/aggregatedList
func (r *AddressesService) AggregatedList(project string) *AddressesAggregatedListCall {
	return &AddressesAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/addresses",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *AddressesAggregatedListCall) Filter(filter string) *AddressesAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *AddressesAggregatedListCall) MaxResults(maxResults int64) *AddressesAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *AddressesAggregatedListCall) PageToken(pageToken string) *AddressesAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *AddressesAggregatedListCall) Context(ctx context.Context) *AddressesAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AddressesAggregatedListCall) Fields(s ...googleapi.Field) *AddressesAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AddressesAggregatedListCall) Do() (*AddressAggregatedList, error) {
	var returnValue *AddressAggregatedList
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
	//   "description": "Retrieves the list of addresses grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.addresses.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/addresses",
	//   "response": {
	//     "$ref": "AddressAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.addresses.delete":

type AddressesDeleteCall struct {
	s             *Service
	project       string
	region        string
	address       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified address resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/addresses/delete
func (r *AddressesService) Delete(project string, region string, address string) *AddressesDeleteCall {
	return &AddressesDeleteCall{
		s:             r.s,
		project:       project,
		region:        region,
		address:       address,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/addresses/{address}",
		context_:      googleapi.NoContext,
	}
}
func (c *AddressesDeleteCall) Context(ctx context.Context) *AddressesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AddressesDeleteCall) Fields(s ...googleapi.Field) *AddressesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AddressesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
		"address": c.address,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified address resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.addresses.delete",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "address"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Name of the address resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/addresses/{address}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.addresses.get":

type AddressesGetCall struct {
	s             *Service
	project       string
	region        string
	address       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified address resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/addresses/get
func (r *AddressesService) Get(project string, region string, address string) *AddressesGetCall {
	return &AddressesGetCall{
		s:             r.s,
		project:       project,
		region:        region,
		address:       address,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/addresses/{address}",
		context_:      googleapi.NoContext,
	}
}
func (c *AddressesGetCall) Context(ctx context.Context) *AddressesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AddressesGetCall) Fields(s ...googleapi.Field) *AddressesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AddressesGetCall) Do() (*Address, error) {
	var returnValue *Address
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
		"address": c.address,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified address resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.addresses.get",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "address"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Name of the address resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/addresses/{address}",
	//   "response": {
	//     "$ref": "Address"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.addresses.insert":

type AddressesInsertCall struct {
	s             *Service
	project       string
	region        string
	address       *Address
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates an address resource in the specified project using
// the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/addresses/insert
func (r *AddressesService) Insert(project string, region string, address *Address) *AddressesInsertCall {
	return &AddressesInsertCall{
		s:             r.s,
		project:       project,
		region:        region,
		address:       address,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/addresses",
		context_:      googleapi.NoContext,
	}
}
func (c *AddressesInsertCall) Context(ctx context.Context) *AddressesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AddressesInsertCall) Fields(s ...googleapi.Field) *AddressesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AddressesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.address,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an address resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.addresses.insert",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/addresses",
	//   "request": {
	//     "$ref": "Address"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.addresses.list":

type AddressesListCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of address resources contained within the
// specified region.

// For details, see https://cloud.google.com/compute/docs/reference/latest/addresses/list
func (r *AddressesService) List(project string, region string) *AddressesListCall {
	return &AddressesListCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/addresses",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *AddressesListCall) Filter(filter string) *AddressesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *AddressesListCall) MaxResults(maxResults int64) *AddressesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *AddressesListCall) PageToken(pageToken string) *AddressesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *AddressesListCall) Context(ctx context.Context) *AddressesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AddressesListCall) Fields(s ...googleapi.Field) *AddressesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AddressesListCall) Do() (*AddressList, error) {
	var returnValue *AddressList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of address resources contained within the specified region.",
	//   "httpMethod": "GET",
	//   "id": "compute.addresses.list",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/addresses",
	//   "response": {
	//     "$ref": "AddressList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.backendServices.delete":

type BackendServicesDeleteCall struct {
	s              *Service
	project        string
	backendService string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Delete: Deletes the specified BackendService resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/delete
func (r *BackendServicesService) Delete(project string, backendService string) *BackendServicesDeleteCall {
	return &BackendServicesDeleteCall{
		s:              r.s,
		project:        project,
		backendService: backendService,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/backendServices/{backendService}",
		context_:       googleapi.NoContext,
	}
}
func (c *BackendServicesDeleteCall) Context(ctx context.Context) *BackendServicesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesDeleteCall) Fields(s ...googleapi.Field) *BackendServicesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"backendService": c.backendService,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified BackendService resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.backendServices.delete",
	//   "parameterOrder": [
	//     "project",
	//     "backendService"
	//   ],
	//   "parameters": {
	//     "backendService": {
	//       "description": "Name of the BackendService resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices/{backendService}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.backendServices.get":

type BackendServicesGetCall struct {
	s              *Service
	project        string
	backendService string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Get: Returns the specified BackendService resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/get
func (r *BackendServicesService) Get(project string, backendService string) *BackendServicesGetCall {
	return &BackendServicesGetCall{
		s:              r.s,
		project:        project,
		backendService: backendService,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/backendServices/{backendService}",
		context_:       googleapi.NoContext,
	}
}
func (c *BackendServicesGetCall) Context(ctx context.Context) *BackendServicesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesGetCall) Fields(s ...googleapi.Field) *BackendServicesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesGetCall) Do() (*BackendService, error) {
	var returnValue *BackendService
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"backendService": c.backendService,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified BackendService resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.backendServices.get",
	//   "parameterOrder": [
	//     "project",
	//     "backendService"
	//   ],
	//   "parameters": {
	//     "backendService": {
	//       "description": "Name of the BackendService resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices/{backendService}",
	//   "response": {
	//     "$ref": "BackendService"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.backendServices.getHealth":

type BackendServicesGetHealthCall struct {
	s                      *Service
	project                string
	backendService         string
	resourcegroupreference *ResourceGroupReference
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
	context_               context.Context
}

// GetHealth: Gets the most recent health check results for this
// BackendService.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/getHealth
func (r *BackendServicesService) GetHealth(project string, backendService string, resourcegroupreference *ResourceGroupReference) *BackendServicesGetHealthCall {
	return &BackendServicesGetHealthCall{
		s:                      r.s,
		project:                project,
		backendService:         backendService,
		resourcegroupreference: resourcegroupreference,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "{project}/global/backendServices/{backendService}/getHealth",
		context_:               googleapi.NoContext,
	}
}
func (c *BackendServicesGetHealthCall) Context(ctx context.Context) *BackendServicesGetHealthCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesGetHealthCall) Fields(s ...googleapi.Field) *BackendServicesGetHealthCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesGetHealthCall) Do() (*BackendServiceGroupHealth, error) {
	var returnValue *BackendServiceGroupHealth
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"backendService": c.backendService,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.resourcegroupreference,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the most recent health check results for this BackendService.",
	//   "httpMethod": "POST",
	//   "id": "compute.backendServices.getHealth",
	//   "parameterOrder": [
	//     "project",
	//     "backendService"
	//   ],
	//   "parameters": {
	//     "backendService": {
	//       "description": "Name of the BackendService resource to which the queried instance belongs.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices/{backendService}/getHealth",
	//   "request": {
	//     "$ref": "ResourceGroupReference"
	//   },
	//   "response": {
	//     "$ref": "BackendServiceGroupHealth"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.backendServices.insert":

type BackendServicesInsertCall struct {
	s              *Service
	project        string
	backendservice *BackendService
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Insert: Creates a BackendService resource in the specified project
// using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/insert
func (r *BackendServicesService) Insert(project string, backendservice *BackendService) *BackendServicesInsertCall {
	return &BackendServicesInsertCall{
		s:              r.s,
		project:        project,
		backendservice: backendservice,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/backendServices",
		context_:       googleapi.NoContext,
	}
}
func (c *BackendServicesInsertCall) Context(ctx context.Context) *BackendServicesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesInsertCall) Fields(s ...googleapi.Field) *BackendServicesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.backendservice,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a BackendService resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.backendServices.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices",
	//   "request": {
	//     "$ref": "BackendService"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.backendServices.list":

type BackendServicesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of BackendService resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/list
func (r *BackendServicesService) List(project string) *BackendServicesListCall {
	return &BackendServicesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/backendServices",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *BackendServicesListCall) Filter(filter string) *BackendServicesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *BackendServicesListCall) MaxResults(maxResults int64) *BackendServicesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *BackendServicesListCall) PageToken(pageToken string) *BackendServicesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *BackendServicesListCall) Context(ctx context.Context) *BackendServicesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesListCall) Fields(s ...googleapi.Field) *BackendServicesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesListCall) Do() (*BackendServiceList, error) {
	var returnValue *BackendServiceList
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
	//   "description": "Retrieves the list of BackendService resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.backendServices.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices",
	//   "response": {
	//     "$ref": "BackendServiceList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.backendServices.patch":

type BackendServicesPatchCall struct {
	s              *Service
	project        string
	backendService string
	backendservice *BackendService
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Patch: Update the entire content of the BackendService resource. This
// method supports patch semantics.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/patch
func (r *BackendServicesService) Patch(project string, backendService string, backendservice *BackendService) *BackendServicesPatchCall {
	return &BackendServicesPatchCall{
		s:              r.s,
		project:        project,
		backendService: backendService,
		backendservice: backendservice,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/backendServices/{backendService}",
		context_:       googleapi.NoContext,
	}
}
func (c *BackendServicesPatchCall) Context(ctx context.Context) *BackendServicesPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesPatchCall) Fields(s ...googleapi.Field) *BackendServicesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesPatchCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"backendService": c.backendService,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.backendservice,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Update the entire content of the BackendService resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "compute.backendServices.patch",
	//   "parameterOrder": [
	//     "project",
	//     "backendService"
	//   ],
	//   "parameters": {
	//     "backendService": {
	//       "description": "Name of the BackendService resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices/{backendService}",
	//   "request": {
	//     "$ref": "BackendService"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.backendServices.update":

type BackendServicesUpdateCall struct {
	s              *Service
	project        string
	backendService string
	backendservice *BackendService
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Update: Update the entire content of the BackendService resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/backendServices/update
func (r *BackendServicesService) Update(project string, backendService string, backendservice *BackendService) *BackendServicesUpdateCall {
	return &BackendServicesUpdateCall{
		s:              r.s,
		project:        project,
		backendService: backendService,
		backendservice: backendservice,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/backendServices/{backendService}",
		context_:       googleapi.NoContext,
	}
}
func (c *BackendServicesUpdateCall) Context(ctx context.Context) *BackendServicesUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackendServicesUpdateCall) Fields(s ...googleapi.Field) *BackendServicesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BackendServicesUpdateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"backendService": c.backendService,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.backendservice,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Update the entire content of the BackendService resource.",
	//   "httpMethod": "PUT",
	//   "id": "compute.backendServices.update",
	//   "parameterOrder": [
	//     "project",
	//     "backendService"
	//   ],
	//   "parameters": {
	//     "backendService": {
	//       "description": "Name of the BackendService resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/backendServices/{backendService}",
	//   "request": {
	//     "$ref": "BackendService"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.diskTypes.aggregatedList":

type DiskTypesAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of disk type resources grouped by
// scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/diskTypes/aggregatedList
func (r *DiskTypesService) AggregatedList(project string) *DiskTypesAggregatedListCall {
	return &DiskTypesAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/diskTypes",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *DiskTypesAggregatedListCall) Filter(filter string) *DiskTypesAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *DiskTypesAggregatedListCall) MaxResults(maxResults int64) *DiskTypesAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *DiskTypesAggregatedListCall) PageToken(pageToken string) *DiskTypesAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DiskTypesAggregatedListCall) Context(ctx context.Context) *DiskTypesAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DiskTypesAggregatedListCall) Fields(s ...googleapi.Field) *DiskTypesAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DiskTypesAggregatedListCall) Do() (*DiskTypeAggregatedList, error) {
	var returnValue *DiskTypeAggregatedList
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
	//   "description": "Retrieves the list of disk type resources grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.diskTypes.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/diskTypes",
	//   "response": {
	//     "$ref": "DiskTypeAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.diskTypes.get":

type DiskTypesGetCall struct {
	s             *Service
	project       string
	zone          string
	diskType      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified disk type resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/diskTypes/get
func (r *DiskTypesService) Get(project string, zone string, diskType string) *DiskTypesGetCall {
	return &DiskTypesGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		diskType:      diskType,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/diskTypes/{diskType}",
		context_:      googleapi.NoContext,
	}
}
func (c *DiskTypesGetCall) Context(ctx context.Context) *DiskTypesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DiskTypesGetCall) Fields(s ...googleapi.Field) *DiskTypesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DiskTypesGetCall) Do() (*DiskType, error) {
	var returnValue *DiskType
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"diskType": c.diskType,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified disk type resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.diskTypes.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "diskType"
	//   ],
	//   "parameters": {
	//     "diskType": {
	//       "description": "Name of the disk type resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/diskTypes/{diskType}",
	//   "response": {
	//     "$ref": "DiskType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.diskTypes.list":

type DiskTypesListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of disk type resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/diskTypes/list
func (r *DiskTypesService) List(project string, zone string) *DiskTypesListCall {
	return &DiskTypesListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/diskTypes",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *DiskTypesListCall) Filter(filter string) *DiskTypesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *DiskTypesListCall) MaxResults(maxResults int64) *DiskTypesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *DiskTypesListCall) PageToken(pageToken string) *DiskTypesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DiskTypesListCall) Context(ctx context.Context) *DiskTypesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DiskTypesListCall) Fields(s ...googleapi.Field) *DiskTypesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DiskTypesListCall) Do() (*DiskTypeList, error) {
	var returnValue *DiskTypeList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of disk type resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.diskTypes.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/diskTypes",
	//   "response": {
	//     "$ref": "DiskTypeList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.disks.aggregatedList":

type DisksAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of disks grouped by scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/disks/aggregatedList
func (r *DisksService) AggregatedList(project string) *DisksAggregatedListCall {
	return &DisksAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/disks",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *DisksAggregatedListCall) Filter(filter string) *DisksAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *DisksAggregatedListCall) MaxResults(maxResults int64) *DisksAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *DisksAggregatedListCall) PageToken(pageToken string) *DisksAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DisksAggregatedListCall) Context(ctx context.Context) *DisksAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DisksAggregatedListCall) Fields(s ...googleapi.Field) *DisksAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DisksAggregatedListCall) Do() (*DiskAggregatedList, error) {
	var returnValue *DiskAggregatedList
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
	//   "description": "Retrieves the list of disks grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.disks.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/disks",
	//   "response": {
	//     "$ref": "DiskAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.disks.createSnapshot":

type DisksCreateSnapshotCall struct {
	s             *Service
	project       string
	zone          string
	disk          string
	snapshot      *Snapshot
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// CreateSnapshot: Creates a snapshot of this disk.

// For details, see https://cloud.google.com/compute/docs/reference/latest/disks/createSnapshot
func (r *DisksService) CreateSnapshot(project string, zone string, disk string, snapshot *Snapshot) *DisksCreateSnapshotCall {
	return &DisksCreateSnapshotCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		disk:          disk,
		snapshot:      snapshot,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/disks/{disk}/createSnapshot",
		context_:      googleapi.NoContext,
	}
}
func (c *DisksCreateSnapshotCall) Context(ctx context.Context) *DisksCreateSnapshotCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DisksCreateSnapshotCall) Fields(s ...googleapi.Field) *DisksCreateSnapshotCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DisksCreateSnapshotCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"disk":    c.disk,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.snapshot,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a snapshot of this disk.",
	//   "httpMethod": "POST",
	//   "id": "compute.disks.createSnapshot",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "disk"
	//   ],
	//   "parameters": {
	//     "disk": {
	//       "description": "Name of the persistent disk to snapshot.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/disks/{disk}/createSnapshot",
	//   "request": {
	//     "$ref": "Snapshot"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.disks.delete":

type DisksDeleteCall struct {
	s             *Service
	project       string
	zone          string
	disk          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified persistent disk.

// For details, see https://cloud.google.com/compute/docs/reference/latest/disks/delete
func (r *DisksService) Delete(project string, zone string, disk string) *DisksDeleteCall {
	return &DisksDeleteCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		disk:          disk,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/disks/{disk}",
		context_:      googleapi.NoContext,
	}
}
func (c *DisksDeleteCall) Context(ctx context.Context) *DisksDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DisksDeleteCall) Fields(s ...googleapi.Field) *DisksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DisksDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"disk":    c.disk,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified persistent disk.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.disks.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "disk"
	//   ],
	//   "parameters": {
	//     "disk": {
	//       "description": "Name of the persistent disk to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/disks/{disk}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.disks.get":

type DisksGetCall struct {
	s             *Service
	project       string
	zone          string
	disk          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns a specified persistent disk.

// For details, see https://cloud.google.com/compute/docs/reference/latest/disks/get
func (r *DisksService) Get(project string, zone string, disk string) *DisksGetCall {
	return &DisksGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		disk:          disk,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/disks/{disk}",
		context_:      googleapi.NoContext,
	}
}
func (c *DisksGetCall) Context(ctx context.Context) *DisksGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DisksGetCall) Fields(s ...googleapi.Field) *DisksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DisksGetCall) Do() (*Disk, error) {
	var returnValue *Disk
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"disk":    c.disk,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns a specified persistent disk.",
	//   "httpMethod": "GET",
	//   "id": "compute.disks.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "disk"
	//   ],
	//   "parameters": {
	//     "disk": {
	//       "description": "Name of the persistent disk to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/disks/{disk}",
	//   "response": {
	//     "$ref": "Disk"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.disks.insert":

type DisksInsertCall struct {
	s             *Service
	project       string
	zone          string
	disk          *Disk
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a persistent disk in the specified project using the
// data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/disks/insert
func (r *DisksService) Insert(project string, zone string, disk *Disk) *DisksInsertCall {
	return &DisksInsertCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		disk:          disk,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/disks",
		context_:      googleapi.NoContext,
	}
}

// SourceImage sets the optional parameter "sourceImage": Source image
// to restore onto a disk.
func (c *DisksInsertCall) SourceImage(sourceImage string) *DisksInsertCall {
	c.params_.Set("sourceImage", fmt.Sprintf("%v", sourceImage))
	return c
}
func (c *DisksInsertCall) Context(ctx context.Context) *DisksInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DisksInsertCall) Fields(s ...googleapi.Field) *DisksInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DisksInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.disk,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a persistent disk in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.disks.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sourceImage": {
	//       "description": "Optional. Source image to restore onto a disk.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/disks",
	//   "request": {
	//     "$ref": "Disk"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.disks.list":

type DisksListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of persistent disks contained within the
// specified zone.

// For details, see https://cloud.google.com/compute/docs/reference/latest/disks/list
func (r *DisksService) List(project string, zone string) *DisksListCall {
	return &DisksListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/disks",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *DisksListCall) Filter(filter string) *DisksListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *DisksListCall) MaxResults(maxResults int64) *DisksListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *DisksListCall) PageToken(pageToken string) *DisksListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *DisksListCall) Context(ctx context.Context) *DisksListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DisksListCall) Fields(s ...googleapi.Field) *DisksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DisksListCall) Do() (*DiskList, error) {
	var returnValue *DiskList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of persistent disks contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "compute.disks.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/disks",
	//   "response": {
	//     "$ref": "DiskList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.firewalls.delete":

type FirewallsDeleteCall struct {
	s             *Service
	project       string
	firewall      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified firewall resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/firewalls/delete
func (r *FirewallsService) Delete(project string, firewall string) *FirewallsDeleteCall {
	return &FirewallsDeleteCall{
		s:             r.s,
		project:       project,
		firewall:      firewall,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/firewalls/{firewall}",
		context_:      googleapi.NoContext,
	}
}
func (c *FirewallsDeleteCall) Context(ctx context.Context) *FirewallsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FirewallsDeleteCall) Fields(s ...googleapi.Field) *FirewallsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FirewallsDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"firewall": c.firewall,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified firewall resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.firewalls.delete",
	//   "parameterOrder": [
	//     "project",
	//     "firewall"
	//   ],
	//   "parameters": {
	//     "firewall": {
	//       "description": "Name of the firewall resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/firewalls/{firewall}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.firewalls.get":

type FirewallsGetCall struct {
	s             *Service
	project       string
	firewall      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified firewall resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/firewalls/get
func (r *FirewallsService) Get(project string, firewall string) *FirewallsGetCall {
	return &FirewallsGetCall{
		s:             r.s,
		project:       project,
		firewall:      firewall,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/firewalls/{firewall}",
		context_:      googleapi.NoContext,
	}
}
func (c *FirewallsGetCall) Context(ctx context.Context) *FirewallsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FirewallsGetCall) Fields(s ...googleapi.Field) *FirewallsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FirewallsGetCall) Do() (*Firewall, error) {
	var returnValue *Firewall
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"firewall": c.firewall,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified firewall resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.firewalls.get",
	//   "parameterOrder": [
	//     "project",
	//     "firewall"
	//   ],
	//   "parameters": {
	//     "firewall": {
	//       "description": "Name of the firewall resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/firewalls/{firewall}",
	//   "response": {
	//     "$ref": "Firewall"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.firewalls.insert":

type FirewallsInsertCall struct {
	s             *Service
	project       string
	firewall      *Firewall
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a firewall resource in the specified project using
// the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/firewalls/insert
func (r *FirewallsService) Insert(project string, firewall *Firewall) *FirewallsInsertCall {
	return &FirewallsInsertCall{
		s:             r.s,
		project:       project,
		firewall:      firewall,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/firewalls",
		context_:      googleapi.NoContext,
	}
}
func (c *FirewallsInsertCall) Context(ctx context.Context) *FirewallsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FirewallsInsertCall) Fields(s ...googleapi.Field) *FirewallsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FirewallsInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.firewall,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a firewall resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.firewalls.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/firewalls",
	//   "request": {
	//     "$ref": "Firewall"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.firewalls.list":

type FirewallsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of firewall resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/firewalls/list
func (r *FirewallsService) List(project string) *FirewallsListCall {
	return &FirewallsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/firewalls",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *FirewallsListCall) Filter(filter string) *FirewallsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *FirewallsListCall) MaxResults(maxResults int64) *FirewallsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *FirewallsListCall) PageToken(pageToken string) *FirewallsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *FirewallsListCall) Context(ctx context.Context) *FirewallsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FirewallsListCall) Fields(s ...googleapi.Field) *FirewallsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FirewallsListCall) Do() (*FirewallList, error) {
	var returnValue *FirewallList
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
	//   "description": "Retrieves the list of firewall resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.firewalls.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/firewalls",
	//   "response": {
	//     "$ref": "FirewallList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.firewalls.patch":

type FirewallsPatchCall struct {
	s             *Service
	project       string
	firewall      string
	firewall2     *Firewall
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates the specified firewall resource with the data included
// in the request. This method supports patch semantics.

// For details, see https://cloud.google.com/compute/docs/reference/latest/firewalls/patch
func (r *FirewallsService) Patch(project string, firewall string, firewall2 *Firewall) *FirewallsPatchCall {
	return &FirewallsPatchCall{
		s:             r.s,
		project:       project,
		firewall:      firewall,
		firewall2:     firewall2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/firewalls/{firewall}",
		context_:      googleapi.NoContext,
	}
}
func (c *FirewallsPatchCall) Context(ctx context.Context) *FirewallsPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FirewallsPatchCall) Fields(s ...googleapi.Field) *FirewallsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FirewallsPatchCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"firewall": c.firewall,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.firewall2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the specified firewall resource with the data included in the request. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "compute.firewalls.patch",
	//   "parameterOrder": [
	//     "project",
	//     "firewall"
	//   ],
	//   "parameters": {
	//     "firewall": {
	//       "description": "Name of the firewall resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/firewalls/{firewall}",
	//   "request": {
	//     "$ref": "Firewall"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.firewalls.update":

type FirewallsUpdateCall struct {
	s             *Service
	project       string
	firewall      string
	firewall2     *Firewall
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates the specified firewall resource with the data
// included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/firewalls/update
func (r *FirewallsService) Update(project string, firewall string, firewall2 *Firewall) *FirewallsUpdateCall {
	return &FirewallsUpdateCall{
		s:             r.s,
		project:       project,
		firewall:      firewall,
		firewall2:     firewall2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/firewalls/{firewall}",
		context_:      googleapi.NoContext,
	}
}
func (c *FirewallsUpdateCall) Context(ctx context.Context) *FirewallsUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FirewallsUpdateCall) Fields(s ...googleapi.Field) *FirewallsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FirewallsUpdateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"firewall": c.firewall,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.firewall2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates the specified firewall resource with the data included in the request.",
	//   "httpMethod": "PUT",
	//   "id": "compute.firewalls.update",
	//   "parameterOrder": [
	//     "project",
	//     "firewall"
	//   ],
	//   "parameters": {
	//     "firewall": {
	//       "description": "Name of the firewall resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/firewalls/{firewall}",
	//   "request": {
	//     "$ref": "Firewall"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.forwardingRules.aggregatedList":

type ForwardingRulesAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of forwarding rules grouped by
// scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/forwardingRules/aggregatedList
func (r *ForwardingRulesService) AggregatedList(project string) *ForwardingRulesAggregatedListCall {
	return &ForwardingRulesAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/forwardingRules",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ForwardingRulesAggregatedListCall) Filter(filter string) *ForwardingRulesAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ForwardingRulesAggregatedListCall) MaxResults(maxResults int64) *ForwardingRulesAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *ForwardingRulesAggregatedListCall) PageToken(pageToken string) *ForwardingRulesAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ForwardingRulesAggregatedListCall) Context(ctx context.Context) *ForwardingRulesAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ForwardingRulesAggregatedListCall) Fields(s ...googleapi.Field) *ForwardingRulesAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ForwardingRulesAggregatedListCall) Do() (*ForwardingRuleAggregatedList, error) {
	var returnValue *ForwardingRuleAggregatedList
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
	//   "description": "Retrieves the list of forwarding rules grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.forwardingRules.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/forwardingRules",
	//   "response": {
	//     "$ref": "ForwardingRuleAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.forwardingRules.delete":

type ForwardingRulesDeleteCall struct {
	s              *Service
	project        string
	region         string
	forwardingRule string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Delete: Deletes the specified ForwardingRule resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/forwardingRules/delete
func (r *ForwardingRulesService) Delete(project string, region string, forwardingRule string) *ForwardingRulesDeleteCall {
	return &ForwardingRulesDeleteCall{
		s:              r.s,
		project:        project,
		region:         region,
		forwardingRule: forwardingRule,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/regions/{region}/forwardingRules/{forwardingRule}",
		context_:       googleapi.NoContext,
	}
}
func (c *ForwardingRulesDeleteCall) Context(ctx context.Context) *ForwardingRulesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ForwardingRulesDeleteCall) Fields(s ...googleapi.Field) *ForwardingRulesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ForwardingRulesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"region":         c.region,
		"forwardingRule": c.forwardingRule,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified ForwardingRule resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.forwardingRules.delete",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "forwardingRule"
	//   ],
	//   "parameters": {
	//     "forwardingRule": {
	//       "description": "Name of the ForwardingRule resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/forwardingRules/{forwardingRule}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.forwardingRules.get":

type ForwardingRulesGetCall struct {
	s              *Service
	project        string
	region         string
	forwardingRule string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Get: Returns the specified ForwardingRule resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/forwardingRules/get
func (r *ForwardingRulesService) Get(project string, region string, forwardingRule string) *ForwardingRulesGetCall {
	return &ForwardingRulesGetCall{
		s:              r.s,
		project:        project,
		region:         region,
		forwardingRule: forwardingRule,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/regions/{region}/forwardingRules/{forwardingRule}",
		context_:       googleapi.NoContext,
	}
}
func (c *ForwardingRulesGetCall) Context(ctx context.Context) *ForwardingRulesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ForwardingRulesGetCall) Fields(s ...googleapi.Field) *ForwardingRulesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ForwardingRulesGetCall) Do() (*ForwardingRule, error) {
	var returnValue *ForwardingRule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"region":         c.region,
		"forwardingRule": c.forwardingRule,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified ForwardingRule resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.forwardingRules.get",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "forwardingRule"
	//   ],
	//   "parameters": {
	//     "forwardingRule": {
	//       "description": "Name of the ForwardingRule resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/forwardingRules/{forwardingRule}",
	//   "response": {
	//     "$ref": "ForwardingRule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.forwardingRules.insert":

type ForwardingRulesInsertCall struct {
	s              *Service
	project        string
	region         string
	forwardingrule *ForwardingRule
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Insert: Creates a ForwardingRule resource in the specified project
// and region using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/forwardingRules/insert
func (r *ForwardingRulesService) Insert(project string, region string, forwardingrule *ForwardingRule) *ForwardingRulesInsertCall {
	return &ForwardingRulesInsertCall{
		s:              r.s,
		project:        project,
		region:         region,
		forwardingrule: forwardingrule,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/regions/{region}/forwardingRules",
		context_:       googleapi.NoContext,
	}
}
func (c *ForwardingRulesInsertCall) Context(ctx context.Context) *ForwardingRulesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ForwardingRulesInsertCall) Fields(s ...googleapi.Field) *ForwardingRulesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ForwardingRulesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.forwardingrule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a ForwardingRule resource in the specified project and region using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.forwardingRules.insert",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/forwardingRules",
	//   "request": {
	//     "$ref": "ForwardingRule"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.forwardingRules.list":

type ForwardingRulesListCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of ForwardingRule resources available to the
// specified project and region.

// For details, see https://cloud.google.com/compute/docs/reference/latest/forwardingRules/list
func (r *ForwardingRulesService) List(project string, region string) *ForwardingRulesListCall {
	return &ForwardingRulesListCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/forwardingRules",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ForwardingRulesListCall) Filter(filter string) *ForwardingRulesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ForwardingRulesListCall) MaxResults(maxResults int64) *ForwardingRulesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *ForwardingRulesListCall) PageToken(pageToken string) *ForwardingRulesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ForwardingRulesListCall) Context(ctx context.Context) *ForwardingRulesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ForwardingRulesListCall) Fields(s ...googleapi.Field) *ForwardingRulesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ForwardingRulesListCall) Do() (*ForwardingRuleList, error) {
	var returnValue *ForwardingRuleList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of ForwardingRule resources available to the specified project and region.",
	//   "httpMethod": "GET",
	//   "id": "compute.forwardingRules.list",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/forwardingRules",
	//   "response": {
	//     "$ref": "ForwardingRuleList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.forwardingRules.setTarget":

type ForwardingRulesSetTargetCall struct {
	s               *Service
	project         string
	region          string
	forwardingRule  string
	targetreference *TargetReference
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// SetTarget: Changes target url for forwarding rule.

// For details, see https://cloud.google.com/compute/docs/reference/latest/forwardingRules/setTarget
func (r *ForwardingRulesService) SetTarget(project string, region string, forwardingRule string, targetreference *TargetReference) *ForwardingRulesSetTargetCall {
	return &ForwardingRulesSetTargetCall{
		s:               r.s,
		project:         project,
		region:          region,
		forwardingRule:  forwardingRule,
		targetreference: targetreference,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/regions/{region}/forwardingRules/{forwardingRule}/setTarget",
		context_:        googleapi.NoContext,
	}
}
func (c *ForwardingRulesSetTargetCall) Context(ctx context.Context) *ForwardingRulesSetTargetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ForwardingRulesSetTargetCall) Fields(s ...googleapi.Field) *ForwardingRulesSetTargetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ForwardingRulesSetTargetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"region":         c.region,
		"forwardingRule": c.forwardingRule,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetreference,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Changes target url for forwarding rule.",
	//   "httpMethod": "POST",
	//   "id": "compute.forwardingRules.setTarget",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "forwardingRule"
	//   ],
	//   "parameters": {
	//     "forwardingRule": {
	//       "description": "Name of the ForwardingRule resource in which target is to be set.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/forwardingRules/{forwardingRule}/setTarget",
	//   "request": {
	//     "$ref": "TargetReference"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalAddresses.delete":

type GlobalAddressesDeleteCall struct {
	s             *Service
	project       string
	address       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified address resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalAddresses/delete
func (r *GlobalAddressesService) Delete(project string, address string) *GlobalAddressesDeleteCall {
	return &GlobalAddressesDeleteCall{
		s:             r.s,
		project:       project,
		address:       address,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/addresses/{address}",
		context_:      googleapi.NoContext,
	}
}
func (c *GlobalAddressesDeleteCall) Context(ctx context.Context) *GlobalAddressesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAddressesDeleteCall) Fields(s ...googleapi.Field) *GlobalAddressesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalAddressesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"address": c.address,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified address resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.globalAddresses.delete",
	//   "parameterOrder": [
	//     "project",
	//     "address"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Name of the address resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/addresses/{address}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalAddresses.get":

type GlobalAddressesGetCall struct {
	s             *Service
	project       string
	address       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified address resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalAddresses/get
func (r *GlobalAddressesService) Get(project string, address string) *GlobalAddressesGetCall {
	return &GlobalAddressesGetCall{
		s:             r.s,
		project:       project,
		address:       address,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/addresses/{address}",
		context_:      googleapi.NoContext,
	}
}
func (c *GlobalAddressesGetCall) Context(ctx context.Context) *GlobalAddressesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAddressesGetCall) Fields(s ...googleapi.Field) *GlobalAddressesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalAddressesGetCall) Do() (*Address, error) {
	var returnValue *Address
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"address": c.address,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified address resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalAddresses.get",
	//   "parameterOrder": [
	//     "project",
	//     "address"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Name of the address resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/addresses/{address}",
	//   "response": {
	//     "$ref": "Address"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.globalAddresses.insert":

type GlobalAddressesInsertCall struct {
	s             *Service
	project       string
	address       *Address
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates an address resource in the specified project using
// the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalAddresses/insert
func (r *GlobalAddressesService) Insert(project string, address *Address) *GlobalAddressesInsertCall {
	return &GlobalAddressesInsertCall{
		s:             r.s,
		project:       project,
		address:       address,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/addresses",
		context_:      googleapi.NoContext,
	}
}
func (c *GlobalAddressesInsertCall) Context(ctx context.Context) *GlobalAddressesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAddressesInsertCall) Fields(s ...googleapi.Field) *GlobalAddressesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalAddressesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.address,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an address resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.globalAddresses.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/addresses",
	//   "request": {
	//     "$ref": "Address"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalAddresses.list":

type GlobalAddressesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of global address resources.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalAddresses/list
func (r *GlobalAddressesService) List(project string) *GlobalAddressesListCall {
	return &GlobalAddressesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/addresses",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *GlobalAddressesListCall) Filter(filter string) *GlobalAddressesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *GlobalAddressesListCall) MaxResults(maxResults int64) *GlobalAddressesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *GlobalAddressesListCall) PageToken(pageToken string) *GlobalAddressesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *GlobalAddressesListCall) Context(ctx context.Context) *GlobalAddressesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAddressesListCall) Fields(s ...googleapi.Field) *GlobalAddressesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalAddressesListCall) Do() (*AddressList, error) {
	var returnValue *AddressList
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
	//   "description": "Retrieves the list of global address resources.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalAddresses.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/addresses",
	//   "response": {
	//     "$ref": "AddressList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.globalForwardingRules.delete":

type GlobalForwardingRulesDeleteCall struct {
	s              *Service
	project        string
	forwardingRule string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Delete: Deletes the specified ForwardingRule resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalForwardingRules/delete
func (r *GlobalForwardingRulesService) Delete(project string, forwardingRule string) *GlobalForwardingRulesDeleteCall {
	return &GlobalForwardingRulesDeleteCall{
		s:              r.s,
		project:        project,
		forwardingRule: forwardingRule,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/forwardingRules/{forwardingRule}",
		context_:       googleapi.NoContext,
	}
}
func (c *GlobalForwardingRulesDeleteCall) Context(ctx context.Context) *GlobalForwardingRulesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalForwardingRulesDeleteCall) Fields(s ...googleapi.Field) *GlobalForwardingRulesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalForwardingRulesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"forwardingRule": c.forwardingRule,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified ForwardingRule resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.globalForwardingRules.delete",
	//   "parameterOrder": [
	//     "project",
	//     "forwardingRule"
	//   ],
	//   "parameters": {
	//     "forwardingRule": {
	//       "description": "Name of the ForwardingRule resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/forwardingRules/{forwardingRule}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalForwardingRules.get":

type GlobalForwardingRulesGetCall struct {
	s              *Service
	project        string
	forwardingRule string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Get: Returns the specified ForwardingRule resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalForwardingRules/get
func (r *GlobalForwardingRulesService) Get(project string, forwardingRule string) *GlobalForwardingRulesGetCall {
	return &GlobalForwardingRulesGetCall{
		s:              r.s,
		project:        project,
		forwardingRule: forwardingRule,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/forwardingRules/{forwardingRule}",
		context_:       googleapi.NoContext,
	}
}
func (c *GlobalForwardingRulesGetCall) Context(ctx context.Context) *GlobalForwardingRulesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalForwardingRulesGetCall) Fields(s ...googleapi.Field) *GlobalForwardingRulesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalForwardingRulesGetCall) Do() (*ForwardingRule, error) {
	var returnValue *ForwardingRule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"forwardingRule": c.forwardingRule,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified ForwardingRule resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalForwardingRules.get",
	//   "parameterOrder": [
	//     "project",
	//     "forwardingRule"
	//   ],
	//   "parameters": {
	//     "forwardingRule": {
	//       "description": "Name of the ForwardingRule resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/forwardingRules/{forwardingRule}",
	//   "response": {
	//     "$ref": "ForwardingRule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.globalForwardingRules.insert":

type GlobalForwardingRulesInsertCall struct {
	s              *Service
	project        string
	forwardingrule *ForwardingRule
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Insert: Creates a ForwardingRule resource in the specified project
// and region using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalForwardingRules/insert
func (r *GlobalForwardingRulesService) Insert(project string, forwardingrule *ForwardingRule) *GlobalForwardingRulesInsertCall {
	return &GlobalForwardingRulesInsertCall{
		s:              r.s,
		project:        project,
		forwardingrule: forwardingrule,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/global/forwardingRules",
		context_:       googleapi.NoContext,
	}
}
func (c *GlobalForwardingRulesInsertCall) Context(ctx context.Context) *GlobalForwardingRulesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalForwardingRulesInsertCall) Fields(s ...googleapi.Field) *GlobalForwardingRulesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalForwardingRulesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.forwardingrule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a ForwardingRule resource in the specified project and region using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.globalForwardingRules.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/forwardingRules",
	//   "request": {
	//     "$ref": "ForwardingRule"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalForwardingRules.list":

type GlobalForwardingRulesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of ForwardingRule resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalForwardingRules/list
func (r *GlobalForwardingRulesService) List(project string) *GlobalForwardingRulesListCall {
	return &GlobalForwardingRulesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/forwardingRules",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *GlobalForwardingRulesListCall) Filter(filter string) *GlobalForwardingRulesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *GlobalForwardingRulesListCall) MaxResults(maxResults int64) *GlobalForwardingRulesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *GlobalForwardingRulesListCall) PageToken(pageToken string) *GlobalForwardingRulesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *GlobalForwardingRulesListCall) Context(ctx context.Context) *GlobalForwardingRulesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalForwardingRulesListCall) Fields(s ...googleapi.Field) *GlobalForwardingRulesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalForwardingRulesListCall) Do() (*ForwardingRuleList, error) {
	var returnValue *ForwardingRuleList
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
	//   "description": "Retrieves the list of ForwardingRule resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalForwardingRules.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/forwardingRules",
	//   "response": {
	//     "$ref": "ForwardingRuleList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.globalForwardingRules.setTarget":

type GlobalForwardingRulesSetTargetCall struct {
	s               *Service
	project         string
	forwardingRule  string
	targetreference *TargetReference
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// SetTarget: Changes target url for forwarding rule.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalForwardingRules/setTarget
func (r *GlobalForwardingRulesService) SetTarget(project string, forwardingRule string, targetreference *TargetReference) *GlobalForwardingRulesSetTargetCall {
	return &GlobalForwardingRulesSetTargetCall{
		s:               r.s,
		project:         project,
		forwardingRule:  forwardingRule,
		targetreference: targetreference,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/forwardingRules/{forwardingRule}/setTarget",
		context_:        googleapi.NoContext,
	}
}
func (c *GlobalForwardingRulesSetTargetCall) Context(ctx context.Context) *GlobalForwardingRulesSetTargetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalForwardingRulesSetTargetCall) Fields(s ...googleapi.Field) *GlobalForwardingRulesSetTargetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalForwardingRulesSetTargetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"forwardingRule": c.forwardingRule,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetreference,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Changes target url for forwarding rule.",
	//   "httpMethod": "POST",
	//   "id": "compute.globalForwardingRules.setTarget",
	//   "parameterOrder": [
	//     "project",
	//     "forwardingRule"
	//   ],
	//   "parameters": {
	//     "forwardingRule": {
	//       "description": "Name of the ForwardingRule resource in which target is to be set.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/forwardingRules/{forwardingRule}/setTarget",
	//   "request": {
	//     "$ref": "TargetReference"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalOperations.aggregatedList":

type GlobalOperationsAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of all operations grouped by
// scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalOperations/aggregatedList
func (r *GlobalOperationsService) AggregatedList(project string) *GlobalOperationsAggregatedListCall {
	return &GlobalOperationsAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/operations",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *GlobalOperationsAggregatedListCall) Filter(filter string) *GlobalOperationsAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *GlobalOperationsAggregatedListCall) MaxResults(maxResults int64) *GlobalOperationsAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *GlobalOperationsAggregatedListCall) PageToken(pageToken string) *GlobalOperationsAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *GlobalOperationsAggregatedListCall) Context(ctx context.Context) *GlobalOperationsAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalOperationsAggregatedListCall) Fields(s ...googleapi.Field) *GlobalOperationsAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalOperationsAggregatedListCall) Do() (*OperationAggregatedList, error) {
	var returnValue *OperationAggregatedList
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
	//   "description": "Retrieves the list of all operations grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalOperations.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/operations",
	//   "response": {
	//     "$ref": "OperationAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.globalOperations.delete":

type GlobalOperationsDeleteCall struct {
	s             *Service
	project       string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified operation resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalOperations/delete
func (r *GlobalOperationsService) Delete(project string, operation string) *GlobalOperationsDeleteCall {
	return &GlobalOperationsDeleteCall{
		s:             r.s,
		project:       project,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *GlobalOperationsDeleteCall) Context(ctx context.Context) *GlobalOperationsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *GlobalOperationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"operation": c.operation,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified operation resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.globalOperations.delete",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations/{operation}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.globalOperations.get":

type GlobalOperationsGetCall struct {
	s             *Service
	project       string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the specified operation resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalOperations/get
func (r *GlobalOperationsService) Get(project string, operation string) *GlobalOperationsGetCall {
	return &GlobalOperationsGetCall{
		s:             r.s,
		project:       project,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *GlobalOperationsGetCall) Context(ctx context.Context) *GlobalOperationsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalOperationsGetCall) Fields(s ...googleapi.Field) *GlobalOperationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalOperationsGetCall) Do() (*Operation, error) {
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
	//   "description": "Retrieves the specified operation resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
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
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.globalOperations.list":

type GlobalOperationsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of operation resources contained within the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/globalOperations/list
func (r *GlobalOperationsService) List(project string) *GlobalOperationsListCall {
	return &GlobalOperationsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/operations",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *GlobalOperationsListCall) Filter(filter string) *GlobalOperationsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *GlobalOperationsListCall) MaxResults(maxResults int64) *GlobalOperationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *GlobalOperationsListCall) PageToken(pageToken string) *GlobalOperationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *GlobalOperationsListCall) Context(ctx context.Context) *GlobalOperationsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalOperationsListCall) Fields(s ...googleapi.Field) *GlobalOperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GlobalOperationsListCall) Do() (*OperationList, error) {
	var returnValue *OperationList
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
	//   "description": "Retrieves the list of operation resources contained within the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.globalOperations.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.httpHealthChecks.delete":

type HttpHealthChecksDeleteCall struct {
	s               *Service
	project         string
	httpHealthCheck string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Delete: Deletes the specified HttpHealthCheck resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/httpHealthChecks/delete
func (r *HttpHealthChecksService) Delete(project string, httpHealthCheck string) *HttpHealthChecksDeleteCall {
	return &HttpHealthChecksDeleteCall{
		s:               r.s,
		project:         project,
		httpHealthCheck: httpHealthCheck,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/httpHealthChecks/{httpHealthCheck}",
		context_:        googleapi.NoContext,
	}
}
func (c *HttpHealthChecksDeleteCall) Context(ctx context.Context) *HttpHealthChecksDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HttpHealthChecksDeleteCall) Fields(s ...googleapi.Field) *HttpHealthChecksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HttpHealthChecksDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"httpHealthCheck": c.httpHealthCheck,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified HttpHealthCheck resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.httpHealthChecks.delete",
	//   "parameterOrder": [
	//     "project",
	//     "httpHealthCheck"
	//   ],
	//   "parameters": {
	//     "httpHealthCheck": {
	//       "description": "Name of the HttpHealthCheck resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/httpHealthChecks/{httpHealthCheck}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.httpHealthChecks.get":

type HttpHealthChecksGetCall struct {
	s               *Service
	project         string
	httpHealthCheck string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Get: Returns the specified HttpHealthCheck resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/httpHealthChecks/get
func (r *HttpHealthChecksService) Get(project string, httpHealthCheck string) *HttpHealthChecksGetCall {
	return &HttpHealthChecksGetCall{
		s:               r.s,
		project:         project,
		httpHealthCheck: httpHealthCheck,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/httpHealthChecks/{httpHealthCheck}",
		context_:        googleapi.NoContext,
	}
}
func (c *HttpHealthChecksGetCall) Context(ctx context.Context) *HttpHealthChecksGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HttpHealthChecksGetCall) Fields(s ...googleapi.Field) *HttpHealthChecksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HttpHealthChecksGetCall) Do() (*HttpHealthCheck, error) {
	var returnValue *HttpHealthCheck
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"httpHealthCheck": c.httpHealthCheck,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified HttpHealthCheck resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.httpHealthChecks.get",
	//   "parameterOrder": [
	//     "project",
	//     "httpHealthCheck"
	//   ],
	//   "parameters": {
	//     "httpHealthCheck": {
	//       "description": "Name of the HttpHealthCheck resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/httpHealthChecks/{httpHealthCheck}",
	//   "response": {
	//     "$ref": "HttpHealthCheck"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.httpHealthChecks.insert":

type HttpHealthChecksInsertCall struct {
	s               *Service
	project         string
	httphealthcheck *HttpHealthCheck
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Insert: Creates a HttpHealthCheck resource in the specified project
// using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/httpHealthChecks/insert
func (r *HttpHealthChecksService) Insert(project string, httphealthcheck *HttpHealthCheck) *HttpHealthChecksInsertCall {
	return &HttpHealthChecksInsertCall{
		s:               r.s,
		project:         project,
		httphealthcheck: httphealthcheck,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/httpHealthChecks",
		context_:        googleapi.NoContext,
	}
}
func (c *HttpHealthChecksInsertCall) Context(ctx context.Context) *HttpHealthChecksInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HttpHealthChecksInsertCall) Fields(s ...googleapi.Field) *HttpHealthChecksInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HttpHealthChecksInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.httphealthcheck,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a HttpHealthCheck resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.httpHealthChecks.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/httpHealthChecks",
	//   "request": {
	//     "$ref": "HttpHealthCheck"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.httpHealthChecks.list":

type HttpHealthChecksListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of HttpHealthCheck resources available to
// the specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/httpHealthChecks/list
func (r *HttpHealthChecksService) List(project string) *HttpHealthChecksListCall {
	return &HttpHealthChecksListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/httpHealthChecks",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *HttpHealthChecksListCall) Filter(filter string) *HttpHealthChecksListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *HttpHealthChecksListCall) MaxResults(maxResults int64) *HttpHealthChecksListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *HttpHealthChecksListCall) PageToken(pageToken string) *HttpHealthChecksListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *HttpHealthChecksListCall) Context(ctx context.Context) *HttpHealthChecksListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HttpHealthChecksListCall) Fields(s ...googleapi.Field) *HttpHealthChecksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HttpHealthChecksListCall) Do() (*HttpHealthCheckList, error) {
	var returnValue *HttpHealthCheckList
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
	//   "description": "Retrieves the list of HttpHealthCheck resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.httpHealthChecks.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/httpHealthChecks",
	//   "response": {
	//     "$ref": "HttpHealthCheckList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.httpHealthChecks.patch":

type HttpHealthChecksPatchCall struct {
	s               *Service
	project         string
	httpHealthCheck string
	httphealthcheck *HttpHealthCheck
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Patch: Updates a HttpHealthCheck resource in the specified project
// using the data included in the request. This method supports patch
// semantics.

// For details, see https://cloud.google.com/compute/docs/reference/latest/httpHealthChecks/patch
func (r *HttpHealthChecksService) Patch(project string, httpHealthCheck string, httphealthcheck *HttpHealthCheck) *HttpHealthChecksPatchCall {
	return &HttpHealthChecksPatchCall{
		s:               r.s,
		project:         project,
		httpHealthCheck: httpHealthCheck,
		httphealthcheck: httphealthcheck,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/httpHealthChecks/{httpHealthCheck}",
		context_:        googleapi.NoContext,
	}
}
func (c *HttpHealthChecksPatchCall) Context(ctx context.Context) *HttpHealthChecksPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HttpHealthChecksPatchCall) Fields(s ...googleapi.Field) *HttpHealthChecksPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HttpHealthChecksPatchCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"httpHealthCheck": c.httpHealthCheck,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.httphealthcheck,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates a HttpHealthCheck resource in the specified project using the data included in the request. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "compute.httpHealthChecks.patch",
	//   "parameterOrder": [
	//     "project",
	//     "httpHealthCheck"
	//   ],
	//   "parameters": {
	//     "httpHealthCheck": {
	//       "description": "Name of the HttpHealthCheck resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/httpHealthChecks/{httpHealthCheck}",
	//   "request": {
	//     "$ref": "HttpHealthCheck"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.httpHealthChecks.update":

type HttpHealthChecksUpdateCall struct {
	s               *Service
	project         string
	httpHealthCheck string
	httphealthcheck *HttpHealthCheck
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Update: Updates a HttpHealthCheck resource in the specified project
// using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/httpHealthChecks/update
func (r *HttpHealthChecksService) Update(project string, httpHealthCheck string, httphealthcheck *HttpHealthCheck) *HttpHealthChecksUpdateCall {
	return &HttpHealthChecksUpdateCall{
		s:               r.s,
		project:         project,
		httpHealthCheck: httpHealthCheck,
		httphealthcheck: httphealthcheck,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/httpHealthChecks/{httpHealthCheck}",
		context_:        googleapi.NoContext,
	}
}
func (c *HttpHealthChecksUpdateCall) Context(ctx context.Context) *HttpHealthChecksUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HttpHealthChecksUpdateCall) Fields(s ...googleapi.Field) *HttpHealthChecksUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HttpHealthChecksUpdateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"httpHealthCheck": c.httpHealthCheck,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.httphealthcheck,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates a HttpHealthCheck resource in the specified project using the data included in the request.",
	//   "httpMethod": "PUT",
	//   "id": "compute.httpHealthChecks.update",
	//   "parameterOrder": [
	//     "project",
	//     "httpHealthCheck"
	//   ],
	//   "parameters": {
	//     "httpHealthCheck": {
	//       "description": "Name of the HttpHealthCheck resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/httpHealthChecks/{httpHealthCheck}",
	//   "request": {
	//     "$ref": "HttpHealthCheck"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.images.delete":

type ImagesDeleteCall struct {
	s             *Service
	project       string
	image         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified image resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/images/delete
func (r *ImagesService) Delete(project string, image string) *ImagesDeleteCall {
	return &ImagesDeleteCall{
		s:             r.s,
		project:       project,
		image:         image,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/images/{image}",
		context_:      googleapi.NoContext,
	}
}
func (c *ImagesDeleteCall) Context(ctx context.Context) *ImagesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImagesDeleteCall) Fields(s ...googleapi.Field) *ImagesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ImagesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"image":   c.image,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified image resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.images.delete",
	//   "parameterOrder": [
	//     "project",
	//     "image"
	//   ],
	//   "parameters": {
	//     "image": {
	//       "description": "Name of the image resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/images/{image}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.images.deprecate":

type ImagesDeprecateCall struct {
	s                 *Service
	project           string
	image             string
	deprecationstatus *DeprecationStatus
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
	context_          context.Context
}

// Deprecate: Sets the deprecation status of an image.
//
// If an empty
// request body is given, clears the deprecation status instead.

// For details, see https://cloud.google.com/compute/docs/reference/latest/images/deprecate
func (r *ImagesService) Deprecate(project string, image string, deprecationstatus *DeprecationStatus) *ImagesDeprecateCall {
	return &ImagesDeprecateCall{
		s:                 r.s,
		project:           project,
		image:             image,
		deprecationstatus: deprecationstatus,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "{project}/global/images/{image}/deprecate",
		context_:          googleapi.NoContext,
	}
}
func (c *ImagesDeprecateCall) Context(ctx context.Context) *ImagesDeprecateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImagesDeprecateCall) Fields(s ...googleapi.Field) *ImagesDeprecateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ImagesDeprecateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"image":   c.image,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.deprecationstatus,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Sets the deprecation status of an image.\n\nIf an empty request body is given, clears the deprecation status instead.",
	//   "httpMethod": "POST",
	//   "id": "compute.images.deprecate",
	//   "parameterOrder": [
	//     "project",
	//     "image"
	//   ],
	//   "parameters": {
	//     "image": {
	//       "description": "Image name.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/images/{image}/deprecate",
	//   "request": {
	//     "$ref": "DeprecationStatus"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.images.get":

type ImagesGetCall struct {
	s             *Service
	project       string
	image         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified image resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/images/get
func (r *ImagesService) Get(project string, image string) *ImagesGetCall {
	return &ImagesGetCall{
		s:             r.s,
		project:       project,
		image:         image,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/images/{image}",
		context_:      googleapi.NoContext,
	}
}
func (c *ImagesGetCall) Context(ctx context.Context) *ImagesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImagesGetCall) Fields(s ...googleapi.Field) *ImagesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ImagesGetCall) Do() (*Image, error) {
	var returnValue *Image
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"image":   c.image,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified image resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.images.get",
	//   "parameterOrder": [
	//     "project",
	//     "image"
	//   ],
	//   "parameters": {
	//     "image": {
	//       "description": "Name of the image resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/images/{image}",
	//   "response": {
	//     "$ref": "Image"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.images.insert":

type ImagesInsertCall struct {
	s             *Service
	project       string
	image         *Image
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates an image resource in the specified project using the
// data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/images/insert
func (r *ImagesService) Insert(project string, image *Image) *ImagesInsertCall {
	return &ImagesInsertCall{
		s:             r.s,
		project:       project,
		image:         image,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/images",
		context_:      googleapi.NoContext,
	}
}
func (c *ImagesInsertCall) Context(ctx context.Context) *ImagesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImagesInsertCall) Fields(s ...googleapi.Field) *ImagesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ImagesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.image,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an image resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.images.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/images",
	//   "request": {
	//     "$ref": "Image"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "compute.images.list":

type ImagesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of image resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/images/list
func (r *ImagesService) List(project string) *ImagesListCall {
	return &ImagesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/images",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ImagesListCall) Filter(filter string) *ImagesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ImagesListCall) MaxResults(maxResults int64) *ImagesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *ImagesListCall) PageToken(pageToken string) *ImagesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ImagesListCall) Context(ctx context.Context) *ImagesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImagesListCall) Fields(s ...googleapi.Field) *ImagesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ImagesListCall) Do() (*ImageList, error) {
	var returnValue *ImageList
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
	//   "description": "Retrieves the list of image resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.images.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/images",
	//   "response": {
	//     "$ref": "ImageList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instanceTemplates.delete":

type InstanceTemplatesDeleteCall struct {
	s                *Service
	project          string
	instanceTemplate string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Delete: Deletes the specified instance template resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instanceTemplates/delete
func (r *InstanceTemplatesService) Delete(project string, instanceTemplate string) *InstanceTemplatesDeleteCall {
	return &InstanceTemplatesDeleteCall{
		s:                r.s,
		project:          project,
		instanceTemplate: instanceTemplate,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/global/instanceTemplates/{instanceTemplate}",
		context_:         googleapi.NoContext,
	}
}
func (c *InstanceTemplatesDeleteCall) Context(ctx context.Context) *InstanceTemplatesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceTemplatesDeleteCall) Fields(s ...googleapi.Field) *InstanceTemplatesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceTemplatesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":          c.project,
		"instanceTemplate": c.instanceTemplate,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified instance template resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.instanceTemplates.delete",
	//   "parameterOrder": [
	//     "project",
	//     "instanceTemplate"
	//   ],
	//   "parameters": {
	//     "instanceTemplate": {
	//       "description": "Name of the instance template resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/instanceTemplates/{instanceTemplate}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instanceTemplates.get":

type InstanceTemplatesGetCall struct {
	s                *Service
	project          string
	instanceTemplate string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Get: Returns the specified instance template resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instanceTemplates/get
func (r *InstanceTemplatesService) Get(project string, instanceTemplate string) *InstanceTemplatesGetCall {
	return &InstanceTemplatesGetCall{
		s:                r.s,
		project:          project,
		instanceTemplate: instanceTemplate,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/global/instanceTemplates/{instanceTemplate}",
		context_:         googleapi.NoContext,
	}
}
func (c *InstanceTemplatesGetCall) Context(ctx context.Context) *InstanceTemplatesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceTemplatesGetCall) Fields(s ...googleapi.Field) *InstanceTemplatesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceTemplatesGetCall) Do() (*InstanceTemplate, error) {
	var returnValue *InstanceTemplate
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":          c.project,
		"instanceTemplate": c.instanceTemplate,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified instance template resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.instanceTemplates.get",
	//   "parameterOrder": [
	//     "project",
	//     "instanceTemplate"
	//   ],
	//   "parameters": {
	//     "instanceTemplate": {
	//       "description": "Name of the instance template resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/instanceTemplates/{instanceTemplate}",
	//   "response": {
	//     "$ref": "InstanceTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instanceTemplates.insert":

type InstanceTemplatesInsertCall struct {
	s                *Service
	project          string
	instancetemplate *InstanceTemplate
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Insert: Creates an instance template resource in the specified
// project using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instanceTemplates/insert
func (r *InstanceTemplatesService) Insert(project string, instancetemplate *InstanceTemplate) *InstanceTemplatesInsertCall {
	return &InstanceTemplatesInsertCall{
		s:                r.s,
		project:          project,
		instancetemplate: instancetemplate,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/global/instanceTemplates",
		context_:         googleapi.NoContext,
	}
}
func (c *InstanceTemplatesInsertCall) Context(ctx context.Context) *InstanceTemplatesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceTemplatesInsertCall) Fields(s ...googleapi.Field) *InstanceTemplatesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceTemplatesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancetemplate,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an instance template resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.instanceTemplates.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/instanceTemplates",
	//   "request": {
	//     "$ref": "InstanceTemplate"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instanceTemplates.list":

type InstanceTemplatesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of instance template resources contained
// within the specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instanceTemplates/list
func (r *InstanceTemplatesService) List(project string) *InstanceTemplatesListCall {
	return &InstanceTemplatesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/instanceTemplates",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *InstanceTemplatesListCall) Filter(filter string) *InstanceTemplatesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *InstanceTemplatesListCall) MaxResults(maxResults int64) *InstanceTemplatesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *InstanceTemplatesListCall) PageToken(pageToken string) *InstanceTemplatesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *InstanceTemplatesListCall) Context(ctx context.Context) *InstanceTemplatesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceTemplatesListCall) Fields(s ...googleapi.Field) *InstanceTemplatesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceTemplatesListCall) Do() (*InstanceTemplateList, error) {
	var returnValue *InstanceTemplateList
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
	//   "description": "Retrieves the list of instance template resources contained within the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.instanceTemplates.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/instanceTemplates",
	//   "response": {
	//     "$ref": "InstanceTemplateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instances.addAccessConfig":

type InstancesAddAccessConfigCall struct {
	s                *Service
	project          string
	zone             string
	instance         string
	networkInterface string
	accessconfig     *AccessConfig
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// AddAccessConfig: Adds an access config to an instance's network
// interface.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/addAccessConfig
func (r *InstancesService) AddAccessConfig(project string, zone string, instance string, networkInterface string, accessconfig *AccessConfig) *InstancesAddAccessConfigCall {
	return &InstancesAddAccessConfigCall{
		s:                r.s,
		project:          project,
		zone:             zone,
		instance:         instance,
		networkInterface: networkInterface,
		accessconfig:     accessconfig,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/zones/{zone}/instances/{instance}/addAccessConfig",
		context_:         googleapi.NoContext,
	}
}
func (c *InstancesAddAccessConfigCall) Context(ctx context.Context) *InstancesAddAccessConfigCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesAddAccessConfigCall) Fields(s ...googleapi.Field) *InstancesAddAccessConfigCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesAddAccessConfigCall) Do() (*Operation, error) {
	var returnValue *Operation
	c.params_.Set("networkInterface", fmt.Sprintf("%v", c.networkInterface))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.accessconfig,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Adds an access config to an instance's network interface.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.addAccessConfig",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance",
	//     "networkInterface"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "The instance name for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "networkInterface": {
	//       "description": "The name of the network interface to add to this instance.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/addAccessConfig",
	//   "request": {
	//     "$ref": "AccessConfig"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.aggregatedList":

type InstancesAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList:

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/aggregatedList
func (r *InstancesService) AggregatedList(project string) *InstancesAggregatedListCall {
	return &InstancesAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/instances",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *InstancesAggregatedListCall) Filter(filter string) *InstancesAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *InstancesAggregatedListCall) MaxResults(maxResults int64) *InstancesAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *InstancesAggregatedListCall) PageToken(pageToken string) *InstancesAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *InstancesAggregatedListCall) Context(ctx context.Context) *InstancesAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesAggregatedListCall) Fields(s ...googleapi.Field) *InstancesAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesAggregatedListCall) Do() (*InstanceAggregatedList, error) {
	var returnValue *InstanceAggregatedList
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
	//   "httpMethod": "GET",
	//   "id": "compute.instances.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/instances",
	//   "response": {
	//     "$ref": "InstanceAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instances.attachDisk":

type InstancesAttachDiskCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	attacheddisk  *AttachedDisk
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AttachDisk: Attaches a Disk resource to an instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/attachDisk
func (r *InstancesService) AttachDisk(project string, zone string, instance string, attacheddisk *AttachedDisk) *InstancesAttachDiskCall {
	return &InstancesAttachDiskCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		attacheddisk:  attacheddisk,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/attachDisk",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesAttachDiskCall) Context(ctx context.Context) *InstancesAttachDiskCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesAttachDiskCall) Fields(s ...googleapi.Field) *InstancesAttachDiskCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesAttachDiskCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.attacheddisk,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Attaches a Disk resource to an instance.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.attachDisk",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Instance name.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/attachDisk",
	//   "request": {
	//     "$ref": "AttachedDisk"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.delete":

type InstancesDeleteCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified Instance resource. For more
// information, see Shutting down an instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/delete
func (r *InstancesService) Delete(project string, zone string, instance string) *InstancesDeleteCall {
	return &InstancesDeleteCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesDeleteCall) Context(ctx context.Context) *InstancesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesDeleteCall) Fields(s ...googleapi.Field) *InstancesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified Instance resource. For more information, see Shutting down an instance.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.instances.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.deleteAccessConfig":

type InstancesDeleteAccessConfigCall struct {
	s                *Service
	project          string
	zone             string
	instance         string
	accessConfig     string
	networkInterface string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// DeleteAccessConfig: Deletes an access config from an instance's
// network interface.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/deleteAccessConfig
func (r *InstancesService) DeleteAccessConfig(project string, zone string, instance string, accessConfig string, networkInterface string) *InstancesDeleteAccessConfigCall {
	return &InstancesDeleteAccessConfigCall{
		s:                r.s,
		project:          project,
		zone:             zone,
		instance:         instance,
		accessConfig:     accessConfig,
		networkInterface: networkInterface,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/zones/{zone}/instances/{instance}/deleteAccessConfig",
		context_:         googleapi.NoContext,
	}
}
func (c *InstancesDeleteAccessConfigCall) Context(ctx context.Context) *InstancesDeleteAccessConfigCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesDeleteAccessConfigCall) Fields(s ...googleapi.Field) *InstancesDeleteAccessConfigCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesDeleteAccessConfigCall) Do() (*Operation, error) {
	var returnValue *Operation
	c.params_.Set("accessConfig", fmt.Sprintf("%v", c.accessConfig))
	c.params_.Set("networkInterface", fmt.Sprintf("%v", c.networkInterface))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes an access config from an instance's network interface.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.deleteAccessConfig",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance",
	//     "accessConfig",
	//     "networkInterface"
	//   ],
	//   "parameters": {
	//     "accessConfig": {
	//       "description": "The name of the access config to delete.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "The instance name for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "networkInterface": {
	//       "description": "The name of the network interface.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/deleteAccessConfig",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.detachDisk":

type InstancesDetachDiskCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	deviceName    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// DetachDisk: Detaches a disk from an instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/detachDisk
func (r *InstancesService) DetachDisk(project string, zone string, instance string, deviceName string) *InstancesDetachDiskCall {
	return &InstancesDetachDiskCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		deviceName:    deviceName,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/detachDisk",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesDetachDiskCall) Context(ctx context.Context) *InstancesDetachDiskCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesDetachDiskCall) Fields(s ...googleapi.Field) *InstancesDetachDiskCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesDetachDiskCall) Do() (*Operation, error) {
	var returnValue *Operation
	c.params_.Set("deviceName", fmt.Sprintf("%v", c.deviceName))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Detaches a disk from an instance.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.detachDisk",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance",
	//     "deviceName"
	//   ],
	//   "parameters": {
	//     "deviceName": {
	//       "description": "Disk device name to detach.",
	//       "location": "query",
	//       "pattern": "\\w[\\w.-]{0,254}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Instance name.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/detachDisk",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.get":

type InstancesGetCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified instance resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/get
func (r *InstancesService) Get(project string, zone string, instance string) *InstancesGetCall {
	return &InstancesGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesGetCall) Context(ctx context.Context) *InstancesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesGetCall) Fields(s ...googleapi.Field) *InstancesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesGetCall) Do() (*Instance, error) {
	var returnValue *Instance
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified instance resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.instances.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the The name of the zone for this request..",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}",
	//   "response": {
	//     "$ref": "Instance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instances.getSerialPortOutput":

type InstancesGetSerialPortOutputCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// GetSerialPortOutput: Returns the specified instance's serial port
// output.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/getSerialPortOutput
func (r *InstancesService) GetSerialPortOutput(project string, zone string, instance string) *InstancesGetSerialPortOutputCall {
	return &InstancesGetSerialPortOutputCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/serialPort",
		context_:      googleapi.NoContext,
	}
}

// Port sets the optional parameter "port": Which COM port to retrieve
// data from.
func (c *InstancesGetSerialPortOutputCall) Port(port int64) *InstancesGetSerialPortOutputCall {
	c.params_.Set("port", fmt.Sprintf("%v", port))
	return c
}
func (c *InstancesGetSerialPortOutputCall) Context(ctx context.Context) *InstancesGetSerialPortOutputCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesGetSerialPortOutputCall) Fields(s ...googleapi.Field) *InstancesGetSerialPortOutputCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesGetSerialPortOutputCall) Do() (*SerialPortOutput, error) {
	var returnValue *SerialPortOutput
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified instance's serial port output.",
	//   "httpMethod": "GET",
	//   "id": "compute.instances.getSerialPortOutput",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "port": {
	//       "default": "1",
	//       "description": "Which COM port to retrieve data from.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "4",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/serialPort",
	//   "response": {
	//     "$ref": "SerialPortOutput"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instances.insert":

type InstancesInsertCall struct {
	s             *Service
	project       string
	zone          string
	instance      *Instance
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates an instance resource in the specified project using
// the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/insert
func (r *InstancesService) Insert(project string, zone string, instance *Instance) *InstancesInsertCall {
	return &InstancesInsertCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesInsertCall) Context(ctx context.Context) *InstancesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesInsertCall) Fields(s ...googleapi.Field) *InstancesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instance,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an instance resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances",
	//   "request": {
	//     "$ref": "Instance"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.list":

type InstancesListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of instance resources contained within the
// specified zone.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/list
func (r *InstancesService) List(project string, zone string) *InstancesListCall {
	return &InstancesListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *InstancesListCall) Filter(filter string) *InstancesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *InstancesListCall) MaxResults(maxResults int64) *InstancesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *InstancesListCall) PageToken(pageToken string) *InstancesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *InstancesListCall) Context(ctx context.Context) *InstancesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesListCall) Fields(s ...googleapi.Field) *InstancesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesListCall) Do() (*InstanceList, error) {
	var returnValue *InstanceList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of instance resources contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "compute.instances.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances",
	//   "response": {
	//     "$ref": "InstanceList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.instances.reset":

type InstancesResetCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Reset: Performs a hard reset on the instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/reset
func (r *InstancesService) Reset(project string, zone string, instance string) *InstancesResetCall {
	return &InstancesResetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/reset",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesResetCall) Context(ctx context.Context) *InstancesResetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesResetCall) Fields(s ...googleapi.Field) *InstancesResetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesResetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Performs a hard reset on the instance.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.reset",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/reset",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.setDiskAutoDelete":

type InstancesSetDiskAutoDeleteCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	autoDelete    bool
	deviceName    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// SetDiskAutoDelete: Sets the auto-delete flag for a disk attached to
// an instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/setDiskAutoDelete
func (r *InstancesService) SetDiskAutoDelete(project string, zone string, instance string, autoDelete bool, deviceName string) *InstancesSetDiskAutoDeleteCall {
	return &InstancesSetDiskAutoDeleteCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		autoDelete:    autoDelete,
		deviceName:    deviceName,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/setDiskAutoDelete",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesSetDiskAutoDeleteCall) Context(ctx context.Context) *InstancesSetDiskAutoDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesSetDiskAutoDeleteCall) Fields(s ...googleapi.Field) *InstancesSetDiskAutoDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesSetDiskAutoDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	c.params_.Set("autoDelete", fmt.Sprintf("%v", c.autoDelete))
	c.params_.Set("deviceName", fmt.Sprintf("%v", c.deviceName))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Sets the auto-delete flag for a disk attached to an instance.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.setDiskAutoDelete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance",
	//     "autoDelete",
	//     "deviceName"
	//   ],
	//   "parameters": {
	//     "autoDelete": {
	//       "description": "Whether to auto-delete the disk when the instance is deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "boolean"
	//     },
	//     "deviceName": {
	//       "description": "The device name of the disk to modify.",
	//       "location": "query",
	//       "pattern": "\\w[\\w.-]{0,254}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "The instance name.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/setDiskAutoDelete",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.setMetadata":

type InstancesSetMetadataCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	metadata      *Metadata
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// SetMetadata: Sets metadata for the specified instance to the data
// included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/setMetadata
func (r *InstancesService) SetMetadata(project string, zone string, instance string, metadata *Metadata) *InstancesSetMetadataCall {
	return &InstancesSetMetadataCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		metadata:      metadata,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/setMetadata",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesSetMetadataCall) Context(ctx context.Context) *InstancesSetMetadataCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesSetMetadataCall) Fields(s ...googleapi.Field) *InstancesSetMetadataCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesSetMetadataCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.metadata,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Sets metadata for the specified instance to the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.setMetadata",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/setMetadata",
	//   "request": {
	//     "$ref": "Metadata"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.setScheduling":

type InstancesSetSchedulingCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	scheduling    *Scheduling
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// SetScheduling: Sets an instance's scheduling options.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/setScheduling
func (r *InstancesService) SetScheduling(project string, zone string, instance string, scheduling *Scheduling) *InstancesSetSchedulingCall {
	return &InstancesSetSchedulingCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		scheduling:    scheduling,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/setScheduling",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesSetSchedulingCall) Context(ctx context.Context) *InstancesSetSchedulingCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesSetSchedulingCall) Fields(s ...googleapi.Field) *InstancesSetSchedulingCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesSetSchedulingCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.scheduling,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Sets an instance's scheduling options.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.setScheduling",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Instance name.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/setScheduling",
	//   "request": {
	//     "$ref": "Scheduling"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.setTags":

type InstancesSetTagsCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	tags          *Tags
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// SetTags: Sets tags for the specified instance to the data included in
// the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/setTags
func (r *InstancesService) SetTags(project string, zone string, instance string, tags *Tags) *InstancesSetTagsCall {
	return &InstancesSetTagsCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		tags:          tags,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/setTags",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesSetTagsCall) Context(ctx context.Context) *InstancesSetTagsCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesSetTagsCall) Fields(s ...googleapi.Field) *InstancesSetTagsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesSetTagsCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.tags,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Sets tags for the specified instance to the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.setTags",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/setTags",
	//   "request": {
	//     "$ref": "Tags"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.start":

type InstancesStartCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Start: This method starts an instance that was stopped using the
// using the instances().stop method. For more information, see Restart
// an instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/start
func (r *InstancesService) Start(project string, zone string, instance string) *InstancesStartCall {
	return &InstancesStartCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/start",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesStartCall) Context(ctx context.Context) *InstancesStartCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesStartCall) Fields(s ...googleapi.Field) *InstancesStartCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesStartCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "This method starts an instance that was stopped using the using the instances().stop method. For more information, see Restart an instance.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.start",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance resource to start.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/start",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.instances.stop":

type InstancesStopCall struct {
	s             *Service
	project       string
	zone          string
	instance      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Stop: This method stops a running instance, shutting it down cleanly,
// and allows you to restart the instance at a later time. Stopped
// instances do not incur per-minute, virtual machine usage charges
// while they are stopped, but any resources that the virtual machine is
// using, such as persistent disks and static IP addresses,will continue
// to be charged until they are deleted. For more information, see
// Stopping an instance.

// For details, see https://cloud.google.com/compute/docs/reference/latest/instances/stop
func (r *InstancesService) Stop(project string, zone string, instance string) *InstancesStopCall {
	return &InstancesStopCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		instance:      instance,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instances/{instance}/stop",
		context_:      googleapi.NoContext,
	}
}
func (c *InstancesStopCall) Context(ctx context.Context) *InstancesStopCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesStopCall) Fields(s ...googleapi.Field) *InstancesStopCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstancesStopCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"zone":     c.zone,
		"instance": c.instance,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "This method stops a running instance, shutting it down cleanly, and allows you to restart the instance at a later time. Stopped instances do not incur per-minute, virtual machine usage charges while they are stopped, but any resources that the virtual machine is using, such as persistent disks and static IP addresses,will continue to be charged until they are deleted. For more information, see Stopping an instance.",
	//   "httpMethod": "POST",
	//   "id": "compute.instances.stop",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Name of the instance resource to start.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instances/{instance}/stop",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.licenses.get":

type LicensesGetCall struct {
	s             *Service
	project       string
	license       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified license resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/licenses/get
func (r *LicensesService) Get(project string, license string) *LicensesGetCall {
	return &LicensesGetCall{
		s:             r.s,
		project:       project,
		license:       license,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/licenses/{license}",
		context_:      googleapi.NoContext,
	}
}
func (c *LicensesGetCall) Context(ctx context.Context) *LicensesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicensesGetCall) Fields(s ...googleapi.Field) *LicensesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicensesGetCall) Do() (*License, error) {
	var returnValue *License
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"license": c.license,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified license resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.licenses.get",
	//   "parameterOrder": [
	//     "project",
	//     "license"
	//   ],
	//   "parameters": {
	//     "license": {
	//       "description": "Name of the license resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/licenses/{license}",
	//   "response": {
	//     "$ref": "License"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.machineTypes.aggregatedList":

type MachineTypesAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of machine type resources grouped
// by scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/machineTypes/aggregatedList
func (r *MachineTypesService) AggregatedList(project string) *MachineTypesAggregatedListCall {
	return &MachineTypesAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/machineTypes",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *MachineTypesAggregatedListCall) Filter(filter string) *MachineTypesAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *MachineTypesAggregatedListCall) MaxResults(maxResults int64) *MachineTypesAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *MachineTypesAggregatedListCall) PageToken(pageToken string) *MachineTypesAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *MachineTypesAggregatedListCall) Context(ctx context.Context) *MachineTypesAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MachineTypesAggregatedListCall) Fields(s ...googleapi.Field) *MachineTypesAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MachineTypesAggregatedListCall) Do() (*MachineTypeAggregatedList, error) {
	var returnValue *MachineTypeAggregatedList
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
	//   "description": "Retrieves the list of machine type resources grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.machineTypes.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/machineTypes",
	//   "response": {
	//     "$ref": "MachineTypeAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.machineTypes.get":

type MachineTypesGetCall struct {
	s             *Service
	project       string
	zone          string
	machineType   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified machine type resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/machineTypes/get
func (r *MachineTypesService) Get(project string, zone string, machineType string) *MachineTypesGetCall {
	return &MachineTypesGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		machineType:   machineType,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/machineTypes/{machineType}",
		context_:      googleapi.NoContext,
	}
}
func (c *MachineTypesGetCall) Context(ctx context.Context) *MachineTypesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MachineTypesGetCall) Fields(s ...googleapi.Field) *MachineTypesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MachineTypesGetCall) Do() (*MachineType, error) {
	var returnValue *MachineType
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"zone":        c.zone,
		"machineType": c.machineType,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified machine type resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.machineTypes.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "machineType"
	//   ],
	//   "parameters": {
	//     "machineType": {
	//       "description": "Name of the machine type resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/machineTypes/{machineType}",
	//   "response": {
	//     "$ref": "MachineType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.machineTypes.list":

type MachineTypesListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of machine type resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/machineTypes/list
func (r *MachineTypesService) List(project string, zone string) *MachineTypesListCall {
	return &MachineTypesListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/machineTypes",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *MachineTypesListCall) Filter(filter string) *MachineTypesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *MachineTypesListCall) MaxResults(maxResults int64) *MachineTypesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *MachineTypesListCall) PageToken(pageToken string) *MachineTypesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *MachineTypesListCall) Context(ctx context.Context) *MachineTypesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MachineTypesListCall) Fields(s ...googleapi.Field) *MachineTypesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MachineTypesListCall) Do() (*MachineTypeList, error) {
	var returnValue *MachineTypeList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of machine type resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.machineTypes.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/machineTypes",
	//   "response": {
	//     "$ref": "MachineTypeList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.networks.delete":

type NetworksDeleteCall struct {
	s             *Service
	project       string
	network       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified network resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/networks/delete
func (r *NetworksService) Delete(project string, network string) *NetworksDeleteCall {
	return &NetworksDeleteCall{
		s:             r.s,
		project:       project,
		network:       network,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/networks/{network}",
		context_:      googleapi.NoContext,
	}
}
func (c *NetworksDeleteCall) Context(ctx context.Context) *NetworksDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NetworksDeleteCall) Fields(s ...googleapi.Field) *NetworksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NetworksDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"network": c.network,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified network resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.networks.delete",
	//   "parameterOrder": [
	//     "project",
	//     "network"
	//   ],
	//   "parameters": {
	//     "network": {
	//       "description": "Name of the network resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/networks/{network}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.networks.get":

type NetworksGetCall struct {
	s             *Service
	project       string
	network       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified network resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/networks/get
func (r *NetworksService) Get(project string, network string) *NetworksGetCall {
	return &NetworksGetCall{
		s:             r.s,
		project:       project,
		network:       network,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/networks/{network}",
		context_:      googleapi.NoContext,
	}
}
func (c *NetworksGetCall) Context(ctx context.Context) *NetworksGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NetworksGetCall) Fields(s ...googleapi.Field) *NetworksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NetworksGetCall) Do() (*Network, error) {
	var returnValue *Network
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"network": c.network,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified network resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.networks.get",
	//   "parameterOrder": [
	//     "project",
	//     "network"
	//   ],
	//   "parameters": {
	//     "network": {
	//       "description": "Name of the network resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/networks/{network}",
	//   "response": {
	//     "$ref": "Network"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.networks.insert":

type NetworksInsertCall struct {
	s             *Service
	project       string
	network       *Network
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a network resource in the specified project using the
// data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/networks/insert
func (r *NetworksService) Insert(project string, network *Network) *NetworksInsertCall {
	return &NetworksInsertCall{
		s:             r.s,
		project:       project,
		network:       network,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/networks",
		context_:      googleapi.NoContext,
	}
}
func (c *NetworksInsertCall) Context(ctx context.Context) *NetworksInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NetworksInsertCall) Fields(s ...googleapi.Field) *NetworksInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NetworksInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.network,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a network resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.networks.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/networks",
	//   "request": {
	//     "$ref": "Network"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.networks.list":

type NetworksListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of network resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/networks/list
func (r *NetworksService) List(project string) *NetworksListCall {
	return &NetworksListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/networks",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *NetworksListCall) Filter(filter string) *NetworksListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *NetworksListCall) MaxResults(maxResults int64) *NetworksListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *NetworksListCall) PageToken(pageToken string) *NetworksListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *NetworksListCall) Context(ctx context.Context) *NetworksListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NetworksListCall) Fields(s ...googleapi.Field) *NetworksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NetworksListCall) Do() (*NetworkList, error) {
	var returnValue *NetworkList
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
	//   "description": "Retrieves the list of network resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.networks.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/networks",
	//   "response": {
	//     "$ref": "NetworkList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.projects.get":

type ProjectsGetCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified project resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/projects/get
func (r *ProjectsService) Get(project string) *ProjectsGetCall {
	return &ProjectsGetCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsGetCall) Context(ctx context.Context) *ProjectsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsGetCall) Do() (*Project, error) {
	var returnValue *Project
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
	//   "description": "Returns the specified project resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.projects.get",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}",
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.projects.moveDisk":

type ProjectsMoveDiskCall struct {
	s               *Service
	project         string
	diskmoverequest *DiskMoveRequest
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// MoveDisk: Moves a persistent disk from one zone to another.

func (r *ProjectsService) MoveDisk(project string, diskmoverequest *DiskMoveRequest) *ProjectsMoveDiskCall {
	return &ProjectsMoveDiskCall{
		s:               r.s,
		project:         project,
		diskmoverequest: diskmoverequest,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/moveDisk",
		context_:        googleapi.NoContext,
	}
}
func (c *ProjectsMoveDiskCall) Context(ctx context.Context) *ProjectsMoveDiskCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMoveDiskCall) Fields(s ...googleapi.Field) *ProjectsMoveDiskCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMoveDiskCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.diskmoverequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Moves a persistent disk from one zone to another.",
	//   "httpMethod": "POST",
	//   "id": "compute.projects.moveDisk",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/moveDisk",
	//   "request": {
	//     "$ref": "DiskMoveRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.projects.moveInstance":

type ProjectsMoveInstanceCall struct {
	s                   *Service
	project             string
	instancemoverequest *InstanceMoveRequest
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
	context_            context.Context
}

// MoveInstance: Moves an instance and its attached persistent disks
// from one zone to another.

func (r *ProjectsService) MoveInstance(project string, instancemoverequest *InstanceMoveRequest) *ProjectsMoveInstanceCall {
	return &ProjectsMoveInstanceCall{
		s:                   r.s,
		project:             project,
		instancemoverequest: instancemoverequest,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "{project}/moveInstance",
		context_:            googleapi.NoContext,
	}
}
func (c *ProjectsMoveInstanceCall) Context(ctx context.Context) *ProjectsMoveInstanceCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMoveInstanceCall) Fields(s ...googleapi.Field) *ProjectsMoveInstanceCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMoveInstanceCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancemoverequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Moves an instance and its attached persistent disks from one zone to another.",
	//   "httpMethod": "POST",
	//   "id": "compute.projects.moveInstance",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/moveInstance",
	//   "request": {
	//     "$ref": "InstanceMoveRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.projects.setCommonInstanceMetadata":

type ProjectsSetCommonInstanceMetadataCall struct {
	s             *Service
	project       string
	metadata      *Metadata
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// SetCommonInstanceMetadata: Sets metadata common to all instances
// within the specified project using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata
func (r *ProjectsService) SetCommonInstanceMetadata(project string, metadata *Metadata) *ProjectsSetCommonInstanceMetadataCall {
	return &ProjectsSetCommonInstanceMetadataCall{
		s:             r.s,
		project:       project,
		metadata:      metadata,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/setCommonInstanceMetadata",
		context_:      googleapi.NoContext,
	}
}
func (c *ProjectsSetCommonInstanceMetadataCall) Context(ctx context.Context) *ProjectsSetCommonInstanceMetadataCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSetCommonInstanceMetadataCall) Fields(s ...googleapi.Field) *ProjectsSetCommonInstanceMetadataCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSetCommonInstanceMetadataCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.metadata,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Sets metadata common to all instances within the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.projects.setCommonInstanceMetadata",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/setCommonInstanceMetadata",
	//   "request": {
	//     "$ref": "Metadata"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.projects.setUsageExportBucket":

type ProjectsSetUsageExportBucketCall struct {
	s                   *Service
	project             string
	usageexportlocation *UsageExportLocation
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
	context_            context.Context
}

// SetUsageExportBucket: Enables the usage export feature and sets the
// usage export bucket where reports are stored. If you provide an empty
// request body using this method, the usage export feature will be
// disabled.

// For details, see https://cloud.google.com/compute/docs/reference/latest/projects/setUsageExportBucket
func (r *ProjectsService) SetUsageExportBucket(project string, usageexportlocation *UsageExportLocation) *ProjectsSetUsageExportBucketCall {
	return &ProjectsSetUsageExportBucketCall{
		s:                   r.s,
		project:             project,
		usageexportlocation: usageexportlocation,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "{project}/setUsageExportBucket",
		context_:            googleapi.NoContext,
	}
}
func (c *ProjectsSetUsageExportBucketCall) Context(ctx context.Context) *ProjectsSetUsageExportBucketCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSetUsageExportBucketCall) Fields(s ...googleapi.Field) *ProjectsSetUsageExportBucketCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSetUsageExportBucketCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.usageexportlocation,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Enables the usage export feature and sets the usage export bucket where reports are stored. If you provide an empty request body using this method, the usage export feature will be disabled.",
	//   "httpMethod": "POST",
	//   "id": "compute.projects.setUsageExportBucket",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/setUsageExportBucket",
	//   "request": {
	//     "$ref": "UsageExportLocation"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "compute.regionOperations.delete":

type RegionOperationsDeleteCall struct {
	s             *Service
	project       string
	region        string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified region-specific operation resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/regionOperations/delete
func (r *RegionOperationsService) Delete(project string, region string, operation string) *RegionOperationsDeleteCall {
	return &RegionOperationsDeleteCall{
		s:             r.s,
		project:       project,
		region:        region,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *RegionOperationsDeleteCall) Context(ctx context.Context) *RegionOperationsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *RegionOperationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"region":    c.region,
		"operation": c.operation,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified region-specific operation resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.regionOperations.delete",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/operations/{operation}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.regionOperations.get":

type RegionOperationsGetCall struct {
	s             *Service
	project       string
	region        string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the specified region-specific operation resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/regionOperations/get
func (r *RegionOperationsService) Get(project string, region string, operation string) *RegionOperationsGetCall {
	return &RegionOperationsGetCall{
		s:             r.s,
		project:       project,
		region:        region,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *RegionOperationsGetCall) Context(ctx context.Context) *RegionOperationsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionOperationsGetCall) Fields(s ...googleapi.Field) *RegionOperationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionOperationsGetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"region":    c.region,
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
	//   "description": "Retrieves the specified region-specific operation resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.regionOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.regionOperations.list":

type RegionOperationsListCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of operation resources contained within the
// specified region.

// For details, see https://cloud.google.com/compute/docs/reference/latest/regionOperations/list
func (r *RegionOperationsService) List(project string, region string) *RegionOperationsListCall {
	return &RegionOperationsListCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/operations",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RegionOperationsListCall) Filter(filter string) *RegionOperationsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RegionOperationsListCall) MaxResults(maxResults int64) *RegionOperationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RegionOperationsListCall) PageToken(pageToken string) *RegionOperationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *RegionOperationsListCall) Context(ctx context.Context) *RegionOperationsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionOperationsListCall) Fields(s ...googleapi.Field) *RegionOperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionOperationsListCall) Do() (*OperationList, error) {
	var returnValue *OperationList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of operation resources contained within the specified region.",
	//   "httpMethod": "GET",
	//   "id": "compute.regionOperations.list",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.regions.get":

type RegionsGetCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified region resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/regions/get
func (r *RegionsService) Get(project string, region string) *RegionsGetCall {
	return &RegionsGetCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}",
		context_:      googleapi.NoContext,
	}
}
func (c *RegionsGetCall) Context(ctx context.Context) *RegionsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionsGetCall) Fields(s ...googleapi.Field) *RegionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionsGetCall) Do() (*Region, error) {
	var returnValue *Region
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified region resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.regions.get",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}",
	//   "response": {
	//     "$ref": "Region"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.regions.list":

type RegionsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of region resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/regions/list
func (r *RegionsService) List(project string) *RegionsListCall {
	return &RegionsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RegionsListCall) Filter(filter string) *RegionsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RegionsListCall) MaxResults(maxResults int64) *RegionsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RegionsListCall) PageToken(pageToken string) *RegionsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *RegionsListCall) Context(ctx context.Context) *RegionsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionsListCall) Fields(s ...googleapi.Field) *RegionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionsListCall) Do() (*RegionList, error) {
	var returnValue *RegionList
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
	//   "description": "Retrieves the list of region resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.regions.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions",
	//   "response": {
	//     "$ref": "RegionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.routes.delete":

type RoutesDeleteCall struct {
	s             *Service
	project       string
	route         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified route resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/routes/delete
func (r *RoutesService) Delete(project string, route string) *RoutesDeleteCall {
	return &RoutesDeleteCall{
		s:             r.s,
		project:       project,
		route:         route,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/routes/{route}",
		context_:      googleapi.NoContext,
	}
}
func (c *RoutesDeleteCall) Context(ctx context.Context) *RoutesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RoutesDeleteCall) Fields(s ...googleapi.Field) *RoutesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RoutesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"route":   c.route,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified route resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.routes.delete",
	//   "parameterOrder": [
	//     "project",
	//     "route"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "route": {
	//       "description": "Name of the route resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/routes/{route}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.routes.get":

type RoutesGetCall struct {
	s             *Service
	project       string
	route         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified route resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/routes/get
func (r *RoutesService) Get(project string, route string) *RoutesGetCall {
	return &RoutesGetCall{
		s:             r.s,
		project:       project,
		route:         route,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/routes/{route}",
		context_:      googleapi.NoContext,
	}
}
func (c *RoutesGetCall) Context(ctx context.Context) *RoutesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RoutesGetCall) Fields(s ...googleapi.Field) *RoutesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RoutesGetCall) Do() (*Route, error) {
	var returnValue *Route
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"route":   c.route,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified route resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.routes.get",
	//   "parameterOrder": [
	//     "project",
	//     "route"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "route": {
	//       "description": "Name of the route resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/routes/{route}",
	//   "response": {
	//     "$ref": "Route"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.routes.insert":

type RoutesInsertCall struct {
	s             *Service
	project       string
	route         *Route
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a route resource in the specified project using the
// data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/routes/insert
func (r *RoutesService) Insert(project string, route *Route) *RoutesInsertCall {
	return &RoutesInsertCall{
		s:             r.s,
		project:       project,
		route:         route,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/routes",
		context_:      googleapi.NoContext,
	}
}
func (c *RoutesInsertCall) Context(ctx context.Context) *RoutesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RoutesInsertCall) Fields(s ...googleapi.Field) *RoutesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RoutesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.route,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a route resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.routes.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/routes",
	//   "request": {
	//     "$ref": "Route"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.routes.list":

type RoutesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of route resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/routes/list
func (r *RoutesService) List(project string) *RoutesListCall {
	return &RoutesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/routes",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RoutesListCall) Filter(filter string) *RoutesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RoutesListCall) MaxResults(maxResults int64) *RoutesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RoutesListCall) PageToken(pageToken string) *RoutesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *RoutesListCall) Context(ctx context.Context) *RoutesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RoutesListCall) Fields(s ...googleapi.Field) *RoutesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RoutesListCall) Do() (*RouteList, error) {
	var returnValue *RouteList
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
	//   "description": "Retrieves the list of route resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.routes.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/routes",
	//   "response": {
	//     "$ref": "RouteList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.snapshots.delete":

type SnapshotsDeleteCall struct {
	s             *Service
	project       string
	snapshot      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified persistent disk snapshot resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/snapshots/delete
func (r *SnapshotsService) Delete(project string, snapshot string) *SnapshotsDeleteCall {
	return &SnapshotsDeleteCall{
		s:             r.s,
		project:       project,
		snapshot:      snapshot,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/snapshots/{snapshot}",
		context_:      googleapi.NoContext,
	}
}
func (c *SnapshotsDeleteCall) Context(ctx context.Context) *SnapshotsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SnapshotsDeleteCall) Fields(s ...googleapi.Field) *SnapshotsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SnapshotsDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"snapshot": c.snapshot,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified persistent disk snapshot resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.snapshots.delete",
	//   "parameterOrder": [
	//     "project",
	//     "snapshot"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshot": {
	//       "description": "Name of the persistent disk snapshot resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/snapshots/{snapshot}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.snapshots.get":

type SnapshotsGetCall struct {
	s             *Service
	project       string
	snapshot      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified persistent disk snapshot resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/snapshots/get
func (r *SnapshotsService) Get(project string, snapshot string) *SnapshotsGetCall {
	return &SnapshotsGetCall{
		s:             r.s,
		project:       project,
		snapshot:      snapshot,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/snapshots/{snapshot}",
		context_:      googleapi.NoContext,
	}
}
func (c *SnapshotsGetCall) Context(ctx context.Context) *SnapshotsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SnapshotsGetCall) Fields(s ...googleapi.Field) *SnapshotsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SnapshotsGetCall) Do() (*Snapshot, error) {
	var returnValue *Snapshot
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":  c.project,
		"snapshot": c.snapshot,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified persistent disk snapshot resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.snapshots.get",
	//   "parameterOrder": [
	//     "project",
	//     "snapshot"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshot": {
	//       "description": "Name of the persistent disk snapshot resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/snapshots/{snapshot}",
	//   "response": {
	//     "$ref": "Snapshot"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.snapshots.list":

type SnapshotsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of persistent disk snapshot resources
// contained within the specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/snapshots/list
func (r *SnapshotsService) List(project string) *SnapshotsListCall {
	return &SnapshotsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/snapshots",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *SnapshotsListCall) Filter(filter string) *SnapshotsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *SnapshotsListCall) MaxResults(maxResults int64) *SnapshotsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *SnapshotsListCall) PageToken(pageToken string) *SnapshotsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *SnapshotsListCall) Context(ctx context.Context) *SnapshotsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SnapshotsListCall) Fields(s ...googleapi.Field) *SnapshotsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SnapshotsListCall) Do() (*SnapshotList, error) {
	var returnValue *SnapshotList
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
	//   "description": "Retrieves the list of persistent disk snapshot resources contained within the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.snapshots.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/snapshots",
	//   "response": {
	//     "$ref": "SnapshotList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetHttpProxies.delete":

type TargetHttpProxiesDeleteCall struct {
	s               *Service
	project         string
	targetHttpProxy string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Delete: Deletes the specified TargetHttpProxy resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetHttpProxies/delete
func (r *TargetHttpProxiesService) Delete(project string, targetHttpProxy string) *TargetHttpProxiesDeleteCall {
	return &TargetHttpProxiesDeleteCall{
		s:               r.s,
		project:         project,
		targetHttpProxy: targetHttpProxy,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/targetHttpProxies/{targetHttpProxy}",
		context_:        googleapi.NoContext,
	}
}
func (c *TargetHttpProxiesDeleteCall) Context(ctx context.Context) *TargetHttpProxiesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetHttpProxiesDeleteCall) Fields(s ...googleapi.Field) *TargetHttpProxiesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetHttpProxiesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"targetHttpProxy": c.targetHttpProxy,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified TargetHttpProxy resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.targetHttpProxies.delete",
	//   "parameterOrder": [
	//     "project",
	//     "targetHttpProxy"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetHttpProxy": {
	//       "description": "Name of the TargetHttpProxy resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/targetHttpProxies/{targetHttpProxy}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetHttpProxies.get":

type TargetHttpProxiesGetCall struct {
	s               *Service
	project         string
	targetHttpProxy string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Get: Returns the specified TargetHttpProxy resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetHttpProxies/get
func (r *TargetHttpProxiesService) Get(project string, targetHttpProxy string) *TargetHttpProxiesGetCall {
	return &TargetHttpProxiesGetCall{
		s:               r.s,
		project:         project,
		targetHttpProxy: targetHttpProxy,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/targetHttpProxies/{targetHttpProxy}",
		context_:        googleapi.NoContext,
	}
}
func (c *TargetHttpProxiesGetCall) Context(ctx context.Context) *TargetHttpProxiesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetHttpProxiesGetCall) Fields(s ...googleapi.Field) *TargetHttpProxiesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetHttpProxiesGetCall) Do() (*TargetHttpProxy, error) {
	var returnValue *TargetHttpProxy
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"targetHttpProxy": c.targetHttpProxy,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified TargetHttpProxy resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetHttpProxies.get",
	//   "parameterOrder": [
	//     "project",
	//     "targetHttpProxy"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetHttpProxy": {
	//       "description": "Name of the TargetHttpProxy resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/targetHttpProxies/{targetHttpProxy}",
	//   "response": {
	//     "$ref": "TargetHttpProxy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetHttpProxies.insert":

type TargetHttpProxiesInsertCall struct {
	s               *Service
	project         string
	targethttpproxy *TargetHttpProxy
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Insert: Creates a TargetHttpProxy resource in the specified project
// using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetHttpProxies/insert
func (r *TargetHttpProxiesService) Insert(project string, targethttpproxy *TargetHttpProxy) *TargetHttpProxiesInsertCall {
	return &TargetHttpProxiesInsertCall{
		s:               r.s,
		project:         project,
		targethttpproxy: targethttpproxy,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/global/targetHttpProxies",
		context_:        googleapi.NoContext,
	}
}
func (c *TargetHttpProxiesInsertCall) Context(ctx context.Context) *TargetHttpProxiesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetHttpProxiesInsertCall) Fields(s ...googleapi.Field) *TargetHttpProxiesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetHttpProxiesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targethttpproxy,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a TargetHttpProxy resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetHttpProxies.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/targetHttpProxies",
	//   "request": {
	//     "$ref": "TargetHttpProxy"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetHttpProxies.list":

type TargetHttpProxiesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of TargetHttpProxy resources available to
// the specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetHttpProxies/list
func (r *TargetHttpProxiesService) List(project string) *TargetHttpProxiesListCall {
	return &TargetHttpProxiesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/targetHttpProxies",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetHttpProxiesListCall) Filter(filter string) *TargetHttpProxiesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetHttpProxiesListCall) MaxResults(maxResults int64) *TargetHttpProxiesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetHttpProxiesListCall) PageToken(pageToken string) *TargetHttpProxiesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetHttpProxiesListCall) Context(ctx context.Context) *TargetHttpProxiesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetHttpProxiesListCall) Fields(s ...googleapi.Field) *TargetHttpProxiesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetHttpProxiesListCall) Do() (*TargetHttpProxyList, error) {
	var returnValue *TargetHttpProxyList
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
	//   "description": "Retrieves the list of TargetHttpProxy resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetHttpProxies.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/targetHttpProxies",
	//   "response": {
	//     "$ref": "TargetHttpProxyList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetHttpProxies.setUrlMap":

type TargetHttpProxiesSetUrlMapCall struct {
	s               *Service
	project         string
	targetHttpProxy string
	urlmapreference *UrlMapReference
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// SetUrlMap: Changes the URL map for TargetHttpProxy.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetHttpProxies/setUrlMap
func (r *TargetHttpProxiesService) SetUrlMap(project string, targetHttpProxy string, urlmapreference *UrlMapReference) *TargetHttpProxiesSetUrlMapCall {
	return &TargetHttpProxiesSetUrlMapCall{
		s:               r.s,
		project:         project,
		targetHttpProxy: targetHttpProxy,
		urlmapreference: urlmapreference,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/targetHttpProxies/{targetHttpProxy}/setUrlMap",
		context_:        googleapi.NoContext,
	}
}
func (c *TargetHttpProxiesSetUrlMapCall) Context(ctx context.Context) *TargetHttpProxiesSetUrlMapCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetHttpProxiesSetUrlMapCall) Fields(s ...googleapi.Field) *TargetHttpProxiesSetUrlMapCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetHttpProxiesSetUrlMapCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":         c.project,
		"targetHttpProxy": c.targetHttpProxy,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.urlmapreference,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Changes the URL map for TargetHttpProxy.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetHttpProxies.setUrlMap",
	//   "parameterOrder": [
	//     "project",
	//     "targetHttpProxy"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetHttpProxy": {
	//       "description": "Name of the TargetHttpProxy resource whose URL map is to be set.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/targetHttpProxies/{targetHttpProxy}/setUrlMap",
	//   "request": {
	//     "$ref": "UrlMapReference"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetInstances.aggregatedList":

type TargetInstancesAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of target instances grouped by
// scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetInstances/aggregatedList
func (r *TargetInstancesService) AggregatedList(project string) *TargetInstancesAggregatedListCall {
	return &TargetInstancesAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/targetInstances",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetInstancesAggregatedListCall) Filter(filter string) *TargetInstancesAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetInstancesAggregatedListCall) MaxResults(maxResults int64) *TargetInstancesAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetInstancesAggregatedListCall) PageToken(pageToken string) *TargetInstancesAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetInstancesAggregatedListCall) Context(ctx context.Context) *TargetInstancesAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetInstancesAggregatedListCall) Fields(s ...googleapi.Field) *TargetInstancesAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetInstancesAggregatedListCall) Do() (*TargetInstanceAggregatedList, error) {
	var returnValue *TargetInstanceAggregatedList
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
	//   "description": "Retrieves the list of target instances grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetInstances.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/targetInstances",
	//   "response": {
	//     "$ref": "TargetInstanceAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetInstances.delete":

type TargetInstancesDeleteCall struct {
	s              *Service
	project        string
	zone           string
	targetInstance string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Delete: Deletes the specified TargetInstance resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetInstances/delete
func (r *TargetInstancesService) Delete(project string, zone string, targetInstance string) *TargetInstancesDeleteCall {
	return &TargetInstancesDeleteCall{
		s:              r.s,
		project:        project,
		zone:           zone,
		targetInstance: targetInstance,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/zones/{zone}/targetInstances/{targetInstance}",
		context_:       googleapi.NoContext,
	}
}
func (c *TargetInstancesDeleteCall) Context(ctx context.Context) *TargetInstancesDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetInstancesDeleteCall) Fields(s ...googleapi.Field) *TargetInstancesDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetInstancesDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"zone":           c.zone,
		"targetInstance": c.targetInstance,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified TargetInstance resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.targetInstances.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "targetInstance"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetInstance": {
	//       "description": "Name of the TargetInstance resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/targetInstances/{targetInstance}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetInstances.get":

type TargetInstancesGetCall struct {
	s              *Service
	project        string
	zone           string
	targetInstance string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Get: Returns the specified TargetInstance resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetInstances/get
func (r *TargetInstancesService) Get(project string, zone string, targetInstance string) *TargetInstancesGetCall {
	return &TargetInstancesGetCall{
		s:              r.s,
		project:        project,
		zone:           zone,
		targetInstance: targetInstance,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/zones/{zone}/targetInstances/{targetInstance}",
		context_:       googleapi.NoContext,
	}
}
func (c *TargetInstancesGetCall) Context(ctx context.Context) *TargetInstancesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetInstancesGetCall) Fields(s ...googleapi.Field) *TargetInstancesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetInstancesGetCall) Do() (*TargetInstance, error) {
	var returnValue *TargetInstance
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":        c.project,
		"zone":           c.zone,
		"targetInstance": c.targetInstance,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified TargetInstance resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetInstances.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "targetInstance"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetInstance": {
	//       "description": "Name of the TargetInstance resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/targetInstances/{targetInstance}",
	//   "response": {
	//     "$ref": "TargetInstance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetInstances.insert":

type TargetInstancesInsertCall struct {
	s              *Service
	project        string
	zone           string
	targetinstance *TargetInstance
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
	context_       context.Context
}

// Insert: Creates a TargetInstance resource in the specified project
// and zone using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetInstances/insert
func (r *TargetInstancesService) Insert(project string, zone string, targetinstance *TargetInstance) *TargetInstancesInsertCall {
	return &TargetInstancesInsertCall{
		s:              r.s,
		project:        project,
		zone:           zone,
		targetinstance: targetinstance,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{project}/zones/{zone}/targetInstances",
		context_:       googleapi.NoContext,
	}
}
func (c *TargetInstancesInsertCall) Context(ctx context.Context) *TargetInstancesInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetInstancesInsertCall) Fields(s ...googleapi.Field) *TargetInstancesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetInstancesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetinstance,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a TargetInstance resource in the specified project and zone using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetInstances.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/targetInstances",
	//   "request": {
	//     "$ref": "TargetInstance"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetInstances.list":

type TargetInstancesListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of TargetInstance resources available to the
// specified project and zone.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetInstances/list
func (r *TargetInstancesService) List(project string, zone string) *TargetInstancesListCall {
	return &TargetInstancesListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/targetInstances",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetInstancesListCall) Filter(filter string) *TargetInstancesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetInstancesListCall) MaxResults(maxResults int64) *TargetInstancesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetInstancesListCall) PageToken(pageToken string) *TargetInstancesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetInstancesListCall) Context(ctx context.Context) *TargetInstancesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetInstancesListCall) Fields(s ...googleapi.Field) *TargetInstancesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetInstancesListCall) Do() (*TargetInstanceList, error) {
	var returnValue *TargetInstanceList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of TargetInstance resources available to the specified project and zone.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetInstances.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/targetInstances",
	//   "response": {
	//     "$ref": "TargetInstanceList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetPools.addHealthCheck":

type TargetPoolsAddHealthCheckCall struct {
	s                                *Service
	project                          string
	region                           string
	targetPool                       string
	targetpoolsaddhealthcheckrequest *TargetPoolsAddHealthCheckRequest
	caller_                          googleapi.Caller
	params_                          url.Values
	pathTemplate_                    string
	context_                         context.Context
}

// AddHealthCheck: Adds health check URL to targetPool.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/addHealthCheck
func (r *TargetPoolsService) AddHealthCheck(project string, region string, targetPool string, targetpoolsaddhealthcheckrequest *TargetPoolsAddHealthCheckRequest) *TargetPoolsAddHealthCheckCall {
	return &TargetPoolsAddHealthCheckCall{
		s:                                r.s,
		project:                          project,
		region:                           region,
		targetPool:                       targetPool,
		targetpoolsaddhealthcheckrequest: targetpoolsaddhealthcheckrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools/{targetPool}/addHealthCheck",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsAddHealthCheckCall) Context(ctx context.Context) *TargetPoolsAddHealthCheckCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsAddHealthCheckCall) Fields(s ...googleapi.Field) *TargetPoolsAddHealthCheckCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsAddHealthCheckCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetpoolsaddhealthcheckrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Adds health check URL to targetPool.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.addHealthCheck",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to which health_check_url is to be added.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}/addHealthCheck",
	//   "request": {
	//     "$ref": "TargetPoolsAddHealthCheckRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetPools.addInstance":

type TargetPoolsAddInstanceCall struct {
	s                             *Service
	project                       string
	region                        string
	targetPool                    string
	targetpoolsaddinstancerequest *TargetPoolsAddInstanceRequest
	caller_                       googleapi.Caller
	params_                       url.Values
	pathTemplate_                 string
	context_                      context.Context
}

// AddInstance: Adds instance url to targetPool.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/addInstance
func (r *TargetPoolsService) AddInstance(project string, region string, targetPool string, targetpoolsaddinstancerequest *TargetPoolsAddInstanceRequest) *TargetPoolsAddInstanceCall {
	return &TargetPoolsAddInstanceCall{
		s:                             r.s,
		project:                       project,
		region:                        region,
		targetPool:                    targetPool,
		targetpoolsaddinstancerequest: targetpoolsaddinstancerequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools/{targetPool}/addInstance",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsAddInstanceCall) Context(ctx context.Context) *TargetPoolsAddInstanceCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsAddInstanceCall) Fields(s ...googleapi.Field) *TargetPoolsAddInstanceCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsAddInstanceCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetpoolsaddinstancerequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Adds instance url to targetPool.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.addInstance",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to which instance_url is to be added.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}/addInstance",
	//   "request": {
	//     "$ref": "TargetPoolsAddInstanceRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetPools.aggregatedList":

type TargetPoolsAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of target pools grouped by scope.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/aggregatedList
func (r *TargetPoolsService) AggregatedList(project string) *TargetPoolsAggregatedListCall {
	return &TargetPoolsAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/targetPools",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetPoolsAggregatedListCall) Filter(filter string) *TargetPoolsAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetPoolsAggregatedListCall) MaxResults(maxResults int64) *TargetPoolsAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetPoolsAggregatedListCall) PageToken(pageToken string) *TargetPoolsAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetPoolsAggregatedListCall) Context(ctx context.Context) *TargetPoolsAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsAggregatedListCall) Fields(s ...googleapi.Field) *TargetPoolsAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsAggregatedListCall) Do() (*TargetPoolAggregatedList, error) {
	var returnValue *TargetPoolAggregatedList
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
	//   "description": "Retrieves the list of target pools grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetPools.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/targetPools",
	//   "response": {
	//     "$ref": "TargetPoolAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetPools.delete":

type TargetPoolsDeleteCall struct {
	s             *Service
	project       string
	region        string
	targetPool    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified TargetPool resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/delete
func (r *TargetPoolsService) Delete(project string, region string, targetPool string) *TargetPoolsDeleteCall {
	return &TargetPoolsDeleteCall{
		s:             r.s,
		project:       project,
		region:        region,
		targetPool:    targetPool,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools/{targetPool}",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsDeleteCall) Context(ctx context.Context) *TargetPoolsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsDeleteCall) Fields(s ...googleapi.Field) *TargetPoolsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified TargetPool resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.targetPools.delete",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetPools.get":

type TargetPoolsGetCall struct {
	s             *Service
	project       string
	region        string
	targetPool    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified TargetPool resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/get
func (r *TargetPoolsService) Get(project string, region string, targetPool string) *TargetPoolsGetCall {
	return &TargetPoolsGetCall{
		s:             r.s,
		project:       project,
		region:        region,
		targetPool:    targetPool,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools/{targetPool}",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsGetCall) Context(ctx context.Context) *TargetPoolsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsGetCall) Fields(s ...googleapi.Field) *TargetPoolsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsGetCall) Do() (*TargetPool, error) {
	var returnValue *TargetPool
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified TargetPool resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetPools.get",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}",
	//   "response": {
	//     "$ref": "TargetPool"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetPools.getHealth":

type TargetPoolsGetHealthCall struct {
	s                 *Service
	project           string
	region            string
	targetPool        string
	instancereference *InstanceReference
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
	context_          context.Context
}

// GetHealth: Gets the most recent health check results for each IP for
// the given instance that is referenced by given TargetPool.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/getHealth
func (r *TargetPoolsService) GetHealth(project string, region string, targetPool string, instancereference *InstanceReference) *TargetPoolsGetHealthCall {
	return &TargetPoolsGetHealthCall{
		s:                 r.s,
		project:           project,
		region:            region,
		targetPool:        targetPool,
		instancereference: instancereference,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "{project}/regions/{region}/targetPools/{targetPool}/getHealth",
		context_:          googleapi.NoContext,
	}
}
func (c *TargetPoolsGetHealthCall) Context(ctx context.Context) *TargetPoolsGetHealthCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsGetHealthCall) Fields(s ...googleapi.Field) *TargetPoolsGetHealthCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsGetHealthCall) Do() (*TargetPoolInstanceHealth, error) {
	var returnValue *TargetPoolInstanceHealth
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancereference,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Gets the most recent health check results for each IP for the given instance that is referenced by given TargetPool.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.getHealth",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to which the queried instance belongs.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}/getHealth",
	//   "request": {
	//     "$ref": "InstanceReference"
	//   },
	//   "response": {
	//     "$ref": "TargetPoolInstanceHealth"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetPools.insert":

type TargetPoolsInsertCall struct {
	s             *Service
	project       string
	region        string
	targetpool    *TargetPool
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a TargetPool resource in the specified project and
// region using the data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/insert
func (r *TargetPoolsService) Insert(project string, region string, targetpool *TargetPool) *TargetPoolsInsertCall {
	return &TargetPoolsInsertCall{
		s:             r.s,
		project:       project,
		region:        region,
		targetpool:    targetpool,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsInsertCall) Context(ctx context.Context) *TargetPoolsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsInsertCall) Fields(s ...googleapi.Field) *TargetPoolsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetpool,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a TargetPool resource in the specified project and region using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.insert",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools",
	//   "request": {
	//     "$ref": "TargetPool"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetPools.list":

type TargetPoolsListCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of TargetPool resources available to the
// specified project and region.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/list
func (r *TargetPoolsService) List(project string, region string) *TargetPoolsListCall {
	return &TargetPoolsListCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetPoolsListCall) Filter(filter string) *TargetPoolsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetPoolsListCall) MaxResults(maxResults int64) *TargetPoolsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetPoolsListCall) PageToken(pageToken string) *TargetPoolsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetPoolsListCall) Context(ctx context.Context) *TargetPoolsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsListCall) Fields(s ...googleapi.Field) *TargetPoolsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsListCall) Do() (*TargetPoolList, error) {
	var returnValue *TargetPoolList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of TargetPool resources available to the specified project and region.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetPools.list",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools",
	//   "response": {
	//     "$ref": "TargetPoolList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetPools.removeHealthCheck":

type TargetPoolsRemoveHealthCheckCall struct {
	s                                   *Service
	project                             string
	region                              string
	targetPool                          string
	targetpoolsremovehealthcheckrequest *TargetPoolsRemoveHealthCheckRequest
	caller_                             googleapi.Caller
	params_                             url.Values
	pathTemplate_                       string
	context_                            context.Context
}

// RemoveHealthCheck: Removes health check URL from targetPool.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/removeHealthCheck
func (r *TargetPoolsService) RemoveHealthCheck(project string, region string, targetPool string, targetpoolsremovehealthcheckrequest *TargetPoolsRemoveHealthCheckRequest) *TargetPoolsRemoveHealthCheckCall {
	return &TargetPoolsRemoveHealthCheckCall{
		s:                                   r.s,
		project:                             project,
		region:                              region,
		targetPool:                          targetPool,
		targetpoolsremovehealthcheckrequest: targetpoolsremovehealthcheckrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools/{targetPool}/removeHealthCheck",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsRemoveHealthCheckCall) Context(ctx context.Context) *TargetPoolsRemoveHealthCheckCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsRemoveHealthCheckCall) Fields(s ...googleapi.Field) *TargetPoolsRemoveHealthCheckCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsRemoveHealthCheckCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetpoolsremovehealthcheckrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Removes health check URL from targetPool.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.removeHealthCheck",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to which health_check_url is to be removed.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}/removeHealthCheck",
	//   "request": {
	//     "$ref": "TargetPoolsRemoveHealthCheckRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetPools.removeInstance":

type TargetPoolsRemoveInstanceCall struct {
	s                                *Service
	project                          string
	region                           string
	targetPool                       string
	targetpoolsremoveinstancerequest *TargetPoolsRemoveInstanceRequest
	caller_                          googleapi.Caller
	params_                          url.Values
	pathTemplate_                    string
	context_                         context.Context
}

// RemoveInstance: Removes instance URL from targetPool.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/removeInstance
func (r *TargetPoolsService) RemoveInstance(project string, region string, targetPool string, targetpoolsremoveinstancerequest *TargetPoolsRemoveInstanceRequest) *TargetPoolsRemoveInstanceCall {
	return &TargetPoolsRemoveInstanceCall{
		s:                                r.s,
		project:                          project,
		region:                           region,
		targetPool:                       targetPool,
		targetpoolsremoveinstancerequest: targetpoolsremoveinstancerequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetPools/{targetPool}/removeInstance",
		context_:      googleapi.NoContext,
	}
}
func (c *TargetPoolsRemoveInstanceCall) Context(ctx context.Context) *TargetPoolsRemoveInstanceCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsRemoveInstanceCall) Fields(s ...googleapi.Field) *TargetPoolsRemoveInstanceCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsRemoveInstanceCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetpoolsremoveinstancerequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Removes instance URL from targetPool.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.removeInstance",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource to which instance_url is to be removed.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}/removeInstance",
	//   "request": {
	//     "$ref": "TargetPoolsRemoveInstanceRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetPools.setBackup":

type TargetPoolsSetBackupCall struct {
	s               *Service
	project         string
	region          string
	targetPool      string
	targetreference *TargetReference
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// SetBackup: Changes backup pool configurations.

// For details, see https://cloud.google.com/compute/docs/reference/latest/targetPools/setBackup
func (r *TargetPoolsService) SetBackup(project string, region string, targetPool string, targetreference *TargetReference) *TargetPoolsSetBackupCall {
	return &TargetPoolsSetBackupCall{
		s:               r.s,
		project:         project,
		region:          region,
		targetPool:      targetPool,
		targetreference: targetreference,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/regions/{region}/targetPools/{targetPool}/setBackup",
		context_:        googleapi.NoContext,
	}
}

// FailoverRatio sets the optional parameter "failoverRatio": New
// failoverRatio value for the containing target pool.
func (c *TargetPoolsSetBackupCall) FailoverRatio(failoverRatio float64) *TargetPoolsSetBackupCall {
	c.params_.Set("failoverRatio", fmt.Sprintf("%v", failoverRatio))
	return c
}
func (c *TargetPoolsSetBackupCall) Context(ctx context.Context) *TargetPoolsSetBackupCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetPoolsSetBackupCall) Fields(s ...googleapi.Field) *TargetPoolsSetBackupCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetPoolsSetBackupCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"region":     c.region,
		"targetPool": c.targetPool,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetreference,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Changes backup pool configurations.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetPools.setBackup",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetPool"
	//   ],
	//   "parameters": {
	//     "failoverRatio": {
	//       "description": "New failoverRatio value for the containing target pool.",
	//       "format": "float",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Name of the region scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetPool": {
	//       "description": "Name of the TargetPool resource for which the backup is to be set.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetPools/{targetPool}/setBackup",
	//   "request": {
	//     "$ref": "TargetReference"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetVpnGateways.aggregatedList":

type TargetVpnGatewaysAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of target VPN gateways grouped by
// scope.

func (r *TargetVpnGatewaysService) AggregatedList(project string) *TargetVpnGatewaysAggregatedListCall {
	return &TargetVpnGatewaysAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/targetVpnGateways",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetVpnGatewaysAggregatedListCall) Filter(filter string) *TargetVpnGatewaysAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetVpnGatewaysAggregatedListCall) MaxResults(maxResults int64) *TargetVpnGatewaysAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetVpnGatewaysAggregatedListCall) PageToken(pageToken string) *TargetVpnGatewaysAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetVpnGatewaysAggregatedListCall) Context(ctx context.Context) *TargetVpnGatewaysAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetVpnGatewaysAggregatedListCall) Fields(s ...googleapi.Field) *TargetVpnGatewaysAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetVpnGatewaysAggregatedListCall) Do() (*TargetVpnGatewayAggregatedList, error) {
	var returnValue *TargetVpnGatewayAggregatedList
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
	//   "description": "Retrieves the list of target VPN gateways grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetVpnGateways.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/targetVpnGateways",
	//   "response": {
	//     "$ref": "TargetVpnGatewayAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetVpnGateways.delete":

type TargetVpnGatewaysDeleteCall struct {
	s                *Service
	project          string
	region           string
	targetVpnGateway string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Delete: Deletes the specified TargetVpnGateway resource.

func (r *TargetVpnGatewaysService) Delete(project string, region string, targetVpnGateway string) *TargetVpnGatewaysDeleteCall {
	return &TargetVpnGatewaysDeleteCall{
		s:                r.s,
		project:          project,
		region:           region,
		targetVpnGateway: targetVpnGateway,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/regions/{region}/targetVpnGateways/{targetVpnGateway}",
		context_:         googleapi.NoContext,
	}
}
func (c *TargetVpnGatewaysDeleteCall) Context(ctx context.Context) *TargetVpnGatewaysDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetVpnGatewaysDeleteCall) Fields(s ...googleapi.Field) *TargetVpnGatewaysDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetVpnGatewaysDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":          c.project,
		"region":           c.region,
		"targetVpnGateway": c.targetVpnGateway,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified TargetVpnGateway resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.targetVpnGateways.delete",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetVpnGateway"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetVpnGateway": {
	//       "description": "Name of the TargetVpnGateway resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetVpnGateways/{targetVpnGateway}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetVpnGateways.get":

type TargetVpnGatewaysGetCall struct {
	s                *Service
	project          string
	region           string
	targetVpnGateway string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Get: Returns the specified TargetVpnGateway resource.

func (r *TargetVpnGatewaysService) Get(project string, region string, targetVpnGateway string) *TargetVpnGatewaysGetCall {
	return &TargetVpnGatewaysGetCall{
		s:                r.s,
		project:          project,
		region:           region,
		targetVpnGateway: targetVpnGateway,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/regions/{region}/targetVpnGateways/{targetVpnGateway}",
		context_:         googleapi.NoContext,
	}
}
func (c *TargetVpnGatewaysGetCall) Context(ctx context.Context) *TargetVpnGatewaysGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetVpnGatewaysGetCall) Fields(s ...googleapi.Field) *TargetVpnGatewaysGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetVpnGatewaysGetCall) Do() (*TargetVpnGateway, error) {
	var returnValue *TargetVpnGateway
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":          c.project,
		"region":           c.region,
		"targetVpnGateway": c.targetVpnGateway,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified TargetVpnGateway resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetVpnGateways.get",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "targetVpnGateway"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetVpnGateway": {
	//       "description": "Name of the TargetVpnGateway resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetVpnGateways/{targetVpnGateway}",
	//   "response": {
	//     "$ref": "TargetVpnGateway"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.targetVpnGateways.insert":

type TargetVpnGatewaysInsertCall struct {
	s                *Service
	project          string
	region           string
	targetvpngateway *TargetVpnGateway
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
	context_         context.Context
}

// Insert: Creates a TargetVpnGateway resource in the specified project
// and region using the data included in the request.

func (r *TargetVpnGatewaysService) Insert(project string, region string, targetvpngateway *TargetVpnGateway) *TargetVpnGatewaysInsertCall {
	return &TargetVpnGatewaysInsertCall{
		s:                r.s,
		project:          project,
		region:           region,
		targetvpngateway: targetvpngateway,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/regions/{region}/targetVpnGateways",
		context_:         googleapi.NoContext,
	}
}
func (c *TargetVpnGatewaysInsertCall) Context(ctx context.Context) *TargetVpnGatewaysInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetVpnGatewaysInsertCall) Fields(s ...googleapi.Field) *TargetVpnGatewaysInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetVpnGatewaysInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.targetvpngateway,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a TargetVpnGateway resource in the specified project and region using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.targetVpnGateways.insert",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetVpnGateways",
	//   "request": {
	//     "$ref": "TargetVpnGateway"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.targetVpnGateways.list":

type TargetVpnGatewaysListCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of TargetVpnGateway resources available to
// the specified project and region.

func (r *TargetVpnGatewaysService) List(project string, region string) *TargetVpnGatewaysListCall {
	return &TargetVpnGatewaysListCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/targetVpnGateways",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *TargetVpnGatewaysListCall) Filter(filter string) *TargetVpnGatewaysListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *TargetVpnGatewaysListCall) MaxResults(maxResults int64) *TargetVpnGatewaysListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *TargetVpnGatewaysListCall) PageToken(pageToken string) *TargetVpnGatewaysListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *TargetVpnGatewaysListCall) Context(ctx context.Context) *TargetVpnGatewaysListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TargetVpnGatewaysListCall) Fields(s ...googleapi.Field) *TargetVpnGatewaysListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TargetVpnGatewaysListCall) Do() (*TargetVpnGatewayList, error) {
	var returnValue *TargetVpnGatewayList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of TargetVpnGateway resources available to the specified project and region.",
	//   "httpMethod": "GET",
	//   "id": "compute.targetVpnGateways.list",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/targetVpnGateways",
	//   "response": {
	//     "$ref": "TargetVpnGatewayList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.urlMaps.delete":

type UrlMapsDeleteCall struct {
	s             *Service
	project       string
	urlMap        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified UrlMap resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/delete
func (r *UrlMapsService) Delete(project string, urlMap string) *UrlMapsDeleteCall {
	return &UrlMapsDeleteCall{
		s:             r.s,
		project:       project,
		urlMap:        urlMap,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/urlMaps/{urlMap}",
		context_:      googleapi.NoContext,
	}
}
func (c *UrlMapsDeleteCall) Context(ctx context.Context) *UrlMapsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsDeleteCall) Fields(s ...googleapi.Field) *UrlMapsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"urlMap":  c.urlMap,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified UrlMap resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.urlMaps.delete",
	//   "parameterOrder": [
	//     "project",
	//     "urlMap"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "urlMap": {
	//       "description": "Name of the UrlMap resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps/{urlMap}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.urlMaps.get":

type UrlMapsGetCall struct {
	s             *Service
	project       string
	urlMap        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified UrlMap resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/get
func (r *UrlMapsService) Get(project string, urlMap string) *UrlMapsGetCall {
	return &UrlMapsGetCall{
		s:             r.s,
		project:       project,
		urlMap:        urlMap,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/urlMaps/{urlMap}",
		context_:      googleapi.NoContext,
	}
}
func (c *UrlMapsGetCall) Context(ctx context.Context) *UrlMapsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsGetCall) Fields(s ...googleapi.Field) *UrlMapsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsGetCall) Do() (*UrlMap, error) {
	var returnValue *UrlMap
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"urlMap":  c.urlMap,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified UrlMap resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.urlMaps.get",
	//   "parameterOrder": [
	//     "project",
	//     "urlMap"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "urlMap": {
	//       "description": "Name of the UrlMap resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps/{urlMap}",
	//   "response": {
	//     "$ref": "UrlMap"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.urlMaps.insert":

type UrlMapsInsertCall struct {
	s             *Service
	project       string
	urlmap        *UrlMap
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a UrlMap resource in the specified project using the
// data included in the request.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/insert
func (r *UrlMapsService) Insert(project string, urlmap *UrlMap) *UrlMapsInsertCall {
	return &UrlMapsInsertCall{
		s:             r.s,
		project:       project,
		urlmap:        urlmap,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/urlMaps",
		context_:      googleapi.NoContext,
	}
}
func (c *UrlMapsInsertCall) Context(ctx context.Context) *UrlMapsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsInsertCall) Fields(s ...googleapi.Field) *UrlMapsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.urlmap,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a UrlMap resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.urlMaps.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps",
	//   "request": {
	//     "$ref": "UrlMap"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.urlMaps.list":

type UrlMapsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of UrlMap resources available to the
// specified project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/list
func (r *UrlMapsService) List(project string) *UrlMapsListCall {
	return &UrlMapsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/urlMaps",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *UrlMapsListCall) Filter(filter string) *UrlMapsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *UrlMapsListCall) MaxResults(maxResults int64) *UrlMapsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *UrlMapsListCall) PageToken(pageToken string) *UrlMapsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *UrlMapsListCall) Context(ctx context.Context) *UrlMapsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsListCall) Fields(s ...googleapi.Field) *UrlMapsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsListCall) Do() (*UrlMapList, error) {
	var returnValue *UrlMapList
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
	//   "description": "Retrieves the list of UrlMap resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.urlMaps.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps",
	//   "response": {
	//     "$ref": "UrlMapList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.urlMaps.patch":

type UrlMapsPatchCall struct {
	s             *Service
	project       string
	urlMap        string
	urlmap        *UrlMap
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Update the entire content of the UrlMap resource. This method
// supports patch semantics.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/patch
func (r *UrlMapsService) Patch(project string, urlMap string, urlmap *UrlMap) *UrlMapsPatchCall {
	return &UrlMapsPatchCall{
		s:             r.s,
		project:       project,
		urlMap:        urlMap,
		urlmap:        urlmap,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/urlMaps/{urlMap}",
		context_:      googleapi.NoContext,
	}
}
func (c *UrlMapsPatchCall) Context(ctx context.Context) *UrlMapsPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsPatchCall) Fields(s ...googleapi.Field) *UrlMapsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsPatchCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"urlMap":  c.urlMap,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.urlmap,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Update the entire content of the UrlMap resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "compute.urlMaps.patch",
	//   "parameterOrder": [
	//     "project",
	//     "urlMap"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "urlMap": {
	//       "description": "Name of the UrlMap resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps/{urlMap}",
	//   "request": {
	//     "$ref": "UrlMap"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.urlMaps.update":

type UrlMapsUpdateCall struct {
	s             *Service
	project       string
	urlMap        string
	urlmap        *UrlMap
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Update the entire content of the UrlMap resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/update
func (r *UrlMapsService) Update(project string, urlMap string, urlmap *UrlMap) *UrlMapsUpdateCall {
	return &UrlMapsUpdateCall{
		s:             r.s,
		project:       project,
		urlMap:        urlMap,
		urlmap:        urlmap,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/global/urlMaps/{urlMap}",
		context_:      googleapi.NoContext,
	}
}
func (c *UrlMapsUpdateCall) Context(ctx context.Context) *UrlMapsUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsUpdateCall) Fields(s ...googleapi.Field) *UrlMapsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsUpdateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"urlMap":  c.urlMap,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.urlmap,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Update the entire content of the UrlMap resource.",
	//   "httpMethod": "PUT",
	//   "id": "compute.urlMaps.update",
	//   "parameterOrder": [
	//     "project",
	//     "urlMap"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "urlMap": {
	//       "description": "Name of the UrlMap resource to update.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps/{urlMap}",
	//   "request": {
	//     "$ref": "UrlMap"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.urlMaps.validate":

type UrlMapsValidateCall struct {
	s                      *Service
	project                string
	urlMap                 string
	urlmapsvalidaterequest *UrlMapsValidateRequest
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
	context_               context.Context
}

// Validate: Run static validation for the UrlMap. In particular, the
// tests of the provided UrlMap will be run. Calling this method does
// NOT create the UrlMap.

// For details, see https://cloud.google.com/compute/docs/reference/latest/urlMaps/validate
func (r *UrlMapsService) Validate(project string, urlMap string, urlmapsvalidaterequest *UrlMapsValidateRequest) *UrlMapsValidateCall {
	return &UrlMapsValidateCall{
		s:                      r.s,
		project:                project,
		urlMap:                 urlMap,
		urlmapsvalidaterequest: urlmapsvalidaterequest,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "{project}/global/urlMaps/{urlMap}/validate",
		context_:               googleapi.NoContext,
	}
}
func (c *UrlMapsValidateCall) Context(ctx context.Context) *UrlMapsValidateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlMapsValidateCall) Fields(s ...googleapi.Field) *UrlMapsValidateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlMapsValidateCall) Do() (*UrlMapsValidateResponse, error) {
	var returnValue *UrlMapsValidateResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"urlMap":  c.urlMap,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.urlmapsvalidaterequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Run static validation for the UrlMap. In particular, the tests of the provided UrlMap will be run. Calling this method does NOT create the UrlMap.",
	//   "httpMethod": "POST",
	//   "id": "compute.urlMaps.validate",
	//   "parameterOrder": [
	//     "project",
	//     "urlMap"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "urlMap": {
	//       "description": "Name of the UrlMap resource to be validated as.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/urlMaps/{urlMap}/validate",
	//   "request": {
	//     "$ref": "UrlMapsValidateRequest"
	//   },
	//   "response": {
	//     "$ref": "UrlMapsValidateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.vpnTunnels.aggregatedList":

type VpnTunnelsAggregatedListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// AggregatedList: Retrieves the list of VPN tunnels grouped by scope.

func (r *VpnTunnelsService) AggregatedList(project string) *VpnTunnelsAggregatedListCall {
	return &VpnTunnelsAggregatedListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/aggregated/vpnTunnels",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *VpnTunnelsAggregatedListCall) Filter(filter string) *VpnTunnelsAggregatedListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *VpnTunnelsAggregatedListCall) MaxResults(maxResults int64) *VpnTunnelsAggregatedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *VpnTunnelsAggregatedListCall) PageToken(pageToken string) *VpnTunnelsAggregatedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *VpnTunnelsAggregatedListCall) Context(ctx context.Context) *VpnTunnelsAggregatedListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VpnTunnelsAggregatedListCall) Fields(s ...googleapi.Field) *VpnTunnelsAggregatedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VpnTunnelsAggregatedListCall) Do() (*VpnTunnelAggregatedList, error) {
	var returnValue *VpnTunnelAggregatedList
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
	//   "description": "Retrieves the list of VPN tunnels grouped by scope.",
	//   "httpMethod": "GET",
	//   "id": "compute.vpnTunnels.aggregatedList",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/aggregated/vpnTunnels",
	//   "response": {
	//     "$ref": "VpnTunnelAggregatedList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.vpnTunnels.delete":

type VpnTunnelsDeleteCall struct {
	s             *Service
	project       string
	region        string
	vpnTunnel     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified VpnTunnel resource.

func (r *VpnTunnelsService) Delete(project string, region string, vpnTunnel string) *VpnTunnelsDeleteCall {
	return &VpnTunnelsDeleteCall{
		s:             r.s,
		project:       project,
		region:        region,
		vpnTunnel:     vpnTunnel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/vpnTunnels/{vpnTunnel}",
		context_:      googleapi.NoContext,
	}
}
func (c *VpnTunnelsDeleteCall) Context(ctx context.Context) *VpnTunnelsDeleteCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VpnTunnelsDeleteCall) Fields(s ...googleapi.Field) *VpnTunnelsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VpnTunnelsDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"region":    c.region,
		"vpnTunnel": c.vpnTunnel,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified VpnTunnel resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.vpnTunnels.delete",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "vpnTunnel"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "vpnTunnel": {
	//       "description": "Name of the VpnTunnel resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/vpnTunnels/{vpnTunnel}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.vpnTunnels.get":

type VpnTunnelsGetCall struct {
	s             *Service
	project       string
	region        string
	vpnTunnel     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified VpnTunnel resource.

func (r *VpnTunnelsService) Get(project string, region string, vpnTunnel string) *VpnTunnelsGetCall {
	return &VpnTunnelsGetCall{
		s:             r.s,
		project:       project,
		region:        region,
		vpnTunnel:     vpnTunnel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/vpnTunnels/{vpnTunnel}",
		context_:      googleapi.NoContext,
	}
}
func (c *VpnTunnelsGetCall) Context(ctx context.Context) *VpnTunnelsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VpnTunnelsGetCall) Fields(s ...googleapi.Field) *VpnTunnelsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VpnTunnelsGetCall) Do() (*VpnTunnel, error) {
	var returnValue *VpnTunnel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"region":    c.region,
		"vpnTunnel": c.vpnTunnel,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified VpnTunnel resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.vpnTunnels.get",
	//   "parameterOrder": [
	//     "project",
	//     "region",
	//     "vpnTunnel"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "vpnTunnel": {
	//       "description": "Name of the VpnTunnel resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/vpnTunnels/{vpnTunnel}",
	//   "response": {
	//     "$ref": "VpnTunnel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.vpnTunnels.insert":

type VpnTunnelsInsertCall struct {
	s             *Service
	project       string
	region        string
	vpntunnel     *VpnTunnel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a VpnTunnel resource in the specified project and
// region using the data included in the request.

func (r *VpnTunnelsService) Insert(project string, region string, vpntunnel *VpnTunnel) *VpnTunnelsInsertCall {
	return &VpnTunnelsInsertCall{
		s:             r.s,
		project:       project,
		region:        region,
		vpntunnel:     vpntunnel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/vpnTunnels",
		context_:      googleapi.NoContext,
	}
}
func (c *VpnTunnelsInsertCall) Context(ctx context.Context) *VpnTunnelsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VpnTunnelsInsertCall) Fields(s ...googleapi.Field) *VpnTunnelsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VpnTunnelsInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.vpntunnel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a VpnTunnel resource in the specified project and region using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "compute.vpnTunnels.insert",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/vpnTunnels",
	//   "request": {
	//     "$ref": "VpnTunnel"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.vpnTunnels.list":

type VpnTunnelsListCall struct {
	s             *Service
	project       string
	region        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of VpnTunnel resources contained in the
// specified project and region.

func (r *VpnTunnelsService) List(project string, region string) *VpnTunnelsListCall {
	return &VpnTunnelsListCall{
		s:             r.s,
		project:       project,
		region:        region,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/regions/{region}/vpnTunnels",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *VpnTunnelsListCall) Filter(filter string) *VpnTunnelsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *VpnTunnelsListCall) MaxResults(maxResults int64) *VpnTunnelsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *VpnTunnelsListCall) PageToken(pageToken string) *VpnTunnelsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *VpnTunnelsListCall) Context(ctx context.Context) *VpnTunnelsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VpnTunnelsListCall) Fields(s ...googleapi.Field) *VpnTunnelsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VpnTunnelsListCall) Do() (*VpnTunnelList, error) {
	var returnValue *VpnTunnelList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"region":  c.region,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of VpnTunnel resources contained in the specified project and region.",
	//   "httpMethod": "GET",
	//   "id": "compute.vpnTunnels.list",
	//   "parameterOrder": [
	//     "project",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "The name of the region for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/regions/{region}/vpnTunnels",
	//   "response": {
	//     "$ref": "VpnTunnelList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.zoneOperations.delete":

type ZoneOperationsDeleteCall struct {
	s             *Service
	project       string
	zone          string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes the specified zone-specific operation resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/zoneOperations/delete
func (r *ZoneOperationsService) Delete(project string, zone string, operation string) *ZoneOperationsDeleteCall {
	return &ZoneOperationsDeleteCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *ZoneOperationsDeleteCall) Context(ctx context.Context) *ZoneOperationsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *ZoneOperationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
		"operation": c.operation,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes the specified zone-specific operation resource.",
	//   "httpMethod": "DELETE",
	//   "id": "compute.zoneOperations.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "compute.zoneOperations.get":

type ZoneOperationsGetCall struct {
	s             *Service
	project       string
	zone          string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Retrieves the specified zone-specific operation resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/zoneOperations/get
func (r *ZoneOperationsService) Get(project string, zone string, operation string) *ZoneOperationsGetCall {
	return &ZoneOperationsGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/operations/{operation}",
		context_:      googleapi.NoContext,
	}
}
func (c *ZoneOperationsGetCall) Context(ctx context.Context) *ZoneOperationsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsGetCall) Fields(s ...googleapi.Field) *ZoneOperationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneOperationsGetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
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
	//   "description": "Retrieves the specified zone-specific operation resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.zoneOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.zoneOperations.list":

type ZoneOperationsListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of operation resources contained within the
// specified zone.

// For details, see https://cloud.google.com/compute/docs/reference/latest/zoneOperations/list
func (r *ZoneOperationsService) List(project string, zone string) *ZoneOperationsListCall {
	return &ZoneOperationsListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/operations",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ZoneOperationsListCall) Filter(filter string) *ZoneOperationsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ZoneOperationsListCall) MaxResults(maxResults int64) *ZoneOperationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *ZoneOperationsListCall) PageToken(pageToken string) *ZoneOperationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ZoneOperationsListCall) Context(ctx context.Context) *ZoneOperationsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsListCall) Fields(s ...googleapi.Field) *ZoneOperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneOperationsListCall) Do() (*OperationList, error) {
	var returnValue *OperationList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Retrieves the list of operation resources contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "compute.zoneOperations.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.zones.get":

type ZonesGetCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the specified zone resource.

// For details, see https://cloud.google.com/compute/docs/reference/latest/zones/get
func (r *ZonesService) Get(project string, zone string) *ZonesGetCall {
	return &ZonesGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}",
		context_:      googleapi.NoContext,
	}
}
func (c *ZonesGetCall) Context(ctx context.Context) *ZonesGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZonesGetCall) Fields(s ...googleapi.Field) *ZonesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZonesGetCall) Do() (*Zone, error) {
	var returnValue *Zone
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the specified zone resource.",
	//   "httpMethod": "GET",
	//   "id": "compute.zones.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}",
	//   "response": {
	//     "$ref": "Zone"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "compute.zones.list":

type ZonesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Retrieves the list of zone resources available to the specified
// project.

// For details, see https://cloud.google.com/compute/docs/reference/latest/zones/list
func (r *ZonesService) List(project string) *ZonesListCall {
	return &ZonesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones",
		context_:      googleapi.NoContext,
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ZonesListCall) Filter(filter string) *ZonesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ZonesListCall) MaxResults(maxResults int64) *ZonesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *ZonesListCall) PageToken(pageToken string) *ZonesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}
func (c *ZonesListCall) Context(ctx context.Context) *ZonesListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZonesListCall) Fields(s ...googleapi.Field) *ZonesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZonesListCall) Do() (*ZoneList, error) {
	var returnValue *ZoneList
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
	//   "description": "Retrieves the list of zone resources available to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "compute.zones.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones",
	//   "response": {
	//     "$ref": "ZoneList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}
