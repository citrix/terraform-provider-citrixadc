resource "citrixadc_cachepolicy" "policy1" {
    policyname = "policy1"
    rule = "true"
    action = "CACHE"
}