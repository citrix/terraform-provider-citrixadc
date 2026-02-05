---
subcategory: "Traffic Management"
---

# Data Source: tmtrafficaction

The tmtrafficaction data source allows you to retrieve information about a TM traffic action.

## Example usage

```terraform
data "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
  name = "my_traffic_action"
}

output "apptimeout" {
  value = data.citrixadc_tmtrafficaction.tf_tmtrafficaction.apptimeout
}

output "sso" {
  value = data.citrixadc_tmtrafficaction.tf_tmtrafficaction.sso
}
```

## Argument Reference

* `name` - (Required) Name for the traffic action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `apptimeout` - Time interval, in minutes, of user inactivity after which the connection is closed.
* `forcedtimeout` - Setting to start, stop or reset TM session force timer.
* `forcedtimeoutval` - Time interval, in minutes, for which force timer should be set.
* `formssoaction` - Name of the configured form-based single sign-on profile.
* `initiatelogout` - Initiate logout for the traffic management (TM) session if the policy evaluates to true. The session is then terminated after two minutes.
* `kcdaccount` - Kerberos constrained delegation account name.
* `passwdexpression` - expression that will be evaluated to obtain password for SingleSignOn.
* `persistentcookie` - Use persistent cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.
* `samlssoprofile` - Profile to be used for doing SAML SSO to remote relying party.
* `sso` - Use single sign-on for the resource that the user is accessing now.
* `userexpression` - expression that will be evaluated to obtain username for SingleSignOn.

## Attribute Reference

* `id` - The id of the tmtrafficaction. It has the same value as the `name` attribute.

## Import

A tmtrafficaction can be imported using its name, e.g.

```shell
terraform import citrixadc_tmtrafficaction.tf_tmtrafficaction my_traffic_action
```
