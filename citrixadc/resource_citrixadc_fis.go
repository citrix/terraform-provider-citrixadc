package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcFis() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFisFunc,
		Read:          readFisFunc,
		Delete:        deleteFisFunc,
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
			"ownernode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createFisFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFisFunc")
	client := meta.(*NetScalerNitroClient).client
	fisName := d.Get("name").(string)
	fis := network.Fis{
		Name:      d.Get("name").(string),
		Ownernode: d.Get("ownernode").(int),
	}

	_, err := client.AddResource(service.Fis.Type(), fisName, &fis)
	if err != nil {
		return err
	}

	d.SetId(fisName)

	err = readFisFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this fis but we can't read it ?? %s", fisName)
		return nil
	}
	return nil
}

func readFisFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFisFunc")
	client := meta.(*NetScalerNitroClient).client
	fisName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading fis state %s", fisName)
	data, err := client.FindResource(service.Fis.Type(), fisName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing fis state %s", fisName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("ownernode", data["ownernode"])

	return nil

}

func deleteFisFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFisFunc")
	client := meta.(*NetScalerNitroClient).client
	fisName := d.Id()
	err := client.DeleteResource(service.Fis.Type(), fisName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
