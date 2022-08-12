resource "citrixadc_aaacertparams" "tf_aaacertparams" {
  usernamefield              = "Subject:CN"
  groupnamefield             = "Subject:OU"
  defaultauthenticationgroup = 50
}
