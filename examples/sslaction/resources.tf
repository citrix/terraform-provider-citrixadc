resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}
