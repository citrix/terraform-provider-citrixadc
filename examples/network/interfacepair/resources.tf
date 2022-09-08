resource "citrixadc_interfacepair" "tf_interfacepair" {
  interface_id = 2
  ifnum        = ["LA/2", "LA/3"]
}
