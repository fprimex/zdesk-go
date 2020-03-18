package zdesk

import (
	"encoding/json"
)

type ChatAccount struct {
	AccountKey string `jsonapi:"attr,account_key"`
	CreateDate string `jsonapi:"attr,create_date"`
	Status string `jsonapi:"attr,status"`
	Plan map[string]interface{} `jsonapi:"attr,plan"`
	Billing string `jsonapi:"attr,billing"`
}

type ChatAgent struct {
	ID int `jsonapi:"attr,id"`
	FirstName string `jsonapi:"attr,first_name"`
	LastName string `jsonapi:"attr,last_name"`
	DisplayName string `jsonapi:"attr,display_name"`
	CreateDate string `jsonapi:"attr,create_date"`
	Email int `jsonapi:"attr,email"`
	Role string `jsonapi:"attr,role"`
	Enabled int `jsonapi:"attr,enabled"`
	Departments []interface{} `jsonapi:"attr,departments"`
	Skills []interface{} `jsonapi:"attr,skills"`
}

type ChatOverview struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	ClientSecret string `jsonapi:"attr,client_secret"`
	ClientType string `jsonapi:"attr,client_type"`
	RedirectUris string `jsonapi:"attr,redirect_uris"`
	Scopes string `jsonapi:"attr,scopes"`
	AgentID int `jsonapi:"attr,agent_id"`
	CreateDate string `jsonapi:"attr,create_date"`
	UpdateDate string `jsonapi:"attr,update_date"`
}

type ChatBan struct {
	ID int `jsonapi:"attr,id"`
	Type string `jsonapi:"attr,type"`
	VisitorID string `jsonapi:"attr,visitor_id"`
	VisitorName string `jsonapi:"attr,visitor_name"`
	IpAddress string `jsonapi:"attr,ip_address"`
	Reason string `jsonapi:"attr,reason"`
}

type Chat struct {
	ID int `jsonapi:"attr,id"`
	Visitor map[string]interface{} `jsonapi:"attr,visitor"`
	Type string `jsonapi:"attr,type"`
	StartedBy string `jsonapi:"attr,started_by"`
	Session map[string]interface{} `jsonapi:"attr,session"`
	Timestamp string `jsonapi:"attr,timestamp"`
	Count map[string]interface{} `jsonapi:"attr,count"`
	Duration string `jsonapi:"attr,duration"`
	DepartmentID int `jsonapi:"attr,department_id"`
	DepartmentName int `jsonapi:"attr,department_name"`
	ResponseTime map[string]interface{} `jsonapi:"attr,response_time"`
	AgentNames []interface{} `jsonapi:"attr,agent_names"`
	AgentIds []interface{} `jsonapi:"attr,agent_ids"`
	Triggered bool `jsonapi:"attr,triggered"`
	TriggeredResponse bool `jsonapi:"attr,triggered_response"`
	Unread bool `jsonapi:"attr,unread"`
	Missed bool `jsonapi:"attr,missed"`
	History []interface{} `jsonapi:"attr,history"`
	Conversions []interface{} `jsonapi:"attr,conversions"`
	Tags []interface{} `jsonapi:"attr,tags"`
	Rating string `jsonapi:"attr,rating"`
	Comment string `jsonapi:"attr,comment"`
	Webpath []interface{} `jsonapi:"attr,webpath"`
	ZendeskTicketID int `jsonapi:"attr,zendesk_ticket_id"`
}

type ChatDepartment struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Enabled int `jsonapi:"attr,enabled"`
	Members []interface{} `jsonapi:"attr,members"`
	Settings map[string]interface{} `jsonapi:"attr,settings"`
}

type ChatGoal struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Enabled int `jsonapi:"attr,enabled"`
	AttributionModel string `jsonapi:"attr,attribution_model"`
	AttributionPeriod int `jsonapi:"attr,attribution_period"`
	Settings map[string]interface{} `jsonapi:"attr,settings"`
}

type ChatShortcut struct {
	Name int `jsonapi:"attr,name"`
	Options int `jsonapi:"attr,options"`
	Message string `jsonapi:"attr,message"`
	Tags []interface{} `jsonapi:"attr,tags"`
}

type ChatSkill struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Enabled int `jsonapi:"attr,enabled"`
	Members []interface{} `jsonapi:"attr,members"`
}

type ChatTrigger struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Enabled int `jsonapi:"attr,enabled"`
	Description string `jsonapi:"attr,description"`
}

