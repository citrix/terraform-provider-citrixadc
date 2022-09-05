resource "citrixadc_auditsyslogparams" "tf_auditsyslogparams" {
  dateformat = "DDMMYYYY"
  loglevel   = ["EMERGENCY"]
  tcp        = "ALL"
}