## Content switching SSL LB with custom HTTP monitor attached to the backend services
This terraform configuration shows an SSL content switch vserver switching to two sets of backends based on the URL pattern.

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
                              |   Production LB   +------+SSL certkey
                              |   cs vserver      |      +----------+
                              |   10.22.24.22:443 |
                              +-------+-----------+
                                      |
                                      |
                        +-------------v-----------+
                        |                         |
  +-------------+       |                         |        +--------------+
  |cs policy    +-------+                         |        | cs policy    |
  |url: /cart/* |       |                         +--------+ url: /catalog/*
  +-------------+       |                         |        +--------------+
                 +------+-------+         +-------+------+
                 + lb vserver   |         | lb vserver   |
                 | cart         |         | catalog      +
                 +--------------+         +--------------+     
                        |                         |
                 +------+-------+         +-------+------+---------------------------+
             +---+ service group|         | service group|                           |
             |   | cart         |         | catalog      +-----+                     |
             |   +----------+---+         +--------------+     |                     |
        +----+--------------|------+              +------------+-----------+         |
        |            |      |      |              |                        |         |
+-------+--+   +-----+-----+|+-----+----+     +---+---------+  +-----------+---+     |
|member    |   |member     |||member    |     | member      |  | member        |     |
|          |   |           |||          |     |             |  |               |     |
+------+---+   +----+------+|+-----+----+     +------+------+  +---------+-----+     |
                            |                        |                   |           |
                            |                        +-----------------+-+           |
         +-------------+    |                       +--------------+                 |
         |lb monitor   |    |                       | lb monitor   +-----------------+
         |cart monitor +----+                       | catalog monitor
         +-------------+                            +--------------+


</pre>


## Experiment
Change the number of backend services by adding / deleting entries in the backend_services list in terraform.tfvars.
Update the netscaler with `terraform plan` and `terraform apply`