type ChatVisitor struct {
	ID int `jsonapi:"attr,id"`
	DisplayName string `jsonapi:"attr,display_name"`
	Created string `jsonapi:"attr,created"`
	Email string `jsonapi:"attr,email"`
	Banned int `jsonapi:"attr,banned"`
	Notes string `jsonapi:"attr,notes"`
	Phone int `jsonapi:"attr,phone"`
}

type AccountSetting struct {
	HeaderColor string `jsonapi:"attr,header_color"`
	PageBackgroundColor string `jsonapi:"attr,page_background_color"`
	TabBackgroundColor string `jsonapi:"attr,tab_background_color"`
	TextColor string `jsonapi:"attr,text_color"`
	HeaderLogoUrl string `jsonapi:"attr,header_logo_url"`
	FaviconUrl string `jsonapi:"attr,favicon_url"`
}

type ActivityStream struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Title string `jsonapi:"attr,title"`
	Verb string `jsonapi:"attr,verb"`
	User User `jsonapi:"attr,user"`
	Actor User `jsonapi:"attr,actor"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Object map[string]interface{} `jsonapi:"attr,object"`
	Target map[string]interface{} `jsonapi:"attr,target"`
}

type AppInstallationLocation struct {
	ID int `jsonapi:"attr,id"`
	LocationName string `jsonapi:"attr,location_name"`
	Installations []interface{} `jsonapi:"attr,installations"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type AppLocation struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	ProductCode int `jsonapi:"attr,product_code"`
	HostApplication string `jsonapi:"attr,host_application"`
	Orderable json.RawMessage `jsonapi:"attr,orderable"`
}

type App struct {
	ID string `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	AppID int `jsonapi:"attr,app_id"`
	Total int `jsonapi:"attr,total"`
	Progress int `jsonapi:"attr,progress"`
	Status string `jsonapi:"attr,status"`
	Message string `jsonapi:"attr,message"`
	RetryIn int `jsonapi:"attr,retry_in"`
}

type Attachment struct {
	ID int `jsonapi:"attr,id"`
	FileName string `jsonapi:"attr,file_name"`
	ContentUrl string `jsonapi:"attr,content_url"`
	ContentType string `jsonapi:"attr,content_type"`
	Size int `jsonapi:"attr,size"`
	Thumbnails []interface{} `jsonapi:"attr,thumbnails"`
	Inline bool `jsonapi:"attr,inline"`
}

type AuditLog struct {
	ID int `jsonapi:"attr,id"`
	Action string `jsonapi:"attr,action"`
	ActorID int `jsonapi:"attr,actor_id"`
	SourceID int `jsonapi:"attr,source_id"`
	SourceType string `jsonapi:"attr,source_type"`
	SourceLabel string `jsonapi:"attr,source_label"`
	ChangesDescription string `jsonapi:"attr,changes_description"`
	IpAddress string `jsonapi:"attr,ip_address"`
	CreatedAt string `jsonapi:"attr,created_at"`
}

type Automation struct {
	ID int `jsonapi:"attr,id"`
	Title string `jsonapi:"attr,title"`
	Active bool `jsonapi:"attr,active"`
	Position int `jsonapi:"attr,position"`
	Conditions json.RawMessage `jsonapi:"attr,conditions"`
	Actions json.RawMessage `jsonapi:"attr,actions"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type Bookmark struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Ticket map[string]interface{} `jsonapi:"attr,ticket"`
	CreatedAt string `jsonapi:"attr,created_at"`
}

type Brand struct {
	Url string `jsonapi:"attr,url"`
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	BrandUrl string `jsonapi:"attr,brand_url"`
	HasHelpCenter bool `jsonapi:"attr,has_help_center"`
	HelpCenterState string `jsonapi:"attr,help_center_state"`
	Active bool `jsonapi:"attr,active"`
	Default bool `jsonapi:"attr,default"`
	Logo Attachment `jsonapi:"attr,logo"`
	TicketFormIds []interface{} `jsonapi:"attr,ticket_form_ids"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Subdomain string `jsonapi:"attr,subdomain"`
	HostMapping string `jsonapi:"attr,host_mapping"`
	SignatureTemplate string `jsonapi:"attr,signature_template"`
}

type EndUser struct {
	ID int `jsonapi:"attr,id"`
	Email string `jsonapi:"attr,email"`
	Name string `jsonapi:"attr,name"`
	CreatedAt string `jsonapi:"attr,created_at"`
	Locale string `jsonapi:"attr,locale"`
	LocaleID int `jsonapi:"attr,locale_id"`
	OrganizationID int `jsonapi:"attr,organization_id"`
	Phone string `jsonapi:"attr,phone"`
	Photo Attachment `jsonapi:"attr,photo"`
	Role string `jsonapi:"attr,role"`
	TimeZone string `jsonapi:"attr,time_zone"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Url string `jsonapi:"attr,url"`
	Verified bool `jsonapi:"attr,verified"`
}

