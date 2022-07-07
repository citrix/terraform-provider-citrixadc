## SSL LB with custom HTTP monitor attached to the backend services
This terraform configuration shows an SSL LB vserver with HTTP backends. Each backend has a customized HTTP monitor that informs the vserver of problems with the backend.

## Topology
<pre>
                           +
                           |
                           |
                           |
                           |
                           |
                           |
                 +---------v---------+      +----------+
                 |   AuctionLB       +------+SSL certkey
                 |   Lb vserver      |      +----------+
                 |   10.22.24.22:443 |
                 +------+------------+
                        |
         +--------------v---+------------------------+
         |                  |                        |
         |                  |                        |
         |                  |                        |
+--------v--------+   +-----v-----------+  +---------v------+
|Service          +   |Service          |  |Service         |
|172.23.33.33:8080    |172.23.44.33:8080|  |172.23.44.34:8080
+---------^-------+   +----^------------+  +-------^--------+
          |                |                       |
          |                |                       |
          +--------------^-------------------------+
                 +--------------------+
                 |LB Monitor (HTTP)   |
                 |                    |
                 +--------------------+

</pre>

## Customization
The Netscaler address, password and login can be customized in provider.tf
LB configuration (persistence type, etc) can be customized in resource.tf
The location of the SSL cert key on the Netscaler appliance can be customized in the terraform.tfvars

## Experiment
Change the number of backend services by adding / deleting entries in the backend_services list in terraform.tfvars.
Update the netscaler with `terraform plan` and `terraform apply`
