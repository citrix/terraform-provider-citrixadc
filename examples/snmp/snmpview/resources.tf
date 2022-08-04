resource "citrixadc_snmpview" "tf_snmpview" {
  name    = "test_name"
  subtree = "1.2.4.7"
  type    = "excluded"
}