type AuthorizedGlobalClient struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Identifier string `jsonapi:"attr,identifier"`
	Company string `jsonapi:"attr,company"`
	Description string `jsonapi:"attr,description"`
	LogoUrl string `jsonapi:"attr,logo_url"`
}

type GroupMembership struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	UserID int `jsonapi:"attr,user_id"`
	GroupID int `jsonapi:"attr,group_id"`
	Default bool `jsonapi:"attr,default"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type Group struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Name string `jsonapi:"attr,name"`
	Deleted bool `jsonapi:"attr,deleted"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type IncrementalExport struct {
	EndTime string `jsonapi:"attr,end_time"`
	NextPage string `jsonapi:"attr,next_page"`
	Count int `jsonapi:"attr,count"`
}

type JobStatus struct {
	ID string `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Total int `jsonapi:"attr,total"`
	Progress int `jsonapi:"attr,progress"`
	Status string `jsonapi:"attr,status"`
	Message string `jsonapi:"attr,message"`
	Results []interface{} `jsonapi:"attr,results"`
}

type Macro struct {
	ID int `jsonapi:"attr,id"`
	Actions json.RawMessage `jsonapi:"attr,actions"`
	Active bool `jsonapi:"attr,active"`
	Description string `jsonapi:"attr,description"`
	Position int `jsonapi:"attr,position"`
	Restriction map[string]interface{} `jsonapi:"attr,restriction"`
	Title string `jsonapi:"attr,title"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type OauthClient struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Identifier string `jsonapi:"attr,identifier"`
	Secret string `jsonapi:"attr,secret"`
	Company string `jsonapi:"attr,company"`
	Description string `jsonapi:"attr,description"`
	RedirectUri []interface{} `jsonapi:"attr,redirect_uri"`
	UserID int `jsonapi:"attr,user_id"`
	Global bool `jsonapi:"attr,global"`
	Url string `jsonapi:"attr,url"`
	LogoUrl string `jsonapi:"attr,logo_url"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type OauthToken struct {
	ID int `jsonapi:"attr,id"`
	UserID int `jsonapi:"attr,user_id"`
	ClientID int `jsonapi:"attr,client_id"`
	Token string `jsonapi:"attr,token"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
	Scopes []interface{} `jsonapi:"attr,scopes"`
	Url string `jsonapi:"attr,url"`
	CreatedAt string `jsonapi:"attr,created_at"`
	ExpiresAt string `jsonapi:"attr,expires_at"`
	UsedAt string `jsonapi:"attr,used_at"`
}

type OrganizationField struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Key string `jsonapi:"attr,key"`
	Type string `jsonapi:"attr,type"`
	Title string `jsonapi:"attr,title"`
	RawTitle string `jsonapi:"attr,raw_title"`
	Description string `jsonapi:"attr,description"`
	RawDescription string `jsonapi:"attr,raw_description"`
	Position int `jsonapi:"attr,position"`
	Active bool `jsonapi:"attr,active"`
	System bool `jsonapi:"attr,system"`
	RegexpForValidation string `jsonapi:"attr,regexp_for_validation"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Tag string `jsonapi:"attr,tag"`
	CustomFieldOptions []interface{} `jsonapi:"attr,custom_field_options"`
}

type OrganizationMembership struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	UserID int `jsonapi:"attr,user_id"`
	OrganizationID int `jsonapi:"attr,organization_id"`
	Default bool `jsonapi:"attr,default"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type Organization struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	ExternalID string `jsonapi:"attr,external_id"`
	Name string `jsonapi:"attr,name"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	DomainNames []interface{} `jsonapi:"attr,domain_names"`
	Details string `jsonapi:"attr,details"`
	Notes string `jsonapi:"attr,notes"`
	GroupID int `jsonapi:"attr,group_id"`
	SharedTickets bool `jsonapi:"attr,shared_tickets"`
	SharedComments bool `jsonapi:"attr,shared_comments"`
	Tags []interface{} `jsonapi:"attr,tags"`
	OrganizationFields string `jsonapi:"attr,organization_fields"`
}

type Request struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Subject string `jsonapi:"attr,subject"`
	Description string `jsonapi:"attr,description"`
	Status string `jsonapi:"attr,status"`
	Priority string `jsonapi:"attr,priority"`
	Type string `jsonapi:"attr,type"`
	CustomFields []interface{} `jsonapi:"attr,custom_fields"`
	OrganizationID int `jsonapi:"attr,organization_id"`
	RequesterID int `jsonapi:"attr,requester_id"`
	AssigneeID int `jsonapi:"attr,assignee_id"`
	GroupID int `jsonapi:"attr,group_id"`
	CollaboratorIds []interface{} `jsonapi:"attr,collaborator_ids"`
	Via json.RawMessage `jsonapi:"attr,via"`
	IsPublic bool `jsonapi:"attr,is_public"`
	DueAt string `jsonapi:"attr,due_at"`
	CanBeSolvedByMe bool `jsonapi:"attr,can_be_solved_by_me"`
	Solved bool `jsonapi:"attr,solved"`
	TicketFormID int `jsonapi:"attr,ticket_form_id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Recipient string `jsonapi:"attr,recipient"`
	FollowupSourceID int `jsonapi:"attr,followup_source_id"`
}

