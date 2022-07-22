resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind2" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key2"
    value = "value2"
}
