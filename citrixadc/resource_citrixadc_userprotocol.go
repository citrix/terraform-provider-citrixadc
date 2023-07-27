package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/user"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcUserprotocol() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createUserprotocolFunc,
		Read:          readUserprotocolFunc,
		Update:        updateUserprotocolFunc,
		Delete:        deleteUserprotocolFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"extension": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transport": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createUserprotocolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Get("name").(string)
	userprotocol := user.Userprotocol{
		Comment:   d.Get("comment").(string),
		Extension: d.Get("extension").(string),
		Name:      d.Get("name").(string),
		Transport: d.Get("transport").(string),
	}

	_, err := client.AddResource("userprotocol", userprotocolName, &userprotocol)
	if err != nil {
		return err
	}

	d.SetId(userprotocolName)

	err = readUserprotocolFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this userprotocol but we can't read it ?? %s", userprotocolName)
		return nil
	}
	return nil
}

func readUserprotocolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading userprotocol state %s", userprotocolName)
	data, err := client.FindResource("userprotocol", userprotocolName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing userprotocol state %s", userprotocolName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("extension", data["extension"])
	d.Set("transport", data["transport"])

	return nil

}

func updateUserprotocolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Get("name").(string)

	userprotocol := user.Userprotocol{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for userprotocol %s, starting update", userprotocolName)
		userprotocol.Comment = d.Get("comment").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("userprotocol", &userprotocol)
		if err != nil {
			return fmt.Errorf("Error updating userprotocol %s", userprotocolName)
		}
	}
	return readUserprotocolFunc(d, meta)
}

func deleteUserprotocolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Id()
	err := client.DeleteResource("userprotocol", userprotocolName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
