resource "citrixadc_dnspolicy" "dnspolicy" {
  name = "policy_A"
  rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
  drop = "YES"
}
resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
  labelname = "blue_label"
  transform = "dns_req"

}
resource "citrixadc_dnspolicylabel_dnspolicy_binding" "dnspolicylabel_dnspolicy_binding" {
  labelname  = citrixadc_dnspolicylabel.dnspolicylabel.labelname
  policyname = citrixadc_dnspolicy.dnspolicy.name
  priority   = 10

}
