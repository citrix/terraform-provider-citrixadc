package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLbmetrictable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbmetrictableFunc,
		Read:          readLbmetrictableFunc,
		Delete:        deleteLbmetrictableFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLbmetrictableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbmetrictableFunc")
	client := meta.(*NetScalerNitroClient).client

	lbmetrictableName := d.Get("metrictable").(string)

	lbmetrictable := lb.Lbmetrictable{
		Metrictable: lbmetrictableName,
	}

	_, err := client.AddResource("lbmetrictable", lbmetrictableName, &lbmetrictable)
	if err != nil {
		return err
	}

	d.SetId(lbmetrictableName)

	err = readLbmetrictableFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbmetrictable but we can't read it ?? %s", lbmetrictableName)
		return nil
	}
	return nil
}

func readLbmetrictableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbmetrictableFunc")
	client := meta.(*NetScalerNitroClient).client
	lbmetrictableName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbmetrictable state %s", lbmetrictableName)
	data, err := client.FindResource("lbmetrictable", lbmetrictableName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbmetrictable state %s", lbmetrictableName)
		d.SetId("")
		return nil
	}
	d.Set("metrictable", data["metrictable"])

	return nil

}

func deleteLbmetrictableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbmetrictableFunc")
	client := meta.(*NetScalerNitroClient).client
	lbmetrictableName := d.Id()
	err := client.DeleteResource("lbmetrictable", lbmetrictableName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
