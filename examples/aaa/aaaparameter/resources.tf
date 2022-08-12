resource "citrixadc_aaaparameter" "tf_aaaparameter" {
  enablestaticpagecaching    = "NO"
  enableenhancedauthfeedback = "YES"
  defaultauthtype            = "LDAP"
  maxaaausers                = 3
  maxloginattempts           = 5
  failedlogintimeout         = 15
}
