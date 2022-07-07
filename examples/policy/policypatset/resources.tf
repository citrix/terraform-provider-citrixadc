resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind1" {
    name = citrixadc_policypatset.tf_patset.name
    string = "Pattern"
}
