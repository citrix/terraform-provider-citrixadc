package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/responder"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcResponderglobal_responderpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderglobal_responderpolicy_bindingFunc,
		Read:          readResponderglobal_responderpolicy_bindingFunc,
		Update:        updateResponderglobal_responderpolicy_bindingFunc,
		Delete:        deleteResponderglobal_responderpolicy_bindingFunc,
		Schema: map[string]*schema.Schema{
			"globalbindtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
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

func createResponderglobal_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderglobal_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	var responderglobal_responderpolicy_bindingName string
	if v, ok := d.GetOk("name"); ok {
		responderglobal_responderpolicy_bindingName = v.(string)
	} else {
		responderglobal_responderpolicy_bindingName = resource.PrefixedUniqueId("tf-responderglobal_responderpolicy_binding-")
		d.Set("name", responderglobal_responderpolicy_bindingName)
	}
	responderglobal_responderpolicy_binding := responder.Responderglobal_responderpolicy_binding{
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	_, err := client.AddResource(netscaler.Responderglobal_responderpolicy_binding.Type(), responderglobal_responderpolicy_bindingName, &responderglobal_responderpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(responderglobal_responderpolicy_bindingName)

	err = readResponderglobal_responderpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderglobal_responderpolicy_binding but we can't read it ?? %s", responderglobal_responderpolicy_bindingName)
		return nil
	}
	return nil
}

func readResponderglobal_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderglobal_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	responderglobal_responderpolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderglobal_responderpolicy_binding state %s", responderglobal_responderpolicy_bindingName)
	data, err := client.FindResource(netscaler.Responderglobal_responderpolicy_binding.Type(), responderglobal_responderpolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderglobal_responderpolicy_binding state %s", responderglobal_responderpolicy_bindingName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])
	d.Set("type", data["type"])

	return nil

}

func updateResponderglobal_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateResponderglobal_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	responderglobal_responderpolicy_bindingName := d.Get("name").(string)

	responderglobal_responderpolicy_binding := responder.Responderglobal_responderpolicy_binding{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("globalbindtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Globalbindtype has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Globalbindtype = d.Get("globalbindtype").(string)
		hasChange = true
	}
	if d.HasChange("gotopriorityexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Gotopriorityexpression has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Gotopriorityexpression = d.Get("gotopriorityexpression").(string)
		hasChange = true
	}
	if d.HasChange("invoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Invoke has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Invoke = d.Get("invoke").(bool)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("labeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Labeltype has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Labeltype = d.Get("labeltype").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for responderglobal_responderpolicy_binding %s, starting update", responderglobal_responderpolicy_bindingName)
		responderglobal_responderpolicy_binding.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Responderglobal_responderpolicy_binding.Type(), responderglobal_responderpolicy_bindingName, &responderglobal_responderpolicy_binding)
		if err != nil {
			return fmt.Errorf("Error updating responderglobal_responderpolicy_binding %s", responderglobal_responderpolicy_bindingName)
		}
	}
	return readResponderglobal_responderpolicy_bindingFunc(d, meta)
}

func deleteResponderglobal_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderglobal_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	responderglobal_responderpolicy_bindingName := d.Id()
	err := client.DeleteResource(netscaler.Responderglobal_responderpolicy_binding.Type(), responderglobal_responderpolicy_bindingName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
