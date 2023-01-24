package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcContentinspectionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createContentinspectionactionFunc,
		Read:          readContentinspectionactionFunc,
		Update:        updateContentinspectionactionFunc,
		Delete:        deleteContentinspectionactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"icapprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ifserverdown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Get("name").(string)
	contentinspectionaction := contentinspection.Contentinspectionaction{
		Icapprofilename: d.Get("icapprofilename").(string),
		Ifserverdown:    d.Get("ifserverdown").(string),
		Name:            d.Get("name").(string),
		Serverip:        d.Get("serverip").(string),
		Servername:      d.Get("servername").(string),
		Serverport:      d.Get("serverport").(int),
		Type:            d.Get("type").(string),
	}

	_, err := client.AddResource("contentinspectionaction", contentinspectionactionName, &contentinspectionaction)
	if err != nil {
		return err
	}

	d.SetId(contentinspectionactionName)

	err = readContentinspectionactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this contentinspectionaction but we can't read it ?? %s", contentinspectionactionName)
		return nil
	}
	return nil
}

func readContentinspectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionaction state %s", contentinspectionactionName)
	data, err := client.FindResource("contentinspectionaction", contentinspectionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionaction state %s", contentinspectionactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("icapprofilename", data["icapprofilename"])
	d.Set("ifserverdown", data["ifserverdown"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	d.Set("serverport", data["serverport"])
	d.Set("type", data["type"])

	return nil

}

func updateContentinspectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Get("name").(string)

	contentinspectionaction := contentinspection.Contentinspectionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("icapprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Icapprofilename has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Icapprofilename = d.Get("icapprofilename").(string)
		hasChange = true
	}
	if d.HasChange("ifserverdown") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifserverdown has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Ifserverdown = d.Get("ifserverdown").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Serverport = d.Get("serverport").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionaction", &contentinspectionaction)
		if err != nil {
			return fmt.Errorf("Error updating contentinspectionaction %s", contentinspectionactionName)
		}
	}
	return readContentinspectionactionFunc(d, meta)
}

func deleteContentinspectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Id()
	err := client.DeleteResource("contentinspectionaction", contentinspectionactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
