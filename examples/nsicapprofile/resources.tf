resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
  name             = "tf_nsicapprofile"
  uri              = "/example"
  mode             = "REQMOD"
  reqtimeout       = 4
  reqtimeoutaction = "RESET"
  preview          = "ENABLED"
  previewlength    = 4096
}