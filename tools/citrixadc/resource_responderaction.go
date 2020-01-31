package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/responder"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcResponderaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderactionFunc,
		Read:          readResponderactionFunc,
		Update:        updateResponderactionFunc,
		Delete:        deleteResponderactionFunc,
		Schema: map[string]*schema.Schema{
			"bypasssafetycheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htmlpage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reasonphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"responsestatuscode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
		Newname:            d.Get("newname").(string),
		Reasonphrase:       d.Get("reasonphrase").(string),
		Responsestatuscode: d.Get("responsestatuscode").(int),
		Target:             d.Get("target").(string),
		Type:               d.Get("type").(string),
	}

	_, err := client.AddResource(netscaler.Responderaction.Type(), responderactionName, &responderaction)
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
	data, err := client.FindResource(netscaler.Responderaction.Type(), responderactionName)
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
	d.Set("newname", data["newname"])
	d.Set("reasonphrase", data["reasonphrase"])
	d.Set("responsestatuscode", data["responsestatuscode"])
	d.Set("target", data["target"])
	d.Set("type", data["type"])

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
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for responderaction %s, starting update", responderactionName)
		responderaction.Newname = d.Get("newname").(string)
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
		_, err := client.UpdateResource(netscaler.Responderaction.Type(), responderactionName, &responderaction)
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
	err := client.DeleteResource(netscaler.Responderaction.Type(), responderactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
