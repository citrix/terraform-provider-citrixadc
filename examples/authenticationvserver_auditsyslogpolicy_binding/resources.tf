resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name       = "tf_syslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
resource "citrixadc_auditsyslogpolicy" "tf_policy" {
  name   = "tf_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}
resource "citrixadc_authenticationvserver_auditsyslogpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_auditsyslogpolicy.tf_policy.name
  priority  = 80
  bindpoint = "REQUEST"
}