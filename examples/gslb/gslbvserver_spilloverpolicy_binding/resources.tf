# Since the spilloverpolicy resource is not yet available on Terraform,
# the tf_spilloverpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add spillover policy tf_spilloverpolicy -rule TRUE -action SPILLOVER

resource "citrixadc_gslbvserver_spilloverpolicy_binding" "tf_gslbvserver_spilloverpolicy_binding" {
  name       = citrixadc_gslbvserver.tf_gslbvserver.name
  policyname = "tf_spilloverpolicy"
  priority   = 100

}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "GSLB-East-Coast-Vserver"
  servicetype   = "HTTP"
  domain {
    domainname = "www.fooco.co"
    ttl        = "60"
  }
  domain {
    domainname = "www.barco.com"
    ttl        = "65"
  }
}

