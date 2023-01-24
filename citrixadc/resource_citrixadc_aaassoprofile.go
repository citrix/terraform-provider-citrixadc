package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaassoprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaassoprofileFunc,
		Read:          readAaassoprofileFunc,
		Update:        updateAaassoprofileFunc,
		Delete:        deleteAaassoprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createAaassoprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Get("name").(string)
	aaassoprofile := aaa.Aaassoprofile{
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
		Username: d.Get("username").(string),
	}

	_, err := client.AddResource("aaassoprofile", aaassoprofileName, &aaassoprofile)
	if err != nil {
		return err
	}

	d.SetId(aaassoprofileName)

	err = readAaassoprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaassoprofile but we can't read it ?? %s", aaassoprofileName)
		return nil
	}
	return nil
}

func readAaassoprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaassoprofile state %s", aaassoprofileName)
	data, err := client.FindResource("aaassoprofile", aaassoprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaassoprofile state %s", aaassoprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	//d.Set("password", data["password"])
	d.Set("username", data["username"])

	return nil

}

func updateAaassoprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Get("name").(string)

	aaassoprofile := aaa.Aaassoprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false

	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for aaassoprofile %s, starting update", aaassoprofileName)
		aaassoprofile.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG]  citrixadc-provider: Username has changed for aaassoprofile %s, starting update", aaassoprofileName)
		aaassoprofile.Username = d.Get("username").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("aaassoprofile", &aaassoprofile)
		if err != nil {
			return fmt.Errorf("Error updating aaassoprofile %s", aaassoprofileName)
		}
	}
	return readAaassoprofileFunc(d, meta)
}

func deleteAaassoprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Id()
	err := client.DeleteResource("aaassoprofile", aaassoprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
