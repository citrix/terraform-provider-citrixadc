resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
  name           = "my_formsso_action"
  actionurl      = "/logon.php"
  userfield      = "loginID"
  passwdfield    = "passwd"
  ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
}
