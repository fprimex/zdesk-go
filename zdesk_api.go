package zdesk

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type AccountListInput struct {
}

func (c *Client) AccountList(i *AccountListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/account")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AccountSettingsListInput struct {
}

func (c *Client) AccountSettingsList(i *AccountSettingsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/account/settings.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AccountSettingsUpdateInput struct {
}

func (c *Client) AccountSettingsUpdate(i *AccountSettingsUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/account/settings.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AccountUpdateInput struct {
}

func (c *Client) AccountUpdate(i *AccountUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/account")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ActivitiesListInput struct {
}

func (c *Client) ActivitiesList(i *ActivitiesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/activities.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ActivityShowInput struct {
	ActivityID string
}

func (c *Client) ActivityShow(i *ActivityShowInput) ([]byte, error) {
	if i.ActivityID == "" {
		return nil, errors.New("Missing required field 'ActivityID'")
	}

	path := fmt.Sprintf("/api/v2/activities/%s.json", i.ActivityID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentCreateInput struct {
}

func (c *Client) AgentCreate(i *AgentCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/agents")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentDeleteInput struct {
	AgentID string
}

func (c *Client) AgentDelete(i *AgentDeleteInput) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	path := fmt.Sprintf("/api/v2/agents/%s", i.AgentID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentShowInput struct {
	AgentID string
}

func (c *Client) AgentShow(i *AgentShowInput) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	path := fmt.Sprintf("/api/v2/agents/%s", i.AgentID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentUpdateInput struct {
	AgentID string
}

func (c *Client) AgentUpdate(i *AgentUpdateInput) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	path := fmt.Sprintf("/api/v2/agents/%s", i.AgentID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentsEmailShowInput struct {
	EmailID string
}

func (c *Client) AgentsEmailShow(i *AgentsEmailShowInput) ([]byte, error) {
	if i.EmailID == "" {
		return nil, errors.New("Missing required field 'EmailID'")
	}

	path := fmt.Sprintf("/api/v2/agents/email/%s", i.EmailID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentsListInput struct {
}

func (c *Client) AgentsList(i *AgentsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/agents")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentsMeInput struct {
}

func (c *Client) AgentsMe(i *AgentsMeInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/agents/me")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AgentsMeUpdateInput struct {
}

func (c *Client) AgentsMeUpdate(i *AgentsMeUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/agents/me")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AnyChannelPushCreateInput struct {
}

func (c *Client) AnyChannelPushCreate(i *AnyChannelPushCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/any_channel/push")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppCreateInput struct {
}

func (c *Client) AppCreate(i *AppCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppDeleteInput struct {
	ID string
}

func (c *Client) AppDelete(i *AppDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppPublicKeyInput struct {
	ID string
}

func (c *Client) AppPublicKey(i *AppPublicKeyInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/%s/public_key.pem", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppShowInput struct {
	ID string
}

func (c *Client) AppShow(i *AppShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppUpdateInput struct {
	ID string
}

func (c *Client) AppUpdate(i *AppUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationCreateInput struct {
}

func (c *Client) AppsInstallationCreate(i *AppsInstallationCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/installations.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationDeleteInput struct {
	ID string
}

func (c *Client) AppsInstallationDelete(i *AppsInstallationDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/installations/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationRequirementsInput struct {
	ID string
}

func (c *Client) AppsInstallationRequirements(i *AppsInstallationRequirementsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/installations/%s/requirements.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationShowInput struct {
	ID string
}

func (c *Client) AppsInstallationShow(i *AppsInstallationShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/installations/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationUpdateInput struct {
	ID string
}

func (c *Client) AppsInstallationUpdate(i *AppsInstallationUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/installations/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationsJobStatusShowInput struct {
	ID string
}

func (c *Client) AppsInstallationsJobStatusShow(i *AppsInstallationsJobStatusShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/installations/job_statuses/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsInstallationsListInput struct {
}

func (c *Client) AppsInstallationsList(i *AppsInstallationsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/installations.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsJobStatusShowInput struct {
	ID string
}

func (c *Client) AppsJobStatusShow(i *AppsJobStatusShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/job_statuses/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsListInput struct {
}

func (c *Client) AppsList(i *AppsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsLocationInstallationsListInput struct {
}

func (c *Client) AppsLocationInstallationsList(i *AppsLocationInstallationsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/location_installations.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsLocationInstallationsReorderInput struct {
}

func (c *Client) AppsLocationInstallationsReorder(i *AppsLocationInstallationsReorderInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/location_installations/reorder.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsLocationShowInput struct {
	ID string
}

func (c *Client) AppsLocationShow(i *AppsLocationShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/apps/locations/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsLocationsListInput struct {
}

func (c *Client) AppsLocationsList(i *AppsLocationsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/locations.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsNotifyCreateInput struct {
}

func (c *Client) AppsNotifyCreate(i *AppsNotifyCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/notify.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsOwnedListInput struct {
}

func (c *Client) AppsOwnedList(i *AppsOwnedListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/owned.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AppsUploadCreateInput struct {
}

func (c *Client) AppsUploadCreate(i *AppsUploadCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/apps/uploads.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AttachmentDeleteInput struct {
	ID string
}

func (c *Client) AttachmentDelete(i *AttachmentDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/attachments/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AttachmentShowInput struct {
	ID string
}

func (c *Client) AttachmentShow(i *AttachmentShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/attachments/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AuditLogShowInput struct {
	ID string
}

func (c *Client) AuditLogShow(i *AuditLogShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/audit_logs/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AuditLogsListInput struct {
}

func (c *Client) AuditLogsList(i *AuditLogsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/audit_logs.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutocompleteTagsInput struct {
}

func (c *Client) AutocompleteTags(i *AutocompleteTagsInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/autocomplete/tags.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationCreateInput struct {
}

func (c *Client) AutomationCreate(i *AutomationCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/automations.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationDeleteInput struct {
	ID string
}

func (c *Client) AutomationDelete(i *AutomationDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/automations/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationShowInput struct {
	ID string
}

func (c *Client) AutomationShow(i *AutomationShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/automations/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationUpdateInput struct {
	ID string
}

func (c *Client) AutomationUpdate(i *AutomationUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/automations/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationsActiveListInput struct {
}

func (c *Client) AutomationsActiveList(i *AutomationsActiveListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/automations/active.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationsDestroyManyInput struct {
}

func (c *Client) AutomationsDestroyMany(i *AutomationsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/automations/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationsListInput struct {
}

func (c *Client) AutomationsList(i *AutomationsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/automations.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationsSearchInput struct {
}

func (c *Client) AutomationsSearch(i *AutomationsSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/automations/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type AutomationsUpdateManyInput struct {
}

func (c *Client) AutomationsUpdateMany(i *AutomationsUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/automations/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BanCreateInput struct {
}

func (c *Client) BanCreate(i *BanCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/bans")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BanDeleteInput struct {
	BanID string
}

func (c *Client) BanDelete(i *BanDeleteInput) ([]byte, error) {
	if i.BanID == "" {
		return nil, errors.New("Missing required field 'BanID'")
	}

	path := fmt.Sprintf("/api/v2/ban/%s", i.BanID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BanShowInput struct {
	BanID string
}

func (c *Client) BanShow(i *BanShowInput) ([]byte, error) {
	if i.BanID == "" {
		return nil, errors.New("Missing required field 'BanID'")
	}

	path := fmt.Sprintf("/api/v2/bans/%s", i.BanID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BansIpListInput struct {
}

func (c *Client) BansIpList(i *BansIpListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/bans/ip")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BansListInput struct {
}

func (c *Client) BansList(i *BansListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/bans")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BookmarkCreateInput struct {
}

func (c *Client) BookmarkCreate(i *BookmarkCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/bookmarks.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BookmarkDeleteInput struct {
	ID string
}

func (c *Client) BookmarkDelete(i *BookmarkDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/bookmarks/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BookmarksListInput struct {
}

func (c *Client) BookmarksList(i *BookmarksListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/bookmarks.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandCheckHostMappingInput struct {
	ID string
}

func (c *Client) BrandCheckHostMapping(i *BrandCheckHostMappingInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/brands/%s/check_host_mapping.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandCreateInput struct {
}

func (c *Client) BrandCreate(i *BrandCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/brands.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandDeleteInput struct {
	ID string
}

func (c *Client) BrandDelete(i *BrandDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/brands/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandShowInput struct {
	ID string
}

func (c *Client) BrandShow(i *BrandShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/brands/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandUpdateInput struct {
	ID string
}

func (c *Client) BrandUpdate(i *BrandUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/brands/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandsCheckHostMappingListInput struct {
}

func (c *Client) BrandsCheckHostMappingList(i *BrandsCheckHostMappingListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/brands/check_host_mapping.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BrandsListInput struct {
}

func (c *Client) BrandsList(i *BrandsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/brands.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleCreateInput struct {
}

func (c *Client) BusinessHoursScheduleCreate(i *BusinessHoursScheduleCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/business_hours/schedules.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleDeleteInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleDelete(i *BusinessHoursScheduleDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleHolidayCreateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleHolidayCreate(i *BusinessHoursScheduleHolidayCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s/holidays.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleHolidayDeleteInput struct {
	ScheduleID string
	ID         string
}

func (c *Client) BusinessHoursScheduleHolidayDelete(i *BusinessHoursScheduleHolidayDeleteInput) ([]byte, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s/holidays/%s.json", i.ScheduleID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleHolidayShowInput struct {
	ScheduleID string
	ID         string
}

func (c *Client) BusinessHoursScheduleHolidayShow(i *BusinessHoursScheduleHolidayShowInput) ([]byte, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s/holidays/%s.json", i.ScheduleID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleHolidayUpdateInput struct {
	ScheduleID string
	ID         string
}

func (c *Client) BusinessHoursScheduleHolidayUpdate(i *BusinessHoursScheduleHolidayUpdateInput) ([]byte, error) {
	if i.ScheduleID == "" {
		return nil, errors.New("Missing required field 'ScheduleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s/holidays/%s.json", i.ScheduleID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleHolidaysInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleHolidays(i *BusinessHoursScheduleHolidaysInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s/holidays.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleShowInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleShow(i *BusinessHoursScheduleShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleUpdateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleUpdate(i *BusinessHoursScheduleUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursScheduleWorkweekUpdateInput struct {
	ID string
}

func (c *Client) BusinessHoursScheduleWorkweekUpdate(i *BusinessHoursScheduleWorkweekUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/business_hours/schedules/%s/workweek.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type BusinessHoursSchedulesListInput struct {
}

func (c *Client) BusinessHoursSchedulesList(i *BusinessHoursSchedulesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/business_hours/schedules.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsTwitterMonitoredTwitterHandleShowInput struct {
	ID string
}

func (c *Client) ChannelsTwitterMonitoredTwitterHandleShow(i *ChannelsTwitterMonitoredTwitterHandleShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/twitter/monitored_twitter_handles/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsTwitterMonitoredTwitterHandlesListInput struct {
}

func (c *Client) ChannelsTwitterMonitoredTwitterHandlesList(i *ChannelsTwitterMonitoredTwitterHandlesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/twitter/monitored_twitter_handles.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsTwitterTicketCreateInput struct {
}

func (c *Client) ChannelsTwitterTicketCreate(i *ChannelsTwitterTicketCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/twitter/tickets.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsTwitterTicketStatusesInput struct {
	ID string
}

func (c *Client) ChannelsTwitterTicketStatuses(i *ChannelsTwitterTicketStatusesInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/twitter/tickets/%s/statuses.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceAgentTicketDisplayCreateInput struct {
	AgentID  string
	TicketID string
}

func (c *Client) ChannelsVoiceAgentTicketDisplayCreate(i *ChannelsVoiceAgentTicketDisplayCreateInput) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/agents/%s/tickets/%s/display.json", i.AgentID, i.TicketID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceAgentUserDisplayCreateInput struct {
	AgentID string
	UserID  string
}

func (c *Client) ChannelsVoiceAgentUserDisplayCreate(i *ChannelsVoiceAgentUserDisplayCreateInput) ([]byte, error) {
	if i.AgentID == "" {
		return nil, errors.New("Missing required field 'AgentID'")
	}

	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/agents/%s/users/%s/display.json", i.AgentID, i.UserID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceAvailabilityShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceAvailabilityShow(i *ChannelsVoiceAvailabilityShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/availabilities/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceAvailabilityUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoiceAvailabilityUpdate(i *ChannelsVoiceAvailabilityUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/availabilities/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingCategoriesListInput struct {
}

func (c *Client) ChannelsVoiceGreetingCategoriesList(i *ChannelsVoiceGreetingCategoriesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/greeting_categories.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingCategoryShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingCategoryShow(i *ChannelsVoiceGreetingCategoryShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/greeting_categories/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingCreateInput struct {
}

func (c *Client) ChannelsVoiceGreetingCreate(i *ChannelsVoiceGreetingCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/greetings.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingDeleteInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingDelete(i *ChannelsVoiceGreetingDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/greetings/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingRecordingInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingRecording(i *ChannelsVoiceGreetingRecordingInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/greetings/%s/recording.mp3", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingShowInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingShow(i *ChannelsVoiceGreetingShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/greetings/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoiceGreetingUpdate(i *ChannelsVoiceGreetingUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/greetings/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceGreetingsListInput struct {
}

func (c *Client) ChannelsVoiceGreetingsList(i *ChannelsVoiceGreetingsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/greetings.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoicePhoneNumberCreateInput struct {
}

func (c *Client) ChannelsVoicePhoneNumberCreate(i *ChannelsVoicePhoneNumberCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/phone_numbers.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoicePhoneNumberShowInput struct {
	ID string
}

func (c *Client) ChannelsVoicePhoneNumberShow(i *ChannelsVoicePhoneNumberShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/phone_numbers/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoicePhoneNumberUpdateInput struct {
	ID string
}

func (c *Client) ChannelsVoicePhoneNumberUpdate(i *ChannelsVoicePhoneNumberUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/channels/voice/phone_numbers/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoicePhoneNumbersDeleteInput struct {
}

func (c *Client) ChannelsVoicePhoneNumbersDelete(i *ChannelsVoicePhoneNumbersDeleteInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/phone_numbers.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoicePhoneNumbersListInput struct {
}

func (c *Client) ChannelsVoicePhoneNumbersList(i *ChannelsVoicePhoneNumbersListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/phone_numbers.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoicePhoneNumbersSearchInput struct {
}

func (c *Client) ChannelsVoicePhoneNumbersSearch(i *ChannelsVoicePhoneNumbersSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/phone_numbers/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceStatsAgentsActivityListInput struct {
}

func (c *Client) ChannelsVoiceStatsAgentsActivityList(i *ChannelsVoiceStatsAgentsActivityListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/stats/agents_activity.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceStatsCurrentQueueActivityListInput struct {
}

func (c *Client) ChannelsVoiceStatsCurrentQueueActivityList(i *ChannelsVoiceStatsCurrentQueueActivityListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/stats/current_queue_activity.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceStatsHistoricalQueueActivityListInput struct {
}

func (c *Client) ChannelsVoiceStatsHistoricalQueueActivityList(i *ChannelsVoiceStatsHistoricalQueueActivityListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/stats/historical_queue_activity.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChannelsVoiceTicketCreateInput struct {
}

func (c *Client) ChannelsVoiceTicketCreate(i *ChannelsVoiceTicketCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/channels/voice/tickets.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChatCreateInput struct {
}

func (c *Client) ChatCreate(i *ChatCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/chats")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChatDeleteInput struct {
	ChatID string
}

func (c *Client) ChatDelete(i *ChatDeleteInput) ([]byte, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	path := fmt.Sprintf("/api/v2/chats/%s", i.ChatID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChatShowInput struct {
	ChatID string
}

func (c *Client) ChatShow(i *ChatShowInput) ([]byte, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	path := fmt.Sprintf("/api/v2/chats/%s", i.ChatID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChatUpdateInput struct {
	ChatID string
}

func (c *Client) ChatUpdate(i *ChatUpdateInput) ([]byte, error) {
	if i.ChatID == "" {
		return nil, errors.New("Missing required field 'ChatID'")
	}

	path := fmt.Sprintf("/api/v2/chats/%s", i.ChatID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChatsListInput struct {
}

func (c *Client) ChatsList(i *ChatsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/chats")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ChatsSearchInput struct {
}

func (c *Client) ChatsSearch(i *ChatsSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/chats/search")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentCreateInput struct {
	PostID string
}

func (c *Client) CommunityPostCommentCreate(i *CommunityPostCommentCreateInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments.json", i.PostID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentDeleteInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentDelete(i *CommunityPostCommentDeleteInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments/%s.json", i.PostID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentDownCreateInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentDownCreate(i *CommunityPostCommentDownCreateInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments/%s/down.json", i.PostID, i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentShowInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentShow(i *CommunityPostCommentShowInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments/%s.json", i.PostID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentUpCreateInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentUpCreate(i *CommunityPostCommentUpCreateInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments/%s/up.json", i.PostID, i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentUpdateInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostCommentUpdate(i *CommunityPostCommentUpdateInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments/%s.json", i.PostID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentVotesInput struct {
	PostID    string
	CommentID string
}

func (c *Client) CommunityPostCommentVotes(i *CommunityPostCommentVotesInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments/%s/votes.json", i.PostID, i.CommentID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCommentsInput struct {
	PostID string
}

func (c *Client) CommunityPostComments(i *CommunityPostCommentsInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/comments.json", i.PostID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostCreateInput struct {
}

func (c *Client) CommunityPostCreate(i *CommunityPostCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/community/posts.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostDeleteInput struct {
	ID string
}

func (c *Client) CommunityPostDelete(i *CommunityPostDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostDownCreateInput struct {
	ID string
}

func (c *Client) CommunityPostDownCreate(i *CommunityPostDownCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/down.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostShowInput struct {
	ID string
}

func (c *Client) CommunityPostShow(i *CommunityPostShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostSubscriptionCreateInput struct {
	PostID string
}

func (c *Client) CommunityPostSubscriptionCreate(i *CommunityPostSubscriptionCreateInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/subscriptions.json", i.PostID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostSubscriptionDeleteInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostSubscriptionDelete(i *CommunityPostSubscriptionDeleteInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/subscriptions/%s.json", i.PostID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostSubscriptionShowInput struct {
	PostID string
	ID     string
}

func (c *Client) CommunityPostSubscriptionShow(i *CommunityPostSubscriptionShowInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/subscriptions/%s.json", i.PostID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostSubscriptionsInput struct {
	PostID string
}

func (c *Client) CommunityPostSubscriptions(i *CommunityPostSubscriptionsInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/subscriptions.json", i.PostID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostUpCreateInput struct {
	ID string
}

func (c *Client) CommunityPostUpCreate(i *CommunityPostUpCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/up.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostUpdateInput struct {
	ID string
}

func (c *Client) CommunityPostUpdate(i *CommunityPostUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostVotesInput struct {
	PostID string
}

func (c *Client) CommunityPostVotes(i *CommunityPostVotesInput) ([]byte, error) {
	if i.PostID == "" {
		return nil, errors.New("Missing required field 'PostID'")
	}

	path := fmt.Sprintf("/api/v2/community/posts/%s/votes.json", i.PostID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityPostsListInput struct {
}

func (c *Client) CommunityPostsList(i *CommunityPostsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/community/posts.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicAccessPolicyInput struct {
	TopicID string
}

func (c *Client) CommunityTopicAccessPolicy(i *CommunityTopicAccessPolicyInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/access_policy.json", i.TopicID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicAccessPolicyUpdateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicAccessPolicyUpdate(i *CommunityTopicAccessPolicyUpdateInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/access_policy.json", i.TopicID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicCreateInput struct {
}

func (c *Client) CommunityTopicCreate(i *CommunityTopicCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/community/topics.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicDeleteInput struct {
	ID string
}

func (c *Client) CommunityTopicDelete(i *CommunityTopicDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicPostsInput struct {
	ID string
}

func (c *Client) CommunityTopicPosts(i *CommunityTopicPostsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/posts.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicShowInput struct {
	ID string
}

func (c *Client) CommunityTopicShow(i *CommunityTopicShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicSubscriptionCreateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptionCreate(i *CommunityTopicSubscriptionCreateInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/subscriptions.json", i.TopicID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicSubscriptionDeleteInput struct {
	TopicID string
	ID      string
}

func (c *Client) CommunityTopicSubscriptionDelete(i *CommunityTopicSubscriptionDeleteInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/subscriptions/%s.json", i.TopicID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicSubscriptionShowInput struct {
	TopicID string
	ID      string
}

func (c *Client) CommunityTopicSubscriptionShow(i *CommunityTopicSubscriptionShowInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/subscriptions/%s.json", i.TopicID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicSubscriptionsInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptions(i *CommunityTopicSubscriptionsInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/subscriptions.json", i.TopicID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicSubscriptionsUpdateInput struct {
	TopicID string
}

func (c *Client) CommunityTopicSubscriptionsUpdate(i *CommunityTopicSubscriptionsUpdateInput) ([]byte, error) {
	if i.TopicID == "" {
		return nil, errors.New("Missing required field 'TopicID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s/subscriptions.json", i.TopicID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicUpdateInput struct {
	ID string
}

func (c *Client) CommunityTopicUpdate(i *CommunityTopicUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/topics/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityTopicsListInput struct {
}

func (c *Client) CommunityTopicsList(i *CommunityTopicsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/community/topics.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityUserCommentsInput struct {
	ID string
}

func (c *Client) CommunityUserComments(i *CommunityUserCommentsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/users/%s/comments.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CommunityUserPostsInput struct {
	ID string
}

func (c *Client) CommunityUserPosts(i *CommunityUserPostsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/community/users/%s/posts.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CustomRolesListInput struct {
}

func (c *Client) CustomRolesList(i *CustomRolesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/custom_roles.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DeletedTicketRestoreUpdateInput struct {
	ID string
}

func (c *Client) DeletedTicketRestoreUpdate(i *DeletedTicketRestoreUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/deleted_tickets/%s/restore.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DeletedTickets1DeleteInput struct {
}

func (c *Client) DeletedTickets1Delete(i *DeletedTickets1DeleteInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/deleted_tickets/1")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DeletedTicketsDestroyManyInput struct {
}

func (c *Client) DeletedTicketsDestroyMany(i *DeletedTicketsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/deleted_tickets/destroy_many")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DeletedTicketsListInput struct {
}

func (c *Client) DeletedTicketsList(i *DeletedTicketsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/deleted_tickets.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DeletedTicketsRestoreManyUpdateInput struct {
}

func (c *Client) DeletedTicketsRestoreManyUpdate(i *DeletedTicketsRestoreManyUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/deleted_tickets/restore_many")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentCreateInput struct {
}

func (c *Client) DepartmentCreate(i *DepartmentCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/departments")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentDeleteInput struct {
	DepartmentID string
}

func (c *Client) DepartmentDelete(i *DepartmentDeleteInput) ([]byte, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	path := fmt.Sprintf("/api/v2/departments/%s", i.DepartmentID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentShowInput struct {
	DepartmentID string
}

func (c *Client) DepartmentShow(i *DepartmentShowInput) ([]byte, error) {
	if i.DepartmentID == "" {
		return nil, errors.New("Missing required field 'DepartmentID'")
	}

	path := fmt.Sprintf("/api/v2/departments/%s", i.DepartmentID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentUpdateInput struct {
	DepartmentIDOrName string
}

func (c *Client) DepartmentUpdate(i *DepartmentUpdateInput) ([]byte, error) {
	if i.DepartmentIDOrName == "" {
		return nil, errors.New("Missing required field 'DepartmentIDOrName'")
	}

	path := fmt.Sprintf("/api/v2/departments/%s", i.DepartmentIDOrName)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentsListInput struct {
}

func (c *Client) DepartmentsList(i *DepartmentsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/departments")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentsNameDeleteInput struct {
	Name string
}

func (c *Client) DepartmentsNameDelete(i *DepartmentsNameDeleteInput) ([]byte, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	path := fmt.Sprintf("/api/v2/departments/name/%s", i.Name)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DepartmentsNameShowInput struct {
	Name string
}

func (c *Client) DepartmentsNameShow(i *DepartmentsNameShowInput) ([]byte, error) {
	if i.Name == "" {
		return nil, errors.New("Missing required field 'Name'")
	}

	path := fmt.Sprintf("/api/v2/departments/name/%s", i.Name)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemCreateInput struct {
}

func (c *Client) DynamicContentItemCreate(i *DynamicContentItemCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/dynamic_content/items.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemDeleteInput struct {
	ID string
}

func (c *Client) DynamicContentItemDelete(i *DynamicContentItemDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemShowInput struct {
	ID string
}

func (c *Client) DynamicContentItemShow(i *DynamicContentItemShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemUpdateInput struct {
	ID string
}

func (c *Client) DynamicContentItemUpdate(i *DynamicContentItemUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantCreateInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantCreate(i *DynamicContentItemVariantCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantDeleteInput struct {
	ItemID string
	ID     string
}

func (c *Client) DynamicContentItemVariantDelete(i *DynamicContentItemVariantDeleteInput) ([]byte, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants/%s.json", i.ItemID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantShowInput struct {
	ItemID string
	ID     string
}

func (c *Client) DynamicContentItemVariantShow(i *DynamicContentItemVariantShowInput) ([]byte, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants/%s.json", i.ItemID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantUpdateInput struct {
	ItemID string
	ID     string
}

func (c *Client) DynamicContentItemVariantUpdate(i *DynamicContentItemVariantUpdateInput) ([]byte, error) {
	if i.ItemID == "" {
		return nil, errors.New("Missing required field 'ItemID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants/%s.json", i.ItemID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantsInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariants(i *DynamicContentItemVariantsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantsCreateManyInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantsCreateMany(i *DynamicContentItemVariantsCreateManyInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants/create_many.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemVariantsUpdateManyInput struct {
	ID string
}

func (c *Client) DynamicContentItemVariantsUpdateMany(i *DynamicContentItemVariantsUpdateManyInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/dynamic_content/items/%s/variants/update_many.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type DynamicContentItemsListInput struct {
}

func (c *Client) DynamicContentItemsList(i *DynamicContentItemsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/dynamic_content/items.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type EndUserIdentityCreateInput struct {
	UserID string
}

func (c *Client) EndUserIdentityCreate(i *EndUserIdentityCreateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/end_users/%s/identities.json", i.UserID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type EndUserIdentityMakePrimaryInput struct {
	UserID string
	ID     string
}

func (c *Client) EndUserIdentityMakePrimary(i *EndUserIdentityMakePrimaryInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/end_users/%s/identities/%s/make_primary", i.UserID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type EndUserShowInput struct {
	ID string
}

func (c *Client) EndUserShow(i *EndUserShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/end_users/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type EndUserUpdateInput struct {
	ID string
}

func (c *Client) EndUserUpdate(i *EndUserUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/end_users/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GoalCreateInput struct {
}

func (c *Client) GoalCreate(i *GoalCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/goals")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GoalDeleteInput struct {
	GoalID string
}

func (c *Client) GoalDelete(i *GoalDeleteInput) ([]byte, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	path := fmt.Sprintf("/api/v2/goals/%s", i.GoalID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GoalShowInput struct {
	GoalID string
}

func (c *Client) GoalShow(i *GoalShowInput) ([]byte, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	path := fmt.Sprintf("/api/v2/goals/%s", i.GoalID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GoalUpdateInput struct {
	GoalID string
}

func (c *Client) GoalUpdate(i *GoalUpdateInput) ([]byte, error) {
	if i.GoalID == "" {
		return nil, errors.New("Missing required field 'GoalID'")
	}

	path := fmt.Sprintf("/api/v2/goals/%s", i.GoalID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GoalsListInput struct {
}

func (c *Client) GoalsList(i *GoalsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/goals")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupCreateInput struct {
}

func (c *Client) GroupCreate(i *GroupCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/groups.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupDeleteInput struct {
	ID string
}

func (c *Client) GroupDelete(i *GroupDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/groups/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipCreateInput struct {
}

func (c *Client) GroupMembershipCreate(i *GroupMembershipCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/group_memberships.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipDeleteInput struct {
	ID string
}

func (c *Client) GroupMembershipDelete(i *GroupMembershipDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/group_memberships/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipShowInput struct {
	ID string
}

func (c *Client) GroupMembershipShow(i *GroupMembershipShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/group_memberships/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipsInput struct {
	GroupID string
}

func (c *Client) GroupMemberships(i *GroupMembershipsInput) ([]byte, error) {
	if i.GroupID == "" {
		return nil, errors.New("Missing required field 'GroupID'")
	}

	path := fmt.Sprintf("/api/v2/groups/%s/memberships.json", i.GroupID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipsAssignableInput struct {
	GroupID string
}

func (c *Client) GroupMembershipsAssignable(i *GroupMembershipsAssignableInput) ([]byte, error) {
	if i.GroupID == "" {
		return nil, errors.New("Missing required field 'GroupID'")
	}

	path := fmt.Sprintf("/api/v2/groups/%s/memberships/assignable.json", i.GroupID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipsAssignableListInput struct {
}

func (c *Client) GroupMembershipsAssignableList(i *GroupMembershipsAssignableListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/group_memberships/assignable.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipsCreateManyInput struct {
}

func (c *Client) GroupMembershipsCreateMany(i *GroupMembershipsCreateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/group_memberships/create_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipsDestroyManyInput struct {
}

func (c *Client) GroupMembershipsDestroyMany(i *GroupMembershipsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/group_memberships/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupMembershipsListInput struct {
}

func (c *Client) GroupMembershipsList(i *GroupMembershipsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/group_memberships.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupShowInput struct {
	ID string
}

func (c *Client) GroupShow(i *GroupShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/groups/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupUpdateInput struct {
	ID string
}

func (c *Client) GroupUpdate(i *GroupUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/groups/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupUsersInput struct {
	ID string
}

func (c *Client) GroupUsers(i *GroupUsersInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/groups/%s/users.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupsAssignableListInput struct {
}

func (c *Client) GroupsAssignableList(i *GroupsAssignableListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/groups/assignable.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type GroupsListInput struct {
}

func (c *Client) GroupsList(i *GroupsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/groups.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleAttachmentCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleAttachmentCreate(i *HelpCenterArticleAttachmentCreateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/attachments.json", i.ArticleID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleAttachmentsInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleAttachments(i *HelpCenterArticleAttachmentsInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/attachments.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleAttachmentsBlockInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleAttachmentsBlock(i *HelpCenterArticleAttachmentsBlockInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/attachments/block.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleAttachmentsInlineInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleAttachmentsInline(i *HelpCenterArticleAttachmentsInlineInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/attachments/inline.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleBulkAttachmentCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleBulkAttachmentCreate(i *HelpCenterArticleBulkAttachmentCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/bulk_attachments.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleCommentCreate(i *HelpCenterArticleCommentCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentDeleteInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentDelete(i *HelpCenterArticleCommentDeleteInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments/%s.json", i.ArticleID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentDownCreateInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentDownCreate(i *HelpCenterArticleCommentDownCreateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments/%s/down.json", i.ArticleID, i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentShowInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentShow(i *HelpCenterArticleCommentShowInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments/%s.json", i.ArticleID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentUpCreateInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentUpCreate(i *HelpCenterArticleCommentUpCreateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments/%s/up.json", i.ArticleID, i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentUpdateInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleCommentUpdate(i *HelpCenterArticleCommentUpdateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments/%s.json", i.ArticleID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentVotesInput struct {
	ArticleID string
	CommentID string
}

func (c *Client) HelpCenterArticleCommentVotes(i *HelpCenterArticleCommentVotesInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments/%s/votes.json", i.ArticleID, i.CommentID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleCommentsInput struct {
	ID string
}

func (c *Client) HelpCenterArticleComments(i *HelpCenterArticleCommentsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/comments.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterArticleDelete(i *HelpCenterArticleDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleDownCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleDownCreate(i *HelpCenterArticleDownCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/down.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleLabelCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleLabelCreate(i *HelpCenterArticleLabelCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/labels.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleLabelDeleteInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleLabelDelete(i *HelpCenterArticleLabelDeleteInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/labels/%s.json", i.ArticleID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleLabelsInput struct {
	ID string
}

func (c *Client) HelpCenterArticleLabels(i *HelpCenterArticleLabelsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/labels.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleShowInput struct {
	Locale string
	ID     string
}

func (c *Client) HelpCenterArticleShow(i *HelpCenterArticleShowInput) ([]byte, error) {
	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/%s/articles/%s.json", i.Locale, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleSourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleSourceLocaleUpdate(i *HelpCenterArticleSourceLocaleUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/source_locale.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleSubscriptionCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleSubscriptionCreate(i *HelpCenterArticleSubscriptionCreateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/subscriptions.json", i.ArticleID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleSubscriptionDeleteInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleSubscriptionDelete(i *HelpCenterArticleSubscriptionDeleteInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/subscriptions/%s.json", i.ArticleID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleSubscriptionShowInput struct {
	ArticleID string
	ID        string
}

func (c *Client) HelpCenterArticleSubscriptionShow(i *HelpCenterArticleSubscriptionShowInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/subscriptions/%s.json", i.ArticleID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleSubscriptionsInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleSubscriptions(i *HelpCenterArticleSubscriptionsInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/subscriptions.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleTranslationCreateInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslationCreate(i *HelpCenterArticleTranslationCreateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/translations.json", i.ArticleID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleTranslationShowInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleTranslationShow(i *HelpCenterArticleTranslationShowInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/translations/%s.json", i.ArticleID, i.Locale)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleTranslationUpdateInput struct {
	ArticleID string
	Locale    string
}

func (c *Client) HelpCenterArticleTranslationUpdate(i *HelpCenterArticleTranslationUpdateInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/translations/%s.json", i.ArticleID, i.Locale)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleTranslationsInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslations(i *HelpCenterArticleTranslationsInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/translations.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleTranslationsMissingInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleTranslationsMissing(i *HelpCenterArticleTranslationsMissingInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/translations/missing.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleUpCreateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleUpCreate(i *HelpCenterArticleUpCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/up.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterArticleUpdate(i *HelpCenterArticleUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticleVotesInput struct {
	ArticleID string
}

func (c *Client) HelpCenterArticleVotes(i *HelpCenterArticleVotesInput) ([]byte, error) {
	if i.ArticleID == "" {
		return nil, errors.New("Missing required field 'ArticleID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/%s/votes.json", i.ArticleID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesAttachmentCreateInput struct {
}

func (c *Client) HelpCenterArticlesAttachmentCreate(i *HelpCenterArticlesAttachmentCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/articles/attachments.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesAttachmentDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesAttachmentDelete(i *HelpCenterArticlesAttachmentDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/attachments/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesAttachmentShowInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesAttachmentShow(i *HelpCenterArticlesAttachmentShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/attachments/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesLabelShowInput struct {
	ID string
}

func (c *Client) HelpCenterArticlesLabelShow(i *HelpCenterArticlesLabelShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/articles/labels/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesLabelsListInput struct {
}

func (c *Client) HelpCenterArticlesLabelsList(i *HelpCenterArticlesLabelsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/articles/labels.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesListInput struct {
}

func (c *Client) HelpCenterArticlesList(i *HelpCenterArticlesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/articles.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterArticlesSearchInput struct {
}

func (c *Client) HelpCenterArticlesSearch(i *HelpCenterArticlesSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/articles/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoriesListInput struct {
}

func (c *Client) HelpCenterCategoriesList(i *HelpCenterCategoriesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/categories.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryArticlesInput struct {
	ID string
}

func (c *Client) HelpCenterCategoryArticles(i *HelpCenterCategoryArticlesInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/articles.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryCreateInput struct {
}

func (c *Client) HelpCenterCategoryCreate(i *HelpCenterCategoryCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/categories.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterCategoryDelete(i *HelpCenterCategoryDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategorySectionCreateInput struct {
	ID string
}

func (c *Client) HelpCenterCategorySectionCreate(i *HelpCenterCategorySectionCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/sections.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategorySectionsInput struct {
	ID string
}

func (c *Client) HelpCenterCategorySections(i *HelpCenterCategorySectionsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/sections.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryShowInput struct {
	ID string
}

func (c *Client) HelpCenterCategoryShow(i *HelpCenterCategoryShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategorySourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterCategorySourceLocaleUpdate(i *HelpCenterCategorySourceLocaleUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/source_locale.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryTranslationCreateInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslationCreate(i *HelpCenterCategoryTranslationCreateInput) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/translations.json", i.CategoryID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryTranslationUpdateInput struct {
	CategoryID string
	Locale     string
}

func (c *Client) HelpCenterCategoryTranslationUpdate(i *HelpCenterCategoryTranslationUpdateInput) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/translations/%s.json", i.CategoryID, i.Locale)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryTranslationsInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslations(i *HelpCenterCategoryTranslationsInput) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/translations.json", i.CategoryID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryTranslationsMissingInput struct {
	CategoryID string
}

func (c *Client) HelpCenterCategoryTranslationsMissing(i *HelpCenterCategoryTranslationsMissingInput) ([]byte, error) {
	if i.CategoryID == "" {
		return nil, errors.New("Missing required field 'CategoryID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s/translations/missing.json", i.CategoryID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterCategoryUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterCategoryUpdate(i *HelpCenterCategoryUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/categories/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterIncrementalArticlesListInput struct {
}

func (c *Client) HelpCenterIncrementalArticlesList(i *HelpCenterIncrementalArticlesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/incremental/articles.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterLocalesListInput struct {
}

func (c *Client) HelpCenterLocalesList(i *HelpCenterLocalesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/locales.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionAccessPolicyInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionAccessPolicy(i *HelpCenterSectionAccessPolicyInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/access_policy.json", i.SectionID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionAccessPolicyUpdateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionAccessPolicyUpdate(i *HelpCenterSectionAccessPolicyUpdateInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/access_policy.json", i.SectionID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionArticleCreateInput struct {
	ID string
}

func (c *Client) HelpCenterSectionArticleCreate(i *HelpCenterSectionArticleCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/articles.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionArticlesInput struct {
	ID string
}

func (c *Client) HelpCenterSectionArticles(i *HelpCenterSectionArticlesInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/articles.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterSectionDelete(i *HelpCenterSectionDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionShowInput struct {
	ID string
}

func (c *Client) HelpCenterSectionShow(i *HelpCenterSectionShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionSourceLocaleUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterSectionSourceLocaleUpdate(i *HelpCenterSectionSourceLocaleUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/source_locale.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionSubscriptionCreateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionSubscriptionCreate(i *HelpCenterSectionSubscriptionCreateInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/subscriptions.json", i.SectionID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionSubscriptionDeleteInput struct {
	SectionID string
	ID        string
}

func (c *Client) HelpCenterSectionSubscriptionDelete(i *HelpCenterSectionSubscriptionDeleteInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/subscriptions/%s.json", i.SectionID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionSubscriptionShowInput struct {
	SectionID string
	ID        string
}

func (c *Client) HelpCenterSectionSubscriptionShow(i *HelpCenterSectionSubscriptionShowInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/subscriptions/%s.json", i.SectionID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionSubscriptionsInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionSubscriptions(i *HelpCenterSectionSubscriptionsInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/subscriptions.json", i.SectionID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionTranslationCreateInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslationCreate(i *HelpCenterSectionTranslationCreateInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/translations.json", i.SectionID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionTranslationUpdateInput struct {
	SectionID string
	Locale    string
}

func (c *Client) HelpCenterSectionTranslationUpdate(i *HelpCenterSectionTranslationUpdateInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	if i.Locale == "" {
		return nil, errors.New("Missing required field 'Locale'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/translations/%s.json", i.SectionID, i.Locale)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionTranslationsInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslations(i *HelpCenterSectionTranslationsInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/translations.json", i.SectionID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionTranslationsMissingInput struct {
	SectionID string
}

func (c *Client) HelpCenterSectionTranslationsMissing(i *HelpCenterSectionTranslationsMissingInput) ([]byte, error) {
	if i.SectionID == "" {
		return nil, errors.New("Missing required field 'SectionID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s/translations/missing.json", i.SectionID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterSectionUpdate(i *HelpCenterSectionUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/sections/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterSectionsListInput struct {
}

func (c *Client) HelpCenterSectionsList(i *HelpCenterSectionsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/sections.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterTranslationDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterTranslationDelete(i *HelpCenterTranslationDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/translations/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserArticlesInput struct {
	ID string
}

func (c *Client) HelpCenterUserArticles(i *HelpCenterUserArticlesInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/users/%s/articles.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserCommentsInput struct {
	ID string
}

func (c *Client) HelpCenterUserComments(i *HelpCenterUserCommentsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/users/%s/comments.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentCreateInput struct {
}

func (c *Client) HelpCenterUserSegmentCreate(i *HelpCenterUserSegmentCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/user_segments.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentDelete(i *HelpCenterUserSegmentDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/user_segments/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentSectionsInput struct {
	UserSegmentID string
}

func (c *Client) HelpCenterUserSegmentSections(i *HelpCenterUserSegmentSectionsInput) ([]byte, error) {
	if i.UserSegmentID == "" {
		return nil, errors.New("Missing required field 'UserSegmentID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/user_segments/%s/sections.json", i.UserSegmentID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentShowInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentShow(i *HelpCenterUserSegmentShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/user_segments/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentTopicsInput struct {
	UserSegmentID string
}

func (c *Client) HelpCenterUserSegmentTopics(i *HelpCenterUserSegmentTopicsInput) ([]byte, error) {
	if i.UserSegmentID == "" {
		return nil, errors.New("Missing required field 'UserSegmentID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/user_segments/%s/topics.json", i.UserSegmentID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentUpdateInput struct {
	ID string
}

func (c *Client) HelpCenterUserSegmentUpdate(i *HelpCenterUserSegmentUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/user_segments/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentsApplicableListInput struct {
}

func (c *Client) HelpCenterUserSegmentsApplicableList(i *HelpCenterUserSegmentsApplicableListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/user_segments/applicable.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSegmentsListInput struct {
}

func (c *Client) HelpCenterUserSegmentsList(i *HelpCenterUserSegmentsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/help_center/user_segments.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserSubscriptionsInput struct {
	UserID string
}

func (c *Client) HelpCenterUserSubscriptions(i *HelpCenterUserSubscriptionsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/users/%s/subscriptions.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterUserVotesInput struct {
	UserID string
}

func (c *Client) HelpCenterUserVotes(i *HelpCenterUserVotesInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/users/%s/votes.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterVoteDeleteInput struct {
	ID string
}

func (c *Client) HelpCenterVoteDelete(i *HelpCenterVoteDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/votes/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HelpCenterVoteShowInput struct {
	ID string
}

func (c *Client) HelpCenterVoteShow(i *HelpCenterVoteShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/help_center/votes/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ImportsTicketInput struct {
}

func (c *Client) ImportsTicket(i *ImportsTicketInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/imports/tickets.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ImportsTicketsCreateManyInput struct {
}

func (c *Client) ImportsTicketsCreateMany(i *ImportsTicketsCreateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/imports/tickets/create_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type IncrementalOrganizationsListInput struct {
}

func (c *Client) IncrementalOrganizationsList(i *IncrementalOrganizationsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/incremental/organizations.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type IncrementalSampleInput struct {
	Item string
}

func (c *Client) IncrementalSample(i *IncrementalSampleInput) ([]byte, error) {
	if i.Item == "" {
		return nil, errors.New("Missing required field 'Item'")
	}

	path := fmt.Sprintf("/api/v2/incremental/%s/sample.json", i.Item)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type IncrementalTicketEventsListInput struct {
}

func (c *Client) IncrementalTicketEventsList(i *IncrementalTicketEventsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/incremental/ticket_events.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type IncrementalTicketMetricEventsListInput struct {
}

func (c *Client) IncrementalTicketMetricEventsList(i *IncrementalTicketMetricEventsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/incremental/ticket_metric_events.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type IncrementalTicketsListInput struct {
}

func (c *Client) IncrementalTicketsList(i *IncrementalTicketsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/incremental/tickets.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type IncrementalUsersListInput struct {
}

func (c *Client) IncrementalUsersList(i *IncrementalUsersListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/incremental/users.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type JobStatusShowInput struct {
	ID string
}

func (c *Client) JobStatusShow(i *JobStatusShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/job_statuses/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type JobStatusesListInput struct {
}

func (c *Client) JobStatusesList(i *JobStatusesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/job_statuses.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type JobStatusesShowManyInput struct {
}

func (c *Client) JobStatusesShowMany(i *JobStatusesShowManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/job_statuses/show_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type LocaleShowInput struct {
	ID string
}

func (c *Client) LocaleShow(i *LocaleShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/locales/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type LocalesAgentListInput struct {
}

func (c *Client) LocalesAgentList(i *LocalesAgentListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/locales/agent.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type LocalesCurrentListInput struct {
}

func (c *Client) LocalesCurrentList(i *LocalesCurrentListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/locales/current.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type LocalesDetectBestLocaleInput struct {
}

func (c *Client) LocalesDetectBestLocale(i *LocalesDetectBestLocaleInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/locales/detect_best_locale.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type LocalesListInput struct {
}

func (c *Client) LocalesList(i *LocalesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/locales.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type LocalesPublicListInput struct {
}

func (c *Client) LocalesPublicList(i *LocalesPublicListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/locales/public.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroApplyInput struct {
	ID string
}

func (c *Client) MacroApply(i *MacroApplyInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/macros/%s/apply.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroAttachmentCreateInput struct {
	MacroID string
}

func (c *Client) MacroAttachmentCreate(i *MacroAttachmentCreateInput) ([]byte, error) {
	if i.MacroID == "" {
		return nil, errors.New("Missing required field 'MacroID'")
	}

	path := fmt.Sprintf("/api/v2/macros/%s/attachments.json", i.MacroID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroAttachmentsInput struct {
	MacroID string
}

func (c *Client) MacroAttachments(i *MacroAttachmentsInput) ([]byte, error) {
	if i.MacroID == "" {
		return nil, errors.New("Missing required field 'MacroID'")
	}

	path := fmt.Sprintf("/api/v2/macros/%s/attachments.json", i.MacroID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroCreateInput struct {
}

func (c *Client) MacroCreate(i *MacroCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroDeleteInput struct {
	ID string
}

func (c *Client) MacroDelete(i *MacroDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/macros/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroShowInput struct {
	ID string
}

func (c *Client) MacroShow(i *MacroShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/macros/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacroUpdateInput struct {
	ID string
}

func (c *Client) MacroUpdate(i *MacroUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/macros/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosActionsListInput struct {
}

func (c *Client) MacrosActionsList(i *MacrosActionsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/actions.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosActiveListInput struct {
}

func (c *Client) MacrosActiveList(i *MacrosActiveListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/active.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosAttachmentCreateInput struct {
}

func (c *Client) MacrosAttachmentCreate(i *MacrosAttachmentCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/attachments.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosAttachmentShowInput struct {
	ID string
}

func (c *Client) MacrosAttachmentShow(i *MacrosAttachmentShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/macros/attachments/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosCategoriesListInput struct {
}

func (c *Client) MacrosCategoriesList(i *MacrosCategoriesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/categories.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosDestroyManyInput struct {
}

func (c *Client) MacrosDestroyMany(i *MacrosDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosListInput struct {
}

func (c *Client) MacrosList(i *MacrosListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosNewListInput struct {
}

func (c *Client) MacrosNewList(i *MacrosNewListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/new.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosSearchInput struct {
}

func (c *Client) MacrosSearch(i *MacrosSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type MacrosUpdateManyInput struct {
}

func (c *Client) MacrosUpdateMany(i *MacrosUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/macros/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsIncrementalRecipientsListInput struct {
}

func (c *Client) NpsIncrementalRecipientsList(i *NpsIncrementalRecipientsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/nps/incremental/recipients.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsIncrementalResponsesListInput struct {
}

func (c *Client) NpsIncrementalResponsesList(i *NpsIncrementalResponsesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/nps/incremental/responses.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyCloseCreateInput struct {
	ID string
}

func (c *Client) NpsSurveyCloseCreate(i *NpsSurveyCloseCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/close", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyInvitationCreateInput struct {
	ID string
}

func (c *Client) NpsSurveyInvitationCreate(i *NpsSurveyInvitationCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/invitations.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyInvitationShowInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyInvitationShow(i *NpsSurveyInvitationShowInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/invitations/%s.json", i.SurveyID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyInvitationsInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyInvitations(i *NpsSurveyInvitationsInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/invitations.json", i.SurveyID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyPreviewInput struct {
	ID string
}

func (c *Client) NpsSurveyPreview(i *NpsSurveyPreviewInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/preview", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyRecipientCreateInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipientCreate(i *NpsSurveyRecipientCreateInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/recipients.json", i.SurveyID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyRecipientShowInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyRecipientShow(i *NpsSurveyRecipientShowInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/recipients/%s.json", i.SurveyID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyRecipientUpdateInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyRecipientUpdate(i *NpsSurveyRecipientUpdateInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/recipients/%s.json", i.SurveyID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyRecipientsInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipients(i *NpsSurveyRecipientsInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/recipients.json", i.SurveyID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyRecipientsSearchInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyRecipientsSearch(i *NpsSurveyRecipientsSearchInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/recipients/search.json", i.SurveyID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyResponseCreateInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyResponseCreate(i *NpsSurveyResponseCreateInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/responses.json", i.SurveyID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyResponseShowInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyResponseShow(i *NpsSurveyResponseShowInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/responses/%s.json", i.SurveyID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyResponseUpdateInput struct {
	SurveyID string
	ID       string
}

func (c *Client) NpsSurveyResponseUpdate(i *NpsSurveyResponseUpdateInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/responses/%s.json", i.SurveyID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyResponsesInput struct {
	SurveyID string
}

func (c *Client) NpsSurveyResponses(i *NpsSurveyResponsesInput) ([]byte, error) {
	if i.SurveyID == "" {
		return nil, errors.New("Missing required field 'SurveyID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s/responses.json", i.SurveyID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveyShowInput struct {
	ID string
}

func (c *Client) NpsSurveyShow(i *NpsSurveyShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/nps/surveys/%s", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveys1UpdateInput struct {
}

func (c *Client) NpsSurveys1Update(i *NpsSurveys1UpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/nps/surveys/1")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type NpsSurveysListInput struct {
}

func (c *Client) NpsSurveysList(i *NpsSurveysListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/nps/surveys")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthClientCreateInput struct {
}

func (c *Client) OauthClientCreate(i *OauthClientCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/clients")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthClientDeleteInput struct {
	ID string
}

func (c *Client) OauthClientDelete(i *OauthClientDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/oauth/clients/%s", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthClientGenerateSecretUpdateInput struct {
	ID string
}

func (c *Client) OauthClientGenerateSecretUpdate(i *OauthClientGenerateSecretUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/oauth/clients/%s/generate_secret.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthClientShowInput struct {
	ID string
}

func (c *Client) OauthClientShow(i *OauthClientShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/oauth/clients/%s", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthClientUpdateInput struct {
	ID string
}

func (c *Client) OauthClientUpdate(i *OauthClientUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/oauth/clients/%s", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthClientsListInput struct {
}

func (c *Client) OauthClientsList(i *OauthClientsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/clients")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthGlobalClientsListInput struct {
}

func (c *Client) OauthGlobalClientsList(i *OauthGlobalClientsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/global_clients.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthTokenCreateInput struct {
}

func (c *Client) OauthTokenCreate(i *OauthTokenCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/tokens.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthTokenDeleteInput struct {
	ID string
}

func (c *Client) OauthTokenDelete(i *OauthTokenDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/oauth/tokens/%s", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthTokenShowInput struct {
	ID string
}

func (c *Client) OauthTokenShow(i *OauthTokenShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/oauth/tokens/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthTokensCurrentDeleteInput struct {
}

func (c *Client) OauthTokensCurrentDelete(i *OauthTokensCurrentDeleteInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/tokens/current.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthTokensCurrentListInput struct {
}

func (c *Client) OauthTokensCurrentList(i *OauthTokensCurrentListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/tokens/current.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OauthTokensListInput struct {
}

func (c *Client) OauthTokensList(i *OauthTokensListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/oauth/tokens")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationCreateInput struct {
}

func (c *Client) OrganizationCreate(i *OrganizationCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationDeleteInput struct {
	ID string
}

func (c *Client) OrganizationDelete(i *OrganizationDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationFieldCreateInput struct {
}

func (c *Client) OrganizationFieldCreate(i *OrganizationFieldCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_fields.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationFieldDeleteInput struct {
	ID string
}

func (c *Client) OrganizationFieldDelete(i *OrganizationFieldDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_fields/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationFieldShowInput struct {
	ID string
}

func (c *Client) OrganizationFieldShow(i *OrganizationFieldShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_fields/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationFieldUpdateInput struct {
	ID string
}

func (c *Client) OrganizationFieldUpdate(i *OrganizationFieldUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_fields/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationFieldsListInput struct {
}

func (c *Client) OrganizationFieldsList(i *OrganizationFieldsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_fields.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationFieldsReorderInput struct {
}

func (c *Client) OrganizationFieldsReorder(i *OrganizationFieldsReorderInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_fields/reorder.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationMembershipCreateInput struct {
}

func (c *Client) OrganizationMembershipCreate(i *OrganizationMembershipCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_memberships.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationMembershipDeleteInput struct {
	ID string
}

func (c *Client) OrganizationMembershipDelete(i *OrganizationMembershipDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_memberships/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationMembershipShowInput struct {
	ID string
}

func (c *Client) OrganizationMembershipShow(i *OrganizationMembershipShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_memberships/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationMembershipsCreateManyInput struct {
}

func (c *Client) OrganizationMembershipsCreateMany(i *OrganizationMembershipsCreateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_memberships/create_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationMembershipsDestroyManyInput struct {
}

func (c *Client) OrganizationMembershipsDestroyMany(i *OrganizationMembershipsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_memberships/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationMembershipsListInput struct {
}

func (c *Client) OrganizationMembershipsList(i *OrganizationMembershipsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_memberships.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationOrganizationMembershipsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationOrganizationMemberships(i *OrganizationOrganizationMembershipsInput) ([]byte, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/organization_memberships.json", i.OrganizationID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationRelatedInput struct {
	ID string
}

func (c *Client) OrganizationRelated(i *OrganizationRelatedInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/related.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationRequestsInput struct {
	ID string
}

func (c *Client) OrganizationRequests(i *OrganizationRequestsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/requests.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationShowInput struct {
	ID string
}

func (c *Client) OrganizationShow(i *OrganizationShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationSubscriptionCreateInput struct {
}

func (c *Client) OrganizationSubscriptionCreate(i *OrganizationSubscriptionCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_subscriptions.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationSubscriptionDeleteInput struct {
	ID string
}

func (c *Client) OrganizationSubscriptionDelete(i *OrganizationSubscriptionDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_subscriptions/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationSubscriptionShowInput struct {
	ID string
}

func (c *Client) OrganizationSubscriptionShow(i *OrganizationSubscriptionShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organization_subscriptions/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationSubscriptionsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationSubscriptions(i *OrganizationSubscriptionsInput) ([]byte, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/subscriptions.json", i.OrganizationID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationSubscriptionsListInput struct {
}

func (c *Client) OrganizationSubscriptionsList(i *OrganizationSubscriptionsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organization_subscriptions.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationTagCreateInput struct {
	ID string
}

func (c *Client) OrganizationTagCreate(i *OrganizationTagCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/tags.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationTagsInput struct {
	ID string
}

func (c *Client) OrganizationTags(i *OrganizationTagsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/tags.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationTagsDeleteInput struct {
	ID string
}

func (c *Client) OrganizationTagsDelete(i *OrganizationTagsDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/tags.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationTagsUpdateInput struct {
	ID string
}

func (c *Client) OrganizationTagsUpdate(i *OrganizationTagsUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/tags.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationTicketsInput struct {
	OrganizationID string
}

func (c *Client) OrganizationTickets(i *OrganizationTicketsInput) ([]byte, error) {
	if i.OrganizationID == "" {
		return nil, errors.New("Missing required field 'OrganizationID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/tickets.json", i.OrganizationID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationUpdateInput struct {
	ID string
}

func (c *Client) OrganizationUpdate(i *OrganizationUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationUsersInput struct {
	ID string
}

func (c *Client) OrganizationUsers(i *OrganizationUsersInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/organizations/%s/users.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsAutocompleteInput struct {
}

func (c *Client) OrganizationsAutocomplete(i *OrganizationsAutocompleteInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/autocomplete.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsCreateManyInput struct {
}

func (c *Client) OrganizationsCreateMany(i *OrganizationsCreateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/create_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsCreateOrUpdateInput struct {
}

func (c *Client) OrganizationsCreateOrUpdate(i *OrganizationsCreateOrUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/create_or_update.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsDestroyManyInput struct {
}

func (c *Client) OrganizationsDestroyMany(i *OrganizationsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsListInput struct {
}

func (c *Client) OrganizationsList(i *OrganizationsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsSearchInput struct {
}

func (c *Client) OrganizationsSearch(i *OrganizationsSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsShowManyInput struct {
}

func (c *Client) OrganizationsShowMany(i *OrganizationsShowManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/show_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type OrganizationsUpdateManyInput struct {
}

func (c *Client) OrganizationsUpdateMany(i *OrganizationsUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/organizations/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ProblemsAutocompleteInput struct {
}

func (c *Client) ProblemsAutocomplete(i *ProblemsAutocompleteInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/problems/autocomplete.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ProblemsListInput struct {
}

func (c *Client) ProblemsList(i *ProblemsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/problems.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type PushNotificationDevicesDestroyManyInput struct {
}

func (c *Client) PushNotificationDevicesDestroyMany(i *PushNotificationDevicesDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/push_notification_devices/destroy_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RecipientAddressCreateInput struct {
}

func (c *Client) RecipientAddressCreate(i *RecipientAddressCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/recipient_addresses.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RecipientAddressDeleteInput struct {
	ID string
}

func (c *Client) RecipientAddressDelete(i *RecipientAddressDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/recipient_addresses/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RecipientAddressShowInput struct {
	ID string
}

func (c *Client) RecipientAddressShow(i *RecipientAddressShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/recipient_addresses/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RecipientAddressUpdateInput struct {
	ID string
}

func (c *Client) RecipientAddressUpdate(i *RecipientAddressUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/recipient_addresses/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RecipientAddressVerifyInput struct {
	ID string
}

func (c *Client) RecipientAddressVerify(i *RecipientAddressVerifyInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/recipient_addresses/%s/verify.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RecipientAddressesListInput struct {
}

func (c *Client) RecipientAddressesList(i *RecipientAddressesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/recipient_addresses.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestCommentShowInput struct {
	RequestID string
	ID        string
}

func (c *Client) RequestCommentShow(i *RequestCommentShowInput) ([]byte, error) {
	if i.RequestID == "" {
		return nil, errors.New("Missing required field 'RequestID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/requests/%s/comments/%s.json", i.RequestID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestCommentsInput struct {
	ID string
}

func (c *Client) RequestComments(i *RequestCommentsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/requests/%s/comments.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestCreateInput struct {
}

func (c *Client) RequestCreate(i *RequestCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/requests.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestShowInput struct {
	ID string
}

func (c *Client) RequestShow(i *RequestShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/requests/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestUpdateInput struct {
	ID string
}

func (c *Client) RequestUpdate(i *RequestUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/requests/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestsCcdListInput struct {
}

func (c *Client) RequestsCcdList(i *RequestsCcdListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/requests/ccd.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestsListInput struct {
}

func (c *Client) RequestsList(i *RequestsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/requests.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestsOpenListInput struct {
}

func (c *Client) RequestsOpenList(i *RequestsOpenListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/requests/open.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestsSearchInput struct {
}

func (c *Client) RequestsSearch(i *RequestsSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/requests/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequestsSolvedListInput struct {
}

func (c *Client) RequestsSolvedList(i *RequestsSolvedListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/requests/solved.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ResourceCollectionCreateInput struct {
}

func (c *Client) ResourceCollectionCreate(i *ResourceCollectionCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/resource_collections.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ResourceCollectionDeleteInput struct {
	ID string
}

func (c *Client) ResourceCollectionDelete(i *ResourceCollectionDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/resource_collections/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ResourceCollectionShowInput struct {
	ID string
}

func (c *Client) ResourceCollectionShow(i *ResourceCollectionShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/resource_collections/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ResourceCollectionsListInput struct {
}

func (c *Client) ResourceCollectionsList(i *ResourceCollectionsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/resource_collections.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ResourceCollectionsUpdateInput struct {
}

func (c *Client) ResourceCollectionsUpdate(i *ResourceCollectionsUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/resource_collections.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SatisfactionRatingShowInput struct {
	ID string
}

func (c *Client) SatisfactionRatingShow(i *SatisfactionRatingShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/satisfaction_ratings/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SatisfactionRatingsListInput struct {
}

func (c *Client) SatisfactionRatingsList(i *SatisfactionRatingsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/satisfaction_ratings.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SatisfactionReasonShowInput struct {
	ID string
}

func (c *Client) SatisfactionReasonShow(i *SatisfactionReasonShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/satisfaction_reasons/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SatisfactionReasonsListInput struct {
}

func (c *Client) SatisfactionReasonsList(i *SatisfactionReasonsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/satisfaction_reasons.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SearchInput struct {
}

func (c *Client) Search(i *SearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SessionsListInput struct {
}

func (c *Client) SessionsList(i *SessionsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/sessions.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SharingAgreementCreateInput struct {
}

func (c *Client) SharingAgreementCreate(i *SharingAgreementCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/sharing_agreements.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SharingAgreementDeleteInput struct {
	ID string
}

func (c *Client) SharingAgreementDelete(i *SharingAgreementDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/sharing_agreements/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SharingAgreementShowInput struct {
	ID string
}

func (c *Client) SharingAgreementShow(i *SharingAgreementShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/sharing_agreements/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SharingAgreementUpdateInput struct {
	ID string
}

func (c *Client) SharingAgreementUpdate(i *SharingAgreementUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/sharing_agreements/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SharingAgreementsListInput struct {
}

func (c *Client) SharingAgreementsList(i *SharingAgreementsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/sharing_agreements.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ShortcutCreateInput struct {
}

func (c *Client) ShortcutCreate(i *ShortcutCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/shortcuts")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ShortcutDeleteInput struct {
	ShortcutName string
}

func (c *Client) ShortcutDelete(i *ShortcutDeleteInput) ([]byte, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	path := fmt.Sprintf("/api/v2/shortcuts/%s", i.ShortcutName)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ShortcutShowInput struct {
	ShortcutName string
}

func (c *Client) ShortcutShow(i *ShortcutShowInput) ([]byte, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	path := fmt.Sprintf("/api/v2/shortcuts/%s", i.ShortcutName)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ShortcutUpdateInput struct {
	ShortcutName string
}

func (c *Client) ShortcutUpdate(i *ShortcutUpdateInput) ([]byte, error) {
	if i.ShortcutName == "" {
		return nil, errors.New("Missing required field 'ShortcutName'")
	}

	path := fmt.Sprintf("/api/v2/shortcuts/%s", i.ShortcutName)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ShortcutsListInput struct {
}

func (c *Client) ShortcutsList(i *ShortcutsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/shortcuts")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SkipCreateInput struct {
}

func (c *Client) SkipCreate(i *SkipCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/skips.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SkipsListInput struct {
}

func (c *Client) SkipsList(i *SkipsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/skips.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPoliciesDefinitionsListInput struct {
}

func (c *Client) SlasPoliciesDefinitionsList(i *SlasPoliciesDefinitionsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/slas/policies/definitions.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPoliciesListInput struct {
}

func (c *Client) SlasPoliciesList(i *SlasPoliciesListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/slas/policies")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPoliciesReorderInput struct {
}

func (c *Client) SlasPoliciesReorder(i *SlasPoliciesReorderInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/slas/policies/reorder.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPolicyCreateInput struct {
}

func (c *Client) SlasPolicyCreate(i *SlasPolicyCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/slas/policies")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPolicyDeleteInput struct {
	ID string
}

func (c *Client) SlasPolicyDelete(i *SlasPolicyDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/slas/policies/%s", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPolicyShowInput struct {
	ID string
}

func (c *Client) SlasPolicyShow(i *SlasPolicyShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/slas/policies/%s", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SlasPolicyUpdateInput struct {
	ID string
}

func (c *Client) SlasPolicyUpdate(i *SlasPolicyUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/slas/policies/%s", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type StreamAgentShowInput struct {
	MetricKey string
}

func (c *Client) StreamAgentShow(i *StreamAgentShowInput) ([]byte, error) {
	if i.MetricKey == "" {
		return nil, errors.New("Missing required field 'MetricKey'")
	}

	path := fmt.Sprintf("/stream/agents/%s", i.MetricKey)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type StreamAgentsAgentsOnlineListInput struct {
}

func (c *Client) StreamAgentsAgentsOnlineList(i *StreamAgentsAgentsOnlineListInput) ([]byte, error) {
	path := fmt.Sprintf("/stream/agents/agents_online")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type StreamAgentsListInput struct {
}

func (c *Client) StreamAgentsList(i *StreamAgentsListInput) ([]byte, error) {
	path := fmt.Sprintf("/stream/agents")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type StreamChatShowInput struct {
	MetricKey string
}

func (c *Client) StreamChatShow(i *StreamChatShowInput) ([]byte, error) {
	if i.MetricKey == "" {
		return nil, errors.New("Missing required field 'MetricKey'")
	}

	path := fmt.Sprintf("/stream/chats/%s", i.MetricKey)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type StreamChatsListInput struct {
}

func (c *Client) StreamChatsList(i *StreamChatsListInput) ([]byte, error) {
	path := fmt.Sprintf("/stream/chats")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type StreamChatsMissedChatsListInput struct {
}

func (c *Client) StreamChatsMissedChatsList(i *StreamChatsMissedChatsListInput) ([]byte, error) {
	path := fmt.Sprintf("/stream/chats/missed_chats")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SuspendedTicketDeleteInput struct {
	ID string
}

func (c *Client) SuspendedTicketDelete(i *SuspendedTicketDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/suspended_tickets/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SuspendedTicketRecoverInput struct {
	ID string
}

func (c *Client) SuspendedTicketRecover(i *SuspendedTicketRecoverInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/suspended_tickets/%s/recover.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SuspendedTicketShowInput struct {
	ID string
}

func (c *Client) SuspendedTicketShow(i *SuspendedTicketShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/suspended_tickets/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SuspendedTicketsDestroyManyInput struct {
}

func (c *Client) SuspendedTicketsDestroyMany(i *SuspendedTicketsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/suspended_tickets/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SuspendedTicketsListInput struct {
}

func (c *Client) SuspendedTicketsList(i *SuspendedTicketsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/suspended_tickets.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SuspendedTicketsRecoverManyInput struct {
}

func (c *Client) SuspendedTicketsRecoverMany(i *SuspendedTicketsRecoverManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/suspended_tickets/recover_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TagsListInput struct {
}

func (c *Client) TagsList(i *TagsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tags.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TargetCreateInput struct {
}

func (c *Client) TargetCreate(i *TargetCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/targets.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TargetDeleteInput struct {
	ID string
}

func (c *Client) TargetDelete(i *TargetDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/targets/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TargetShowInput struct {
	ID string
}

func (c *Client) TargetShow(i *TargetShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/targets/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TargetUpdateInput struct {
	ID string
}

func (c *Client) TargetUpdate(i *TargetUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/targets/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TargetsListInput struct {
}

func (c *Client) TargetsList(i *TargetsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/targets.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketAuditMakePrivateInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketAuditMakePrivate(i *TicketAuditMakePrivateInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/audits/%s/make_private.json", i.TicketID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketAuditShowInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketAuditShow(i *TicketAuditShowInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/audits/%s.json", i.TicketID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketAuditsInput struct {
	TicketID string
}

func (c *Client) TicketAudits(i *TicketAuditsInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/audits.json", i.TicketID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketCollaboratorsInput struct {
	ID string
}

func (c *Client) TicketCollaborators(i *TicketCollaboratorsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/collaborators.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketCommentAttachmentRedactInput struct {
	TicketID     string
	CommentID    string
	AttachmentID string
}

func (c *Client) TicketCommentAttachmentRedact(i *TicketCommentAttachmentRedactInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.CommentID == "" {
		return nil, errors.New("Missing required field 'CommentID'")
	}

	if i.AttachmentID == "" {
		return nil, errors.New("Missing required field 'AttachmentID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/comments/%s/attachments/%s/redact.json", i.TicketID, i.CommentID, i.AttachmentID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketCommentMakePrivateInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketCommentMakePrivate(i *TicketCommentMakePrivateInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/comments/%s/make_private.json", i.TicketID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketCommentRedactInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketCommentRedact(i *TicketCommentRedactInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/comments/%s/redact.json", i.TicketID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketCommentsInput struct {
	TicketID string
}

func (c *Client) TicketComments(i *TicketCommentsInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/comments.json", i.TicketID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketCreateInput struct {
}

func (c *Client) TicketCreate(i *TicketCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketDeleteInput struct {
	ID string
}

func (c *Client) TicketDelete(i *TicketDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldCreateInput struct {
}

func (c *Client) TicketFieldCreate(i *TicketFieldCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_fields.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldDeleteInput struct {
	ID string
}

func (c *Client) TicketFieldDelete(i *TicketFieldDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldOptionCreateInput struct {
	FieldID string
	ID      string
}

func (c *Client) TicketFieldOptionCreate(i *TicketFieldOptionCreateInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s/options/%s.json", i.FieldID, i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldOptionDeleteInput struct {
	FieldID string
	ID      string
}

func (c *Client) TicketFieldOptionDelete(i *TicketFieldOptionDeleteInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s/options/%s.json", i.FieldID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldOptionShowInput struct {
	FieldID string
	ID      string
}

func (c *Client) TicketFieldOptionShow(i *TicketFieldOptionShowInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s/options/%s.json", i.FieldID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldOptionsInput struct {
	FieldID string
}

func (c *Client) TicketFieldOptions(i *TicketFieldOptionsInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s/options.json", i.FieldID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldShowInput struct {
	ID string
}

func (c *Client) TicketFieldShow(i *TicketFieldShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldUpdateInput struct {
	ID string
}

func (c *Client) TicketFieldUpdate(i *TicketFieldUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_fields/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFieldsListInput struct {
}

func (c *Client) TicketFieldsList(i *TicketFieldsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_fields.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormCloneInput struct {
	ID string
}

func (c *Client) TicketFormClone(i *TicketFormCloneInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_forms/%s/clone.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormCreateInput struct {
}

func (c *Client) TicketFormCreate(i *TicketFormCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_forms.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormDeleteInput struct {
	ID string
}

func (c *Client) TicketFormDelete(i *TicketFormDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_forms/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormShowInput struct {
	ID string
}

func (c *Client) TicketFormShow(i *TicketFormShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_forms/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormUpdateInput struct {
	ID string
}

func (c *Client) TicketFormUpdate(i *TicketFormUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_forms/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormsListInput struct {
}

func (c *Client) TicketFormsList(i *TicketFormsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_forms.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormsReorderInput struct {
}

func (c *Client) TicketFormsReorder(i *TicketFormsReorderInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_forms/reorder.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketFormsShowManyInput struct {
}

func (c *Client) TicketFormsShowMany(i *TicketFormsShowManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_forms/show_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketIncidentsInput struct {
	ID string
}

func (c *Client) TicketIncidents(i *TicketIncidentsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/incidents.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketMacroApplyInput struct {
	TicketID string
	ID       string
}

func (c *Client) TicketMacroApply(i *TicketMacroApplyInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/macros/%s/apply.json", i.TicketID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketMarkAsSpamInput struct {
	ID string
}

func (c *Client) TicketMarkAsSpam(i *TicketMarkAsSpamInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/mark_as_spam.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketMergeInput struct {
	ID string
}

func (c *Client) TicketMerge(i *TicketMergeInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/merge.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketMetricShowInput struct {
	TicketMetricID string
}

func (c *Client) TicketMetricShow(i *TicketMetricShowInput) ([]byte, error) {
	if i.TicketMetricID == "" {
		return nil, errors.New("Missing required field 'TicketMetricID'")
	}

	path := fmt.Sprintf("/api/v2/ticket_metrics/%s.json", i.TicketMetricID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketMetricsInput struct {
	TicketID string
}

func (c *Client) TicketMetrics(i *TicketMetricsInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/metrics.json", i.TicketID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketMetricsListInput struct {
}

func (c *Client) TicketMetricsList(i *TicketMetricsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/ticket_metrics.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketRelatedInput struct {
	ID string
}

func (c *Client) TicketRelated(i *TicketRelatedInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/related.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketSatisfactionRatingCreateInput struct {
	TicketID string
}

func (c *Client) TicketSatisfactionRatingCreate(i *TicketSatisfactionRatingCreateInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/satisfaction_rating.json", i.TicketID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketShowInput struct {
	ID string
}

func (c *Client) TicketShow(i *TicketShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketSkipsInput struct {
	TicketID string
}

func (c *Client) TicketSkips(i *TicketSkipsInput) ([]byte, error) {
	if i.TicketID == "" {
		return nil, errors.New("Missing required field 'TicketID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/skips.json", i.TicketID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketTagCreateInput struct {
	ID string
}

func (c *Client) TicketTagCreate(i *TicketTagCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/tags.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketTagsInput struct {
	ID string
}

func (c *Client) TicketTags(i *TicketTagsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/tags.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketTagsDeleteInput struct {
	ID string
}

func (c *Client) TicketTagsDelete(i *TicketTagsDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/tags.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketTagsUpdateInput struct {
	ID string
}

func (c *Client) TicketTagsUpdate(i *TicketTagsUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s/tags.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketUpdateInput struct {
	ID string
}

func (c *Client) TicketUpdate(i *TicketUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsCreateManyInput struct {
}

func (c *Client) TicketsCreateMany(i *TicketsCreateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets/create_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsDestroyManyInput struct {
}

func (c *Client) TicketsDestroyMany(i *TicketsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsListInput struct {
}

func (c *Client) TicketsList(i *TicketsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsMarkManyAsSpamInput struct {
}

func (c *Client) TicketsMarkManyAsSpam(i *TicketsMarkManyAsSpamInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets/mark_many_as_spam.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsRecentListInput struct {
}

func (c *Client) TicketsRecentList(i *TicketsRecentListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets/recent.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsShowManyInput struct {
}

func (c *Client) TicketsShowMany(i *TicketsShowManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets/show_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TicketsUpdateManyInput struct {
}

func (c *Client) TicketsUpdateMany(i *TicketsUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/tickets/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TopicTagCreateInput struct {
	ID string
}

func (c *Client) TopicTagCreate(i *TopicTagCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/topics/%s/tags.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TopicTagsInput struct {
	ID string
}

func (c *Client) TopicTags(i *TopicTagsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/topics/%s/tags.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TopicTagsDeleteInput struct {
	ID string
}

func (c *Client) TopicTagsDelete(i *TopicTagsDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/topics/%s/tags.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TopicTagsUpdateInput struct {
	ID string
}

func (c *Client) TopicTagsUpdate(i *TopicTagsUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/topics/%s/tags.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerCreateInput struct {
}

func (c *Client) TriggerCreate(i *TriggerCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerDeleteByIDInput struct {
	ID string
}

func (c *Client) TriggerDeleteByID(i *TriggerDeleteByIDInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/triggers/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerDeleteByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerDeleteByTriggerName(i *TriggerDeleteByTriggerNameInput) ([]byte, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	path := fmt.Sprintf("/api/v2/triggers/%s", i.TriggerName)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerShowByIDInput struct {
	ID string
}

func (c *Client) TriggerShowByID(i *TriggerShowByIDInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/triggers/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerShowByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerShowByTriggerName(i *TriggerShowByTriggerNameInput) ([]byte, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	path := fmt.Sprintf("/api/v2/triggers/%s", i.TriggerName)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerUpdateByIDInput struct {
	ID string
}

func (c *Client) TriggerUpdateByID(i *TriggerUpdateByIDInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/triggers/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggerUpdateByTriggerNameInput struct {
	TriggerName string
}

func (c *Client) TriggerUpdateByTriggerName(i *TriggerUpdateByTriggerNameInput) ([]byte, error) {
	if i.TriggerName == "" {
		return nil, errors.New("Missing required field 'TriggerName'")
	}

	path := fmt.Sprintf("/api/v2/triggers/%s", i.TriggerName)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggersActiveListInput struct {
}

func (c *Client) TriggersActiveList(i *TriggersActiveListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers/active.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggersDestroyManyInput struct {
}

func (c *Client) TriggersDestroyMany(i *TriggersDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggersListInput struct {
}

func (c *Client) TriggersList(i *TriggersListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggersReorderInput struct {
}

func (c *Client) TriggersReorder(i *TriggersReorderInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers/reorder.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggersSearchInput struct {
}

func (c *Client) TriggersSearch(i *TriggersSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type TriggersUpdateManyInput struct {
}

func (c *Client) TriggersUpdateMany(i *TriggersUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/triggers/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UploadCreateInput struct {
}

func (c *Client) UploadCreate(i *UploadCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/uploads.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UploadDeleteInput struct {
	Token string
}

func (c *Client) UploadDelete(i *UploadDeleteInput) ([]byte, error) {
	if i.Token == "" {
		return nil, errors.New("Missing required field 'Token'")
	}

	path := fmt.Sprintf("/api/v2/uploads/%s.json", i.Token)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserCreateInput struct {
}

func (c *Client) UserCreate(i *UserCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserDeleteInput struct {
	ID string
}

func (c *Client) UserDelete(i *UserDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldCreateInput struct {
}

func (c *Client) UserFieldCreate(i *UserFieldCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/user_fields.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldDeleteInput struct {
	ID string
}

func (c *Client) UserFieldDelete(i *UserFieldDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldOptionCreateInput struct {
	FieldID string
}

func (c *Client) UserFieldOptionCreate(i *UserFieldOptionCreateInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s/options.json", i.FieldID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldOptionDeleteInput struct {
	FieldID string
	ID      string
}

func (c *Client) UserFieldOptionDelete(i *UserFieldOptionDeleteInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s/options/%s.json", i.FieldID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldOptionShowInput struct {
	FieldID string
	ID      string
}

func (c *Client) UserFieldOptionShow(i *UserFieldOptionShowInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s/options/%s.json", i.FieldID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldOptionsInput struct {
	FieldID string
}

func (c *Client) UserFieldOptions(i *UserFieldOptionsInput) ([]byte, error) {
	if i.FieldID == "" {
		return nil, errors.New("Missing required field 'FieldID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s/options.json", i.FieldID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldShowInput struct {
	ID string
}

func (c *Client) UserFieldShow(i *UserFieldShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldUpdateInput struct {
	ID string
}

func (c *Client) UserFieldUpdate(i *UserFieldUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/user_fields/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldsListInput struct {
}

func (c *Client) UserFieldsList(i *UserFieldsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/user_fields.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserFieldsReorderInput struct {
}

func (c *Client) UserFieldsReorder(i *UserFieldsReorderInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/user_fields/reorder.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserGroupMembershipCreateInput struct {
	UserID string
}

func (c *Client) UserGroupMembershipCreate(i *UserGroupMembershipCreateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/group_memberships.json", i.UserID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserGroupMembershipDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserGroupMembershipDelete(i *UserGroupMembershipDeleteInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/group_memberships/%s.json", i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserGroupMembershipMakeDefaultInput struct {
	UserID       string
	MembershipID string
}

func (c *Client) UserGroupMembershipMakeDefault(i *UserGroupMembershipMakeDefaultInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.MembershipID == "" {
		return nil, errors.New("Missing required field 'MembershipID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/group_memberships/%s/make_default.json", i.UserID, i.MembershipID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserGroupMembershipShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserGroupMembershipShow(i *UserGroupMembershipShowInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/group_memberships/%s.json", i.UserID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserGroupMembershipsInput struct {
	UserID string
}

func (c *Client) UserGroupMemberships(i *UserGroupMembershipsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/group_memberships.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserGroupsInput struct {
	UserID string
}

func (c *Client) UserGroups(i *UserGroupsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/groups.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentitiesInput struct {
	UserID string
}

func (c *Client) UserIdentities(i *UserIdentitiesInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityCreateInput struct {
	UserID string
}

func (c *Client) UserIdentityCreate(i *UserIdentityCreateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities.json", i.UserID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityDelete(i *UserIdentityDeleteInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities/%s.json", i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityMakePrimaryInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityMakePrimary(i *UserIdentityMakePrimaryInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities/%s/make_primary", i.UserID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityRequestVerificationUpdateInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityRequestVerificationUpdate(i *UserIdentityRequestVerificationUpdateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities/%s/request_verification.json", i.UserID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityShow(i *UserIdentityShowInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities/%s.json", i.UserID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityUpdateInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityUpdate(i *UserIdentityUpdateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities/%s.json", i.UserID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserIdentityVerifyInput struct {
	UserID string
	ID     string
}

func (c *Client) UserIdentityVerify(i *UserIdentityVerifyInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/identities/%s/verify", i.UserID, i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserMergeInput struct {
	ID string
}

func (c *Client) UserMerge(i *UserMergeInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/merge.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationMembershipCreateInput struct {
	UserID string
}

func (c *Client) UserOrganizationMembershipCreate(i *UserOrganizationMembershipCreateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organization_memberships.json", i.UserID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationMembershipDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserOrganizationMembershipDelete(i *UserOrganizationMembershipDeleteInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organization_memberships/%s.json", i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationMembershipMakeDefaultInput struct {
	ID           string
	MembershipID string
}

func (c *Client) UserOrganizationMembershipMakeDefault(i *UserOrganizationMembershipMakeDefaultInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	if i.MembershipID == "" {
		return nil, errors.New("Missing required field 'MembershipID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organization_memberships/%s/make_default.json", i.ID, i.MembershipID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationMembershipShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserOrganizationMembershipShow(i *UserOrganizationMembershipShowInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organization_memberships/%s.json", i.UserID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationMembershipsInput struct {
	UserID string
}

func (c *Client) UserOrganizationMemberships(i *UserOrganizationMembershipsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organization_memberships.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationSubscriptionsInput struct {
	UserID string
}

func (c *Client) UserOrganizationSubscriptions(i *UserOrganizationSubscriptionsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organization_subscriptions.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserOrganizationsInput struct {
	UserID string
}

func (c *Client) UserOrganizations(i *UserOrganizationsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/organizations.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserPasswordCreateInput struct {
	UserID string
}

func (c *Client) UserPasswordCreate(i *UserPasswordCreateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/password.json", i.UserID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserPasswordRequirementsInput struct {
	UserID string
}

func (c *Client) UserPasswordRequirements(i *UserPasswordRequirementsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/password/requirements.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserPasswordUpdateInput struct {
	UserID string
}

func (c *Client) UserPasswordUpdate(i *UserPasswordUpdateInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/password.json", i.UserID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserRelatedInput struct {
	ID string
}

func (c *Client) UserRelated(i *UserRelatedInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/related.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserRequestsInput struct {
	ID string
}

func (c *Client) UserRequests(i *UserRequestsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/requests.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserSessionDeleteInput struct {
	UserID string
	ID     string
}

func (c *Client) UserSessionDelete(i *UserSessionDeleteInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/sessions/%s.json", i.UserID, i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserSessionShowInput struct {
	UserID string
	ID     string
}

func (c *Client) UserSessionShow(i *UserSessionShowInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/sessions/%s.json", i.UserID, i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserSessionsInput struct {
	UserID string
}

func (c *Client) UserSessions(i *UserSessionsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/sessions.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserSessionsDeleteInput struct {
	UserID string
}

func (c *Client) UserSessionsDelete(i *UserSessionsDeleteInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/sessions.json", i.UserID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserShowInput struct {
	ID string
}

func (c *Client) UserShow(i *UserShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserSkipsInput struct {
	UserID string
}

func (c *Client) UserSkips(i *UserSkipsInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/skips.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTagCreateInput struct {
	ID string
}

func (c *Client) UserTagCreate(i *UserTagCreateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tags.json", i.ID)
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTagsInput struct {
	ID string
}

func (c *Client) UserTags(i *UserTagsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tags.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTagsDeleteInput struct {
	ID string
}

func (c *Client) UserTagsDelete(i *UserTagsDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tags.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTagsUpdateInput struct {
	ID string
}

func (c *Client) UserTagsUpdate(i *UserTagsUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tags.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTicketsAssignedInput struct {
	UserID string
}

func (c *Client) UserTicketsAssigned(i *UserTicketsAssignedInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tickets/assigned.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTicketsCcdInput struct {
	UserID string
}

func (c *Client) UserTicketsCcd(i *UserTicketsCcdInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tickets/ccd.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserTicketsRequestedInput struct {
	UserID string
}

func (c *Client) UserTicketsRequested(i *UserTicketsRequestedInput) ([]byte, error) {
	if i.UserID == "" {
		return nil, errors.New("Missing required field 'UserID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s/tickets/requested.json", i.UserID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserUpdateInput struct {
	ID string
}

func (c *Client) UserUpdate(i *UserUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/users/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersAutocompleteInput struct {
}

func (c *Client) UsersAutocomplete(i *UsersAutocompleteInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/autocomplete.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersCreateManyInput struct {
}

func (c *Client) UsersCreateMany(i *UsersCreateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/create_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersCreateOrUpdateInput struct {
}

func (c *Client) UsersCreateOrUpdate(i *UsersCreateOrUpdateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/create_or_update.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersCreateOrUpdateManyInput struct {
}

func (c *Client) UsersCreateOrUpdateMany(i *UsersCreateOrUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/create_or_update_many.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersDestroyManyInput struct {
}

func (c *Client) UsersDestroyMany(i *UsersDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersListInput struct {
}

func (c *Client) UsersList(i *UsersListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersMeInput struct {
}

func (c *Client) UsersMe(i *UsersMeInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/me.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersMeLogoutInput struct {
}

func (c *Client) UsersMeLogout(i *UsersMeLogoutInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/me/logout.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersMeMergeInput struct {
}

func (c *Client) UsersMeMerge(i *UsersMeMergeInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/me/merge.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersMeOauthClientsInput struct {
}

func (c *Client) UsersMeOauthClients(i *UsersMeOauthClientsInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/me/oauth/clients.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersMeSessionInput struct {
}

func (c *Client) UsersMeSession(i *UsersMeSessionInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/me/session.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersRequestCreateInput struct {
}

func (c *Client) UsersRequestCreate(i *UsersRequestCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/request_create.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersSearchInput struct {
}

func (c *Client) UsersSearch(i *UsersSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersShowManyInput struct {
}

func (c *Client) UsersShowMany(i *UsersShowManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/show_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UsersUpdateManyInput struct {
}

func (c *Client) UsersUpdateMany(i *UsersUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/users/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewCountInput struct {
	ID string
}

func (c *Client) ViewCount(i *ViewCountInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s/count.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewCreateInput struct {
}

func (c *Client) ViewCreate(i *ViewCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewDeleteInput struct {
	ID string
}

func (c *Client) ViewDelete(i *ViewDeleteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s.json", i.ID)
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewExecuteInput struct {
	ID string
}

func (c *Client) ViewExecute(i *ViewExecuteInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s/execute.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewExportInput struct {
	ID string
}

func (c *Client) ViewExport(i *ViewExportInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s/export.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewShowInput struct {
	ID string
}

func (c *Client) ViewShow(i *ViewShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewTicketsInput struct {
	ID string
}

func (c *Client) ViewTickets(i *ViewTicketsInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s/tickets.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewUpdateInput struct {
	ID string
}

func (c *Client) ViewUpdate(i *ViewUpdateInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/views/%s.json", i.ID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsActiveListInput struct {
}

func (c *Client) ViewsActiveList(i *ViewsActiveListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/active.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsCompactInput struct {
}

func (c *Client) ViewsCompact(i *ViewsCompactInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/compact.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsCountManyListInput struct {
}

func (c *Client) ViewsCountManyList(i *ViewsCountManyListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/count_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsDestroyManyInput struct {
}

func (c *Client) ViewsDestroyMany(i *ViewsDestroyManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/destroy_many.json")
	resp, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsListInput struct {
}

func (c *Client) ViewsList(i *ViewsListInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsPreviewInput struct {
}

func (c *Client) ViewsPreview(i *ViewsPreviewInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/preview.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsPreviewCountInput struct {
}

func (c *Client) ViewsPreviewCount(i *ViewsPreviewCountInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/preview/count.json")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsSearchInput struct {
}

func (c *Client) ViewsSearch(i *ViewsSearchInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/search.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsShowManyInput struct {
}

func (c *Client) ViewsShowMany(i *ViewsShowManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/show_many.json")
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ViewsUpdateManyInput struct {
}

func (c *Client) ViewsUpdateMany(i *ViewsUpdateManyInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/views/update_many.json")
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type VisitorCreateInput struct {
}

func (c *Client) VisitorCreate(i *VisitorCreateInput) ([]byte, error) {
	path := fmt.Sprintf("/api/v2/visitors")
	resp, err := c.Request("POST", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type VisitorShowInput struct {
	VisitorID string
}

func (c *Client) VisitorShow(i *VisitorShowInput) ([]byte, error) {
	if i.VisitorID == "" {
		return nil, errors.New("Missing required field 'VisitorID'")
	}

	path := fmt.Sprintf("/api/v2/visitors/%s", i.VisitorID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type VisitorUpdateInput struct {
	VisitorID string
}

func (c *Client) VisitorUpdate(i *VisitorUpdateInput) ([]byte, error) {
	if i.VisitorID == "" {
		return nil, errors.New("Missing required field 'VisitorID'")
	}

	path := fmt.Sprintf("/api/v2/visitors/%s", i.VisitorID)
	resp, err := c.Request("PUT", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
