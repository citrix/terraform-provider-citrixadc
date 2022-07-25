resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}
resource "citrixadc_crvserver_filterpolicy_binding" "crvserver_filterpolicy_binding" {
  name = citrixadc_crvserver.crvserver.name
  policyname = citrixadc_filterpolicy.tf_filterpolicy.name
  priority = 10
}