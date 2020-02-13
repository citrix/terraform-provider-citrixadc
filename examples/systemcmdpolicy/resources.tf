resource "citrixadc_systemcmdpolicy" "tf_policy" {
    policyname = "test_policy"
    action = "ALLOW"
    cmdspec = "show.*"
}
