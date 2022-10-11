<!-- ## Citrix ADC  -->
## Configure a secure content switching server in Citrix ADC

## Structure
* `main.tf` describes the actual Citrix ADC config objects to be created. The attributes of these resources are either hard coded or looked up from input variables in `examples.tfvars`
* `variables.tf` describes the input variables to the terraform config. These can have defaults
* `provider.tf` is used to specify the username, password and endpoint of the Citrix ADC. Alternatively, you can set the NS_URL, NS_LOGIN and NS_PASSWORD environment variables.
* `examples.tfvars` has the variable inputs specified in `variables.tf`

## Using
Modify the `examples.tfvars` and `provider.tf` to suit your own Citrix ADC deployment. Use `terraform plan -var-file examples.tfvars` and `terraform apply -var-file examples.tfvars` to configure the Citrix ADC.

## Updating your configuration
Modify the set of backend services and use `terraform plan -var-file examples.tfvars` and `terraform apply -var-file examples.tfvars` to verify the changes

## Destroying your Configuration
To destroy the configuration that you built now use `terraform destroy -var-file examples.tfvars` to destroy the configuration.