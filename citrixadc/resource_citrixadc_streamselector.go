package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcStreamselector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createStreamselectorFunc,
		Read:          readStreamselectorFunc,
		Update:        updateStreamselectorFunc,
		Delete:        deleteStreamselectorFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createStreamselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Get("name").(string)
	streamselector := stream.Streamselector{
		Name: d.Get("name").(string),
		Rule: toStringList(d.Get("rule").([]interface{})),
	}

	_, err := client.AddResource(service.Streamselector.Type(), streamselectorName, &streamselector)
	if err != nil {
		return err
	}

	d.SetId(streamselectorName)

	err = readStreamselectorFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this streamselector but we can't read it ?? %s", streamselectorName)
		return nil
	}
	return nil
}

func readStreamselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading streamselector state %s", streamselectorName)
	data, err := client.FindResource(service.Streamselector.Type(), streamselectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing streamselector state %s", streamselectorName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateStreamselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Get("name").(string)

	streamselector := stream.Streamselector{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for streamselector %s, starting update", streamselectorName)
		streamselector.Rule = toStringList(d.Get("rule").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Streamselector.Type(), &streamselector)
		if err != nil {
			return fmt.Errorf("Error updating streamselector %s", streamselectorName)
		}
	}
	return readStreamselectorFunc(d, meta)
}

func deleteStreamselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Id()
	err := client.DeleteResource(service.Streamselector.Type(), streamselectorName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
