---
subcategory: "Load Balancing"
---

# Resource: lbprofile

The Lbprofile resource is used to set load balancing parameters in a profile.


## Example usage

```hcl
resource "citrixadc_lbprofile" "tf_lbprofile" {
    lbprofilename = "tf_lbprofile"
    dbslb = "ENABLED"
	processlocal = "DISABLED"
	httponlycookieflag = "ENABLED"
	lbhashfingers = 257
	lbhashalgorithm = "PRAC"
	storemqttclientidandusername = "YES"
}
```


## Argument Reference

* `lbprofilename` - (Required) Name of the LB profile.
* `dbslb` - (Optional) Enable database specific load balancing for MySQL and MSSQL service types. Possible values: [ ENABLED, DISABLED ]
* `processlocal` - (Optional) By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single pa cket request response mode or when the upstream device is performing a proper RSS for connection based distribution. Possible values: [ ENABLED, DISABLED ]
* `httponlycookieflag` - (Optional) Include the HttpOnly attribute in persistence cookies. The HttpOnly attribute limits the scope of a cookie to HTTP requests and helps mitigate the risk of cross-site scripting attacks. Possible values: [ ENABLED, DISABLED ]
* `cookiepassphrase` - (Optional) Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.
* `usesecuredpersistencecookie` - (Optional) Encode persistence cookie values using SHA2 hash. Possible values: [ ENABLED, DISABLED ]
* `useencryptedpersistencecookie` - (Optional) Encode persistence cookie values using SHA2 hash. Possible values: [ ENABLED, DISABLED ]
* `literaladccookieattribute` - (Optional) String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence , GSLB site persistence, CS cookie persistence, LB group cookie persistence). Sample usage - add lb profile lbprof -LiteralADCCookieAttribute ";SameSite=None".
* `computedadccookieattribute` - (Optional) ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set. Sample usage - add ns variable lbvar -type TEXT(100) -scope Transaction add ns assignment lbassign -variable $lbvar -set "\\";SameSite=Strict\\"" add rewrite policy lbpol <valid policy expression> lbassign bind rewrite global lbpol 100 next -type RES_OVERRIDE add lb profile lbprof -ComputedADCCookieAttribute "$lbvar" For incoming client request, if above policy evaluates TRUE, then SameSite=Strict will be appended to ADC generated cookie.
* `storemqttclientidandusername` - (Optional) This option allows to store the MQTT clientid and username in transactional logs.
Default value: NO
Possible values : [YES, NO]
* `lbhashalgorithm` - (Optional) This option dictates the hashing algorithm used for hash based LB methods (URLHASH, DOMAINHASH, SOURCEIPHASH, DESTINATIONIPHASH, SRCIPDESTIPHASH, SRCIPSRCPORTHASH, TOKEN, USER_TOKEN, CALLIDHASH).
Default value: [DEFAULT]
Possible values : [DEFAULT, PRAC, JARH]
* `lbhashfingers` - (Optional) This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory.
Default value: 256
Minimum value = 1
Maximum value = 1024

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbprofile. It has the same value as the `lbprofilename` attribute.


## Import

A lbprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lbprofile.tf_lbprofile tf_lbprofile
```
