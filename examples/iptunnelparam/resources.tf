resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
  dropfrag             = "NO"
  dropfragcputhreshold = 1
  srciproundrobin      = "NO"
  enablestrictrx       = "NO"
  enablestricttx       = "NO"
  useclientsourceip    = "NO"
}