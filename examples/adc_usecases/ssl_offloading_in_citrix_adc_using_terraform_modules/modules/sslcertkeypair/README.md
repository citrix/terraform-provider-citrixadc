Install SSL certificate
======================================================


Requirements
------------
Terraform version >= 1.1.0 (this module developed with version v1.1.5)
-	[Terraform](https://www.terraform.io/downloads.html)

Citrix ADC Terraform Provider 
-	https://github.com/citrix/terraform-provider-citrixadc


Building The Modules
---------------------
This child module creates below resources  -
-	Creates SSL Certkey pair

Note: This child module take ssl files uploaded in /nsconfig/ssl manually and install the certificate.


Variable inputs
----------------------
The require variable values need to provide in terraform.tfvars to create resources in root module.



