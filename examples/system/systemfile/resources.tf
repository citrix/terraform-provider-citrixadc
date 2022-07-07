resource "citrixadc_systemfile" "tf_file" {
    filename = "resources_copy.tf"
    filelocation = "/var/tmp"
    filecontent = file("resources.tf")
}
