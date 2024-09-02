package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaapreauthenticationaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaapreauthenticationactionFunc,
		Read:          readAaapreauthenticationactionFunc,
		Update:        updateAaapreauthenticationactionFunc,
		Delete:        deleteAaapreauthenticationactionFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultepagroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deletefiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killprocess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preauthenticationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaapreauthenticationactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Get("name").(string)

	aaapreauthenticationaction := aaa.Aaapreauthenticationaction{
		Defaultepagroup:         d.Get("defaultepagroup").(string),
		Deletefiles:             d.Get("deletefiles").(string),
		Killprocess:             d.Get("killprocess").(string),
		Name:                    d.Get("name").(string),
		Preauthenticationaction: d.Get("preauthenticationaction").(string),
	}

	_, err := client.AddResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName, &aaapreauthenticationaction)
	if err != nil {
		return err
	}

	d.SetId(aaapreauthenticationactionName)

	err = readAaapreauthenticationactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaapreauthenticationaction but we can't read it ?? %s", aaapreauthenticationactionName)
		return nil
	}
	return nil
}

func readAaapreauthenticationactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading aaapreauthenticationaction state %s", aaapreauthenticationactionName)
	data, err := client.FindResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaapreauthenticationaction state %s", aaapreauthenticationactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("defaultepagroup", data["defaultepagroup"])
	d.Set("deletefiles", data["deletefiles"])
	d.Set("killprocess", data["killprocess"])
	d.Set("name", data["name"])
	d.Set("preauthenticationaction", data["preauthenticationaction"])

	return nil

}

func updateAaapreauthenticationactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Get("name").(string)

	aaapreauthenticationaction := aaa.Aaapreauthenticationaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultepagroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultepagroup has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Defaultepagroup = d.Get("defaultepagroup").(string)
		hasChange = true
	}
	if d.HasChange("deletefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefiles has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Deletefiles = d.Get("deletefiles").(string)
		hasChange = true
	}
	if d.HasChange("killprocess") {
		log.Printf("[DEBUG]  citrixadc-provider: Killprocess has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Killprocess = d.Get("killprocess").(string)
		hasChange = true
	}
	if d.HasChange("preauthenticationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Preauthenticationaction has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Preauthenticationaction = d.Get("preauthenticationaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName, &aaapreauthenticationaction)
		if err != nil {
			return fmt.Errorf("Error updating aaapreauthenticationaction %s", aaapreauthenticationactionName)
		}
	}
	return readAaapreauthenticationactionFunc(d, meta)
}

func deleteAaapreauthenticationactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Id()
	err := client.DeleteResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
