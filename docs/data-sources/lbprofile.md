---
subcategory: "Load Balancing"
---

# Data Source `lbprofile`

The lbprofile data source allows you to retrieve information about a load balancing profile.


## Example usage

```terraform
data "citrixadc_lbprofile" "tf_lbprofile" {
  lbprofilename = "my_lbprofile"
}

output "dbslb" {
  value = data.citrixadc_lbprofile.tf_lbprofile.dbslb
}

output "httponlycookieflag" {
  value = data.citrixadc_lbprofile.tf_lbprofile.httponlycookieflag
}
```


## Argument Reference

* `lbprofilename` - (Required) Name of the LB profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbprofile. It has the same value as the `lbprofilename` attribute.
* `computedadccookieattribute` - ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set.
* `cookiepassphrase` - Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.
* `dbslb` - Enable database specific load balancing for MySQL and MSSQL service types.
* `httponlycookieflag` - Include the HttpOnly attribute in persistence cookies. The HttpOnly attribute limits the scope of a cookie to HTTP requests and helps mitigate the risk of cross-site scripting attacks.
* `lbhashalgorithm` - This option dictates the hashing algorithm used for hash based LB methods (URLHASH, DOMAINHASH, SOURCEIPHASH, DESTINATIONIPHASH, SRCIPDESTIPHASH, SRCIPSRCPORTHASH, TOKEN, USER_TOKEN, CALLIDHASH).
* `lbhashfingers` - This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory.
* `literaladccookieattribute` - String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence , GSLB site persistence, CS cookie persistence, LB group cookie persistence).
* `processlocal` - By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.
* `proximityfromself` - Use the ADC location instead of client IP for static proximity LB or GSLB decision.
* `storemqttclientidandusername` - This option allows to store the MQTT clientid and username in transactional logs.
* `useencryptedpersistencecookie` - Encode persistence cookie values using SHA2 hash.
* `usesecuredpersistencecookie` - Encode persistence cookie values using SHA2 hash.


## Import

A lbprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lbprofile.tf_lbprofile my_lbprofile
```
