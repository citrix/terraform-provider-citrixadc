resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}

resource "citrixadc_auditsyslogpolicy" "tf_policy" {
    name = "tf_auditsyslogpolicy"
    rule = "ns_true"
    action = citrixadc_auditsyslogaction.tf_syslogaction.name

    globalbinding {
        priority = 120
        feature = "SYSTEM"
        globalbindtype = "SYSTEM_GLOBAL"
    }
}
