resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
  name = "my_policy"
  rule = "REQ.VLANID == 5"
  reqaction = "my_action"
}
