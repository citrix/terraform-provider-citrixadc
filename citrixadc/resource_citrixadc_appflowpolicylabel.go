package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppflowpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppflowpolicylabelFunc,
		Read:          readAppflowpolicylabelFunc,
		Delete:        deleteAppflowpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppflowpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicylabelName := d.Get("labelname").(string)
	
	appflowpolicylabel := appflow.Appflowpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Appflowpolicylabel.Type(), appflowpolicylabelName, &appflowpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(appflowpolicylabelName)

	err = readAppflowpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appflowpolicylabel but we can't read it ?? %s", appflowpolicylabelName)
		return nil
	}
	return nil
}

func readAppflowpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowpolicylabel state %s", appflowpolicylabelName)
	data, err := client.FindResource(service.Appflowpolicylabel.Type(), appflowpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowpolicylabel state %s", appflowpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func deleteAppflowpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicylabelName := d.Id()
	err := client.DeleteResource(service.Appflowpolicylabel.Type(), appflowpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
