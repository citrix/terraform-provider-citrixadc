package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/rewrite"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRewriteglobal_rewritepolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewriteglobal_rewritepolicy_bindingFunc,
		Read:          readRewriteglobal_rewritepolicy_bindingFunc,
		Update:        updateRewriteglobal_rewritepolicy_bindingFunc,
		Delete:        deleteRewriteglobal_rewritepolicy_bindingFunc,
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

func createRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	var rewriteglobal_rewritepolicy_bindingName string
	if v, ok := d.GetOk("name"); ok {
		rewriteglobal_rewritepolicy_bindingName = v.(string)
	} else {
		rewriteglobal_rewritepolicy_bindingName = resource.PrefixedUniqueId("tf-rewriteglobal_rewritepolicy_binding-")
		d.Set("name", rewriteglobal_rewritepolicy_bindingName)
	}
	rewriteglobal_rewritepolicy_binding := rewrite.Rewriteglobal_rewritepolicy_binding{
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	_, err := client.AddResource(netscaler.Rewriteglobal_rewritepolicy_binding.Type(), rewriteglobal_rewritepolicy_bindingName, &rewriteglobal_rewritepolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(rewriteglobal_rewritepolicy_bindingName)

	err = readRewriteglobal_rewritepolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewriteglobal_rewritepolicy_binding but we can't read it ?? %s", rewriteglobal_rewritepolicy_bindingName)
		return nil
	}
	return nil
}

func readRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	rewriteglobal_rewritepolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rewriteglobal_rewritepolicy_binding state %s", rewriteglobal_rewritepolicy_bindingName)
	data, err := client.FindResource(netscaler.Rewriteglobal_rewritepolicy_binding.Type(), rewriteglobal_rewritepolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewriteglobal_rewritepolicy_binding state %s", rewriteglobal_rewritepolicy_bindingName)
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

func updateRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	rewriteglobal_rewritepolicy_bindingName := d.Get("name").(string)

	rewriteglobal_rewritepolicy_binding := rewrite.Rewriteglobal_rewritepolicy_binding{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("globalbindtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Globalbindtype has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Globalbindtype = d.Get("globalbindtype").(string)
		hasChange = true
	}
	if d.HasChange("gotopriorityexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Gotopriorityexpression has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Gotopriorityexpression = d.Get("gotopriorityexpression").(string)
		hasChange = true
	}
	if d.HasChange("invoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Invoke has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Invoke = d.Get("invoke").(bool)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("labeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Labeltype has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Labeltype = d.Get("labeltype").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for rewriteglobal_rewritepolicy_binding %s, starting update", rewriteglobal_rewritepolicy_bindingName)
		rewriteglobal_rewritepolicy_binding.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Rewriteglobal_rewritepolicy_binding.Type(), rewriteglobal_rewritepolicy_bindingName, &rewriteglobal_rewritepolicy_binding)
		if err != nil {
			return fmt.Errorf("Error updating rewriteglobal_rewritepolicy_binding %s", rewriteglobal_rewritepolicy_bindingName)
		}
	}
	return readRewriteglobal_rewritepolicy_bindingFunc(d, meta)
}

func deleteRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	rewriteglobal_rewritepolicy_bindingName := d.Id()
	err := client.DeleteResource(netscaler.Rewriteglobal_rewritepolicy_binding.Type(), rewriteglobal_rewritepolicy_bindingName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
