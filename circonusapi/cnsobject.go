// Copyright 2016 Alem Abreha <alem.abreha@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
package circonusapi



// Account API Object
type Account struct {
	Cid           string        `json:"_cid"`
	ContactGroups []string      `json:"_contact_groups"`
	Owner         string        `json:"_owner"`
	UiBaseUrl     string        `json:"_ui_base_url"`
	Usage         []interface{} `json:"_usage"`
	Address1      string        `json:"address1"`
	Address2      string        `json:"address2"`
	City          string        `json:"city"`
	CountryCode   string        `json:"country_code"`
	Description   string        `json:"description"`
	Invites       []interface{} `json:"invites"`
	Name          string        `json:"name"`
	State         string        `json:"state_prov"`
	TimeZone      string        `json:"timezone"`
	Users         []interface{} `json:"users"`
}

// Alert API Object
type Alert struct {
	Acknowledgement string   `json:"_acknowledgement"`
	AlertUrl        string   `json:"_alert_url"`
	Broker          string   `json:"_broker"`
	Check           string   `json:"_check"`
	CheckName       string   `json:"_check_name"`
	Cid             string   `json:"_cid"`
	ClearedOn       float64  `json:"_cleared_on"`
	ClearedValue    float64  `json:"_cleared_value"`
	Maintenance     []string `json:"_maintenance"`
	MetricLink      string   `json:"_metric_link"`
	MetricName      string   `json:"_metric_name"`
	MetricNotes     string   `json:"_metric_notes"`
	OccurredOn      float64  `json:"_occurred_on"`
	RuleSet         string   `json:"_rule_set"`
	Severity        int      `json:"_severity"`
	Tags            []string `json:"_tags"`
	Value           float64  `json:"_value"`
}

//Annotation API Object
type Annotation struct {
	Cid            string  `json:"_cid"`
	Created        float64 `json:"_created"`
	LastModified   float64 `json:"_last_modified"`
	LastModifiedBy string  `json:"_last_modified_by"`
	Category       string  `json:"category"`
	Description    string  `json:"description"`
	Start          float64 `json:"start"`
	Stop           float64 `json:"stop"`
	Title          string  `json:"title"`
}

// Broker API Object
type Broker struct {
	Cid       string        `json:"_cid"`
	Details   []interface{} `json:"_details"`
	Latitude  string        `json:"_latitude"`
	Longitude string        `json:"_longitude"`
	Name      string        `json:"_name"`
	Tags      []string      `json:"_tags"`
	Type      string        `json:"_type"`
}

//Broker Detail API Object
type BrokerDetail struct {
	Cn                     string   `json:"cn"`
	ExternalHost           string   `json:"external_host"`
	ExternalPort           float64  `json:"external_port"`
	IpAddress              string   `json:"ipaddress"`
	MinimumVersionRequired float64  `json:"minimum_version_required"`
	Modules                []string `json:"modules"`
	Port                   float64  `json:"port"`
	Skew                   float64  `json:"skew"`
	Status                 string   `json:"status"`
	Version                float64  `json:"version"`
}

//Data API Object
type Data struct {
	Timestamp         float64 `json:"timestamp"`
	Value             float64 `json:"value"`
	Count             float64 `json:"count"`
	Counter           float64 `json:"counter"`
	Counter2          float64 `json:"counter2"`
	Counter2Stddev    float64 `json:"counter2_stddev"`
	CounterStddev     float64 `json:"counter_stddev"`
	Derivative        float64 `json:"derivative"`
	Derivative2       float64 `json:"derivative2"`
	Derivative2Stddev float64 `json:"derivative2_stddev"`
	DerivativeStddev  float64 `json:"derivative_stddev"`
	Stddev            float64 `json:"stddev"`
}

//Check API Object
type Check struct {
	Active      bool        `json:"_active"`
	Broker      string      `json:"_broker"`
	CheckBundle string      `json:"_check_bundle"`
	CheckUuid   string      `json:"_check_uuid"`
	Cid         string      `json:"_cid"`
	Details     interface{} `json:"_details"` // the map needs to be verified
}

//Check Bundle API Object
type CheckBundle struct {
	CheckUuids            []interface{}     `json:"_check_uuids"`
	Checks                []interface{}     `json:"_checks"`
	Cid                   string            `json:"_cid"`
	Created               float64           `json:"_created"`
	LastModified          float64           `json:"_last_modified"`
	LastModifiedBy        string            `json:"_last_modified_by"`
	ReverseConnectionUrls []interface{}     `json:"_reverse_connection_urls"`
	Brokers               []interface{}     `json:"brokers"`
	Config                map[string]string `json:"config"`
	DisplayName           string            `json:"display_name"`
	Metrics               []interface{}     `json:"metrics"`
	Notes                 string            `json:"notes"`
	Period                float64           `json:"period"`
	Status                string            `json:"status"`
	Tags                  []interface{}     `json:"tags"`
	Target                string            `json:"string"`
	TimeOut               float64           `json:"timeout"`
	Type                  string            `json:"type"`
}

