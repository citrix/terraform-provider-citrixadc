resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
  sessionlife         = 300
  avgwaitingclient    = 400
  maxaltrespbandwidth = 50
  dosattackthresh     = 100
}
