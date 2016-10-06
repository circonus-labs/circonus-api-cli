package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/misale/circonus-api-go/circonusapi"
	"io/ioutil"
	"log"
)

// circli Usage
var object = flag.String("object", "", "Circonus Object Type. Possible object types : \n"+
	"\taccount :\tbasic contact details associated with the account and description \n"+
	"\talert :\t\tRepresentation of an Alert that occurred (Readonly) \n"+
	"\tannotation :\tMark important events used for correlation \n"+
	"\tbroker :\tRemote software agent that collects the data from monitored hosts\n"+
	"\tcheck :\t\tIndividual elements of a check bundle (Readonly)\n"+
	"\tcheck_bundle :\tCollection of checks that have the same configuration and target, but "+
	"collected from different brokers\n"+
	"\tcheck_move :\tRequest that a Check be moved from one Broker to another\n"+
	"\tdata :\t\tReadonly endpoint to pull the values of a single data point for a given time "+
	"range\n"+
	"\tdashboard :\tProvides access for creating, reading, updating and deleting Dashboards.\n"+
	"\tgraph :\t\tAllows mass creation, editing and removal of graphs\n"+
	"\tmaintenance :\tSchedule a maintenance window for your account, check bundle, rule set "+
	"or host\n"+
	"\tmetric_cluster :\tA metric cluster is a cluster of metrics defined by a set of queries\n"+
	"\trule_set :\tdefine a collection of rules to apply to a given metric\n"+
	"\trule_set_group :\tAllows togroup together rule sets and trigger alerts based on "+
	"combinations of those rule sets faulting\n"+
	"\ttag :\t\tList all tags in use in your account (don't have any fields, Readonly)\n"+
	"\ttemplate :\tA means to setup a mass number of checks quickly through both the API and "+
	"UI\n"+
	"\tuser :\t\tGet basic information about the users on your account or a single user\n"+
	"\tworksheet :\tCollection of graphs and allow quick correlation across them\n")
var call = flag.String("call", "", "Circonus Call Type. Example : get, list, create, update, delete\n")
var file = flag.String("file", "", "JSON formatted data input file\n")
var oid = flag.String("oid", "", "Circonus object ID. Value of _cid in the API object without the \"/<object_type>/\" prefix. This flag is used with \"update\" and \"delete\" calls\n")
var where = flag.String("where", "", "JSON string used for querying where clause\n")
var host_removal_action = flag.String("host_removal_action", "", "Host removal action used with template update actions: possible values are \"unbind\", \"deactivate\", \"remove\"\n"+
	"\tunbind :\tMarks the check(s) as unbound from the template, you can then modify them as if they were a \"normal\" check\n"+
	"\tdeactivate :\tSets the check(s) as inactive, they will still show up on the interface and you can view historic data for them \n"+
	"\tremove :\tDeletes the check(s) from the system, they will no longer show in the UI and historic data will be gone \n")
var bundle_removal_action = flag.String("bundle_removal_action", "", "Bundle removal action used with template update actions: possible values are \"unbind\", \"deactivate\", \"remove\"\n"+
	"\tunbind :\tMarks the check(s) as unbound from the template, you can then modify them as if they were a \"normal\" check\n"+
	"\tdeactivate :\tSets the check(s) as inactive, they will still show up on the interface and you can view historic data for them \n"+
	"\tremove :\tDeletes the check(s) from the system, they will no longer show in the UI and historic data will be gone \n")

