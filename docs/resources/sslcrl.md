---
subcategory: "SSL"
---

# Resource: sslcrl

The sslcrl resource is used to Configure Certificate Revocation List resource


## Example usage

```hcl
resource "citrixadc_sslcrl" "tf_sslcrl" {
  crlname = "tf_sslcrl"
  crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
  cacert  = "rootrsa_cert1"
}
```


## Argument Reference

* `crlname` - (Required) Name for the Certificate Revocation List (CRL). Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the CRL is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my crl" or 'my crl'). Minimum length =  1
* `crlpath` - (Required) Path to the CRL file. /var/netscaler/ssl/ is the default path. Minimum length =  1
* `inform` - (Optional) Input format of the CRL file. The two formats supported on the appliance are: PEM - Privacy Enhanced Mail. DER - Distinguished Encoding Rule. Possible values: [ DER, PEM ]
* `refresh` - (Optional) Set CRL auto refresh. Possible values: [ ENABLED, DISABLED ]
* `cacert` - (Optional) CA certificate that has issued the CRL. Required if CRL Auto Refresh is selected. Install the CA certificate on the appliance before adding the CRL. Minimum length =  1
* `method` - (Optional) Method for CRL refresh. If LDAP is selected, specify the method, CA certificate, base DN, port, and LDAP server name. If HTTP is selected, specify the CA certificate, method, URL, and port. Cannot be changed after a CRL is added. Possible values: [ HTTP, LDAP ]
* `server` - (Optional) IP address of the LDAP server from which to fetch the CRLs. Minimum length =  1
* `url` - (Optional) URL of the CRL distribution point.
* `port` - (Optional) Port for the LDAP server. Minimum value =  1
* `basedn` - (Optional) Base distinguished name (DN), which is used in an LDAP search to search for a CRL. Citrix recommends searching for the Base DN instead of the Issuer Name from the CA certificate, because the Issuer Name field might not exactly match the LDAP directory structure's DN. Minimum length =  1
* `scope` - (Optional) Extent of the search operation on the LDAP server. Available settings function as follows: One - One level below Base DN. Base - Exactly the same level as Base DN. Possible values: [ Base, One ]
* `interval` - (Optional) CRL refresh interval. Use the NONE setting to unset this parameter. Possible values: [ MONTHLY, WEEKLY, DAILY, NOW, NONE ]
* `day` - (Optional) Day on which to refresh the CRL, or, if the Interval parameter is not set, the number of days after which to refresh the CRL. If Interval is set to MONTHLY, specify the date. If Interval is set to WEEKLY, specify the day of the week (for example, Sun=0 and Sat=6). This parameter is not applicable if the Interval is set to DAILY. Minimum value =  0 Maximum value =  31
* `time` - (Optional) Time, in hours (1-24) and minutes (1-60), at which to refresh the CRL.
* `binddn` - (Optional) Bind distinguished name (DN) to be used to access the CRL object in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed. Minimum length =  1
* `password` - (Optional) Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed. Minimum length =  1
* `binary` - (Optional) Set the LDAP-based CRL retrieval mode to binary. Possible values: [ YES, NO ]
* `cacertfile` - (Optional) Name of and, optionally, path to the CA certificate file. /nsconfig/ssl/ is the default path. Maximum length =  63
* `cakeyfile` - (Optional) Name of and, optionally, path to the CA key file. /nsconfig/ssl/ is the default path. Maximum length =  63
* `indexfile` - (Optional) Name of and, optionally, path to the file containing the serial numbers of all the certificates that are revoked. Revoked certificates are appended to the file. /nsconfig/ssl/ is the default path. Maximum length =  63
* `revoke` - (Optional) Name of and, optionally, path to the certificate to be revoked. /nsconfig/ssl/ is the default path. Maximum length =  63
* `gencrl` - (Optional) Name of and, optionally, path to the CRL file to be generated. The list of certificates that have been revoked is obtained from the index file. /nsconfig/ssl/ is the default path. Maximum length =  63


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcrl. It has the same value as the `crlname` attribute.


## Import

A sslcrl can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcrl.tf_sslcrl tf_sslcrl
```
