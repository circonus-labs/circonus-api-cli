# Circonus CLI (circli) 
Go implementation of CLI wrapper for Circonus API (based on the documentation at https://login.circonus.com/resources/api)

## Building CLI Binary
- Update Circonus API endpoint settings(Token,CirconusURL and AppName) in **circonusapi/endpoint.go** 
```go
package circonusapi

const (
	Token       = "ffffffff-ffff-ffff-ffff-ffffffffffff"
	CirconusURL = "https://api.circonus.com/v2/"
	Accept      = "application/json"
	ContentType = "application/json"
	AppName     = "xyz"
)
```
- Compile **calls/circli.go**
```
$ go build circli.go 
```
## Usage :

```
Usage of circli:
  -bundle_removal_action string
        Bundle removal action used with template update actions: possible values are "unbind", "deactivate", "remove"
        unbind :        Marks the check(s) as unbound from the template, you can then modify them as if they were a "normal" check
        deactivate :    Sets the check(s) as inactive, they will still show up on the interface and you can view historic data for them 
        remove :        Deletes the check(s) from the system, they will no longer show in the UI and historic data will be gone 

  -call string
        Circonus Call Type. Example : get, list, create, update, delete

  -file string
        JSON formatted data input file

  -host_removal_action string
        Host removal action used with template update actions: possible values are "unbind", "deactivate", "remove"
        unbind :        Marks the check(s) as unbound from the template, you can then modify them as if they were a "normal" check
        deactivate :    Sets the check(s) as inactive, they will still show up on the interface and you can view historic data for them 
        remove :        Deletes the check(s) from the system, they will no longer show in the UI and historic data will be gone 

  -object string
        Circonus Object Type. Possible object types : 
        account :       basic contact details associated with the account and description 
        alert :         Representation of an Alert that occurred (Readonly) 
        annotation :    Mark important events used for correlation 
        broker :        Remote software agent that collects the data from monitored hosts
        caql :          Provides a way to extact data from Circonus using a CAQL query
        check :         Individual elements of a check bundle (Readonly)
        check_bundle :  Collection of checks that have the same configuration and target, but collected from different brokers
        check_bundle_metrics :  Provides a way to add or remove individual metrics under a Check Bundle rather than replacing the entire Check Bundle object
        check_move :    Request that a Check be moved from one Broker to another
        contact_group : Provides means of being notified about alerts. Each contact_group can have one to many users and means of contact
        data :          Readonly endpoint to pull the values of a single data point for a given time range
        dashboard :     Provides access for creating, reading, updating and deleting Dashboards.
        graph :         Allows mass creation, editing and removal of graphs
        maintenance :   Schedule a maintenance window for your account, check bundle, rule set or host
        metric :        Provides API access to individual Metrics. Generally readonly but units and tags fields can be updated
        metric_cluster :        A metric cluster is a cluster of metrics defined by a set of queries
        rule_set :      define a collection of rules to apply to a given metric
        rule_set_group :        Allows togroup together rule sets and trigger alerts based on combinations of those rule sets faulting
        tag :           List all tags in use in your account (don't have any fields, Readonly)
        template :      A means to setup a mass number of checks quickly through both the API and UI
        user :          Get basic information about the users on your account or a single user
        worksheet :     Collection of graphs and allow quick correlation across them

  -oid string
        Circonus object ID. Value of _cid in the API object without the "/<object_type>/" prefix. This flag is used with "update" and "delete" calls

  -where string
        JSON string used for querying where clause
```

## Circonus CLI (circli) user manual
#####[Get](https://github.com/misale/circonus-api-go#get-calls) 
#####[List](https://github.com/misale/circonus-api-go#list-calls)  
#####[Create](https://github.com/misale/circonus-api-go#create-calls)
#####[Update](https://github.com/misale/circonus-api-go#update-calls)
#####[Delete](https://github.com/misale/circonus-api-go#delete-calls) 
#

### Get Calls
Get calls are made to filter/search Circonus API objects using a json formatted where string provided by -where flag or as a file with -file flag. The possible where keys are listed under "Query format" for each type of API object, the 
values of the where keys can be used in combination for filtering/searching.
#### CLI Syntax
```
circli -object <circonus_object_name> -call get [ -where <json_format_where_string> | -file <json_file> ]
```
#
**account** : basic contact details associated with the account and description
 - **Query format** 
```json
 {"account_id":"<account_id_number>","name":"<account_name>"}
```
  - **Query field (key) description**
    - **account_id** : the integer part of the value of _cid in the account API object 
    - **name** : string value of name in account API object
 - **Example**
```json
        $ circli -object account -call get -where '{"account_id":2317}'
        {
                "_cid": "/account/2317",
                "_contact_groups": [
                        "/contact_group/2365",
                        "/contact_group/2321",
                        "/contact_group/2324"
                ],
                "_owner": "/user/5474",
                "_ui_base_url": "https://xyz.circonus.com/",
                "_usage": [
                        {
                                "_limit": 10000,
                                "_type": "Metrics",
                                "_used": 7426
                        }
                ],
                "address1": "",
                "address2": "",
                "city": "",
                "country_code": "",
                "description": "XYZ",
                "invites": [
                        {
                                "email": "jane_doe@xyz.com",
                                "role": "Read Only"
                        }
                ],
                "name": "xyz",
                "state_prov": "",
                "timezone": "US/Eastern",
                "users": [
                        {
                                "role": "Admin",
                                "user": "/user/3142"
                        },
                        {
                                "role": "Normal",
                                "user": "/user/5856"
                        }
                ]
        }
        $ 
```
#
 **alert** : Representation of an Alert that occurred (Readonly) 
 - **Query format**
```json
    {"alert_id":"<alert_id>" ,"check":"<check_id>","metric_name":"<metric_name>", "severity" : "<alert_severity_num>" , "occurred_on_lt":"<epoch_time_upper_bound>","occurred_on_gt":"<epoch_time_lower_bound>","tags_has":"<tag>" } 
```
 - **Query field (key) description**
    - **alert_id** : the integer part of the value of _cid in the alert API object 
    - **severity** : the value of _severity in the alert API object
    - **check** : integer part of check API object _cid (for filtering alerts by check ID)
    - **metric_name** : metric name string for filtering alerts associated with a specific metric ("name"+"_"+"check_id")
    - **occurred_on_lt** : upper bound epoch timestamp for _occurred_on (lt = less than)
    - **occurred_on_le** : upper bound epoch timestamp for _occurred_on (le = less than or equal to)
    - **occurred_on_gt** : lower bound epoch timestamp for _occurred_on (gt = greater than)
    - **occurred_on_ge** : lower bound epoch timestamp for _occurred_on (ge = greater than or equal to)
    - **tags_has** : tag string
 - **Example**
```json
 $ circli -object alert -call get -where '{"occurred_on_gt":1460000000,"occurred_on_lt":1470048105}'
 [
         {
                 "_acknowledgement": "",
                 "_alert_url": "https://xyz.circonus.com/fault-detection?alert_id=32855508",
                 "_broker": "/broker/1",
                 "_check": "/check/165697",
                 "_check_name": "1.1.1.1 http",
                 "_cid": "/alert/32855508",
                 "_cleared_on": 0,
                 "_cleared_value": 0,
                 "_maintenance": [],
                 "_metric_link": "",
                 "_metric_name": "tt_connect",
                 "_metric_notes": "",
                 "_occurred_on": 0,
                 "_rule_set": "/rule_set/165697_tt_connect",
                 "_severity": 1,
                 "_tags": [
                         "test"
                 ],
                 "_value": 0
         }
 ]
 $ 
```
#
**annotation** : Mark important events used for correlation
 - **Query format**
```json
  {"annotation_id":"<annotation_id>","category":"<category_string>","start":"<lower_bound_epoch_timestamp>","stop":"<upper_bound_epoch_timestamp>"}
```
  - **Query field (key) description**
    - **annotation_id** : The integer part of the _cid value of annotation API object
    - **category** : The category string of annotation API object
    - **start_gt** : Lower bound (gt = greater than) for epoch start timestamp of annotation API object 
    - **start_ge** : Lower bound (ge = greater than or equal to) for epoch start timestamp of annotation API object 
    - **start_lt** : Upper bound (lt = less than) for epoch start timestamp of annotation API object 
    - **start_le** : Upper bound (le = less than or equal to) for epoch start timestamp of annotation API object 
    - **stop_lt** : Upper bound (lt = less than) for epoch stop timestamp of annotation API object
    - **stop_le** : Upper bound (lt = less than or equal to) for epoch stop timestamp of annotation API object
    - **stop_gt** : Lower bound (gt = greater than) for epoch stop timestamp of annotation API object
    - **stop_ge** : Lower bound (ge = greater than or equal to) for epoch stop timestamp of annotation API object
 - **Example**
```json
  $ circli -object annotation -call get -where '{"start_gt":1467000000,"stop_lt":1470048105}'
  [
          {
                  "_cid": "/annotation/124216",
                  "_created": 1.467047764e+09,
                  "_last_modified": 1.467047764e+09,
                  "_last_modified_by": "/user/5483",
                  "category": "network maintenance",
                  "description": "Testing Annotationthrough API",
                  "start": 1.467047744e+09,
                  "stop": 1.467048744e+09,
                  "title": "Test Annotation"
          },
          {
                  "_cid": "/annotation/124964",
                  "_created": 1.467127597e+09,
                  "_last_modified": 1.467128656e+09,
                  "_last_modified_by": "/user/5483",
                  "category": "network maintenance",
                  "description": "Testing Annotationthrough API",
                  "start": 1.467047744e+09,
                  "stop": 1.468827165e+09,
                  "title": "Test Annotation updated 3"
          }
  ]
  $ 
```
#
**broker** : Remote software agent that collects the data from monitored hosts
 - **Query format**
```json
  {"broker_id":"<broker_id>","name":"<broker_name>","type":"<broker_type>"}
```
  - **Query field (key) description**
    - **broker_id** : The integer part of the _cid value in broker API object
    - **name** : The string value of _name in broker API object
    - **type** : The string value of _type in broker API object
 - **Example**
```json
  $ circli -object broker -call get -where '{"type":"enterprise"}'
   [
           {
                   "_cid": "/broker/1289",
                   "_details": [
                           {
                                   "cn": "a9999-x999999.noit.circonus.net",
                                   "external_host": null,
                                   "external_port": 43191,
                                   "ipaddress": "1.1.1.1",
                                   "minimum_version_required": 1.461358415e+09,
                                   "modules": [
                                           "cim",
                                           "circonuswindowsagent",
                                           "cloudwatch",
                                           "collectd",
                                           "dcm",
                                           "dhcp",
                                           "dns",
                                           "ec_console",
                                           "elasticsearch",
                                           "external",
                                           "ganglia",
                                           "googleanalytics:m1",
                                           "googleanalytics:m2",
                                           "googleanalytics:m3",
                                           "googleanalytics:m4",
                                           "googleanalytics:m5",
                                           "googleanalytics:m6",
                                           "googleanalytics:m7",
                                           "haproxy",
                                           "http",
                                           "httptrap",
                                           "imap",
                                           "jmx",
                                           "json",
                                           "keynote",
                                           "keynote_pulse",
                                           "ldap",
                                           "memcached",
                                           "munin",
                                           "mysql",
                                           "newrelic_rpm",
                                           "nginx",
                                           "nrpe",
                                           "ntp",
                                           "oracle",
                                           "ping_icmp",
                                           "pop3",
                                           "postgres",
                                           "redis",
                                           "resmon",
                                           "selfcheck",
                                           "smtp",
                                           "snmp",
                                           "sparkpost",
                                           "sqlserver",
                                           "ssh2",
                                           "statsd",
                                           "tcp",
                                           "varnish"
                                   ],
                                   "port": null,
                                   "skew": "-0.2035",
                                   "status": "active",
                                   "version": 1.468274736e+09
                           }
                   ],
                   "_latitude": "39.9607",
                   "_longitude": "-75.6055",
                   "_name": "a9999-x999999",
                   "_tags": [],
                   "_type": "enterprise"
           }
   ]
   $ 
```
#
**caql** : The CAQL API endpoint provides a way to extact data from Circonus using a CAQL query
 - **Query format**
```json
 {"query":"<caql_query>","start":"<start_epoch>","end":"<end_epoch>","period":"<period>"}
```
 - **Query field (key) description**
    - **query** : CAQL query string
    - **start** : lower bound epoch time (integer) for CAQL query
    - **end** : upper bound epoch time (integer) for CAQL query
    - **period** : time period in seconds for data extraction
 - **Example**
```
$ circli -call get -object caql -where '{"query":"metric:average(\"05ef308d-e19c-6515-9e43-c4fffff51823\",\"sessions\")","start":1478558000,"end":1478559000,"period":60}'
{"_data":[[1478557980,[1]],[1478558040,[4]],[1478558100,[2]],[1478558160,[3]],[1478558220,[6]],[1478558280,[2]],[1478558340,[5]],[1478558400,[2]],[1478558460,[2]],[1478558520,[3]],[1478558580,[6]],[1478558640,[4]],[1478558700,[4]],[1478558760,
[6]],[1478558820,[2]],[1478558880,[10]],[1478558940,[3]]],"_start":1478558000,"_end":1478559000,"_query":"metric:average(\"05ef308d-e19c-6515-9e43-c4fffff51823\",\"sessions\")","_period":60}
$ 
```
#
**check** : Individual elements of a check bundle (Readonly)
 - **Query format**
```json
  {"check_id":"<check_id>","check_bundle_id":"<check_bundle_id>","check_uuid":"<check_uuid>"}
```
 - **Query field (key) description**
    - **check_id** : The integer part of the value of _cid in a check API object
    - **check_bundle_id** : The integer part of the value of _check_bundle in a check API object
    - **check_uuid** : The string value of _check_uuid in a check API object
 - **Example**
```json
  $ circli -object check -call get -where '{"check_uuid":"ff8f62f8-ff87-4ffd-f27b-ff4ff34fff54"}'
  [
          {
                  "_active": true,
                  "_broker": "/broker/1289",
                  "_check_bundle": "/check_bundle/135179",
                  "_check_uuid": "ff8f62f8-ff87-4ffd-f27b-ff4ff34fff54",
                  "_cid": "/check/166356",
                  "_details": {
                          "submission_url": "https://1.1.1.1:43191/module/httptrap/ff8f62f8-ff87-4ffd-f27b-ff4ff34fff54/a469c069fa6830f2"
                  }
          }
  ]
  $ 
```
#
**check_bundle** : Collection of checks that have the same configuration and target, but collected from different brokers
 - **Query format**
```json
  {"bundle_id":"<check_bundle_id>","type":"<check_bundle_type>","target":"<target_host>","target_like":"<target_string>","display_name":"<display_name>","display_name_like":"<display_name_string>","tags_has":"<tag>","checks_has":"<check_id>",
  "brokers_has":"<broker_id>"} 
```
 - **Query field (key) description**
    - **check_bundle_id** : The integer part of _cid value in check_bundle API object (known check_bundle cid)
    - **type** : The string value of type in check_bundle API object (to list all check bundles that has a specific type field)
    - **target** : The string value of target in check_bundle API object (to list all check bundles that target a particular server)
    - **target_like** : String pattern of target to filter check_bundles that have target matching the pattern.
    - **display_name** : The string value of display_name in check_bundle API object (to find a check bundle with a particular name)
    - **display_name_like** : String pattern of display_name to filter check_bundles that have display_name matching the pattern.
    - **tags_has** : String element of the list value of tags in check_bundle API object (to list all check bundles with a particular tag)
    - **checks_has** : Integer part of element of the list value of checks in check_bundle API object (to find a check bundle that has a particular check in it)
    - **brokers_has** : Integer part of element of the list value of brokers in check_bundle API object (to list all check bundles using a particular broker)
 - **Example**
```json
  $ circli -object checkbundle -call get -where '{"target":"xyz-web-01.wash.dc.xyz.net"}'
  [
          {
                  "_check_uuids": [
                          "ff306363-9ff0-ff0f-91f2-f1052ff9f838"
                  ],
                  "_checks": [
                          "/check/167194"
                  ],
                  "_cid": "/check_bundle/135930",
                  "_created": 1.470319944e+09,
                  "_last_modified": 1.470319944e+09,
                  "_last_modified_by": "/user/5483",
                  "_reverse_connection_urls": [
                          "mtev_reverse://1.1.1.1:43191/check/ff306363-9ff0-ff0f-91f2-f1052ff9f838"
                  ],
                  "brokers": [
                          "/broker/1289"
                  ],
                  "config": {
                          "header_Host": "xyz-web-01.wash.dc.xyz.net",
                          "http_version": "1.1",
                          "method": "GET",
                          "payload": "",
                          "port": "8088",
                          "read_limit": "0",
                          "reverse:secret_key": "1929ffff-6870-f188-f5bf-f13f7f9193f1",
                          "url": "http://xyz-web-01.wash.dc.xyz.net/_metrics"
                  },
                  "display_name": "xyz-web-01.wash.dc.xyz.net json",
                  "metrics": [
                          {
                                  "name": "web`proxy.process.http.407_responses",
                                  "status": "active",
                                  "tags": [],
                                  "type": "numeric",
                                  "units": null
                          },
                          {
                                  "name": "web`proxy.process.cache.volume_1.lookup.active",
                                  "status": "active",
                                  "tags": [],
                                  "type": "numeric",
                                  "units": null
                          },
                          {
                                  "name": "web`proxy.process.cache.lookup.success",
                                  "status": "active",
                                  "tags": [],
                                  "type": "numeric",
                                  "units": null
                          },
                          {
                                  "name": "web`proxy.process.http.305_responses",
                                  "status": "active",
                                  "tags": [],
                                  "type": "numeric",
                                  "units": null
                          }
                  ],
                  "notes": "",
                  "period": 60,
                  "status": "active",
                  "tags": [],
                  "string": "",
                  "timeout": 10,
                  "type": "json"
          }
  ]
  $ 
