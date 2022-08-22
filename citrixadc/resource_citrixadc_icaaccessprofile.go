package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ica"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIcaaccessprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIcaaccessprofileFunc,
		Read:          readIcaaccessprofileFunc,
		Update:        updateIcaaccessprofileFunc,
		Delete:        deleteIcaaccessprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"clientaudioredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientclipboardredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientcomportredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientdriveredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientprinterredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientusbdriveredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connectclientlptports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localremotedatasharing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multistream": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcaaccessprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Get("name").(string)
	icaaccessprofile := ica.Icaaccessprofile{
		Clientaudioredirection:     d.Get("clientaudioredirection").(string),
		Clientclipboardredirection: d.Get("clientclipboardredirection").(string),
		Clientcomportredirection:   d.Get("clientcomportredirection").(string),
		Clientdriveredirection:     d.Get("clientdriveredirection").(string),
		Clientprinterredirection:   d.Get("clientprinterredirection").(string),
		Clientusbdriveredirection:  d.Get("clientusbdriveredirection").(string),
		Connectclientlptports:      d.Get("connectclientlptports").(string),
		Localremotedatasharing:     d.Get("localremotedatasharing").(string),
		Multistream:                d.Get("multistream").(string),
		Name:                       d.Get("name").(string),
	}

	_, err := client.AddResource("icaaccessprofile", icaaccessprofileName, &icaaccessprofile)
	if err != nil {
		return err
	}

	d.SetId(icaaccessprofileName)

	err = readIcaaccessprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this icaaccessprofile but we can't read it ?? %s", icaaccessprofileName)
		return nil
	}
	return nil
}

func readIcaaccessprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icaaccessprofile state %s", icaaccessprofileName)
	data, err := client.FindResource("icaaccessprofile", icaaccessprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icaaccessprofile state %s", icaaccessprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("clientaudioredirection", data["clientaudioredirection"])
	d.Set("clientclipboardredirection", data["clientclipboardredirection"])
	d.Set("clientcomportredirection", data["clientcomportredirection"])
	d.Set("clientdriveredirection", data["clientdriveredirection"])
	d.Set("clientprinterredirection", data["clientprinterredirection"])
	d.Set("clientusbdriveredirection", data["clientusbdriveredirection"])
	d.Set("connectclientlptports", data["connectclientlptports"])
	d.Set("localremotedatasharing", data["localremotedatasharing"])
	d.Set("multistream", data["multistream"])

	return nil

}

func updateIcaaccessprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Get("name").(string)

	icaaccessprofile := ica.Icaaccessprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientaudioredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientaudioredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientaudioredirection = d.Get("clientaudioredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientclipboardredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientclipboardredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientclipboardredirection = d.Get("clientclipboardredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientcomportredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcomportredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientcomportredirection = d.Get("clientcomportredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientdriveredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientdriveredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientdriveredirection = d.Get("clientdriveredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientprinterredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientprinterredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientprinterredirection = d.Get("clientprinterredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientusbdriveredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientusbdriveredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientusbdriveredirection = d.Get("clientusbdriveredirection").(string)
		hasChange = true
	}
	if d.HasChange("connectclientlptports") {
		log.Printf("[DEBUG]  citrixadc-provider: Connectclientlptports has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Connectclientlptports = d.Get("connectclientlptports").(string)
		hasChange = true
	}
	if d.HasChange("localremotedatasharing") {
		log.Printf("[DEBUG]  citrixadc-provider: Localremotedatasharing has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Localremotedatasharing = d.Get("localremotedatasharing").(string)
		hasChange = true
	}
	if d.HasChange("multistream") {
		log.Printf("[DEBUG]  citrixadc-provider: Multistream has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Multistream = d.Get("multistream").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icaaccessprofile",  &icaaccessprofile)
		if err != nil {
			return fmt.Errorf("Error updating icaaccessprofile %s", icaaccessprofileName)
		}
	}
	return readIcaaccessprofileFunc(d, meta)
}

func deleteIcaaccessprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Id()
	err := client.DeleteResource("icaaccessprofile", icaaccessprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
