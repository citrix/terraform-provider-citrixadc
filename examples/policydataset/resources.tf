resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value" {
  name = citrixadc_policydataset.tf_dataset.name

  for_each = var.inputdata
  value    = each.key
  index    = each.value["index"]
}
