<!-- ## Citrix ADC  -->
# Using `local-exec` to Configure Citrix ADC as the last-resort

## When to Use

All the supported Citrix ADC terraform resource can be found in the [Terraform Registry](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs)
For any not yet supported resource you can use the below steps.

## Files Structure

The folder contains two files
* main.tf
* do_nitro_request.sh

### main.tf

It contains `null_resource` and inside that we have a local-exec `provisioner` to know more about local-exec provisioner [here](https://www.terraform.io/language/resources/provisioners/local-exec). The local-exec provisioner will run the script `do_nitro_request.sh` which contains the curl commands to be executed.


### do_nitro_request.sh

In this script we give curl command to be executed and response is received. It contains the URL, Request Headers, payload that is accepted by the resource. To know more about [curl](https://curl.se/docs/) 

For Example :   

In the `do_nitro_request.sh` -  

* Updates the parameter of the `systemparameter`. 
* Adds the lbvserver `my_lbvserver`
* Adds service `my_service`
* Binds the resources `lbvserver` to `service`(`lbvserver_service_binding`)
* Get the details of binding `lbvserver_service_binding`
    
    
Similarly you can use this script to make API calls for any resources that are not yet supported in our terraform registry.

## Usage

### Step-1 Install the Required Plugins
* The terraform needs plugins to be installed in local folder so, use `terraform init` - It automatically installs the required plugins from the Terraform Registry.

### Step-2 Applying the Configuration 
* Modify the `main.tf` and `do_nitro_request.sh` to suit your resource configuration.
* Use `terraform plan` to review the plan
* Use `terraform apply` to apply the configuration.

### Step-3 Updating your configuration
* Refer `Troubleshoot` below

### Step-4 Destroying your Configuration
* To destroy the configuration use `terraform destroy`.

## Troubleshoot

* Since the `local-exec` provisoner is used in `null_resource`. The `null_resource` executes only once so if there is any error in the curl commands than you need to run `terraform destroy` and modify the changes and apply the new changes by `terraform apply`

## Further Reading

* **Curl**: https://curl.se/docs/
* **Null Resource**: https://www.terraform.io/language/resources/provisioners/null_resource and https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource
* **local-exec Provisioner**: https://www.terraform.io/language/resources/provisioners/local-exec
