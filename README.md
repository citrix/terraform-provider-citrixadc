[![CircleCI](https://circleci.com/gh/citrix/terraform-provider-citrixadc/tree/master.svg?style=shield)](https://circleci.com/gh/citrix/terraform-provider-citrixadc/tree/master)

## Terraform provider for Citrix ADC
Citrix has developed a custom Terraform provider for automating [Citrix ADC](https://www.citrix.com/products/netscaler-adc/) deployments and configurations. Using [Terraform](https://www.terraform.io), you can [custom configure your ADCs](https://www.youtube.com/watch?v=IJIIWm5rzpQ&ab_channel=Citrix) for different use-cases such as Load Balancing, SSL, Content Switching, GSLB, WAF etc. 

For users new to terraform provider for Citrix ADC, check out the [_**installation steps**_](#installation) and [_**getting started with configuring adc**_](#get-started-on-configuring-adc-through-terraform).

For deploying Citrix ADC in Public Cloud - AWS and Azure, check out cloud scripts in github repo [terraform-cloud-scripts](https://github.com/citrix/terraform-cloud-scripts).

Learn more about Citrix ADC Automation [here](https://docs.citrix.com/en-us/citrix-adc/current-release/deploying-vpx/citrix-adc-automation.html) 

**Important note: The provider will not commit the config changes to Citrix ADC's persistent store.**

## Table of contents

* [Why Terraform for Citrix ADC ?](#why-terraform-for-citrix-adc?)
* [Navigating Repository](#navigating-the-repository)
* [Installating Terraform and Citrix ADC Provider](#installation)
* [Get Started on using terraform to configure Citrix ADC](#get-started-on-configuring-adc-through-terraform)
* Usage Guidelines
  - [Understanding Provider Configuration](#understanding-provider-configuration)
  - [Understanding Resource Configuration](#resource-configuration)
  - [General guidelines on ADC configurations](#general-guidelines-on-configuring-adc)
  - [Commiting changes to Citrix ADC's persistent store](#commiting-changes-to-citrix-adc's-persistent-store)
  - [List of ADC use-cases supported through Terraform](#adc-use-case-supported-through-terraform)
  - [Using `remote-exec` for one-time tasks](#using-`remote-exec`-for-one-time-tasks)
  - [Building your own provider](#building)


## Why Terraform for Citrix ADC ?

[Terraform](https://www.terraform.io/) is an open-source infrastructure as code software tool that provides a consistent CLI workflow to manage hundreds of cloud services.Terraform codifies cloud APIs into declarative configuration files.
Terraform can be used to **_deploy_** and **_configure_** ADC. Configuring Citrix ADC through Terraform provides multiple benefits.
1. Infrastucture as Code approach to ADC -You can store the ADC configs in scm tools like GitHub and version and track it like just other code repositories you have.
2. Declarative Approach to ADC automation - Users just need to defined the target state of ADC. ADC terraform resources will make the appropriate API calls to achieve the target state.
3. ADC resources files in Terraform are human friendly and easy to understand.
4. Abstract away the complexity associated with Citrix ADC internals architecture.
5. Detect the configuration drifts on ADC through Terraform easily.

## Requirement [Do we need this ????????]

* [hashicorp/terraform](https://github.com/hashicorp/terraform)

## Navigating the repository

1. _citrixadc folder_ - Contains all the ADC resources library that we support through Terraform. These resource libraries will internally call NITRO APIS to configure target ADC.
2. _examples folder_ - Contain the examples for users to use various ADC resources e.g [simple_lb](https://github.com/citrix/terraform-provider-citrixadc/blob/master/examples/simple_lb/) folder contains the resources.tf that illustrates how citrixadc_lbvserver resource can be used to create a Load Balancing vserver on target ADC. Similarly , different folders contains examples on defining different resources. Users are expected to review these examples and define their desired ADC configurations.
3. _docs folder_ - https://github.com/citrix/terraform-provider-citrixadc/tree/master/docs/resources  - contains the documentation of all resources confgirations supported through Terraform. Refer this to understand the different arguments, values that a particular resource takes.


## Installation

### **Step 1. Installing Terraform CLI:**
First step is to install Terraform CLI. Refer the https://learn.hashicorp.com/tutorials/terraform/install-cli for installing Terraform CLI. 

### **Step 2. Installing Citrix ADC Provider:**
Terraform provider for Citrix ADC is not available through terrform.registry.io as of now. Hence users have to install the provider manually.

#### **Follow below steps to install citrix adc provider for Terraform CLI version < 13.0**
1. Download the citrix adc terraform binary in your local machine where you have terraform installed from the [Releases section of the github repo](https://github.com/citrix/terraform-provider-citrixadc/releases).Untar the files and you can find the binary file terraform-provider-ctxadc.

2. Edit .terraformrc for the base directory of plugins:
```
plugin_cache_dir = "/home/user/.terraform.d/plugins"
```
3. Copy terrafom-provider-citrixadc binary in appropriate location - `$plugin_cache_dir/<platform>/terraform-provider-citrixadc`.
e.g. `/home/user/.terraform.d/plugins/linux_amd64/terraform-provider-citrixadc`

#### **Follow below steps to install citrix adc provider for Terraform CLI version >13.0**
1. Download the citrix adc terraform binary in your local machine where you have terraform installed from the [Releases section of the github repo](https://github.com/citrix/terraform-provider-citrixadc/releases).Untar the files and you can find the binary file terraform-provider-ctxadc.

2. Create a following directory in your local machine and save the citrix adc terraform binary. e.g. in Ubuntu machine. Note that the directory structure has to be same as below, you can edit the version -0.12.43 to the citrix adc version you downloaded.
```
mkdir -p /home/user/.terraform.d/plugins/registry.terraform.io/citrix/citrixadc/0.12.43/linux_amd64/
```
3. Copy the terraform-provider-citrixadc to the above created folder as shown below
```
cp terraform-provider-citrixadc /home/user/.terraform.d/plugins/registry.terraform.io/citrix/citrixadc/0.12.43/linux_amd64/
```

## Get Started on Configuring ADC through Terraform
_In order to familiarize with citrix adc configuration through terraform, lets get started with basic configuration of setting up server in ADC through Terraform._

Before we configure, clone the github repository in your local machine as follows:
```
git clone https://github.com/citrix/terraform-provider-citrixadc/
```
**Step-1** : Now navigate to examples folder as below. Here you can find many ready to use examples for you to get started:
```
cd terraform-provider-citrixadc/examples/
```
Lets configure a simple server in citrix ADC.
```
cd terraform-provider-citrixadc/examples/simple_server/
```
**Step-2** : Provider.tf contains the details of the target Citrix ADC.Edit the simple_server/provider.tf as follows and add details of your target adc.
For **terraform version > 13.0** edit the provider.tf as follows
```
terraform {
    required_providers {
        citrixadc = {
            source = "citrix/citrixadc"
        }
    }
}
provider "citrixadc" {
  endpoint = "http://10.1.1.3:80"
  username = "UsernameOfYourADC"
  password = "PasswordOfYourADC"
 }
```
For **terraform version < 13.0**, edit the provider.tf as follows
```
provider "citrixadc" {
  endpoint = "http://10.1.1.3:80"
  username = "UsernameOfYourADC"
  password = "PasswordOfYourADC"
 }
 ```
**Step-3** : Resources.tf contains the desired state of the resources that you want to manage through terraform.Here we want to create simple server. Edit the simple_server/resources.tf with your configuration values - name,ipaddress as below. 
```
resource "citrixadc_server" "test_server" {
  name      = "test_server"
  ipaddress = "192.168.2.2"
}
```
**Step-4** : Once the provider.tf and resources.tf is edited and saved with the desired values in the simple_server folder, you are good to run terraform and configure ADC.Initialize the terraform by running terraform-init inside the simple_server folder as follow:
```
terraform-provider-citrixadc/examples/simple_server$ terraform init
```
You should see following output if terraform was able to successfully find citrix adc provider and initialize it -
![image](https://user-images.githubusercontent.com/68320753/111422447-ba528d00-8714-11eb-91a6-02a1418b73eb.png)

**Step-5** : Now run the terraform-plan command. This will fetch the true state of your target ADC and will show you the changes/additions it need to make to achieve the desired configuration given in resources.tf. As we see below, terraform plans to add a new resource :
```
terraform-provider-citrixadc/examples/simple_server$ terraform plan
```

![image](https://user-images.githubusercontent.com/68320753/111422516-d5250180-8714-11eb-89e2-bc3d3432c9c7.png)

**Step-6** : If the above plan looks good, then go ahead and run terraform-apply to apply the configurations. Type yes, when prompted.**
```
terraform-provider-citrixadc/examples/simple_server$ terraform apply
```

![image](https://user-images.githubusercontent.com/68320753/111423045-b410e080-8715-11eb-9845-741b6398efbb.png)
![image](https://user-images.githubusercontent.com/68320753/111423077-bf640c00-8715-11eb-835a-fe36b90576db.png)
As you see above, terraform successfully created server with name test_server3 and given ipaddress on your target ADC. You can validate it by going to ADC GUI, and navigating to Traffic Management -> Load Balancing -> Servers. 

_Similary repeat steps 1-6 for different resource configurations on Citrix ADC. Also refer to [general guidelines on configuring ADC](#general-guidelines-on-configuring-adc)_


## Usage Guidelines

### Understanding Provider Configuration
Provider.tf contains the information on target ADC where you want to apply configuration.
```
provider "citrixadc" {
    username = "${var.ns_user}"
    password = "${var.ns_password}"
    endpoint = "http://10.71.136.250/"
}
```

We can use a `https` URL and accept the untrusted authority certificate on the Citrix ADC by specifying `insecure_skip_verify = true`

To use `https` without the need to set `insecure_skip_verify = true` follow this [guide](https://support.citrix.com/article/CTX122521) on
how to replace the default TLS certificate with one from a trusted Certifcate Authority.

Use of `https` is preferred. Using `http` will result in all provider configuration variables as well as resource variables
to be transmitted in cleartext. Anyone observing the HTTP data stream will be able to parse sensitive values such as the provider password.

Avoid storing provider credentials in the local state by using a backend that supports encryption.
The hasicorp [vault provider](https://registry.terraform.io/providers/hashicorp/vault/latest/docs) is also recommended for
storing sensitive data.

##### Argument Reference

The following arguments are supported.

* `username` - This is the user name to access to Citrix ADC. Defaults to `nsroot` unless environment variable `NS_LOGIN` has been set
* `password` - This is the password to access to Citrix ADC. Defaults to `nsroot` unless environment variable `NS_PASSWORD` has been set
* `endpoint` - (Required) Nitro API endpoint in the form `http://<NS_IP>/` or `http://<NS_IP>:<PORT>/`. Can be specified in environment variable `NS_URL`
* `insecure_skip_verify` - (Optional, true/false) Whether to accept the untrusted certificate on the Citrix ADC when the Citrix ADC endpoint is `https`
* `proxied_ns` - (Optional, NSIP) The target Citrix ADC NSIP for MAS proxied calls. When this option is defined, `username`, `password` and `endpoint` must refer to the MAS proxy.

The username, password and endpoint can be provided in environment variables `NS_LOGIN`, `NS_PASSWORD` and `NS_URL`. 

### Resource Configuration
Resources.tf contains the desired state of the resources that you want on target ADC. E.g. For creating a Load Balancing vserver in ADC following resource.tf contains the desired configs of lbvserver 

**`citrixadc_lbvserver`**
```
resource "citrixadc_lbvserver" "foo" {
  name = "sample_lb"
  ipv46 = "10.71.136.150"
  port = 443
  servicetype = "SSL"
  lbmethod = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  sslcertkey = "${citrixadc_sslcertkey.foo.certkey}"
  sslprofile = "ns_default_ssl_profile_secure_frontend"
}
```
In order to understand the arguments, possible values, and other arguments available for a given resource, refer the NITRO API documentation <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/load-balancing/lbvserver/lbvserver/>  and the Terraform documentation such as https://github.com/citrix/terraform-provider-citrixadc/blob/master/docs/resources/lbvserver.md .

??????**Note that the attribute `state` is not synced with the remote object.
If the state of the lb vserver is out of sync with the terraform configuration you will need to manually taint the resource and apply the configuration again.
**

## General guidelines on configuring ADC
The subfolders in the [example folder](https://github.com/citrix/terraform-provider-citrixadc/tree/master/examples) contains examples of different ADC configurations through terraform

### Structure
* `resources.tf` describes the actual NetScaler config objects to be created. The attributes of these resources are either hard coded or looked up from input variables in `terraform.tfvars`
* `variables.tf` describes the input variables to the terraform config. These can have defaults
* `provider.tf` is used to specify the username, password and endpoint of the NetScaler. Alternatively, you can set the NS_URL, NS_LOGIN and NS_PASSWORD environment variables.
* `terraform.tfvars` has the variable inputs specified in `variables.tf`

### Using
Modify the `terraform.tfvars` and `provider.tf` to suit your own NetScaler deployment. Use `terraform plan` and `terraform apply` to configure the NetScaler.

### Updating your configuration
Modify the set of backend services and use `terraform plan` and `terraform apply` to verify the changes

### Commiting changes to Citrix ADC's persistent store
The provider will not commit the config changes to Citrix ADC's persistent store. To do this, run the shell script `ns_commit.sh`:

```
export NS_URL=http://<host>:<port>/
export NS_LOGIN=nsroot
export NS_PASSWORD=nsroot
./ns_commit.sh
```

To ensure that the config is saved on every run, we can use something like `terraform apply && ns_commit.sh`

## ADC Use-Case supported through Terraform

ADC Use-Case -  Configuration examples (resource.tf )

1. Load Balancing -  
2. Content Switching
3. Responder/Rewrite Policies
4. SSL
5. Global Load Server Balancing (GSLB)
6. Web Application Firewall (WAF)
7. Core ADC features
8. Pool Licensing         


## Using `remote-exec` for one-time tasks
Terraform is useful for maintaining desired state for a set of resources. It is less useful for tasks such as network configuration which don't change. Network configuration is like using a provisioner inside Terraform. The directory `examples/remote-exec` show examples of how Terraform can use ssh to accomplish these one-time tasks.

## Building
### Assumption
* You have (some) experience with Terraform, the different provisioners and providers that come out of the box,
its configuration files, tfstate files, etc.
* You are comfortable with the Go language and its code organization.

1. Install `terraform` from <https://www.terraform.io/downloads.html>
2. Install `dep` (<https://github.com/golang/dep>)
3. Check out this code: `git clone https://<>`
4. Build this code using `make build`
