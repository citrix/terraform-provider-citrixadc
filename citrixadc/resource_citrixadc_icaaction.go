package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIcaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIcaactionFunc,
		Read:          readIcaactionFunc,
		Update:        updateIcaactionFunc,
		Delete:        deleteIcaactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accessprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"latencyprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Get("name").(string)
	icaaction := ica.Icaaction{
		Accessprofilename:  d.Get("accessprofilename").(string),
		Latencyprofilename: d.Get("latencyprofilename").(string),
		Name:               d.Get("name").(string),
	}

	_, err := client.AddResource("icaaction", icaactionName, &icaaction)
	if err != nil {
		return err
	}

	d.SetId(icaactionName)

	err = readIcaactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this icaaction but we can't read it ?? %s", icaactionName)
		return nil
	}
	return nil
}

func readIcaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icaaction state %s", icaactionName)
	data, err := client.FindResource("icaaction", icaactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icaaction state %s", icaactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("accessprofilename", data["accessprofilename"])
	d.Set("latencyprofilename", data["latencyprofilename"])

	return nil

}

func updateIcaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Get("name").(string)

	icaaction := ica.Icaaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("accessprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Accessprofilename has changed for icaaction %s, starting update", icaactionName)
		icaaction.Accessprofilename = d.Get("accessprofilename").(string)
		hasChange = true
	}
	if d.HasChange("latencyprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Latencyprofilename has changed for icaaction %s, starting update", icaactionName)
		icaaction.Latencyprofilename = d.Get("latencyprofilename").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icaaction", &icaaction)
		if err != nil {
			return fmt.Errorf("Error updating icaaction %s", icaactionName)
		}
	}
	return readIcaactionFunc(d, meta)
}

func deleteIcaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Id()
	err := client.DeleteResource("icaaction", icaactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
