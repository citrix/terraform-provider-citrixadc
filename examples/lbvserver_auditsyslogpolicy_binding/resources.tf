resource "citrixadc_lbvserver" "tf_lbvserver1" {
  name        = "tf_lbvserver1"
  servicetype = "HTTP"
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction1" {
    name = "tf_syslogaction1"
    serverip = "10.124.67.92"
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy1" {
    name = "tf_syslogpolicy"
    rule = "true"
    action = citrixadc_auditsyslogaction.tf_syslogaction1.name

}

resource "citrixadc_lbvserver_auditsyslogpolicy_binding" "demo" {
    name = citrixadc_lbvserver.tf_lbvserver1.name
    policyname = citrixadc_auditsyslogpolicy.tf_syslogpolicy1.name
    priority = 100
}
