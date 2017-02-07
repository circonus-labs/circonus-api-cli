// Copyright 2016 Alem Abreha <alem.abreha@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
package circonusapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	custom_url "net/url" // Importing net/url as custom_url since "url" is used as variable
	// in this code
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//HTTPCall function takes a url, header setting, http call method and data payload and makes HTTP
// request (CRUD operation)
func HTTPCall(payload []byte, method string, url string) (*http.Response, error) {

	//because http.NewRequest takes io.Reader interface convert payload into Buffer type since
	// this type implements Read function >>> func (b *Buffer) Read(p []byte) (n int, err error)

	byte_buffer := bytes.NewBuffer(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, byte_buffer)
	if err != nil {
		fmt.Println(err)
		return &http.Response{}, err
	} else {
		req.Header.Set("X-Circonus-Auth-Token", os.Getenv("CIRONUS_API_TOKEN"))
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Circonus-App-Name", os.Getenv("CIRCONUS_APP_NAME"))

		return client.Do(req)
	}
}

// GetCns function takes a filter interface and object type string , returns an interface that can be marshalled into Circonus objects
func GetCns(filter interface{}, object string) ([]byte, error) {

	URL, err := UrlMaker(object, filter)
	if err != nil {
		fmt.Println(err)
		os.Exit(251)
	}
	result, err := CirconusCall(URL, "GET", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err

	} else {
		return result, nil
	}
}

// GetData function takes a DataFilter struct and returns lists of timestamp : metrics value
// pairs that match the search/filter
func GetData(filter DataFilter) (interface{}, error) {

	URL, err := UrlMaker("data", filter)
	if err != nil {
		fmt.Println(err)
		os.Exit(251)
	}
	result, err := CirconusCall(URL, "GET", nil)

	if err != nil {
		fmt.Println(err)
		return nil, err

	} else {
		//Output gets marshalled into an Broker struct
		var data map[string][]Data

		json.Unmarshal(result, &data)
		return data["data"], nil
	}

}

// CreateCns function takes []byte payload and object_url string, returns []byte
func CreateCns(payload []byte, object_url string) ([]byte, error) {
	URL := os.Getenv("CIRCONUS_API_URL") + object_url
	result, err := CirconusCall(URL, "POST", payload)

	if err != nil {
		fmt.Println(err)
		return nil, err

	} else {

		return result, nil
	}

}

// DeleteCns function takes oid string and object_url string, returns []byte
func DeleteCns(oid string, object_url string) ([]byte, error) {

	URL := os.Getenv("CIRCONUS_API_URL") + object_url + "/" + oid

	result, err := CirconusCall(URL, "DELETE", nil)

	if err != nil {
		fmt.Println(err)
		return nil, err

	} else {

		return result, nil
	}

}

// UpdateCns function takes oid string, object_url string, payload []byte and returns []byte
func UpdateCns(oid string, object_url string, payload []byte) ([]byte, error) {

	URL := os.Getenv("CIRCONUS_API_URL") + object_url + "/" + custom_url.QueryEscape(oid)

	result, err := CirconusCall(URL, "PUT", payload)

	if err != nil {
		fmt.Println(err)
		return nil, err

	} else {

		return result, nil
	}

}

// UpdateTemplate function takes oid string, object_url string, host_removal_action string, bundle_removal_action and payload []byte and returns []byte
func UpdateTemplate(oid string, object_url string, host_removal_action string, bundle_removal_action string, payload []byte) ([]byte, error) {

	URL := os.Getenv("CIRCONUS_API_URL") + object_url + "/" + oid + "?"

	if host_removal_action != "" {
		URL += "host_removal_action=" + host_removal_action + "&"
	}
	if bundle_removal_action != "" {
		URL += "bundle_removal_action=" + bundle_removal_action + "&"
	}
	result, err := CirconusCall(URL, "PUT", payload)

	if err != nil {
		fmt.Println(err)
		return nil, err

	} else {

		return result, nil
	}

}

