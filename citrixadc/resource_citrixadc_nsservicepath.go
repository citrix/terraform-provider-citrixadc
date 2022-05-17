package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcNsservicepath() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsservicepathFunc,
		Read:          readNsservicepathFunc,
		Delete:        deleteNsservicepathFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"servicepathname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNsservicepathFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsservicepathFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicepathName := d.Get("servicepathname").(string)
	nsservicepath := ns.Nsservicepath{
		Servicepathname: d.Get("servicepathname").(string),
	}

	_, err := client.AddResource(service.Nsservicepath.Type(), nsservicepathName, &nsservicepath)
	if err != nil {
		return err
	}

	d.SetId(nsservicepathName)

	err = readNsservicepathFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsservicepath but we can't read it ?? %s", nsservicepathName)
		return nil
	}
	return nil
}

func readNsservicepathFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsservicepathFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicepathName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsservicepath state %s", nsservicepathName)
	data, err := client.FindResource(service.Nsservicepath.Type(), nsservicepathName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicepath state %s", nsservicepathName)
		d.SetId("")
		return nil
	}
	d.Set("servicepathname", data["servicepathname"])

	return nil

}

func deleteNsservicepathFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsservicepathFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicepathName := d.Id()
	err := client.DeleteResource(service.Nsservicepath.Type(), nsservicepathName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
