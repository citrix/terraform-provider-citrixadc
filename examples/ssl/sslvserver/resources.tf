resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "test_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "test_csvserver"
  port        = 443
  servicetype = "SSL"

  sslpolicybinding {
      policyname = "tf_policy"
     priority = 200
  }
  sslpolicybinding {
      policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
  }

}

resource "citrixadc_lbvserver" "test_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "test_lbvserver"
  port        = 443
  servicetype = "SSL"

  sslpolicybinding {
      policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 101
  }
  sslpolicybinding {
      policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
  }

}
