resource "citrixadc_appfwfieldtype" "demo_appfwfieldtype" {
  name = "demo_appfwfieldtype"
  regex = "test_.*regex"
  priority = "100"
  # comment = "test comment"
}