type ResourceCollection struct {
	ID int `jsonapi:"attr,id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type SatisfactionRating struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	AssigneeID int `jsonapi:"attr,assignee_id"`
	GroupID int `jsonapi:"attr,group_id"`
	RequesterID int `jsonapi:"attr,requester_id"`
	TicketID int `jsonapi:"attr,ticket_id"`
	Score string `jsonapi:"attr,score"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Comment string `jsonapi:"attr,comment"`
	Reason string `jsonapi:"attr,reason"`
}

type SatisfactionReason struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	ReasonCode int `jsonapi:"attr,reason_code"`
	Value string `jsonapi:"attr,value"`
	RawValue string `jsonapi:"attr,raw_value"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	DeletedAt string `jsonapi:"attr,deleted_at"`
}

type Schedule struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	TimeZone string `jsonapi:"attr,time_zone"`
	Intervals []interface{} `jsonapi:"attr,intervals"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type Search struct {
	Count int `jsonapi:"attr,count"`
	NextPage string `jsonapi:"attr,next_page"`
	PrevPage string `jsonapi:"attr,prev_page"`
	Results []interface{} `jsonapi:"attr,results"`
}

type Session struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	UserID string `jsonapi:"attr,user_id"`
	AuthenticatedAt string `jsonapi:"attr,authenticated_at"`
	LastSeenAt string `jsonapi:"attr,last_seen_at"`
}

type SharingAgreement struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Type string `jsonapi:"attr,type"`
	Status string `jsonapi:"attr,status"`
	PartnerName string `jsonapi:"attr,partner_name"`
	RemoteSubdomain string `jsonapi:"attr,remote_subdomain"`
	CreatedAt string `jsonapi:"attr,created_at"`
}

type SLAPolicy struct {
	ID int `jsonapi:"attr,id"`
	Title string `jsonapi:"attr,title"`
	Description string `jsonapi:"attr,description"`
	Position int `jsonapi:"attr,position"`
	Filter json.RawMessage `jsonapi:"attr,filter"`
	PolicyMetrics []interface{} `jsonapi:"attr,policy_metrics"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type SupportAddress struct {
	ID int `jsonapi:"attr,id"`
	Email string `jsonapi:"attr,email"`
	Name string `jsonapi:"attr,name"`
	Default bool `jsonapi:"attr,default"`
	BrandID int `jsonapi:"attr,brand_id"`
	ForwardingStatus string `jsonapi:"attr,forwarding_status"`
	SpfStatus string `jsonapi:"attr,spf_status"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type TargetFailure struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	TargetName string `jsonapi:"attr,target_name"`
	StatusCode string `jsonapi:"attr,status_code"`
	Message string `jsonapi:"attr,message"`
	CreatedAt string `jsonapi:"attr,created_at"`
	ConsecutiveFailureCount int `jsonapi:"attr,consecutive_failure_count"`
	RawRequest string `jsonapi:"attr,raw_request"`
	RawResponse string `jsonapi:"attr,raw_response"`
}

type Target struct {
	ID int `jsonapi:"attr,id"`
	Title string `jsonapi:"attr,title"`
	Type string `jsonapi:"attr,type"`
	Active bool `jsonapi:"attr,active"`
	CreatedAt string `jsonapi:"attr,created_at"`
}

type TicketAudit struct {
	ID int `jsonapi:"attr,id"`
	TicketID int `jsonapi:"attr,ticket_id"`
	Metadata string `jsonapi:"attr,metadata"`
	Via json.RawMessage `jsonapi:"attr,via"`
	CreatedAt string `jsonapi:"attr,created_at"`
	AuthorID int `jsonapi:"attr,author_id"`
	Events []interface{} `jsonapi:"attr,events"`
}

