resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
  name              = "my_rdpclientprofile"
  rdpurloverride    = "ENABLE"
  redirectclipboard = "ENABLE"
  redirectdrives    = "ENABLE"
}
