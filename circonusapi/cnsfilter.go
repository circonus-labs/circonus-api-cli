// Copyright 2016 Alem Abreha <alem.abreha@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
package circonusapi

//structs used for filtering API objects

type AnnotationFilter struct {
	AnnotationID        float64 `json:"annotation_id"`
	Description         string  `json:"description"`
	DescriptionWildcard string  `json:"description_like"`
	Category            string  `json:"category"`
	CategoryWildcard    string  `json:"category_like"`
	StartGt             float64 `json:"start_gt"`
	StartGe             float64 `json:"start_ge"`
	StartLt             float64 `json:"start_lt"`
	StartLe             float64 `json:"start_le"`
	StopLt              float64 `json:"stop_lt"`
	StopLe              float64 `json:"stop_le"`
	StopGe              float64 `json:"stop_ge"`
	StopGt              float64 `json:"stop_gt"`
}

type AccountFilter struct {
	AccountID float64 `json:"account_id"`
	Name      string  `json:"name"`
}

type AlertFilter struct {
	AlertID      float64 `json:"alert_id"`
	Severity     float64 `json:"severity"`
	Check        float64 `json:"check"`
	MetricName   string  `json:"metric_name"`
	OccurredOnLt float64 `json:"occurred_on_lt"`
	OccurredOnLe float64 `json:"occurred_on_le"`
	OccurredOnGt float64 `json:"occurred_on_gt"`
	OccurredOnGe float64 `json:"occurred_on_ge"`
	TagsHas      string  `json:"tags_has"`
}

type UserFilter struct {
	UserID    float64 `json:"user_id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
}

type CheckFilter struct {
	CheckID       float64 `json:"check_id"`
	CheckBundleId float64 `json:"check_bundle_id"`
	CheckUuid     string  `json:"check_uuid"`
}

type CheckBundleFilter struct {
	CheckBundleID   float64 `json:"check_bundle_id"` //known check_bundle cid
	Type            string  `json:"type"`            //to list all check bundles that has a specific type field
	Target          string  `json:"target"`          //check bundles that target a particular server
	TargetLike      string  `json:"target_like"`
	DisplayName     string  `json:"display_name"` // check bundle with a particular name
	DisplayNameLike string  `json:"display_name_like"`
	TagsHas         string  `json:"tags_has"`    //check bundles with a particular tag
	ChecksHas       float64 `json:"checks_has"`  //check bundle that has a particular check in it
	BrokersHas      float64 `json:"brokers_has"` //all check bundles using a particular broker
}

type BrokerFilter struct {
	BrokerID float64 `json:"broker_id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
}

type CheckMoveFilter struct {
	CheckMoveID float64 `json:"check_move_id"`
}

//filtering only in object format for timestamp and value fields
//check = checkID_MetricName (combination of checkID and MetricName)
type DataFilter struct {
	CheckId    float64 `json:"check_id"`
	MetricName string  `json:"metric_name"`
	Type       string  `json:"type"`
	Period     float64 `json:"period"`
	Start      float64 `json:"start"`
	Stop       float64 `json:"stop"`
}

type GraphFilter struct {
	GraphID       string `json:"graph_id"`
	Title         string `json:"title"`
	TitleWildcard string `json:"title_like"`
	TagsHas       string `json:"tags_has"`
}

type DashboardFilter struct {
	DashboardId   float64 `json:"dashboard_id"`
	Title         string  `json:"title"`
	TitleWildcard string  `json:"title_like"`
}

type MaintenanceFilter struct {
	MaintenanceID float64 `json:"maintenance_id"`
	Item          string  `json:"item"`
	ItemLike      string  `json:"item_like"`
	TagsHas       string  `json:"tags_has"`
	StartGe       float64 `json:"start_ge"`
	StartGt       float64 `json:"start_gt"`
	StopLe        float64 `json:"stop_le"`
	StopLt        float64 `json:"stop_lt"`
	Type          string  `json:"type"`
}

type MetricClusterFilter struct {
	MetricClusterID float64 `json:"metric_cluster_id"`
	Name            string  `json:"name"`
	NameWildcard    string  `json:"name_like"`
	TagsHas         string  `json:"tags_has"`
}

type GenericFilter struct {
	OID string `json:"oid"`
}

type TagFilter struct {
	TagId string `json:"tag_id"`
}

type RuleSetFilter struct {
	Name  string  `json:"name"`
	Check float64 `json:"check"`
}

type RuleSetGroupFilter struct {
	RuleSetGroupId float64 `json:"rule_set_group_id"`
	Name           string  `json:"name"`
	NameWildcard   string  `json:"name_like"`
	TagsHas        string  `json:"tags_has"`
}

type WorksheetFilter struct {
	WorksheetId   string `json:"worksheet_id"`
	Title         string `json:"title"`
	TitleWildcard string `json:"title_like"`
	TagsHas       string `json:"tags_has"`
}

type TemplateFilter struct {
	TemplateId   float64 `json:"template_id"`
	Name         string  `json:"name"`
	NameWildcard string  `json:"name_like"`
	TagsHas      string  `json:"tags_has"`
}

type MetricFilter struct {
	MetricID      string `json:"metric_id"`
	Name          string `json:"name"`
	NameWildcard  string `json:"name_like"`
	CheckID       string `json:"check_id"`
	CheckBundleID string `json:"check_bundle_id"`
	CheckUuid     string `json:"check_uuid"`
	TagsHas       string `json:"tags_has"`
	CheckTagsHas  string `json:"check_tags_has"`
}

type CaqlDataFilter struct {
	Query  string  `json:"query"`
	Start  float64 `json:"start"`
	End    float64 `json:"end"`
	Period float64 `json:"period"`
}

type CheckBundleMetricsFilter struct {
	CheckBundleId float64 `json:"check_bundle_id"`
}

type ContactGroupFilter struct {
	ContactGroupId float64 `json:"contact_group_id"`
	Name           string  `json:"name"`
	NameLike       string  `json:"name_like"`
	TagsHas        string  `json:"tags_has"`
}