type TicketComment struct {
	ID int `jsonapi:"attr,id"`
	Type string `jsonapi:"attr,type"`
	Body string `jsonapi:"attr,body"`
	HtmlBody string `jsonapi:"attr,html_body"`
	PlainBody string `jsonapi:"attr,plain_body"`
	Public bool `jsonapi:"attr,public"`
	AuthorID int `jsonapi:"attr,author_id"`
	Attachments []interface{} `jsonapi:"attr,attachments"`
	Via map[string]interface{} `jsonapi:"attr,via"`
	Metadata map[string]interface{} `jsonapi:"attr,metadata"`
	CreatedAt string `jsonapi:"attr,created_at"`
}

type TicketField struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Type string `jsonapi:"attr,type"`
	Title string `jsonapi:"attr,title"`
	RawTitle string `jsonapi:"attr,raw_title"`
	Description string `jsonapi:"attr,description"`
	RawDescription string `jsonapi:"attr,raw_description"`
	Position int `jsonapi:"attr,position"`
	Active bool `jsonapi:"attr,active"`
	Required bool `jsonapi:"attr,required"`
	CollapsedForAgents bool `jsonapi:"attr,collapsed_for_agents"`
	RegexpForValidation string `jsonapi:"attr,regexp_for_validation"`
	TitleInPortal string `jsonapi:"attr,title_in_portal"`
	RawTitleInPortal string `jsonapi:"attr,raw_title_in_portal"`
	VisibleInPortal bool `jsonapi:"attr,visible_in_portal"`
	EditableInPortal bool `jsonapi:"attr,editable_in_portal"`
	RequiredInPortal bool `jsonapi:"attr,required_in_portal"`
	Tag string `jsonapi:"attr,tag"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	SystemFieldOptions []interface{} `jsonapi:"attr,system_field_options"`
	CustomFieldOptions []interface{} `jsonapi:"attr,custom_field_options"`
	Removable bool `jsonapi:"attr,removable"`
}

type TicketForm struct {
	Name string `jsonapi:"attr,name"`
	RawName string `jsonapi:"attr,raw_name"`
	DisplayName string `jsonapi:"attr,display_name"`
	RawDisplayName string `jsonapi:"attr,raw_display_name"`
	Position int `jsonapi:"attr,position"`
	Active bool `jsonapi:"attr,active"`
	EndUserVisible bool `jsonapi:"attr,end_user_visible"`
	Default bool `jsonapi:"attr,default"`
	TicketFieldIds []interface{} `jsonapi:"attr,ticket_field_ids"`
	InAllBrands bool `jsonapi:"attr,in_all_brands"`
	RestrictedBrandIds []interface{} `jsonapi:"attr,restricted_brand_ids"`
	InAllOrganizations bool `jsonapi:"attr,in_all_organizations"`
	RestrictedOrganizationIds []interface{} `jsonapi:"attr,restricted_organization_ids"`
}

type TicketMetricEvent struct {
	ID int `jsonapi:"attr,id"`
	TicketID int `jsonapi:"attr,ticket_id"`
	Metric string `jsonapi:"attr,metric"`
	InstanceID int `jsonapi:"attr,instance_id"`
	Type string `jsonapi:"attr,type"`
	Time string `jsonapi:"attr,time"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type TicketMetric struct {
	ID int `jsonapi:"attr,id"`
	TicketID int `jsonapi:"attr,ticket_id"`
	Url string `jsonapi:"attr,url"`
	GroupStations int `jsonapi:"attr,group_stations"`
	AssigneeStations int `jsonapi:"attr,assignee_stations"`
	Reopens int `jsonapi:"attr,reopens"`
	Replies int `jsonapi:"attr,replies"`
	AssigneeUpdatedAt string `jsonapi:"attr,assignee_updated_at"`
	RequesterUpdatedAt string `jsonapi:"attr,requester_updated_at"`
	StatusUpdatedAt string `jsonapi:"attr,status_updated_at"`
	InitiallyAssignedAt string `jsonapi:"attr,initially_assigned_at"`
	AssignedAt string `jsonapi:"attr,assigned_at"`
	SolvedAt string `jsonapi:"attr,solved_at"`
	LatestCommentAddedAt string `jsonapi:"attr,latest_comment_added_at"`
	FirstResolutionTimeInMinutes map[string]interface{} `jsonapi:"attr,first_resolution_time_in_minutes"`
	ReplyTimeInMinutes map[string]interface{} `jsonapi:"attr,reply_time_in_minutes"`
	FullResolutionTimeInMinutes map[string]interface{} `jsonapi:"attr,full_resolution_time_in_minutes"`
	AgentWaitTimeInMinutes map[string]interface{} `jsonapi:"attr,agent_wait_time_in_minutes"`
	RequesterWaitTimeInMinutes map[string]interface{} `jsonapi:"attr,requester_wait_time_in_minutes"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type TicketSkip struct {
	ID int `jsonapi:"attr,id"`
	TicketID int `jsonapi:"attr,ticket_id"`
	UserID int `jsonapi:"attr,user_id"`
	Reason string `jsonapi:"attr,reason"`
	Ticket Ticket `jsonapi:"attr,ticket"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type Ticket struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	ExternalID string `jsonapi:"attr,external_id"`
	Type string `jsonapi:"attr,type"`
	Subject string `jsonapi:"attr,subject"`
	RawSubject string `jsonapi:"attr,raw_subject"`
	Description string `jsonapi:"attr,description"`
	Priority string `jsonapi:"attr,priority"`
	Status string `jsonapi:"attr,status"`
	Recipient string `jsonapi:"attr,recipient"`
	RequesterID int `jsonapi:"attr,requester_id"`
	SubmitterID int `jsonapi:"attr,submitter_id"`
	AssigneeID int `jsonapi:"attr,assignee_id"`
	OrganizationID int `jsonapi:"attr,organization_id"`
	GroupID int `jsonapi:"attr,group_id"`
	CollaboratorIds []interface{} `jsonapi:"attr,collaborator_ids"`
	FollowerIds []interface{} `jsonapi:"attr,follower_ids"`
	ForumTopicID int `jsonapi:"attr,forum_topic_id"`
	ProblemID int `jsonapi:"attr,problem_id"`
	HasIncidents bool `jsonapi:"attr,has_incidents"`
	DueAt string `jsonapi:"attr,due_at"`
	Tags []interface{} `jsonapi:"attr,tags"`
	Via json.RawMessage `jsonapi:"attr,via"`
	CustomFields []interface{} `jsonapi:"attr,custom_fields"`
	SatisfactionRating map[string]interface{} `jsonapi:"attr,satisfaction_rating"`
	SharingAgreementIds []interface{} `jsonapi:"attr,sharing_agreement_ids"`
	FollowupIds []interface{} `jsonapi:"attr,followup_ids"`
	TicketFormID int `jsonapi:"attr,ticket_form_id"`
	BrandID int `jsonapi:"attr,brand_id"`
	AllowChannelback bool `jsonapi:"attr,allow_channelback"`
	IsPublic bool `jsonapi:"attr,is_public"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type Trigger struct {
	ID int `jsonapi:"attr,id"`
	Title string `jsonapi:"attr,title"`
	Active bool `jsonapi:"attr,active"`
	Position int `jsonapi:"attr,position"`
	Conditions json.RawMessage `jsonapi:"attr,conditions"`
	Actions json.RawMessage `jsonapi:"attr,actions"`
	Description string `jsonapi:"attr,description"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type TwitterChannel struct {
	ID int `jsonapi:"attr,id"`
	ScreenName string `jsonapi:"attr,screen_name"`
	TwitterUserID int `jsonapi:"attr,twitter_user_id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	AvatarUrl string `jsonapi:"attr,avatar_url"`
	Name string `jsonapi:"attr,name"`
	AllowReply bool `jsonapi:"attr,allow_reply"`
	CanReply bool `jsonapi:"attr,can_reply"`
	BrandID int `jsonapi:"attr,brand_id"`
}

type UserField struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Key string `jsonapi:"attr,key"`
	Type string `jsonapi:"attr,type"`
	Title string `jsonapi:"attr,title"`
	RawTitle string `jsonapi:"attr,raw_title"`
	Description string `jsonapi:"attr,description"`
	RawDescription string `jsonapi:"attr,raw_description"`
	Position int `jsonapi:"attr,position"`
	Active bool `jsonapi:"attr,active"`
	System bool `jsonapi:"attr,system"`
	RegexpForValidation string `jsonapi:"attr,regexp_for_validation"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Tag string `jsonapi:"attr,tag"`
	CustomFieldOptions []interface{} `jsonapi:"attr,custom_field_options"`
}

type UserIdentity struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	UserID int `jsonapi:"attr,user_id"`
	Type string `jsonapi:"attr,type"`
	Value string `jsonapi:"attr,value"`
	Verified bool `jsonapi:"attr,verified"`
	Primary bool `jsonapi:"attr,primary"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	UndeliverableCount int `jsonapi:"attr,undeliverable_count"`
	DeliverableState string `jsonapi:"attr,deliverable_state"`
}

type User struct {
	ID int `jsonapi:"attr,id"`
	Email string `jsonapi:"attr,email"`
	Name string `jsonapi:"attr,name"`
	CreatedAt string `jsonapi:"attr,created_at"`
	Locale string `jsonapi:"attr,locale"`
	LocaleID int `jsonapi:"attr,locale_id"`
	OrganizationID int `jsonapi:"attr,organization_id"`
	Phone string `jsonapi:"attr,phone"`
	Photo Attachment `jsonapi:"attr,photo"`
	Role string `jsonapi:"attr,role"`
	TimeZone string `jsonapi:"attr,time_zone"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	Url string `jsonapi:"attr,url"`
	Verified bool `jsonapi:"attr,verified"`
}