```
#
**check_bundle_metrics** : Provides interface to update/list metrics under a check_bundle.
 - **Query format**
```json
  {"check_bundle_id":"<check_bundle_id>"} 
```
 - **Query field (key) description**
    - **check_bundle_id** : The integer part of _cid value in check_bundle API object (known check_bundle cid)
 - **Example**
```json
$ circli -call get -object check_bundle_metrics -where '{"check_bundle_id":158234}' 
{
        "_cid": "/check_bundle_metrics/158234",
        "metrics": [
                {
                        "name": "cpu`num",
                        "result": "success",
                        "status": "active",
                        "tags": [],
                        "type": "numeric",
                        "units": null
                },
                {
                        "name": "mem`pct",
                        "result": "success",
                        "status": "active",
                        "tags": [],
                        "type": "numeric",
                        "units": null
                },
                {
                        "name": "status",
                        "result": "success",
                        "status": "active",
                        "tags": [],
                        "type": "text",
                        "units": null
                }
        ]
}
$ 
```
#
**check_move** : Request that a Check be moved from one Broker to another
 - **Query format**
```json
  {"move_id":"<check_move_id>"}
```
 - **Query field (key) description**
    - **check_move_id** : The integer part of the _cid value in a check_move API object
 - **Example**
```json
 $ circli -object checkmove -call get -where '{"check_move_id":167194}'
 {
         "_broker": "/broker/1289",
         "_cid": "/check_move/167194",
         "_error": "",
         "_status": "Pending",
         "check_id": 167194,
         "new_broker": "/broker/1"
 }
 $ 
