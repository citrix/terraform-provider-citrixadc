package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/rdp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRdpserverprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRdpserverprofileFunc,
		Read:          readRdpserverprofileFunc,
		Update:        updateRdpserverprofileFunc,
		Delete:        deleteRdpserverprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"psk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rdpredirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRdpserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Get("name").(string)
	rdpserverprofile := rdp.Rdpserverprofile{
		Name:           d.Get("name").(string),
		Psk:            d.Get("psk").(string),
		Rdpip:          d.Get("rdpip").(string),
		Rdpport:        d.Get("rdpport").(int),
		Rdpredirection: d.Get("rdpredirection").(string),
	}

	_, err := client.AddResource("rdpserverprofile", rdpserverprofileName, &rdpserverprofile)
	if err != nil {
		return err
	}

	d.SetId(rdpserverprofileName)

	err = readRdpserverprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rdpserverprofile but we can't read it ?? %s", rdpserverprofileName)
		return nil
	}
	return nil
}

func readRdpserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rdpserverprofile state %s", rdpserverprofileName)
	data, err := client.FindResource("rdpserverprofile", rdpserverprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rdpserverprofile state %s", rdpserverprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rdpip", data["rdpip"])
	d.Set("rdpport", data["rdpport"])
	d.Set("rdpredirection", data["rdpredirection"])

	return nil

}

func updateRdpserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Get("name").(string)

	rdpserverprofile := rdp.Rdpserverprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("psk") {
		log.Printf("[DEBUG]  citrixadc-provider: Psk has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Psk = d.Get("psk").(string)
		hasChange = true
	}
	if d.HasChange("rdpip") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpip has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Rdpip = d.Get("rdpip").(string)
		hasChange = true
	}
	if d.HasChange("rdpport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpport has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Rdpport = d.Get("rdpport").(int)
		hasChange = true
	}
	if d.HasChange("rdpredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpredirection has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Rdpredirection = d.Get("rdpredirection").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("rdpserverprofile", &rdpserverprofile)
		if err != nil {
			return fmt.Errorf("Error updating rdpserverprofile %s", rdpserverprofileName)
		}
	}
	return readRdpserverprofileFunc(d, meta)
}

func deleteRdpserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Id()
	err := client.DeleteResource("rdpserverprofile", rdpserverprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
