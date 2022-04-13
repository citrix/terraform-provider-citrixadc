package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcVxlanvlanmap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVxlanvlanmapFunc,
		Read:          readVxlanvlanmapFunc,
		Delete:        deleteVxlanvlanmapFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVxlanvlanmapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlanvlanmapFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanvlanmapName := d.Get("name").(string)
	vxlanvlanmap := network.Vxlanvlanmap{
		Name: d.Get("name").(string),
	}

	_, err := client.AddResource("vxlanvlanmap", vxlanvlanmapName, &vxlanvlanmap)
	if err != nil {
		return err
	}

	d.SetId(vxlanvlanmapName)

	err = readVxlanvlanmapFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vxlanvlanmap but we can't read it ?? %s", vxlanvlanmapName)
		return nil
	}
	return nil
}

func readVxlanvlanmapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlanvlanmapFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanvlanmapName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vxlanvlanmap state %s", vxlanvlanmapName)
	data, err := client.FindResource("vxlanvlanmap", vxlanvlanmapName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vxlanvlanmap state %s", vxlanvlanmapName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteVxlanvlanmapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlanvlanmapFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanvlanmapName := d.Id()
	err := client.DeleteResource("vxlanvlanmap", vxlanvlanmapName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