```
#
**contact_group** : Provides means of being notified about alerts. Each contact_group can have one to many users and means of contact
 - **Query format**
```json
  {"contact_group_id":"<contact_group_id>","name":"<contact_group name>","name_like":"<name string pattern>","tags_has":"<tag_string>"}
```
 - **Query field (key) description**
    - **contact_group_id** : The integer part of the _cid value in a contact_group API object
    - **name** : String value of name in a contact_group API object
    - **name_like** : String pattern of name to find all contact_groups with matching name
    - **tags_has** : Tag string to find contact_groups whose tags list includes this.
 - **Example**
```json
 $ circli -call get -object contact_group -where '{"contact_group_id":1234}'
 {
         "_cid": "/contact_group/1234",
         "_last_modified": 1.465402911e+09,
         "_last_modified_by": "/user/1007",
         "aggregation_window": 300,
         "alert_formats": {
                 "long_message": null,
                 "long_subject": null,
                 "long_summary": null,
                 "short_message": null,
                 "short_summary": null
         },
         "contacts": {
                 "external": [
                         {
                                 "contact_info": "slack://slack.com?token=xoxb-88888888888-FFFFFjqryFFFFFFFFFFFFF2v\u0026channel=D1F5FFFF4\u0026username=circonus_o_matic",
                                 "method": "slack"
                         }
                 ],
                 "users": []
         },
         "escalations": [
                 null,
                 null,
                 null,
                 null,
                 null
         ],
         "name": "testGroup",
         "reminders": [
                 0,
                 0,
                 0,
                 0,
                 300
         ],
         "tags": [
                 "service:test"
         ]
 }
 $ 
