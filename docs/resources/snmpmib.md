---
subcategory: "SNMP"
---

# Resource: snmpmib

The snmpmib resource is used to create snmpmib.


## Example usage

```hcl
resource "citrixadc_snmpmib" "tf_snmpmib" {
  contact  = "phone_number"
  name     = "my_name"
  location = "LOCATION"
  customid = "CUSTOMER_ID"
}

```


## Argument Reference

* `contact` - (Optional) Name of the administrator for this Citrix ADC. Along with the name, you can include information on how to contact this person, such as a phone number or an email address. Can consist of 1 to 127 characters that include uppercase and  lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  The following requirement applies only to the Citrix ADC CLI: If the information includes one or more spaces, enclose it in double or single quotation marks (for example, "my contact" or 'my contact').
* `customid` - (Optional) Custom identification number for the Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a custom identification that helps identify the Citrix ADC appliance.  The following requirement applies only to the Citrix ADC CLI: If the ID includes one or more spaces, enclose it in double or single quotation marks (for example, "my ID" or 'my ID').
* `location` - (Optional) Physical location of the Citrix ADC. For example, you can specify building name, lab number, and rack number. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  The following requirement applies only to the Citrix ADC CLI: If the location includes one or more spaces, enclose it in double or single quotation marks (for example, "my location" or 'my location').
* `name` - (Optional) Name for this Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the Citrix ADC appliance.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my name" or 'my name').
* `ownernode` - (Optional) ID of the cluster node for which we are setting the mib. This is a mandatory argument to set snmp mib on CLIP.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpmib. It is a unique string prefixed with "tf-snmpmib-".

