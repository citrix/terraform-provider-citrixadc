package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVideooptimizationpacingaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVideooptimizationpacingactionFunc,
		Read:          readVideooptimizationpacingactionFunc,
		Update:        updateVideooptimizationpacingactionFunc,
		Delete:        deleteVideooptimizationpacingactionFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rate": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVideooptimizationpacingactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Get("name").(string)

	videooptimizationpacingaction := videooptimization.Videooptimizationpacingaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Newname: d.Get("newname").(string),
		Rate:    d.Get("rate").(int),
	}

	_, err := client.AddResource("videooptimizationpacingaction", videooptimizationpacingactionName, &videooptimizationpacingaction)
	if err != nil {
		return err
	}

	d.SetId(videooptimizationpacingactionName)

	err = readVideooptimizationpacingactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this videooptimizationpacingaction but we can't read it ?? %s", videooptimizationpacingactionName)
		return nil
	}
	return nil
}

func readVideooptimizationpacingactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationpacingaction state %s", videooptimizationpacingactionName)
	data, err := client.FindResource("videooptimizationpacingaction", videooptimizationpacingactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationpacingaction state %s", videooptimizationpacingactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("newname", data["newname"])
	d.Set("rate", data["rate"])

	return nil

}

func updateVideooptimizationpacingactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Get("name").(string)

	videooptimizationpacingaction := videooptimization.Videooptimizationpacingaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationpacingaction %s, starting update", videooptimizationpacingactionName)
		videooptimizationpacingaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for videooptimizationpacingaction %s, starting update", videooptimizationpacingactionName)
		videooptimizationpacingaction.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rate") {
		log.Printf("[DEBUG]  citrixadc-provider: Rate has changed for videooptimizationpacingaction %s, starting update", videooptimizationpacingactionName)
		videooptimizationpacingaction.Rate = d.Get("rate").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationpacingaction", videooptimizationpacingactionName, &videooptimizationpacingaction)
		if err != nil {
			return fmt.Errorf("Error updating videooptimizationpacingaction %s", videooptimizationpacingactionName)
		}
	}
	return readVideooptimizationpacingactionFunc(d, meta)
}

func deleteVideooptimizationpacingactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Id()
	err := client.DeleteResource("videooptimizationpacingaction", videooptimizationpacingactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