```
#
**dashboard** : Provides API access for creating, reading, updating and deleting Dashboards
  - **Query format**
```json
   {"dashboard_id":"<dashboard_id>","title":"<title>","title_like":"<title_like>"}
```
  - **Query field (key) description**
    - **dashboard_id** : Integer part of _cid value in dashboard API object
    - **title** : String value of title in dashboard API object
    - **title_like** : String pattern to search all dashboards that have titles matching the string pattern
  - **Example**
```json
  $ circli -object dashboard -call get -where '{"title_like":"Fancy*"}'
  [
          { 
                  "_created": 1472738070,
                  "options": {
                          "fullscreen_hide_title": false,
                          "text_size": 16,
                          "linkages": [],
                          "access_configs": [
                                  {
                                          "fullscreen_hide_title": false,
                                          "text_size": 16,
                                          "scale_text": true,
                                          "nickname": "nfd_test",
                                          "black_dash": false,
                                          "fullscreen": false,
                                          "shared_id": "5s2Cf7",
                                          "enabled": true
                                  }
                          ],
                          "scale_text": true,
                          "hide_grid": false
                  },
                  "_cid": "/dashboard/1749",
                  "shared": true,
                  "_active": false,
                  "title": "New Fancy Dashboard",
                  "_dashboard_uuid": "38f13373-675f-458f-9141-ffff498f6f13",
                  "account_default": false,
                  "grid_layout": {
                          "width": 8,
                          "height": 4
                  },
                  "_created_by": "/user/5352",
                  "_last_modified": 1474570596,
                  "widgets": [
                          {
                                  "active": true,
                                  "height": 1,
                                  "name": "Graph",
                                  "origin": "e0",
                                  "settings": {
                                          "_graph_title": "PRESENTATION DATA: xyz-web Server CPU Usage",
                                          "account_id": "2317",
                                          "date_window": "2w",
                                          "graph_id": "2ffff924-7978-f6f9-f9f9-f48151ff4ff8",
                                          "hide_xaxis": false,
                                          "hide_yaxis": false,
                                          "key_inline": false,
                                          "key_loc": "noop",
                                          "key_size": "1",
                                          "key_wrap": false,
                                          "label": "",
                                          "period": "2000",
                                          "realtime": false,
                                          "show_flags": false
                                  },
                                  "type": "graph",
                                  "widget_id": "w4",
                                  "width": 2
                          },
                          {
                                  "active": true,
                                  "height": 1,
                                  "name": "Gauge",
                                  "origin": "d0",
                                  "settings": {
                                          "_check_id": 163862,
                                          "account_id": "2317",
                                          "check_uuid": "f0f04f62-6800-f88f-ff25-f03ff8f15134",
                                          "disable_autoformat": false,
                                          "formula": "",
                                          "metric_display_name": "xyz-web-01.wash.dc.xyz.net cosi/system: cpu`wait_io (on 1.2.3.4, from a9999-x999999)",
                                          "metric_name": "cpu`wait_io",
                                          "period": 0,
                                          "range_high": 3,
                                          "range_low": 0,
                                          "thresholds": {
                                                  "colors": [
                                                          "#008000",
                                                          "#ffcc00",
                                                          "#ee0000"
                                                  ],
                                                  "flip": false,
                                                  "values": [
                                                          "75%",
                                                          "87.5%"
                                                  ]
                                          },
                                          "title": "CPU% IO Wait",
                                          "type": "dial",
                                          "value_type": "counter"
                                  },
                                  "type": "gauge",
                                  "widget_id": "w6",
                                  "width": 1
                          }
                  ]
          }
  ]
  $ 
```
#
**data** : Readonly endpoint to pull the values of a single data point for a given time range
 - **Query format**
```json
  {"check_id":"<check_id>","metric_name":"<metric_name>","period":"<period>","start":"<start_epoch>","stop":"<stop_epoch>","type":"<type_of_metric>"}
