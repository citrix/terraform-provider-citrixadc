---
subcategory: "Audit"
---

# Data Source: citrixadc_auditnslogaction

Use this data source to retrieve information about an existing Audit NS Log Action.

The `citrixadc_auditnslogaction` data source allows you to retrieve details of an audit nslog action by its name. This is useful for referencing existing audit nslog actions in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing audit nslog action
data "citrixadc_auditnslogaction" "example" {
  name = "my_nslog_action"
}

# Use the retrieved action data in a policy
resource "citrixadc_auditnslogpolicy" "example_policy" {
  name   = "example_policy"
  rule   = "true"
  action = data.citrixadc_auditnslogaction.example.name
}

# Reference action attributes
output "server_ip" {
  value = data.citrixadc_auditnslogaction.example.serverip
}

output "log_levels" {
  value = data.citrixadc_auditnslogaction.example.loglevel
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the nslog action to retrieve. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the audit nslog action (same as name).

### Server Configuration

* `serverip` - IP address of the nslog server.

* `serverport` - Port on which the nslog server accepts connections.

* `serverdomainname` - Auditserver name as a FQDN. Mutually exclusive with serverIP.

* `domainresolvenow` - Immediately send a DNS query to resolve the server's domain name.

* `domainresolveretry` - Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the audit server if the last query failed.

### Log Configuration

* `loglevel` - Audit log level, which specifies the types of events to log. Available settings:
  * `ALL` - All events.
  * `EMERGENCY` - Events that indicate an immediate crisis on the server.
  * `ALERT` - Events that might require action.
  * `CRITICAL` - Events that indicate an imminent server crisis.
  * `ERROR` - Events that indicate some type of error.
  * `WARNING` - Events that require action in the near future.
  * `NOTICE` - Events that the administrator should know about.
  * `INFORMATIONAL` - All but low-level events.
  * `DEBUG` - All events, in extreme detail.
  * `NONE` - No events.

* `logfacility` - Facility value, as defined in RFC 3164, assigned to the log message. Log facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.

* `dateformat` - Format of dates in the logs. Supported formats:
  * `MMDDYYYY` - U.S. style month/date/year format.
  * `DDMMYYYY` - European style date/month/year format.
  * `YYYYMMDD` - ISO style year/month/date format.

* `timezone` - Time zone used for date and timestamps in the logs:
  * `GMT_TIME` - Coordinated Universal Time.
  * `LOCAL_TIME` - The server's timezone setting.

### Event Type Configuration

* `tcp` - Log TCP messages. Possible values: `NONE`, `ALL`.

* `acl` - Log access control list (ACL) messages. Possible values: `ENABLED`, `DISABLED`.

* `alg` - Log the ALG messages. Possible values: `ENABLED`, `DISABLED`.

* `appflowexport` - Export log messages to AppFlow collectors. AppFlow collectors are entities to which log messages can be sent so that some action can be performed on them. Possible values: `ENABLED`, `DISABLED`.

* `contentinspectionlog` - Log Content Inspection event information. Possible values: `ENABLED`, `DISABLED`.

* `lsn` - Log the LSN messages. Possible values: `ENABLED`, `DISABLED`.

* `protocolviolations` - Log protocol violations. Possible values: `NONE`, `ALL`.

* `sslinterception` - Log SSL Interception event information. Possible values: `ENABLED`, `DISABLED`.

* `subscriberlog` - Log subscriber session event information. Possible values: `ENABLED`, `DISABLED`.

* `urlfiltering` - Log URL filtering event information. Possible values: `ENABLED`, `DISABLED`.

* `userdefinedauditlog` - Log user-configurable log messages to nslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria. Possible values: `YES`, `NO`.

## Notes

* Audit nslog actions define how audit logs are sent to external syslog servers.
* The action specifies the server IP/FQDN, port, log levels, and what types of events to log.
* When using with policies, the action determines where and how logs are sent when the policy rule matches.
* Either `serverip` or `serverdomainname` must be specified, but not both.
* The `loglevel` attribute is a list that can contain multiple log levels to enable.
