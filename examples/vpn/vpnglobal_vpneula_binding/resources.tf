resource "citrixadc_vpneula" "tf_vpneula" {
	name = "tf_vpneula"	
}
resource "citrixadc_vpnglobal_vpneula_binding" "tf_bind" {
  eula = citrixadc_vpneula.tf_vpneula.name
}