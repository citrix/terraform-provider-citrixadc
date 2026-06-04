---
subcategory: "AAA"
---

# Data Source: aaasession

Retrieves information about an active AAA-TM/VPN session on the Citrix ADC without terminating it. The data source queries the NITRO `get(all)` endpoint and returns the first active session that matches every supplied filter (user, group, intranet IP range, session key, or cluster node).

Use this data source to inspect live session state. To terminate sessions, use the `citrixadc_aaasession` resource instead.


## Example usage

Look up the active session for a specific user:

```hcl
data "citrixadc_aaasession" "by_user" {
  username = "jdoe"
}

output "aaasession_groupname" {
  value = data.citrixadc_aaasession.by_user.groupname
}
```

Look up a session by its intranet IP range:

```hcl
data "citrixadc_aaasession" "by_iprange" {
  iip     = "10.102.1.0"
  netmask = "255.255.255.0"
}
```


## Argument Reference

All arguments are optional filters. The data source returns the first active session that matches every filter you supply. If no session matches, an error is raised.

* `all` - (Optional) Filter on the all flag.
* `username` - (Optional) Name of the AAA user.
* `groupname` - (Optional) Name of the AAA group.
* `iip` - (Optional) IP address or the first address in the intranet IP range.
* `netmask` - (Optional) Subnet mask for the intranet IP range specified by `iip`.
* `sessionkey` - (Optional) Show the AAA session associated with the given session key.
* `nodeid` - (Optional) Unique number that identifies the cluster node. This is a GET-only cluster filter, valid for reading sessions but not for the kill action exposed by the resource.


## Attribute Reference

In addition to the arguments, the following read-only attributes are populated from the matched session. Any argument left unset in the configuration is also exported with the value returned by the ADC:

* `id` - A synthetic identifier with the constant value `"aaasession-query"`.
* `all` - The all flag of the matched session.
* `username` - Name of the AAA user for the matched session.
* `groupname` - Name of the AAA group for the matched session.
* `iip` - IP address or the first address in the intranet IP range for the matched session.
* `netmask` - Subnet mask for the intranet IP range of the matched session.
* `sessionkey` - Session key of the matched AAA session.
* `nodeid` - Cluster node identifier for the matched session.
