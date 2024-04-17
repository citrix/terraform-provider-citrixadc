package citrixadc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbactionFunc,
		Read:          readLbactionFunc,
		Update:        updateLbactionFunc,
		Delete:        deleteLbactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"value": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbactionFunc")
	client := meta.(*NetScalerNitroClient).client

	lbactionName := d.Get("name").(string)

	lbaction := lbaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Newname: d.Get("newname").(string),
		Type:    d.Get("type").(string),
		Value:   toIntegerList(d.Get("value").([]interface{})),
	}

	_, err := client.AddResource("lbaction", lbactionName, &lbaction)
	if err != nil {
		return err
	}

	d.SetId(lbactionName)

	err = readLbactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbaction but we can't read it ?? %s", lbactionName)
		return nil
	}
	return nil
}

func readLbactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbactionFunc")
	client := meta.(*NetScalerNitroClient).client
	lbactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbaction state %s", lbactionName)
	data, err := client.FindResource("lbaction", lbactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbaction state %s", lbactionName)
		d.SetId("")
		return nil
	}

	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("newname", data["newname"])
	d.Set("type", data["type"])
	d.Set("value", stringListToIntList(data["value"].([]interface{})))

	return nil

}

func updateLbactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbactionFunc")
	client := meta.(*NetScalerNitroClient).client
	lbactionName := d.Get("name").(string)

	lbaction := lbaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for lbaction %s, starting update", lbactionName)
		lbaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for lbaction %s, starting update", lbactionName)
		lbaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for lbaction %s, starting update", lbactionName)
		lbaction.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for lbaction %s, starting update", lbactionName)
		lbaction.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("value") {
		log.Printf("[DEBUG]  citrixadc-provider: Value has changed for lbaction %s, starting update", lbactionName)
		lbaction.Value = toIntegerList(d.Get("value").([]interface{}))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lbaction", lbactionName, &lbaction)
		if err != nil {
			return fmt.Errorf("Error updating lbaction %s", lbactionName)
		}
	}
	return readLbactionFunc(d, meta)
}

func deleteLbactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbactionFunc")
	client := meta.(*NetScalerNitroClient).client
	lbactionName := d.Id()
	err := client.DeleteResource("lbaction", lbactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

/**
* Configuration for lb action resource.
 */
type lbaction struct {
	/**
	* Name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

	* The following requirement applies only to the Citrix ADC CLI:
	* If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb action" or 'my lb action').
	 */
	Name string `json:"name,omitempty"`
	/**
	* Type of an LB action. Available settings function as follows:
	* NOLBACTION - Does not consider LB action in making LB decision.
	* SELECTIONORDER - services bound to vserver with order specified in value parameter is considerd for lb/gslb decision.
	Possible values = NOLBACTION, SELECTIONORDER
	*/
	Type string `json:"type,omitempty"`
	/**
	* The selection order list used during lb/gslb decision. Preference of services during lb/gslb decision is as follows - services corresponding to first order specified in the sequence is considered first, services corresponding to second order specified in the sequence is considered next and so on.
	* For example, if -value 2 1 3 is specified here and service-1 bound to a vserver with order 1, service-2 bound to a vserver with order 2 and service-3 bound to a vserver with order 3. Then preference of selecting services in LB decision is as follows: service-2, service-1, service-3.
	* Minimum value = 1
	 */
	Value []int `json:"value,omitempty"`
	/**
	* Comment. Any type of information about this LB action.
	 */
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the load balancing virtual server group.
	 */
	Newname string `json:"newname,omitempty"`
}
