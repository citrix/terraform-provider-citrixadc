## NetScaler configuration examples
The subfolders have examples of terraform configurations

## Structure
* `resources.tf` describes the actual NetScaler config objects to be created. The attributes of these resources are either hard coded or looked up from input variables in `terraform.tfvars`
* `variables.tf` describes the input variables to the terraform config. These can have defaults
* `provider.tf` is used to specify the username, password and endpoint of the NetScaler. Alternatively, you can set the NS_URL, NS_LOGIN and NS_PASSWORD environment variables.
* `terraform.tfvars` has the variable inputs specified in `variables.tf`

## Using
Modify the `terraform.tfvars` and `provider.tf` to suit your own NetScaler deployment. Use `terraform plan` and `terraform apply` to configure the NetScaler.

## Updating your configuration
Modify the set of backend services and use `terraform plan` and `terraform apply` to verify the changes

