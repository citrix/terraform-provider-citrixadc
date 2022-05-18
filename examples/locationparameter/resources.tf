resource "citrixadc_locationparameter" "tf_locationpara" {
  context            = "geographic"
  q1label            = "asia"
  matchwildcardtoany = "YES"
}