# Configure SSL offloading with end-to-end encryption via Terraform

## What is SSL offloading

A simple SSL offloading setup terminates SSL traffic (HTTPS), decrypts the SSL records, and forwards the clear text (HTTP) traffic to the back-end web servers. Clear text traffic is vulnerable to being spoofed, read, stolen, or compromised by individuals who succeed in gaining access to the back-end network devices or web servers.

You can, therefore, configure SSL offloading with end-to-end security by re-encrypting the clear text data and using secure SSL sessions to communicate with the back-end Web servers.

Configure the back-end SSL transactions so that the appliance uses SSL session multiplexing to reuse existing SSL sessions with the back-end web servers. It helps in avoiding CPU-intensive key exchange (full handshake) operations and also reduces the overall number of SSL sessions on the server. As a result, it accelerates the SSL transaction while maintaining end-to-end security.

## How to use the terraform scripts

1. Download the `ssl-offloading-with-end-to-end-encryption` folder into your local machine.
2. Change the directory to `ssl-offloading-with-end-to-end-encryption` folder.
3. Fill the target Citrix ADC in `provider.tf`
4. Fill the service, lb-vserver, certkey details in `input.auto.tfvars`
5. Run `terraform init` to initialise the terraform environment
6. Run `terraform apply` to apply the ssl offloading configuration to the target Citrix ADC via terraform.

## Further reading

* **Citrix Docs**: [Configure SSL offloading with end-to-end encryption](https://docs.citrix.com/en-us/citrix-adc/current-release/ssl/how-to-articles/end-to-end-encrypt.html)