//CheckMove API Object
type CheckMove struct {
	Broker    string  `json:"_broker"`
	Cid       string  `json:"_cid"`
	Error     string  `json:"_error"`
	Status    string  `json:"_status"`
	CheckID   float64 `json:"check_id"`
	NewBroker string  `json:"new_broker"`
}

//Circonus Usage struct
type CnsUsage struct {
	Limit float64 `json:"_limit"`
	Unit  string  `json:"_type"`
	Used  int64   `json:"_used"`
}

// Circonus Invite struct
type CnsInvite struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// Metric API Object
type Metric struct {
	Active      bool          `json:"_active"`
	Check       string        `json:"_check"`
	CheckActive bool          `json:"_check_active"`
	CheckBundle string        `json:"_check_bundle"`
	CheckTags   []string      `json:"_check_tags"`
	CheckUUID   string        `json:"_check_uuid"`
	Cid         string        `json:"_cid"`
	Histogram   string        `json:"_histogram"`
	MetricName  string        `json:"_metric_name"`
	MetricType  string        `json:"_metric_type"`
	Tags        []interface{} `json:"tags"`
	Units       interface{}   `json:"units"`
}

// User Role struct
type Role struct {
	Role string `json:"role"`
	User string `json:"user"`
}

// Contact Info struct
type ContactInfo struct {
	SMS  string `json:"sms"`
	XMPP string `json:"xmpp"`
}

//User API Object
type User struct {
	Cid         string      `json:"_cid"`
	ContactInfo ContactInfo `json:"contact_info"`
	Email       string      `json:"email"`
	FirstName   string      `json:"firstname"`
	LastName    string      `json:"lastname"`
}

// Graphs are the workhorse of visualizations in the system. This endpoint allows mass creation,
// editing and removal of graphs. Unlike other systems graphs do not store the data with them,
// so a graph created today or a month from now will still show the same data.
type Graph struct {
	Cid               string        `json:"_cid"`
	AccessKeys        []interface{} `json:"access_keys"`
	Composites        []interface{} `json:"composites"`
	Datapoints        []interface{} `json:"datapoints"`
	Description       string        `json:"description"`
	Guides            []interface{} `json:"guides"`
	LineStyle         interface{}   `json:"line_style"`
	LogarithmicLeftY  interface{}   `json:"logarithmic_left_y"`
	LogarithmicRightY interface{}   `json:"logarithmic_right_y"`
	MaxLeftY          interface{}   `json:"max_left_y"`
	MaxRightY         interface{}   `json:"max_right_y"`
	MetricClusters    []interface{} `json:"metric_clusters"`
	MinLeftY          interface{}   `json:"min_left_y"`
	MinRightY         interface{}   `json:"min_right_y"`
	Notes             interface{}   `json:"notes"`
	Style             interface{}   `json:"style"`
	Tags              []interface{} `json:"tags"`
	Title             string        `json:"title"`
}

// DataPoint struct (graph)
type DataPoint struct {
	Alpha         string      `json:"alpha"`
	Axis          string      `json:"axis"`
	Caql          string      `json:"caql"`
	CheckID       string      `json:"check_id"`
	Color         string      `json:"color"`
	DataFormula   string      `json:"data_formula"`
	Derive        bool        `json:"derive"`
	Hidden        bool        `json:"hidden"`
	LegendFormula string      `json:"legend_formula"`
	MetricName    string      `json:"metric_name"`
	MetricType    string      `json:"metric_type"`
	Name          string      `json:"name"`
	Stack         interface{} `json:"stack"`
}

//AccessKey struct (graph)
type AccessKey struct {
	Width          int    `json:"width"`
	LockShowTimes  bool   `json:"lock_show_times"`
	Nickname       string `json:"nickname"`
	Active         bool   `json:"active"`
	Legend         bool   `json:"legend"`
	Height         int    `json:"height"`
	LockRangeStart int    `json:"lock_range_start"`
	Key            string `json:"key"`
	LockRangeEnd   int    `json:"lock_range_end"`
	YLabels        bool   `json:"y_labels"`
	XLabels        bool   `json:"x_labels"`
	Title          bool   `json:"title"`
	LockDate       bool   `json:"lock_date"`
	LockMode       string `json:"lock_mode"`
	LockZoom       string `json:"lock_zoom"`
}

