resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  ipv46       = "10.102.80.55"
  port        = 8090
  cachetype   = "REVERSE"
}
resource "citrixadc_policymap" "tf_policymap" {
  mappolicyname = "ia_mappol123"
  sd            = "amazon.com"
  td            = "apple.com"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_lbvserver"
  servicetype = "HTTP"
  ipv46       = "192.122.3.31"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_service" "tf_service" {
  lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
  name        = "tf_service"
  port        = 8081
  ip          = "10.33.4.5"
  servicetype = "HTTP"
  cachetype   = "TRANSPARENT"
}
resource "citrixadc_crvserver_policymap_binding" "crvserver_policymap_binding" {
  name          = citrixadc_crvserver.crvserver.name
  policyname    = citrixadc_policymap.tf_policymap.mappolicyname
  targetvserver = citrixadc_lbvserver.foo_lbvserver.name
  depends_on = [
    citrixadc_service.tf_service
  ]
}
