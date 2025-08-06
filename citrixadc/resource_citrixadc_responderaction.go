package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcResponderaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderactionFunc,
		Read:          readResponderactionFunc,
		Update:        updateResponderactionFunc,
		Delete:        deleteResponderactionFunc,
		CustomizeDiff: customizeResponderactionDiff,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bypasssafetycheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htmlpage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"reasonphrase": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"responsestatuscode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createResponderactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderactionFunc")
	client := meta.(*NetScalerNitroClient).client
	var responderactionName string
	if v, ok := d.GetOk("name"); ok {
		responderactionName = v.(string)
	} else {
		responderactionName = resource.PrefixedUniqueId("tf-responderaction-")
		d.Set("name", responderactionName)
	}
	responderaction := responder.Responderaction{
		Bypasssafetycheck:  d.Get("bypasssafetycheck").(string),
		Comment:            d.Get("comment").(string),
		Htmlpage:           d.Get("htmlpage").(string),
		Name:               d.Get("name").(string),
		Reasonphrase:       d.Get("reasonphrase").(string),
		Responsestatuscode: d.Get("responsestatuscode").(int),
		Target:             d.Get("target").(string),
		Type:               d.Get("type").(string),
	}

	_, err := client.AddResource(service.Responderaction.Type(), responderactionName, &responderaction)
	if err != nil {
		return err
	}

	d.SetId(responderactionName)

	err = readResponderactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderaction but we can't read it ?? %s", responderactionName)
		return nil
	}
	return nil
}

func readResponderactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderactionFunc")
	client := meta.(*NetScalerNitroClient).client
	responderactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderaction state %s", responderactionName)
	data, err := client.FindResource(service.Responderaction.Type(), responderactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderaction state %s", responderactionName)
		d.SetId("")
		return nil
	}

	d.Set("name", data["name"])
	d.Set("bypasssafetycheck", data["bypasssafetycheck"])
	d.Set("comment", data["comment"])
	d.Set("htmlpage", data["htmlpage"])
	d.Set("name", data["name"])
	d.Set("reasonphrase", data["reasonphrase"])
	d.Set("target", data["target"])
	d.Set("type", data["type"])
	// Check if the value from the API is a string and convert it to an int
	// before setting it in the Terraform state.
	if val, ok := data["responsestatuscode"].(string); ok {
		if valInt, err := strconv.Atoi(val); err == nil {
			d.Set("responsestatuscode", valInt)
		} else {
			// Log an error if the string cannot be converted to an integer.
			log.Printf("[ERROR] citrixadc-provider: Failed to convert responsestatuscode string to int: %v", err)
		}
	} else {
		// If the value is not a string, assume it's already an integer and set it directly.
		d.Set("responsestatuscode", data["responsestatuscode"])
	}

	return nil

}

func updateResponderactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateResponderactionFunc")
	client := meta.(*NetScalerNitroClient).client
	responderactionName := d.Get("name").(string)

	responderaction := responder.Responderaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bypasssafetycheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Bypasssafetycheck has changed for responderaction %s, starting update", responderactionName)
		responderaction.Bypasssafetycheck = d.Get("bypasssafetycheck").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for responderaction %s, starting update", responderactionName)
		responderaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("htmlpage") {
		log.Printf("[DEBUG]  citrixadc-provider: Htmlpage has changed for responderaction %s, starting update", responderactionName)
		responderaction.Htmlpage = d.Get("htmlpage").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for responderaction %s, starting update", responderactionName)
		responderaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("reasonphrase") {
		log.Printf("[DEBUG]  citrixadc-provider: Reasonphrase has changed for responderaction %s, starting update", responderactionName)
		responderaction.Reasonphrase = d.Get("reasonphrase").(string)
		hasChange = true
	}
	if d.HasChange("responsestatuscode") {
		log.Printf("[DEBUG]  citrixadc-provider: Responsestatuscode has changed for responderaction %s, starting update", responderactionName)
		responderaction.Responsestatuscode = d.Get("responsestatuscode").(int)
		hasChange = true
	}
	if d.HasChange("target") {
		log.Printf("[DEBUG]  citrixadc-provider: Target has changed for responderaction %s, starting update", responderactionName)
		responderaction.Target = d.Get("target").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for responderaction %s, starting update", responderactionName)
		responderaction.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Responderaction.Type(), responderactionName, &responderaction)
		if err != nil {
			return fmt.Errorf("Error updating responderaction %s", responderactionName)
		}
	}
	return readResponderactionFunc(d, meta)
}

func deleteResponderactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderactionFunc")
	client := meta.(*NetScalerNitroClient).client
	responderactionName := d.Id()
	err := client.DeleteResource(service.Responderaction.Type(), responderactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func customizeResponderactionDiff(diff *schema.ResourceDiff, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In customizeDiff")
	o := diff.GetChangedKeysPrefix("")

	// Check if target and bypasssafetycheck is in changed keys
	targetDefined := false
	bypasssafetycheckDefined := false

	for _, element := range o {

		if element == "target" {
			targetDefined = true
		}

		if element == "bypasssafetycheck" {
			bypasssafetycheckDefined = true
		}
	}

	// Clear bypasssafetycheck when present without target
	if bypasssafetycheckDefined && !targetDefined {
		diff.Clear("bypasssafetycheck")
	}

	return nil
}
