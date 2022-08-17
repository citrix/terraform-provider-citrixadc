resource "citrixadc_aaaglobal_aaapreauthenticationpolicy_binding" "tf_aaaglobal_aaapreauthenticationpolicy_binding" {
  policy    = "my_preauthentication_policy"
  priority  = 50
  }