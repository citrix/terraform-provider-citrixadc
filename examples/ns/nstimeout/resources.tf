resource "citrixadc_nstimeout" "tf_nstimeout" {
  zombie     = 70
  client     = 2300
  server     = 2400
  httpclient = 2500
  reducedrsttimeout = 15
}
