---
subcategory: "Audit"
---

# Data Source: citrixadc_auditnslogparams

The `citrixadc_auditnslogparams` data source allows you to retrieve information about the global Audit NS Log parameters configuration. These parameters control the behavior of audit logging to NS Log servers on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_auditnslogparams" "tf_auditnslogparams" {
}

output "dateformat" {
  value = data.citrixadc_auditnslogparams.tf_auditnslogparams.dateformat
}

output "loglevel" {
  value = data.citrixadc_auditnslogparams.tf_auditnslogparams.loglevel
}

output "tcp_logging" {
  value = data.citrixadc_auditnslogparams.tf_auditnslogparams.tcp
}
```

## Argument Reference

This datasource does not require any arguments. It retrieves the global audit nslog parameters configuration.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the auditnslogparams. It is a system-generated identifier.

### Server Configuration

* `serverip` - IP address of the nslog server.

* `serverport` - Port on which the nslog server accepts connections.

### Log Configuration

* `loglevel` - Types of information to be logged. Available settings:
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

* `timezone` - Time zone used for date and timestamps in the logs. Supported settings:
  * `GMT_TIME` - Coordinated Universal Time.
  * `LOCAL_TIME` - Use the server's timezone setting.

### Feature-Specific Logging

* `acl` - Configure auditing to log access control list (ACL) messages.

* `alg` - Log the ALG messages.

* `tcp` - Configure auditing to log TCP messages.

* `protocolviolations` - Log protocol violations.

* `lsn` - Log the LSN messages.

* `contentinspectionlog` - Log Content Inspection event information.

* `sslinterception` - Log SSL Interception event information.

* `subscriberlog` - Log subscriber session event information.

* `urlfiltering` - Log URL filtering event information.

* `userdefinedauditlog` - Log user-configurable log messages to nslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.

### AppFlow Integration

* `appflowexport` - Export log messages to AppFlow collectors. AppFlow collectors are entities to which log messages can be sent so that some action can be performed on them.
