Load Balancing Services with specified service port
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
-	Creates load balancing servers
-	Creates require load balancing services on given dedicated port
-	Set the require service parameters 


Variable inputs
----------------------
The require variable values need to provide in variables.tf and terraform.tfvars to create resources in root module.