```
 - **Query field (key) description** 
    - **check_id** : The integer part of the value of _cid in a check API object
    - **metric_name** : String name of the metric that data extraction will be executed for
    - **period** : For numeric and histogram types only, and not text. The resolution of data you want to get back (in seconds), valid values are: 60, 300, 1800, 10800, and 86400. 
    - **start** : The start time, in epoch seconds, of the duration you wish to export data from
    - **stop** : The end time, in epoch seconds, of the duration you wish to export data from
    - **type** : The type of data you wish to extract. This must be text, numeric, or histogram
 - **Example**
```json
  $ circli -object data -call get -where '{"check_id":167720,"metric_name":"analytics`BACKEND`bout","period":300,"start":1471862000,"stop":1471863124,"type":"numeric"}'
  [
          {
                  "timestamp": 1.4718618e+09,
                  "value": 6.9220423e+07,
                  "count": 5,
                  "counter": 0,
                  "counter2": 0,
                  "counter2_stddev": 0,
                  "counter_stddev": 0,
                  "derivative": 0,
                  "derivative2": 0,
                  "derivative2_stddev": 0,
                  "derivative_stddev": 0,
                  "stddev": 0
          },
          {
                  "timestamp": 1.4718621e+09,
                  "value": 6.9220423e+07,
                  "count": 5,
                  "counter": 0,
                  "counter2": 0,
                  "counter2_stddev": 0,
                  "counter_stddev": 0,
                  "derivative": 0,
                  "derivative2": 0,
                  "derivative2_stddev": 0,
                  "derivative_stddev": 0,
                  "stddev": 0
          },
          {
                  "timestamp": 1.4718624e+09,
                  "value": 6.9220423e+07,
                  "count": 5,
                  "counter": 0,
                  "counter2": 0,
                  "counter2_stddev": 0,
                  "counter_stddev": 0,
                  "derivative": 0,
                  "derivative2": 0,
                  "derivative2_stddev": 0,
                  "derivative_stddev": 0,
                  "stddev": 0
          },
          {
                  "timestamp": 1.4718627e+09,
                  "value": 6.9220423e+07,
                  "count": 5,
                  "counter": 0,
                  "counter2": 0,
                  "counter2_stddev": 0,
                  "counter_stddev": 0,
                  "derivative": 0,
                  "derivative2": 0,
                  "derivative2_stddev": 0,
                  "derivative_stddev": 0,
                  "stddev": 0
          }
  ]
  $  
```
#
**graph** : Allows mass creation, editing and removal of graphs
 - **Query format** 
```json
  {"graph_id":"<graph_id>","title":"<title>","title_like":"<title_like>","tags_has":"<tag>"}
```
 - **Query field (key) description**
    - **graph_id** : The string value of _cid in a graph API object
    - **title** : The string value of title in a graph API object
    - **title_like** : String pattern to search all graphs that have titles containing the string pattern
    - **tags_has** : Tag string to search graphs that have matching tag
 - **Example**
```json
  $ circli -object graph -call get -where '{"graph_id":"98ff11f9-ff16-f74f-f446-f8f0f7533ff4"}'
  {
          "_cid": "/graph/98ff11f9-ff16-f74f-f446-f8f0f7533ff4",
          "access_keys": [],
          "composites": [],
          "datapoints": [
                  {
                          "alpha": "0.3",
                          "axis": "l",
                          "caql": null,
                          "check_id": 167984,
                          "color": "#657aa6",
                          "data_formula": "=-1 * VAL",
                          "derive": "counter",
                          "hidden": false,
                          "legend_formula": null,
                          "metric_name": "if`eth4`in_errors",
                          "metric_type": "numeric",
                          "name": "rx errors",
                          "stack": null
                  },
                  {
                          "alpha": "0.3",
                          "axis": "l",
                          "caql": null,
                          "check_id": 167984,
                          "color": "#4fa18e",
                          "data_formula": null,
                          "derive": "counter",
                          "hidden": false,
                          "legend_formula": null,
                          "metric_name": "if`eth4`out_errors",
                          "metric_type": "numeric",
                          "name": "tx errors",
                          "stack": null
                  }
          ],
          "description": "Network interface errors for eth4",
          "guides": [],
          "line_style": "interpolated",
          "logarithmic_left_y": null,
          "logarithmic_right_y": null,
          "max_left_y": null,
          "max_right_y": null,
          "metric_clusters": [],
          "min_left_y": null,
          "min_right_y": null,
          "notes": "cosi:register,cosi_id:7ff595f0-2414-444f-f34f-06ff6f61f683",
          "style": "area",
          "tags": [
                  "arch:x86_64",
                  "cosi:install",
                  "distro:redhat-5.1",
                  "os:linux"
          ],
          "title": "xyz-app-07 eth4 Errors"
  }
  $ 
```
#
**maintenance** : Schedule a maintenance window for your account, check bundle, rule set or host
 - **Query format**
```json
  {"maintenance_id":"<maintenance_id>","item":"<item>","type":"<type>","tags_has":"<tag>","start_ge":"<start>","stop_le":"<stop>"}
```
 - **Query field (key) description** 
    - **maintenance_id** : Integer part of the value of _cid in maintenance API object
    - **item** : String value of item in maintenance API object
    - **item_like** : Item string pattern to find maintenance API objects with matching item values 
    - **tags_has** : Tag string to search maintenance API objects that have matching tag
    - **start_ge** : Integer lower bound for the value of start in maintenance API object (ge = greater than or equal to)
    - **start_gt** : Integer lower bound for the value of start in maintenance API object (gt = greater than)
    - **stop_le** : Integer upper bound for the value if stop in maintenance API object (le = less than or equal to)
    - **stop_lt** : Integer upper bound for the value if stop in maintenance API object (lt = less than )
    - **type** : String value of type in maintenance API object
 - **Example**
```json
  $ circli -object maintenance -call get -where '{"maintenance_id":66799}'
  {
          "_cid": "/maintenance/66799",
          "item": "1.1.1.1",
          "notes": "Maintenance Window Test Through API",
          "severities": [
                  "1",
                  "2",
                  "3",
                  "4",
                  "5"
          ],
          "start": 1.4708559e+09,
          "stop": 1.471977817e+09,
          "tags": [
                  "web-cache:wash_202"
          ],
          "type": "host"
  }
  $ 
```
#
**metric_cluster** : A metric cluster is a cluster of metrics defined by a set of queries
 - **Query format**
```json
  {"metric_cluster_id":"<metric_cluster_id>","name":"<metric_cluster name>","name_like":"<metric_cluster_name_string_pattern>","tags_has":"<metric_cluster_tag>"}
