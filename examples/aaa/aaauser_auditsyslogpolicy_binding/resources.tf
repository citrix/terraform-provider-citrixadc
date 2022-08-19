resource "citrixadc_aaauser_auditsyslogpolicy_binding" "tf_aaauser_auditsyslogpolicy_binding" {
  username = "user1"
  policy    = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
  priority  = 100
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
resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
  name   = "tf_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}