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