type View struct {
	ID int `jsonapi:"attr,id"`
	Title string `jsonapi:"attr,title"`
	Active bool `jsonapi:"attr,active"`
	SLAID int `jsonapi:"attr,sla_id"`
	Restriction map[string]interface{} `jsonapi:"attr,restriction"`
	Position int `jsonapi:"attr,position"`
	Execution json.RawMessage `jsonapi:"attr,execution"`
	Conditions json.RawMessage `jsonapi:"attr,conditions"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterAccessPolicy struct {
	ViewableBy string `jsonapi:"attr,viewable_by"`
	ManageableBy string `jsonapi:"attr,manageable_by"`
	RestrictedToGroupIds []interface{} `jsonapi:"attr,restricted_to_group_ids"`
	RestrictedToOrganizationIds []interface{} `jsonapi:"attr,restricted_to_organization_ids"`
	RequiredTags []interface{} `jsonapi:"attr,required_tags"`
}

type HelpCenterArticleAttachment struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	ArticleID int `jsonapi:"attr,article_id"`
	FileName string `jsonapi:"attr,file_name"`
	ContentUrl string `jsonapi:"attr,content_url"`
	ContentType string `jsonapi:"attr,content_type"`
	Size int `jsonapi:"attr,size"`
	Inline bool `jsonapi:"attr,inline"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterArticle struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	Title string `jsonapi:"attr,title"`
	Body string `jsonapi:"attr,body"`
	Locale string `jsonapi:"attr,locale"`
	SourceLocale string `jsonapi:"attr,source_locale"`
	AuthorID int `jsonapi:"attr,author_id"`
	CommentsDisabled bool `jsonapi:"attr,comments_disabled"`
	OutdatedLocales []interface{} `jsonapi:"attr,outdated_locales"`
	Outdated bool `jsonapi:"attr,outdated"`
	LabelNames string `jsonapi:"attr,label_names"`
	Draft bool `jsonapi:"attr,draft"`
	Promoted bool `jsonapi:"attr,promoted"`
	Position int `jsonapi:"attr,position"`
	VoteSum int `jsonapi:"attr,vote_sum"`
	VoteCount int `jsonapi:"attr,vote_count"`
	SectionID int `jsonapi:"attr,section_id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterCategory struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Locale string `jsonapi:"attr,locale"`
	SourceLocale string `jsonapi:"attr,source_locale"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	Outdated bool `jsonapi:"attr,outdated"`
	Position int `jsonapi:"attr,position"`
	TranslationIds []interface{} `jsonapi:"attr,translation_ids"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterArticleComment struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Body string `jsonapi:"attr,body"`
	AuthorID int `jsonapi:"attr,author_id"`
	SourceID int `jsonapi:"attr,source_id"`
	SourceType string `jsonapi:"attr,source_type"`
	Locale string `jsonapi:"attr,locale"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	VoteSum int `jsonapi:"attr,vote_sum"`
	VoteCount int `jsonapi:"attr,vote_count"`
}

