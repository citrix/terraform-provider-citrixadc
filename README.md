# Terraform `Citrix ADC` Provider

Citrix has developed a custom Terraform provider for automating [Citrix ADC](https://www.citrix.com/products/netscaler-adc/) deployments and configurations. Using [Terraform](https://www.terraform.io), you can [custom configure your ADCs](https://www.youtube.com/watch?v=IJIIWm5rzpQ&ab_channel=Citrix) for different use-cases such as Load Balancing, SSL, Content Switching, GSLB, WAF etc.

Learn more about Citrix ADC Automation [here](https://docs.citrix.com/en-us/citrix-adc/current-release/deploying-vpx/citrix-adc-automation.html)

> :round_pushpin:For deploying Citrix ADC in Public Cloud - AWS and Azure, check out cloud scripts in github repo [terraform-cloud-scripts](https://github.com/citrix/terraform-cloud-scripts).

> :envelope: For any immediate issues or help , reach out to us at appmodernization@citrix.com !

## Terrraform Provider Documentation

1. [Why Terraform for Citrix ADC ?](#why-terraform-for-citrix-adc-)
2. [Navigating Repository](#navigating-the-repository)
3. [Understanding Provider Configuration](#understanding-provider-configuration)
4. [Understanding Resource Configuration](#resource-configuration)
5. [General guidelines on ADC configurations](#general-guidelines-on-configuring-adc)
6. [Commiting changes to Citrix ADC's persistent store](#commiting-changes-to-citrix-adcs-persistent-store)
7. [List of ADC use-cases supported through Terraform](#adc-use-case-supported-through-terraform)
8. [Using `remote-exec` for one-time tasks](#using-remote-exec-for-one-time-tasks)


## Beginners Guide to Automating ADC with Terraform

1. [Hands-on lab with ADC automation with Terraform](#hands-on-lab-with-adc-automation-with-terraform)
2. [Install Terraform in your own setup](#install-terraform-in-your-own-setup)
3. [Understanding the ADC terraform provider repository](#understanding-the-adc-terraform-provider-repository))
4. [Get your first terraform config into ADC](#get-your-first-terraform-config-into-adc)
5. [How to write terraform resources file for ADC](#how-to-write-terraform-resources-file-for-adc)
6. [Set up SSL-Offloading use-case in ADC](#set-up-ssl-offloading-use-case-in-adc)
7. [Committing changes to Citrix ADC's persistent store](#commiting-changes-to-citrix-adcs-persistent-store)
8. [Managing ADC configs drifts in terraform](#managing-adc-configs-drifts-in-terraform)
9. [Importing ADC configs into Terraform resources file](#importing-adc-configs-into-terraform-resources-file)

## Advanced guide on Automating ADC with Terraform

1. [Deploy ADC in AWS using Terraform](#deploy-adc-in-aws-using-terraform)
2. [Leveraging Terraform workspaces to manage multiple ADCs](#leveraging-terraform-workspaces-to-manage-multiple-adcs)
3. [Dynamically updates Services using Consul-Terraform-Sync](#dynamically-updates-services-using-consul-terraform-sync)
4. [Blue-Green Deployment with Citrix ADC and Azure Pipelines](#blue-green-deployment-with-citrix-adc-and-azure-pipelines)

### Why Terraform for Citrix ADC ?

[Terraform](https://www.terraform.io/) is an open-source infrastructure as code software tool that provides a consistent CLI workflow to manage hundreds of cloud services.Terraform codifies cloud APIs into declarative configuration files.
Terraform can be used to **_deploy_** and **_configure_** ADC. Configuring Citrix ADC through Terraform provides multiple benefits.

1. Infrastucture as Code approach to ADC -You can store the ADC configs in scm tools like GitHub and version and track it like just other code repositories you have.
2. Declarative Approach to ADC automation - Users just need to defined the target state of ADC. ADC terraform resources will make the appropriate API calls to achieve the target state.
3. ADC resources files in Terraform are human friendly and easy to understand.
4. Abstract away the complexity associated with Citrix ADC internals architecture.
5. Detect the configuration drifts on ADC through Terraform easily.

### Navigating the repository

1. `citrixadc` folder - Contains all the ADC resources library that we support through Terraform. These resource libraries will internally call NITRO APIS to configure target ADC.
2. `examples` folder - Contain the examples for users to use various ADC resources e.g [simple_lb](https://github.com/citrix/terraform-provider-citrixadc/blob/master/examples/simple_lb/) folder contains the resources.tf that illustrates how citrixadc_lbvserver resource can be used to create a Load Balancing vserver on target ADC. Similarly , different folders contains examples on defining different resources. Users are expected to review these examples and define their desired ADC configurations.
3. `docs` folder` - https://github.com/citrix/terraform-provider-citrixadc/tree/master/docs/resources  - contains the documentation of all resources confgirations supported through Terraform. Refer this to understand the different arguments, values that a particular resource takes.


### Understanding Provider Configuration

`provider.tf` contains the information on target ADC where you want to apply configuration.

```hcl
provider "citrixadc" {
    username = "${var.ns_user}"  # You can optionally use `NS_LOGIN` environment variables.
    password = "${var.ns_password}"  # You can optionally use `NS_PASSWORD` environment variables.
    endpoint = "http://10.71.136.250/"  # You can optionally use `NS_URL` environment variables.
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

You can also use environment variables as stated in the comments above.

The following arguments are supported.
- `username` - This is the user name to access to Citrix ADC. Defaults to `nsroot` unless environment variable `NS_LOGIN` has been set
- `password` - This is the password to access to Citrix ADC. Defaults to `nsroot` unless environment variable `NS_PASSWORD` has been set
- `endpoint` - (Required) Nitro API endpoint in the form `http://<NS_IP>/` or `http://<NS_IP>:<PORT>/`. Can be specified in environment variable `NS_URL`
* `insecure_skip_verify` - (Optional, true/false) Whether to accept the untrusted certificate on the Citrix ADC when the Citrix ADC endpoint is `https`
- `proxied_ns` - (Optional, NSIP) The target Citrix ADC NSIP for MAS proxied calls. When this option is defined, `username`, `password` and `endpoint` must refer to the MAS proxy.

The username, password and endpoint can be provided in environment variables `NS_LOGIN`, `NS_PASSWORD` and `NS_URL`.

### Resource Configuration

Resources.tf contains the desired state of the resources that you want on target ADC. E.g. For creating a Load Balancing vserver in ADC following resource.tf contains the desired configs of lbvserver

**`citrixadc_lbvserver`**

```hcl
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

> the attribute `state` is not synced with the remote object. If the state of the lb vserver is out of sync with the terraform configuration you will need to manually taint the resource and apply the configuration again.

### General guidelines on configuring ADC

The subfolders in the example folder contains examples of different ADC configurations through terraform. Refer to [simple_lb](https://github.com/citrix/terraform-provider-citrixadc/tree/master/examples/simple_lb) example to understand below structure and usage.

##### Structure

- `resources.tf` describes the actual NetScaler config objects to be created. The attributes of these resources are either hard coded or looked up from input variables in `terraform.tfvars`
- `variables.tf` describes the input variables to the terraform config. These can have defaults
- `provider.tf` is used to specify the username, password and endpoint of the NetScaler. Alternatively, you can set the NS_URL, NS_LOGIN and NS_PASSWORD environment variables.
- `terraform.tfvars` has the variable inputs specified in `variables.tf`

#### Using

- Modify the `terraform.tfvars` and `provider.tf` to suit your own NetScaler deployment.
- Use `terraform plan` and `terraform apply` to configure the NetScaler.

#### Updating your configuration

Modify the set of backend services and use `terraform plan` and `terraform apply` to verify the changes

### Commiting changes to Citrix ADC's persistent store

The provider will not commit the config changes to Citrix ADC's persistent store. To do this, run the shell script `ns_commit.sh`:

```bash
export NS_URL=http://<host>:<port>/
export NS_LOGIN=nsroot
export NS_PASSWORD=nsroot
./ns_commit.sh
```

To ensure that the config is saved on every run, we can use something like `terraform apply && ns_commit.sh`

### ADC Use-Case supported through Terraform

List of Use-Cases supported in ADC can be found in here https://registry.terraform.io/providers/citrix/citrixadc/latest/docs .


### Using `remote-exec` for one-time tasks

Terraform is useful for maintaining desired state for a set of resources. It is less useful for tasks such as network configuration which don't change. Network configuration is like using a provisioner inside Terraform. The directory `examples/remote-exec` show examples of how Terraform can use ssh to accomplish these one-time tasks.

---------

## Beginners Guide to ADC Automation with Terraform

### Hands-on lab with ADC automation with Terraform

Try out our [Hands on Lab](https://forum.developer.cloud.com/s/netscaler-labs?labId=000001068) to experience whats it like using Terraform for ADC.

### Install Terraform in your own setup

First step to using Terraform for ADC is to install Terraform CLI. Refer the [Hashicorp documentation](https://learn.hashicorp.com/tutorials/terraform/install-cli) for installing Terraform CLI for your own environment.

### Understanding the ADC terraform provider repository

Refer the [Navigating our repository](#navigating-the-repository) section

### Get your first terraform config into ADC

Follow the article on [Getting started with Terraform on NetScaler](https://forum.developer.cloud.com/s/article/Get-Started-with-NetScaler-Automation-using-Terraform) to get your first configuration

### How to write terraform resources file for ADC

To write the Terraform resources for Citrix ADC, refer the following links
  [NITRO API documentation](https://developer-docs.citrix.com/projects/citrix-adc-nitro-api-reference/en/latest/configuration/configuration/) or [terraform registry documentation](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs).

### Set up SSL-Offloading use-case in ADC

Here is the [Terraform Template](https://forum.developer.cloud.com/s/article/SSL-Offloading-using-Terraform)that you follow to configure SSL Offloading.

### Committing changes to Citrix ADC's persistent store

Refer the [commiting changes](#commiting-changes-to-citrix-adcs-persistent-store) section

### Managing ADC configs drifts in terraform

You want to see the current state of ADC entities in Terraform 
- Use **terraform refresh** to update your local terraform state file to match with existing ADC state
- Use **terraform show** to show the current state for your entire configuration
- Use **terraform state list** to show the resources that are being tracked/managed via Terraform
- To inspect a particular entity use **terraform state show <entity_name>** e.g. 
**terraform state show citrixadc_servicegroup.tf_servicegroup**

If you want to override the ADC configuration with the configs you have in Terraform resource file then 
- You can run **terraform plan** to see the drifts/diff between the two state 
- Run **terraform apply** to push the desired configs( in your Terraform resource file) to your ADC.

Update your terraform state file to reflect the current/true state of ADC
- Use **terraform refresh** to update your local terraform state file to match with existing ADC state

### Importing ADC configs into Terraform resources file
Learn how to import existing NetScaler configurations into Terraform resources [here](https://forum.developer.cloud.com/s/article/Importing-existing-NetScaler-configuration-into-Terraform)

-----------
## Advanced guide on Automating ADC with Terraform

### Deploy ADC in AWS using Terraform
Refer our [terraform cloud scripts for AWS](https://github.com/citrix/terraform-cloud-scripts) and [demo video](https://www.youtube.com/watch?v=LgGS0-Q5ODE&list=PLrUklKi1o_Zny9cgvjJ7xrBtcdOY_Kc6N&index=15&ab_channel=Citrix)


### Leveraging Terraform workspaces to manage multiple ADCs

TBD

### Dynamically updates Services using Consul-Terraform-Sync

Our [Consul-Terraform-Sync integration](https://github.com/citrix/terraform-citrixadc-servicegroup-consul-sync-nia) allows users to automatically create, update and delete Service groups in Citrix ADC that are synced with the Consul Terraform Sync framework.

### Blue-Green Deployment with Citrix ADC and Azure Pipelines

Integrate ADC and configure it faciliate Blue-Green deployment using Terraform.Check out the [solution](https://github.com/citrix/citrix-k8s-ingress-controller/tree/master/cd/blue-green)
