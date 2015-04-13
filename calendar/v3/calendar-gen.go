// Package calendar provides access to the Calendar API.
//
// See https://developers.google.com/google-apps/calendar/firstapp
//
// Usage example:
//
//   import "github.com/jfcote87/api2/calendar/v3"
//   ...
//   calendarService, err := calendar.New(oauthHttpClient)
package calendar

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

const apiId = "calendar:v3"
const apiName = "calendar"
const apiVersion = "v3"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/calendar/v3/"}

// OAuth2 scopes used by this API.
const (
	// Manage your calendars
	CalendarScope = "https://www.googleapis.com/auth/calendar"

	// View your calendars
	CalendarReadonlyScope = "https://www.googleapis.com/auth/calendar.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Acl = NewAclService(s)
	s.CalendarList = NewCalendarListService(s)
	s.Calendars = NewCalendarsService(s)
	s.Channels = NewChannelsService(s)
	s.Colors = NewColorsService(s)
	s.Events = NewEventsService(s)
	s.Freebusy = NewFreebusyService(s)
	s.Settings = NewSettingsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Acl *AclService

	CalendarList *CalendarListService

	Calendars *CalendarsService

	Channels *ChannelsService

	Colors *ColorsService

	Events *EventsService

	Freebusy *FreebusyService

	Settings *SettingsService
}

func NewAclService(s *Service) *AclService {
	rs := &AclService{s: s}
	return rs
}

type AclService struct {
	s *Service
}

func NewCalendarListService(s *Service) *CalendarListService {
	rs := &CalendarListService{s: s}
	return rs
}

type CalendarListService struct {
	s *Service
}

func NewCalendarsService(s *Service) *CalendarsService {
	rs := &CalendarsService{s: s}
	return rs
}

type CalendarsService struct {
	s *Service
}

func NewChannelsService(s *Service) *ChannelsService {
	rs := &ChannelsService{s: s}
	return rs
}

type ChannelsService struct {
	s *Service
}

func NewColorsService(s *Service) *ColorsService {
	rs := &ColorsService{s: s}
	return rs
}

type ColorsService struct {
	s *Service
}

func NewEventsService(s *Service) *EventsService {
	rs := &EventsService{s: s}
	return rs
}

type EventsService struct {
	s *Service
}

func NewFreebusyService(s *Service) *FreebusyService {
	rs := &FreebusyService{s: s}
	return rs
}

type FreebusyService struct {
	s *Service
}

func NewSettingsService(s *Service) *SettingsService {
	rs := &SettingsService{s: s}
	return rs
}

type SettingsService struct {
	s *Service
}

type Acl struct {
	// Etag: ETag of the collection.
	Etag string `json:"etag,omitempty"`

	// Items: List of rules on the access control list.
	Items []*AclRule `json:"items,omitempty"`

	// Kind: Type of the collection ("calendar#acl").
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result.
	// Omitted if no further results are available, in which case
	// nextSyncToken is provided.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NextSyncToken: Token used at a later point in time to retrieve only
	// the entries that have changed since this result was returned. Omitted
	// if further results are available, in which case nextPageToken is
	// provided.
	NextSyncToken string `json:"nextSyncToken,omitempty"`
}

type AclRule struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: Identifier of the ACL rule.
	Id string `json:"id,omitempty"`

	// Kind: Type of the resource ("calendar#aclRule").
	Kind string `json:"kind,omitempty"`

	// Role: The role assigned to the scope. Possible values are:
	// - "none"
	// - Provides no access.
	// - "freeBusyReader" - Provides read access to
	// free/busy information.
	// - "reader" - Provides read access to the
	// calendar. Private events will appear to users with reader access, but
	// event details will be hidden.
	// - "writer" - Provides read and write
	// access to the calendar. Private events will appear to users with
	// writer access, and event details will be visible.
	// - "owner" -
	// Provides ownership of the calendar. This role has all of the
	// permissions of the writer role with the additional ability to see and
	// manipulate ACLs.
	Role string `json:"role,omitempty"`

	// Scope: The scope of the rule.
	Scope *AclRuleScope `json:"scope,omitempty"`
}

type AclRuleScope struct {
	// Type: The type of the scope. Possible values are:
	// - "default" - The
	// public scope. This is the default value.
	// - "user" - Limits the scope
	// to a single user.
	// - "group" - Limits the scope to a group.
	// -
	// "domain" - Limits the scope to a domain.  Note: The permissions
	// granted to the "default", or public, scope apply to any user,
	// authenticated or not.
	Type string `json:"type,omitempty"`

	// Value: The email address of a user or group, or the name of a domain,
	// depending on the scope type. Omitted for type "default".
	Value string `json:"value,omitempty"`
}

type Calendar struct {
	// Description: Description of the calendar. Optional.
	Description string `json:"description,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: Identifier of the calendar.
	Id string `json:"id,omitempty"`

	// Kind: Type of the resource ("calendar#calendar").
	Kind string `json:"kind,omitempty"`

	// Location: Geographic location of the calendar as free-form text.
	// Optional.
	Location string `json:"location,omitempty"`

	// Summary: Title of the calendar.
	Summary string `json:"summary,omitempty"`

	// TimeZone: The time zone of the calendar. (Formatted as an IANA Time
	// Zone Database name, e.g. "Europe/Zurich".) Optional.
	TimeZone string `json:"timeZone,omitempty"`
}

type CalendarList struct {
	// Etag: ETag of the collection.
	Etag string `json:"etag,omitempty"`

	// Items: Calendars that are present on the user's calendar list.
	Items []*CalendarListEntry `json:"items,omitempty"`

	// Kind: Type of the collection ("calendar#calendarList").
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result.
	// Omitted if no further results are available, in which case
	// nextSyncToken is provided.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NextSyncToken: Token used at a later point in time to retrieve only
	// the entries that have changed since this result was returned. Omitted
	// if further results are available, in which case nextPageToken is
	// provided.
	NextSyncToken string `json:"nextSyncToken,omitempty"`
}

type CalendarListEntry struct {
	// AccessRole: The effective access role that the authenticated user has
	// on the calendar. Read-only. Possible values are:
	// - "freeBusyReader"
	// - Provides read access to free/busy information.
	// - "reader" -
	// Provides read access to the calendar. Private events will appear to
	// users with reader access, but event details will be hidden.
	// -
	// "writer" - Provides read and write access to the calendar. Private
	// events will appear to users with writer access, and event details
	// will be visible.
	// - "owner" - Provides ownership of the calendar.
	// This role has all of the permissions of the writer role with the
	// additional ability to see and manipulate ACLs.
	AccessRole string `json:"accessRole,omitempty"`

	// BackgroundColor: The main color of the calendar in the hexadecimal
	// format "#0088aa". This property supersedes the index-based colorId
	// property. Optional.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// ColorId: The color of the calendar. This is an ID referring to an
	// entry in the calendar section of the colors definition (see the
	// colors endpoint). Optional.
	ColorId string `json:"colorId,omitempty"`

	// DefaultReminders: The default reminders that the authenticated user
	// has for this calendar.
	DefaultReminders []*EventReminder `json:"defaultReminders,omitempty"`

	// Deleted: Whether this calendar list entry has been deleted from the
	// calendar list. Read-only. Optional. The default is False.
	Deleted bool `json:"deleted,omitempty"`

	// Description: Description of the calendar. Optional. Read-only.
	Description string `json:"description,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// ForegroundColor: The foreground color of the calendar in the
	// hexadecimal format "#ffffff". This property supersedes the
	// index-based colorId property. Optional.
	ForegroundColor string `json:"foregroundColor,omitempty"`

	// Hidden: Whether the calendar has been hidden from the list. Optional.
	// The default is False.
	Hidden bool `json:"hidden,omitempty"`

	// Id: Identifier of the calendar.
	Id string `json:"id,omitempty"`

	// Kind: Type of the resource ("calendar#calendarListEntry").
	Kind string `json:"kind,omitempty"`

	// Location: Geographic location of the calendar as free-form text.
	// Optional. Read-only.
	Location string `json:"location,omitempty"`

	// NotificationSettings: The notifications that the authenticated user
	// is receiving for this calendar.
	NotificationSettings *CalendarListEntryNotificationSettings `json:"notificationSettings,omitempty"`

	// Primary: Whether the calendar is the primary calendar of the
	// authenticated user. Read-only. Optional. The default is False.
	Primary bool `json:"primary,omitempty"`

	// Selected: Whether the calendar content shows up in the calendar UI.
	// Optional. The default is False.
	Selected bool `json:"selected,omitempty"`

	// Summary: Title of the calendar. Read-only.
	Summary string `json:"summary,omitempty"`

	// SummaryOverride: The summary that the authenticated user has set for
	// this calendar. Optional.
	SummaryOverride string `json:"summaryOverride,omitempty"`

	// TimeZone: The time zone of the calendar. Optional. Read-only.
	TimeZone string `json:"timeZone,omitempty"`
}

type CalendarListEntryNotificationSettings struct {
	// Notifications: The list of notifications set for this calendar.
	Notifications []*CalendarNotification `json:"notifications,omitempty"`
}

type CalendarNotification struct {
	// Method: The method used to deliver the notification. Possible values
	// are:
	// - "email" - Reminders are sent via email.
	// - "sms" - Reminders
	// are sent via SMS. This value is read-only and is ignored on inserts
	// and updates.
	Method string `json:"method,omitempty"`

	// Type: The type of notification. Possible values are:
	// -
	// "eventCreation" - Notification sent when a new event is put on the
	// calendar.
	// - "eventChange" - Notification sent when an event is
	// changed.
	// - "eventCancellation" - Notification sent when an event is
	// cancelled.
	// - "eventResponse" - Notification sent when an event is
	// changed.
	// - "agenda" - An agenda with the events of the day (sent out
	// in the morning).
	Type string `json:"type,omitempty"`
}

type Channel struct {
	// Address: The address where notifications are delivered for this
	// channel.
	Address string `json:"address,omitempty"`

	// Expiration: Date and time of notification channel expiration,
	// expressed as a Unix timestamp, in milliseconds. Optional.
	Expiration int64 `json:"expiration,omitempty,string"`

	// Id: A UUID or similar unique string that identifies this channel.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this as a notification channel used to watch for
	// changes to a resource. Value: the fixed string "api#channel".
	Kind string `json:"kind,omitempty"`

	// Params: Additional parameters controlling delivery channel behavior.
	// Optional.
	Params map[string]string `json:"params,omitempty"`

	// Payload: A Boolean value to indicate whether payload is wanted.
	// Optional.
	Payload bool `json:"payload,omitempty"`

	// ResourceId: An opaque ID that identifies the resource being watched
	// on this channel. Stable across different API versions.
	ResourceId string `json:"resourceId,omitempty"`

	// ResourceUri: A version-specific identifier for the watched resource.
	ResourceUri string `json:"resourceUri,omitempty"`

	// Token: An arbitrary string delivered to the target address with each
	// notification delivered over this channel. Optional.
	Token string `json:"token,omitempty"`

	// Type: The type of delivery mechanism used for this channel.
	Type string `json:"type,omitempty"`
}

type ColorDefinition struct {
	// Background: The background color associated with this color
	// definition.
	Background string `json:"background,omitempty"`

	// Foreground: The foreground color that can be used to write on top of
	// a background with 'background' color.
	Foreground string `json:"foreground,omitempty"`
}

type Colors struct {
	// Calendar: Palette of calendar colors, mapping from the color ID to
	// its definition. A calendarListEntry resource refers to one of these
	// color IDs in its color field. Read-only.
	Calendar map[string]ColorDefinition `json:"calendar,omitempty"`

	// Event: Palette of event colors, mapping from the color ID to its
	// definition. An event resource may refer to one of these color IDs in
	// its color field. Read-only.
	Event map[string]ColorDefinition `json:"event,omitempty"`

	// Kind: Type of the resource ("calendar#colors").
	Kind string `json:"kind,omitempty"`

	// Updated: Last modification time of the color palette (as a RFC 3339
	// timestamp). Read-only.
	Updated string `json:"updated,omitempty"`
}