type HelpCenterArticleLabel struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	Name string `jsonapi:"attr,name"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterPostComment struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	Body string `jsonapi:"attr,body"`
	AuthorID int `jsonapi:"attr,author_id"`
	PostID int `jsonapi:"attr,post_id"`
	Official bool `jsonapi:"attr,official"`
	VoteSum int `jsonapi:"attr,vote_sum"`
	VoteCount int `jsonapi:"attr,vote_count"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterPost struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	Title string `jsonapi:"attr,title"`
	Details string `jsonapi:"attr,details"`
	AuthorID int `jsonapi:"attr,author_id"`
	Pinned bool `jsonapi:"attr,pinned"`
	Featured bool `jsonapi:"attr,featured"`
	Closed bool `jsonapi:"attr,closed"`
	Status string `jsonapi:"attr,status"`
	VoteSum int `jsonapi:"attr,vote_sum"`
	VoteCount int `jsonapi:"attr,vote_count"`
	CommentCount int `jsonapi:"attr,comment_count"`
	FollowerCount int `jsonapi:"attr,follower_count"`
	TopicID int `jsonapi:"attr,topic_id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterSection struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Locale string `jsonapi:"attr,locale"`
	SourceLocale string `jsonapi:"attr,source_locale"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	CategoryID int `jsonapi:"attr,category_id"`
	Outdated bool `jsonapi:"attr,outdated"`
	Position int `jsonapi:"attr,position"`
	TranslationIds []interface{} `jsonapi:"attr,translation_ids"`
	ManageableBy string `jsonapi:"attr,manageable_by"`
	UserSegmentID int `jsonapi:"attr,user_segment_id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterSubscription struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	UserID int `jsonapi:"attr,user_id"`
	ContentID int `jsonapi:"attr,content_id"`
	ContentType string `jsonapi:"attr,content_type"`
	Locale string `jsonapi:"attr,locale"`
	IncludeComments bool `jsonapi:"attr,include_comments"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterTopic struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	Name string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Position int `jsonapi:"attr,position"`
	FollowerCount int `jsonapi:"attr,follower_count"`
	ManageableBy string `jsonapi:"attr,manageable_by"`
	UserSegmentID int `jsonapi:"attr,user_segment_id"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type HelpCenterTranslation struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	HtmlUrl string `jsonapi:"attr,html_url"`
	SourceID int `jsonapi:"attr,source_id"`
	SourceType string `jsonapi:"attr,source_type"`
	Locale string `jsonapi:"attr,locale"`
	Title string `jsonapi:"attr,title"`
	Body string `jsonapi:"attr,body"`
	Outdated bool `jsonapi:"attr,outdated"`
	Draft bool `jsonapi:"attr,draft"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	UpdatedByID int `jsonapi:"attr,updated_by_id"`
	CreatedByID int `jsonapi:"attr,created_by_id"`
}

