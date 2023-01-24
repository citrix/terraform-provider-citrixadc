package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ica"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIcalatencyprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIcalatencyprofileFunc,
		Read:          readIcalatencyprofileFunc,
		Update:        updateIcalatencyprofileFunc,
		Delete:        deleteIcalatencyprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"l7latencymaxnotifycount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencymonitoring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l7latencynotifyinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencythresholdfactor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencywaittime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcalatencyprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Get("name").(string)
	icalatencyprofile := ica.Icalatencyprofile{
		L7latencymaxnotifycount:  d.Get("l7latencymaxnotifycount").(int),
		L7latencymonitoring:      d.Get("l7latencymonitoring").(string),
		L7latencynotifyinterval:  d.Get("l7latencynotifyinterval").(int),
		L7latencythresholdfactor: d.Get("l7latencythresholdfactor").(int),
		L7latencywaittime:        d.Get("l7latencywaittime").(int),
		Name:                     d.Get("name").(string),
	}

	_, err := client.AddResource("icalatencyprofile", icalatencyprofileName, &icalatencyprofile)
	if err != nil {
		return err
	}

	d.SetId(icalatencyprofileName)

	err = readIcalatencyprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this icalatencyprofile but we can't read it ?? %s", icalatencyprofileName)
		return nil
	}
	return nil
}

func readIcalatencyprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icalatencyprofile state %s", icalatencyprofileName)
	data, err := client.FindResource("icalatencyprofile", icalatencyprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icalatencyprofile state %s", icalatencyprofileName)
		d.SetId("")
		return nil
	}
	d.Set("l7latencymaxnotifycount", data["l7latencymaxnotifycount"])
	d.Set("l7latencymonitoring", data["l7latencymonitoring"])
	d.Set("l7latencynotifyinterval", data["l7latencynotifyinterval"])
	d.Set("l7latencythresholdfactor", data["l7latencythresholdfactor"])
	d.Set("l7latencywaittime", data["l7latencywaittime"])
	d.Set("name", data["name"])

	return nil

}

func updateIcalatencyprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Get("name").(string)

	icalatencyprofile := ica.Icalatencyprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("l7latencymaxnotifycount") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencymaxnotifycount has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencymaxnotifycount = d.Get("l7latencymaxnotifycount").(int)
		hasChange = true
	}
	if d.HasChange("l7latencymonitoring") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencymonitoring has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencymonitoring = d.Get("l7latencymonitoring").(string)
		hasChange = true
	}
	if d.HasChange("l7latencynotifyinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencynotifyinterval has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencynotifyinterval = d.Get("l7latencynotifyinterval").(int)
		hasChange = true
	}
	if d.HasChange("l7latencythresholdfactor") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencythresholdfactor has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencythresholdfactor = d.Get("l7latencythresholdfactor").(int)
		hasChange = true
	}
	if d.HasChange("l7latencywaittime") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencywaittime has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencywaittime = d.Get("l7latencywaittime").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icalatencyprofile", &icalatencyprofile)
		if err != nil {
			return fmt.Errorf("Error updating icalatencyprofile %s", icalatencyprofileName)
		}
	}
	return readIcalatencyprofileFunc(d, meta)
}

func deleteIcalatencyprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Id()
	err := client.DeleteResource("icalatencyprofile", icalatencyprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
