resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}

resource "citrixadc_lbvserver_filterpolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_filterpolicy.tf_filterpolicy.name
    priority = 100
}
