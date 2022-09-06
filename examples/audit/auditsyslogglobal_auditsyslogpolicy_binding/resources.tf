resource "citrixadc_auditsyslogglobal_auditsyslogpolicy_binding" "tf_auditsyslogglobal_auditsyslogpolicy_binding" {
  policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
  priority   = 100
  globalbindtype = "SYSTEM_GLOBAL"
}

resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "true"
    action = citrixadc_auditsyslogaction.tf_syslogaction.name
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}