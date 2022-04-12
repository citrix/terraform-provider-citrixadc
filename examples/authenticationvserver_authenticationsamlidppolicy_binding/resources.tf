resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert    = "/var/tmp/certificate1.crt"
  key     = "/var/tmp/key1.pem"
}
resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
  name                        = "tf_samlidpprofile"
  samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
  assertionconsumerserviceurl = "http://www.example.com"
  sendpassword                = "OFF"
  samlissuername              = "new_user"
  rejectunsignedrequests      = "ON"
  signaturealg                = "RSA-SHA1"
  digestmethod                = "SHA1"
  nameidformat                = "Unspecified"
}
resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
  name    = "tf_samlidppolicy"
  rule    = "false"
  action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
  comment = "aSimpleTesting"
}
resource "citrixadc_authenticationvserver_authenticationsamlidppolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationsamlidppolicy.tf_samlidppolicy.name
  bindpoint = "REQUEST"
  priority  = 88
  secondary = "false"
}