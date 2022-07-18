resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
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
resource "citrixadc_vpnvserver_authenticationsamlidppolicy_binding" "tf_binding" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_authenticationsamlidppolicy.tf_samlidppolicy.name
  priority  = 9
  bindpoint = "REQUEST" //doesnot unbind for other values
}