```
 - **Query field (key) description**
    - **metric_cluster_id** : Integer part of the value of _cid in metric_cluster API object
    - **name** : String value of name in metric_cluster API object 
    - **name_like** : String pattern to search metric_clusters by matching name field values
    - **tags_has** : Tag string to search metric_clusters that have a matching tag
 - **Example**
```json
  $ circli -object metriccluster -call get -where '{"name_like":"202"}'
  [
          {
                  "_cid": "/metric_cluster/30517",
                  "description": "Compare all servers' CPU usage",
                  "name": "Wash 202 Servers CPU",
                  "queries": [
                          {
                                  "where": "cpu`idle 54.130",
                                  "type": "average"
                          }
                  ],
                  "tags": [],
                  "_matching_metrics": null
          },
          {
                  "_cid": "/metric_cluster/30518",
                  "description": "",
                  "name": "202 Server CPU Idle",
                  "queries": [
                          {
                                  "where": "(tag:cluster:202), cpuidle",
                                  "type": "counter"
                          }
                  ],
                  "tags": [],
                  "_matching_metrics": null
          },
          {
                  "_cid": "/metric_cluster/30519",
                  "description": "Sum of bps per server",
                  "name": "202 Server Egress",
                  "queries": [
                          {
                                  "where": "if`eth*`out_bytes (tags:cluster:202)",
                                  "type": "average"
                          }
                  ],
                  "tags": [],
                  "_matching_metrics": null
          }
  ]
  $ 
```
#
**rule_set** : define a collection of rules to apply to a given metric
 - **Query format**
```json
  {"name":"<rule_set_name>","check":"<check_id>"}
```
 - **Query field (key) description** 
    - **name** : The string value of _cid field in a rule_set API object
    - **check** : Integer part of the value of _cid in a check API object, to list rule set that are monitoring a particular check 
 - **Example**  
```json
  $ go run circli.go -object rule_set -call get -where '{"check":163862}'
  [
          {
                  "_cid": "/rule_set/163862_if`bond0`in_bytes",
                  "check": "/check/163862",
                  "contact_groups": {
                          "1": [],
                          "2": [],
                          "3": [],
                          "4": [],
                          "5": []
                  },
                  "derive": "mixed",
                  "link": "",
                  "metric_name": "if`bond0`in_bytes",
                  "metric_type": "numeric",
                  "notes": "Do something",
                  "parent": null,
                  "rules": [
                          {
                                  "criteria": "min value",
                                  "severity": 5,
                                  "transform": null,
                                  "transform_options": {},
                                  "transform_selection": null,
                                  "value": "40",
                                  "wait": 0,
                                  "windowing_duration": 300,
                                  "windowing_function": null
                          },
                          {
                                  "criteria": "on absence",
                                  "severity": 5,
                                  "transform": null,
                                  "transform_options": {},
                                  "transform_selection": null,
                                  "value": 120,
                                  "wait": 0,
                                  "windowing_duration": 300,
                                  "windowing_function": null
                          },
                          {
                                  "criteria": "max value",
                                  "severity": 5,
                                  "transform": null,
                                  "transform_options": {},
                                  "transform_selection": null,
                                  "value": "40",
                                  "wait": 0,
                                  "windowing_duration": 300,
                                  "windowing_function": null
                          },
                          {
                                  "criteria": "max value",
                                  "severity": 3,
                                  "transform": null,
                                  "transform_options": {
                                          "caql": "anomaly_detection%2850%2C%20model%3D%27auto%27%29",
                                          "model": "auto",
                                          "rt_analytics": "ad",
                                          "sensitivity": "50"
                                  },
                                  "transform_selection": null,
                                  "value": "99.99",
                                  "wait": 0,
                                  "windowing_duration": 60,
                                  "windowing_function": "average"
                          }
                  ]
          }
  ]
  $ 
```
#
**rule_set_group** : Allows togroup together rule sets and trigger alerts based on combinations of those rule sets faulting
 - **Query format**
```json
  {"rule_set_group_id":"<rule set group id>","name":"<rule_set_group name>","name_like":"<name like>","tags_has":"<tags has>"}
```
 - **Query field (key) description**
    - **rule_set_group_id** : Integer part of the value of _cid in rule_set_group API object
    - **name** : String value of name in rule_set_group API object
    - **name_like** : String pattern to search rule_set_groups by matching name field values
    - **tags_has** : Tag string to search rule_set_groups that have a matching tag
 - **Example**
```json
  $ circli -object rule_set_group -call get -where '{"name":"test"}'
  [
          {
                  "_cid": "/rule_set_group/175410",
                  "contact_groups": {
                          "1": [
                                  "/contact_group/2365"
                          ],
                          "2": [
                                  "/contact_group/2365"
                          ],
                          "3": [
                                  "/contact_group/2365"
                          ],
                          "4": [],
                          "5": []
                  },
                  "formulas": [
                          {
                                  "expression": 2,
                                  "raise_severity": 1,
                                  "wait": 1
                          }
                  ],
                  "name": "test",
                  "rule_set_conditions": [
                          {
                                  "matching_severities": [
                                          "3"
                                  ],
                                  "rule_set": "/rule_set/163862_if`bond0`in_bytes"
                          },
                          {
                                  "matching_severities": [
                                          "1"
                                  ],
                                  "rule_set": "/rule_set/167614_load`1min"
                          }
                  ]
          }
  ]
  $ 
```
#
**tags** : List all tags in use in your account (don't have any fields, Readonly)
 - **Query format**
```json
  {"tag_id":"<tag>"}
```
 - **Query field (key) description**
    - **tag_id**: Tag name string from the value of _cid in a Tag API object
 - **Example**
```json
 $ circli -object tag -call get -where '{"tag_id" : "haproxy"}'
 {
         "_cid": "/tag/haproxy"
 }
 $ 
```
#
**template** : A means to setup a mass number of checks quickly through both the API and UI
 - **Query format**
```json
  {"template_id":"<template id>","name":"<template name>","name_like":"<name pattern>","tags_has":"<tag>"}
```
 - **Query field (key) description**
    - **template_id** : Integer part of _cid value in template API object
  	- **name** : String value of name in template API object
  	- **name_like** : String pattern to search templates by matching name field values
  	- **tags_has** : Tag string to search templates that have a matching tag
 
 - **Example**
```json
  $ circli -object template -call get -where '{"tags_has":"cluster:web"}'
  [
          {
                  "_cid": "/template/970",
                  "_last_modified": 1474470113,
                  "_last_modified_by": "/user/5483",
                  "check_bundles": [
                          {
                                  "bundle_id": "/check_bundle/135930",
                                  "name": "xyz-web-01.wash.dc.xyz.net json"
                          }
                  ],
                    "hosts": [
                          "xyz-web-01.wash.dc.xyz.net",
                          "xyz-web-02.wash.dc.xyz.net",
                          "xyz-web-03.wash.dc.xyz.net"
                  ],
                  "master_host": "xyz-web-01.wash.dc.xyz.net",
                  "name": "test",
                  "notes": "",
                  "status": "active",
                  "tags": [
                          "cluster:web"
                  ]
          }
  ]
  $
