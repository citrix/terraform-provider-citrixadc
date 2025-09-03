package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"
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
				ForceNew: true,
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
				ForceNew: true,
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

	lbaction := lb.Lbaction{
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

	lbaction := lb.Lbaction{
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
