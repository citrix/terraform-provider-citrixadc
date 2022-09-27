<!-- ## Citrix ADC  -->
# Redirect an external URL to an internal URL using Citrix ADC

Refer [here](https://docs.citrix.com/en-us/citrix-adc/current-release/appexpert/rewrite/rewrite-action-policy-examples/example-redirecting-external-url.html) for the use-case.

Here is the link to the [demonstration video](https://youtu.be/ioEX3olGxZ4)

## Folder Structure
* `main.tf` describes the actual Citrix ADC config objects to be created. The attributes of these resources are hard coded.
* `provider.tf` is used to specify the username, password and endpoint of the Citrix ADC. Alternatively, you can set the NS_URL, NS_LOGIN and NS_PASSWORD environment variables.

## Usage

### Step-1 Install the Required Plugins
* The terraform needs plugins to be installed in the local folder so, use `terraform init` - It automatically installs the required plugins from the Terraform Registry.

### Step-2 Applying the Configuration 
* Modify the `resources.tf` (if necessary) to suit your own configuration.
* Use `terraform plan` to review the plan
* Use `terraform apply` to apply the configuration.

### Step-3 Updating your configuration
* Modify the set of resources (if necessary)
* Use `terraform plan` and `terraform apply` to review and update the changes respectively.

### Step-4 Destroying your Configuration
* To destroy the configuration use `terraform destroy`.
