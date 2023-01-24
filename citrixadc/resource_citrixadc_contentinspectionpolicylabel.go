package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcContentinspectionpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createContentinspectionpolicylabelFunc,
		Read:          readContentinspectionpolicylabelFunc,
		Delete:        deleteContentinspectionpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createContentinspectionpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicylabelName := d.Get("labelname").(string)
	contentinspectionpolicylabel := contentinspection.Contentinspectionpolicylabel{
		Comment:   d.Get("comment").(string),
		Labelname: d.Get("labelname").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource("contentinspectionpolicylabel", contentinspectionpolicylabelName, &contentinspectionpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(contentinspectionpolicylabelName)

	err = readContentinspectionpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this contentinspectionpolicylabel but we can't read it ?? %s", contentinspectionpolicylabelName)
		return nil
	}
	return nil
}

func readContentinspectionpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionpolicylabel state %s", contentinspectionpolicylabelName)
	data, err := client.FindResource("contentinspectionpolicylabel", contentinspectionpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionpolicylabel state %s", contentinspectionpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("type", data["type"])

	return nil

}

func deleteContentinspectionpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicylabelName := d.Id()
	err := client.DeleteResource("contentinspectionpolicylabel", contentinspectionpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
