resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
  dateformat = "DDMMYYYY"
  loglevel   = ["EMERGENCY"]
  tcp        = "ALL"
}
