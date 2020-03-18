package zdesk

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) AccountList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/account"
	return c.Request("GET", path, ro)
}

func (c *Client) AccountSettingsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/account/settings.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AccountSettingsUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/account/settings.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) AccountUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/account"
	return c.Request("PUT", path, ro)
}

func (c *Client) ActivitiesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/activities.json"
	return c.Request("GET", path, ro)
}

type ActivityShowInput struct {
	ActivityID string
}

func (c *Client) ActivityShow(i *ActivityShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ActivityID == "" {
		return nil, errors.New("Missing required field 'ActivityID'")
	}

	api_path := "/api/v2/activities/%s.json"
	path := fmt.Sprintf(api_path, i.ActivityID)
	return c.Request("GET", path, ro)
}

func (c *Client) AgentCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/agents"
	return c.Request("POST", path, ro)
}

type AgentDeleteInput struct {
	AgentID string
}

func (c *Client) AgentDelete(i *AgentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	api_path := "/api/v2/agents/%s"
	path := fmt.Sprintf(api_path, i.AgentID)
	return c.Request("DELETE", path, ro)
}

type AgentShowInput struct {
	AgentID string
}

func (c *Client) AgentShow(i *AgentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	api_path := "/api/v2/agents/%s"
	path := fmt.Sprintf(api_path, i.AgentID)
	return c.Request("GET", path, ro)
}

type AgentUpdateInput struct {
	AgentID string
}

func (c *Client) AgentUpdate(i *AgentUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	api_path := "/api/v2/agents/%s"
	path := fmt.Sprintf(api_path, i.AgentID)
	return c.Request("PUT", path, ro)
}

type AgentsEmailShowInput struct {
	EmailID string
}

func (c *Client) AgentsEmailShow(i *AgentsEmailShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.EmailID == "" {
		return nil, errors.New("Missing required field 'EmailID'")
	}

	api_path := "/api/v2/agents/email/%s"
	path := fmt.Sprintf(api_path, i.EmailID)
	return c.Request("GET", path, ro)
}

func (c *Client) AgentsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/agents"
	return c.Request("GET", path, ro)
}

func (c *Client) AgentsMe(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/agents/me"
	return c.Request("GET", path, ro)
}

func (c *Client) AgentsMeUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/agents/me"
	return c.Request("PUT", path, ro)
}

func (c *Client) AnyChannelPushCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/any_channel/push"
	return c.Request("POST", path, ro)
}

func (c *Client) AppCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps.json"
	return c.Request("POST", path, ro)
}

type AppDeleteInput struct {
	ID string
}

func (c *Client) AppDelete(i *AppDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type AppPublicKeyInput struct {
	ID string
}

func (c *Client) AppPublicKey(i *AppPublicKeyInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s/public_key.pem"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type AppShowInput struct {
	ID string
}

func (c *Client) AppShow(i *AppShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type AppUpdateInput struct {
	ID string
}

func (c *Client) AppUpdate(i *AppUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) AppsInstallationCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/installations.json"
	return c.Request("POST", path, ro)
}

type AppsInstallationDeleteInput struct {
	ID string
}

func (c *Client) AppsInstallationDelete(i *AppsInstallationDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type AppsInstallationRequirementsInput struct {
	ID string
}

func (c *Client) AppsInstallationRequirements(i *AppsInstallationRequirementsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s/requirements.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type AppsInstallationShowInput struct {
	ID string
}

func (c *Client) AppsInstallationShow(i *AppsInstallationShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type AppsInstallationUpdateInput struct {
	ID string
}

func (c *Client) AppsInstallationUpdate(i *AppsInstallationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type AppsInstallationsJobStatusShowInput struct {
	ID string
}

func (c *Client) AppsInstallationsJobStatusShow(i *AppsInstallationsJobStatusShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/job_statuses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) AppsInstallationsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/installations.json"
	return c.Request("GET", path, ro)
}

type AppsJobStatusShowInput struct {
	ID string
}

func (c *Client) AppsJobStatusShow(i *AppsJobStatusShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/job_statuses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) AppsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AppsLocationInstallationsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/location_installations.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AppsLocationInstallationsReorder(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/location_installations/reorder.json"
	return c.Request("POST", path, ro)
}

type AppsLocationShowInput struct {
	ID string
}

func (c *Client) AppsLocationShow(i *AppsLocationShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/locations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) AppsLocationsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/locations.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AppsNotifyCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/notify.json"
	return c.Request("POST", path, ro)
}

func (c *Client) AppsOwnedList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/owned.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AppsUploadCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/apps/uploads.json"
	return c.Request("POST", path, ro)
}

type AttachmentDeleteInput struct {
	ID string
}

func (c *Client) AttachmentDelete(i *AttachmentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type AttachmentShowInput struct {
	ID string
}

func (c *Client) AttachmentShow(i *AttachmentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type AuditLogShowInput struct {
	ID string
}

func (c *Client) AuditLogShow(i *AuditLogShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/audit_logs/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) AuditLogsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/audit_logs.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AutocompleteTags(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/autocomplete/tags.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AutomationCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/automations.json"
	return c.Request("POST", path, ro)
}

type AutomationDeleteInput struct {
	ID string
}

func (c *Client) AutomationDelete(i *AutomationDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/automations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type AutomationShowInput struct {
	ID string
}

func (c *Client) AutomationShow(i *AutomationShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/automations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type AutomationUpdateInput struct {
	ID string
}

func (c *Client) AutomationUpdate(i *AutomationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/automations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) AutomationsActiveList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/automations/active.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AutomationsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/automations/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) AutomationsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/automations.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AutomationsSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/automations/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) AutomationsUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/automations/update_many.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) BanCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/bans"
	return c.Request("POST", path, ro)
}

type BanDeleteInput struct {
	BanID string
}

func (c *Client) BanDelete(i *BanDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.BanID == "" {
		return nil, errors.New("Missing required field 'BanID'")
	}

	api_path := "/api/v2/ban/%s"
	path := fmt.Sprintf(api_path, i.BanID)
	return c.Request("DELETE", path, ro)
}

type BanShowInput struct {
	BanID string
}

func (c *Client) BanShow(i *BanShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.BanID == "" {
		return nil, errors.New("Missing required field 'BanID'")
	}

	api_path := "/api/v2/bans/%s"
	path := fmt.Sprintf(api_path, i.BanID)
	return c.Request("GET", path, ro)
}

func (c *Client) BansIpList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/bans/ip"
	return c.Request("GET", path, ro)
}

func (c *Client) BansList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/bans"
	return c.Request("GET", path, ro)
}

func (c *Client) BookmarkCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/bookmarks.json"
	return c.Request("POST", path, ro)
}

type BookmarkDeleteInput struct {
	ID string
}

func (c *Client) BookmarkDelete(i *BookmarkDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/bookmarks/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

func (c *Client) BookmarksList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/bookmarks.json"
	return c.Request("GET", path, ro)
}

type BrandCheckHostMappingInput struct {
	ID string
}

func (c *Client) BrandCheckHostMapping(i *BrandCheckHostMappingInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s/check_host_mapping.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) BrandCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/brands.json"
	return c.Request("POST", path, ro)
}

type BrandDeleteInput struct {
	ID string
}

func (c *Client) BrandDelete(i *BrandDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type BrandShowInput struct {
	ID string
}

func (c *Client) BrandShow(i *BrandShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type BrandUpdateInput struct {
	ID string
}

func (c *Client) BrandUpdate(i *BrandUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) BrandsCheckHostMappingList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/brands/check_host_mapping.json"
	return c.Request("GET", path, ro)
}

func (c *Client) BrandsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/brands.json"
	return c.Request("GET", path, ro)
}

func (c *Client) BusinessHoursScheduleCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/business_hours/schedules.json"
	return c.Request("POST", path, ro)
}

type BusinessHoursScheduleDeleteInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleDelete(i *BusinessHoursScheduleDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type BusinessHoursScheduleHolidayCreateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleHolidayCreate(i *BusinessHoursScheduleHolidayCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type BusinessHoursScheduleHolidayDeleteInput struct {
	ScheduleID string
	ID string
}

func (c *Client) BusinessHoursScheduleHolidayDelete(i *BusinessHoursScheduleHolidayDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays/%s.json"
	path := fmt.Sprintf(api_path, i.ScheduleID, i.ID)
	return c.Request("DELETE", path, ro)
}

type BusinessHoursScheduleHolidayShowInput struct {
	ScheduleID string
	ID string
}

func (c *Client) BusinessHoursScheduleHolidayShow(i *BusinessHoursScheduleHolidayShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays/%s.json"
	path := fmt.Sprintf(api_path, i.ScheduleID, i.ID)
	return c.Request("GET", path, ro)
}

type BusinessHoursScheduleHolidayUpdateInput struct {
	ScheduleID string
	ID string
}

func (c *Client) BusinessHoursScheduleHolidayUpdate(i *BusinessHoursScheduleHolidayUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays/%s.json"
	path := fmt.Sprintf(api_path, i.ScheduleID, i.ID)
	return c.Request("PUT", path, ro)
}

type BusinessHoursScheduleHolidaysInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleHolidays(i *BusinessHoursScheduleHolidaysInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type BusinessHoursScheduleShowInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleShow(i *BusinessHoursScheduleShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type BusinessHoursScheduleUpdateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleUpdate(i *BusinessHoursScheduleUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type BusinessHoursScheduleWorkweekUpdateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleWorkweekUpdate(i *BusinessHoursScheduleWorkweekUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/workweek.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) BusinessHoursSchedulesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/business_hours/schedules.json"
	return c.Request("GET", path, ro)
}

type ChannelsTwitterMonitoredTwitterHandleShowInput struct {
	ID string
}

func (c *Client) ChannelsTwitterMonitoredTwitterHandleShow(i *ChannelsTwitterMonitoredTwitterHandleShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/twitter/monitored_twitter_handles/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsTwitterMonitoredTwitterHandlesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/twitter/monitored_twitter_handles.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsTwitterTicketCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/twitter/tickets.json"
	return c.Request("POST", path, ro)
}

type ChannelsTwitterTicketStatusesInput struct {
	ID string
}

func (c *Client) ChannelsTwitterTicketStatuses(i *ChannelsTwitterTicketStatusesInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/twitter/tickets/%s/statuses.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ChannelsVoiceAgentTicketDisplayCreateInput struct {
	AgentID string
	TicketID string
}

func (c *Client) ChannelsVoiceAgentTicketDisplayCreate(i *ChannelsVoiceAgentTicketDisplayCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/channels/voice/agents/%s/tickets/%s/display.json"
	path := fmt.Sprintf(api_path, i.AgentID, i.TicketID)
	return c.Request("POST", path, ro)
}

type ChannelsVoiceAgentUserDisplayCreateInput struct {
	AgentID string
	UserID string
}

func (c *Client) ChannelsVoiceAgentUserDisplayCreate(i *ChannelsVoiceAgentUserDisplayCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/channels/voice/agents/%s/users/%s/display.json"
	path := fmt.Sprintf(api_path, i.AgentID, i.UserID)
	return c.Request("POST", path, ro)
}

type ChannelsVoiceAvailabilityShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceAvailabilityShow(i *ChannelsVoiceAvailabilityShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/availabilities/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ChannelsVoiceAvailabilityUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoiceAvailabilityUpdate(i *ChannelsVoiceAvailabilityUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/availabilities/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) ChannelsVoiceGreetingCategoriesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/greeting_categories.json"
	return c.Request("GET", path, ro)
}

type ChannelsVoiceGreetingCategoryShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingCategoryShow(i *ChannelsVoiceGreetingCategoryShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greeting_categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoiceGreetingCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/greetings.json"
	return c.Request("POST", path, ro)
}

type ChannelsVoiceGreetingDeleteInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingDelete(i *ChannelsVoiceGreetingDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type ChannelsVoiceGreetingRecordingInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingRecording(i *ChannelsVoiceGreetingRecordingInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s/recording.mp3"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ChannelsVoiceGreetingShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingShow(i *ChannelsVoiceGreetingShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ChannelsVoiceGreetingUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingUpdate(i *ChannelsVoiceGreetingUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) ChannelsVoiceGreetingsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/greetings.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoicePhoneNumberCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/phone_numbers.json"
	return c.Request("POST", path, ro)
}

type ChannelsVoicePhoneNumberShowInput struct {
	ID string
}

func (c *Client) ChannelsVoicePhoneNumberShow(i *ChannelsVoicePhoneNumberShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/phone_numbers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ChannelsVoicePhoneNumberUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoicePhoneNumberUpdate(i *ChannelsVoicePhoneNumberUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/phone_numbers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) ChannelsVoicePhoneNumbersDelete(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/phone_numbers.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) ChannelsVoicePhoneNumbersList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/phone_numbers.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoicePhoneNumbersSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/phone_numbers/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoiceStatsAgentsActivityList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/stats/agents_activity.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoiceStatsCurrentQueueActivityList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/stats/current_queue_activity.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoiceStatsHistoricalQueueActivityList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/stats/historical_queue_activity.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ChannelsVoiceTicketCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/channels/voice/tickets.json"
	return c.Request("POST", path, ro)
}

func (c *Client) ChatCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/chats"
	return c.Request("POST", path, ro)
}

type ChatDeleteInput struct {
	ChatID string
}

func (c *Client) ChatDelete(i *ChatDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	api_path := "/api/v2/chats/%s"
	path := fmt.Sprintf(api_path, i.ChatID)
	return c.Request("DELETE", path, ro)
}

type ChatShowInput struct {
	ChatID string
}

func (c *Client) ChatShow(i *ChatShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	api_path := "/api/v2/chats/%s"
	path := fmt.Sprintf(api_path, i.ChatID)
	return c.Request("GET", path, ro)
}

type ChatUpdateInput struct {
	ChatID string
}

func (c *Client) ChatUpdate(i *ChatUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	api_path := "/api/v2/chats/%s"
	path := fmt.Sprintf(api_path, i.ChatID)
	return c.Request("PUT", path, ro)
}

func (c *Client) ChatsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/chats"
	return c.Request("GET", path, ro)
}

func (c *Client) ChatsSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/chats/search"
	return c.Request("GET", path, ro)
}

type CommunityPostCommentCreateInput struct {
	PostID string
}

func (c *Client) CommunityPostCommentCreate(i *CommunityPostCommentCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/comments.json"
	path := fmt.Sprintf(api_path, i.PostID)
	return c.Request("POST", path, ro)
}

type CommunityPostCommentDeleteInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostCommentDelete(i *CommunityPostCommentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("DELETE", path, ro)
}

type CommunityPostCommentDownCreateInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostCommentDownCreate(i *CommunityPostCommentDownCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s/down.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("POST", path, ro)
}

type CommunityPostCommentShowInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostCommentShow(i *CommunityPostCommentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityPostCommentUpCreateInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostCommentUpCreate(i *CommunityPostCommentUpCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s/up.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("POST", path, ro)
}

type CommunityPostCommentUpdateInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostCommentUpdate(i *CommunityPostCommentUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("PUT", path, ro)
}

type CommunityPostCommentVotesInput struct {
	PostID string
	CommentID string
}

func (c *Client) CommunityPostCommentVotes(i *CommunityPostCommentVotesInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s/votes.json"
	path := fmt.Sprintf(api_path, i.PostID, i.CommentID)
	return c.Request("GET", path, ro)
}

type CommunityPostCommentsInput struct {
	PostID string
}

func (c *Client) CommunityPostComments(i *CommunityPostCommentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/comments.json"
	path := fmt.Sprintf(api_path, i.PostID)
	return c.Request("GET", path, ro)
}

func (c *Client) CommunityPostCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/community/posts.json"
	return c.Request("POST", path, ro)
}

type CommunityPostDeleteInput struct {
	ID string
}

func (c *Client) CommunityPostDelete(i *CommunityPostDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type CommunityPostDownCreateInput struct {
	ID string
}

func (c *Client) CommunityPostDownCreate(i *CommunityPostDownCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/down.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type CommunityPostShowInput struct {
	ID string
}

func (c *Client) CommunityPostShow(i *CommunityPostShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityPostSubscriptionCreateInput struct {
	PostID string
}

func (c *Client) CommunityPostSubscriptionCreate(i *CommunityPostSubscriptionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.PostID)
	return c.Request("POST", path, ro)
}

type CommunityPostSubscriptionDeleteInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostSubscriptionDelete(i *CommunityPostSubscriptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("DELETE", path, ro)
}

type CommunityPostSubscriptionShowInput struct {
	PostID string
	ID string
}

func (c *Client) CommunityPostSubscriptionShow(i *CommunityPostSubscriptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityPostSubscriptionsInput struct {
	PostID string
}

func (c *Client) CommunityPostSubscriptions(i *CommunityPostSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.PostID)
	return c.Request("GET", path, ro)
}

type CommunityPostUpCreateInput struct {
	ID string
}

func (c *Client) CommunityPostUpCreate(i *CommunityPostUpCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/up.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type CommunityPostUpdateInput struct {
	ID string
}

func (c *Client) CommunityPostUpdate(i *CommunityPostUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type CommunityPostVotesInput struct {
	PostID string
}

func (c *Client) CommunityPostVotes(i *CommunityPostVotesInput, ro *RequestOptions) (*http.Response, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/votes.json"
	path := fmt.Sprintf(api_path, i.PostID)
	return c.Request("GET", path, ro)
}

func (c *Client) CommunityPostsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/community/posts.json"
	return c.Request("GET", path, ro)
}

type CommunityTopicAccessPolicyInput struct {
	TopicID string
}

func (c *Client) CommunityTopicAccessPolicy(i *CommunityTopicAccessPolicyInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	return c.Request("GET", path, ro)
}

type CommunityTopicAccessPolicyUpdateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicAccessPolicyUpdate(i *CommunityTopicAccessPolicyUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	return c.Request("PUT", path, ro)
}

func (c *Client) CommunityTopicCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/community/topics.json"
	return c.Request("POST", path, ro)
}

type CommunityTopicDeleteInput struct {
	ID string
}

func (c *Client) CommunityTopicDelete(i *CommunityTopicDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type CommunityTopicPostsInput struct {
	ID string
}

func (c *Client) CommunityTopicPosts(i *CommunityTopicPostsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s/posts.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityTopicShowInput struct {
	ID string
}

func (c *Client) CommunityTopicShow(i *CommunityTopicShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityTopicSubscriptionCreateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptionCreate(i *CommunityTopicSubscriptionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	return c.Request("POST", path, ro)
}

type CommunityTopicSubscriptionDeleteInput struct {
	TopicID string
	ID string
}

func (c *Client) CommunityTopicSubscriptionDelete(i *CommunityTopicSubscriptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.TopicID, i.ID)
	return c.Request("DELETE", path, ro)
}

type CommunityTopicSubscriptionShowInput struct {
	TopicID string
	ID string
}

func (c *Client) CommunityTopicSubscriptionShow(i *CommunityTopicSubscriptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.TopicID, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityTopicSubscriptionsInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptions(i *CommunityTopicSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	return c.Request("GET", path, ro)
}

type CommunityTopicSubscriptionsUpdateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptionsUpdate(i *CommunityTopicSubscriptionsUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	return c.Request("PUT", path, ro)
}

type CommunityTopicUpdateInput struct {
	ID string
}

func (c *Client) CommunityTopicUpdate(i *CommunityTopicUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) CommunityTopicsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/community/topics.json"
	return c.Request("GET", path, ro)
}

type CommunityUserCommentsInput struct {
	ID string
}

func (c *Client) CommunityUserComments(i *CommunityUserCommentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/users/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type CommunityUserPostsInput struct {
	ID string
}

func (c *Client) CommunityUserPosts(i *CommunityUserPostsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/users/%s/posts.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) CustomRolesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/custom_roles.json"
	return c.Request("GET", path, ro)
}

type DeletedTicketRestoreUpdateInput struct {
	ID string
}

func (c *Client) DeletedTicketRestoreUpdate(i *DeletedTicketRestoreUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/deleted_tickets/%s/restore.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) DeletedTickets1Delete(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/deleted_tickets/1"
	return c.Request("DELETE", path, ro)
}

func (c *Client) DeletedTicketsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/deleted_tickets/destroy_many"
	return c.Request("DELETE", path, ro)
}

func (c *Client) DeletedTicketsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/deleted_tickets.json"
	return c.Request("GET", path, ro)
}

func (c *Client) DeletedTicketsRestoreManyUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/deleted_tickets/restore_many"
	return c.Request("PUT", path, ro)
}

func (c *Client) DepartmentCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/departments"
	return c.Request("POST", path, ro)
}

type DepartmentDeleteInput struct {
	DepartmentID string
}

func (c *Client) DepartmentDelete(i *DepartmentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	api_path := "/api/v2/departments/%s"
	path := fmt.Sprintf(api_path, i.DepartmentID)
	return c.Request("DELETE", path, ro)
}

type DepartmentShowInput struct {
	DepartmentID string
}

func (c *Client) DepartmentShow(i *DepartmentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	api_path := "/api/v2/departments/%s"
	path := fmt.Sprintf(api_path, i.DepartmentID)
	return c.Request("GET", path, ro)
}

type DepartmentUpdateInput struct {
	DepartmentID string
}

func (c *Client) DepartmentUpdate(i *DepartmentUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	api_path := "/api/v2/departments/%s"
	path := fmt.Sprintf(api_path, i.DepartmentID)
	return c.Request("PUT", path, ro)
}

func (c *Client) DepartmentsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/departments"
	return c.Request("GET", path, ro)
}

type DepartmentsNameDeleteInput struct {
	Name string
}

func (c *Client) DepartmentsNameDelete(i *DepartmentsNameDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/departments/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	return c.Request("DELETE", path, ro)
}

type DepartmentsNameShowInput struct {
	Name string
}

func (c *Client) DepartmentsNameShow(i *DepartmentsNameShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/departments/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	return c.Request("GET", path, ro)
}

type DepartmentsNameUpdateInput struct {
	Name string
}

func (c *Client) DepartmentsNameUpdate(i *DepartmentsNameUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/departments/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	return c.Request("PUT", path, ro)
}

func (c *Client) DynamicContentItemCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/dynamic_content/items.json"
	return c.Request("POST", path, ro)
}

type DynamicContentItemDeleteInput struct {
	ID string
}

func (c *Client) DynamicContentItemDelete(i *DynamicContentItemDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type DynamicContentItemShowInput struct {
	ID string
}

func (c *Client) DynamicContentItemShow(i *DynamicContentItemShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type DynamicContentItemUpdateInput struct {
	ID string
}

func (c *Client) DynamicContentItemUpdate(i *DynamicContentItemUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type DynamicContentItemVariantCreateInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantCreate(i *DynamicContentItemVariantCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type DynamicContentItemVariantDeleteInput struct {
	ItemID string
	ID string
}

func (c *Client) DynamicContentItemVariantDelete(i *DynamicContentItemVariantDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/%s.json"
	path := fmt.Sprintf(api_path, i.ItemID, i.ID)
	return c.Request("DELETE", path, ro)
}

type DynamicContentItemVariantShowInput struct {
	ItemID string
	ID string
}

func (c *Client) DynamicContentItemVariantShow(i *DynamicContentItemVariantShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/%s.json"
	path := fmt.Sprintf(api_path, i.ItemID, i.ID)
	return c.Request("GET", path, ro)
}

type DynamicContentItemVariantUpdateInput struct {
	ItemID string
	ID string
}

func (c *Client) DynamicContentItemVariantUpdate(i *DynamicContentItemVariantUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/%s.json"
	path := fmt.Sprintf(api_path, i.ItemID, i.ID)
	return c.Request("PUT", path, ro)
}

type DynamicContentItemVariantsInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariants(i *DynamicContentItemVariantsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type DynamicContentItemVariantsCreateManyInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantsCreateMany(i *DynamicContentItemVariantsCreateManyInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/create_many.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type DynamicContentItemVariantsUpdateManyInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantsUpdateMany(i *DynamicContentItemVariantsUpdateManyInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/update_many.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) DynamicContentItemsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/dynamic_content/items.json"
	return c.Request("GET", path, ro)
}

type EndUserIdentityCreateInput struct {
	UserID string
}

func (c *Client) EndUserIdentityCreate(i *EndUserIdentityCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/end_users/%s/identities.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("POST", path, ro)
}

type EndUserIdentityMakePrimaryInput struct {
	UserID string
	ID string
}

func (c *Client) EndUserIdentityMakePrimary(i *EndUserIdentityMakePrimaryInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/end_users/%s/identities/%s/make_primary"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("PUT", path, ro)
}

type EndUserShowInput struct {
	ID string
}

func (c *Client) EndUserShow(i *EndUserShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/end_users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type EndUserUpdateInput struct {
	ID string
}

func (c *Client) EndUserUpdate(i *EndUserUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/end_users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) GoalCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/goals"
	return c.Request("POST", path, ro)
}

type GoalDeleteInput struct {
	GoalID string
}

func (c *Client) GoalDelete(i *GoalDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	api_path := "/api/v2/goals/%s"
	path := fmt.Sprintf(api_path, i.GoalID)
	return c.Request("DELETE", path, ro)
}

type GoalShowInput struct {
	GoalID string
}

func (c *Client) GoalShow(i *GoalShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	api_path := "/api/v2/goals/%s"
	path := fmt.Sprintf(api_path, i.GoalID)
	return c.Request("GET", path, ro)
}

type GoalUpdateInput struct {
	GoalID string
}

func (c *Client) GoalUpdate(i *GoalUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	api_path := "/api/v2/goals/%s"
	path := fmt.Sprintf(api_path, i.GoalID)
	return c.Request("PUT", path, ro)
}

func (c *Client) GoalsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/goals"
	return c.Request("GET", path, ro)
}

func (c *Client) GroupCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/groups.json"
	return c.Request("POST", path, ro)
}

type GroupDeleteInput struct {
	ID string
}

func (c *Client) GroupDelete(i *GroupDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

func (c *Client) GroupMembershipCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/group_memberships.json"
	return c.Request("POST", path, ro)
}

type GroupMembershipDeleteInput struct {
	ID string
}

func (c *Client) GroupMembershipDelete(i *GroupMembershipDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type GroupMembershipShowInput struct {
	ID string
}

func (c *Client) GroupMembershipShow(i *GroupMembershipShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type GroupMembershipsInput struct {
	GroupID string
}

func (c *Client) GroupMemberships(i *GroupMembershipsInput, ro *RequestOptions) (*http.Response, error) {
	if i.GroupID == "" {
		return nil, errors.New("Missing required field 'GroupID'")
	}

	api_path := "/api/v2/groups/%s/memberships.json"
	path := fmt.Sprintf(api_path, i.GroupID)
	return c.Request("GET", path, ro)
}

type GroupMembershipsAssignableInput struct {
	GroupID string
}

func (c *Client) GroupMembershipsAssignable(i *GroupMembershipsAssignableInput, ro *RequestOptions) (*http.Response, error) {
	if i.GroupID == "" {
		return nil, errors.New("Missing required field 'GroupID'")
	}

	api_path := "/api/v2/groups/%s/memberships/assignable.json"
	path := fmt.Sprintf(api_path, i.GroupID)
	return c.Request("GET", path, ro)
}

func (c *Client) GroupMembershipsAssignableList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/group_memberships/assignable.json"
	return c.Request("GET", path, ro)
}

func (c *Client) GroupMembershipsCreateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/group_memberships/create_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) GroupMembershipsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/group_memberships/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) GroupMembershipsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/group_memberships.json"
	return c.Request("GET", path, ro)
}

type GroupShowInput struct {
	ID string
}

func (c *Client) GroupShow(i *GroupShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type GroupUpdateInput struct {
	ID string
}

func (c *Client) GroupUpdate(i *GroupUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type GroupUsersInput struct {
	ID string
}

func (c *Client) GroupUsers(i *GroupUsersInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s/users.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) GroupsAssignableList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/groups/assignable.json"
	return c.Request("GET", path, ro)
}

func (c *Client) GroupsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/groups.json"
	return c.Request("GET", path, ro)
}

type HelpCenterArticleAttachmentCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleAttachmentCreate(i *HelpCenterArticleAttachmentCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleAttachmentsInput struct {
	ArticleID string
Locale string
}

func (c *Client) HelpCenterArticleAttachments(i *HelpCenterArticleAttachmentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.ArticleID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/attachments.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleAttachmentsBlockInput struct {
	ArticleID string
Locale string
}

func (c *Client) HelpCenterArticleAttachmentsBlock(i *HelpCenterArticleAttachmentsBlockInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments/block.json"
	path := fmt.Sprintf(api_path, i.ArticleID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/attachments/block.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleAttachmentsInlineInput struct {
	ArticleID string
Locale string
}

func (c *Client) HelpCenterArticleAttachmentsInline(i *HelpCenterArticleAttachmentsInlineInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments/inline.json"
	path := fmt.Sprintf(api_path, i.ArticleID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/attachments/inline.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleBulkAttachmentCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleBulkAttachmentCreate(i *HelpCenterArticleBulkAttachmentCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/bulk_attachments.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleCommentCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleCommentCreate(i *HelpCenterArticleCommentCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleCommentDeleteInput struct {
	ArticleID string
	ID string
}

func (c *Client) HelpCenterArticleCommentDelete(i *HelpCenterArticleCommentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterArticleCommentDownCreateInput struct {
	ArticleID string
	ID string
}

func (c *Client) HelpCenterArticleCommentDownCreate(i *HelpCenterArticleCommentDownCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s/down.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleCommentShowInput struct {
	ArticleID string
	ID string
Locale string
}

func (c *Client) HelpCenterArticleCommentShow(i *HelpCenterArticleCommentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/comments/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleCommentUpCreateInput struct {
	ArticleID string
	ID string
}

func (c *Client) HelpCenterArticleCommentUpCreate(i *HelpCenterArticleCommentUpCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s/up.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleCommentUpdateInput struct {
	ArticleID string
	ID string
}

func (c *Client) HelpCenterArticleCommentUpdate(i *HelpCenterArticleCommentUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	return c.Request("PUT", path, ro)
}

type HelpCenterArticleCommentVotesInput struct {
	ArticleID string
	CommentID string
Locale string
}

func (c *Client) HelpCenterArticleCommentVotes(i *HelpCenterArticleCommentVotesInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s/votes.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.CommentID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/comments/{comment_id}/votes.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleCommentsInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterArticleComments(i *HelpCenterArticleCommentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{id}/comments.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterArticleDelete(i *HelpCenterArticleDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterArticleDownCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleDownCreate(i *HelpCenterArticleDownCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/down.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleLabelCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleLabelCreate(i *HelpCenterArticleLabelCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/labels.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleLabelDeleteInput struct {
	ArticleID string
	ID string
}

func (c *Client) HelpCenterArticleLabelDelete(i *HelpCenterArticleLabelDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/labels/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterArticleLabelsInput struct {
	ID string
}

func (c *Client) HelpCenterArticleLabels(i *HelpCenterArticleLabelsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/labels.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type HelpCenterArticleShowInput struct {
	Locale string
	ID string
}

func (c *Client) HelpCenterArticleShow(i *HelpCenterArticleShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/%s/articles/%s.json"
	path := fmt.Sprintf(api_path, i.Locale, i.ID)
	return c.Request("GET", path, ro)
}

type HelpCenterArticleSourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleSourceLocaleUpdate(i *HelpCenterArticleSourceLocaleUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/source_locale.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type HelpCenterArticleSubscriptionCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleSubscriptionCreate(i *HelpCenterArticleSubscriptionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleSubscriptionDeleteInput struct {
	ArticleID string
	ID string
}

func (c *Client) HelpCenterArticleSubscriptionDelete(i *HelpCenterArticleSubscriptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterArticleSubscriptionShowInput struct {
	ArticleID string
	ID string
Locale string
}

func (c *Client) HelpCenterArticleSubscriptionShow(i *HelpCenterArticleSubscriptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/subscriptions/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleSubscriptionsInput struct {
	ArticleID string
Locale string
}

func (c *Client) HelpCenterArticleSubscriptions(i *HelpCenterArticleSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.ArticleID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/subscriptions.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticleTranslationCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslationCreate(i *HelpCenterArticleTranslationCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleTranslationShowInput struct {
	ArticleID string
	Locale string
}

func (c *Client) HelpCenterArticleTranslationShow(i *HelpCenterArticleTranslationShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.Locale)
	return c.Request("GET", path, ro)
}

type HelpCenterArticleTranslationUpdateInput struct {
	ArticleID string
	Locale string
}

func (c *Client) HelpCenterArticleTranslationUpdate(i *HelpCenterArticleTranslationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.Locale)
	return c.Request("PUT", path, ro)
}

type HelpCenterArticleTranslationsInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslations(i *HelpCenterArticleTranslationsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	return c.Request("GET", path, ro)
}

type HelpCenterArticleTranslationsMissingInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslationsMissing(i *HelpCenterArticleTranslationsMissingInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations/missing.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	return c.Request("GET", path, ro)
}

type HelpCenterArticleUpCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleUpCreate(i *HelpCenterArticleUpCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/up.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type HelpCenterArticleUpdateInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterArticleUpdate(i *HelpCenterArticleUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("PUT", path, ro)
}

type HelpCenterArticleVotesInput struct {
	ArticleID string
Locale string
}

func (c *Client) HelpCenterArticleVotes(i *HelpCenterArticleVotesInput, ro *RequestOptions) (*http.Response, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/votes.json"
	path := fmt.Sprintf(api_path, i.ArticleID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/votes.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

func (c *Client) HelpCenterArticlesAttachmentCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/articles/attachments.json"
	return c.Request("POST", path, ro)
}

type HelpCenterArticlesAttachmentDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesAttachmentDelete(i *HelpCenterArticlesAttachmentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterArticlesAttachmentShowInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterArticlesAttachmentShow(i *HelpCenterArticlesAttachmentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/attachments/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticlesLabelShowInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesLabelShow(i *HelpCenterArticlesLabelShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/labels/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type HelpCenterArticlesLabelsListInput struct {
Locale string
}

func (c *Client) HelpCenterArticlesLabelsList(i *HelpCenterArticlesLabelsListInput, ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/articles/labels.json"

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/labels.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterArticlesListInput struct {
Locale string
}

func (c *Client) HelpCenterArticlesList(i *HelpCenterArticlesListInput, ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/articles.json"

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

func (c *Client) HelpCenterArticlesSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/articles/search.json"
	return c.Request("GET", path, ro)
}

type HelpCenterCategoriesListInput struct {
Locale string
}

func (c *Client) HelpCenterCategoriesList(i *HelpCenterCategoriesListInput, ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/categories.json"

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterCategoryArticlesInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterCategoryArticles(i *HelpCenterCategoryArticlesInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterCategoryCreateInput struct {
Locale string
}

func (c *Client) HelpCenterCategoryCreate(i *HelpCenterCategoryCreateInput, ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/categories.json"

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("POST", path, ro)
}

type HelpCenterCategoryDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterCategoryDelete(i *HelpCenterCategoryDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterCategorySectionCreateInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterCategorySectionCreate(i *HelpCenterCategorySectionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/sections.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}/sections.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("POST", path, ro)
}

type HelpCenterCategorySectionsInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterCategorySections(i *HelpCenterCategorySectionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/sections.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}/sections.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterCategoryShowInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterCategoryShow(i *HelpCenterCategoryShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterCategorySourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterCategorySourceLocaleUpdate(i *HelpCenterCategorySourceLocaleUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/source_locale.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type HelpCenterCategoryTranslationCreateInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslationCreate(i *HelpCenterCategoryTranslationCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations.json"
	path := fmt.Sprintf(api_path, i.CategoryID)
	return c.Request("POST", path, ro)
}

type HelpCenterCategoryTranslationUpdateInput struct {
	CategoryID string
	Locale string
}

func (c *Client) HelpCenterCategoryTranslationUpdate(i *HelpCenterCategoryTranslationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.CategoryID, i.Locale)
	return c.Request("PUT", path, ro)
}

type HelpCenterCategoryTranslationsInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslations(i *HelpCenterCategoryTranslationsInput, ro *RequestOptions) (*http.Response, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations.json"
	path := fmt.Sprintf(api_path, i.CategoryID)
	return c.Request("GET", path, ro)
}

type HelpCenterCategoryTranslationsMissingInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslationsMissing(i *HelpCenterCategoryTranslationsMissingInput, ro *RequestOptions) (*http.Response, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations/missing.json"
	path := fmt.Sprintf(api_path, i.CategoryID)
	return c.Request("GET", path, ro)
}

type HelpCenterCategoryUpdateInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterCategoryUpdate(i *HelpCenterCategoryUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("PUT", path, ro)
}

func (c *Client) HelpCenterIncrementalArticlesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/incremental/articles.json"
	return c.Request("GET", path, ro)
}

func (c *Client) HelpCenterLocalesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/locales.json"
	return c.Request("GET", path, ro)
}

type HelpCenterSectionAccessPolicyInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionAccessPolicy(i *HelpCenterSectionAccessPolicyInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	return c.Request("GET", path, ro)
}

type HelpCenterSectionAccessPolicyUpdateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionAccessPolicyUpdate(i *HelpCenterSectionAccessPolicyUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	return c.Request("PUT", path, ro)
}

type HelpCenterSectionArticleCreateInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterSectionArticleCreate(i *HelpCenterSectionArticleCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("POST", path, ro)
}

type HelpCenterSectionArticlesInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterSectionArticles(i *HelpCenterSectionArticlesInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterSectionDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterSectionDelete(i *HelpCenterSectionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterSectionShowInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterSectionShow(i *HelpCenterSectionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterSectionSourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterSectionSourceLocaleUpdate(i *HelpCenterSectionSourceLocaleUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/source_locale.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type HelpCenterSectionSubscriptionCreateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionSubscriptionCreate(i *HelpCenterSectionSubscriptionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	return c.Request("POST", path, ro)
}

type HelpCenterSectionSubscriptionDeleteInput struct {
	SectionID string
	ID string
}

func (c *Client) HelpCenterSectionSubscriptionDelete(i *HelpCenterSectionSubscriptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.SectionID, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterSectionSubscriptionShowInput struct {
	SectionID string
	ID string
Locale string
}

func (c *Client) HelpCenterSectionSubscriptionShow(i *HelpCenterSectionSubscriptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.SectionID, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{section_id}/subscriptions/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterSectionSubscriptionsInput struct {
	SectionID string
Locale string
}

func (c *Client) HelpCenterSectionSubscriptions(i *HelpCenterSectionSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.SectionID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{section_id}/subscriptions.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterSectionTranslationCreateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslationCreate(i *HelpCenterSectionTranslationCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	return c.Request("POST", path, ro)
}

type HelpCenterSectionTranslationUpdateInput struct {
	SectionID string
	Locale string
}

func (c *Client) HelpCenterSectionTranslationUpdate(i *HelpCenterSectionTranslationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.SectionID, i.Locale)
	return c.Request("PUT", path, ro)
}

type HelpCenterSectionTranslationsInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslations(i *HelpCenterSectionTranslationsInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	return c.Request("GET", path, ro)
}

type HelpCenterSectionTranslationsMissingInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslationsMissing(i *HelpCenterSectionTranslationsMissingInput, ro *RequestOptions) (*http.Response, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations/missing.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	return c.Request("GET", path, ro)
}

type HelpCenterSectionUpdateInput struct {
	ID string
Locale string
}

func (c *Client) HelpCenterSectionUpdate(i *HelpCenterSectionUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("PUT", path, ro)
}

type HelpCenterSectionsListInput struct {
Locale string
}

func (c *Client) HelpCenterSectionsList(i *HelpCenterSectionsListInput, ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/sections.json"

	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	return c.Request("GET", path, ro)
}

type HelpCenterTranslationDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterTranslationDelete(i *HelpCenterTranslationDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/translations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterUserArticlesInput struct {
	ID string
}

func (c *Client) HelpCenterUserArticles(i *HelpCenterUserArticlesInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/users/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type HelpCenterUserCommentsInput struct {
	ID string
}

func (c *Client) HelpCenterUserComments(i *HelpCenterUserCommentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/users/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) HelpCenterUserSegmentCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/user_segments.json"
	return c.Request("POST", path, ro)
}

type HelpCenterUserSegmentDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentDelete(i *HelpCenterUserSegmentDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterUserSegmentSectionsInput struct {
	UserSegmentID string
}

func (c *Client) HelpCenterUserSegmentSections(i *HelpCenterUserSegmentSectionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserSegmentID == "" {
		return nil, errors.New("Missing required field 'UserSegmentID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s/sections.json"
	path := fmt.Sprintf(api_path, i.UserSegmentID)
	return c.Request("GET", path, ro)
}

type HelpCenterUserSegmentShowInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentShow(i *HelpCenterUserSegmentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type HelpCenterUserSegmentTopicsInput struct {
	UserSegmentID string
}

func (c *Client) HelpCenterUserSegmentTopics(i *HelpCenterUserSegmentTopicsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserSegmentID == "" {
		return nil, errors.New("Missing required field 'UserSegmentID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s/topics.json"
	path := fmt.Sprintf(api_path, i.UserSegmentID)
	return c.Request("GET", path, ro)
}

type HelpCenterUserSegmentUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentUpdate(i *HelpCenterUserSegmentUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) HelpCenterUserSegmentsApplicableList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/user_segments/applicable.json"
	return c.Request("GET", path, ro)
}

func (c *Client) HelpCenterUserSegmentsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/help_center/user_segments.json"
	return c.Request("GET", path, ro)
}

type HelpCenterUserSubscriptionsInput struct {
	UserID string
}

func (c *Client) HelpCenterUserSubscriptions(i *HelpCenterUserSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/help_center/users/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type HelpCenterUserVotesInput struct {
	UserID string
}

func (c *Client) HelpCenterUserVotes(i *HelpCenterUserVotesInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/help_center/users/%s/votes.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type HelpCenterVoteDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterVoteDelete(i *HelpCenterVoteDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/votes/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type HelpCenterVoteShowInput struct {
	ID string
}

func (c *Client) HelpCenterVoteShow(i *HelpCenterVoteShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/votes/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) ImportsTicket(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/imports/tickets.json"
	return c.Request("POST", path, ro)
}

func (c *Client) ImportsTicketsCreateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/imports/tickets/create_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) IncrementalOrganizationsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/incremental/organizations.json"
	return c.Request("GET", path, ro)
}

type IncrementalSampleInput struct {
	Item string
}

func (c *Client) IncrementalSample(i *IncrementalSampleInput, ro *RequestOptions) (*http.Response, error) {
	if i.Item == "" {
		return nil, errors.New("Missing required field 'Item'")
	}

	api_path := "/api/v2/incremental/%s/sample.json"
	path := fmt.Sprintf(api_path, i.Item)
	return c.Request("GET", path, ro)
}

func (c *Client) IncrementalTicketEventsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/incremental/ticket_events.json"
	return c.Request("GET", path, ro)
}

func (c *Client) IncrementalTicketMetricEventsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/incremental/ticket_metric_events.json"
	return c.Request("GET", path, ro)
}

func (c *Client) IncrementalTicketsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/incremental/tickets.json"
	return c.Request("GET", path, ro)
}

func (c *Client) IncrementalUsersList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/incremental/users.json"
	return c.Request("GET", path, ro)
}

type JobStatusShowInput struct {
	ID string
}

func (c *Client) JobStatusShow(i *JobStatusShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/job_statuses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) JobStatusesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/job_statuses.json"
	return c.Request("GET", path, ro)
}

func (c *Client) JobStatusesShowMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/job_statuses/show_many.json"
	return c.Request("GET", path, ro)
}

type LocaleShowInput struct {
	ID string
}

func (c *Client) LocaleShow(i *LocaleShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/locales/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) LocalesAgentList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/locales/agent.json"
	return c.Request("GET", path, ro)
}

func (c *Client) LocalesCurrentList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/locales/current.json"
	return c.Request("GET", path, ro)
}

func (c *Client) LocalesDetectBestLocale(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/locales/detect_best_locale.json"
	return c.Request("GET", path, ro)
}

func (c *Client) LocalesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/locales.json"
	return c.Request("GET", path, ro)
}

func (c *Client) LocalesPublicList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/locales/public.json"
	return c.Request("GET", path, ro)
}

type MacroApplyInput struct {
	ID string
}

func (c *Client) MacroApply(i *MacroApplyInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s/apply.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type MacroAttachmentCreateInput struct {
	MacroID string
}

func (c *Client) MacroAttachmentCreate(i *MacroAttachmentCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.MacroID == "" {
		return nil, errors.New("Missing required field 'MacroID'")
	}

	api_path := "/api/v2/macros/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.MacroID)
	return c.Request("POST", path, ro)
}

type MacroAttachmentsInput struct {
	MacroID string
}

func (c *Client) MacroAttachments(i *MacroAttachmentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.MacroID == "" {
		return nil, errors.New("Missing required field 'MacroID'")
	}

	api_path := "/api/v2/macros/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.MacroID)
	return c.Request("GET", path, ro)
}

func (c *Client) MacroCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros.json"
	return c.Request("POST", path, ro)
}

type MacroDeleteInput struct {
	ID string
}

func (c *Client) MacroDelete(i *MacroDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type MacroShowInput struct {
	ID string
}

func (c *Client) MacroShow(i *MacroShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type MacroUpdateInput struct {
	ID string
}

func (c *Client) MacroUpdate(i *MacroUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) MacrosActionsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/actions.json"
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosActiveList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/active.json"
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosAttachmentCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/attachments.json"
	return c.Request("POST", path, ro)
}

type MacrosAttachmentShowInput struct {
	ID string
}

func (c *Client) MacrosAttachmentShow(i *MacrosAttachmentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosCategoriesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/categories.json"
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) MacrosList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros.json"
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosNewList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/new.json"
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) MacrosUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/macros/update_many.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) NpsIncrementalRecipientsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/nps/incremental/recipients.json"
	return c.Request("GET", path, ro)
}

func (c *Client) NpsIncrementalResponsesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/nps/incremental/responses.json"
	return c.Request("GET", path, ro)
}

type NpsSurveyCloseCreateInput struct {
	ID string
}

func (c *Client) NpsSurveyCloseCreate(i *NpsSurveyCloseCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/close"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type NpsSurveyInvitationCreateInput struct {
	ID string
}

func (c *Client) NpsSurveyInvitationCreate(i *NpsSurveyInvitationCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/invitations.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type NpsSurveyInvitationShowInput struct {
	SurveyID string
	ID string
}

func (c *Client) NpsSurveyInvitationShow(i *NpsSurveyInvitationShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/invitations/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	return c.Request("GET", path, ro)
}

type NpsSurveyInvitationsInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyInvitations(i *NpsSurveyInvitationsInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/invitations.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	return c.Request("GET", path, ro)
}

type NpsSurveyPreviewInput struct {
	ID string
}

func (c *Client) NpsSurveyPreview(i *NpsSurveyPreviewInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/preview"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type NpsSurveyRecipientCreateInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipientCreate(i *NpsSurveyRecipientCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	return c.Request("POST", path, ro)
}

type NpsSurveyRecipientShowInput struct {
	SurveyID string
	ID string
}

func (c *Client) NpsSurveyRecipientShow(i *NpsSurveyRecipientShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	return c.Request("GET", path, ro)
}

type NpsSurveyRecipientUpdateInput struct {
	SurveyID string
	ID string
}

func (c *Client) NpsSurveyRecipientUpdate(i *NpsSurveyRecipientUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	return c.Request("PUT", path, ro)
}

type NpsSurveyRecipientsInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipients(i *NpsSurveyRecipientsInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	return c.Request("GET", path, ro)
}

type NpsSurveyRecipientsSearchInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipientsSearch(i *NpsSurveyRecipientsSearchInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients/search.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	return c.Request("GET", path, ro)
}

type NpsSurveyResponseCreateInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyResponseCreate(i *NpsSurveyResponseCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	return c.Request("POST", path, ro)
}

type NpsSurveyResponseShowInput struct {
	SurveyID string
	ID string
}

func (c *Client) NpsSurveyResponseShow(i *NpsSurveyResponseShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	return c.Request("GET", path, ro)
}

type NpsSurveyResponseUpdateInput struct {
	SurveyID string
	ID string
}

func (c *Client) NpsSurveyResponseUpdate(i *NpsSurveyResponseUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	return c.Request("PUT", path, ro)
}

type NpsSurveyResponsesInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyResponses(i *NpsSurveyResponsesInput, ro *RequestOptions) (*http.Response, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	return c.Request("GET", path, ro)
}

type NpsSurveyShowInput struct {
	ID string
}

func (c *Client) NpsSurveyShow(i *NpsSurveyShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) NpsSurveys1Update(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/nps/surveys/1"
	return c.Request("PUT", path, ro)
}

func (c *Client) NpsSurveysList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/nps/surveys"
	return c.Request("GET", path, ro)
}

func (c *Client) OauthClientCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/clients"
	return c.Request("POST", path, ro)
}

type OauthClientDeleteInput struct {
	ID string
}

func (c *Client) OauthClientDelete(i *OauthClientDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type OauthClientGenerateSecretUpdateInput struct {
	ID string
}

func (c *Client) OauthClientGenerateSecretUpdate(i *OauthClientGenerateSecretUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s/generate_secret.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type OauthClientShowInput struct {
	ID string
}

func (c *Client) OauthClientShow(i *OauthClientShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type OauthClientUpdateInput struct {
	ID string
}

func (c *Client) OauthClientUpdate(i *OauthClientUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) OauthClientsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/clients"
	return c.Request("GET", path, ro)
}

func (c *Client) OauthGlobalClientsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/global_clients.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OauthTokenCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/tokens.json"
	return c.Request("POST", path, ro)
}

type OauthTokenDeleteInput struct {
	ID string
}

func (c *Client) OauthTokenDelete(i *OauthTokenDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/tokens/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type OauthTokenShowInput struct {
	ID string
}

func (c *Client) OauthTokenShow(i *OauthTokenShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/tokens/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) OauthTokensCurrentDelete(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/tokens/current.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) OauthTokensCurrentList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/tokens/current.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OauthTokensList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/oauth/tokens"
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations.json"
	return c.Request("POST", path, ro)
}

type OrganizationDeleteInput struct {
	ID string
}

func (c *Client) OrganizationDelete(i *OrganizationDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

func (c *Client) OrganizationFieldCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_fields.json"
	return c.Request("POST", path, ro)
}

type OrganizationFieldDeleteInput struct {
	ID string
}

func (c *Client) OrganizationFieldDelete(i *OrganizationFieldDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type OrganizationFieldShowInput struct {
	ID string
}

func (c *Client) OrganizationFieldShow(i *OrganizationFieldShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type OrganizationFieldUpdateInput struct {
	ID string
}

func (c *Client) OrganizationFieldUpdate(i *OrganizationFieldUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) OrganizationFieldsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_fields.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationFieldsReorder(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_fields/reorder.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) OrganizationMembershipCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_memberships.json"
	return c.Request("POST", path, ro)
}

type OrganizationMembershipDeleteInput struct {
	ID string
}

func (c *Client) OrganizationMembershipDelete(i *OrganizationMembershipDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type OrganizationMembershipShowInput struct {
	ID string
}

func (c *Client) OrganizationMembershipShow(i *OrganizationMembershipShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationMembershipsCreateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_memberships/create_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) OrganizationMembershipsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_memberships/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) OrganizationMembershipsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_memberships.json"
	return c.Request("GET", path, ro)
}

type OrganizationOrganizationMembershipsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationOrganizationMemberships(i *OrganizationOrganizationMembershipsInput, ro *RequestOptions) (*http.Response, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	api_path := "/api/v2/organizations/%s/organization_memberships.json"
	path := fmt.Sprintf(api_path, i.OrganizationID)
	return c.Request("GET", path, ro)
}

type OrganizationRelatedInput struct {
	ID string
}

func (c *Client) OrganizationRelated(i *OrganizationRelatedInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/related.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type OrganizationRequestsInput struct {
	ID string
}

func (c *Client) OrganizationRequests(i *OrganizationRequestsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/requests.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type OrganizationShowInput struct {
	ID string
}

func (c *Client) OrganizationShow(i *OrganizationShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationSubscriptionCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_subscriptions.json"
	return c.Request("POST", path, ro)
}

type OrganizationSubscriptionDeleteInput struct {
	ID string
}

func (c *Client) OrganizationSubscriptionDelete(i *OrganizationSubscriptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type OrganizationSubscriptionShowInput struct {
	ID string
}

func (c *Client) OrganizationSubscriptionShow(i *OrganizationSubscriptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type OrganizationSubscriptionsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationSubscriptions(i *OrganizationSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	api_path := "/api/v2/organizations/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.OrganizationID)
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationSubscriptionsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organization_subscriptions.json"
	return c.Request("GET", path, ro)
}

type OrganizationTagCreateInput struct {
	ID string
}

func (c *Client) OrganizationTagCreate(i *OrganizationTagCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type OrganizationTagsInput struct {
	ID string
}

func (c *Client) OrganizationTags(i *OrganizationTagsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type OrganizationTagsDeleteInput struct {
	ID string
}

func (c *Client) OrganizationTagsDelete(i *OrganizationTagsDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type OrganizationTagsUpdateInput struct {
	ID string
}

func (c *Client) OrganizationTagsUpdate(i *OrganizationTagsUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type OrganizationTicketsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationTickets(i *OrganizationTicketsInput, ro *RequestOptions) (*http.Response, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	api_path := "/api/v2/organizations/%s/tickets.json"
	path := fmt.Sprintf(api_path, i.OrganizationID)
	return c.Request("GET", path, ro)
}

type OrganizationUpdateInput struct {
	ID string
}

func (c *Client) OrganizationUpdate(i *OrganizationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type OrganizationUsersInput struct {
	ID string
}

func (c *Client) OrganizationUsers(i *OrganizationUsersInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/users.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationsAutocomplete(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/autocomplete.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationsCreateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/create_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) OrganizationsCreateOrUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/create_or_update.json"
	return c.Request("POST", path, ro)
}

func (c *Client) OrganizationsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) OrganizationsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationsSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationsShowMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/show_many.json"
	return c.Request("GET", path, ro)
}

func (c *Client) OrganizationsUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/organizations/update_many.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) ProblemsAutocomplete(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/problems/autocomplete.json"
	return c.Request("POST", path, ro)
}

func (c *Client) ProblemsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/problems.json"
	return c.Request("GET", path, ro)
}

func (c *Client) PushNotificationDevicesDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/push_notification_devices/destroy_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) RecipientAddressCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/recipient_addresses.json"
	return c.Request("POST", path, ro)
}

type RecipientAddressDeleteInput struct {
	ID string
}

func (c *Client) RecipientAddressDelete(i *RecipientAddressDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type RecipientAddressShowInput struct {
	ID string
}

func (c *Client) RecipientAddressShow(i *RecipientAddressShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type RecipientAddressUpdateInput struct {
	ID string
}

func (c *Client) RecipientAddressUpdate(i *RecipientAddressUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type RecipientAddressVerifyInput struct {
	ID string
}

func (c *Client) RecipientAddressVerify(i *RecipientAddressVerifyInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s/verify.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) RecipientAddressesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/recipient_addresses.json"
	return c.Request("GET", path, ro)
}

type RequestCommentShowInput struct {
	RequestID string
	ID string
}

func (c *Client) RequestCommentShow(i *RequestCommentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.RequestID == "" {
		return nil, errors.New("Missing required field 'RequestID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.RequestID, i.ID)
	return c.Request("GET", path, ro)
}

type RequestCommentsInput struct {
	ID string
}

func (c *Client) RequestComments(i *RequestCommentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) RequestCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/requests.json"
	return c.Request("POST", path, ro)
}

type RequestShowInput struct {
	ID string
}

func (c *Client) RequestShow(i *RequestShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type RequestUpdateInput struct {
	ID string
}

func (c *Client) RequestUpdate(i *RequestUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) RequestsCcdList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/requests/ccd.json"
	return c.Request("GET", path, ro)
}

func (c *Client) RequestsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/requests.json"
	return c.Request("GET", path, ro)
}

func (c *Client) RequestsOpenList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/requests/open.json"
	return c.Request("GET", path, ro)
}

func (c *Client) RequestsSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/requests/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) RequestsSolvedList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/requests/solved.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ResourceCollectionCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/resource_collections.json"
	return c.Request("POST", path, ro)
}

type ResourceCollectionDeleteInput struct {
	ID string
}

func (c *Client) ResourceCollectionDelete(i *ResourceCollectionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/resource_collections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type ResourceCollectionShowInput struct {
	ID string
}

func (c *Client) ResourceCollectionShow(i *ResourceCollectionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/resource_collections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) ResourceCollectionsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/resource_collections.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ResourceCollectionsUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/resource_collections.json"
	return c.Request("PUT", path, ro)
}

type SatisfactionRatingShowInput struct {
	ID string
}

func (c *Client) SatisfactionRatingShow(i *SatisfactionRatingShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/satisfaction_ratings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) SatisfactionRatingsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/satisfaction_ratings.json"
	return c.Request("GET", path, ro)
}

type SatisfactionReasonShowInput struct {
	ID string
}

func (c *Client) SatisfactionReasonShow(i *SatisfactionReasonShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/satisfaction_reasons/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) SatisfactionReasonsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/satisfaction_reasons.json"
	return c.Request("GET", path, ro)
}

func (c *Client) Search(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) SessionsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/sessions.json"
	return c.Request("GET", path, ro)
}

func (c *Client) SharingAgreementCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/sharing_agreements.json"
	return c.Request("POST", path, ro)
}

type SharingAgreementDeleteInput struct {
	ID string
}

func (c *Client) SharingAgreementDelete(i *SharingAgreementDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/sharing_agreements/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type SharingAgreementShowInput struct {
	ID string
}

func (c *Client) SharingAgreementShow(i *SharingAgreementShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/sharing_agreements/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type SharingAgreementUpdateInput struct {
	ID string
}

func (c *Client) SharingAgreementUpdate(i *SharingAgreementUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/sharing_agreements/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) SharingAgreementsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/sharing_agreements.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ShortcutCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/shortcuts"
	return c.Request("POST", path, ro)
}

type ShortcutDeleteInput struct {
	ShortcutName string
}

func (c *Client) ShortcutDelete(i *ShortcutDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	api_path := "/api/v2/shortcuts/%s"
	path := fmt.Sprintf(api_path, i.ShortcutName)
	return c.Request("DELETE", path, ro)
}

type ShortcutShowInput struct {
	ShortcutName string
}

func (c *Client) ShortcutShow(i *ShortcutShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	api_path := "/api/v2/shortcuts/%s"
	path := fmt.Sprintf(api_path, i.ShortcutName)
	return c.Request("GET", path, ro)
}

type ShortcutUpdateInput struct {
	ShortcutName string
}

func (c *Client) ShortcutUpdate(i *ShortcutUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	api_path := "/api/v2/shortcuts/%s"
	path := fmt.Sprintf(api_path, i.ShortcutName)
	return c.Request("PUT", path, ro)
}

func (c *Client) ShortcutsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/shortcuts"
	return c.Request("GET", path, ro)
}

func (c *Client) SkillCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/skills"
	return c.Request("POST", path, ro)
}

type SkillDeleteInput struct {
	SkillID string
}

func (c *Client) SkillDelete(i *SkillDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.SkillID == "" {
		return nil, errors.New("Missing required field 'SkillID'")
	}

	api_path := "/api/v2/skills/%s"
	path := fmt.Sprintf(api_path, i.SkillID)
	return c.Request("DELETE", path, ro)
}

type SkillShowInput struct {
	SkillID string
}

func (c *Client) SkillShow(i *SkillShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.SkillID == "" {
		return nil, errors.New("Missing required field 'SkillID'")
	}

	api_path := "/api/v2/skills/%s"
	path := fmt.Sprintf(api_path, i.SkillID)
	return c.Request("GET", path, ro)
}

type SkillUpdateInput struct {
	SkillID string
}

func (c *Client) SkillUpdate(i *SkillUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.SkillID == "" {
		return nil, errors.New("Missing required field 'SkillID'")
	}

	api_path := "/api/v2/skills/%s"
	path := fmt.Sprintf(api_path, i.SkillID)
	return c.Request("PUT", path, ro)
}

func (c *Client) SkillsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/skills"
	return c.Request("GET", path, ro)
}

type SkillsNameDeleteInput struct {
	Name string
}

func (c *Client) SkillsNameDelete(i *SkillsNameDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/skills/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	return c.Request("DELETE", path, ro)
}

type SkillsNameShowInput struct {
	Name string
}

func (c *Client) SkillsNameShow(i *SkillsNameShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/skills/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	return c.Request("GET", path, ro)
}

type SkillsNameUpdateInput struct {
	Name string
}

func (c *Client) SkillsNameUpdate(i *SkillsNameUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/skills/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	return c.Request("PUT", path, ro)
}

func (c *Client) SkipCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/skips.json"
	return c.Request("POST", path, ro)
}

func (c *Client) SkipsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/skips.json"
	return c.Request("GET", path, ro)
}

func (c *Client) SlasPoliciesDefinitionsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/slas/policies/definitions.json"
	return c.Request("GET", path, ro)
}

func (c *Client) SlasPoliciesList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/slas/policies"
	return c.Request("GET", path, ro)
}

func (c *Client) SlasPoliciesReorder(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/slas/policies/reorder.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) SlasPolicyCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/slas/policies"
	return c.Request("POST", path, ro)
}

type SlasPolicyDeleteInput struct {
	ID string
}

func (c *Client) SlasPolicyDelete(i *SlasPolicyDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/slas/policies/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type SlasPolicyShowInput struct {
	ID string
}

func (c *Client) SlasPolicyShow(i *SlasPolicyShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/slas/policies/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type SlasPolicyUpdateInput struct {
	ID string
}

func (c *Client) SlasPolicyUpdate(i *SlasPolicyUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/slas/policies/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type StreamAgentShowInput struct {
	MetricKey string
}

func (c *Client) StreamAgentShow(i *StreamAgentShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.MetricKey == "" {
		return nil, errors.New("Missing required field 'MetricKey'")
	}

	api_path := "/stream/agents/%s"
	path := fmt.Sprintf(api_path, i.MetricKey)
	return c.Request("GET", path, ro)
}

func (c *Client) StreamAgentsAgentsOnlineList(ro *RequestOptions) (*http.Response, error) {
	path := "/stream/agents/agents_online"
	return c.Request("GET", path, ro)
}

func (c *Client) StreamAgentsList(ro *RequestOptions) (*http.Response, error) {
	path := "/stream/agents"
	return c.Request("GET", path, ro)
}

type StreamChatShowInput struct {
	MetricKey string
}

func (c *Client) StreamChatShow(i *StreamChatShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.MetricKey == "" {
		return nil, errors.New("Missing required field 'MetricKey'")
	}

	api_path := "/stream/chats/%s"
	path := fmt.Sprintf(api_path, i.MetricKey)
	return c.Request("GET", path, ro)
}

func (c *Client) StreamChatsList(ro *RequestOptions) (*http.Response, error) {
	path := "/stream/chats"
	return c.Request("GET", path, ro)
}

func (c *Client) StreamChatsMissedChatsList(ro *RequestOptions) (*http.Response, error) {
	path := "/stream/chats/missed_chats"
	return c.Request("GET", path, ro)
}

type SuspendedTicketDeleteInput struct {
	ID string
}

func (c *Client) SuspendedTicketDelete(i *SuspendedTicketDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/suspended_tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type SuspendedTicketRecoverInput struct {
	ID string
}

func (c *Client) SuspendedTicketRecover(i *SuspendedTicketRecoverInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/suspended_tickets/%s/recover.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type SuspendedTicketShowInput struct {
	ID string
}

func (c *Client) SuspendedTicketShow(i *SuspendedTicketShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/suspended_tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) SuspendedTicketsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/suspended_tickets/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) SuspendedTicketsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/suspended_tickets.json"
	return c.Request("GET", path, ro)
}

func (c *Client) SuspendedTicketsRecoverMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/suspended_tickets/recover_many.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) TagsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tags.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TargetCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/targets.json"
	return c.Request("POST", path, ro)
}

type TargetDeleteInput struct {
	ID string
}

func (c *Client) TargetDelete(i *TargetDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/targets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type TargetFailureShowInput struct {
	ID string
}

func (c *Client) TargetFailureShow(i *TargetFailureShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/target_failures/%s"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) TargetFailuresList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/target_failures"
	return c.Request("GET", path, ro)
}

type TargetShowInput struct {
	ID string
}

func (c *Client) TargetShow(i *TargetShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/targets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TargetUpdateInput struct {
	ID string
}

func (c *Client) TargetUpdate(i *TargetUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/targets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) TargetsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/targets.json"
	return c.Request("GET", path, ro)
}

type TicketAuditMakePrivateInput struct {
	TicketID string
	ID string
}

func (c *Client) TicketAuditMakePrivate(i *TicketAuditMakePrivateInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/audits/%s/make_private.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	return c.Request("PUT", path, ro)
}

type TicketAuditShowInput struct {
	TicketID string
	ID string
}

func (c *Client) TicketAuditShow(i *TicketAuditShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/audits/%s.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	return c.Request("GET", path, ro)
}

type TicketAuditsInput struct {
	TicketID string
}

func (c *Client) TicketAudits(i *TicketAuditsInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/audits.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	return c.Request("GET", path, ro)
}

func (c *Client) TicketAuditsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_audits.json"
	return c.Request("GET", path, ro)
}

type TicketCollaboratorsInput struct {
	ID string
}

func (c *Client) TicketCollaborators(i *TicketCollaboratorsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/collaborators.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketCommentAttachmentRedactInput struct {
	TicketID string
	CommentID string
	AttachmentID string
}

func (c *Client) TicketCommentAttachmentRedact(i *TicketCommentAttachmentRedactInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	if i.AttachmentID == "" {
		return nil, errors.New("Missing required field 'AttachmentID'")
	}

	api_path := "/api/v2/tickets/%s/comments/%s/attachments/%s/redact.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.CommentID, i.AttachmentID)
	return c.Request("PUT", path, ro)
}

type TicketCommentMakePrivateInput struct {
	TicketID string
	ID string
}

func (c *Client) TicketCommentMakePrivate(i *TicketCommentMakePrivateInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/comments/%s/make_private.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	return c.Request("PUT", path, ro)
}

type TicketCommentRedactInput struct {
	TicketID string
	ID string
}

func (c *Client) TicketCommentRedact(i *TicketCommentRedactInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/comments/%s/redact.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	return c.Request("PUT", path, ro)
}

type TicketCommentsInput struct {
	TicketID string
}

func (c *Client) TicketComments(i *TicketCommentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/comments.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	return c.Request("GET", path, ro)
}

func (c *Client) TicketCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets.json"
	return c.Request("POST", path, ro)
}

type TicketDeleteInput struct {
	ID string
}

func (c *Client) TicketDelete(i *TicketDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

func (c *Client) TicketFieldCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_fields.json"
	return c.Request("POST", path, ro)
}

type TicketFieldDeleteInput struct {
	ID string
}

func (c *Client) TicketFieldDelete(i *TicketFieldDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type TicketFieldOptionCreateInput struct {
	FieldID string
	ID string
}

func (c *Client) TicketFieldOptionCreate(i *TicketFieldOptionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	return c.Request("POST", path, ro)
}

type TicketFieldOptionDeleteInput struct {
	FieldID string
	ID string
}

func (c *Client) TicketFieldOptionDelete(i *TicketFieldOptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	return c.Request("DELETE", path, ro)
}

type TicketFieldOptionShowInput struct {
	FieldID string
	ID string
}

func (c *Client) TicketFieldOptionShow(i *TicketFieldOptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	return c.Request("GET", path, ro)
}

type TicketFieldOptionsInput struct {
	FieldID string
}

func (c *Client) TicketFieldOptions(i *TicketFieldOptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options.json"
	path := fmt.Sprintf(api_path, i.FieldID)
	return c.Request("GET", path, ro)
}

type TicketFieldShowInput struct {
	ID string
}

func (c *Client) TicketFieldShow(i *TicketFieldShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketFieldUpdateInput struct {
	ID string
}

func (c *Client) TicketFieldUpdate(i *TicketFieldUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) TicketFieldsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_fields.json"
	return c.Request("GET", path, ro)
}

type TicketFormCloneInput struct {
	ID string
}

func (c *Client) TicketFormClone(i *TicketFormCloneInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s/clone.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

func (c *Client) TicketFormCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_forms.json"
	return c.Request("POST", path, ro)
}

type TicketFormDeleteInput struct {
	ID string
}

func (c *Client) TicketFormDelete(i *TicketFormDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type TicketFormShowInput struct {
	ID string
}

func (c *Client) TicketFormShow(i *TicketFormShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketFormUpdateInput struct {
	ID string
}

func (c *Client) TicketFormUpdate(i *TicketFormUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) TicketFormsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_forms.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TicketFormsReorder(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_forms/reorder.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) TicketFormsShowMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_forms/show_many.json"
	return c.Request("GET", path, ro)
}

type TicketIncidentsInput struct {
	ID string
}

func (c *Client) TicketIncidents(i *TicketIncidentsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/incidents.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketMacroApplyInput struct {
	TicketID string
	ID string
}

func (c *Client) TicketMacroApply(i *TicketMacroApplyInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/macros/%s/apply.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	return c.Request("GET", path, ro)
}

type TicketMarkAsSpamInput struct {
	ID string
}

func (c *Client) TicketMarkAsSpam(i *TicketMarkAsSpamInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/mark_as_spam.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type TicketMergeInput struct {
	ID string
}

func (c *Client) TicketMerge(i *TicketMergeInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/merge.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type TicketMetricShowInput struct {
	TicketMetricID string
}

func (c *Client) TicketMetricShow(i *TicketMetricShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketMetricID == "" {
		return nil, errors.New("Missing required field 'TicketMetricID'")
	}

	api_path := "/api/v2/ticket_metrics/%s.json"
	path := fmt.Sprintf(api_path, i.TicketMetricID)
	return c.Request("GET", path, ro)
}

type TicketMetricsInput struct {
	TicketID string
}

func (c *Client) TicketMetrics(i *TicketMetricsInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/metrics.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	return c.Request("GET", path, ro)
}

func (c *Client) TicketMetricsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/ticket_metrics.json"
	return c.Request("GET", path, ro)
}

type TicketRelatedInput struct {
	ID string
}

func (c *Client) TicketRelated(i *TicketRelatedInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/related.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketSatisfactionRatingCreateInput struct {
	TicketID string
}

func (c *Client) TicketSatisfactionRatingCreate(i *TicketSatisfactionRatingCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/satisfaction_rating.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	return c.Request("POST", path, ro)
}

type TicketShowInput struct {
	ID string
}

func (c *Client) TicketShow(i *TicketShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketSkipsInput struct {
	TicketID string
}

func (c *Client) TicketSkips(i *TicketSkipsInput, ro *RequestOptions) (*http.Response, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/skips.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	return c.Request("GET", path, ro)
}

type TicketTagCreateInput struct {
	ID string
}

func (c *Client) TicketTagCreate(i *TicketTagCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type TicketTagsInput struct {
	ID string
}

func (c *Client) TicketTags(i *TicketTagsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TicketTagsDeleteInput struct {
	ID string
}

func (c *Client) TicketTagsDelete(i *TicketTagsDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type TicketTagsUpdateInput struct {
	ID string
}

func (c *Client) TicketTagsUpdate(i *TicketTagsUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type TicketUpdateInput struct {
	ID string
}

func (c *Client) TicketUpdate(i *TicketUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) TicketsCreateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets/create_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) TicketsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) TicketsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TicketsMarkManyAsSpam(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets/mark_many_as_spam.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) TicketsRecentList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets/recent.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TicketsShowMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets/show_many.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TicketsUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/tickets/update_many.json"
	return c.Request("PUT", path, ro)
}

type TopicTagCreateInput struct {
	ID string
}

func (c *Client) TopicTagCreate(i *TopicTagCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type TopicTagsInput struct {
	ID string
}

func (c *Client) TopicTags(i *TopicTagsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TopicTagsDeleteInput struct {
	ID string
}

func (c *Client) TopicTagsDelete(i *TopicTagsDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type TopicTagsUpdateInput struct {
	ID string
}

func (c *Client) TopicTagsUpdate(i *TopicTagsUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) TriggerCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers"
	return c.Request("POST", path, ro)
}

type TriggerDeleteByIDInput struct {
	ID string
}

func (c *Client) TriggerDeleteByID(i *TriggerDeleteByIDInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/triggers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type TriggerDeleteByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerDeleteByTriggerName(i *TriggerDeleteByTriggerNameInput, ro *RequestOptions) (*http.Response, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	api_path := "/api/v2/triggers/%s"
	path := fmt.Sprintf(api_path, i.TriggerName)
	return c.Request("DELETE", path, ro)
}

type TriggerShowByIDInput struct {
	ID string
}

func (c *Client) TriggerShowByID(i *TriggerShowByIDInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/triggers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type TriggerShowByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerShowByTriggerName(i *TriggerShowByTriggerNameInput, ro *RequestOptions) (*http.Response, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	api_path := "/api/v2/triggers/%s"
	path := fmt.Sprintf(api_path, i.TriggerName)
	return c.Request("GET", path, ro)
}

type TriggerUpdateByIDInput struct {
	ID string
}

func (c *Client) TriggerUpdateByID(i *TriggerUpdateByIDInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/triggers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type TriggerUpdateByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerUpdateByTriggerName(i *TriggerUpdateByTriggerNameInput, ro *RequestOptions) (*http.Response, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	api_path := "/api/v2/triggers/%s"
	path := fmt.Sprintf(api_path, i.TriggerName)
	return c.Request("PUT", path, ro)
}

func (c *Client) TriggersActiveList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers/active.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TriggersDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) TriggersList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers"
	return c.Request("GET", path, ro)
}

func (c *Client) TriggersReorder(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers/reorder.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) TriggersSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) TriggersUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/triggers/update_many.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) UploadCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/uploads.json"
	return c.Request("POST", path, ro)
}

type UploadDeleteInput struct {
	Token string
}

func (c *Client) UploadDelete(i *UploadDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.Token == "" {
		return nil, errors.New("Missing required field 'Token'")
	}

	api_path := "/api/v2/uploads/%s.json"
	path := fmt.Sprintf(api_path, i.Token)
	return c.Request("DELETE", path, ro)
}

func (c *Client) UserCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users.json"
	return c.Request("POST", path, ro)
}

type UserDeleteInput struct {
	ID string
}

func (c *Client) UserDelete(i *UserDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

func (c *Client) UserFieldCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/user_fields.json"
	return c.Request("POST", path, ro)
}

type UserFieldDeleteInput struct {
	ID string
}

func (c *Client) UserFieldDelete(i *UserFieldDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserFieldOptionCreateInput struct {
	FieldID string
}

func (c *Client) UserFieldOptionCreate(i *UserFieldOptionCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	api_path := "/api/v2/user_fields/%s/options.json"
	path := fmt.Sprintf(api_path, i.FieldID)
	return c.Request("POST", path, ro)
}

type UserFieldOptionDeleteInput struct {
	FieldID string
	ID string
}

func (c *Client) UserFieldOptionDelete(i *UserFieldOptionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserFieldOptionShowInput struct {
	FieldID string
	ID string
}

func (c *Client) UserFieldOptionShow(i *UserFieldOptionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	return c.Request("GET", path, ro)
}

type UserFieldOptionsInput struct {
	FieldID string
}

func (c *Client) UserFieldOptions(i *UserFieldOptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	api_path := "/api/v2/user_fields/%s/options.json"
	path := fmt.Sprintf(api_path, i.FieldID)
	return c.Request("GET", path, ro)
}

type UserFieldShowInput struct {
	ID string
}

func (c *Client) UserFieldShow(i *UserFieldShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type UserFieldUpdateInput struct {
	ID string
}

func (c *Client) UserFieldUpdate(i *UserFieldUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) UserFieldsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/user_fields.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UserFieldsReorder(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/user_fields/reorder.json"
	return c.Request("PUT", path, ro)
}

type UserGroupMembershipCreateInput struct {
	UserID string
}

func (c *Client) UserGroupMembershipCreate(i *UserGroupMembershipCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/group_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("POST", path, ro)
}

type UserGroupMembershipDeleteInput struct {
	UserID string
	ID string
}

func (c *Client) UserGroupMembershipDelete(i *UserGroupMembershipDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserGroupMembershipMakeDefaultInput struct {
	UserID string
	MembershipID string
}

func (c *Client) UserGroupMembershipMakeDefault(i *UserGroupMembershipMakeDefaultInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.MembershipID == "" {
		return nil, errors.New("Missing required field 'MembershipID'")
	}

	api_path := "/api/v2/users/%s/group_memberships/%s/make_default.json"
	path := fmt.Sprintf(api_path, i.UserID, i.MembershipID)
	return c.Request("PUT", path, ro)
}

type UserGroupMembershipShowInput struct {
	UserID string
	ID string
}

func (c *Client) UserGroupMembershipShow(i *UserGroupMembershipShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("GET", path, ro)
}

type UserGroupMembershipsInput struct {
	UserID string
}

func (c *Client) UserGroupMemberships(i *UserGroupMembershipsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/group_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserGroupsInput struct {
	UserID string
}

func (c *Client) UserGroups(i *UserGroupsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/groups.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserIdentitiesInput struct {
	UserID string
}

func (c *Client) UserIdentities(i *UserIdentitiesInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/identities.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserIdentityCreateInput struct {
	UserID string
}

func (c *Client) UserIdentityCreate(i *UserIdentityCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/identities.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("POST", path, ro)
}

type UserIdentityDeleteInput struct {
	UserID string
	ID string
}

func (c *Client) UserIdentityDelete(i *UserIdentityDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserIdentityMakePrimaryInput struct {
	UserID string
	ID string
}

func (c *Client) UserIdentityMakePrimary(i *UserIdentityMakePrimaryInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s/make_primary"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("PUT", path, ro)
}

type UserIdentityRequestVerificationUpdateInput struct {
	UserID string
	ID string
}

func (c *Client) UserIdentityRequestVerificationUpdate(i *UserIdentityRequestVerificationUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s/request_verification.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("PUT", path, ro)
}

type UserIdentityShowInput struct {
	UserID string
	ID string
}

func (c *Client) UserIdentityShow(i *UserIdentityShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("GET", path, ro)
}

type UserIdentityUpdateInput struct {
	UserID string
	ID string
}

func (c *Client) UserIdentityUpdate(i *UserIdentityUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("PUT", path, ro)
}

type UserIdentityVerifyInput struct {
	UserID string
	ID string
}

func (c *Client) UserIdentityVerify(i *UserIdentityVerifyInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s/verify"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("PUT", path, ro)
}

type UserMergeInput struct {
	ID string
}

func (c *Client) UserMerge(i *UserMergeInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/merge.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type UserOrganizationMembershipCreateInput struct {
	UserID string
}

func (c *Client) UserOrganizationMembershipCreate(i *UserOrganizationMembershipCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("POST", path, ro)
}

type UserOrganizationMembershipDeleteInput struct {
	UserID string
	ID string
}

func (c *Client) UserOrganizationMembershipDelete(i *UserOrganizationMembershipDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserOrganizationMembershipMakeDefaultInput struct {
	ID string
	MembershipID string
}

func (c *Client) UserOrganizationMembershipMakeDefault(i *UserOrganizationMembershipMakeDefaultInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	if i.MembershipID == "" {
		return nil, errors.New("Missing required field 'MembershipID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships/%s/make_default.json"
	path := fmt.Sprintf(api_path, i.ID, i.MembershipID)
	return c.Request("PUT", path, ro)
}

type UserOrganizationMembershipShowInput struct {
	UserID string
	ID string
}

func (c *Client) UserOrganizationMembershipShow(i *UserOrganizationMembershipShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("GET", path, ro)
}

type UserOrganizationMembershipsInput struct {
	UserID string
}

func (c *Client) UserOrganizationMemberships(i *UserOrganizationMembershipsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserOrganizationSubscriptionsInput struct {
	UserID string
}

func (c *Client) UserOrganizationSubscriptions(i *UserOrganizationSubscriptionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organization_subscriptions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserOrganizationsInput struct {
	UserID string
}

func (c *Client) UserOrganizations(i *UserOrganizationsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organizations.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserPasswordCreateInput struct {
	UserID string
}

func (c *Client) UserPasswordCreate(i *UserPasswordCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/password.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("POST", path, ro)
}

type UserPasswordRequirementsInput struct {
	UserID string
}

func (c *Client) UserPasswordRequirements(i *UserPasswordRequirementsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/password/requirements.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserPasswordUpdateInput struct {
	UserID string
}

func (c *Client) UserPasswordUpdate(i *UserPasswordUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/password.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("PUT", path, ro)
}

type UserRelatedInput struct {
	ID string
}

func (c *Client) UserRelated(i *UserRelatedInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/related.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type UserRequestsInput struct {
	ID string
}

func (c *Client) UserRequests(i *UserRequestsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/requests.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type UserSessionDeleteInput struct {
	UserID string
	ID string
}

func (c *Client) UserSessionDelete(i *UserSessionDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/sessions/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserSessionShowInput struct {
	UserID string
	ID string
}

func (c *Client) UserSessionShow(i *UserSessionShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/sessions/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	return c.Request("GET", path, ro)
}

type UserSessionsInput struct {
	UserID string
}

func (c *Client) UserSessions(i *UserSessionsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/sessions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserSessionsDeleteInput struct {
	UserID string
}

func (c *Client) UserSessionsDelete(i *UserSessionsDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/sessions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("DELETE", path, ro)
}

type UserShowInput struct {
	ID string
}

func (c *Client) UserShow(i *UserShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type UserSkipsInput struct {
	UserID string
}

func (c *Client) UserSkips(i *UserSkipsInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/skips.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserTagCreateInput struct {
	ID string
}

func (c *Client) UserTagCreate(i *UserTagCreateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("POST", path, ro)
}

type UserTagsInput struct {
	ID string
}

func (c *Client) UserTags(i *UserTagsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type UserTagsDeleteInput struct {
	ID string
}

func (c *Client) UserTagsDelete(i *UserTagsDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type UserTagsUpdateInput struct {
	ID string
}

func (c *Client) UserTagsUpdate(i *UserTagsUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

type UserTicketsAssignedInput struct {
	UserID string
}

func (c *Client) UserTicketsAssigned(i *UserTicketsAssignedInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/tickets/assigned.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserTicketsCcdInput struct {
	UserID string
}

func (c *Client) UserTicketsCcd(i *UserTicketsCcdInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/tickets/ccd.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserTicketsRequestedInput struct {
	UserID string
}

func (c *Client) UserTicketsRequested(i *UserTicketsRequestedInput, ro *RequestOptions) (*http.Response, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/tickets/requested.json"
	path := fmt.Sprintf(api_path, i.UserID)
	return c.Request("GET", path, ro)
}

type UserUpdateInput struct {
	ID string
}

func (c *Client) UserUpdate(i *UserUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) UsersAutocomplete(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/autocomplete.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersCreateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/create_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) UsersCreateOrUpdate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/create_or_update.json"
	return c.Request("POST", path, ro)
}

func (c *Client) UsersCreateOrUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/create_or_update_many.json"
	return c.Request("POST", path, ro)
}

func (c *Client) UsersDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) UsersList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersMe(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/me.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersMeLogout(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/me/logout.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) UsersMeMerge(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/me/merge.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) UsersMeOauthClients(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/me/oauth/clients.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersMeSession(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/me/session.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersRequestCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/request_create.json"
	return c.Request("POST", path, ro)
}

func (c *Client) UsersSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersShowMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/show_many.json"
	return c.Request("GET", path, ro)
}

func (c *Client) UsersUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/users/update_many.json"
	return c.Request("PUT", path, ro)
}

type ViewCountInput struct {
	ID string
}

func (c *Client) ViewCount(i *ViewCountInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/count.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

func (c *Client) ViewCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views.json"
	return c.Request("POST", path, ro)
}

type ViewDeleteInput struct {
	ID string
}

func (c *Client) ViewDelete(i *ViewDeleteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("DELETE", path, ro)
}

type ViewExecuteInput struct {
	ID string
}

func (c *Client) ViewExecute(i *ViewExecuteInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/execute.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ViewExportInput struct {
	ID string
}

func (c *Client) ViewExport(i *ViewExportInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/export.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ViewShowInput struct {
	ID string
}

func (c *Client) ViewShow(i *ViewShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ViewTicketsInput struct {
	ID string
}

func (c *Client) ViewTickets(i *ViewTicketsInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/tickets.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("GET", path, ro)
}

type ViewUpdateInput struct {
	ID string
}

func (c *Client) ViewUpdate(i *ViewUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	return c.Request("PUT", path, ro)
}

func (c *Client) ViewsActiveList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/active.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ViewsCompact(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/compact.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ViewsCountManyList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/count_many.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ViewsDestroyMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/destroy_many.json"
	return c.Request("DELETE", path, ro)
}

func (c *Client) ViewsList(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ViewsPreview(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/preview.json"
	return c.Request("POST", path, ro)
}

func (c *Client) ViewsPreviewCount(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/preview/count.json"
	return c.Request("POST", path, ro)
}

func (c *Client) ViewsSearch(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/search.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ViewsShowMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/show_many.json"
	return c.Request("GET", path, ro)
}

func (c *Client) ViewsUpdateMany(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/views/update_many.json"
	return c.Request("PUT", path, ro)
}

func (c *Client) VisitorCreate(ro *RequestOptions) (*http.Response, error) {
	path := "/api/v2/visitors"
	return c.Request("POST", path, ro)
}

type VisitorShowInput struct {
	VisitorID string
}

func (c *Client) VisitorShow(i *VisitorShowInput, ro *RequestOptions) (*http.Response, error) {
	if i.VisitorID == "" {
		return nil, errors.New("Missing required field 'VisitorID'")
	}

	api_path := "/api/v2/visitors/%s"
	path := fmt.Sprintf(api_path, i.VisitorID)
	return c.Request("GET", path, ro)
}

type VisitorUpdateInput struct {
	VisitorID string
}

func (c *Client) VisitorUpdate(i *VisitorUpdateInput, ro *RequestOptions) (*http.Response, error) {
	if i.VisitorID == "" {
		return nil, errors.New("Missing required field 'VisitorID'")
	}

	api_path := "/api/v2/visitors/%s"
	path := fmt.Sprintf(api_path, i.VisitorID)
	return c.Request("PUT", path, ro)
}


