---
subcategory: "Audit"
---

# Data Source: citrixadc_auditmessageaction

Use this data source to retrieve information about an existing Audit Message Action.

The `citrixadc_auditmessageaction` data source allows you to retrieve details of an audit message action by its name. This is useful for referencing existing audit message actions in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing audit message action
data "citrixadc_auditmessageaction" "example" {
  name = "my_audit_action"
}

# Use the retrieved action data in a policy
resource "citrixadc_auditmessagepolicy" "example_policy" {
  name   = "example_policy"
  rule   = "true"
  action = data.citrixadc_auditmessageaction.example.name
}

# Reference action attributes
output "log_level" {
  value = data.citrixadc_auditmessageaction.example.loglevel
}

output "log_expression" {
  value = data.citrixadc_auditmessageaction.example.stringbuilderexpr
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the audit message action to retrieve. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the audit message action (same as name).

### Action Configuration

* `loglevel` - Audit log level, which specifies the severity level of the log message being generated. The following loglevels are valid:
  * `EMERGENCY` - Events that indicate an immediate crisis on the server.
  * `ALERT` - Events that might require action.
  * `CRITICAL` - Events that indicate an imminent server crisis.
  * `ERROR` - Events that indicate some type of error.
  * `WARNING` - Events that require action in the near future.
  * `NOTICE` - Events that the administrator should know about.
  * `INFORMATIONAL` - All but low-level events.
  * `DEBUG` - All events, in extreme detail.

* `stringbuilderexpr` - Default-syntax expression that defines the format and content of the log message. This expression can use string operations, HTTP headers, and other data from the request/response to build custom log messages.

* `logtonewnslog` - Send the message to the new nslog. Possible values: `YES`, `NO`.

* `bypasssafetycheck` - Bypass the safety check and allow unsafe expressions. This should be used with caution as it allows expressions that might have security implications. Possible values: `YES`, `NO`.

## Notes

* Audit message actions are used to define custom log messages that can be generated during policy evaluation.
* The `stringbuilderexpr` parameter accepts default syntax expressions, allowing you to create dynamic log messages based on request/response data.
* When using with policies, the action defines what log message to generate when the policy rule matches.
