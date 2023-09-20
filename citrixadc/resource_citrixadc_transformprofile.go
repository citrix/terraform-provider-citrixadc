package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/transform"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTransformprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTransformprofileFunc,
		Read:          readTransformprofileFunc,
		Update:        updateTransformprofileFunc,
		Delete:        deleteTransformprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"onlytransformabsurlinbody": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTransformprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTransformprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	transformprofileName := d.Get("name").(string)

	transformprofileNew := transform.Transformprofile{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}

	_, err := client.AddResource(service.Transformprofile.Type(), transformprofileName, &transformprofileNew)
	if err != nil {
		return err
	}

	// Need to also update to include the parameters that are
	// invalid for the create operation
	transformprofile := transform.Transformprofile{
		Comment:                   d.Get("comment").(string),
		Name:                      d.Get("name").(string),
		Onlytransformabsurlinbody: d.Get("onlytransformabsurlinbody").(string),
		Type:                      d.Get("type").(string),
	}

	// Update will fail if only the name attribute is present
	doUpdate := false
	if transformprofile.Comment != "" {
		doUpdate = true
	}
	if transformprofile.Onlytransformabsurlinbody != "" {
		doUpdate = true
	}
	if transformprofile.Type != "" {
		doUpdate = true
	}
	if doUpdate {
		_, err := client.UpdateResource(service.Transformprofile.Type(), transformprofileName, &transformprofile)
		if err != nil {
			return err
		}
	}

	d.SetId(transformprofileName)

	err = readTransformprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this transformprofile but we can't read it ?? %s", transformprofileName)
		return nil
	}
	return nil
}

func readTransformprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTransformprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	transformprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading transformprofile state %s", transformprofileName)
	data, err := client.FindResource(service.Transformprofile.Type(), transformprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing transformprofile state %s", transformprofileName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("onlytransformabsurlinbody", data["onlytransformabsurlinbody"])
	d.Set("type", data["type"])

	return nil

}

func updateTransformprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTransformprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	transformprofileName := d.Get("name").(string)

	transformprofile := transform.Transformprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for transformprofile %s, starting update", transformprofileName)
		transformprofile.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for transformprofile %s, starting update", transformprofileName)
		transformprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("onlytransformabsurlinbody") {
		log.Printf("[DEBUG]  citrixadc-provider: Onlytransformabsurlinbody has changed for transformprofile %s, starting update", transformprofileName)
		transformprofile.Onlytransformabsurlinbody = d.Get("onlytransformabsurlinbody").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for transformprofile %s, starting update", transformprofileName)
		transformprofile.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Transformprofile.Type(), transformprofileName, &transformprofile)
		if err != nil {
			return fmt.Errorf("Error updating transformprofile %s", transformprofileName)
		}
	}
	return readTransformprofileFunc(d, meta)
}

func deleteTransformprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTransformprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	transformprofileName := d.Id()
	err := client.DeleteResource(service.Transformprofile.Type(), transformprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
