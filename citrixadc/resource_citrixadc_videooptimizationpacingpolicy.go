package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVideooptimizationpacingpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVideooptimizationpacingpolicyFunc,
		Read:          readVideooptimizationpacingpolicyFunc,
		Update:        updateVideooptimizationpacingpolicyFunc,
		Delete:        deleteVideooptimizationpacingpolicyFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVideooptimizationpacingpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Get("name").(string)

	videooptimizationpacingpolicy := videooptimization.Videooptimizationpacingpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName, &videooptimizationpacingpolicy)
	if err != nil {
		return err
	}

	d.SetId(videooptimizationpacingpolicyName)

	err = readVideooptimizationpacingpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this videooptimizationpacingpolicy but we can't read it ?? %s", videooptimizationpacingpolicyName)
		return nil
	}
	return nil
}

func readVideooptimizationpacingpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationpacingpolicy state %s", videooptimizationpacingpolicyName)
	data, err := client.FindResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationpacingpolicy state %s", videooptimizationpacingpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateVideooptimizationpacingpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Get("name").(string)

	videooptimizationpacingpolicy := videooptimization.Videooptimizationpacingpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName, &videooptimizationpacingpolicy)
		if err != nil {
			return fmt.Errorf("Error updating videooptimizationpacingpolicy %s", videooptimizationpacingpolicyName)
		}
	}
	return readVideooptimizationpacingpolicyFunc(d, meta)
}

func deleteVideooptimizationpacingpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Id()
	err := client.DeleteResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
