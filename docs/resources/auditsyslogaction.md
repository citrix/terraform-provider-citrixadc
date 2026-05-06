---
subcategory: "Audit"
---

# Resource: auditsyslogaction

This resource is used to create audit syslog actions.


## Example usage

### Using httpauthtoken (sensitive attribute - persisted in state)

```hcl
variable "auditsyslogaction_httpauthtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name            = "tf_syslogaction"
  serverip        = "10.78.60.33"
  serverport      = 514
  transport       = "HTTP"
  httpendpointurl = "https://logs.example.com/ingest"
  httpauthtoken   = var.auditsyslogaction_httpauthtoken
  loglevel        = ["ERROR", "NOTICE"]
}
```

### Using httpauthtoken_wo (write-only/ephemeral - NOT persisted in state)

The `httpauthtoken_wo` attribute provides an ephemeral path for the HTTP authentication token. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `httpauthtoken_wo_version`.

```hcl
variable "auditsyslogaction_httpauthtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name                     = "tf_syslogaction"
  serverip                 = "10.78.60.33"
  serverport               = 514
  transport                = "HTTP"
  httpendpointurl          = "https://logs.example.com/ingest"
  httpauthtoken_wo         = var.auditsyslogaction_httpauthtoken
  httpauthtoken_wo_version = 1
  loglevel                 = ["ERROR", "NOTICE"]
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name                     = "tf_syslogaction"
  serverip                 = "10.78.60.33"
  serverport               = 514
  transport                = "HTTP"
  httpendpointurl          = "https://logs.example.com/ingest"
  httpauthtoken_wo         = var.auditsyslogaction_httpauthtoken
  httpauthtoken_wo_version = 2  # Bumped to trigger update
  loglevel                 = ["ERROR", "NOTICE"]
}
```


## Argument Reference

* `name` - (Required) Name of the syslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog action is added. Changing this attribute forces a new resource to be created.
* `loglevel` - (Required) Audit log level, which specifies the types of events to log. Possible values: [ ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE ]
* `acl` - (Optional) Log access control list (ACL) messages. Possible values: [ ENABLED, DISABLED ]
* `alg` - (Optional) Log alg info. Possible values: [ ENABLED, DISABLED ]
* `appflowexport` - (Optional) Export log messages to AppFlow collectors. Appflow collectors are entities to which log messages can be sent so that some action can be performed on them. Possible values: [ ENABLED, DISABLED ]
* `contentinspectionlog` - (Optional) Log Content Inspection event information. Possible values: [ ENABLED, DISABLED ]
* `dateformat` - (Optional) Format of dates in the logs. Possible values: [ MMDDYYYY, DDMMYYYY, YYYYMMDD ]
* `dns` - (Optional) Log DNS related syslog messages. Possible values: [ ENABLED, DISABLED ]
* `domainresolvenow` - (Optional) Immediately send a DNS query to resolve the server's domain name.
* `domainresolveretry` - (Optional) Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.
* `httpauthtoken` - (Optional, Sensitive) Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorization header is required to be of the form - Splunk <auth-token>. The value is persisted in Terraform state (encrypted). See also `httpauthtoken_wo` for an ephemeral alternative.
* `httpauthtoken_wo` - (Optional, Sensitive, WriteOnly) Same as `httpauthtoken`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `httpauthtoken_wo_version`. If both `httpauthtoken` and `httpauthtoken_wo` are set, `httpauthtoken_wo` takes precedence.
* `httpauthtoken_wo_version` - (Optional) An integer version tracker for `httpauthtoken_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `httpendpointurl` - (Optional) The URL at which to upload the logs messages on the endpoint.
* `lbvservername` - (Optional) Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName.
* `logfacility` - (Optional) Facility value, as defined in RFC 3164, assigned to the log message. Possible values: [ LOCAL0, LOCAL1, LOCAL2, LOCAL3, LOCAL4, LOCAL5, LOCAL6, LOCAL7 ]
* `lsn` - (Optional) Log lsn info. Possible values: [ ENABLED, DISABLED ]
* `managementlog` - (Optional) Management log specifies the categories of log files to be exported. It uses destination and transport from PE params. Possible values: [ ALL, SHELL, ACCESS, NSMGMT, NONE ]
* `maxlogdatasizetohold` - (Optional) Max size of log data that can be held in NSB chain of server info.
* `mgmtloglevel` - (Optional) Management log level, which specifies the types of events to log. Possible values: [ ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE ]
* `netprofile` - (Optional) Name of the network profile. The SNIP configured in the network profile will be used as source IP while sending log messages.
* `protocolviolations` - (Optional) Log protocol violations.
* `serverdomainname` - (Optional) SYSLOG server name as a FQDN. Mutually exclusive with serverIP/lbVserverName.
* `serverip` - (Optional) IP address of the syslog server.
* `serverport` - (Optional) Port on which the syslog server accepts connections.
* `sslinterception` - (Optional) Log SSL Interception event information. Possible values: [ ENABLED, DISABLED ]
* `streamanalytics` - (Optional) Export log stream analytics statistics to syslog server.
* `subscriberlog` - (Optional) Log subscriber session event information. Possible values: [ ENABLED, DISABLED ]
* `syslogcompliance` - (Optional) Setting this parameter ensures that all the Audit Logs generated for this Syslog Action comply with an RFC. For example, set it to RFC5424 to ensure RFC 5424 compliance.
* `tcp` - (Optional) Log TCP messages. Possible values: [ NONE, ALL ]
* `tcpprofilename` - (Optional) Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.
* `timezone` - (Optional) Time zone used for date and timestamps in the logs. Possible values: [ GMT_TIME, LOCAL_TIME ]
* `transport` - (Optional) Transport type used to send auditlogs to syslog server. Default type is UDP. Possible values: [ TCP, UDP, HTTP ]. Changing this attribute forces a new resource to be created.
* `urlfiltering` - (Optional) Log URL filtering event information. Possible values: [ ENABLED, DISABLED ]
* `userdefinedauditlog` - (Optional) Log user-configurable log messages to syslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditsyslogaction. It has the same value as the `name` attribute.


## Import

An instance of the resource can be imported using its name, e.g.

```shell
terraform import citrixadc_auditsyslogaction.tf_syslogaction tf_syslogaction
```
