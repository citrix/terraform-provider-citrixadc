resource "citrixadc_csvserver" "test_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "test_csvserver"
  port        = 80
  servicetype = "HTTP"

  lbvserverbinding = citrixadc_lbvserver.test_lbvserver_old.name
}

resource "citrixadc_lbvserver" "test_lbvserver_old" {
  ipv46       = "10.10.10.33"
  name        = "test_lbvserver_old"
  port        = 80
  servicetype = "HTTP"
}
resource "citrixadc_lbvserver" "test_lbvserver_new" {
  ipv46       = "10.10.10.44"
  name        = "test_lbvserver_new"
  port        = 80
  servicetype = "HTTP"
}
