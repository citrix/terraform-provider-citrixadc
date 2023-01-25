package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcContentinspectioncallout() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createContentinspectioncalloutFunc,
		Read:          readContentinspectioncalloutFunc,
		Update:        updateContentinspectioncalloutFunc,
		Delete:        deleteContentinspectioncalloutFunc,
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
			"resultexpr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"returntype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profilename": &schema.Schema{
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

func createContentinspectioncalloutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Get("name").(string)
	contentinspectioncallout := contentinspection.Contentinspectioncallout{
		Comment:     d.Get("comment").(string),
		Name:        d.Get("name").(string),
		Profilename: d.Get("profilename").(string),
		Resultexpr:  d.Get("resultexpr").(string),
		Returntype:  d.Get("returntype").(string),
		Serverip:    d.Get("serverip").(string),
		Servername:  d.Get("servername").(string),
		Serverport:  d.Get("serverport").(int),
		Type:        d.Get("type").(string),
	}

	_, err := client.AddResource("contentinspectioncallout", contentinspectioncalloutName, &contentinspectioncallout)
	if err != nil {
		return err
	}

	d.SetId(contentinspectioncalloutName)

	err = readContentinspectioncalloutFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this contentinspectioncallout but we can't read it ?? %s", contentinspectioncalloutName)
		return nil
	}
	return nil
}

func readContentinspectioncalloutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectioncallout state %s", contentinspectioncalloutName)
	data, err := client.FindResource("contentinspectioncallout", contentinspectioncalloutName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectioncallout state %s", contentinspectioncalloutName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("profilename", data["profilename"])
	d.Set("resultexpr", data["resultexpr"])
	d.Set("returntype", data["returntype"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	d.Set("serverport", data["serverport"])
	d.Set("type", data["type"])

	return nil

}

func updateContentinspectioncalloutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Get("name").(string)

	contentinspectioncallout := contentinspection.Contentinspectioncallout{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("resultexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Resultexpr has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Resultexpr = d.Get("resultexpr").(string)
		hasChange = true
	}
	if d.HasChange("returntype") {
		log.Printf("[DEBUG]  citrixadc-provider: Returntype has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Returntype = d.Get("returntype").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Serverport = d.Get("serverport").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectioncallout", &contentinspectioncallout)
		if err != nil {
			return fmt.Errorf("Error updating contentinspectioncallout %s", contentinspectioncalloutName)
		}
	}
	return readContentinspectioncalloutFunc(d, meta)
}

func deleteContentinspectioncalloutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Id()
	err := client.DeleteResource("contentinspectioncallout", contentinspectioncalloutName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
