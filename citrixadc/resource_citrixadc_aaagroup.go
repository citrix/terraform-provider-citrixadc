package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAaagroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaagroupFunc,
		Read:          readAaagroupFunc,
		Delete:        deleteAaagroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loggedin": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createAaagroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroupFunc")
	client := meta.(*NetScalerNitroClient).client
	aaagroupName := d.Get("groupname").(string)
	aaagroup := aaa.Aaagroup{
		Groupname: d.Get("groupname").(string),
		Weight:    d.Get("weight").(int),
	}

	_, err := client.AddResource(service.Aaagroup.Type(), aaagroupName, &aaagroup)
	if err != nil {
		return err
	}

	d.SetId(aaagroupName)

	err = readAaagroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaagroup but we can't read it ?? %s", aaagroupName)
		return nil
	}
	return nil
}

func readAaagroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroupFunc")
	client := meta.(*NetScalerNitroClient).client
	aaagroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup state %s", aaagroupName)
	data, err := client.FindResource(service.Aaagroup.Type(), aaagroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup state %s", aaagroupName)
		d.SetId("")
		return nil
	}
	d.Set("groupname", data["groupname"])
	d.Set("loggedin", data["loggedin"])
	setToInt("weight", d, data["weight"])

	return nil

}
func deleteAaagroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroupFunc")
	client := meta.(*NetScalerNitroClient).client
	aaagroupName := d.Id()
	err := client.DeleteResource(service.Aaagroup.Type(), aaagroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
