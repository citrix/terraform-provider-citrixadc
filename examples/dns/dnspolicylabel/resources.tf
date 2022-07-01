resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
  labelname = "blue_label"
  transform = "dns_req"

}