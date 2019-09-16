resource "citrixadc_lbvserver" "test_lb" {
  name = "testLB"
  ipv46 = "1.2.3.4"
  port = "80"
  servicetype = "HTTP"
}

