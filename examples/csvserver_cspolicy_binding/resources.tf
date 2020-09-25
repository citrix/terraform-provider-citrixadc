resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname      = "tf_cspolicy"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cspolicy.tf_cspolicy.policyname
    priority = 100
    targetlbvserver = citrixadc_lbvserver.tf_lbvserver.name
}

/*
resource "citrixadc_cspolicy" "tf_cspolicy_extra" {
  policyname      = "tf_cspolicy_extra"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.86.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind_extra" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cspolicy.tf_cspolicy_extra.policyname
    priority = 110
    targetlbvserver = citrixadc_lbvserver.tf_lbvserver.name
}
*/
