resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
  rtspalgprofilename = "my_lsn_rtspalgprofile"
  rtspportrange      = 4200
  rtspidletimeout    = 150
}