func main() {
	flag.Parse()
	//user commandline input sanity check
	if *where == "" && *file == "" && *call != "list" && *call != "delete" {
		log.Fatal("circli needs either where JSON string with -where flag or JSON file " +
			"with the -file flag")
	}
	//lookup is either where JSON string or contents of JSON file
	var lookup string
	if *where != "" && *call != "list" {
		lookup = *where
	}
	if *file != "" && *call != "list" {
		//verify if file is valid
		file_content, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal("Error while reading ", *file, " : ", err)
		} else {
			lookup = string(file_content)
		}
	}
	if *object == "" {
		log.Fatal("object type is missing, it has to be specified : -object value " +
			"is mandatory")
	}
	if (*oid == "") && (*call == "update" || *call == "delete") {
		log.Fatal("Updating existing objects needs a non-nil value passed to -oid")
	}

	switch *call {

	case "":
		log.Fatal("call type is missing, available call types are : get, list, create, " +
			"update, delete")

	case "get":
		switch *object {
		case "account":
			var account_filter circonusapi.AccountFilter
			err := json.Unmarshal([]byte(lookup), &account_filter)
			// nil_account for verifying unmarshalling
			nil_account := circonusapi.AccountFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if account_filter == nil_account {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(account_filter, "account")
					if err != nil {
						log.Fatal("account get call failure ", err)
					} else {
						// output is marshalled into an Account struct
						if account_filter.AccountID == 0 {
							var data []circonusapi.Account
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Account
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}

					}
				}
			}

		/*
		 */
		case "alert":
			var alert_filter circonusapi.AlertFilter
			err := json.Unmarshal([]byte(lookup), &alert_filter)
			// nil_alert for verifying unmarshalling
			nil_alert := circonusapi.AlertFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if alert_filter == nil_alert {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(alert_filter, "alert")
					if err != nil {
						log.Fatal("alert get call failure ", err)
					} else {
						if alert_filter.AlertID == 0 {
							var data []circonusapi.Alert
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Alert
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}

		case "annotation":
			var annotation_filter circonusapi.AnnotationFilter
			err := json.Unmarshal([]byte(lookup), &annotation_filter)
			// nil_annotation for verifying unmarshalling
			nil_annotation := circonusapi.AnnotationFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if annotation_filter == nil_annotation {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(annotation_filter, "annotation")
					if err != nil {
						log.Fatal("annotation get call failure ", err)
					} else {
						if annotation_filter.AnnotationID == 0 {
							var data []circonusapi.Annotation
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Annotation
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}

		case "broker":
			var broker_filter circonusapi.BrokerFilter
			err := json.Unmarshal([]byte(lookup), &broker_filter)
			// nil_broker for verifying unmarshalling
			nil_broker := circonusapi.BrokerFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if broker_filter == nil_broker {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(broker_filter, "broker")
					if err != nil {
						log.Fatal("broker get call failure ", err)
					} else {
						if broker_filter.BrokerID == 0 {
							var data []circonusapi.Broker
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Broker
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}

		case "check":
			var check_filter circonusapi.CheckFilter
			err := json.Unmarshal([]byte(lookup), &check_filter)
			// nil_check for verifying unmarshalling
			nil_check := circonusapi.CheckFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if check_filter == nil_check {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(check_filter, "check")
					if err != nil {
						log.Fatal("check get call failure ", err)
					} else {
						if check_filter.CheckID == 0 {
							var data []circonusapi.Check
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Check
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}

		case "check_bundle":
			var checkbundle_filter circonusapi.CheckBundleFilter
			err := json.Unmarshal([]byte(lookup), &checkbundle_filter)
			// nil_checkbundle for verifying unmarshalling
			nil_checkbundle := circonusapi.CheckBundleFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if checkbundle_filter == nil_checkbundle {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(checkbundle_filter, "check_bundle")
					if err != nil {
						log.Fatal("check_bundle get call failure ", err)
					} else {
						if checkbundle_filter.CheckBundleID == 0 {
							var data []circonusapi.CheckBundle
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.CheckBundle
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}

		case "check_move":
			var checkmove_filter circonusapi.CheckMoveFilter
			err := json.Unmarshal([]byte(lookup), &checkmove_filter)
			// nil_checkmove for verifying unmarshalling
			nil_checkmove := circonusapi.CheckMoveFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if checkmove_filter == nil_checkmove {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(checkmove_filter, "check_move")
					if err != nil {
						log.Fatal("check_move get call failure ", err)
					} else {
						if checkmove_filter.CheckMoveID == 0 {
							var data []circonusapi.CheckMove
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.CheckMove
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}

		case "data":
			var data_filter circonusapi.DataFilter
			err := json.Unmarshal([]byte(lookup), &data_filter)
			// nil_data for verifying unmarshalling
			nil_data := circonusapi.DataFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if data_filter == nil_data {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetData(data_filter)
					if err != nil {
						log.Fatal("data get call failure ", err)
					} else {
						rendered_data, _ := json.MarshalIndent(result, "", "\t")

						fmt.Println(string(rendered_data))
					}
				}
			}
		case "dashboard":
			var dashboard_filter circonusapi.DashboardFilter
			err := json.Unmarshal([]byte(lookup), &dashboard_filter)
			// nil_dashboard for verifying unmarshalling
			nil_dashboard := circonusapi.DashboardFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if dashboard_filter == nil_dashboard {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(dashboard_filter, "dashboard")
					if err != nil {
						log.Fatal("dashboard get call failure ", err)
					} else {
						if dashboard_filter.DashboardId == 0 {
							var data []circonusapi.Dashboard
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Dashboard
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "graph":
			var graph_filter circonusapi.GraphFilter
			err := json.Unmarshal([]byte(lookup), &graph_filter)
			// nil_graph for verifying unmarshalling
			nil_graph := circonusapi.GraphFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if graph_filter == nil_graph {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(graph_filter, "graph")
					if err != nil {
						log.Fatal("data get call failure ", err)
					} else {
						if graph_filter.GraphID == "" {
							var data []circonusapi.Graph
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Graph
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "maintenance":
			var maintenance_filter circonusapi.MaintenanceFilter
			err := json.Unmarshal([]byte(lookup), &maintenance_filter)
			// nil_maintenance for verifying unmarshalling
			nil_maintenance := circonusapi.MaintenanceFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if maintenance_filter == nil_maintenance {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(maintenance_filter, "maintenance")
					if err != nil {
						log.Fatal("maintenance get call failure ", err)
					} else {
						if maintenance_filter.MaintenanceID == 0 {
							var data []circonusapi.Maintenance
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Maintenance
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "metric_cluster":
			var metriccluster_filter circonusapi.MetricClusterFilter
			err := json.Unmarshal([]byte(lookup), &metriccluster_filter)
			// nil_metriccluster for verifying unmarshalling
			nil_metriccluster := circonusapi.MetricClusterFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if metriccluster_filter == nil_metriccluster {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(metriccluster_filter, "metric_cluster")
					if err != nil {
						log.Fatal("metric_cluster get call failure ", err)
					} else {
						if metriccluster_filter.MetricClusterID == 0 {
							var data []circonusapi.MetricCluster
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.MetricCluster
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "rule_set":
			var ruleset_filter circonusapi.RuleSetFilter
			err := json.Unmarshal([]byte(lookup), &ruleset_filter)
			// nil_rulesetfilter for verifying unmarshalling
			nil_rulesetfilter := circonusapi.RuleSetFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if ruleset_filter == nil_rulesetfilter {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(ruleset_filter, "rule_set")
					if err != nil {
						log.Fatal("rule_set get call failure ", err)
					} else {
						if ruleset_filter.Name == "" {
							var data []circonusapi.RuleSet
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.RuleSet
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "rule_set_group":
			var rulesetgroup_filter circonusapi.RuleSetGroupFilter
			err := json.Unmarshal([]byte(lookup), &rulesetgroup_filter)
			// nil_rulesetgroupfilter for verifying unmarshalling
			nil_rulesetgroupfilter := circonusapi.RuleSetGroupFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if rulesetgroup_filter == nil_rulesetgroupfilter {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(rulesetgroup_filter, "rule_set_group")
					if err != nil {
						log.Fatal("rule_set_group get call failure ", err)
					} else {
						if rulesetgroup_filter.RuleSetGroupId == 0 {
							var data []circonusapi.RuleSetGroup
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.RuleSetGroup
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "tag":
			var tag_filter circonusapi.TagFilter
			err := json.Unmarshal([]byte(lookup), &tag_filter)
			// nil_tag for verifying unmarshalling
			nil_tag := circonusapi.TagFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if tag_filter == nil_tag {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(tag_filter, "tag")
					if err != nil {
						log.Fatal("tag get call failure ", err)
					} else {
						if tag_filter.TagId == "" {
							var data []circonusapi.Tag
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Tag
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "template":
			var template_filter circonusapi.TemplateFilter
			err := json.Unmarshal([]byte(lookup), &template_filter)
			// nil_templatefilter for verifying unmarshalling
			nil_templatefilter := circonusapi.TemplateFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if template_filter == nil_templatefilter {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(template_filter, "template")
					if err != nil {
						log.Fatal("rule_set_group get call failure ", err)
					} else {
						if template_filter.TemplateId == 0 {
							var data []circonusapi.Template
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Template
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "user":
			var user_filter circonusapi.UserFilter
			err := json.Unmarshal([]byte(lookup), &user_filter)
			// nil_user for verifying unmarshalling
			nil_user := circonusapi.UserFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if user_filter == nil_user {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(user_filter, "user")
					if err != nil {
						log.Fatal("user get call failure ", err)
					} else {
						if user_filter.UserID == 0 {
							var data []circonusapi.User
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.User
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		case "worksheet":
			var worksheet_filter circonusapi.WorksheetFilter
			err := json.Unmarshal([]byte(lookup), &worksheet_filter)
			// nil_worksheetfilter for verifying unmarshalling
			nil_worksheetfilter := circonusapi.WorksheetFilter{}

			if err != nil {
				log.Fatal(err)

			} else {
				if worksheet_filter == nil_worksheetfilter {
					log.Fatal("The Value of -where string or -file content "+
						"flag is not properly formatted : \n", lookup, "\n")
				} else {
					result, err := circonusapi.GetCns(worksheet_filter, "worksheet")
					if err != nil {
						log.Fatal("worksheet get call failure ", err)
					} else {
						if worksheet_filter.WorksheetId == "" {
							var data []circonusapi.Worksheet
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						} else {
							var data circonusapi.Worksheet
							json.Unmarshal(result, &data)
							rendered_data, _ := json.MarshalIndent(data, "", "\t")
							fmt.Println(string(rendered_data))
						}
					}
				}
			}
		default:
			log.Println("unrecognized value for -object flag : ", *object)
		}

	case "list":
		switch *object {
		case "account":
			result, err := circonusapi.GetCns(circonusapi.AccountFilter{}, "account")
			if err != nil {
				log.Fatal("account list call failure ", err)
			} else {
				var data []circonusapi.Account
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		/*
		 */
		case "alert":
			result, err := circonusapi.GetCns(circonusapi.AlertFilter{}, "alert")
			if err != nil {
				log.Fatal("alert list call failure ", err)
			} else {
				var data []circonusapi.Alert
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "annotation":
			result, err := circonusapi.GetCns(circonusapi.AnnotationFilter{}, "annotation")
			if err != nil {
				log.Fatal("annotation list call failure ", err)
			} else {
				var data []circonusapi.Annotation
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "broker":
			result, err := circonusapi.GetCns(circonusapi.BrokerFilter{}, "broker")
			if err != nil {
				log.Fatal("broker list call failure ", err)
			} else {
				var data []circonusapi.Broker
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "check":
			result, err := circonusapi.GetCns(circonusapi.CheckFilter{}, "check")
			if err != nil {
				log.Fatal("check list call failure ", err)
			} else {
				var data []circonusapi.Check
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "check_bundle":
			result, err := circonusapi.GetCns(circonusapi.CheckBundleFilter{}, "check_bundle")
			if err != nil {
				log.Fatal("check_bundle list call failure ", err)
			} else {
				var data []circonusapi.CheckBundle
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "check_move":
			result, err := circonusapi.GetCns(circonusapi.CheckMoveFilter{}, "check_move")
			if err != nil {
				log.Fatal("check_move list call failure ", err)
			} else {
				var data []circonusapi.CheckMove
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "data":
			result, err := circonusapi.GetData(circonusapi.DataFilter{})
			if err != nil {
				log.Fatal("data list call failure ", err)
			} else {
				rendered_data, _ := json.MarshalIndent(result, "", "\t")

				fmt.Println(string(rendered_data))
			}
		case "dashboard":
			result, err := circonusapi.GetCns(circonusapi.DashboardFilter{}, "dashboard")
			if err != nil {
				log.Fatal("dashboard get call failure ", err)
			} else {
				//fmt.Println(string(result))
				var data []circonusapi.Dashboard
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "graph":
			result, err := circonusapi.GetCns(circonusapi.GraphFilter{}, "graph")
			if err != nil {
				log.Fatal("graph get call failure ", err)
			} else {
				var data []circonusapi.Graph
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "maintenance":
			result, err := circonusapi.GetCns(circonusapi.MaintenanceFilter{}, "maintenance")
			if err != nil {
				log.Fatal("maintenance list call failure ", err)
			} else {
				var data []circonusapi.Maintenance
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "metric_cluster":
			result, err := circonusapi.GetCns(circonusapi.MetricClusterFilter{}, "metric_cluster")
			//fmt.Println(result)
			if err != nil {
				log.Fatal("metric_cluster list call failure ", err)
			} else {
				var data []circonusapi.MetricCluster
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "rule_set":
			result, err := circonusapi.GetCns(circonusapi.RuleSetFilter{}, "rule_set")
			if err != nil {
				log.Fatal("rule_set list call failure ", err)
			} else {
				var data []circonusapi.RuleSet
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "rule_set_group":
			result, err := circonusapi.GetCns(circonusapi.RuleSetGroupFilter{}, "rule_set_group")
			if err != nil {
				log.Fatal("rule_set_group list call failure ", err)
			} else {
				var data []circonusapi.RuleSetGroup
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "tag":
			result, err := circonusapi.GetCns(circonusapi.TagFilter{}, "tag")
			//fmt.Println(result)
			if err != nil {
				log.Fatal("metriccluster list call failure ", err)
			} else {
				var data []circonusapi.Tag
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "template":
			result, err := circonusapi.GetCns(circonusapi.TemplateFilter{}, "template")
			//fmt.Println(result)
			if err != nil {
				log.Fatal("template list call failure ", err)
			} else {
				var data []circonusapi.Template
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "user":
			result, err := circonusapi.GetCns(circonusapi.UserFilter{}, "user")
			if err != nil {
				log.Fatal("user list call failure ", err)
			} else {
				var data []circonusapi.User
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		case "worksheet":
			result, err := circonusapi.GetCns(circonusapi.WorksheetFilter{}, "worksheet")
			if err != nil {
				log.Fatal("worksheet list call failure ", err)
			} else {
				var data []circonusapi.Worksheet
				json.Unmarshal(result, &data)
				rendered_data, _ := json.MarshalIndent(data, "", "\t")
				fmt.Println(string(rendered_data))
			}
		default:
			log.Println("unrecognized value for -object flag : ", *object)
		}

	case "create":
		switch *object {
		case "account":
			log.Fatal("Currently accounts can only be opened and closed via the GUI. It is" +
				" not foreseen that there will be a need to do this on an automated basis.")
		case "alert":
			log.Fatal("Alerts are Readonly, cannot create new alert through the " +
				"alert API endpoint")
		case "annotation":
			result, err := circonusapi.CreateCns([]byte(lookup), "annotation")
			if err != nil {
				log.Fatal("annotation create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_annotation circonusapi.Annotation
					json.Unmarshal([]byte(rendered_data.String()), &new_annotation)
					if new_annotation.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("Annotation creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}

		case "broker":
			log.Fatal("You can't modify (or for that matter create or delete brokers) with " +
				"the API. Please use the web interface for this functionality ")
		case "check":
			log.Fatal("Checks are read only, to make a change to one you must change to all the " +
				"checks in a bundle, so any modifications should be done to the check_bundle endpoint")
		case "check_bundle":
			result, err := circonusapi.CreateCns([]byte(lookup), "check_bundle")
			if err != nil {
				log.Fatal("check_bundle create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_checkbundle circonusapi.CheckBundle
					json.Unmarshal([]byte(rendered_data.String()), &new_checkbundle)
					if new_checkbundle.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("check_bundle creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}

		case "check_move":
			result, err := circonusapi.CreateCns([]byte(lookup), "check_move")
			if err != nil {
				log.Fatal("check_move create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_checkmove circonusapi.CheckMove
					json.Unmarshal([]byte(rendered_data.String()), &new_checkmove)
					if new_checkmove.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("check_move creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "data":
			log.Fatal("Data points are Readonly, cannot create new data point through" +
				" the Data API endpoint")
		case "dashboard":
			result, err := circonusapi.CreateCns([]byte(lookup), "dashboard")
			if err != nil {
				log.Fatal("Dashboard create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_dashboard circonusapi.Dashboard
					json.Unmarshal([]byte(rendered_data.String()), &new_dashboard)
					if new_dashboard.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("Dashboard creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "graph":
			result, err := circonusapi.CreateCns([]byte(lookup), "graph")
			if err != nil {
				log.Fatal("Graph create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_graph circonusapi.Graph
					json.Unmarshal([]byte(rendered_data.String()), &new_graph)
					if new_graph.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("Graph creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "maintenance":
			result, err := circonusapi.CreateCns([]byte(lookup), "maintenance")
			if err != nil {
				log.Fatal("Maintenance create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_maintenance circonusapi.Maintenance
					json.Unmarshal([]byte(rendered_data.String()), &new_maintenance)
					if new_maintenance.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("Maintenance creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "metric_cluster":
			result, err := circonusapi.CreateCns([]byte(lookup), "metric_cluster")
			if err != nil {
				log.Fatal("metric_cluster create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_metriccluster circonusapi.MetricCluster
					json.Unmarshal([]byte(rendered_data.String()), &new_metriccluster)
					if new_metriccluster.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("metric_cluster creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "rule_set":
			result, err := circonusapi.CreateCns([]byte(lookup), "rule_set")
			if err != nil {
				log.Fatal("rule_set create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_rule_set circonusapi.RuleSet
					json.Unmarshal([]byte(rendered_data.String()), &new_rule_set)
					if new_rule_set.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("rule_set creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "rule_set_group":
			result, err := circonusapi.CreateCns([]byte(lookup), "rule_set_group")
			if err != nil {
				log.Fatal("rule_set_group create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_rule_set_group circonusapi.RuleSetGroup
					json.Unmarshal([]byte(rendered_data.String()), &new_rule_set_group)
					if new_rule_set_group.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("rule_set_group creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "tag":
			log.Fatal("Tags are Readonly, Tags can only be listed throug the Tag " +
				"API endpoint")
		case "template":
			result, err := circonusapi.CreateCns([]byte(lookup), "template")
			if err != nil {
				log.Fatal("template create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_template circonusapi.Template
					json.Unmarshal([]byte(rendered_data.String()), &new_template)
					if new_template.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("template creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		case "user":
			log.Fatal("Currently users can only be created via people signing up " +
				"through the website.")
		case "worksheet":
			result, err := circonusapi.CreateCns([]byte(lookup), "worksheet")
			if err != nil {
				log.Fatal("worksheet create call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					var new_worksheet circonusapi.Worksheet
					json.Unmarshal([]byte(rendered_data.String()), &new_worksheet)
					if new_worksheet.Cid != "" {
						fmt.Println(rendered_data.String())
					} else {
						fmt.Println("worksheet creation failed")
						log.Fatal(rendered_data.String())
					}
				}
			}
		}
	case "update":
		switch *object {
		case "account":
			result, err := circonusapi.UpdateCns(*oid, "account", []byte(lookup))
			if err != nil {
				log.Fatal("Account update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}

		case "alert":
			log.Fatal("Alerts are Readonly, cannot update alerts through the " +
				"alert API endpoint")
		case "annotation":
			result, err := circonusapi.UpdateCns(*oid, "annotation", []byte(lookup))
			if err != nil {
				log.Fatal("Annotation update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "broker":
			log.Fatal("You can't modify (or for that matter create or delete brokers) " +
				"with the API. Please use the web interface for this functionality ")

		case "check":
			log.Fatal("Checks are Readonly, cannot update Checks through the " +
				"Check API endpoint")
		case "check_bundle":
			result, err := circonusapi.UpdateCns(*oid, "check_bundle", []byte(lookup))
			if err != nil {
				log.Fatal("check_bundle update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "check_move":
			result, err := circonusapi.UpdateCns(*oid, "check_move", []byte(lookup))
			if err != nil {
				log.Fatal("check_move update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "data":
			log.Fatal("Data points are Readonly, cannot update data point through" +
				" the Data API endpoint")
		case "dashboard":
			result, err := circonusapi.UpdateCns(*oid, "dashboard", []byte(lookup))
			if err != nil {
				log.Fatal("dashboard update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "graph":
			result, err := circonusapi.UpdateCns(*oid, "graph", []byte(lookup))
			if err != nil {
				log.Fatal("graph update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "maintenance":
			result, err := circonusapi.UpdateCns(*oid, "maintenance", []byte(lookup))
			if err != nil {
				log.Fatal("check_move update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "metric_cluster":
			result, err := circonusapi.UpdateCns(*oid, "metric_cluster", []byte(lookup))
			if err != nil {
				log.Fatal("metric_cluster update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "rule_set":
			result, err := circonusapi.UpdateCns(*oid, "rule_set", []byte(lookup))
			if err != nil {
				log.Fatal("rule_set update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "rule_set_group":
			result, err := circonusapi.UpdateCns(*oid, "rule_set_group", []byte(lookup))
			if err != nil {
				log.Fatal("rule_set_group update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "tag":
			log.Fatal("Tags are Readonly, Tags can only be listed throug the Tag " +
				"API endpoint")
		case "template":
			result, err := circonusapi.UpdateTemplate(*oid, "template", *host_removal_action, *bundle_removal_action, []byte(lookup))
			if err != nil {
				log.Fatal("template update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "user":
			result, err := circonusapi.UpdateCns(*oid, "user", []byte(lookup))
			if err != nil {
				log.Fatal("user update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		case "worksheet":
			result, err := circonusapi.UpdateCns(*oid, "worksheet", []byte(lookup))
			if err != nil {
				log.Fatal("worksheet update call error ", err)
			} else {
				var rendered_data bytes.Buffer
				err := json.Indent(&rendered_data, result, "", "\t")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(rendered_data.String())
				}
			}
		}
	case "delete":
		switch *object {
		case "account":
			log.Fatal("Currently accounts can only be opened and closed via the GUI. It is not foreseen that there will be a need to do this on an automated basis.")
		case "alert":
			log.Fatal("Alerts are Readonly, cannot delete alert through the " +
				"alert API endpoint")
		case "annotation":
			result, err := circonusapi.DeleteCns(*oid, "annotation")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("annotation delete : ", *oid, string(result))
			}
		case "broker":
			log.Fatal("You can't modify (or for that matter create or delete brokers) with the API. Please use the web interface for this functionality ")
		case "check":
			log.Fatal("Checks are Readonly, cannot delete a Check through the " +
				"Check API endpoint")
		case "check_bundle":
			result, err := circonusapi.DeleteCns(*oid, "check_bundle")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("check_bundle delete : ", *oid, string(result))
			}
			log.Println("You cannot completely remove check bundles from the system, this sets the check_bundle status to disabled and hides the check_bundle from API listing and UI")
		case "check_move":
			result, err := circonusapi.DeleteCns(*oid, "check_move")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("check_move delete : ", *oid, string(result))
			}
		case "data":
			log.Fatal("Data points are Readonly, cannot delete a data point through" +
				" the Data API endpoint")
		case "dashboard":
			result, err := circonusapi.DeleteCns(*oid, "dashboard")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("dashboard delete : ", *oid, string(result))
			}
		case "graph":
			result, err := circonusapi.DeleteCns(*oid, "graph")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("graph delete : ", *oid, string(result))
			}

		case "maintenance":
			result, err := circonusapi.DeleteCns(*oid, "maintenance")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("maintenance delete : ", *oid, string(result))
			}
		case "metric_cluster":
			result, err := circonusapi.DeleteCns(*oid, "metric_cluster")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("metric_cluster delete : ", *oid, string(result))
			}
		case "rule_set":
			result, err := circonusapi.DeleteCns(*oid, "rule_set")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("rule_set delete : ", *oid, string(result))
			}
		case "rule_set_group":
			result, err := circonusapi.DeleteCns(*oid, "rule_set_group")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("rule_set_group delete : ", *oid, string(result))
			}
		case "tag":
			log.Fatal("Tags are Readonly, Tags can only be listed throug the Tag " +
				"API endpoint")
		case "template":
			result, err := circonusapi.DeleteCns(*oid, "template")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("template delete : ", *oid, string(result))
			}
		case "user":
			log.Println("Users cannot be deleted directly, though they can be removed from an account with the API by modifying the roles field of the account with the Circonus account API")
		case "worksheet":
			result, err := circonusapi.DeleteCns(*oid, "worksheet")
			if err != nil {
				log.Fatal("Delete call errored : \n", err)
			}
			if result != nil {
				log.Println("worksheet delete : ", *oid, string(result))
			}
		}
	}

}
