package citrixadc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

// Quicbridge Profile is a new resource. Go Nitro is to be updated for this new resource
type quicbridgeprofile struct {
	Name             string `json:"name,omitempty"`
	Routingalgorithm string `json:"routingalgorithm,omitempty"`
	Serveridlength   int    `json:"serveridlength,omitempty"`
}

func resourceCitrixAdcQuicbridgeprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createQuicbridgeprofileFunc,
		Read:          readQuicbridgeprofileFunc,
		Update:        updateQuicbridgeprofileFunc,
		Delete:        deleteQuicbridgeprofileFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"routingalgorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serveridlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createQuicbridgeprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Get("name").(string)

	quicbridgeprofile := quicbridgeprofile{
		Name:             d.Get("name").(string),
		Routingalgorithm: d.Get("routingalgorithm").(string),
		Serveridlength:   d.Get("serveridlength").(int),
	}

	_, err := client.AddResource("quicbridgeprofile", quicbridgeprofileName, &quicbridgeprofile)
	if err != nil {
		return err
	}

	d.SetId(quicbridgeprofileName)

	err = readQuicbridgeprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this quicbridgeprofile but we can't read it ?? %s", quicbridgeprofileName)
		return nil
	}
	return nil
}

func readQuicbridgeprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading quicbridgeprofile state %s", quicbridgeprofileName)
	data, err := client.FindResource("quicbridgeprofile", quicbridgeprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing quicbridgeprofile state %s", quicbridgeprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("routingalgorithm", data["routingalgorithm"])
	d.Set("serveridlength", data["serveridlength"])

	return nil

}

func updateQuicbridgeprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Get("name").(string)

	quicbridgeprofile := quicbridgeprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for quicbridgeprofile %s, starting update", quicbridgeprofileName)
		quicbridgeprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("routingalgorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Routingalgorithm has changed for quicbridgeprofile %s, starting update", quicbridgeprofileName)
		quicbridgeprofile.Routingalgorithm = d.Get("routingalgorithm").(string)
		hasChange = true
	}
	if d.HasChange("serveridlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Serveridlength has changed for quicbridgeprofile %s, starting update", quicbridgeprofileName)
		quicbridgeprofile.Serveridlength = d.Get("serveridlength").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("quicbridgeprofile", quicbridgeprofileName, &quicbridgeprofile)
		if err != nil {
			return fmt.Errorf("Error updating quicbridgeprofile %s", quicbridgeprofileName)
		}
	}
	return readQuicbridgeprofileFunc(d, meta)
}

func deleteQuicbridgeprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Id()
	err := client.DeleteResource("quicbridgeprofile", quicbridgeprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
