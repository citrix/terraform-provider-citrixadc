Etrade Specific FrontEnd Cipher group
===================================================


Requirements
------------
Terraform version >= 1.1.0 (this module developed with version v1.1.5)
-	[Terraform](https://www.terraform.io/downloads.html)

Citrix ADC Terraform Provider 
-	https://github.com/citrix/terraform-provider-citrixadc


Building The Modules
---------------------
This child module creates below resources  -
-	Creates specific cipher group
-	Bind require ciphersuites to group.


Variable inputs
----------------------
The require variable values need to provide in resources.tf file, then call the module in root module.