//CinconusCall function is a wrapper ontop of HTTP call that extracts only the message body from
// HTTP request.
func CirconusCall(url string, method string, payload []byte) ([]byte, error) {

	resp, err := HTTPCall(payload, method, url)

	if err != nil {
		return nil, err

	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		return body, nil

	}

}

// UrlMaker function generates the URL string for CirconusCall depending on the API object type
// thats is being queried.
func UrlMaker(item string, filter interface{}) (string, error) {
	var url string

	CirconusURL := os.Getenv("CIRCONUS_API_URL")
	switch strings.ToLower(item) {

	case "account":
		account_filter := filter.(AccountFilter)
		if account_filter.AccountID == 0 {
			url = CirconusURL + "account?"
			if account_filter.Name != "" {
				url += "f_name=" + custom_url.QueryEscape(account_filter.Name) + "&"
			}
		} else {
			url = CirconusURL + "account/" + strconv.FormatFloat(account_filter.AccountID, 'f', -1, 64)
		}
	case "alert":
		alert_filter := filter.(AlertFilter)
		if alert_filter.AlertID != 0 {
			url = CirconusURL + "alert/" + strconv.FormatFloat(alert_filter.AlertID, 'f', -1, 64)
		} else {
			url = CirconusURL + "alert?"
			if alert_filter.TagsHas != "" {
				url += "f__tags_has=" + custom_url.QueryEscape(alert_filter.TagsHas) + "&"
			}
			if alert_filter.OccurredOnLt != 0 {
				url += "f__occurred_on_lt=" + strconv.FormatFloat(alert_filter.OccurredOnLt, 'f', -1, 64) + "&"
			}
			if alert_filter.OccurredOnGt != 0 {
				url += "f__occurred_on_gt=" + strconv.FormatFloat(alert_filter.OccurredOnGt, 'f', -1, 64) + "&"
			}
			if alert_filter.Severity != 0 {
				url += "f__severity=" + strconv.FormatFloat(alert_filter.Severity, 'f', -1, 64) + "&"
			}
			if alert_filter.Check != 0 {
				url += "f__check=/check/" + strconv.FormatFloat(alert_filter.Check, 'f', -1, 64) + "&"
			}
			if alert_filter.MetricName != "" {
				url += "f__metric_name=" + custom_url.QueryEscape(alert_filter.MetricName) + "&"
			}
			if alert_filter.OccurredOnLe != 0 {
				url += "f__occurred_on_le=" + strconv.FormatFloat(alert_filter.OccurredOnLe, 'f', -1, 64) + "&"
			}
			if alert_filter.OccurredOnGe != 0 {
				url += "f__occurred_on_ge=" + strconv.FormatFloat(alert_filter.OccurredOnGe, 'f', -1, 64) + "&"
			}
		}
	case "user":
		user_filter := filter.(UserFilter)
		if user_filter.UserID == 0 {
			url = CirconusURL + "user?"
			if user_filter.FirstName != "" {
				url += "f_firstname=" + custom_url.QueryEscape(user_filter.FirstName) + "&"
			}
			if user_filter.LastName != "" {
				url += "f_lastname=" + custom_url.QueryEscape(user_filter.LastName) + "&"
			}
			if user_filter.Email != "" {
				url += "f_email=" + custom_url.QueryEscape(user_filter.Email) + "&"
			}
		} else {
			url = CirconusURL + "user/" + strconv.FormatFloat(user_filter.UserID, 'f', -1, 64)
		}

	case "annotation":
		annotation_filter := filter.(AnnotationFilter)
		if annotation_filter.AnnotationID != 0 {
			url = CirconusURL + "annotation/" + strconv.FormatFloat(annotation_filter.AnnotationID, 'f', -1, 64)
		} else {
			url = CirconusURL + "annotation?"
			if annotation_filter.Category != "" {
				url += "f_category=" + annotation_filter.Category + "&"
			}
			if annotation_filter.CategoryWildcard != "" {
				url += "f_category_wildcard=*" + custom_url.QueryEscape(annotation_filter.CategoryWildcard) + "*&"
			}
			if annotation_filter.Description != "" {
				url += "f_description=" + custom_url.QueryEscape(annotation_filter.Description) + "&"
			}
			if annotation_filter.DescriptionWildcard != "" {
				url += "f_description_wildcard=*" + custom_url.QueryEscape(annotation_filter.DescriptionWildcard) + "*&"
			}
			if annotation_filter.StartGt != 0 {
				url += "f_start_gt=" + strconv.FormatFloat(annotation_filter.StartGt, 'f', -1, 64) + "&"
			}
			if annotation_filter.StartGe != 0 {
				url += "f_start_ge=" + strconv.FormatFloat(annotation_filter.StartGe, 'f', -1, 64) + "&"
			}
			if annotation_filter.StartLe != 0 {
				url += "f_start_le=" + strconv.FormatFloat(annotation_filter.StartLe, 'f', -1, 64) + "&"
			}
			if annotation_filter.StartLt != 0 {
				url += "f_start_lt=" + strconv.FormatFloat(annotation_filter.StartLt, 'f', -1, 64) + "&"
			}
			if annotation_filter.StopLt != 0 {
				url += "f_stop_lt=" + strconv.FormatFloat(annotation_filter.StopLt, 'f', -1, 64) + "&"
			}
			if annotation_filter.StopLe != 0 {
				url += "f_stop_le=" + strconv.FormatFloat(annotation_filter.StopLe, 'f', -1, 64) + "&"
			}
			if annotation_filter.StopGe != 0 {
				url += "f_stop_ge=" + strconv.FormatFloat(annotation_filter.StopGe, 'f', -1, 64) + "&"
			}
			if annotation_filter.StopGt != 0 {
				url += "f_stop_gt=" + strconv.FormatFloat(annotation_filter.StopGt, 'f', -1, 64) + "&"
			}
		}

	case "check":
		check_filter := filter.(CheckFilter)
		if check_filter.CheckID == 0 {
			url = CirconusURL + "check?"
			if check_filter.CheckBundleId != 0 {
				url += "f__check_bundle=/check_bundle/" + strconv.FormatFloat(check_filter.CheckBundleId, 'f', -1, 64) + "&"
			}
			if check_filter.CheckUuid != "" {
				url += "f__check_uuid=" + check_filter.CheckUuid + "&"
			}
		} else {
			url = CirconusURL + "check/" + strconv.FormatFloat(check_filter.CheckID, 'f', -1, 64)
		}

	case "check_bundle":
		bundle_filter := filter.(CheckBundleFilter)
		if bundle_filter.CheckBundleID == 0 {
			url = CirconusURL + "check_bundle?"
			if bundle_filter.BrokersHas != 0 {
				url += "f_broker_has=/broker/" + strconv.FormatFloat(bundle_filter.BrokersHas, 'f', -1, 64) + "&"
			}
			if bundle_filter.ChecksHas != 0 {
				url += "f__checks_has=/check/" + strconv.FormatFloat(bundle_filter.ChecksHas, 'f', -1, 64) + "&"
			}
			if bundle_filter.DisplayName != "" {
				url += "f_display_name=" + custom_url.QueryEscape(bundle_filter.DisplayName) + "&"
			}
			if bundle_filter.TagsHas != "" {
				url += "f_tags_has=" + bundle_filter.TagsHas + "&"
			}
			if bundle_filter.Target != "" {
				url += "f_target=" + bundle_filter.Target + "&"
			}
			if bundle_filter.Type != "" {
				url += "f_type=" + bundle_filter.Type + "&"
			}
			if bundle_filter.TargetLike != "" {
				url += "f_target_wildcard=*" + custom_url.QueryEscape(bundle_filter.TargetLike) + "*&"
			}
			if bundle_filter.DisplayNameLike != "" {
				url += "f_display_name_wildcard=*" + custom_url.QueryEscape(bundle_filter.DisplayNameLike) + "*&"
			}
		} else {
			url = CirconusURL + "check_bundle/" + strconv.FormatFloat(bundle_filter.CheckBundleID, 'f', -1, 64)
		}

	case "check_bundle_metrics":
		check_bundle_metrics_filter := filter.(CheckBundleMetricsFilter)
		url = CirconusURL + "check_bundle_metrics/" + strconv.FormatFloat(check_bundle_metrics_filter.CheckBundleId, 'f', -1, 64)

	case "broker":
		broker_filter := filter.(BrokerFilter)
		if broker_filter.BrokerID != 0 {
			url = CirconusURL + "broker/" + strconv.FormatFloat(broker_filter.BrokerID, 'f', -1, 64)
		} else {
			url = CirconusURL + "broker?"
			if broker_filter.Name != "" {
				url += "f__name=" + custom_url.QueryEscape(broker_filter.Name) + "&"
			}
			if broker_filter.Type != "" {
				url += "f__type=" + broker_filter.Type + "&"
			}
		}

	case "checkmove":
		checkmove_filter := filter.(CheckMoveFilter)
		if checkmove_filter.CheckMoveID != 0 {
			url = CirconusURL + "check_move/" + strconv.FormatFloat(checkmove_filter.CheckMoveID, 'f', -1, 64)
		} else {
			url = CirconusURL + "check_move"
		}

	case "contact_group":
		contact_group_filter := filter.(ContactGroupFilter)
		if contact_group_filter.ContactGroupId != 0 {
			url = CirconusURL + "contact_group/" + strconv.FormatFloat(contact_group_filter.ContactGroupId, 'f', -1, 64)
		} else {
			url = CirconusURL + "contact_group?"
			if contact_group_filter.Name != "" {
				url += "f_name=" + custom_url.QueryEscape(contact_group_filter.Name) + "&"
			}
			if contact_group_filter.NameLike != "" {
				url += "f_name_wildcard=*" + custom_url.QueryEscape(contact_group_filter.NameLike) + "*&"
			}
			if contact_group_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(contact_group_filter.TagsHas) + "&"
			}
		}

	case "data":
		data_filter := filter.(DataFilter)
		url = CirconusURL + "data/" + strconv.FormatFloat(data_filter.CheckId, 'f', -1, 64) + "_" + custom_url.QueryEscape(data_filter.MetricName) +
			"?format=object" + "&type=" + data_filter.Type +
			"&period=" + strconv.FormatFloat(data_filter.Period, 'f', -1, 64) +
			"&start=" + strconv.FormatFloat(data_filter.Start, 'f', -1, 64) +
			"&end=" + strconv.FormatFloat(data_filter.Stop, 'f', -1, 64)

	case "dashboard":
		dashboard_filter := filter.(DashboardFilter)
		if dashboard_filter.DashboardId != 0 {
			url = CirconusURL + "dashboard/" + strconv.FormatFloat(dashboard_filter.DashboardId, 'f', -1, 64)
		} else {
			url = CirconusURL + "dashboard?"
			if dashboard_filter.TitleWildcard != "" {
				url += "f_title_wildcard=*" + custom_url.QueryEscape(dashboard_filter.TitleWildcard) + "*&"
			}
			if dashboard_filter.Title != "" {
				url += "f_title=" + custom_url.QueryEscape(dashboard_filter.Title) + "&"
			}
		}
	case "graph":
		graph_filter := filter.(GraphFilter)
		if graph_filter.GraphID == "" {
			url = CirconusURL + "graph?"
			if graph_filter.Title != "" {
				url += "f_title=" + custom_url.QueryEscape(graph_filter.Title) + "&"
			}
			if graph_filter.TitleWildcard != "" {
				url += "f_title_wildcard=*" + custom_url.QueryEscape(graph_filter.TitleWildcard) + "*&"
			}
			if graph_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(graph_filter.TagsHas) + "&"
			}
		} else {
			url = CirconusURL + "graph/" + graph_filter.GraphID
		}

	case "metric_cluster":
		metriccluster_filter := filter.(MetricClusterFilter)
		if metriccluster_filter.MetricClusterID == 0 {
			url = CirconusURL + "metric_cluster?"
			if metriccluster_filter.TagsHas != "" {
				url += "f_tags_has=" + metriccluster_filter.TagsHas + "&"
			}
			if metriccluster_filter.NameWildcard != "" {
				url += "f_name_wildcard=*" + custom_url.QueryEscape(metriccluster_filter.NameWildcard) + "*&"
			}
			if metriccluster_filter.Name != "" {
				url += "f_name=" + custom_url.QueryEscape(metriccluster_filter.Name) + "&"
			}
		} else {
			url = CirconusURL + "metric_cluster/" + strconv.FormatFloat(metriccluster_filter.MetricClusterID, 'f', -1, 64) + "?extra=_matching_metrics"
		}

	case "metric":
		metric_filter := filter.(MetricFilter)
		if metric_filter.MetricID == "" {
			url = CirconusURL + "metric?"
			if metric_filter.CheckBundleID != "" {
				url += "f__check_bundle=/check_bundle/" + custom_url.QueryEscape(metric_filter.CheckBundleID) + "&"
			}
			if metric_filter.CheckID != "" {
				url += "f__check=/check/" + custom_url.QueryEscape(metric_filter.CheckID) + "&"
			}
			if metric_filter.CheckUuid != "" {
				url += "f__check_uuid=" + custom_url.QueryEscape(metric_filter.CheckUuid) + "&"
			}
			if metric_filter.Name != "" {
				url += "f__metric_name=" + custom_url.QueryEscape(metric_filter.Name) + "&"
			}
			if metric_filter.NameWildcard != "" {
				url += "f__metric_name_wildcard=*" + custom_url.QueryEscape(metric_filter.NameWildcard) + "*&"
			}
			if metric_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(metric_filter.TagsHas) + "&"
			}
			if metric_filter.CheckTagsHas != "" {
				url += "f__check_tags_has=" + custom_url.QueryEscape(metric_filter.CheckTagsHas) + "&"
			}

		} else {
			url = CirconusURL + "metric/" + custom_url.QueryEscape(metric_filter.MetricID)
		}

	case "maintenance":
		maintenance_filter := filter.(MaintenanceFilter)
		if maintenance_filter.MaintenanceID == 0 {
			url = CirconusURL + "maintenance?"
			if maintenance_filter.Item != "" {
				url += "f_item=" + custom_url.QueryEscape(maintenance_filter.Item) + "&"
			}
			if maintenance_filter.ItemLike != "" {
				url += "f_item_wildcard=*" + custom_url.QueryEscape(maintenance_filter.ItemLike) + "*&"
			}
			if maintenance_filter.Type != "" {
				url += "f_type=" + custom_url.QueryEscape(maintenance_filter.Type) + "&"
			}
			if maintenance_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(maintenance_filter.TagsHas) + "&"
			}
			if maintenance_filter.StartGe != 0 {
				url += "f_start_ge=" + strconv.FormatFloat(maintenance_filter.StartGe, 'f', -1, 64) + "&"
			}
			if maintenance_filter.StartGt != 0 {
				url += "f_start_gt=" + strconv.FormatFloat(maintenance_filter.StartGt, 'f', -1, 64) + "&"
			}
			if maintenance_filter.StopLe != 0 {
				url += "f_stop_le=" + strconv.FormatFloat(maintenance_filter.StopLe, 'f', -1, 64) + "&"
			}
			if maintenance_filter.StopLt != 0 {
				url += "f_stop_lt=" + strconv.FormatFloat(maintenance_filter.StopLt, 'f', -1, 64) + "&"
			}
		} else {
			url = CirconusURL + "maintenance/" + strconv.FormatFloat(maintenance_filter.MaintenanceID, 'f', -1, 64)
		}
	case "tag":
		tag_filter := filter.(TagFilter)
		if tag_filter.TagId != "" {
			url = CirconusURL + "tag/" + custom_url.QueryEscape(tag_filter.TagId)
		} else {
			url = CirconusURL + "tag"
		}
	case "rule_set":
		ruleset_filter := filter.(RuleSetFilter)
		if ruleset_filter.Name == "" {
			url = CirconusURL + "rule_set?"
			if ruleset_filter.Check != 0 {
				url += "f_check=/check/" + strconv.FormatFloat(ruleset_filter.Check, 'f', -1, 64)
			}
		} else {
			url = CirconusURL + "rule_set/" + custom_url.QueryEscape(ruleset_filter.Name)
		}
	case "rule_set_group":
		rulesetgroup_filter := filter.(RuleSetGroupFilter)
		if rulesetgroup_filter.RuleSetGroupId == 0 {
			url = CirconusURL + "rule_set_group?"
			if rulesetgroup_filter.NameWildcard != "" {
				url += "f_name_wildcard=*" + custom_url.QueryEscape(rulesetgroup_filter.NameWildcard) + "*&"
			}
			if rulesetgroup_filter.Name != "" {
				url += "f_name=" + custom_url.QueryEscape(rulesetgroup_filter.Name) + "&"
			}
			if rulesetgroup_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(rulesetgroup_filter.TagsHas) + "&"
			}
		} else {
			url = CirconusURL + "rule_set_group/" + strconv.FormatFloat(rulesetgroup_filter.RuleSetGroupId, 'f', -1, 64)
		}
	case "worksheet":
		worksheet_filter := filter.(WorksheetFilter)
		if worksheet_filter.WorksheetId == "" {
			url = CirconusURL + "worksheet?"
			if worksheet_filter.TitleWildcard != "" {
				url += "f_title_wildcard=*" + custom_url.QueryEscape(worksheet_filter.TitleWildcard) + "*&"
			}
			if worksheet_filter.Title != "" {
				url += "f_title=" + custom_url.QueryEscape(worksheet_filter.Title) + "&"
			}
			if worksheet_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(worksheet_filter.TagsHas) + "&"
			}
		} else {
			url = CirconusURL + "worksheet/" + custom_url.QueryEscape(worksheet_filter.WorksheetId)
		}
	case "template":
		template_filter := filter.(TemplateFilter)
		if template_filter.TemplateId == 0 {
			url = CirconusURL + "template?"
			if template_filter.NameWildcard != "" {
				url += "f_name_wildcard=*" + custom_url.QueryEscape(template_filter.NameWildcard) + "*&"
			}
			if template_filter.Name != "" {
				url += "f_name=" + custom_url.QueryEscape(template_filter.Name) + "&"
			}
			if template_filter.TagsHas != "" {
				url += "f_tags_has=" + custom_url.QueryEscape(template_filter.TagsHas) + "&"
			}
		} else {
			url = CirconusURL + "template/" + strconv.FormatFloat(template_filter.TemplateId, 'f', -1, 64)
		}
	case "caql":
		caql_filter := filter.(CaqlDataFilter)
		url = CirconusURL + "caql?query=" + custom_url.QueryEscape(caql_filter.Query) + "&start=" + strconv.FormatFloat(caql_filter.Start, 'f', -1, 64) +
			"&end=" + strconv.FormatFloat(caql_filter.End, 'f', -1, 64) + "&period=" + strconv.FormatFloat(caql_filter.Period, 'f', -1, 64)
	default:
		url = CirconusURL
	}

	return url, nil
}
