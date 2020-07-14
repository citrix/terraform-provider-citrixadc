resource "citrixadc_csvserver" "foo_csvserver" {

  ipv46       = "10.202.11.11"
  name        = "tst_policy_cs"
  port        = 9090
  servicetype = "SSL"
  comment     = "hello"
  sslprofile = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_lbvserver" "foo_lbvserver" {

  name        = "tst_policy_lb"
  servicetype = "HTTP"
  ipv46       = "192.122.3.3"
  port        = 8000
  comment     = "hello"
}

resource "citrixadc_cspolicy" "foo_cspolicy" {
  csvserver       = citrixadc_csvserver.foo_csvserver.name
  targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
  policyname      = "test_policy"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
  priority        = 10

  # Any change in the following id set will force recreation of the cs policy
  forcenew_id_set = [
    citrixadc_lbvserver.foo_lbvserver.id,
    citrixadc_csvserver.foo_csvserver.id,
  ]
}
