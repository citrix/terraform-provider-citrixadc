resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "New_vserver"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationauthnprofile" "tf_authenticationauthnprofile" {
  name                = "tf_name"
  authnvsname         = citrixadc_authenticationvserver.tf_authenticationvserver.name
  authenticationhost  = "hostname"
  authenticationlevel = "20"
}