```
#
**user** : Get basic information about the users on your account or a single user
 - **Query format**
```json
  {"user_id":"<user_id>","first_name":"<first_name>","last_name":"<last_name>","email":"<email_address>"}
```
 - **Query field (key) description**
    - **user_id** : Integer value of _cid in user API object
    - **first_name** : String value of firstname in user API object
    - **last_name** : String Value of lastname in user API object
    - **email** : String value of email in user API object
 - **Example**
```json
  $ circli -object user -call get -where '{"last_name":"Smith"}'
  [
          {
                  "_cid": "/user/5483",
                  "contact_info": {
                          "sms": "",
                          "xmpp": ""
                  },
                  "email": "alen_smith@xyz.com",
                  "firstname": "Alen",
                  "lastname": "Smith"
          }
  ]
  $ 
```
#
**worksheet** : Collection of graphs and allow quick correlation across them
 - **Query format**
```json
  {"worksheet_id":"<worksheet id>","title":"<worksheet title>","title_like":"<title string pattern>","tags_has":"<tag>"}
```
 - **Query field (key) description**
    - **worksheet_id** : Hex value of _cid in worksheet API object 
    - **title** : String value of title in worksheet API object
    - **title_like** : String pattern to search worksheets by matching title field values
    - **tags_has** : Tag string to search worksheets that have a matching tag
 - **Example**
```json
  $ circli -object worksheet -call get -where '{"title_like":"xyz-app-07"}'
  [
          {
                  "_cid": "/worksheet/f87f7957-7f3f-4f03-fff1-f0f8f5f126ff",
                  "description": "COSI worksheet for xyz-app-07",
                  "favorite": false,
                  "graphs": [],
                  "notes": "cosi:register,cosi_id:7ff595f0-2414-444f-f34f-06ff6f61f683",
                  "smart_queries": [
                          {
                                  "name": "Circonus One Step Install",
                                  "order": [],
                                  "where": "(notes:\"cosi:register,cosi_id:7ff595f0-2414-444f-f34f-06ff6f61f683*\")"
                          }
                  ],
                  "tags": [
                          "arch:x86_64",
                          "cosi:install",
                          "distro:redhat-5.1",
                          "os:linux"
                  ],
                  "title": "COSI xyz-app-07 1.2.3.4"
          }
  ]
  $ 
```
#
### List Calls
List calls are made to get the complete listing of all objects of a given type. 
#### CLI Syntax
```
circli -object <object_type> -call list
```
Possible Object Types : account, alert, annotation, broker, check, check_bundle, check_move, data, dashboard, graph, maintenance, metric_cluster, rule_set, rule_set_group, tag, template, user, worksheet 
####Example : List all annotations
```json
$ circli -object annotation -call list
[
        {
                "_cid": "/annotation/112378",
                "_created": 1.465228384e+09,
                "_last_modified": 1.465228384e+09,
                "_last_modified_by": "/user/2",
                "category": "samples",
                "description": "Hi there.",
                "start": 1.465226308e+09,
                "stop": 1.465227556e+09,
                "title": "Sample Annotation"
        },
        {
                "_cid": "/annotation/139620",
                "_created": 1.470837799e+09,
                "_last_modified": 1.470837799e+09,
                "_last_modified_by": "/user/5483",
                "category": "moratorium moratorium",
                "description": "No changes allowed",
                "start": 1.4701824e+09,
                "stop": 1.471823999e+09,
                "title": "Moratorium 2016"
        },
        {
                "_cid": "/annotation/139636",
                "_created": 1.470859743e+09,
                "_last_modified": 1.471271903e+09,
                "_last_modified_by": "/user/5483",
                "category": "server maintenance",
                "description": "Annotation update through API test",
                "start": 1.470853853e+09,
                "stop": 1.47204923e+09,
                "title": "Moratorium is extended"
        }
]
$ 

```
#
### Create Calls
#### CLI Syntax
```
circli -object <object_type> -call create [-where <json_string> | -file <json_file>]
```
#### Example : Create a graph
```json
$ cat new_graph.json 
{
    "access_keys": [],
    "composites": [],
    "datapoints": [
        {
            "legend_formula": null,
            "caql": null,
            "check_id": 167983,
            "metric_type": "numeric",
            "stack": null,
            "name": "% Used (bytes)",
            "axis": "l",
            "data_formula": null,
            "color": "#657aa6",
            "metric_name": "fs`/`df_used_percent",
            "alpha": null,
            "derive": "gauge",
            "hidden": false
        },
        {
            "legend_formula": null,
            "caql": null,
            "check_id": 167983,
            "metric_type": "numeric",
            "stack": null,
            "name": "% Used (inode)",
            "axis": "l",
            "data_formula": null,
            "color": "#4fa18e",
            "metric_name": "fs`/`df_used_inode_percent",
            "alpha": null,
            "derive": "gauge",
            "hidden": false
        }
    ],
    "description": "Filesystem space used and inodes used.",
    "guides": [],
    "line_style": "stepped",
    "logarithmic_left_y": null,
    "logarithmic_right_y": null,
    "max_left_y": "100",
    "max_right_y": null,
    "metric_clusters": [],
    "min_left_y": null,
    "min_right_y": null,
    "notes": "cosi:register,cosi_id:870604f6-9041-41ff-f1ff-9ff2ff0ff560",
    "style": "line",
    "tags": [
        "arch:x86_64",
        "cosi:install",
        "distro:redhat-5.1",
        "os:linux"
    ],
    "title": "Testing graph create through circli"
}
$ 
$ circli -object graph -call create -file new_graph.json 
{
        "access_keys": [],
        "composites": [],
        "_cid": "/graph/ffff2629-0999-fff8-f89f-fffff29f57f25",
        "style": "line",
        "description": null,
        "line_style": "stepped",
        "tags": [],
        "logarithmic_right_y": null,
        "logarithmic_left_y": null,
        "datapoints": [
                {
                        "legend_formula": null,
                        "caql": null,
                        "check_id": 167983,
                        "metric_type": "numeric",
                        "stack": null,
                        "name": "% Used (bytes)",
                        "axis": "l",
                        "data_formula": null,
                        "color": "#657aa6",
                        "metric_name": "fs`/`df_used_percent",
                        "alpha": null,
                        "derive": "gauge",
                        "hidden": false
                },
                {
                        "legend_formula": null,
                        "caql": null,
                        "check_id": 167983,
                        "metric_type": "numeric",
                        "stack": null,
                        "name": "% Used (inode)",
                        "axis": "l",
                        "data_formula": null,
                        "color": "#4fa18e",
                        "metric_name": "fs`/`df_used_inode_percent",
                        "alpha": null,
                        "derive": "gauge",
                        "hidden": false
                }
        ],
        "max_left_y": "100",
        "notes": null,
        "title": null,
        "max_right_y": null,
        "min_left_y": null,
        "guides": [],
        "metric_clusters": [],
        "min_right_y": null
}
$ 
```
#
### Update Calls
#### CLI Syntax
```
circli -object <object_type> -call update -oid <object_id> [-where <json_string> | -file <json_file> ]
```
#### Example 1: Update a check bundle with a new set of tags
```json
$ cat bundle_update.json 
{"tags": [
        "web-layer:server",
        "web-region:dc",
        "cluster:202",
        "api_create"
    ]}
