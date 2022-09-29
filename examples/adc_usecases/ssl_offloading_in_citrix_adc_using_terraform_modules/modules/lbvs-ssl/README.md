SSL virtual server configuration and parameter settings (Child configuration module)
===================================================


Requirements
------------
Terraform version >= 1.1.0 (this module developed with version v1.1.5)
-	[Terraform](https://www.terraform.io/downloads.html)

Citrix ADC Terraform provider files
-   [Citrix ADC Terraform Provider](https://github.com/citrix/terraform-provider-citrixadc)



Building The Modules
---------------------
This child configuration module helps to creates below resources in root/parent module 
-	Creates load balancing SSL virtual server on specified service port
-	Set the require SSL virtual server parameters 
-   The require parameter attributes need to be enabled or disabled as needed.


Variable inputs
----------------------
- The require variable, ssl virtual server parameter values need to provide in variables.tf and terraform.tfvars in root/parent module to create resources in root module.
- provide require parameters values in terraform.tfvars of root module.