type Error struct {
	// Domain: Domain, or broad category, of the error.
	Domain string `json:"domain,omitempty"`

	// Reason: Specific reason for the error. Some of the possible values
	// are:
	// - "groupTooBig" - The group of users requested is too large
	// for a single query.
	// - "tooManyCalendarsRequested" - The number of
	// calendars requested is too large for a single query.
	// - "notFound" -
	// The requested resource was not found.
	// - "internalError" - The API
	// service has encountered an internal error.  Additional error types
	// may be added in the future, so clients should gracefully handle
	// additional error statuses not included in this list.
	Reason string `json:"reason,omitempty"`
}

type Event struct {
	// AnyoneCanAddSelf: Whether anyone can invite themselves to the event.
	// Optional. The default is False.
	AnyoneCanAddSelf bool `json:"anyoneCanAddSelf,omitempty"`

	// Attendees: The attendees of the event.
	Attendees []*EventAttendee `json:"attendees,omitempty"`

	// AttendeesOmitted: Whether attendees may have been omitted from the
	// event's representation. When retrieving an event, this may be due to
	// a restriction specified by the maxAttendee query parameter. When
	// updating an event, this can be used to only update the participant's
	// response. Optional. The default is False.
	AttendeesOmitted bool `json:"attendeesOmitted,omitempty"`

	// ColorId: The color of the event. This is an ID referring to an entry
	// in the event section of the colors definition (see the  colors
	// endpoint). Optional.
	ColorId string `json:"colorId,omitempty"`

	// Created: Creation time of the event (as a RFC 3339 timestamp).
	// Read-only.
	Created string `json:"created,omitempty"`

	// Creator: The creator of the event. Read-only.
	Creator *EventCreator `json:"creator,omitempty"`

	// Description: Description of the event. Optional.
	Description string `json:"description,omitempty"`

	// End: The (exclusive) end time of the event. For a recurring event,
	// this is the end time of the first instance.
	End *EventDateTime `json:"end,omitempty"`

	// EndTimeUnspecified: Whether the end time is actually unspecified. An
	// end time is still provided for compatibility reasons, even if this
	// attribute is set to True. The default is False.
	EndTimeUnspecified bool `json:"endTimeUnspecified,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// ExtendedProperties: Extended properties of the event.
	ExtendedProperties *EventExtendedProperties `json:"extendedProperties,omitempty"`

	// Gadget: A gadget that extends this event.
	Gadget *EventGadget `json:"gadget,omitempty"`

	// GuestsCanInviteOthers: Whether attendees other than the organizer can
	// invite others to the event. Optional. The default is True.
	GuestsCanInviteOthers bool `json:"guestsCanInviteOthers,omitempty"`

	// GuestsCanModify: Whether attendees other than the organizer can
	// modify the event. Optional. The default is False.
	GuestsCanModify bool `json:"guestsCanModify,omitempty"`

	// GuestsCanSeeOtherGuests: Whether attendees other than the organizer
	// can see who the event's attendees are. Optional. The default is True.
	GuestsCanSeeOtherGuests bool `json:"guestsCanSeeOtherGuests,omitempty"`

	// HangoutLink: An absolute link to the Google+ hangout associated with
	// this event. Read-only.
	HangoutLink string `json:"hangoutLink,omitempty"`

	// HtmlLink: An absolute link to this event in the Google Calendar Web
	// UI. Read-only.
	HtmlLink string `json:"htmlLink,omitempty"`

	// ICalUID: Event ID in the iCalendar format.
	ICalUID string `json:"iCalUID,omitempty"`

	// Id: Identifier of the event. When creating new single or recurring
	// events, you can specify their IDs. Provided IDs must follow these
	// rules:
	// - characters allowed in the ID are those used in base32hex
	// encoding, i.e. lowercase letters a-v and digits 0-9, see section
	// 3.1.2 in RFC2938
	// - the length of the ID must be between 5 and 1024
	// characters
	// - the ID must be unique per calendar  Due to the globally
	// distributed nature of the system, we cannot guarantee that ID
	// collisions will be detected at event creation time. To minimize the
	// risk of collisions we recommend using an established UUID algorithm
	// such as one described in RFC4122.
	Id string `json:"id,omitempty"`

	// Kind: Type of the resource ("calendar#event").
	Kind string `json:"kind,omitempty"`

	// Location: Geographic location of the event as free-form text.
	// Optional.
	Location string `json:"location,omitempty"`

	// Locked: Whether this is a locked event copy where no changes can be
	// made to the main event fields "summary", "description", "location",
	// "start", "end" or "recurrence". The default is False. Read-Only.
	Locked bool `json:"locked,omitempty"`

	// Organizer: The organizer of the event. If the organizer is also an
	// attendee, this is indicated with a separate entry in attendees with
	// the organizer field set to True. To change the organizer, use the
	// move operation. Read-only, except when importing an event.
	Organizer *EventOrganizer `json:"organizer,omitempty"`

	// OriginalStartTime: For an instance of a recurring event, this is the
	// time at which this event would start according to the recurrence data
	// in the recurring event identified by recurringEventId. Immutable.
	OriginalStartTime *EventDateTime `json:"originalStartTime,omitempty"`

	// PrivateCopy: Whether this is a private event copy where changes are
	// not shared with other copies on other calendars. Optional. Immutable.
	// The default is False.
	PrivateCopy bool `json:"privateCopy,omitempty"`

	// Recurrence: List of RRULE, EXRULE, RDATE and EXDATE lines for a
	// recurring event. This field is omitted for single events or instances
	// of recurring events.
	Recurrence []string `json:"recurrence,omitempty"`

	// RecurringEventId: For an instance of a recurring event, this is the
	// event ID of the recurring event itself. Immutable.
	RecurringEventId string `json:"recurringEventId,omitempty"`

	// Reminders: Information about the event's reminders for the
	// authenticated user.
	Reminders *EventReminders `json:"reminders,omitempty"`

	// Sequence: Sequence number as per iCalendar.
	Sequence int64 `json:"sequence,omitempty"`

	// Source: Source of an event from which it was created; for example a
	// web page, an email message or any document identifiable by an URL
	// using HTTP/HTTPS protocol. Accessible only by the creator of the
	// event.
	Source *EventSource `json:"source,omitempty"`

	// Start: The (inclusive) start time of the event. For a recurring
	// event, this is the start time of the first instance.
	Start *EventDateTime `json:"start,omitempty"`

	// Status: Status of the event. Optional. Possible values are:
	// -
	// "confirmed" - The event is confirmed. This is the default status.
	// -
	// "tentative" - The event is tentatively confirmed.
	// - "cancelled" -
	// The event is cancelled.
	Status string `json:"status,omitempty"`

	// Summary: Title of the event.
	Summary string `json:"summary,omitempty"`

	// Transparency: Whether the event blocks time on the calendar.
	// Optional. Possible values are:
	// - "opaque" - The event blocks time
	// on the calendar. This is the default value.
	// - "transparent" - The
	// event does not block time on the calendar.
	Transparency string `json:"transparency,omitempty"`

	// Updated: Last modification time of the event (as a RFC 3339
	// timestamp). Read-only.
	Updated string `json:"updated,omitempty"`

	// Visibility: Visibility of the event. Optional. Possible values are:
	//
	// - "default" - Uses the default visibility for events on the
	// calendar. This is the default value.
	// - "public" - The event is
	// public and event details are visible to all readers of the calendar.
	//
	// - "private" - The event is private and only event attendees may view
	// event details.
	// - "confidential" - The event is private. This value
	// is provided for compatibility reasons.
	Visibility string `json:"visibility,omitempty"`
}

type EventCreator struct {
	// DisplayName: The creator's name, if available.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The creator's email address, if available.
	Email string `json:"email,omitempty"`

	// Id: The creator's Profile ID, if available.
	Id string `json:"id,omitempty"`

	// Self: Whether the creator corresponds to the calendar on which this
	// copy of the event appears. Read-only. The default is False.
	Self bool `json:"self,omitempty"`
}

type EventExtendedProperties struct {
	// Private: Properties that are private to the copy of the event that
	// appears on this calendar.
	Private map[string]string `json:"private,omitempty"`

	// Shared: Properties that are shared between copies of the event on
	// other attendees' calendars.
	Shared map[string]string `json:"shared,omitempty"`
}

type EventGadget struct {
	// Display: The gadget's display mode. Optional. Possible values are:
	//
	// - "icon" - The gadget displays next to the event's title in the
	// calendar view.
	// - "chip" - The gadget displays when the event is
	// clicked.
	Display string `json:"display,omitempty"`

	// Height: The gadget's height in pixels. Optional.
	Height int64 `json:"height,omitempty"`

	// IconLink: The gadget's icon URL.
	IconLink string `json:"iconLink,omitempty"`

	// Link: The gadget's URL.
	Link string `json:"link,omitempty"`

	// Preferences: Preferences.
	Preferences map[string]string `json:"preferences,omitempty"`

	// Title: The gadget's title.
	Title string `json:"title,omitempty"`

	// Type: The gadget's type.
	Type string `json:"type,omitempty"`

	// Width: The gadget's width in pixels. Optional.
	Width int64 `json:"width,omitempty"`
}

type EventOrganizer struct {
	// DisplayName: The organizer's name, if available.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The organizer's email address, if available.
	Email string `json:"email,omitempty"`

	// Id: The organizer's Profile ID, if available.
	Id string `json:"id,omitempty"`

	// Self: Whether the organizer corresponds to the calendar on which this
	// copy of the event appears. Read-only. The default is False.
	Self bool `json:"self,omitempty"`
}

type EventReminders struct {
	// Overrides: If the event doesn't use the default reminders, this lists
	// the reminders specific to the event, or, if not set, indicates that
	// no reminders are set for this event.
	Overrides []*EventReminder `json:"overrides,omitempty"`

	// UseDefault: Whether the default reminders of the calendar apply to
	// the event.
	UseDefault bool `json:"useDefault,omitempty"`
}

type EventSource struct {
	// Title: Title of the source; for example a title of a web page or an
	// email subject.
	Title string `json:"title,omitempty"`

	// Url: URL of the source pointing to a resource. URL's protocol must be
	// HTTP or HTTPS.
	Url string `json:"url,omitempty"`
}

type EventAttachment struct {
}

type EventAttendee struct {
	// AdditionalGuests: Number of additional guests. Optional. The default
	// is 0.
	AdditionalGuests int64 `json:"additionalGuests,omitempty"`

	// Comment: The attendee's response comment. Optional.
	Comment string `json:"comment,omitempty"`

	// DisplayName: The attendee's name, if available. Optional.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The attendee's email address, if available. This field must be
	// present when adding an attendee.
	Email string `json:"email,omitempty"`

	// Id: The attendee's Profile ID, if available.
	Id string `json:"id,omitempty"`

	// Optional: Whether this is an optional attendee. Optional. The default
	// is False.
	Optional bool `json:"optional,omitempty"`

	// Organizer: Whether the attendee is the organizer of the event.
	// Read-only. The default is False.
	Organizer bool `json:"organizer,omitempty"`

	// Resource: Whether the attendee is a resource. Read-only. The default
	// is False.
	Resource bool `json:"resource,omitempty"`

	// ResponseStatus: The attendee's response status. Possible values are:
	//
	// - "needsAction" - The attendee has not responded to the invitation.
	//
	// - "declined" - The attendee has declined the invitation.
	// -
	// "tentative" - The attendee has tentatively accepted the invitation.
	//
	// - "accepted" - The attendee has accepted the invitation.
	ResponseStatus string `json:"responseStatus,omitempty"`

	// Self: Whether this entry represents the calendar on which this copy
	// of the event appears. Read-only. The default is False.
	Self bool `json:"self,omitempty"`
}

type EventDateTime struct {
	// Date: The date, in the format "yyyy-mm-dd", if this is an all-day
	// event.
	Date string `json:"date,omitempty"`

	// DateTime: The time, as a combined date-time value (formatted
	// according to RFC 3339). A time zone offset is required unless a time
	// zone is explicitly specified in timeZone.
	DateTime string `json:"dateTime,omitempty"`

	// TimeZone: The time zone in which the time is specified. (Formatted as
	// an IANA Time Zone Database name, e.g. "Europe/Zurich".) For recurring
	// events this field is required and specifies the time zone in which
	// the recurrence is expanded. For single events this field is optional
	// and indicates a custom time zone for the event start/end.
	TimeZone string `json:"timeZone,omitempty"`
}

type EventReminder struct {
	// Method: The method used by this reminder. Possible values are:
	// -
	// "email" - Reminders are sent via email.
	// - "sms" - Reminders are sent
	// via SMS.
	// - "popup" - Reminders are sent via a UI popup.
	Method string `json:"method,omitempty"`

	// Minutes: Number of minutes before the start of the event when the
	// reminder should trigger.
	Minutes int64 `json:"minutes,omitempty"`
}

type Events struct {
	// AccessRole: The user's access role for this calendar. Read-only.
	// Possible values are:
	// - "none" - The user has no access.
	// -
	// "freeBusyReader" - The user has read access to free/busy information.
	//
	// - "reader" - The user has read access to the calendar. Private
	// events will appear to users with reader access, but event details
	// will be hidden.
	// - "writer" - The user has read and write access to
	// the calendar. Private events will appear to users with writer access,
	// and event details will be visible.
	// - "owner" - The user has
	// ownership of the calendar. This role has all of the permissions of
	// the writer role with the additional ability to see and manipulate
	// ACLs.
	AccessRole string `json:"accessRole,omitempty"`

	// DefaultReminders: The default reminders on the calendar for the
	// authenticated user. These reminders apply to all events on this
	// calendar that do not explicitly override them (i.e. do not have
	// reminders.useDefault set to True).
	DefaultReminders []*EventReminder `json:"defaultReminders,omitempty"`

	// Description: Description of the calendar. Read-only.
	Description string `json:"description,omitempty"`

	// Etag: ETag of the collection.
	Etag string `json:"etag,omitempty"`

	// Items: List of events on the calendar.
	Items []*Event `json:"items,omitempty"`

	// Kind: Type of the collection ("calendar#events").
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result.
	// Omitted if no further results are available, in which case
	// nextSyncToken is provided.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NextSyncToken: Token used at a later point in time to retrieve only
	// the entries that have changed since this result was returned. Omitted
	// if further results are available, in which case nextPageToken is
	// provided.
	NextSyncToken string `json:"nextSyncToken,omitempty"`

	// Summary: Title of the calendar. Read-only.
	Summary string `json:"summary,omitempty"`

	// TimeZone: The time zone of the calendar. Read-only.
	TimeZone string `json:"timeZone,omitempty"`

	// Updated: Last modification time of the calendar (as a RFC 3339
	// timestamp). Read-only.
	Updated string `json:"updated,omitempty"`
}

type FreeBusyCalendar struct {
	// Busy: List of time ranges during which this calendar should be
	// regarded as busy.
	Busy []*TimePeriod `json:"busy,omitempty"`

	// Errors: Optional error(s) (if computation for the calendar failed).
	Errors []*Error `json:"errors,omitempty"`
}

type FreeBusyGroup struct {
	// Calendars: List of calendars' identifiers within a group.
	Calendars []string `json:"calendars,omitempty"`

	// Errors: Optional error(s) (if computation for the group failed).
	Errors []*Error `json:"errors,omitempty"`
}

type FreeBusyRequest struct {
	// CalendarExpansionMax: Maximal number of calendars for which FreeBusy
	// information is to be provided. Optional.
	CalendarExpansionMax int64 `json:"calendarExpansionMax,omitempty"`

	// GroupExpansionMax: Maximal number of calendar identifiers to be
	// provided for a single group. Optional. An error will be returned for
	// a group with more members than this value.
	GroupExpansionMax int64 `json:"groupExpansionMax,omitempty"`

	// Items: List of calendars and/or groups to query.
	Items []*FreeBusyRequestItem `json:"items,omitempty"`

	// TimeMax: The end of the interval for the query.
	TimeMax string `json:"timeMax,omitempty"`

	// TimeMin: The start of the interval for the query.
	TimeMin string `json:"timeMin,omitempty"`

	// TimeZone: Time zone used in the response. Optional. The default is
	// UTC.
	TimeZone string `json:"timeZone,omitempty"`
}

type FreeBusyRequestItem struct {
	// Id: The identifier of a calendar or a group.
	Id string `json:"id,omitempty"`
}

type FreeBusyResponse struct {
	// Calendars: List of free/busy information for calendars.
	Calendars map[string]FreeBusyCalendar `json:"calendars,omitempty"`

	// Groups: Expansion of groups.
	Groups map[string]FreeBusyGroup `json:"groups,omitempty"`

	// Kind: Type of the resource ("calendar#freeBusy").
	Kind string `json:"kind,omitempty"`

	// TimeMax: The end of the interval.
	TimeMax string `json:"timeMax,omitempty"`

	// TimeMin: The start of the interval.
	TimeMin string `json:"timeMin,omitempty"`
}

type Setting struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The id of the user setting.
	Id string `json:"id,omitempty"`

	// Kind: Type of the resource ("calendar#setting").
	Kind string `json:"kind,omitempty"`

	// Value: Value of the user setting. The format of the value depends on
	// the ID of the setting. It must always be a UTF-8 string of length up
	// to 1024 characters.
	Value string `json:"value,omitempty"`
}

type Settings struct {
	// Etag: Etag of the collection.
	Etag string `json:"etag,omitempty"`

	// Items: List of user settings.
	Items []*Setting `json:"items,omitempty"`

	// Kind: Type of the collection ("calendar#settings").
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result.
	// Omitted if no further results are available, in which case
	// nextSyncToken is provided.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NextSyncToken: Token used at a later point in time to retrieve only
	// the entries that have changed since this result was returned. Omitted
	// if further results are available, in which case nextPageToken is
	// provided.
	NextSyncToken string `json:"nextSyncToken,omitempty"`
}

type TimePeriod struct {
	// End: The (exclusive) end of the time period.
	End string `json:"end,omitempty"`

	// Start: The (inclusive) start of the time period.
	Start string `json:"start,omitempty"`
}

// method id "calendar.acl.delete":

type AclDeleteCall struct {
	s             *Service
	calendarId    string
	ruleId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes an access control rule.

func (r *AclService) Delete(calendarId string, ruleId string) *AclDeleteCall {
	return &AclDeleteCall{
		s:             r.s,
		calendarId:    calendarId,
		ruleId:        ruleId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl/{ruleId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AclDeleteCall) Context(ctx context.Context) *AclDeleteCall {
	c.context_ = ctx
	return c
}

func (c *AclDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"ruleId":     c.ruleId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes an access control rule.",
	//   "httpMethod": "DELETE",
	//   "id": "calendar.acl.delete",
	//   "parameterOrder": [
	//     "calendarId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "ACL rule identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl/{ruleId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.acl.get":

type AclGetCall struct {
	s             *Service
	calendarId    string
	ruleId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns an access control rule.

func (r *AclService) Get(calendarId string, ruleId string) *AclGetCall {
	return &AclGetCall{
		s:             r.s,
		calendarId:    calendarId,
		ruleId:        ruleId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl/{ruleId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AclGetCall) Context(ctx context.Context) *AclGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AclGetCall) Fields(s ...googleapi.Field) *AclGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AclGetCall) Do() (*AclRule, error) {
	var returnValue *AclRule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"ruleId":     c.ruleId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns an access control rule.",
	//   "httpMethod": "GET",
	//   "id": "calendar.acl.get",
	//   "parameterOrder": [
	//     "calendarId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "ACL rule identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl/{ruleId}",
	//   "response": {
	//     "$ref": "AclRule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.acl.insert":

type AclInsertCall struct {
	s             *Service
	calendarId    string
	aclrule       *AclRule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates an access control rule.

func (r *AclService) Insert(calendarId string, aclrule *AclRule) *AclInsertCall {
	return &AclInsertCall{
		s:             r.s,
		calendarId:    calendarId,
		aclrule:       aclrule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl",
		context_:      googleapi.NoContext,
	}
}
func (c *AclInsertCall) Context(ctx context.Context) *AclInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AclInsertCall) Fields(s ...googleapi.Field) *AclInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AclInsertCall) Do() (*AclRule, error) {
	var returnValue *AclRule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.aclrule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an access control rule.",
	//   "httpMethod": "POST",
	//   "id": "calendar.acl.insert",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl",
	//   "request": {
	//     "$ref": "AclRule"
	//   },
	//   "response": {
	//     "$ref": "AclRule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.acl.list":

type AclListCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Returns the rules in the access control list for the calendar.

func (r *AclService) List(calendarId string) *AclListCall {
	return &AclListCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. By default the value is 100
// entries. The page size can never be larger than 250 entries.
func (c *AclListCall) MaxResults(maxResults int64) *AclListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *AclListCall) PageToken(pageToken string) *AclListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted ACLs in the result. Deleted ACLs are represented by
// role equal to "none". Deleted ACLs will always be included if
// syncToken is provided.  The default is False.
func (c *AclListCall) ShowDeleted(showDeleted bool) *AclListCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then. All
// entries deleted since the previous list request will always be in the
// result set and it is not allowed to set showDeleted to False.
// If the
// syncToken expires, the server will respond with a 410 GONE response
// code and the client should clear its storage and perform a full
// synchronization without any syncToken.
// Learn more about incremental
// synchronization.
//  The default is to return all entries.
func (c *AclListCall) SyncToken(syncToken string) *AclListCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}
func (c *AclListCall) Context(ctx context.Context) *AclListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AclListCall) Fields(s ...googleapi.Field) *AclListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AclListCall) Do() (*Acl, error) {
	var returnValue *Acl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the rules in the access control list for the calendar.",
	//   "httpMethod": "GET",
	//   "id": "calendar.acl.list",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted ACLs in the result. Deleted ACLs are represented by role equal to \"none\". Deleted ACLs will always be included if syncToken is provided. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. All entries deleted since the previous list request will always be in the result set and it is not allowed to set showDeleted to False.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl",
	//   "response": {
	//     "$ref": "Acl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.acl.patch":

type AclPatchCall struct {
	s             *Service
	calendarId    string
	ruleId        string
	aclrule       *AclRule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates an access control rule. This method supports patch
// semantics.

func (r *AclService) Patch(calendarId string, ruleId string, aclrule *AclRule) *AclPatchCall {
	return &AclPatchCall{
		s:             r.s,
		calendarId:    calendarId,
		ruleId:        ruleId,
		aclrule:       aclrule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl/{ruleId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AclPatchCall) Context(ctx context.Context) *AclPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AclPatchCall) Fields(s ...googleapi.Field) *AclPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AclPatchCall) Do() (*AclRule, error) {
	var returnValue *AclRule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"ruleId":     c.ruleId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.aclrule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an access control rule. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "calendar.acl.patch",
	//   "parameterOrder": [
	//     "calendarId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "ACL rule identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl/{ruleId}",
	//   "request": {
	//     "$ref": "AclRule"
	//   },
	//   "response": {
	//     "$ref": "AclRule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.acl.update":

type AclUpdateCall struct {
	s             *Service
	calendarId    string
	ruleId        string
	aclrule       *AclRule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates an access control rule.

func (r *AclService) Update(calendarId string, ruleId string, aclrule *AclRule) *AclUpdateCall {
	return &AclUpdateCall{
		s:             r.s,
		calendarId:    calendarId,
		ruleId:        ruleId,
		aclrule:       aclrule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl/{ruleId}",
		context_:      googleapi.NoContext,
	}
}
func (c *AclUpdateCall) Context(ctx context.Context) *AclUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AclUpdateCall) Fields(s ...googleapi.Field) *AclUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AclUpdateCall) Do() (*AclRule, error) {
	var returnValue *AclRule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"ruleId":     c.ruleId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.aclrule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an access control rule.",
	//   "httpMethod": "PUT",
	//   "id": "calendar.acl.update",
	//   "parameterOrder": [
	//     "calendarId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "ACL rule identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl/{ruleId}",
	//   "request": {
	//     "$ref": "AclRule"
	//   },
	//   "response": {
	//     "$ref": "AclRule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.acl.watch":

type AclWatchCall struct {
	s             *Service
	calendarId    string
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Watch: Watch for changes to ACL resources.

func (r *AclService) Watch(calendarId string, channel *Channel) *AclWatchCall {
	return &AclWatchCall{
		s:             r.s,
		calendarId:    calendarId,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/acl/watch",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. By default the value is 100
// entries. The page size can never be larger than 250 entries.
func (c *AclWatchCall) MaxResults(maxResults int64) *AclWatchCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *AclWatchCall) PageToken(pageToken string) *AclWatchCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted ACLs in the result. Deleted ACLs are represented by
// role equal to "none". Deleted ACLs will always be included if
// syncToken is provided.  The default is False.
func (c *AclWatchCall) ShowDeleted(showDeleted bool) *AclWatchCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then. All
// entries deleted since the previous list request will always be in the
// result set and it is not allowed to set showDeleted to False.
// If the
// syncToken expires, the server will respond with a 410 GONE response
// code and the client should clear its storage and perform a full
// synchronization without any syncToken.
// Learn more about incremental
// synchronization.
//  The default is to return all entries.
func (c *AclWatchCall) SyncToken(syncToken string) *AclWatchCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}
func (c *AclWatchCall) Context(ctx context.Context) *AclWatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AclWatchCall) Fields(s ...googleapi.Field) *AclWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AclWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Watch for changes to ACL resources.",
	//   "httpMethod": "POST",
	//   "id": "calendar.acl.watch",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted ACLs in the result. Deleted ACLs are represented by role equal to \"none\". Deleted ACLs will always be included if syncToken is provided. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. All entries deleted since the previous list request will always be in the result set and it is not allowed to set showDeleted to False.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/acl/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.calendarList.delete":

type CalendarListDeleteCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes an entry on the user's calendar list.

func (r *CalendarListService) Delete(calendarId string) *CalendarListDeleteCall {
	return &CalendarListDeleteCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/calendarList/{calendarId}",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarListDeleteCall) Context(ctx context.Context) *CalendarListDeleteCall {
	c.context_ = ctx
	return c
}

func (c *CalendarListDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes an entry on the user's calendar list.",
	//   "httpMethod": "DELETE",
	//   "id": "calendar.calendarList.delete",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/calendarList/{calendarId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendarList.get":

type CalendarListGetCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns an entry on the user's calendar list.

func (r *CalendarListService) Get(calendarId string) *CalendarListGetCall {
	return &CalendarListGetCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/calendarList/{calendarId}",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarListGetCall) Context(ctx context.Context) *CalendarListGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarListGetCall) Fields(s ...googleapi.Field) *CalendarListGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarListGetCall) Do() (*CalendarListEntry, error) {
	var returnValue *CalendarListEntry
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns an entry on the user's calendar list.",
	//   "httpMethod": "GET",
	//   "id": "calendar.calendarList.get",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/calendarList/{calendarId}",
	//   "response": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.calendarList.insert":

type CalendarListInsertCall struct {
	s                 *Service
	calendarlistentry *CalendarListEntry
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
	context_          context.Context
}

// Insert: Adds an entry to the user's calendar list.

func (r *CalendarListService) Insert(calendarlistentry *CalendarListEntry) *CalendarListInsertCall {
	return &CalendarListInsertCall{
		s:                 r.s,
		calendarlistentry: calendarlistentry,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "users/me/calendarList",
		context_:          googleapi.NoContext,
	}
}

// ColorRgbFormat sets the optional parameter "colorRgbFormat": Whether
// to use the foregroundColor and backgroundColor fields to write the
// calendar colors (RGB). If this feature is used, the index-based
// colorId field will be set to the best matching option automatically.
// The default is False.
func (c *CalendarListInsertCall) ColorRgbFormat(colorRgbFormat bool) *CalendarListInsertCall {
	c.params_.Set("colorRgbFormat", fmt.Sprintf("%v", colorRgbFormat))
	return c
}
func (c *CalendarListInsertCall) Context(ctx context.Context) *CalendarListInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarListInsertCall) Fields(s ...googleapi.Field) *CalendarListInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarListInsertCall) Do() (*CalendarListEntry, error) {
	var returnValue *CalendarListEntry
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.calendarlistentry,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Adds an entry to the user's calendar list.",
	//   "httpMethod": "POST",
	//   "id": "calendar.calendarList.insert",
	//   "parameters": {
	//     "colorRgbFormat": {
	//       "description": "Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB). If this feature is used, the index-based colorId field will be set to the best matching option automatically. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "users/me/calendarList",
	//   "request": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "response": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendarList.list":

type CalendarListListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Returns entries on the user's calendar list.

func (r *CalendarListService) List() *CalendarListListCall {
	return &CalendarListListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/calendarList",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. By default the value is 100
// entries. The page size can never be larger than 250 entries.
func (c *CalendarListListCall) MaxResults(maxResults int64) *CalendarListListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// MinAccessRole sets the optional parameter "minAccessRole": The
// minimum access role for the user in the returned entires.  The
// default is no restriction.
func (c *CalendarListListCall) MinAccessRole(minAccessRole string) *CalendarListListCall {
	c.params_.Set("minAccessRole", fmt.Sprintf("%v", minAccessRole))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *CalendarListListCall) PageToken(pageToken string) *CalendarListListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted calendar list entries in the result.  The default is
// False.
func (c *CalendarListListCall) ShowDeleted(showDeleted bool) *CalendarListListCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// ShowHidden sets the optional parameter "showHidden": Whether to show
// hidden entries.  The default is False.
func (c *CalendarListListCall) ShowHidden(showHidden bool) *CalendarListListCall {
	c.params_.Set("showHidden", fmt.Sprintf("%v", showHidden))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then. If only
// read-only fields such as calendar properties or ACLs have changed,
// the entry won't be returned. All entries deleted and hidden since the
// previous list request will always be in the result set and it is not
// allowed to set showDeleted neither showHidden to False.
// To ensure
// client state consistency minAccessRole query parameter cannot be
// specified together with nextSyncToken.
// If the syncToken expires, the
// server will respond with a 410 GONE response code and the client
// should clear its storage and perform a full synchronization without
// any syncToken.
// Learn more about incremental synchronization.
//  The
// default is to return all entries.
func (c *CalendarListListCall) SyncToken(syncToken string) *CalendarListListCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}
func (c *CalendarListListCall) Context(ctx context.Context) *CalendarListListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarListListCall) Fields(s ...googleapi.Field) *CalendarListListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarListListCall) Do() (*CalendarList, error) {
	var returnValue *CalendarList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns entries on the user's calendar list.",
	//   "httpMethod": "GET",
	//   "id": "calendar.calendarList.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "minAccessRole": {
	//       "description": "The minimum access role for the user in the returned entires. Optional. The default is no restriction.",
	//       "enum": [
	//         "freeBusyReader",
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read free/busy information.",
	//         "The user can read and modify events and access control lists.",
	//         "The user can read events that are not private.",
	//         "The user can read and modify events."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted calendar list entries in the result. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "showHidden": {
	//       "description": "Whether to show hidden entries. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. If only read-only fields such as calendar properties or ACLs have changed, the entry won't be returned. All entries deleted and hidden since the previous list request will always be in the result set and it is not allowed to set showDeleted neither showHidden to False.\nTo ensure client state consistency minAccessRole query parameter cannot be specified together with nextSyncToken.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/calendarList",
	//   "response": {
	//     "$ref": "CalendarList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.calendarList.patch":

type CalendarListPatchCall struct {
	s                 *Service
	calendarId        string
	calendarlistentry *CalendarListEntry
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
	context_          context.Context
}

// Patch: Updates an entry on the user's calendar list. This method
// supports patch semantics.

func (r *CalendarListService) Patch(calendarId string, calendarlistentry *CalendarListEntry) *CalendarListPatchCall {
	return &CalendarListPatchCall{
		s:                 r.s,
		calendarId:        calendarId,
		calendarlistentry: calendarlistentry,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "users/me/calendarList/{calendarId}",
		context_:          googleapi.NoContext,
	}
}

// ColorRgbFormat sets the optional parameter "colorRgbFormat": Whether
// to use the foregroundColor and backgroundColor fields to write the
// calendar colors (RGB). If this feature is used, the index-based
// colorId field will be set to the best matching option automatically.
// The default is False.
func (c *CalendarListPatchCall) ColorRgbFormat(colorRgbFormat bool) *CalendarListPatchCall {
	c.params_.Set("colorRgbFormat", fmt.Sprintf("%v", colorRgbFormat))
	return c
}
func (c *CalendarListPatchCall) Context(ctx context.Context) *CalendarListPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarListPatchCall) Fields(s ...googleapi.Field) *CalendarListPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarListPatchCall) Do() (*CalendarListEntry, error) {
	var returnValue *CalendarListEntry
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.calendarlistentry,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an entry on the user's calendar list. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "calendar.calendarList.patch",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "colorRgbFormat": {
	//       "description": "Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB). If this feature is used, the index-based colorId field will be set to the best matching option automatically. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "users/me/calendarList/{calendarId}",
	//   "request": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "response": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendarList.update":

type CalendarListUpdateCall struct {
	s                 *Service
	calendarId        string
	calendarlistentry *CalendarListEntry
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
	context_          context.Context
}

// Update: Updates an entry on the user's calendar list.

func (r *CalendarListService) Update(calendarId string, calendarlistentry *CalendarListEntry) *CalendarListUpdateCall {
	return &CalendarListUpdateCall{
		s:                 r.s,
		calendarId:        calendarId,
		calendarlistentry: calendarlistentry,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "users/me/calendarList/{calendarId}",
		context_:          googleapi.NoContext,
	}
}

// ColorRgbFormat sets the optional parameter "colorRgbFormat": Whether
// to use the foregroundColor and backgroundColor fields to write the
// calendar colors (RGB). If this feature is used, the index-based
// colorId field will be set to the best matching option automatically.
// The default is False.
func (c *CalendarListUpdateCall) ColorRgbFormat(colorRgbFormat bool) *CalendarListUpdateCall {
	c.params_.Set("colorRgbFormat", fmt.Sprintf("%v", colorRgbFormat))
	return c
}
func (c *CalendarListUpdateCall) Context(ctx context.Context) *CalendarListUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarListUpdateCall) Fields(s ...googleapi.Field) *CalendarListUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarListUpdateCall) Do() (*CalendarListEntry, error) {
	var returnValue *CalendarListEntry
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.calendarlistentry,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an entry on the user's calendar list.",
	//   "httpMethod": "PUT",
	//   "id": "calendar.calendarList.update",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "colorRgbFormat": {
	//       "description": "Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB). If this feature is used, the index-based colorId field will be set to the best matching option automatically. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "users/me/calendarList/{calendarId}",
	//   "request": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "response": {
	//     "$ref": "CalendarListEntry"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendarList.watch":

type CalendarListWatchCall struct {
	s             *Service
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Watch: Watch for changes to CalendarList resources.

func (r *CalendarListService) Watch(channel *Channel) *CalendarListWatchCall {
	return &CalendarListWatchCall{
		s:             r.s,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/calendarList/watch",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. By default the value is 100
// entries. The page size can never be larger than 250 entries.
func (c *CalendarListWatchCall) MaxResults(maxResults int64) *CalendarListWatchCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// MinAccessRole sets the optional parameter "minAccessRole": The
// minimum access role for the user in the returned entires.  The
// default is no restriction.
func (c *CalendarListWatchCall) MinAccessRole(minAccessRole string) *CalendarListWatchCall {
	c.params_.Set("minAccessRole", fmt.Sprintf("%v", minAccessRole))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *CalendarListWatchCall) PageToken(pageToken string) *CalendarListWatchCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted calendar list entries in the result.  The default is
// False.
func (c *CalendarListWatchCall) ShowDeleted(showDeleted bool) *CalendarListWatchCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// ShowHidden sets the optional parameter "showHidden": Whether to show
// hidden entries.  The default is False.
func (c *CalendarListWatchCall) ShowHidden(showHidden bool) *CalendarListWatchCall {
	c.params_.Set("showHidden", fmt.Sprintf("%v", showHidden))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then. If only
// read-only fields such as calendar properties or ACLs have changed,
// the entry won't be returned. All entries deleted and hidden since the
// previous list request will always be in the result set and it is not
// allowed to set showDeleted neither showHidden to False.
// To ensure
// client state consistency minAccessRole query parameter cannot be
// specified together with nextSyncToken.
// If the syncToken expires, the
// server will respond with a 410 GONE response code and the client
// should clear its storage and perform a full synchronization without
// any syncToken.
// Learn more about incremental synchronization.
//  The
// default is to return all entries.
func (c *CalendarListWatchCall) SyncToken(syncToken string) *CalendarListWatchCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}
func (c *CalendarListWatchCall) Context(ctx context.Context) *CalendarListWatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarListWatchCall) Fields(s ...googleapi.Field) *CalendarListWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarListWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Watch for changes to CalendarList resources.",
	//   "httpMethod": "POST",
	//   "id": "calendar.calendarList.watch",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "minAccessRole": {
	//       "description": "The minimum access role for the user in the returned entires. Optional. The default is no restriction.",
	//       "enum": [
	//         "freeBusyReader",
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read free/busy information.",
	//         "The user can read and modify events and access control lists.",
	//         "The user can read events that are not private.",
	//         "The user can read and modify events."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted calendar list entries in the result. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "showHidden": {
	//       "description": "Whether to show hidden entries. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. If only read-only fields such as calendar properties or ACLs have changed, the entry won't be returned. All entries deleted and hidden since the previous list request will always be in the result set and it is not allowed to set showDeleted neither showHidden to False.\nTo ensure client state consistency minAccessRole query parameter cannot be specified together with nextSyncToken.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/calendarList/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.calendars.clear":

type CalendarsClearCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Clear: Clears a primary calendar. This operation deletes all events
// associated with the primary calendar of an account.

func (r *CalendarsService) Clear(calendarId string) *CalendarsClearCall {
	return &CalendarsClearCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/clear",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarsClearCall) Context(ctx context.Context) *CalendarsClearCall {
	c.context_ = ctx
	return c
}

func (c *CalendarsClearCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Clears a primary calendar. This operation deletes all events associated with the primary calendar of an account.",
	//   "httpMethod": "POST",
	//   "id": "calendar.calendars.clear",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/clear",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendars.delete":

type CalendarsDeleteCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes a secondary calendar. Use calendars.clear for
// clearing all events on primary calendars.

func (r *CalendarsService) Delete(calendarId string) *CalendarsDeleteCall {
	return &CalendarsDeleteCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarsDeleteCall) Context(ctx context.Context) *CalendarsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *CalendarsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes a secondary calendar. Use calendars.clear for clearing all events on primary calendars.",
	//   "httpMethod": "DELETE",
	//   "id": "calendar.calendars.delete",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendars.get":

type CalendarsGetCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns metadata for a calendar.

func (r *CalendarsService) Get(calendarId string) *CalendarsGetCall {
	return &CalendarsGetCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarsGetCall) Context(ctx context.Context) *CalendarsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarsGetCall) Fields(s ...googleapi.Field) *CalendarsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarsGetCall) Do() (*Calendar, error) {
	var returnValue *Calendar
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns metadata for a calendar.",
	//   "httpMethod": "GET",
	//   "id": "calendar.calendars.get",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}",
	//   "response": {
	//     "$ref": "Calendar"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.calendars.insert":

type CalendarsInsertCall struct {
	s             *Service
	calendar      *Calendar
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates a secondary calendar.

func (r *CalendarsService) Insert(calendar *Calendar) *CalendarsInsertCall {
	return &CalendarsInsertCall{
		s:             r.s,
		calendar:      calendar,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarsInsertCall) Context(ctx context.Context) *CalendarsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarsInsertCall) Fields(s ...googleapi.Field) *CalendarsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarsInsertCall) Do() (*Calendar, error) {
	var returnValue *Calendar
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.calendar,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates a secondary calendar.",
	//   "httpMethod": "POST",
	//   "id": "calendar.calendars.insert",
	//   "path": "calendars",
	//   "request": {
	//     "$ref": "Calendar"
	//   },
	//   "response": {
	//     "$ref": "Calendar"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendars.patch":

type CalendarsPatchCall struct {
	s             *Service
	calendarId    string
	calendar      *Calendar
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates metadata for a calendar. This method supports patch
// semantics.

func (r *CalendarsService) Patch(calendarId string, calendar *Calendar) *CalendarsPatchCall {
	return &CalendarsPatchCall{
		s:             r.s,
		calendarId:    calendarId,
		calendar:      calendar,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarsPatchCall) Context(ctx context.Context) *CalendarsPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarsPatchCall) Fields(s ...googleapi.Field) *CalendarsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarsPatchCall) Do() (*Calendar, error) {
	var returnValue *Calendar
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.calendar,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates metadata for a calendar. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "calendar.calendars.patch",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}",
	//   "request": {
	//     "$ref": "Calendar"
	//   },
	//   "response": {
	//     "$ref": "Calendar"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.calendars.update":

type CalendarsUpdateCall struct {
	s             *Service
	calendarId    string
	calendar      *Calendar
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates metadata for a calendar.

func (r *CalendarsService) Update(calendarId string, calendar *Calendar) *CalendarsUpdateCall {
	return &CalendarsUpdateCall{
		s:             r.s,
		calendarId:    calendarId,
		calendar:      calendar,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}",
		context_:      googleapi.NoContext,
	}
}
func (c *CalendarsUpdateCall) Context(ctx context.Context) *CalendarsUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CalendarsUpdateCall) Fields(s ...googleapi.Field) *CalendarsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CalendarsUpdateCall) Do() (*Calendar, error) {
	var returnValue *Calendar
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.calendar,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates metadata for a calendar.",
	//   "httpMethod": "PUT",
	//   "id": "calendar.calendars.update",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}",
	//   "request": {
	//     "$ref": "Calendar"
	//   },
	//   "response": {
	//     "$ref": "Calendar"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.channels.stop":

type ChannelsStopCall struct {
	s             *Service
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Stop: Stop watching resources through this channel

func (r *ChannelsService) Stop(channel *Channel) *ChannelsStopCall {
	return &ChannelsStopCall{
		s:             r.s,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "channels/stop",
		context_:      googleapi.NoContext,
	}
}
func (c *ChannelsStopCall) Context(ctx context.Context) *ChannelsStopCall {
	c.context_ = ctx
	return c
}

func (c *ChannelsStopCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Stop watching resources through this channel",
	//   "httpMethod": "POST",
	//   "id": "calendar.channels.stop",
	//   "path": "channels/stop",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.colors.get":

type ColorsGetCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns the color definitions for calendars and events.

func (r *ColorsService) Get() *ColorsGetCall {
	return &ColorsGetCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "colors",
		context_:      googleapi.NoContext,
	}
}
func (c *ColorsGetCall) Context(ctx context.Context) *ColorsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColorsGetCall) Fields(s ...googleapi.Field) *ColorsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ColorsGetCall) Do() (*Colors, error) {
	var returnValue *Colors
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns the color definitions for calendars and events.",
	//   "httpMethod": "GET",
	//   "id": "calendar.colors.get",
	//   "path": "colors",
	//   "response": {
	//     "$ref": "Colors"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.events.delete":

type EventsDeleteCall struct {
	s             *Service
	calendarId    string
	eventId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Delete: Deletes an event.

func (r *EventsService) Delete(calendarId string, eventId string) *EventsDeleteCall {
	return &EventsDeleteCall{
		s:             r.s,
		calendarId:    calendarId,
		eventId:       eventId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/{eventId}",
		context_:      googleapi.NoContext,
	}
}

// SendNotifications sets the optional parameter "sendNotifications":
// Whether to send notifications about the deletion of the event.  The
// default is False.
func (c *EventsDeleteCall) SendNotifications(sendNotifications bool) *EventsDeleteCall {
	c.params_.Set("sendNotifications", fmt.Sprintf("%v", sendNotifications))
	return c
}
func (c *EventsDeleteCall) Context(ctx context.Context) *EventsDeleteCall {
	c.context_ = ctx
	return c
}

func (c *EventsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"eventId":    c.eventId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Deletes an event.",
	//   "httpMethod": "DELETE",
	//   "id": "calendar.events.delete",
	//   "parameterOrder": [
	//     "calendarId",
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "eventId": {
	//       "description": "Event identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sendNotifications": {
	//       "description": "Whether to send notifications about the deletion of the event. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/{eventId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.get":

type EventsGetCall struct {
	s             *Service
	calendarId    string
	eventId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns an event.

func (r *EventsService) Get(calendarId string, eventId string) *EventsGetCall {
	return &EventsGetCall{
		s:             r.s,
		calendarId:    calendarId,
		eventId:       eventId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/{eventId}",
		context_:      googleapi.NoContext,
	}
}

// AlwaysIncludeEmail sets the optional parameter "alwaysIncludeEmail":
// Whether to always include a value in the email field for the
// organizer, creator and attendees, even if no real email is available
// (i.e. a generated, non-working value will be provided). The use of
// this option is discouraged and should only be used by clients which
// cannot handle the absence of an email address value in the mentioned
// places.  The default is False.
func (c *EventsGetCall) AlwaysIncludeEmail(alwaysIncludeEmail bool) *EventsGetCall {
	c.params_.Set("alwaysIncludeEmail", fmt.Sprintf("%v", alwaysIncludeEmail))
	return c
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsGetCall) MaxAttendees(maxAttendees int64) *EventsGetCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// TimeZone sets the optional parameter "timeZone": Time zone used in
// the response.  The default is the time zone of the calendar.
func (c *EventsGetCall) TimeZone(timeZone string) *EventsGetCall {
	c.params_.Set("timeZone", fmt.Sprintf("%v", timeZone))
	return c
}
func (c *EventsGetCall) Context(ctx context.Context) *EventsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsGetCall) Fields(s ...googleapi.Field) *EventsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsGetCall) Do() (*Event, error) {
	var returnValue *Event
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"eventId":    c.eventId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns an event.",
	//   "httpMethod": "GET",
	//   "id": "calendar.events.get",
	//   "parameterOrder": [
	//     "calendarId",
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "alwaysIncludeEmail": {
	//       "description": "Whether to always include a value in the email field for the organizer, creator and attendees, even if no real email is available (i.e. a generated, non-working value will be provided). The use of this option is discouraged and should only be used by clients which cannot handle the absence of an email address value in the mentioned places. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "eventId": {
	//       "description": "Event identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "timeZone": {
	//       "description": "Time zone used in the response. Optional. The default is the time zone of the calendar.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/{eventId}",
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.events.import":

type EventsImportCall struct {
	s             *Service
	calendarId    string
	event         *Event
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Import: Imports an event. This operation is used to add a private
// copy of an existing event to a calendar.

func (r *EventsService) Import(calendarId string, event *Event) *EventsImportCall {
	return &EventsImportCall{
		s:             r.s,
		calendarId:    calendarId,
		event:         event,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/import",
		context_:      googleapi.NoContext,
	}
}
func (c *EventsImportCall) Context(ctx context.Context) *EventsImportCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsImportCall) Fields(s ...googleapi.Field) *EventsImportCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsImportCall) Do() (*Event, error) {
	var returnValue *Event
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.event,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Imports an event. This operation is used to add a private copy of an existing event to a calendar.",
	//   "httpMethod": "POST",
	//   "id": "calendar.events.import",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/import",
	//   "request": {
	//     "$ref": "Event"
	//   },
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.insert":

type EventsInsertCall struct {
	s             *Service
	calendarId    string
	event         *Event
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Insert: Creates an event.

func (r *EventsService) Insert(calendarId string, event *Event) *EventsInsertCall {
	return &EventsInsertCall{
		s:             r.s,
		calendarId:    calendarId,
		event:         event,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events",
		context_:      googleapi.NoContext,
	}
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsInsertCall) MaxAttendees(maxAttendees int64) *EventsInsertCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// SendNotifications sets the optional parameter "sendNotifications":
// Whether to send notifications about the creation of the new event.
// The default is False.
func (c *EventsInsertCall) SendNotifications(sendNotifications bool) *EventsInsertCall {
	c.params_.Set("sendNotifications", fmt.Sprintf("%v", sendNotifications))
	return c
}
func (c *EventsInsertCall) Context(ctx context.Context) *EventsInsertCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsInsertCall) Fields(s ...googleapi.Field) *EventsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsInsertCall) Do() (*Event, error) {
	var returnValue *Event
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.event,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an event.",
	//   "httpMethod": "POST",
	//   "id": "calendar.events.insert",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "sendNotifications": {
	//       "description": "Whether to send notifications about the creation of the new event. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events",
	//   "request": {
	//     "$ref": "Event"
	//   },
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.instances":

type EventsInstancesCall struct {
	s             *Service
	calendarId    string
	eventId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Instances: Returns instances of the specified recurring event.

func (r *EventsService) Instances(calendarId string, eventId string) *EventsInstancesCall {
	return &EventsInstancesCall{
		s:             r.s,
		calendarId:    calendarId,
		eventId:       eventId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/{eventId}/instances",
		context_:      googleapi.NoContext,
	}
}

// AlwaysIncludeEmail sets the optional parameter "alwaysIncludeEmail":
// Whether to always include a value in the email field for the
// organizer, creator and attendees, even if no real email is available
// (i.e. a generated, non-working value will be provided). The use of
// this option is discouraged and should only be used by clients which
// cannot handle the absence of an email address value in the mentioned
// places.  The default is False.
func (c *EventsInstancesCall) AlwaysIncludeEmail(alwaysIncludeEmail bool) *EventsInstancesCall {
	c.params_.Set("alwaysIncludeEmail", fmt.Sprintf("%v", alwaysIncludeEmail))
	return c
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsInstancesCall) MaxAttendees(maxAttendees int64) *EventsInstancesCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of events returned on one result page. By default the value is 250
// events. The page size can never be larger than 2500 events.
func (c *EventsInstancesCall) MaxResults(maxResults int64) *EventsInstancesCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OriginalStart sets the optional parameter "originalStart": The
// original start time of the instance in the result.
func (c *EventsInstancesCall) OriginalStart(originalStart string) *EventsInstancesCall {
	c.params_.Set("originalStart", fmt.Sprintf("%v", originalStart))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *EventsInstancesCall) PageToken(pageToken string) *EventsInstancesCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted events (with status equals "cancelled") in the
// result. Cancelled instances of recurring events will still be
// included if singleEvents is False.  The default is False.
func (c *EventsInstancesCall) ShowDeleted(showDeleted bool) *EventsInstancesCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// TimeMax sets the optional parameter "timeMax": Upper bound
// (exclusive) for an event's start time to filter by.  The default is
// not to filter by start time.
func (c *EventsInstancesCall) TimeMax(timeMax string) *EventsInstancesCall {
	c.params_.Set("timeMax", fmt.Sprintf("%v", timeMax))
	return c
}

// TimeMin sets the optional parameter "timeMin": Lower bound
// (inclusive) for an event's end time to filter by.  The default is not
// to filter by end time.
func (c *EventsInstancesCall) TimeMin(timeMin string) *EventsInstancesCall {
	c.params_.Set("timeMin", fmt.Sprintf("%v", timeMin))
	return c
}

// TimeZone sets the optional parameter "timeZone": Time zone used in
// the response.  The default is the time zone of the calendar.
func (c *EventsInstancesCall) TimeZone(timeZone string) *EventsInstancesCall {
	c.params_.Set("timeZone", fmt.Sprintf("%v", timeZone))
	return c
}
func (c *EventsInstancesCall) Context(ctx context.Context) *EventsInstancesCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsInstancesCall) Fields(s ...googleapi.Field) *EventsInstancesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsInstancesCall) Do() (*Events, error) {
	var returnValue *Events
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"eventId":    c.eventId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns instances of the specified recurring event.",
	//   "httpMethod": "GET",
	//   "id": "calendar.events.instances",
	//   "parameterOrder": [
	//     "calendarId",
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "alwaysIncludeEmail": {
	//       "description": "Whether to always include a value in the email field for the organizer, creator and attendees, even if no real email is available (i.e. a generated, non-working value will be provided). The use of this option is discouraged and should only be used by clients which cannot handle the absence of an email address value in the mentioned places. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "eventId": {
	//       "description": "Recurring event identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of events returned on one result page. By default the value is 250 events. The page size can never be larger than 2500 events. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "originalStart": {
	//       "description": "The original start time of the instance in the result. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted events (with status equals \"cancelled\") in the result. Cancelled instances of recurring events will still be included if singleEvents is False. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "timeMax": {
	//       "description": "Upper bound (exclusive) for an event's start time to filter by. Optional. The default is not to filter by start time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeMin": {
	//       "description": "Lower bound (inclusive) for an event's end time to filter by. Optional. The default is not to filter by end time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeZone": {
	//       "description": "Time zone used in the response. Optional. The default is the time zone of the calendar.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/{eventId}/instances",
	//   "response": {
	//     "$ref": "Events"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.events.list":

type EventsListCall struct {
	s             *Service
	calendarId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Returns events on the specified calendar.

func (r *EventsService) List(calendarId string) *EventsListCall {
	return &EventsListCall{
		s:             r.s,
		calendarId:    calendarId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events",
		context_:      googleapi.NoContext,
	}
}

// AlwaysIncludeEmail sets the optional parameter "alwaysIncludeEmail":
// Whether to always include a value in the email field for the
// organizer, creator and attendees, even if no real email is available
// (i.e. a generated, non-working value will be provided). The use of
// this option is discouraged and should only be used by clients which
// cannot handle the absence of an email address value in the mentioned
// places.  The default is False.
func (c *EventsListCall) AlwaysIncludeEmail(alwaysIncludeEmail bool) *EventsListCall {
	c.params_.Set("alwaysIncludeEmail", fmt.Sprintf("%v", alwaysIncludeEmail))
	return c
}

// ICalUID sets the optional parameter "iCalUID": Specifies event ID in
// the iCalendar format to be included in the response.
func (c *EventsListCall) ICalUID(iCalUID string) *EventsListCall {
	c.params_.Set("iCalUID", fmt.Sprintf("%v", iCalUID))
	return c
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsListCall) MaxAttendees(maxAttendees int64) *EventsListCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of events returned on one result page. By default the value is 250
// events. The page size can never be larger than 2500 events.
func (c *EventsListCall) MaxResults(maxResults int64) *EventsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": The order of the
// events returned in the result.  The default is an unspecified, stable
// order.
func (c *EventsListCall) OrderBy(orderBy string) *EventsListCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *EventsListCall) PageToken(pageToken string) *EventsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// PrivateExtendedProperty sets the optional parameter
// "privateExtendedProperty": Extended properties constraint specified
// as propertyName=value. Matches only private properties. This
// parameter might be repeated multiple times to return events that
// match all given constraints.
func (c *EventsListCall) PrivateExtendedProperty(privateExtendedProperty ...string) *EventsListCall {
	c.params_["privateExtendedProperty"] = privateExtendedProperty
	return c
}

// Q sets the optional parameter "q": Free text search terms to find
// events that match these terms in any field, except for extended
// properties.
func (c *EventsListCall) Q(q string) *EventsListCall {
	c.params_.Set("q", fmt.Sprintf("%v", q))
	return c
}

// SharedExtendedProperty sets the optional parameter
// "sharedExtendedProperty": Extended properties constraint specified as
// propertyName=value. Matches only shared properties. This parameter
// might be repeated multiple times to return events that match all
// given constraints.
func (c *EventsListCall) SharedExtendedProperty(sharedExtendedProperty ...string) *EventsListCall {
	c.params_["sharedExtendedProperty"] = sharedExtendedProperty
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted events (with status equals "cancelled") in the
// result. Cancelled instances of recurring events (but not the
// underlying recurring event) will still be included if showDeleted and
// singleEvents are both False. If showDeleted and singleEvents are both
// True, only single instances of deleted events (but not the underlying
// recurring events) are returned.  The default is False.
func (c *EventsListCall) ShowDeleted(showDeleted bool) *EventsListCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// ShowHiddenInvitations sets the optional parameter
// "showHiddenInvitations": Whether to include hidden invitations in the
// result.  The default is False.
func (c *EventsListCall) ShowHiddenInvitations(showHiddenInvitations bool) *EventsListCall {
	c.params_.Set("showHiddenInvitations", fmt.Sprintf("%v", showHiddenInvitations))
	return c
}

// SingleEvents sets the optional parameter "singleEvents": Whether to
// expand recurring events into instances and only return single one-off
// events and instances of recurring events, but not the underlying
// recurring events themselves.  The default is False.
func (c *EventsListCall) SingleEvents(singleEvents bool) *EventsListCall {
	c.params_.Set("singleEvents", fmt.Sprintf("%v", singleEvents))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then. All events
// deleted since the previous list request will always be in the result
// set and it is not allowed to set showDeleted to False.
// There are
// several query parameters that cannot be specified together with
// nextSyncToken to ensure consistency of the client state.
//
// These are:
//
// - iCalUID
// - orderBy
// - privateExtendedProperty
// - q
// -
// sharedExtendedProperty
// - timeMin
// - timeMax
// - updatedMin If the
// syncToken expires, the server will respond with a 410 GONE response
// code and the client should clear its storage and perform a full
// synchronization without any syncToken.
// Learn more about incremental
// synchronization.
//  The default is to return all entries.
func (c *EventsListCall) SyncToken(syncToken string) *EventsListCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}

// TimeMax sets the optional parameter "timeMax": Upper bound
// (exclusive) for an event's start time to filter by.  The default is
// not to filter by start time.
func (c *EventsListCall) TimeMax(timeMax string) *EventsListCall {
	c.params_.Set("timeMax", fmt.Sprintf("%v", timeMax))
	return c
}

// TimeMin sets the optional parameter "timeMin": Lower bound
// (inclusive) for an event's end time to filter by.  The default is not
// to filter by end time.
func (c *EventsListCall) TimeMin(timeMin string) *EventsListCall {
	c.params_.Set("timeMin", fmt.Sprintf("%v", timeMin))
	return c
}

// TimeZone sets the optional parameter "timeZone": Time zone used in
// the response.  The default is the time zone of the calendar.
func (c *EventsListCall) TimeZone(timeZone string) *EventsListCall {
	c.params_.Set("timeZone", fmt.Sprintf("%v", timeZone))
	return c
}

// UpdatedMin sets the optional parameter "updatedMin": Lower bound for
// an event's last modification time (as a RFC 3339 timestamp) to filter
// by. When specified, entries deleted since this time will always be
// included regardless of showDeleted.  The default is not to filter by
// last modification time.
func (c *EventsListCall) UpdatedMin(updatedMin string) *EventsListCall {
	c.params_.Set("updatedMin", fmt.Sprintf("%v", updatedMin))
	return c
}
func (c *EventsListCall) Context(ctx context.Context) *EventsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsListCall) Fields(s ...googleapi.Field) *EventsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsListCall) Do() (*Events, error) {
	var returnValue *Events
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns events on the specified calendar.",
	//   "httpMethod": "GET",
	//   "id": "calendar.events.list",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "alwaysIncludeEmail": {
	//       "description": "Whether to always include a value in the email field for the organizer, creator and attendees, even if no real email is available (i.e. a generated, non-working value will be provided). The use of this option is discouraged and should only be used by clients which cannot handle the absence of an email address value in the mentioned places. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "iCalUID": {
	//       "description": "Specifies event ID in the iCalendar format to be included in the response. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of events returned on one result page. By default the value is 250 events. The page size can never be larger than 2500 events. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "The order of the events returned in the result. Optional. The default is an unspecified, stable order.",
	//       "enum": [
	//         "startTime",
	//         "updated"
	//       ],
	//       "enumDescriptions": [
	//         "Order by the start date/time (ascending). This is only available when querying single events (i.e. the parameter singleEvents is True)",
	//         "Order by last modification time (ascending)."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "privateExtendedProperty": {
	//       "description": "Extended properties constraint specified as propertyName=value. Matches only private properties. This parameter might be repeated multiple times to return events that match all given constraints.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Free text search terms to find events that match these terms in any field, except for extended properties. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sharedExtendedProperty": {
	//       "description": "Extended properties constraint specified as propertyName=value. Matches only shared properties. This parameter might be repeated multiple times to return events that match all given constraints.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted events (with status equals \"cancelled\") in the result. Cancelled instances of recurring events (but not the underlying recurring event) will still be included if showDeleted and singleEvents are both False. If showDeleted and singleEvents are both True, only single instances of deleted events (but not the underlying recurring events) are returned. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "showHiddenInvitations": {
	//       "description": "Whether to include hidden invitations in the result. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "singleEvents": {
	//       "description": "Whether to expand recurring events into instances and only return single one-off events and instances of recurring events, but not the underlying recurring events themselves. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. All events deleted since the previous list request will always be in the result set and it is not allowed to set showDeleted to False.\nThere are several query parameters that cannot be specified together with nextSyncToken to ensure consistency of the client state.\n\nThese are: \n- iCalUID \n- orderBy \n- privateExtendedProperty \n- q \n- sharedExtendedProperty \n- timeMin \n- timeMax \n- updatedMin If the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeMax": {
	//       "description": "Upper bound (exclusive) for an event's start time to filter by. Optional. The default is not to filter by start time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeMin": {
	//       "description": "Lower bound (inclusive) for an event's end time to filter by. Optional. The default is not to filter by end time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeZone": {
	//       "description": "Time zone used in the response. Optional. The default is the time zone of the calendar.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "updatedMin": {
	//       "description": "Lower bound for an event's last modification time (as a RFC 3339 timestamp) to filter by. When specified, entries deleted since this time will always be included regardless of showDeleted. Optional. The default is not to filter by last modification time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events",
	//   "response": {
	//     "$ref": "Events"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.events.move":

type EventsMoveCall struct {
	s             *Service
	calendarId    string
	eventId       string
	destinationid string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Move: Moves an event to another calendar, i.e. changes an event's
// organizer.

func (r *EventsService) Move(calendarId string, eventId string, destinationid string) *EventsMoveCall {
	return &EventsMoveCall{
		s:             r.s,
		calendarId:    calendarId,
		eventId:       eventId,
		destinationid: destinationid,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/{eventId}/move",
		context_:      googleapi.NoContext,
	}
}

// SendNotifications sets the optional parameter "sendNotifications":
// Whether to send notifications about the change of the event's
// organizer.  The default is False.
func (c *EventsMoveCall) SendNotifications(sendNotifications bool) *EventsMoveCall {
	c.params_.Set("sendNotifications", fmt.Sprintf("%v", sendNotifications))
	return c
}
func (c *EventsMoveCall) Context(ctx context.Context) *EventsMoveCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsMoveCall) Fields(s ...googleapi.Field) *EventsMoveCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsMoveCall) Do() (*Event, error) {
	var returnValue *Event
	c.params_.Set("destination", fmt.Sprintf("%v", c.destinationid))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"eventId":    c.eventId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Moves an event to another calendar, i.e. changes an event's organizer.",
	//   "httpMethod": "POST",
	//   "id": "calendar.events.move",
	//   "parameterOrder": [
	//     "calendarId",
	//     "eventId",
	//     "destination"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier of the source calendar where the event currently is on.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destination": {
	//       "description": "Calendar identifier of the target calendar where the event is to be moved to.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "eventId": {
	//       "description": "Event identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sendNotifications": {
	//       "description": "Whether to send notifications about the change of the event's organizer. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/{eventId}/move",
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.patch":

type EventsPatchCall struct {
	s             *Service
	calendarId    string
	eventId       string
	event         *Event
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Patch: Updates an event. This method supports patch semantics.

func (r *EventsService) Patch(calendarId string, eventId string, event *Event) *EventsPatchCall {
	return &EventsPatchCall{
		s:             r.s,
		calendarId:    calendarId,
		eventId:       eventId,
		event:         event,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/{eventId}",
		context_:      googleapi.NoContext,
	}
}

// AlwaysIncludeEmail sets the optional parameter "alwaysIncludeEmail":
// Whether to always include a value in the email field for the
// organizer, creator and attendees, even if no real email is available
// (i.e. a generated, non-working value will be provided). The use of
// this option is discouraged and should only be used by clients which
// cannot handle the absence of an email address value in the mentioned
// places.  The default is False.
func (c *EventsPatchCall) AlwaysIncludeEmail(alwaysIncludeEmail bool) *EventsPatchCall {
	c.params_.Set("alwaysIncludeEmail", fmt.Sprintf("%v", alwaysIncludeEmail))
	return c
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsPatchCall) MaxAttendees(maxAttendees int64) *EventsPatchCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// SendNotifications sets the optional parameter "sendNotifications":
// Whether to send notifications about the event update (e.g. attendee's
// responses, title changes, etc.).  The default is False.
func (c *EventsPatchCall) SendNotifications(sendNotifications bool) *EventsPatchCall {
	c.params_.Set("sendNotifications", fmt.Sprintf("%v", sendNotifications))
	return c
}
func (c *EventsPatchCall) Context(ctx context.Context) *EventsPatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsPatchCall) Fields(s ...googleapi.Field) *EventsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsPatchCall) Do() (*Event, error) {
	var returnValue *Event
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"eventId":    c.eventId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.event,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an event. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "calendar.events.patch",
	//   "parameterOrder": [
	//     "calendarId",
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "alwaysIncludeEmail": {
	//       "description": "Whether to always include a value in the email field for the organizer, creator and attendees, even if no real email is available (i.e. a generated, non-working value will be provided). The use of this option is discouraged and should only be used by clients which cannot handle the absence of an email address value in the mentioned places. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "eventId": {
	//       "description": "Event identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "sendNotifications": {
	//       "description": "Whether to send notifications about the event update (e.g. attendee's responses, title changes, etc.). Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/{eventId}",
	//   "request": {
	//     "$ref": "Event"
	//   },
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.quickAdd":

type EventsQuickAddCall struct {
	s             *Service
	calendarId    string
	text          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// QuickAdd: Creates an event based on a simple text string.

func (r *EventsService) QuickAdd(calendarId string, text string) *EventsQuickAddCall {
	return &EventsQuickAddCall{
		s:             r.s,
		calendarId:    calendarId,
		text:          text,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/quickAdd",
		context_:      googleapi.NoContext,
	}
}

// SendNotifications sets the optional parameter "sendNotifications":
// Whether to send notifications about the creation of the event.  The
// default is False.
func (c *EventsQuickAddCall) SendNotifications(sendNotifications bool) *EventsQuickAddCall {
	c.params_.Set("sendNotifications", fmt.Sprintf("%v", sendNotifications))
	return c
}
func (c *EventsQuickAddCall) Context(ctx context.Context) *EventsQuickAddCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsQuickAddCall) Fields(s ...googleapi.Field) *EventsQuickAddCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsQuickAddCall) Do() (*Event, error) {
	var returnValue *Event
	c.params_.Set("text", fmt.Sprintf("%v", c.text))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Creates an event based on a simple text string.",
	//   "httpMethod": "POST",
	//   "id": "calendar.events.quickAdd",
	//   "parameterOrder": [
	//     "calendarId",
	//     "text"
	//   ],
	//   "parameters": {
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sendNotifications": {
	//       "description": "Whether to send notifications about the creation of the event. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "text": {
	//       "description": "The text describing the event to be created.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/quickAdd",
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.update":

type EventsUpdateCall struct {
	s             *Service
	calendarId    string
	eventId       string
	event         *Event
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Update: Updates an event.

func (r *EventsService) Update(calendarId string, eventId string, event *Event) *EventsUpdateCall {
	return &EventsUpdateCall{
		s:             r.s,
		calendarId:    calendarId,
		eventId:       eventId,
		event:         event,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/{eventId}",
		context_:      googleapi.NoContext,
	}
}

// AlwaysIncludeEmail sets the optional parameter "alwaysIncludeEmail":
// Whether to always include a value in the email field for the
// organizer, creator and attendees, even if no real email is available
// (i.e. a generated, non-working value will be provided). The use of
// this option is discouraged and should only be used by clients which
// cannot handle the absence of an email address value in the mentioned
// places.  The default is False.
func (c *EventsUpdateCall) AlwaysIncludeEmail(alwaysIncludeEmail bool) *EventsUpdateCall {
	c.params_.Set("alwaysIncludeEmail", fmt.Sprintf("%v", alwaysIncludeEmail))
	return c
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsUpdateCall) MaxAttendees(maxAttendees int64) *EventsUpdateCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// SendNotifications sets the optional parameter "sendNotifications":
// Whether to send notifications about the event update (e.g. attendee's
// responses, title changes, etc.).  The default is False.
func (c *EventsUpdateCall) SendNotifications(sendNotifications bool) *EventsUpdateCall {
	c.params_.Set("sendNotifications", fmt.Sprintf("%v", sendNotifications))
	return c
}
func (c *EventsUpdateCall) Context(ctx context.Context) *EventsUpdateCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsUpdateCall) Fields(s ...googleapi.Field) *EventsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsUpdateCall) Do() (*Event, error) {
	var returnValue *Event
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
		"eventId":    c.eventId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.event,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates an event.",
	//   "httpMethod": "PUT",
	//   "id": "calendar.events.update",
	//   "parameterOrder": [
	//     "calendarId",
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "alwaysIncludeEmail": {
	//       "description": "Whether to always include a value in the email field for the organizer, creator and attendees, even if no real email is available (i.e. a generated, non-working value will be provided). The use of this option is discouraged and should only be used by clients which cannot handle the absence of an email address value in the mentioned places. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "eventId": {
	//       "description": "Event identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "sendNotifications": {
	//       "description": "Whether to send notifications about the event update (e.g. attendee's responses, title changes, etc.). Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/{eventId}",
	//   "request": {
	//     "$ref": "Event"
	//   },
	//   "response": {
	//     "$ref": "Event"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar"
	//   ]
	// }

}

// method id "calendar.events.watch":

type EventsWatchCall struct {
	s             *Service
	calendarId    string
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Watch: Watch for changes to Events resources.

func (r *EventsService) Watch(calendarId string, channel *Channel) *EventsWatchCall {
	return &EventsWatchCall{
		s:             r.s,
		calendarId:    calendarId,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "calendars/{calendarId}/events/watch",
		context_:      googleapi.NoContext,
	}
}

// AlwaysIncludeEmail sets the optional parameter "alwaysIncludeEmail":
// Whether to always include a value in the email field for the
// organizer, creator and attendees, even if no real email is available
// (i.e. a generated, non-working value will be provided). The use of
// this option is discouraged and should only be used by clients which
// cannot handle the absence of an email address value in the mentioned
// places.  The default is False.
func (c *EventsWatchCall) AlwaysIncludeEmail(alwaysIncludeEmail bool) *EventsWatchCall {
	c.params_.Set("alwaysIncludeEmail", fmt.Sprintf("%v", alwaysIncludeEmail))
	return c
}

// ICalUID sets the optional parameter "iCalUID": Specifies event ID in
// the iCalendar format to be included in the response.
func (c *EventsWatchCall) ICalUID(iCalUID string) *EventsWatchCall {
	c.params_.Set("iCalUID", fmt.Sprintf("%v", iCalUID))
	return c
}

// MaxAttendees sets the optional parameter "maxAttendees": The maximum
// number of attendees to include in the response. If there are more
// than the specified number of attendees, only the participant is
// returned.
func (c *EventsWatchCall) MaxAttendees(maxAttendees int64) *EventsWatchCall {
	c.params_.Set("maxAttendees", fmt.Sprintf("%v", maxAttendees))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of events returned on one result page. By default the value is 250
// events. The page size can never be larger than 2500 events.
func (c *EventsWatchCall) MaxResults(maxResults int64) *EventsWatchCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": The order of the
// events returned in the result.  The default is an unspecified, stable
// order.
func (c *EventsWatchCall) OrderBy(orderBy string) *EventsWatchCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *EventsWatchCall) PageToken(pageToken string) *EventsWatchCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// PrivateExtendedProperty sets the optional parameter
// "privateExtendedProperty": Extended properties constraint specified
// as propertyName=value. Matches only private properties. This
// parameter might be repeated multiple times to return events that
// match all given constraints.
func (c *EventsWatchCall) PrivateExtendedProperty(privateExtendedProperty ...string) *EventsWatchCall {
	c.params_["privateExtendedProperty"] = privateExtendedProperty
	return c
}

// Q sets the optional parameter "q": Free text search terms to find
// events that match these terms in any field, except for extended
// properties.
func (c *EventsWatchCall) Q(q string) *EventsWatchCall {
	c.params_.Set("q", fmt.Sprintf("%v", q))
	return c
}

// SharedExtendedProperty sets the optional parameter
// "sharedExtendedProperty": Extended properties constraint specified as
// propertyName=value. Matches only shared properties. This parameter
// might be repeated multiple times to return events that match all
// given constraints.
func (c *EventsWatchCall) SharedExtendedProperty(sharedExtendedProperty ...string) *EventsWatchCall {
	c.params_["sharedExtendedProperty"] = sharedExtendedProperty
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Whether to
// include deleted events (with status equals "cancelled") in the
// result. Cancelled instances of recurring events (but not the
// underlying recurring event) will still be included if showDeleted and
// singleEvents are both False. If showDeleted and singleEvents are both
// True, only single instances of deleted events (but not the underlying
// recurring events) are returned.  The default is False.
func (c *EventsWatchCall) ShowDeleted(showDeleted bool) *EventsWatchCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// ShowHiddenInvitations sets the optional parameter
// "showHiddenInvitations": Whether to include hidden invitations in the
// result.  The default is False.
func (c *EventsWatchCall) ShowHiddenInvitations(showHiddenInvitations bool) *EventsWatchCall {
	c.params_.Set("showHiddenInvitations", fmt.Sprintf("%v", showHiddenInvitations))
	return c
}

// SingleEvents sets the optional parameter "singleEvents": Whether to
// expand recurring events into instances and only return single one-off
// events and instances of recurring events, but not the underlying
// recurring events themselves.  The default is False.
func (c *EventsWatchCall) SingleEvents(singleEvents bool) *EventsWatchCall {
	c.params_.Set("singleEvents", fmt.Sprintf("%v", singleEvents))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then. All events
// deleted since the previous list request will always be in the result
// set and it is not allowed to set showDeleted to False.
// There are
// several query parameters that cannot be specified together with
// nextSyncToken to ensure consistency of the client state.
//
// These are:
//
// - iCalUID
// - orderBy
// - privateExtendedProperty
// - q
// -
// sharedExtendedProperty
// - timeMin
// - timeMax
// - updatedMin If the
// syncToken expires, the server will respond with a 410 GONE response
// code and the client should clear its storage and perform a full
// synchronization without any syncToken.
// Learn more about incremental
// synchronization.
//  The default is to return all entries.
func (c *EventsWatchCall) SyncToken(syncToken string) *EventsWatchCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}

// TimeMax sets the optional parameter "timeMax": Upper bound
// (exclusive) for an event's start time to filter by.  The default is
// not to filter by start time.
func (c *EventsWatchCall) TimeMax(timeMax string) *EventsWatchCall {
	c.params_.Set("timeMax", fmt.Sprintf("%v", timeMax))
	return c
}

// TimeMin sets the optional parameter "timeMin": Lower bound
// (inclusive) for an event's end time to filter by.  The default is not
// to filter by end time.
func (c *EventsWatchCall) TimeMin(timeMin string) *EventsWatchCall {
	c.params_.Set("timeMin", fmt.Sprintf("%v", timeMin))
	return c
}

// TimeZone sets the optional parameter "timeZone": Time zone used in
// the response.  The default is the time zone of the calendar.
func (c *EventsWatchCall) TimeZone(timeZone string) *EventsWatchCall {
	c.params_.Set("timeZone", fmt.Sprintf("%v", timeZone))
	return c
}

// UpdatedMin sets the optional parameter "updatedMin": Lower bound for
// an event's last modification time (as a RFC 3339 timestamp) to filter
// by. When specified, entries deleted since this time will always be
// included regardless of showDeleted.  The default is not to filter by
// last modification time.
func (c *EventsWatchCall) UpdatedMin(updatedMin string) *EventsWatchCall {
	c.params_.Set("updatedMin", fmt.Sprintf("%v", updatedMin))
	return c
}
func (c *EventsWatchCall) Context(ctx context.Context) *EventsWatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventsWatchCall) Fields(s ...googleapi.Field) *EventsWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventsWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"calendarId": c.calendarId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Watch for changes to Events resources.",
	//   "httpMethod": "POST",
	//   "id": "calendar.events.watch",
	//   "parameterOrder": [
	//     "calendarId"
	//   ],
	//   "parameters": {
	//     "alwaysIncludeEmail": {
	//       "description": "Whether to always include a value in the email field for the organizer, creator and attendees, even if no real email is available (i.e. a generated, non-working value will be provided). The use of this option is discouraged and should only be used by clients which cannot handle the absence of an email address value in the mentioned places. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "calendarId": {
	//       "description": "Calendar identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "iCalUID": {
	//       "description": "Specifies event ID in the iCalendar format to be included in the response. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxAttendees": {
	//       "description": "The maximum number of attendees to include in the response. If there are more than the specified number of attendees, only the participant is returned. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of events returned on one result page. By default the value is 250 events. The page size can never be larger than 2500 events. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "The order of the events returned in the result. Optional. The default is an unspecified, stable order.",
	//       "enum": [
	//         "startTime",
	//         "updated"
	//       ],
	//       "enumDescriptions": [
	//         "Order by the start date/time (ascending). This is only available when querying single events (i.e. the parameter singleEvents is True)",
	//         "Order by last modification time (ascending)."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "privateExtendedProperty": {
	//       "description": "Extended properties constraint specified as propertyName=value. Matches only private properties. This parameter might be repeated multiple times to return events that match all given constraints.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Free text search terms to find events that match these terms in any field, except for extended properties. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sharedExtendedProperty": {
	//       "description": "Extended properties constraint specified as propertyName=value. Matches only shared properties. This parameter might be repeated multiple times to return events that match all given constraints.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Whether to include deleted events (with status equals \"cancelled\") in the result. Cancelled instances of recurring events (but not the underlying recurring event) will still be included if showDeleted and singleEvents are both False. If showDeleted and singleEvents are both True, only single instances of deleted events (but not the underlying recurring events) are returned. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "showHiddenInvitations": {
	//       "description": "Whether to include hidden invitations in the result. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "singleEvents": {
	//       "description": "Whether to expand recurring events into instances and only return single one-off events and instances of recurring events, but not the underlying recurring events themselves. Optional. The default is False.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then. All events deleted since the previous list request will always be in the result set and it is not allowed to set showDeleted to False.\nThere are several query parameters that cannot be specified together with nextSyncToken to ensure consistency of the client state.\n\nThese are: \n- iCalUID \n- orderBy \n- privateExtendedProperty \n- q \n- sharedExtendedProperty \n- timeMin \n- timeMax \n- updatedMin If the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeMax": {
	//       "description": "Upper bound (exclusive) for an event's start time to filter by. Optional. The default is not to filter by start time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeMin": {
	//       "description": "Lower bound (inclusive) for an event's end time to filter by. Optional. The default is not to filter by end time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeZone": {
	//       "description": "Time zone used in the response. Optional. The default is the time zone of the calendar.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "updatedMin": {
	//       "description": "Lower bound for an event's last modification time (as a RFC 3339 timestamp) to filter by. When specified, entries deleted since this time will always be included regardless of showDeleted. Optional. The default is not to filter by last modification time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "calendars/{calendarId}/events/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.freebusy.query":

type FreebusyQueryCall struct {
	s               *Service
	freebusyrequest *FreeBusyRequest
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
	context_        context.Context
}

// Query: Returns free/busy information for a set of calendars.

func (r *FreebusyService) Query(freebusyrequest *FreeBusyRequest) *FreebusyQueryCall {
	return &FreebusyQueryCall{
		s:               r.s,
		freebusyrequest: freebusyrequest,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "freeBusy",
		context_:        googleapi.NoContext,
	}
}
func (c *FreebusyQueryCall) Context(ctx context.Context) *FreebusyQueryCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FreebusyQueryCall) Fields(s ...googleapi.Field) *FreebusyQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FreebusyQueryCall) Do() (*FreeBusyResponse, error) {
	var returnValue *FreeBusyResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.freebusyrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns free/busy information for a set of calendars.",
	//   "httpMethod": "POST",
	//   "id": "calendar.freebusy.query",
	//   "path": "freeBusy",
	//   "request": {
	//     "$ref": "FreeBusyRequest"
	//   },
	//   "response": {
	//     "$ref": "FreeBusyResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.settings.get":

type SettingsGetCall struct {
	s             *Service
	setting       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Get: Returns a single user setting.

func (r *SettingsService) Get(setting string) *SettingsGetCall {
	return &SettingsGetCall{
		s:             r.s,
		setting:       setting,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/settings/{setting}",
		context_:      googleapi.NoContext,
	}
}
func (c *SettingsGetCall) Context(ctx context.Context) *SettingsGetCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsGetCall) Fields(s ...googleapi.Field) *SettingsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SettingsGetCall) Do() (*Setting, error) {
	var returnValue *Setting
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"setting": c.setting,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns a single user setting.",
	//   "httpMethod": "GET",
	//   "id": "calendar.settings.get",
	//   "parameterOrder": [
	//     "setting"
	//   ],
	//   "parameters": {
	//     "setting": {
	//       "description": "The id of the user setting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/settings/{setting}",
	//   "response": {
	//     "$ref": "Setting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ]
	// }

}

// method id "calendar.settings.list":

type SettingsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// List: Returns all user settings for the authenticated user.

func (r *SettingsService) List() *SettingsListCall {
	return &SettingsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/settings",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. By default the value is 100
// entries. The page size can never be larger than 250 entries.
func (c *SettingsListCall) MaxResults(maxResults int64) *SettingsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *SettingsListCall) PageToken(pageToken string) *SettingsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then.
// If the
// syncToken expires, the server will respond with a 410 GONE response
// code and the client should clear its storage and perform a full
// synchronization without any syncToken.
// Learn more about incremental
// synchronization.
//  The default is to return all entries.
func (c *SettingsListCall) SyncToken(syncToken string) *SettingsListCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}
func (c *SettingsListCall) Context(ctx context.Context) *SettingsListCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsListCall) Fields(s ...googleapi.Field) *SettingsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SettingsListCall) Do() (*Settings, error) {
	var returnValue *Settings
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Returns all user settings for the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "calendar.settings.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/settings",
	//   "response": {
	//     "$ref": "Settings"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "calendar.settings.watch":

type SettingsWatchCall struct {
	s             *Service
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
}

// Watch: Watch for changes to Settings resources.

func (r *SettingsService) Watch(channel *Channel) *SettingsWatchCall {
	return &SettingsWatchCall{
		s:             r.s,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/me/settings/watch",
		context_:      googleapi.NoContext,
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of entries returned on one result page. By default the value is 100
// entries. The page size can never be larger than 250 entries.
func (c *SettingsWatchCall) MaxResults(maxResults int64) *SettingsWatchCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token specifying
// which result page to return.
func (c *SettingsWatchCall) PageToken(pageToken string) *SettingsWatchCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SyncToken sets the optional parameter "syncToken": Token obtained
// from the nextSyncToken field returned on the last page of results
// from the previous list request. It makes the result of this list
// request contain only entries that have changed since then.
// If the
// syncToken expires, the server will respond with a 410 GONE response
// code and the client should clear its storage and perform a full
// synchronization without any syncToken.
// Learn more about incremental
// synchronization.
//  The default is to return all entries.
func (c *SettingsWatchCall) SyncToken(syncToken string) *SettingsWatchCall {
	c.params_.Set("syncToken", fmt.Sprintf("%v", syncToken))
	return c
}
func (c *SettingsWatchCall) Context(ctx context.Context) *SettingsWatchCall {
	c.context_ = ctx
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsWatchCall) Fields(s ...googleapi.Field) *SettingsWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SettingsWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Watch for changes to Settings resources.",
	//   "httpMethod": "POST",
	//   "id": "calendar.settings.watch",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of entries returned on one result page. By default the value is 100 entries. The page size can never be larger than 250 entries. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "syncToken": {
	//       "description": "Token obtained from the nextSyncToken field returned on the last page of results from the previous list request. It makes the result of this list request contain only entries that have changed since then.\nIf the syncToken expires, the server will respond with a 410 GONE response code and the client should clear its storage and perform a full synchronization without any syncToken.\nLearn more about incremental synchronization.\nOptional. The default is to return all entries.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/me/settings/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/calendar",
	//     "https://www.googleapis.com/auth/calendar.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}