// Schedule a maintenance window for your account, check bundle, rule set or host.
type Maintenance struct {
	Cid        string   `json:"_cid"`
	Item       string   `json:"item"`
	Notes      string   `json:"notes"`
	Severities []string `json:"severities"`
	Start      float64  `json:"start"`
	Stop       float64  `json:"stop"`
	Tags       []string `json:"tags"`
	Type       string   `json:"type"`
}

// A metric cluster is a cluster of metrics defined by a set of queries.
type MetricCluster struct {
	Cid         string `json:"_cid"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Queries     []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"queries"`
	Tags            []string `json:"tags"`
	MatchingMetrics []string `json:"_matching_metrics"`
}

// Rule sets let you define a collection of rules to apply to a given metric. The rules are
// processed in order, and the first one to be found in violation generates an alert, stopping
// further processing. Each rule has a severity which links to a contact group or groups to be
// notified about the alert.
type RuleSet struct {
	Cid           string `json:"_cid"`
	Check         string `json:"check"`
	ContactGroups struct {
			      Num1 []string `json:"1"`
			      Num2 []string `json:"2"`
			      Num3 []string `json:"3"`
			      Num4 []string `json:"4"`
			      Num5 []string `json:"5"`
		      } `json:"contact_groups"`
	Derive     interface{}   `json:"derive"`
	Link       interface{}   `json:"link"`
	MetricName string        `json:"metric_name"`
	MetricType string        `json:"metric_type"`
	Notes      string        `json:"notes"`
	Parent     interface{}   `json:"parent"`
	Rules      []interface{} `json:"rules"`
}

// Rule struct (RuleSet)
type Rule struct {
	Criteria           string      `json:"criteria"`
	Severity           int         `json:"severity"`
	Transform          interface{} `json:"transform"`
	TransformOptions   interface{} `json:"transform_options"`
	TransformSelection interface{} `json:"transform_selection"`
	Value              string      `json:"value"`
	Wait               int         `json:"wait"`
	WindowingDuration  int         `json:"windowing_duration"`
	WindowingFunction  interface{} `json:"windowing_function"`
}

// Rule Set Groups allow you to group together rule sets and trigger alerts based on combinations
// of those rule sets faulting, for example triggering a further alert if a certain number of the
// rule sets themselves are faulting or if a explicit combination of the rulesets are
// in an specific fault state
type RuleSetGroup struct {
	Cid           string `json:"_cid"`
	ContactGroups struct {
			      Num1 []string `json:"1"`
			      Num2 []string `json:"2"`
			      Num3 []string `json:"3"`
			      Num4 []string `json:"4"`
			      Num5 []string `json:"5"`
		      } `json:"contact_groups"`
	Formulas []struct {
		Expression    int `json:"expression"`
		RaiseSeverity int `json:"raise_severity"`
		Wait          int `json:"wait"`
	} `json:"formulas"`
	Name              string `json:"name"`
	RuleSetConditions []struct {
		MatchingSeverities []string `json:"matching_severities"`
		RuleSet            string   `json:"rule_set"`
	} `json:"rule_set_conditions"`
}

// String identifier
type Tag struct {
	Cid string `json:"_cid"`
}

// Templates are a means to setup a mass number of checks quickly through both the API and UI.
// A master host and check bundles are items already being collected by the system,
// linking them to a template and then adding new hosts will apply those checks with their
// same settings across the list of new targets. Hosts and check bundles can then be removed,
// deactivated and unlinked allowing flexibility to work on one off servers or remove them from
// your monitoring infrastructure.
type Template struct {
	Cid            string `json:"_cid"`
	LastModified   int    `json:"_last_modified"`
	LastModifiedBy string `json:"_last_modified_by"`
	CheckBundles   []struct {
		BundleID string `json:"bundle_id"`
		Name     string `json:"name"`
	} `json:"check_bundles"`
	Hosts      []string `json:"hosts"`
	MasterHost string   `json:"master_host"`
	Name       string   `json:"name"`
	Notes      string   `json:"notes"`
	Status     string   `json:"status"`
	Tags       []string `json:"tags"`
}

