package zdesk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handle(resp *http.Response, ro *RequestOptions, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	if ro != nil {
		if ro.ResponseOption == "location" {
			return []byte(resp.Header.Get("Location")), nil
		}

		if ro.ResponseOption == "code" {
			return []byte(fmt.Sprintf("%d", resp.StatusCode)), nil
		}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) AccountList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/account"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AccountSettingsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/account/settings.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AccountSettingsUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/account/settings.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AccountUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/account"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ActivitiesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/activities.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ActivityShowInput struct {
	ActivityID string
}

func (c *Client) ActivityShow(i *ActivityShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ActivityID == "" {
		return nil, errors.New("Missing required field 'ActivityID'")
	}

	api_path := "/api/v2/activities/%s.json"
	path := fmt.Sprintf(api_path, i.ActivityID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AgentCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/agents"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type AgentDeleteInput struct {
	AgentID string
}

func (c *Client) AgentDelete(i *AgentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	api_path := "/api/v2/agents/%s"
	path := fmt.Sprintf(api_path, i.AgentID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type AgentShowInput struct {
	AgentID string
}

func (c *Client) AgentShow(i *AgentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	api_path := "/api/v2/agents/%s"
	path := fmt.Sprintf(api_path, i.AgentID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AgentUpdateInput struct {
	AgentID string
}

func (c *Client) AgentUpdate(i *AgentUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	api_path := "/api/v2/agents/%s"
	path := fmt.Sprintf(api_path, i.AgentID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type AgentsEmailShowInput struct {
	EmailID string
}

func (c *Client) AgentsEmailShow(i *AgentsEmailShowInput, ro *RequestOptions) ([]byte, error) {
	if i.EmailID == "" {
		return nil, errors.New("Missing required field 'EmailID'")
	}

	api_path := "/api/v2/agents/email/%s"
	path := fmt.Sprintf(api_path, i.EmailID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AgentsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/agents"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AgentsMe(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/agents/me"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AgentsMeUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/agents/me"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AnyChannelPushCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/any_channel/push"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type AppDeleteInput struct {
	ID string
}

func (c *Client) AppDelete(i *AppDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type AppPublicKeyInput struct {
	ID string
}

func (c *Client) AppPublicKey(i *AppPublicKeyInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s/public_key.pem"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AppShowInput struct {
	ID string
}

func (c *Client) AppShow(i *AppShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AppUpdateInput struct {
	ID string
}

func (c *Client) AppUpdate(i *AppUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsInstallationCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/installations.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type AppsInstallationDeleteInput struct {
	ID string
}

func (c *Client) AppsInstallationDelete(i *AppsInstallationDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type AppsInstallationRequirementsInput struct {
	ID string
}

func (c *Client) AppsInstallationRequirements(i *AppsInstallationRequirementsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s/requirements.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AppsInstallationShowInput struct {
	ID string
}

func (c *Client) AppsInstallationShow(i *AppsInstallationShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AppsInstallationUpdateInput struct {
	ID string
}

func (c *Client) AppsInstallationUpdate(i *AppsInstallationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type AppsInstallationsJobStatusShowInput struct {
	ID string
}

func (c *Client) AppsInstallationsJobStatusShow(i *AppsInstallationsJobStatusShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/installations/job_statuses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsInstallationsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/installations.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AppsJobStatusShowInput struct {
	ID string
}

func (c *Client) AppsJobStatusShow(i *AppsJobStatusShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/job_statuses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsLocationInstallationsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/location_installations.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsLocationInstallationsReorder(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/location_installations/reorder.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type AppsLocationShowInput struct {
	ID string
}

func (c *Client) AppsLocationShow(i *AppsLocationShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/apps/locations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsLocationsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/locations.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsNotifyCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/notify.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsOwnedList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/owned.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AppsUploadCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/apps/uploads.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type AttachmentDeleteInput struct {
	ID string
}

func (c *Client) AttachmentDelete(i *AttachmentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type AttachmentShowInput struct {
	ID string
}

func (c *Client) AttachmentShow(i *AttachmentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AuditLogShowInput struct {
	ID string
}

func (c *Client) AuditLogShow(i *AuditLogShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/audit_logs/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AuditLogsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/audit_logs.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutocompleteTags(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/autocomplete/tags.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutomationCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/automations.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type AutomationDeleteInput struct {
	ID string
}

func (c *Client) AutomationDelete(i *AutomationDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/automations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type AutomationShowInput struct {
	ID string
}

func (c *Client) AutomationShow(i *AutomationShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/automations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type AutomationUpdateInput struct {
	ID string
}

func (c *Client) AutomationUpdate(i *AutomationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/automations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutomationsActiveList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/automations/active.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutomationsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/automations/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutomationsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/automations.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutomationsSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/automations/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) AutomationsUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/automations/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BanCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/bans"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type BanDeleteInput struct {
	BanID string
}

func (c *Client) BanDelete(i *BanDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.BanID == "" {
		return nil, errors.New("Missing required field 'BanID'")
	}

	api_path := "/api/v2/ban/%s"
	path := fmt.Sprintf(api_path, i.BanID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type BanShowInput struct {
	BanID string
}

func (c *Client) BanShow(i *BanShowInput, ro *RequestOptions) ([]byte, error) {
	if i.BanID == "" {
		return nil, errors.New("Missing required field 'BanID'")
	}

	api_path := "/api/v2/bans/%s"
	path := fmt.Sprintf(api_path, i.BanID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BansIpList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/bans/ip"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BansList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/bans"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BookmarkCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/bookmarks.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type BookmarkDeleteInput struct {
	ID string
}

func (c *Client) BookmarkDelete(i *BookmarkDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/bookmarks/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BookmarksList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/bookmarks.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type BrandCheckHostMappingInput struct {
	ID string
}

func (c *Client) BrandCheckHostMapping(i *BrandCheckHostMappingInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s/check_host_mapping.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BrandCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/brands.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type BrandDeleteInput struct {
	ID string
}

func (c *Client) BrandDelete(i *BrandDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type BrandShowInput struct {
	ID string
}

func (c *Client) BrandShow(i *BrandShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type BrandUpdateInput struct {
	ID string
}

func (c *Client) BrandUpdate(i *BrandUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/brands/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BrandsCheckHostMappingList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/brands/check_host_mapping.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BrandsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/brands.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BusinessHoursScheduleCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/business_hours/schedules.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleDeleteInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleDelete(i *BusinessHoursScheduleDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleHolidayCreateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleHolidayCreate(i *BusinessHoursScheduleHolidayCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleHolidayDeleteInput struct {
	ScheduleID string
	ID         string
}

func (c *Client) BusinessHoursScheduleHolidayDelete(i *BusinessHoursScheduleHolidayDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays/%s.json"
	path := fmt.Sprintf(api_path, i.ScheduleID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleHolidayShowInput struct {
	ScheduleID string
	ID         string
}

func (c *Client) BusinessHoursScheduleHolidayShow(i *BusinessHoursScheduleHolidayShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays/%s.json"
	path := fmt.Sprintf(api_path, i.ScheduleID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleHolidayUpdateInput struct {
	ScheduleID string
	ID         string
}

func (c *Client) BusinessHoursScheduleHolidayUpdate(i *BusinessHoursScheduleHolidayUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays/%s.json"
	path := fmt.Sprintf(api_path, i.ScheduleID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleHolidaysInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleHolidays(i *BusinessHoursScheduleHolidaysInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/holidays.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleShowInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleShow(i *BusinessHoursScheduleShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleUpdateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleUpdate(i *BusinessHoursScheduleUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type BusinessHoursScheduleWorkweekUpdateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleWorkweekUpdate(i *BusinessHoursScheduleWorkweekUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/business_hours/schedules/%s/workweek.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) BusinessHoursSchedulesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/business_hours/schedules.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsTwitterMonitoredTwitterHandleShowInput struct {
	ID string
}

func (c *Client) ChannelsTwitterMonitoredTwitterHandleShow(i *ChannelsTwitterMonitoredTwitterHandleShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/twitter/monitored_twitter_handles/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsTwitterMonitoredTwitterHandlesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/twitter/monitored_twitter_handles.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsTwitterTicketCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/twitter/tickets.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ChannelsTwitterTicketStatusesInput struct {
	ID string
}

func (c *Client) ChannelsTwitterTicketStatuses(i *ChannelsTwitterTicketStatusesInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/twitter/tickets/%s/statuses.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceAgentTicketDisplayCreateInput struct {
	AgentID  string
	TicketID string
}

func (c *Client) ChannelsVoiceAgentTicketDisplayCreate(i *ChannelsVoiceAgentTicketDisplayCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/channels/voice/agents/%s/tickets/%s/display.json"
	path := fmt.Sprintf(api_path, i.AgentID, i.TicketID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceAgentUserDisplayCreateInput struct {
	AgentID string
	UserID  string
}

func (c *Client) ChannelsVoiceAgentUserDisplayCreate(i *ChannelsVoiceAgentUserDisplayCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/channels/voice/agents/%s/users/%s/display.json"
	path := fmt.Sprintf(api_path, i.AgentID, i.UserID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceAvailabilityShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceAvailabilityShow(i *ChannelsVoiceAvailabilityShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/availabilities/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceAvailabilityUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoiceAvailabilityUpdate(i *ChannelsVoiceAvailabilityUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/availabilities/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceGreetingCategoriesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/greeting_categories.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceGreetingCategoryShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingCategoryShow(i *ChannelsVoiceGreetingCategoryShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greeting_categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceGreetingCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/greetings.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceGreetingDeleteInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingDelete(i *ChannelsVoiceGreetingDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceGreetingRecordingInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingRecording(i *ChannelsVoiceGreetingRecordingInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s/recording.mp3"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceGreetingShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingShow(i *ChannelsVoiceGreetingShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoiceGreetingUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingUpdate(i *ChannelsVoiceGreetingUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/greetings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceGreetingsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/greetings.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoicePhoneNumberCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/phone_numbers.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoicePhoneNumberShowInput struct {
	ID string
}

func (c *Client) ChannelsVoicePhoneNumberShow(i *ChannelsVoicePhoneNumberShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/phone_numbers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChannelsVoicePhoneNumberUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoicePhoneNumberUpdate(i *ChannelsVoicePhoneNumberUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/channels/voice/phone_numbers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoicePhoneNumbersDelete(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/phone_numbers.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoicePhoneNumbersList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/phone_numbers.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoicePhoneNumbersSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/phone_numbers/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceStatsAgentsActivityList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/stats/agents_activity.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceStatsCurrentQueueActivityList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/stats/current_queue_activity.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceStatsHistoricalQueueActivityList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/stats/historical_queue_activity.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChannelsVoiceTicketCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/channels/voice/tickets.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChatCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/chats"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ChatDeleteInput struct {
	ChatID string
}

func (c *Client) ChatDelete(i *ChatDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	api_path := "/api/v2/chats/%s"
	path := fmt.Sprintf(api_path, i.ChatID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type ChatShowInput struct {
	ChatID string
}

func (c *Client) ChatShow(i *ChatShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	api_path := "/api/v2/chats/%s"
	path := fmt.Sprintf(api_path, i.ChatID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ChatUpdateInput struct {
	ChatID string
}

func (c *Client) ChatUpdate(i *ChatUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	api_path := "/api/v2/chats/%s"
	path := fmt.Sprintf(api_path, i.ChatID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChatsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/chats"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ChatsSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/chats/search"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentCreateInput struct {
	PostID string
}

func (c *Client) CommunityPostCommentCreate(i *CommunityPostCommentCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/comments.json"
	path := fmt.Sprintf(api_path, i.PostID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentDeleteInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentDelete(i *CommunityPostCommentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentDownCreateInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentDownCreate(i *CommunityPostCommentDownCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s/down.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentShowInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentShow(i *CommunityPostCommentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentUpCreateInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentUpCreate(i *CommunityPostCommentUpCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s/up.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentUpdateInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentUpdate(i *CommunityPostCommentUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentVotesInput struct {
	PostID    string
	CommentID string
}

func (c *Client) CommunityPostCommentVotes(i *CommunityPostCommentVotesInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	api_path := "/api/v2/community/posts/%s/comments/%s/votes.json"
	path := fmt.Sprintf(api_path, i.PostID, i.CommentID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostCommentsInput struct {
	PostID string
}

func (c *Client) CommunityPostComments(i *CommunityPostCommentsInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/comments.json"
	path := fmt.Sprintf(api_path, i.PostID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) CommunityPostCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/community/posts.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostDeleteInput struct {
	ID string
}

func (c *Client) CommunityPostDelete(i *CommunityPostDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostDownCreateInput struct {
	ID string
}

func (c *Client) CommunityPostDownCreate(i *CommunityPostDownCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/down.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostShowInput struct {
	ID string
}

func (c *Client) CommunityPostShow(i *CommunityPostShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostSubscriptionCreateInput struct {
	PostID string
}

func (c *Client) CommunityPostSubscriptionCreate(i *CommunityPostSubscriptionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.PostID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostSubscriptionDeleteInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostSubscriptionDelete(i *CommunityPostSubscriptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostSubscriptionShowInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostSubscriptionShow(i *CommunityPostSubscriptionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.PostID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostSubscriptionsInput struct {
	PostID string
}

func (c *Client) CommunityPostSubscriptions(i *CommunityPostSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.PostID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostUpCreateInput struct {
	ID string
}

func (c *Client) CommunityPostUpCreate(i *CommunityPostUpCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s/up.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostUpdateInput struct {
	ID string
}

func (c *Client) CommunityPostUpdate(i *CommunityPostUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/posts/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type CommunityPostVotesInput struct {
	PostID string
}

func (c *Client) CommunityPostVotes(i *CommunityPostVotesInput, ro *RequestOptions) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	api_path := "/api/v2/community/posts/%s/votes.json"
	path := fmt.Sprintf(api_path, i.PostID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) CommunityPostsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/community/posts.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicAccessPolicyInput struct {
	TopicID string
}

func (c *Client) CommunityTopicAccessPolicy(i *CommunityTopicAccessPolicyInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicAccessPolicyUpdateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicAccessPolicyUpdate(i *CommunityTopicAccessPolicyUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) CommunityTopicCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/community/topics.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicDeleteInput struct {
	ID string
}

func (c *Client) CommunityTopicDelete(i *CommunityTopicDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicPostsInput struct {
	ID string
}

func (c *Client) CommunityTopicPosts(i *CommunityTopicPostsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s/posts.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicShowInput struct {
	ID string
}

func (c *Client) CommunityTopicShow(i *CommunityTopicShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicSubscriptionCreateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptionCreate(i *CommunityTopicSubscriptionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicSubscriptionDeleteInput struct {
	TopicID string
	ID      string
}

func (c *Client) CommunityTopicSubscriptionDelete(i *CommunityTopicSubscriptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.TopicID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicSubscriptionShowInput struct {
	TopicID string
	ID      string
}

func (c *Client) CommunityTopicSubscriptionShow(i *CommunityTopicSubscriptionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.TopicID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicSubscriptionsInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptions(i *CommunityTopicSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicSubscriptionsUpdateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptionsUpdate(i *CommunityTopicSubscriptionsUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	api_path := "/api/v2/community/topics/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.TopicID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type CommunityTopicUpdateInput struct {
	ID string
}

func (c *Client) CommunityTopicUpdate(i *CommunityTopicUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/topics/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) CommunityTopicsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/community/topics.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityUserCommentsInput struct {
	ID string
}

func (c *Client) CommunityUserComments(i *CommunityUserCommentsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/users/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type CommunityUserPostsInput struct {
	ID string
}

func (c *Client) CommunityUserPosts(i *CommunityUserPostsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/community/users/%s/posts.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) CustomRolesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/custom_roles.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type DeletedTicketRestoreUpdateInput struct {
	ID string
}

func (c *Client) DeletedTicketRestoreUpdate(i *DeletedTicketRestoreUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/deleted_tickets/%s/restore.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DeletedTickets1Delete(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/deleted_tickets/1"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DeletedTicketsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/deleted_tickets/destroy_many"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DeletedTicketsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/deleted_tickets.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DeletedTicketsRestoreManyUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/deleted_tickets/restore_many"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DepartmentCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/departments"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type DepartmentDeleteInput struct {
	DepartmentID string
}

func (c *Client) DepartmentDelete(i *DepartmentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	api_path := "/api/v2/departments/%s"
	path := fmt.Sprintf(api_path, i.DepartmentID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type DepartmentShowInput struct {
	DepartmentID string
}

func (c *Client) DepartmentShow(i *DepartmentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	api_path := "/api/v2/departments/%s"
	path := fmt.Sprintf(api_path, i.DepartmentID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type DepartmentUpdateInput struct {
	DepartmentIDOrName string
}

func (c *Client) DepartmentUpdate(i *DepartmentUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.DepartmentIDOrName == "" {
		return nil, errors.New("Missing required field 'DepartmentIDOrName'")
	}

	api_path := "/api/v2/departments/%s"
	path := fmt.Sprintf(api_path, i.DepartmentIDOrName)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DepartmentsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/departments"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type DepartmentsNameDeleteInput struct {
	Name string
}

func (c *Client) DepartmentsNameDelete(i *DepartmentsNameDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/departments/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type DepartmentsNameShowInput struct {
	Name string
}

func (c *Client) DepartmentsNameShow(i *DepartmentsNameShowInput, ro *RequestOptions) ([]byte, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	api_path := "/api/v2/departments/name/%s"
	path := fmt.Sprintf(api_path, i.Name)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DynamicContentItemCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/dynamic_content/items.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemDeleteInput struct {
	ID string
}

func (c *Client) DynamicContentItemDelete(i *DynamicContentItemDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemShowInput struct {
	ID string
}

func (c *Client) DynamicContentItemShow(i *DynamicContentItemShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemUpdateInput struct {
	ID string
}

func (c *Client) DynamicContentItemUpdate(i *DynamicContentItemUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantCreateInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantCreate(i *DynamicContentItemVariantCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantDeleteInput struct {
	ItemID string
	ID     string
}

func (c *Client) DynamicContentItemVariantDelete(i *DynamicContentItemVariantDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/%s.json"
	path := fmt.Sprintf(api_path, i.ItemID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantShowInput struct {
	ItemID string
	ID     string
}

func (c *Client) DynamicContentItemVariantShow(i *DynamicContentItemVariantShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/%s.json"
	path := fmt.Sprintf(api_path, i.ItemID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantUpdateInput struct {
	ItemID string
	ID     string
}

func (c *Client) DynamicContentItemVariantUpdate(i *DynamicContentItemVariantUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/%s.json"
	path := fmt.Sprintf(api_path, i.ItemID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantsInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariants(i *DynamicContentItemVariantsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantsCreateManyInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantsCreateMany(i *DynamicContentItemVariantsCreateManyInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/create_many.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type DynamicContentItemVariantsUpdateManyInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantsUpdateMany(i *DynamicContentItemVariantsUpdateManyInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/dynamic_content/items/%s/variants/update_many.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) DynamicContentItemsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/dynamic_content/items.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type EndUserIdentityCreateInput struct {
	UserID string
}

func (c *Client) EndUserIdentityCreate(i *EndUserIdentityCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/end_users/%s/identities.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type EndUserIdentityMakePrimaryInput struct {
	UserID string
	ID     string
}

func (c *Client) EndUserIdentityMakePrimary(i *EndUserIdentityMakePrimaryInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/end_users/%s/identities/%s/make_primary"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type EndUserShowInput struct {
	ID string
}

func (c *Client) EndUserShow(i *EndUserShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/end_users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type EndUserUpdateInput struct {
	ID string
}

func (c *Client) EndUserUpdate(i *EndUserUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/end_users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GoalCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/goals"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type GoalDeleteInput struct {
	GoalID string
}

func (c *Client) GoalDelete(i *GoalDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	api_path := "/api/v2/goals/%s"
	path := fmt.Sprintf(api_path, i.GoalID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type GoalShowInput struct {
	GoalID string
}

func (c *Client) GoalShow(i *GoalShowInput, ro *RequestOptions) ([]byte, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	api_path := "/api/v2/goals/%s"
	path := fmt.Sprintf(api_path, i.GoalID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type GoalUpdateInput struct {
	GoalID string
}

func (c *Client) GoalUpdate(i *GoalUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	api_path := "/api/v2/goals/%s"
	path := fmt.Sprintf(api_path, i.GoalID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GoalsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/goals"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/groups.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type GroupDeleteInput struct {
	ID string
}

func (c *Client) GroupDelete(i *GroupDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupMembershipCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/group_memberships.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type GroupMembershipDeleteInput struct {
	ID string
}

func (c *Client) GroupMembershipDelete(i *GroupMembershipDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type GroupMembershipShowInput struct {
	ID string
}

func (c *Client) GroupMembershipShow(i *GroupMembershipShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type GroupMembershipsInput struct {
	GroupID string
}

func (c *Client) GroupMemberships(i *GroupMembershipsInput, ro *RequestOptions) ([]byte, error) {
	if i.GroupID == "" {
		return nil, errors.New("Missing required field 'GroupID'")
	}

	api_path := "/api/v2/groups/%s/memberships.json"
	path := fmt.Sprintf(api_path, i.GroupID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type GroupMembershipsAssignableInput struct {
	GroupID string
}

func (c *Client) GroupMembershipsAssignable(i *GroupMembershipsAssignableInput, ro *RequestOptions) ([]byte, error) {
	if i.GroupID == "" {
		return nil, errors.New("Missing required field 'GroupID'")
	}

	api_path := "/api/v2/groups/%s/memberships/assignable.json"
	path := fmt.Sprintf(api_path, i.GroupID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupMembershipsAssignableList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/group_memberships/assignable.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupMembershipsCreateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/group_memberships/create_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupMembershipsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/group_memberships/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupMembershipsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/group_memberships.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type GroupShowInput struct {
	ID string
}

func (c *Client) GroupShow(i *GroupShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type GroupUpdateInput struct {
	ID string
}

func (c *Client) GroupUpdate(i *GroupUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type GroupUsersInput struct {
	ID string
}

func (c *Client) GroupUsers(i *GroupUsersInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/groups/%s/users.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupsAssignableList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/groups/assignable.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) GroupsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/groups.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleAttachmentCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleAttachmentCreate(i *HelpCenterArticleAttachmentCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleAttachmentsInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleAttachments(i *HelpCenterArticleAttachmentsInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/attachments.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleAttachmentsBlockInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleAttachmentsBlock(i *HelpCenterArticleAttachmentsBlockInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments/block.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/attachments/block.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleAttachmentsInlineInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleAttachmentsInline(i *HelpCenterArticleAttachmentsInlineInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/attachments/inline.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/attachments/inline.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleBulkAttachmentCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleBulkAttachmentCreate(i *HelpCenterArticleBulkAttachmentCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/bulk_attachments.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleCommentCreate(i *HelpCenterArticleCommentCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentDeleteInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentDelete(i *HelpCenterArticleCommentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentDownCreateInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentDownCreate(i *HelpCenterArticleCommentDownCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s/down.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentShowInput struct {
	ArticleID string
	ID        string
	Locale    string
}

func (c *Client) HelpCenterArticleCommentShow(i *HelpCenterArticleCommentShowInput, ro *RequestOptions) ([]byte, error) {
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
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentUpCreateInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentUpCreate(i *HelpCenterArticleCommentUpCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s/up.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentUpdateInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentUpdate(i *HelpCenterArticleCommentUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentVotesInput struct {
	ArticleID string
	CommentID string
	Locale    string
}

func (c *Client) HelpCenterArticleCommentVotes(i *HelpCenterArticleCommentVotesInput, ro *RequestOptions) ([]byte, error) {
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
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleCommentsInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterArticleComments(i *HelpCenterArticleCommentsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{id}/comments.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterArticleDelete(i *HelpCenterArticleDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleDownCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleDownCreate(i *HelpCenterArticleDownCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/down.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleLabelCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleLabelCreate(i *HelpCenterArticleLabelCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/labels.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleLabelDeleteInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleLabelDelete(i *HelpCenterArticleLabelDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/labels/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleLabelsInput struct {
	ID string
}

func (c *Client) HelpCenterArticleLabels(i *HelpCenterArticleLabelsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/labels.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleShowInput struct {
	Locale string
	ID     string
}

func (c *Client) HelpCenterArticleShow(i *HelpCenterArticleShowInput, ro *RequestOptions) ([]byte, error) {
	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/%s/articles/%s.json"
	path := fmt.Sprintf(api_path, i.Locale, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleSourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleSourceLocaleUpdate(i *HelpCenterArticleSourceLocaleUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/source_locale.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleSubscriptionCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleSubscriptionCreate(i *HelpCenterArticleSubscriptionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleSubscriptionDeleteInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleSubscriptionDelete(i *HelpCenterArticleSubscriptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleSubscriptionShowInput struct {
	ArticleID string
	ID        string
	Locale    string
}

func (c *Client) HelpCenterArticleSubscriptionShow(i *HelpCenterArticleSubscriptionShowInput, ro *RequestOptions) ([]byte, error) {
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
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleSubscriptionsInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleSubscriptions(i *HelpCenterArticleSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/subscriptions.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleTranslationCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslationCreate(i *HelpCenterArticleTranslationCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleTranslationShowInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleTranslationShow(i *HelpCenterArticleTranslationShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.Locale)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleTranslationUpdateInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleTranslationUpdate(i *HelpCenterArticleTranslationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.ArticleID, i.Locale)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleTranslationsInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslations(i *HelpCenterArticleTranslationsInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleTranslationsMissingInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslationsMissing(i *HelpCenterArticleTranslationsMissingInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/translations/missing.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleUpCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleUpCreate(i *HelpCenterArticleUpCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s/up.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleUpdateInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterArticleUpdate(i *HelpCenterArticleUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticleVotesInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleVotes(i *HelpCenterArticleVotesInput, ro *RequestOptions) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	api_path := "/api/v2/help_center/articles/%s/votes.json"
	path := fmt.Sprintf(api_path, i.ArticleID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/{article_id}/votes.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterArticlesAttachmentCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/articles/attachments.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticlesAttachmentDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesAttachmentDelete(i *HelpCenterArticlesAttachmentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticlesAttachmentShowInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterArticlesAttachmentShow(i *HelpCenterArticlesAttachmentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/attachments/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticlesLabelShowInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesLabelShow(i *HelpCenterArticlesLabelShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/articles/labels/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticlesLabelsListInput struct {
	Locale string
}

func (c *Client) HelpCenterArticlesLabelsList(i *HelpCenterArticlesLabelsListInput, ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/articles/labels.json"
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles/labels.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterArticlesListInput struct {
	Locale string
}

func (c *Client) HelpCenterArticlesList(i *HelpCenterArticlesListInput, ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/articles.json"
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterArticlesSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/articles/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoriesListInput struct {
	Locale string
}

func (c *Client) HelpCenterCategoriesList(i *HelpCenterCategoriesListInput, ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/categories.json"
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryArticlesInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterCategoryArticles(i *HelpCenterCategoryArticlesInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryCreateInput struct {
	Locale string
}

func (c *Client) HelpCenterCategoryCreate(i *HelpCenterCategoryCreateInput, ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/categories.json"
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterCategoryDelete(i *HelpCenterCategoryDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategorySectionCreateInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterCategorySectionCreate(i *HelpCenterCategorySectionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/sections.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}/sections.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategorySectionsInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterCategorySections(i *HelpCenterCategorySectionsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/sections.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}/sections.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryShowInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterCategoryShow(i *HelpCenterCategoryShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategorySourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterCategorySourceLocaleUpdate(i *HelpCenterCategorySourceLocaleUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s/source_locale.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryTranslationCreateInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslationCreate(i *HelpCenterCategoryTranslationCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations.json"
	path := fmt.Sprintf(api_path, i.CategoryID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryTranslationUpdateInput struct {
	CategoryID string
	Locale     string
}

func (c *Client) HelpCenterCategoryTranslationUpdate(i *HelpCenterCategoryTranslationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.CategoryID, i.Locale)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryTranslationsInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslations(i *HelpCenterCategoryTranslationsInput, ro *RequestOptions) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations.json"
	path := fmt.Sprintf(api_path, i.CategoryID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryTranslationsMissingInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslationsMissing(i *HelpCenterCategoryTranslationsMissingInput, ro *RequestOptions) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	api_path := "/api/v2/help_center/categories/%s/translations/missing.json"
	path := fmt.Sprintf(api_path, i.CategoryID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterCategoryUpdateInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterCategoryUpdate(i *HelpCenterCategoryUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/categories/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/categories/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterIncrementalArticlesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/incremental/articles.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterLocalesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/locales.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionAccessPolicyInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionAccessPolicy(i *HelpCenterSectionAccessPolicyInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionAccessPolicyUpdateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionAccessPolicyUpdate(i *HelpCenterSectionAccessPolicyUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/access_policy.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionArticleCreateInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterSectionArticleCreate(i *HelpCenterSectionArticleCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionArticlesInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterSectionArticles(i *HelpCenterSectionArticlesInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}/articles.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterSectionDelete(i *HelpCenterSectionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionShowInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterSectionShow(i *HelpCenterSectionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionSourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterSectionSourceLocaleUpdate(i *HelpCenterSectionSourceLocaleUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/source_locale.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionSubscriptionCreateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionSubscriptionCreate(i *HelpCenterSectionSubscriptionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionSubscriptionDeleteInput struct {
	SectionID string
	ID        string
}

func (c *Client) HelpCenterSectionSubscriptionDelete(i *HelpCenterSectionSubscriptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.SectionID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionSubscriptionShowInput struct {
	SectionID string
	ID        string
	Locale    string
}

func (c *Client) HelpCenterSectionSubscriptionShow(i *HelpCenterSectionSubscriptionShowInput, ro *RequestOptions) ([]byte, error) {
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
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionSubscriptionsInput struct {
	SectionID string
	Locale    string
}

func (c *Client) HelpCenterSectionSubscriptions(i *HelpCenterSectionSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{section_id}/subscriptions.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionTranslationCreateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslationCreate(i *HelpCenterSectionTranslationCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionTranslationUpdateInput struct {
	SectionID string
	Locale    string
}

func (c *Client) HelpCenterSectionTranslationUpdate(i *HelpCenterSectionTranslationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations/%s.json"
	path := fmt.Sprintf(api_path, i.SectionID, i.Locale)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionTranslationsInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslations(i *HelpCenterSectionTranslationsInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionTranslationsMissingInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslationsMissing(i *HelpCenterSectionTranslationsMissingInput, ro *RequestOptions) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	api_path := "/api/v2/help_center/sections/%s/translations/missing.json"
	path := fmt.Sprintf(api_path, i.SectionID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionUpdateInput struct {
	ID     string
	Locale string
}

func (c *Client) HelpCenterSectionUpdate(i *HelpCenterSectionUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/sections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections/{id}.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterSectionsListInput struct {
	Locale string
}

func (c *Client) HelpCenterSectionsList(i *HelpCenterSectionsListInput, ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/sections.json"
	if i.Locale != "" {
		api_opt_path := "/api/v2/help_center/{locale}/sections.json"
		path = fmt.Sprintf(api_opt_path, i.Locale)
	}
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterTranslationDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterTranslationDelete(i *HelpCenterTranslationDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/translations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserArticlesInput struct {
	ID string
}

func (c *Client) HelpCenterUserArticles(i *HelpCenterUserArticlesInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/users/%s/articles.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserCommentsInput struct {
	ID string
}

func (c *Client) HelpCenterUserComments(i *HelpCenterUserCommentsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/users/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterUserSegmentCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/user_segments.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserSegmentDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentDelete(i *HelpCenterUserSegmentDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserSegmentSectionsInput struct {
	UserSegmentID string
}

func (c *Client) HelpCenterUserSegmentSections(i *HelpCenterUserSegmentSectionsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserSegmentID == "" {
		return nil, errors.New("Missing required field 'UserSegmentID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s/sections.json"
	path := fmt.Sprintf(api_path, i.UserSegmentID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserSegmentShowInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentShow(i *HelpCenterUserSegmentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserSegmentTopicsInput struct {
	UserSegmentID string
}

func (c *Client) HelpCenterUserSegmentTopics(i *HelpCenterUserSegmentTopicsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserSegmentID == "" {
		return nil, errors.New("Missing required field 'UserSegmentID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s/topics.json"
	path := fmt.Sprintf(api_path, i.UserSegmentID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserSegmentUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentUpdate(i *HelpCenterUserSegmentUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/user_segments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterUserSegmentsApplicableList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/user_segments/applicable.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) HelpCenterUserSegmentsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/help_center/user_segments.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserSubscriptionsInput struct {
	UserID string
}

func (c *Client) HelpCenterUserSubscriptions(i *HelpCenterUserSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/help_center/users/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterUserVotesInput struct {
	UserID string
}

func (c *Client) HelpCenterUserVotes(i *HelpCenterUserVotesInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/help_center/users/%s/votes.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterVoteDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterVoteDelete(i *HelpCenterVoteDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/votes/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type HelpCenterVoteShowInput struct {
	ID string
}

func (c *Client) HelpCenterVoteShow(i *HelpCenterVoteShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/help_center/votes/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ImportsTicket(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/imports/tickets.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ImportsTicketsCreateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/imports/tickets/create_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) IncrementalOrganizationsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/incremental/organizations.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type IncrementalSampleInput struct {
	Item string
}

func (c *Client) IncrementalSample(i *IncrementalSampleInput, ro *RequestOptions) ([]byte, error) {
	if i.Item == "" {
		return nil, errors.New("Missing required field 'Item'")
	}

	api_path := "/api/v2/incremental/%s/sample.json"
	path := fmt.Sprintf(api_path, i.Item)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) IncrementalTicketEventsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/incremental/ticket_events.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) IncrementalTicketMetricEventsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/incremental/ticket_metric_events.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) IncrementalTicketsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/incremental/tickets.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) IncrementalUsersList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/incremental/users.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type JobStatusShowInput struct {
	ID string
}

func (c *Client) JobStatusShow(i *JobStatusShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/job_statuses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) JobStatusesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/job_statuses.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) JobStatusesShowMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/job_statuses/show_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type LocaleShowInput struct {
	ID string
}

func (c *Client) LocaleShow(i *LocaleShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/locales/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) LocalesAgentList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/locales/agent.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) LocalesCurrentList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/locales/current.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) LocalesDetectBestLocale(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/locales/detect_best_locale.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) LocalesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/locales.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) LocalesPublicList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/locales/public.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type MacroApplyInput struct {
	ID string
}

func (c *Client) MacroApply(i *MacroApplyInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s/apply.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type MacroAttachmentCreateInput struct {
	MacroID string
}

func (c *Client) MacroAttachmentCreate(i *MacroAttachmentCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.MacroID == "" {
		return nil, errors.New("Missing required field 'MacroID'")
	}

	api_path := "/api/v2/macros/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.MacroID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type MacroAttachmentsInput struct {
	MacroID string
}

func (c *Client) MacroAttachments(i *MacroAttachmentsInput, ro *RequestOptions) ([]byte, error) {
	if i.MacroID == "" {
		return nil, errors.New("Missing required field 'MacroID'")
	}

	api_path := "/api/v2/macros/%s/attachments.json"
	path := fmt.Sprintf(api_path, i.MacroID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacroCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type MacroDeleteInput struct {
	ID string
}

func (c *Client) MacroDelete(i *MacroDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type MacroShowInput struct {
	ID string
}

func (c *Client) MacroShow(i *MacroShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type MacroUpdateInput struct {
	ID string
}

func (c *Client) MacroUpdate(i *MacroUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosActionsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/actions.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosActiveList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/active.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosAttachmentCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/attachments.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type MacrosAttachmentShowInput struct {
	ID string
}

func (c *Client) MacrosAttachmentShow(i *MacrosAttachmentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/macros/attachments/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosCategoriesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/categories.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosNewList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/new.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) MacrosUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/macros/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) NpsIncrementalRecipientsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/nps/incremental/recipients.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) NpsIncrementalResponsesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/nps/incremental/responses.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyCloseCreateInput struct {
	ID string
}

func (c *Client) NpsSurveyCloseCreate(i *NpsSurveyCloseCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/close"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyInvitationCreateInput struct {
	ID string
}

func (c *Client) NpsSurveyInvitationCreate(i *NpsSurveyInvitationCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/invitations.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyInvitationShowInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyInvitationShow(i *NpsSurveyInvitationShowInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/invitations/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyInvitationsInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyInvitations(i *NpsSurveyInvitationsInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/invitations.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyPreviewInput struct {
	ID string
}

func (c *Client) NpsSurveyPreview(i *NpsSurveyPreviewInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/preview"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyRecipientCreateInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipientCreate(i *NpsSurveyRecipientCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyRecipientShowInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyRecipientShow(i *NpsSurveyRecipientShowInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyRecipientUpdateInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyRecipientUpdate(i *NpsSurveyRecipientUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyRecipientsInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipients(i *NpsSurveyRecipientsInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyRecipientsSearchInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipientsSearch(i *NpsSurveyRecipientsSearchInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/recipients/search.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyResponseCreateInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyResponseCreate(i *NpsSurveyResponseCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyResponseShowInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyResponseShow(i *NpsSurveyResponseShowInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyResponseUpdateInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyResponseUpdate(i *NpsSurveyResponseUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses/%s.json"
	path := fmt.Sprintf(api_path, i.SurveyID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyResponsesInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyResponses(i *NpsSurveyResponsesInput, ro *RequestOptions) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	api_path := "/api/v2/nps/surveys/%s/responses.json"
	path := fmt.Sprintf(api_path, i.SurveyID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type NpsSurveyShowInput struct {
	ID string
}

func (c *Client) NpsSurveyShow(i *NpsSurveyShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/nps/surveys/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) NpsSurveys1Update(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/nps/surveys/1"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) NpsSurveysList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/nps/surveys"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthClientCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/clients"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OauthClientDeleteInput struct {
	ID string
}

func (c *Client) OauthClientDelete(i *OauthClientDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type OauthClientGenerateSecretUpdateInput struct {
	ID string
}

func (c *Client) OauthClientGenerateSecretUpdate(i *OauthClientGenerateSecretUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s/generate_secret.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type OauthClientShowInput struct {
	ID string
}

func (c *Client) OauthClientShow(i *OauthClientShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OauthClientUpdateInput struct {
	ID string
}

func (c *Client) OauthClientUpdate(i *OauthClientUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/clients/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthClientsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/clients"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthGlobalClientsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/global_clients.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthTokenCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/tokens.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OauthTokenDeleteInput struct {
	ID string
}

func (c *Client) OauthTokenDelete(i *OauthTokenDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/tokens/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type OauthTokenShowInput struct {
	ID string
}

func (c *Client) OauthTokenShow(i *OauthTokenShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/oauth/tokens/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthTokensCurrentDelete(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/tokens/current.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthTokensCurrentList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/tokens/current.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OauthTokensList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/oauth/tokens"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OrganizationDeleteInput struct {
	ID string
}

func (c *Client) OrganizationDelete(i *OrganizationDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationFieldCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_fields.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OrganizationFieldDeleteInput struct {
	ID string
}

func (c *Client) OrganizationFieldDelete(i *OrganizationFieldDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type OrganizationFieldShowInput struct {
	ID string
}

func (c *Client) OrganizationFieldShow(i *OrganizationFieldShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationFieldUpdateInput struct {
	ID string
}

func (c *Client) OrganizationFieldUpdate(i *OrganizationFieldUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationFieldsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_fields.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationFieldsReorder(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_fields/reorder.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationMembershipCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_memberships.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OrganizationMembershipDeleteInput struct {
	ID string
}

func (c *Client) OrganizationMembershipDelete(i *OrganizationMembershipDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type OrganizationMembershipShowInput struct {
	ID string
}

func (c *Client) OrganizationMembershipShow(i *OrganizationMembershipShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationMembershipsCreateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_memberships/create_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationMembershipsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_memberships/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationMembershipsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_memberships.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationOrganizationMembershipsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationOrganizationMemberships(i *OrganizationOrganizationMembershipsInput, ro *RequestOptions) ([]byte, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	api_path := "/api/v2/organizations/%s/organization_memberships.json"
	path := fmt.Sprintf(api_path, i.OrganizationID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationRelatedInput struct {
	ID string
}

func (c *Client) OrganizationRelated(i *OrganizationRelatedInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/related.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationRequestsInput struct {
	ID string
}

func (c *Client) OrganizationRequests(i *OrganizationRequestsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/requests.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationShowInput struct {
	ID string
}

func (c *Client) OrganizationShow(i *OrganizationShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationSubscriptionCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_subscriptions.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OrganizationSubscriptionDeleteInput struct {
	ID string
}

func (c *Client) OrganizationSubscriptionDelete(i *OrganizationSubscriptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type OrganizationSubscriptionShowInput struct {
	ID string
}

func (c *Client) OrganizationSubscriptionShow(i *OrganizationSubscriptionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organization_subscriptions/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationSubscriptionsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationSubscriptions(i *OrganizationSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	api_path := "/api/v2/organizations/%s/subscriptions.json"
	path := fmt.Sprintf(api_path, i.OrganizationID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationSubscriptionsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organization_subscriptions.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationTagCreateInput struct {
	ID string
}

func (c *Client) OrganizationTagCreate(i *OrganizationTagCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type OrganizationTagsInput struct {
	ID string
}

func (c *Client) OrganizationTags(i *OrganizationTagsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationTagsDeleteInput struct {
	ID string
}

func (c *Client) OrganizationTagsDelete(i *OrganizationTagsDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type OrganizationTagsUpdateInput struct {
	ID string
}

func (c *Client) OrganizationTagsUpdate(i *OrganizationTagsUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type OrganizationTicketsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationTickets(i *OrganizationTicketsInput, ro *RequestOptions) ([]byte, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	api_path := "/api/v2/organizations/%s/tickets.json"
	path := fmt.Sprintf(api_path, i.OrganizationID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type OrganizationUpdateInput struct {
	ID string
}

func (c *Client) OrganizationUpdate(i *OrganizationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type OrganizationUsersInput struct {
	ID string
}

func (c *Client) OrganizationUsers(i *OrganizationUsersInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/organizations/%s/users.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsAutocomplete(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/autocomplete.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsCreateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/create_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsCreateOrUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/create_or_update.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsShowMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/show_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) OrganizationsUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/organizations/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ProblemsAutocomplete(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/problems/autocomplete.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ProblemsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/problems.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) PushNotificationDevicesDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/push_notification_devices/destroy_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RecipientAddressCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/recipient_addresses.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type RecipientAddressDeleteInput struct {
	ID string
}

func (c *Client) RecipientAddressDelete(i *RecipientAddressDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type RecipientAddressShowInput struct {
	ID string
}

func (c *Client) RecipientAddressShow(i *RecipientAddressShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type RecipientAddressUpdateInput struct {
	ID string
}

func (c *Client) RecipientAddressUpdate(i *RecipientAddressUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type RecipientAddressVerifyInput struct {
	ID string
}

func (c *Client) RecipientAddressVerify(i *RecipientAddressVerifyInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/recipient_addresses/%s/verify.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RecipientAddressesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/recipient_addresses.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type RequestCommentShowInput struct {
	RequestID string
	ID        string
}

func (c *Client) RequestCommentShow(i *RequestCommentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.RequestID == "" {
		return nil, errors.New("Missing required field 'RequestID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s/comments/%s.json"
	path := fmt.Sprintf(api_path, i.RequestID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type RequestCommentsInput struct {
	ID string
}

func (c *Client) RequestComments(i *RequestCommentsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s/comments.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RequestCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/requests.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type RequestShowInput struct {
	ID string
}

func (c *Client) RequestShow(i *RequestShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type RequestUpdateInput struct {
	ID string
}

func (c *Client) RequestUpdate(i *RequestUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/requests/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RequestsCcdList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/requests/ccd.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RequestsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/requests.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RequestsOpenList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/requests/open.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RequestsSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/requests/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) RequestsSolvedList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/requests/solved.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ResourceCollectionCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/resource_collections.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ResourceCollectionDeleteInput struct {
	ID string
}

func (c *Client) ResourceCollectionDelete(i *ResourceCollectionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/resource_collections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type ResourceCollectionShowInput struct {
	ID string
}

func (c *Client) ResourceCollectionShow(i *ResourceCollectionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/resource_collections/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ResourceCollectionsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/resource_collections.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ResourceCollectionsUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/resource_collections.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type SatisfactionRatingShowInput struct {
	ID string
}

func (c *Client) SatisfactionRatingShow(i *SatisfactionRatingShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/satisfaction_ratings/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SatisfactionRatingsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/satisfaction_ratings.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type SatisfactionReasonShowInput struct {
	ID string
}

func (c *Client) SatisfactionReasonShow(i *SatisfactionReasonShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/satisfaction_reasons/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SatisfactionReasonsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/satisfaction_reasons.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) Search(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SessionsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/sessions.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SharingAgreementCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/sharing_agreements.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type SharingAgreementDeleteInput struct {
	ID string
}

func (c *Client) SharingAgreementDelete(i *SharingAgreementDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/sharing_agreements/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type SharingAgreementShowInput struct {
	ID string
}

func (c *Client) SharingAgreementShow(i *SharingAgreementShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/sharing_agreements/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type SharingAgreementUpdateInput struct {
	ID string
}

func (c *Client) SharingAgreementUpdate(i *SharingAgreementUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/sharing_agreements/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SharingAgreementsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/sharing_agreements.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ShortcutCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/shortcuts"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ShortcutDeleteInput struct {
	ShortcutName string
}

func (c *Client) ShortcutDelete(i *ShortcutDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	api_path := "/api/v2/shortcuts/%s"
	path := fmt.Sprintf(api_path, i.ShortcutName)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type ShortcutShowInput struct {
	ShortcutName string
}

func (c *Client) ShortcutShow(i *ShortcutShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	api_path := "/api/v2/shortcuts/%s"
	path := fmt.Sprintf(api_path, i.ShortcutName)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ShortcutUpdateInput struct {
	ShortcutName string
}

func (c *Client) ShortcutUpdate(i *ShortcutUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	api_path := "/api/v2/shortcuts/%s"
	path := fmt.Sprintf(api_path, i.ShortcutName)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ShortcutsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/shortcuts"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SkipCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/skips.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SkipsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/skips.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SlasPoliciesDefinitionsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/slas/policies/definitions.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SlasPoliciesList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/slas/policies"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SlasPoliciesReorder(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/slas/policies/reorder.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SlasPolicyCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/slas/policies"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type SlasPolicyDeleteInput struct {
	ID string
}

func (c *Client) SlasPolicyDelete(i *SlasPolicyDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/slas/policies/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type SlasPolicyShowInput struct {
	ID string
}

func (c *Client) SlasPolicyShow(i *SlasPolicyShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/slas/policies/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type SlasPolicyUpdateInput struct {
	ID string
}

func (c *Client) SlasPolicyUpdate(i *SlasPolicyUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/slas/policies/%s"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type StreamAgentShowInput struct {
	MetricKey string
}

func (c *Client) StreamAgentShow(i *StreamAgentShowInput, ro *RequestOptions) ([]byte, error) {
	if i.MetricKey == "" {
		return nil, errors.New("Missing required field 'MetricKey'")
	}

	api_path := "/stream/agents/%s"
	path := fmt.Sprintf(api_path, i.MetricKey)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) StreamAgentsAgentsOnlineList(ro *RequestOptions) ([]byte, error) {
	path := "/stream/agents/agents_online"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) StreamAgentsList(ro *RequestOptions) ([]byte, error) {
	path := "/stream/agents"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type StreamChatShowInput struct {
	MetricKey string
}

func (c *Client) StreamChatShow(i *StreamChatShowInput, ro *RequestOptions) ([]byte, error) {
	if i.MetricKey == "" {
		return nil, errors.New("Missing required field 'MetricKey'")
	}

	api_path := "/stream/chats/%s"
	path := fmt.Sprintf(api_path, i.MetricKey)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) StreamChatsList(ro *RequestOptions) ([]byte, error) {
	path := "/stream/chats"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) StreamChatsMissedChatsList(ro *RequestOptions) ([]byte, error) {
	path := "/stream/chats/missed_chats"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type SuspendedTicketDeleteInput struct {
	ID string
}

func (c *Client) SuspendedTicketDelete(i *SuspendedTicketDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/suspended_tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type SuspendedTicketRecoverInput struct {
	ID string
}

func (c *Client) SuspendedTicketRecover(i *SuspendedTicketRecoverInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/suspended_tickets/%s/recover.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type SuspendedTicketShowInput struct {
	ID string
}

func (c *Client) SuspendedTicketShow(i *SuspendedTicketShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/suspended_tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SuspendedTicketsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/suspended_tickets/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SuspendedTicketsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/suspended_tickets.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) SuspendedTicketsRecoverMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/suspended_tickets/recover_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TagsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tags.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TargetCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/targets.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TargetDeleteInput struct {
	ID string
}

func (c *Client) TargetDelete(i *TargetDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/targets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TargetShowInput struct {
	ID string
}

func (c *Client) TargetShow(i *TargetShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/targets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TargetUpdateInput struct {
	ID string
}

func (c *Client) TargetUpdate(i *TargetUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/targets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TargetsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/targets.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketAuditMakePrivateInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketAuditMakePrivate(i *TicketAuditMakePrivateInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/audits/%s/make_private.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TicketAuditShowInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketAuditShow(i *TicketAuditShowInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/audits/%s.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketAuditsInput struct {
	TicketID string
}

func (c *Client) TicketAudits(i *TicketAuditsInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/audits.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketCollaboratorsInput struct {
	ID string
}

func (c *Client) TicketCollaborators(i *TicketCollaboratorsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/collaborators.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketCommentAttachmentRedactInput struct {
	TicketID     string
	CommentID    string
	AttachmentID string
}

func (c *Client) TicketCommentAttachmentRedact(i *TicketCommentAttachmentRedactInput, ro *RequestOptions) ([]byte, error) {
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
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TicketCommentMakePrivateInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketCommentMakePrivate(i *TicketCommentMakePrivateInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/comments/%s/make_private.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TicketCommentRedactInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketCommentRedact(i *TicketCommentRedactInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/comments/%s/redact.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TicketCommentsInput struct {
	TicketID string
}

func (c *Client) TicketComments(i *TicketCommentsInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/comments.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketDeleteInput struct {
	ID string
}

func (c *Client) TicketDelete(i *TicketDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketFieldCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_fields.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldDeleteInput struct {
	ID string
}

func (c *Client) TicketFieldDelete(i *TicketFieldDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldOptionCreateInput struct {
	FieldID string
	ID      string
}

func (c *Client) TicketFieldOptionCreate(i *TicketFieldOptionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldOptionDeleteInput struct {
	FieldID string
	ID      string
}

func (c *Client) TicketFieldOptionDelete(i *TicketFieldOptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldOptionShowInput struct {
	FieldID string
	ID      string
}

func (c *Client) TicketFieldOptionShow(i *TicketFieldOptionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldOptionsInput struct {
	FieldID string
}

func (c *Client) TicketFieldOptions(i *TicketFieldOptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	api_path := "/api/v2/ticket_fields/%s/options.json"
	path := fmt.Sprintf(api_path, i.FieldID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldShowInput struct {
	ID string
}

func (c *Client) TicketFieldShow(i *TicketFieldShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketFieldUpdateInput struct {
	ID string
}

func (c *Client) TicketFieldUpdate(i *TicketFieldUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketFieldsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_fields.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketFormCloneInput struct {
	ID string
}

func (c *Client) TicketFormClone(i *TicketFormCloneInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s/clone.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketFormCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_forms.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketFormDeleteInput struct {
	ID string
}

func (c *Client) TicketFormDelete(i *TicketFormDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TicketFormShowInput struct {
	ID string
}

func (c *Client) TicketFormShow(i *TicketFormShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketFormUpdateInput struct {
	ID string
}

func (c *Client) TicketFormUpdate(i *TicketFormUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/ticket_forms/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketFormsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_forms.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketFormsReorder(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_forms/reorder.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketFormsShowMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_forms/show_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketIncidentsInput struct {
	ID string
}

func (c *Client) TicketIncidents(i *TicketIncidentsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/incidents.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketMacroApplyInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketMacroApply(i *TicketMacroApplyInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/macros/%s/apply.json"
	path := fmt.Sprintf(api_path, i.TicketID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketMarkAsSpamInput struct {
	ID string
}

func (c *Client) TicketMarkAsSpam(i *TicketMarkAsSpamInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/mark_as_spam.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TicketMergeInput struct {
	ID string
}

func (c *Client) TicketMerge(i *TicketMergeInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/merge.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketMetricShowInput struct {
	TicketMetricID string
}

func (c *Client) TicketMetricShow(i *TicketMetricShowInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketMetricID == "" {
		return nil, errors.New("Missing required field 'TicketMetricID'")
	}

	api_path := "/api/v2/ticket_metrics/%s.json"
	path := fmt.Sprintf(api_path, i.TicketMetricID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketMetricsInput struct {
	TicketID string
}

func (c *Client) TicketMetrics(i *TicketMetricsInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/metrics.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketMetricsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/ticket_metrics.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketRelatedInput struct {
	ID string
}

func (c *Client) TicketRelated(i *TicketRelatedInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/related.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketSatisfactionRatingCreateInput struct {
	TicketID string
}

func (c *Client) TicketSatisfactionRatingCreate(i *TicketSatisfactionRatingCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/satisfaction_rating.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketShowInput struct {
	ID string
}

func (c *Client) TicketShow(i *TicketShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketSkipsInput struct {
	TicketID string
}

func (c *Client) TicketSkips(i *TicketSkipsInput, ro *RequestOptions) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	api_path := "/api/v2/tickets/%s/skips.json"
	path := fmt.Sprintf(api_path, i.TicketID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketTagCreateInput struct {
	ID string
}

func (c *Client) TicketTagCreate(i *TicketTagCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TicketTagsInput struct {
	ID string
}

func (c *Client) TicketTags(i *TicketTagsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TicketTagsDeleteInput struct {
	ID string
}

func (c *Client) TicketTagsDelete(i *TicketTagsDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TicketTagsUpdateInput struct {
	ID string
}

func (c *Client) TicketTagsUpdate(i *TicketTagsUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TicketUpdateInput struct {
	ID string
}

func (c *Client) TicketUpdate(i *TicketUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/tickets/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsCreateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets/create_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsMarkManyAsSpam(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets/mark_many_as_spam.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsRecentList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets/recent.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsShowMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets/show_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TicketsUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/tickets/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TopicTagCreateInput struct {
	ID string
}

func (c *Client) TopicTagCreate(i *TopicTagCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TopicTagsInput struct {
	ID string
}

func (c *Client) TopicTags(i *TopicTagsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TopicTagsDeleteInput struct {
	ID string
}

func (c *Client) TopicTagsDelete(i *TopicTagsDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TopicTagsUpdateInput struct {
	ID string
}

func (c *Client) TopicTagsUpdate(i *TopicTagsUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/topics/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggerCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type TriggerDeleteByIDInput struct {
	ID string
}

func (c *Client) TriggerDeleteByID(i *TriggerDeleteByIDInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/triggers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TriggerDeleteByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerDeleteByTriggerName(i *TriggerDeleteByTriggerNameInput, ro *RequestOptions) ([]byte, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	api_path := "/api/v2/triggers/%s"
	path := fmt.Sprintf(api_path, i.TriggerName)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type TriggerShowByIDInput struct {
	ID string
}

func (c *Client) TriggerShowByID(i *TriggerShowByIDInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/triggers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TriggerShowByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerShowByTriggerName(i *TriggerShowByTriggerNameInput, ro *RequestOptions) ([]byte, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	api_path := "/api/v2/triggers/%s"
	path := fmt.Sprintf(api_path, i.TriggerName)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type TriggerUpdateByIDInput struct {
	ID string
}

func (c *Client) TriggerUpdateByID(i *TriggerUpdateByIDInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/triggers/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type TriggerUpdateByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerUpdateByTriggerName(i *TriggerUpdateByTriggerNameInput, ro *RequestOptions) ([]byte, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	api_path := "/api/v2/triggers/%s"
	path := fmt.Sprintf(api_path, i.TriggerName)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggersActiveList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers/active.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggersDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggersList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggersReorder(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers/reorder.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggersSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) TriggersUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/triggers/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UploadCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/uploads.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UploadDeleteInput struct {
	Token string
}

func (c *Client) UploadDelete(i *UploadDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.Token == "" {
		return nil, errors.New("Missing required field 'Token'")
	}

	api_path := "/api/v2/uploads/%s.json"
	path := fmt.Sprintf(api_path, i.Token)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UserCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserDeleteInput struct {
	ID string
}

func (c *Client) UserDelete(i *UserDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UserFieldCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/user_fields.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserFieldDeleteInput struct {
	ID string
}

func (c *Client) UserFieldDelete(i *UserFieldDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserFieldOptionCreateInput struct {
	FieldID string
}

func (c *Client) UserFieldOptionCreate(i *UserFieldOptionCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	api_path := "/api/v2/user_fields/%s/options.json"
	path := fmt.Sprintf(api_path, i.FieldID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserFieldOptionDeleteInput struct {
	FieldID string
	ID      string
}

func (c *Client) UserFieldOptionDelete(i *UserFieldOptionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserFieldOptionShowInput struct {
	FieldID string
	ID      string
}

func (c *Client) UserFieldOptionShow(i *UserFieldOptionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s/options/%s.json"
	path := fmt.Sprintf(api_path, i.FieldID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserFieldOptionsInput struct {
	FieldID string
}

func (c *Client) UserFieldOptions(i *UserFieldOptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	api_path := "/api/v2/user_fields/%s/options.json"
	path := fmt.Sprintf(api_path, i.FieldID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserFieldShowInput struct {
	ID string
}

func (c *Client) UserFieldShow(i *UserFieldShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserFieldUpdateInput struct {
	ID string
}

func (c *Client) UserFieldUpdate(i *UserFieldUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/user_fields/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UserFieldsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/user_fields.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UserFieldsReorder(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/user_fields/reorder.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserGroupMembershipCreateInput struct {
	UserID string
}

func (c *Client) UserGroupMembershipCreate(i *UserGroupMembershipCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/group_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserGroupMembershipDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserGroupMembershipDelete(i *UserGroupMembershipDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserGroupMembershipMakeDefaultInput struct {
	UserID       string
	MembershipID string
}

func (c *Client) UserGroupMembershipMakeDefault(i *UserGroupMembershipMakeDefaultInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.MembershipID == "" {
		return nil, errors.New("Missing required field 'MembershipID'")
	}

	api_path := "/api/v2/users/%s/group_memberships/%s/make_default.json"
	path := fmt.Sprintf(api_path, i.UserID, i.MembershipID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserGroupMembershipShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserGroupMembershipShow(i *UserGroupMembershipShowInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/group_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserGroupMembershipsInput struct {
	UserID string
}

func (c *Client) UserGroupMemberships(i *UserGroupMembershipsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/group_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserGroupsInput struct {
	UserID string
}

func (c *Client) UserGroups(i *UserGroupsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/groups.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserIdentitiesInput struct {
	UserID string
}

func (c *Client) UserIdentities(i *UserIdentitiesInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/identities.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityCreateInput struct {
	UserID string
}

func (c *Client) UserIdentityCreate(i *UserIdentityCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/identities.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityDelete(i *UserIdentityDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityMakePrimaryInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityMakePrimary(i *UserIdentityMakePrimaryInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s/make_primary"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityRequestVerificationUpdateInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityRequestVerificationUpdate(i *UserIdentityRequestVerificationUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s/request_verification.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityShow(i *UserIdentityShowInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityUpdateInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityUpdate(i *UserIdentityUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserIdentityVerifyInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityVerify(i *UserIdentityVerifyInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/identities/%s/verify"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserMergeInput struct {
	ID string
}

func (c *Client) UserMerge(i *UserMergeInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/merge.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationMembershipCreateInput struct {
	UserID string
}

func (c *Client) UserOrganizationMembershipCreate(i *UserOrganizationMembershipCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationMembershipDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserOrganizationMembershipDelete(i *UserOrganizationMembershipDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationMembershipMakeDefaultInput struct {
	ID           string
	MembershipID string
}

func (c *Client) UserOrganizationMembershipMakeDefault(i *UserOrganizationMembershipMakeDefaultInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	if i.MembershipID == "" {
		return nil, errors.New("Missing required field 'MembershipID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships/%s/make_default.json"
	path := fmt.Sprintf(api_path, i.ID, i.MembershipID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationMembershipShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserOrganizationMembershipShow(i *UserOrganizationMembershipShowInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationMembershipsInput struct {
	UserID string
}

func (c *Client) UserOrganizationMemberships(i *UserOrganizationMembershipsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organization_memberships.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationSubscriptionsInput struct {
	UserID string
}

func (c *Client) UserOrganizationSubscriptions(i *UserOrganizationSubscriptionsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organization_subscriptions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserOrganizationsInput struct {
	UserID string
}

func (c *Client) UserOrganizations(i *UserOrganizationsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/organizations.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserPasswordCreateInput struct {
	UserID string
}

func (c *Client) UserPasswordCreate(i *UserPasswordCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/password.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserPasswordRequirementsInput struct {
	UserID string
}

func (c *Client) UserPasswordRequirements(i *UserPasswordRequirementsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/password/requirements.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserPasswordUpdateInput struct {
	UserID string
}

func (c *Client) UserPasswordUpdate(i *UserPasswordUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/password.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserRelatedInput struct {
	ID string
}

func (c *Client) UserRelated(i *UserRelatedInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/related.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserRequestsInput struct {
	ID string
}

func (c *Client) UserRequests(i *UserRequestsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/requests.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserSessionDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserSessionDelete(i *UserSessionDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/sessions/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserSessionShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserSessionShow(i *UserSessionShowInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/sessions/%s.json"
	path := fmt.Sprintf(api_path, i.UserID, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserSessionsInput struct {
	UserID string
}

func (c *Client) UserSessions(i *UserSessionsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/sessions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserSessionsDeleteInput struct {
	UserID string
}

func (c *Client) UserSessionsDelete(i *UserSessionsDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/sessions.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserShowInput struct {
	ID string
}

func (c *Client) UserShow(i *UserShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserSkipsInput struct {
	UserID string
}

func (c *Client) UserSkips(i *UserSkipsInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/skips.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserTagCreateInput struct {
	ID string
}

func (c *Client) UserTagCreate(i *UserTagCreateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type UserTagsInput struct {
	ID string
}

func (c *Client) UserTags(i *UserTagsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserTagsDeleteInput struct {
	ID string
}

func (c *Client) UserTagsDelete(i *UserTagsDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type UserTagsUpdateInput struct {
	ID string
}

func (c *Client) UserTagsUpdate(i *UserTagsUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s/tags.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type UserTicketsAssignedInput struct {
	UserID string
}

func (c *Client) UserTicketsAssigned(i *UserTicketsAssignedInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/tickets/assigned.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserTicketsCcdInput struct {
	UserID string
}

func (c *Client) UserTicketsCcd(i *UserTicketsCcdInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/tickets/ccd.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserTicketsRequestedInput struct {
	UserID string
}

func (c *Client) UserTicketsRequested(i *UserTicketsRequestedInput, ro *RequestOptions) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	api_path := "/api/v2/users/%s/tickets/requested.json"
	path := fmt.Sprintf(api_path, i.UserID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type UserUpdateInput struct {
	ID string
}

func (c *Client) UserUpdate(i *UserUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/users/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersAutocomplete(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/autocomplete.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersCreateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/create_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersCreateOrUpdate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/create_or_update.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersCreateOrUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/create_or_update_many.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersMe(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/me.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersMeLogout(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/me/logout.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersMeMerge(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/me/merge.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersMeOauthClients(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/me/oauth/clients.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersMeSession(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/me/session.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersRequestCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/request_create.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersShowMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/show_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) UsersUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/users/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}

type ViewCountInput struct {
	ID string
}

func (c *Client) ViewCount(i *ViewCountInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/count.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type ViewDeleteInput struct {
	ID string
}

func (c *Client) ViewDelete(i *ViewDeleteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}

type ViewExecuteInput struct {
	ID string
}

func (c *Client) ViewExecute(i *ViewExecuteInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/execute.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ViewExportInput struct {
	ID string
}

func (c *Client) ViewExport(i *ViewExportInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/export.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ViewShowInput struct {
	ID string
}

func (c *Client) ViewShow(i *ViewShowInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ViewTicketsInput struct {
	ID string
}

func (c *Client) ViewTickets(i *ViewTicketsInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s/tickets.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type ViewUpdateInput struct {
	ID string
}

func (c *Client) ViewUpdate(i *ViewUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	api_path := "/api/v2/views/%s.json"
	path := fmt.Sprintf(api_path, i.ID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsActiveList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/active.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsCompact(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/compact.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsCountManyList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/count_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsDestroyMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/destroy_many.json"
	resp, err := c.Request("DELETE", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsList(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsPreview(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/preview.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsPreviewCount(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/preview/count.json"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsSearch(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/search.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsShowMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/show_many.json"
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) ViewsUpdateMany(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/views/update_many.json"
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
func (c *Client) VisitorCreate(ro *RequestOptions) ([]byte, error) {
	path := "/api/v2/visitors"
	resp, err := c.Request("POST", path, ro)
	return handle(resp, ro, err)
}

type VisitorShowInput struct {
	VisitorID string
}

func (c *Client) VisitorShow(i *VisitorShowInput, ro *RequestOptions) ([]byte, error) {
	if i.VisitorID == "" {
		return nil, errors.New("Missing required field 'VisitorID'")
	}

	api_path := "/api/v2/visitors/%s"
	path := fmt.Sprintf(api_path, i.VisitorID)
	resp, err := c.Request("GET", path, ro)
	return handle(resp, ro, err)
}

type VisitorUpdateInput struct {
	VisitorID string
}

func (c *Client) VisitorUpdate(i *VisitorUpdateInput, ro *RequestOptions) ([]byte, error) {
	if i.VisitorID == "" {
		return nil, errors.New("Missing required field 'VisitorID'")
	}

	api_path := "/api/v2/visitors/%s"
	path := fmt.Sprintf(api_path, i.VisitorID)
	resp, err := c.Request("PUT", path, ro)
	return handle(resp, ro, err)
}
