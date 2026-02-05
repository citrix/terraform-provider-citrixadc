---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationnegotiateaction

The `citrixadc_authenticationnegotiateaction` data source is used to retrieve information about an existing Authentication Negotiate Action configured on a Citrix ADC appliance. This action represents the AD KDC server profile used for Kerberos/NTLM authentication.

## Example usage

```hcl
# Retrieve an authentication negotiate action by name
data "citrixadc_authenticationnegotiateaction" "example" {
  name = "demo_negotiateaction"
}

# Use the retrieved data in other resources
output "action_domain" {
  value = data.citrixadc_authenticationnegotiateaction.example.domain
}

output "action_ntlmpath" {
  value = data.citrixadc_authenticationnegotiateaction.example.ntlmpath
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the AD KDC server profile (negotiate action) to retrieve. This is the unique identifier for the action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the authentication negotiate action. It has the same value as the `name` attribute.
* `defaultauthenticationgroup` - The default group that is chosen when the authentication succeeds in addition to extracted groups.
* `domain` - Domain name of the service principal that represents Citrix ADC.
* `domainuser` - User name of the account that is mapped with Citrix ADC principal. This can be given along with domain and password when keytab file is not available. If username is given along with keytab file, then that keytab file will be searched for this user's credentials.
* `domainuserpasswd` - Password of the account that is mapped to the Citrix ADC principal.
* `keytab` - The path to the keytab file that is used to decrypt kerberos tickets presented to Citrix ADC. If keytab is not available, domain/username/password can be specified in the negotiate action configuration.
* `ntlmpath` - The path to the site that is enabled for NTLM authentication, including FQDN of the server. This is used when clients fallback to NTLM.
* `ou` - Active Directory organizational units (OU) attribute.
