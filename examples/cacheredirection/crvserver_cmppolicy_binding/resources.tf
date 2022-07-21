resource "citrixadc_cmppolicy" "tf_cmppolicy" {
  name      = "tf_cmppolicy1"
  rule      = "HTTP.REQ.HEADER(\"Content-Type\").CONTAINS(\"text\")"
  resaction = "COMPRESS"
}
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_cmppolicy_binding" "crvserver_cmppolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = citrixadc_cmppolicy.tf_cmppolicy.name
  priority   = 10
  bindpoint  = "REQUEST"

}