// Worksheets are simply a collection of graphs and allow quick correlation across them.
type WorkSheet struct {
	Cid         string `json:"_cid"`
	Description string `json:"description"`
	Favorite    bool   `json:"favorite"`
	Graphs      []struct {
		Graph string `json:"graph"`
	} `json:"graphs"`
	Notes        string `json:"notes"`
	SmartQueries []struct {
		Name  string   `json:"name"`
		Order []string `json:"order"`
		Query string   `json:"query"`
	} `json:"smart_queries"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

//Provides API access for creating, reading, updating and deleting Dashboards.
type Dashboard struct {
	Created int `json:"_created"`
	Options struct {
			FullscreenHideTitle bool          `json:"fullscreen_hide_title"`
			TextSize            int           `json:"text_size"`
			Linkages            []interface{} `json:"linkages"`
			AccessConfigs       []struct {
				FullscreenHideTitle bool   `json:"fullscreen_hide_title"`
				TextSize            int    `json:"text_size"`
				ScaleText           bool   `json:"scale_text"`
				Nickname            string `json:"nickname"`
				BlackDash           bool   `json:"black_dash"`
				Fullscreen          bool   `json:"fullscreen"`
				SharedID            string `json:"shared_id"`
				Enabled             bool   `json:"enabled"`
			} `json:"access_configs"`
			ScaleText bool `json:"scale_text"`
			HideGrid  bool `json:"hide_grid"`
		} `json:"options"`
	Cid            string `json:"_cid"`
	Shared         bool   `json:"shared"`
	Active         bool   `json:"_active"`
	Title          string `json:"title"`
	DashboardUUID  string `json:"_dashboard_uuid"`
	AccountDefault bool   `json:"account_default"`
	GridLayout     struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"grid_layout"`
	CreatedBy    string        `json:"_created_by"`
	LastModified int           `json:"_last_modified"`
	Widgets      []interface{} `json:"widgets"`
}

//hashes defining the widgets for this Dashboard
type Widget struct {
	Width    float64 `json:"width"`
	Name     string  `json:"name"`
	Active   bool    `json:"active"`
	Origin   string  `json:"origin"`
	Height   float64 `json:"height"`
	Settings struct {
			 KeyInline  bool   `json:"key_inline"`
			 Period     string `json:"period"`
			 KeySize    string `json:"key_size"`
			 HideYaxis  bool   `json:"hide_yaxis"`
			 GraphID    string `json:"graph_id"`
			 AccountID  string `json:"account_id"`
			 ShowFlags  bool   `json:"show_flags"`
			 DateWindow string `json:"date_window"`
			 HideXaxis  bool   `json:"hide_xaxis"`
			 KeyWrap    bool   `json:"key_wrap"`
			 Label      string `json:"label"`
			 KeyLoc     string `json:"key_loc"`
			 GraphTitle string `json:"_graph_title"`
			 Realtime   bool   `json:"realtime"`
		 } `json:"settings"`
	Type     string `json:"type"`
	WidgetID string `json:"widget_id"`
}

//Worksheet API Object
type Worksheet struct {
	Cid          string        `json:"_cid"`
	Description  string        `json:"description"`
	Favorite     bool          `json:"favorite"`
	Graphs       []interface{} `json:"graphs"`
	Notes        string        `json:"notes"`
	SmartQueries []struct {
		Name  string        `json:"name"`
		Order []interface{} `json:"order"`
		Query string        `json:"query"`
	} `json:"smart_queries"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

//CAQL data struct
type CaqlData struct {
	Data []struct {
		TimeStamp float64   `json:"time_stamp"`
		DataPoint []float64 `json:"data_point"`
	} `json:"_data"`
	End    float64 `json:"_end"`
	Period float64 `json:"_period"`
	Query  string  `json:"_query"`
	Start  float64 `json:"_start"`
}

//Check Bundle Metrics API Object
type CheckBundleMetrics struct {
	Cid     string `json:"_cid"`
	Metrics []struct {
		Name   string        `json:"name"`
		Result string        `json:"result"`
		Status string        `json:"status"`
		Tags   []interface{} `json:"tags"`
		Type   string        `json:"type"`
		Units  interface{}   `json:"units"`
	} `json:"metrics"`
}

// Contact Group API Object
type ContactGroup struct {
	Cid               string  `json:"_cid"`
	LastModified      float64 `json:"_last_modified"`
	LastModifiedBy    string  `json:"_last_modified_by"`
	AggregationWindow float64 `json:"aggregation_window"`
	AlertFormats      struct {
				  LongMessage  interface{} `json:"long_message"`
				  LongSubject  interface{} `json:"long_subject"`
				  LongSummary  interface{} `json:"long_summary"`
				  ShortMessage interface{} `json:"short_message"`
				  ShortSummary interface{} `json:"short_summary"`
			  } `json:"alert_formats"`
	Contacts struct {
				  External []struct {
					  ContactInfo string `json:"contact_info"`
					  Method      string `json:"method"`
				  } `json:"external"`
				  Users []struct {
					  ContactInfo string `json:"_contact_info"`
					  Method      string `json:"method"`
					  User        string `json:"user"`
				  } `json:"users"`
			  } `json:"contacts"`
	Escalations []interface{} `json:"escalations"`
	Name        string        `json:"name"`
	Reminders   []int         `json:"reminders"`
	Tags        []string      `json:"tags"`
}
