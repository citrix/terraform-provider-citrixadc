package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcClusternodegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodegroupFunc,
		Read:          readClusternodegroupFunc,
		Update:        updateClusternodegroupFunc,
		Delete:        deleteClusternodegroupFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"strict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createClusternodegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodegroupName := d.Get("name").(string)

	clusternodegroup := cluster.Clusternodegroup{
		Name:     d.Get("name").(string),
		Priority: d.Get("priority").(int),
		State:    d.Get("state").(string),
		Sticky:   d.Get("sticky").(string),
		Strict:   d.Get("strict").(string),
	}

	_, err := client.AddResource(service.Clusternodegroup.Type(), clusternodegroupName, &clusternodegroup)
	if err != nil {
		return err
	}

	d.SetId(clusternodegroupName)

	err = readClusternodegroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternodegroup but we can't read it ?? %s", clusternodegroupName)
		return nil
	}
	return nil
}

func readClusternodegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodegroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup state %s", clusternodegroupName)
	data, err := client.FindResource(service.Clusternodegroup.Type(), clusternodegroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup state %s", clusternodegroupName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	//setToInt("priority", d, data["priority"])
	//d.Set("state", data["state"])
	//d.Set("sticky", data["sticky"])
	//d.Set("strict", data["strict"])

	return nil

}

func updateClusternodegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusternodegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodegroupName := d.Get("name").(string)

	clusternodegroup := cluster.Clusternodegroup{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for clusternodegroup %s, starting update", clusternodegroupName)
		clusternodegroup.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for clusternodegroup %s, starting update", clusternodegroupName)
		clusternodegroup.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("sticky") {
		log.Printf("[DEBUG]  citrixadc-provider: Sticky has changed for clusternodegroup %s, starting update", clusternodegroupName)
		clusternodegroup.Sticky = d.Get("sticky").(string)
		hasChange = true
	}
	if d.HasChange("strict") {
		log.Printf("[DEBUG]  citrixadc-provider: Strict has changed for clusternodegroup %s, starting update", clusternodegroupName)
		clusternodegroup.Strict = d.Get("strict").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Clusternodegroup.Type(), &clusternodegroup)
		if err != nil {
			return fmt.Errorf("Error updating clusternodegroup %s", clusternodegroupName)
		}
	}
	return readClusternodegroupFunc(d, meta)
}

func deleteClusternodegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodegroupName := d.Id()
	err := client.DeleteResource(service.Clusternodegroup.Type(), clusternodegroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
