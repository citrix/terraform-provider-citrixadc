package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcContentinspectionprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createContentinspectionprofileFunc,
		Read:          readContentinspectionprofileFunc,
		Update:        updateContentinspectionprofileFunc,
		Delete:        deleteContentinspectionprofileFunc,
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
			"egressinterface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egressvlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ingressinterface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingressvlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iptunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectionprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Get("name").(string)
	contentinspectionprofile := contentinspection.Contentinspectionprofile{
		Egressinterface:  d.Get("egressinterface").(string),
		Egressvlan:       d.Get("egressvlan").(int),
		Ingressinterface: d.Get("ingressinterface").(string),
		Ingressvlan:      d.Get("ingressvlan").(int),
		Iptunnel:         d.Get("iptunnel").(string),
		Name:             d.Get("name").(string),
		Type:             d.Get("type").(string),
	}

	_, err := client.AddResource("contentinspectionprofile", contentinspectionprofileName, &contentinspectionprofile)
	if err != nil {
		return err
	}

	d.SetId(contentinspectionprofileName)

	err = readContentinspectionprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this contentinspectionprofile but we can't read it ?? %s", contentinspectionprofileName)
		return nil
	}
	return nil
}

func readContentinspectionprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionprofile state %s", contentinspectionprofileName)
	data, err := client.FindResource("contentinspectionprofile", contentinspectionprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionprofile state %s", contentinspectionprofileName)
		d.SetId("")
		return nil
	}
	d.Set("egressinterface", data["egressinterface"])
	d.Set("egressvlan", data["egressvlan"])
	d.Set("ingressinterface", data["ingressinterface"])
	d.Set("ingressvlan", data["ingressvlan"])
	d.Set("iptunnel", data["iptunnel"])
	d.Set("name", data["name"])
	d.Set("type", data["type"])

	return nil

}

func updateContentinspectionprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Get("name").(string)

	contentinspectionprofile := contentinspection.Contentinspectionprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("egressinterface") {
		log.Printf("[DEBUG]  citrixadc-provider: Egressinterface has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Egressinterface = d.Get("egressinterface").(string)
		hasChange = true
	}
	if d.HasChange("egressvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Egressvlan has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Egressvlan = d.Get("egressvlan").(int)
		hasChange = true
	}
	if d.HasChange("ingressinterface") {
		log.Printf("[DEBUG]  citrixadc-provider: Ingressinterface has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Ingressinterface = d.Get("ingressinterface").(string)
		hasChange = true
	}
	if d.HasChange("ingressvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Ingressvlan has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Ingressvlan = d.Get("ingressvlan").(int)
		hasChange = true
	}
	if d.HasChange("iptunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Iptunnel has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Iptunnel = d.Get("iptunnel").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionprofile", &contentinspectionprofile)
		if err != nil {
			return fmt.Errorf("Error updating contentinspectionprofile %s", contentinspectionprofileName)
		}
	}
	return readContentinspectionprofileFunc(d, meta)
}

func deleteContentinspectionprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Id()
	err := client.DeleteResource("contentinspectionprofile", contentinspectionprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
