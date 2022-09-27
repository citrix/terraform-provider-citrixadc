resource "null_resource" "do_nitro_request" {

  provisioner "local-exec" {
    environment = {
      NSIP       = "10.10.10.10"
      NITRO_PASS = "secret"
      PROMPT     = "Hello Again"
      SERVICE    = "my_service"
      LBVSERVER  = "my_lbvserver"
    }

    interpreter = ["bash"]
    command     = "do_nitro_request.sh"
  }

}
