package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVideooptimizationdetectionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVideooptimizationdetectionactionFunc,
		Read:          readVideooptimizationdetectionactionFunc,
		Update:        updateVideooptimizationdetectionactionFunc,
		Delete:        deleteVideooptimizationdetectionactionFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
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

func createVideooptimizationdetectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Get("name").(string)
	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Newname: d.Get("newname").(string),
		Type:    d.Get("type").(string),
	}

	_, err := client.AddResource("videooptimizationdetectionaction", videooptimizationdetectionactionName, &videooptimizationdetectionaction)
	if err != nil {
		return err
	}

	d.SetId(videooptimizationdetectionactionName)

	err = readVideooptimizationdetectionactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this videooptimizationdetectionaction but we can't read it ?? %s", videooptimizationdetectionactionName)
		return nil
	}
	return nil
}

func readVideooptimizationdetectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationdetectionaction state %s", videooptimizationdetectionactionName)
	data, err := client.FindResource("videooptimizationdetectionaction", videooptimizationdetectionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationdetectionaction state %s", videooptimizationdetectionactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("type", data["type"])

	return nil

}

func updateVideooptimizationdetectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Get("name").(string)

	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationdetectionaction %s, starting update", videooptimizationdetectionactionName)
		videooptimizationdetectionaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for videooptimizationdetectionaction %s, starting update", videooptimizationdetectionactionName)
		videooptimizationdetectionaction.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for videooptimizationdetectionaction %s, starting update", videooptimizationdetectionactionName)
		videooptimizationdetectionaction.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationdetectionaction", videooptimizationdetectionactionName, &videooptimizationdetectionaction)
		if err != nil {
			return fmt.Errorf("Error updating videooptimizationdetectionaction %s", videooptimizationdetectionactionName)
		}
	}
	return readVideooptimizationdetectionactionFunc(d, meta)
}

func deleteVideooptimizationdetectionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Id()
	err := client.DeleteResource("videooptimizationdetectionaction", videooptimizationdetectionactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
