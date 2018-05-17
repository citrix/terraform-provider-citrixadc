package netscaler

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerNsacls() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsaclsFunc,
		Read:          readNsaclsFunc,
		Update:        updateNsaclsFunc,
		Delete:        deleteNsaclsFunc,
		Schema:        map[string]*schema.Schema{},
	}
}

func createNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	nsacls := ns.Nsacls{}

	_, err := client.ApplyResource(netscaler.Nsacls.Type(), &nsacls)
	if err != nil {
		return err
	}

	d.SetId(nsaclsName)

	return nil
}

func readNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readNsaclsFunc")

	return nil
}

func updateNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateNsaclsFunc")
	return nil
}

func deleteNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteNsaclsFunc")

	d.SetId("")
	return nil
}
