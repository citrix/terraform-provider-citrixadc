package citrixadc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbpolicyFunc,
		Read:          readLbpolicyFunc,
		Update:        updateLbpolicyFunc,
		Delete:        deleteLbpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client

	lbpolicyName := d.Get("name").(string)

	lbpolicy := lbpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("lbpolicy", lbpolicyName, &lbpolicy)
	if err != nil {
		return err
	}

	d.SetId(lbpolicyName)

	err = readLbpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbpolicy but we can't read it ?? %s", lbpolicyName)
		return nil
	}
	return nil
}

func readLbpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	lbpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbpolicy state %s", lbpolicyName)
	data, err := client.FindResource("lbpolicy", lbpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbpolicy state %s", lbpolicyName)
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

func updateLbpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	lbpolicyName := d.Get("name").(string)

	lbpolicy := lbpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lbpolicy", lbpolicyName, &lbpolicy)
		if err != nil {
			return fmt.Errorf("Error updating lbpolicy %s", lbpolicyName)
		}
	}
	return readLbpolicyFunc(d, meta)
}

func deleteLbpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	lbpolicyName := d.Id()
	err := client.DeleteResource("lbpolicy", lbpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

/**
* Configuration for lb policy resource.
 */
type lbpolicy struct {
	/**
	* Name of the LB policy.
	* Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added.
	* The following requirement applies only to the Citrix ADC CLI:
	* If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy" or 'my lb policy').
	 */
	Name string `json:"name,omitempty"`
	/**
	* Expression against which traffic is evaluated.
	 */
	Rule string `json:"rule,omitempty"`
	/**
	* Name of action to use if the request matches this LB policy.
	 */
	Action string `json:"action,omitempty"`
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows:
	* NOLBACTION - Does not consider LB actions in making LB decision.
	* RESET - Reset the request and notify the user, so that the user can resend the request.
	* DROP - Drop the request without sending a response to the user.
	 */
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Name of the messagelog action to use for requests that match this policy.
	 */
	Logaction string `json:"logaction,omitempty"`
	/**
	* Any type of information about this LB policy.
	 */
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	* The following requirement applies only to the Citrix ADC CLI:
	* If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy" or 'my lb policy').
	* Minimum length = 1
	 */
	Newname string `json:"newname,omitempty"`
}