$ 
$ circli -object check_bundle -call update -oid "137665"  -file bundle_update.json 
{
        "_created": 1471979522,
        "status": "active",
        "_reverse_connection_urls": [
                "mtev_reverse://1.1.1.1:43191/check/44582ff5-ff73-4ff3-ff31-f7f28f5659f9"
        ],
        "target": "1.1.1.1",
        "_checks": [
                "/check/169196"
        ],
        "timeout": 10,
        "metrics": [
                {
                        "status": "active",
                        "name": "FooDiskReadSessionsActive",
                        "type": "numeric",
                        "units": null,
                        "tags": []
                }],        
        "_cid": "/check_bundle/137665",
        "display_name": "xyz-app-03 Session/Cache Metrics",
        "tags": [
                "api_create",
                "web-layer:server",
                "web-region:dc",
                "cluster:202"
        ],
        "_check_uuids": [
                "44582ff5-ff73-4ff3-ff31-f7f28f5659f9"
        ],
        "notes": null,
        "type": "snmp",
        "config": {
                "oid_FooSessions_bps": "1.3.6.1.2.27.4.25.699.1.2.2.0",
                "oid_FooSessionsActive": "1.3.6.1.2.27.4.25.699.1.2.1.0",
                "oid_FooUniqueSessionsActive": "1.3.6.1.2.27.4.25.699.1.2.3.0",
                "oid_FooUniqueSessions_bps": "1.3.6.1.2.27.4.25.699.1.2.4.0",
                "reverse:secret_key": "ffffffff-ffff-ffff-ffff-ffffffffffff",
                "oid_FooDiskReadSessions_bps": "1.3.6.1.2.27.4.25.699.1.1.7.0",
                "community": "Net9M9mt",
                "oid_FooDiskReadSessionsActive": "1.3.6.1.2.27.4.25.699.1.1.6.0",
                "oid_FooCacheSessionsActive": "1.3.6.1.2.27.4.25.699.1.1.3.0",
                "auth_protocol": "MD5",
                "oid_FooCacheSessions_bps": "1.3.6.1.2.27.4.25.699.1.1.5.0",
                "security_level": "authPriv",
                "separate_queries": "false",
                "version": "2c",
                "port": "161",
                "privacy_protocol": "DES"
        }
}
$ 
```
#### Example 2: Update template and unbind a host
- before update
```json
{
        "_cid": "/template/970",
        "_last_modified": 1474470113,
        "_last_modified_by": "/user/5483",
        "check_bundles": [
                {
                        "bundle_id": "/check_bundle/135930",
                        "name": "xyz-web-01.wash.dc.xyz.net json"
                }
        ],
        "hosts": [
                "xyz-web-01.wash.dc.xyz.net",
                "xyz-web-02.wash.dc.xyz.net",
                "xyz-web-03.wash.dc.xyz.net"
        ],
        "master_host": "xyz-web-01.wash.dc.xyz.net",
        "name": "test",
        "notes": "",
        "status": "active",
        "tags": [
                "cluster:web"
        ]
}
```
- updating template with host unbinding (removing host *xyz-web-03.wash.dc.xyz.net*)
```json
$ cat template_update.json 
{
"hosts":[
                "xyz-web-01.wash.dc.xyz.net",
                "xyz-web-02.wash.dc.xyz.net"
        ],
"name":"test", 
"master_host":"xyz-web-01.wash.dc.xyz.net", 
"check_bundles":[
                        {       "bundle_id":"/check_bundle/135930",
                                "name":"xyz-web-01.wash.dc.xyz.net json"
                        }
                ]
}
$
```
- after updating
```json
$ circli -object template -call update -oid "970" -file template_update.json -host_removal_action "unbind"
{
        "hosts": [
                "xyz-web-01.wash.dc.xyz.net",
                "xyz-web-02.wash.dc.xyz.net"
        ],
        "_last_modified_by": "/user/5483",
        "status": "active",
        "name": "test",
        "_cid": "/template/970",
        "sync_rules": false,
        "tags": [
                "cluster:web"
        ],
        "master_host": "xyz-web-01.wash.dc.xyz.net",
        "notes": null,
        "check_bundles": [
                {
                        "bundle_id": "/check_bundle/135930",
                        "name": "xyz-web-01.wash.dc.xyz.net json"
                }
        ],
        "_last_modified": 1475166917
}
$
```
#
### Delete Calls
#### CLI Syntax
```
circli -object <object_type> -call delete -oid <object_id>
```
#### Example : Delete a graph
```json
$ circli -object graph -call delete -oid fff3219f-7fd5-f2f5-f0f1-f2ff4f291670
2016/08/25 20:50:47 graph delete :  fff3219f-7fd5-f2f5-f0f1-f2ff4f291670 
$ circli -object graph -call get -where '{"graph_id":"fff3219f-7fd5-f2f5-f0f1-f2ff4f291670"}'
{
        "_cid": "",
        "access_keys": null,
        "composites": null,
        "datapoints": null,
        "description": "",
        "guides": null,
        "line_style": null,
        "logarithmic_left_y": null,
        "logarithmic_right_y": null,
        "max_left_y": null,
        "max_right_y": null,
        "metric_clusters": null,
        "min_left_y": null,
        "min_right_y": null,
        "notes": null,
        "style": null,
        "tags": null,
        "title": ""
}
$ 
```