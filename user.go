package nsone

type User struct {
	Name        string             `json:"name"`
	Username    string             `json:"username"`
	Notify      map[string]bool    `json:"notify"`
	LastAccess  int                `json:"last_access,omitempty"`
	Email       string             `json:"email"`
	Teams       []string           `json:"teams"`
	Permissions UserPermissionsMap `json:"permissions"`
}

type UserPermissionsMap struct {
	Dns        UserPermissionsDns        `json:"dns"`
	Data       UserPermissionsData       `json:"data"`
	Account    UserPermissionsAccount    `json:"account"`
	Monitoring UserPermissionsMonitoring `json:"monitoring"`
}

type UserPermissionsDns struct {
	ViewZones           bool     `json:"view_zones"`
	ManageZones         bool     `json:"manage_zones"`
	ZonesAllowByDefault bool     `json:"zones_allow_by_default"`
	ZonesDeny           []string `json:"zones_deny"`
	ZonesAllow          []string `json:"zones_allow"`
}

type UserPermissionsData struct {
	PushToDatafeeds   bool `json:"push_to_datafeeds"`
	ManageDatasources bool `json:"manage_datasources"`
	ManageDatafeeds   bool `json:"manage_datafeeds"`
}

type UserPermissionsAccount struct {
	ManageUsers           bool `json:"manage_users"`
	ManagePaymentMethods  bool `json:"manage_payment_methods"`
	ManagePlan            bool `json:"manage_plan"`
	ManageTeams           bool `json:"manage_teams"`
	ManageApikeys         bool `json:"manage_apikeys"`
	ManageAccountSettings bool `json:"manage_account_settings"`
	ViewActivityLog       bool `json:"view_activity_log"`
	ViewInvoices          bool `json:"view_invoices"`
}
type UserPermissionsMonitoring struct {
	ManageLists bool `json:"manage_lists"`
	ManageJobs  bool `json:"manage_jobs"`
	ViewJobs    bool `json:"view_jobs"`
	ManageUsers bool `json:"manage_users"`
}