type HelpCenterUserSegment struct {
	ID int `jsonapi:"attr,id"`
	UserType string `jsonapi:"attr,user_type"`
	GroupIds []interface{} `jsonapi:"attr,group_ids"`
	OrganizationIds []interface{} `jsonapi:"attr,organization_ids"`
	Tags []interface{} `jsonapi:"attr,tags"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	BuiltIn bool `jsonapi:"attr,built_in"`
}

type HelpCenterVote struct {
	ID int `jsonapi:"attr,id"`
	Url string `jsonapi:"attr,url"`
	UserID int `jsonapi:"attr,user_id"`
	Value int `jsonapi:"attr,value"`
	ItemID int `jsonapi:"attr,item_id"`
	ItemType string `jsonapi:"attr,item_type"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type NpsNpsInvitation struct {
	ID int `jsonapi:"attr,id"`
	Status string `jsonapi:"attr,status"`
	RecipientsCount int `jsonapi:"attr,recipients_count"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	DeliveredAt string `jsonapi:"attr,delivered_at"`
}

type NpsNpsRecipient struct {
	ID int `jsonapi:"attr,id"`
	UserID int `jsonapi:"attr,user_id"`
	Email string `jsonapi:"attr,email"`
	Name string `jsonapi:"attr,name"`
	Locale string `jsonapi:"attr,locale"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
	DeliveredAt string `jsonapi:"attr,delivered_at"`
}

type NpsNpsResponse struct {
	ID int `jsonapi:"attr,id"`
	RecipientID int `jsonapi:"attr,recipient_id"`
	Rating int `jsonapi:"attr,rating"`
	Comment string `jsonapi:"attr,comment"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

type NpsNpsSurvey struct {
	ID int `jsonapi:"attr,id"`
	Name string `jsonapi:"attr,name"`
	Subject string `jsonapi:"attr,subject"`
	EmailSubject string `jsonapi:"attr,email_subject"`
	FromEmailID int `jsonapi:"attr,from_email_id"`
	IntroText string `jsonapi:"attr,intro_text"`
	HighlightColor string `jsonapi:"attr,highlight_color"`
	RelationshipID int `jsonapi:"attr,relationship_id"`
	CommentsQuestion string `jsonapi:"attr,comments_question"`
	DeliveryMethod string `jsonapi:"attr,delivery_method"`
	Status string `jsonapi:"attr,status"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}


