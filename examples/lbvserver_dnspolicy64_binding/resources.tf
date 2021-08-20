# Since the dnspolicy64 resource is not yet available on Terraform,
# the tf_dnspolicy64 policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add dns policy64 tf_dnspolicy64 -rule true -action default_DNS64_action

resource "citrixadc_lbvserver_dnspolicy64_binding" "tf_lbvserver_dnspolicy64_binding" {
      name = citrixadc_lbvserver.tf_lbvserver.name
      policyname = "tf_dnspolicy64"
      priority = 1
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "DNS_TCP"
}