package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/lb"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbvserver_rewritepolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbvserver_rewritepolicy_bindingFunc,
		Read:          readLbvserver_rewritepolicy_bindingFunc,
		Update:        updateLbvserver_rewritepolicy_bindingFunc,
		Delete:        deleteLbvserver_rewritepolicy_bindingFunc,
		Schema: map[string]*schema.Schema{
			"bindpoint": &schema.Schema{
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
			"name": &schema.Schema{
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
		},
	}
}

func createLbvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	var lbvserver_rewritepolicy_bindingName string
	if v, ok := d.GetOk("name"); ok {
		lbvserver_rewritepolicy_bindingName = v.(string)
	} else {
		lbvserver_rewritepolicy_bindingName = resource.PrefixedUniqueId("tf-lbvserver_rewritepolicy_binding-")
		d.Set("name", lbvserver_rewritepolicy_bindingName)
	}
	lbvserver_rewritepolicy_binding := lb.Lbvserver_rewritepolicy_binding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   d.Get("name").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	_, err := client.AddResource(netscaler.Lbvserver_rewritepolicy_binding.Type(), lbvserver_rewritepolicy_bindingName, &lbvserver_rewritepolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(lbvserver_rewritepolicy_bindingName)

	err = readLbvserver_rewritepolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbvserver_rewritepolicy_binding but we can't read it ?? %s", lbvserver_rewritepolicy_bindingName)
		return nil
	}
	return nil
}

func readLbvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserver_rewritepolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_rewritepolicy_binding state %s", lbvserver_rewritepolicy_bindingName)
	data, err := client.FindResource(netscaler.Lbvserver_rewritepolicy_binding.Type(), lbvserver_rewritepolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_rewritepolicy_binding state %s", lbvserver_rewritepolicy_bindingName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("bindpoint", data["bindpoint"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("name", data["name"])
	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])

	return nil

}

func updateLbvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserver_rewritepolicy_bindingName := d.Get("name").(string)

	lbvserver_rewritepolicy_binding := lb.Lbvserver_rewritepolicy_binding{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bindpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Bindpoint has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Bindpoint = d.Get("bindpoint").(string)
		hasChange = true
	}
	if d.HasChange("gotopriorityexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Gotopriorityexpression has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Gotopriorityexpression = d.Get("gotopriorityexpression").(string)
		hasChange = true
	}
	if d.HasChange("invoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Invoke has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Invoke = d.Get("invoke").(bool)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("labeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Labeltype has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Labeltype = d.Get("labeltype").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for lbvserver_rewritepolicy_binding %s, starting update", lbvserver_rewritepolicy_bindingName)
		lbvserver_rewritepolicy_binding.Priority = d.Get("priority").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Lbvserver_rewritepolicy_binding.Type(), lbvserver_rewritepolicy_bindingName, &lbvserver_rewritepolicy_binding)
		if err != nil {
			return fmt.Errorf("Error updating lbvserver_rewritepolicy_binding %s", lbvserver_rewritepolicy_bindingName)
		}
	}
	return readLbvserver_rewritepolicy_bindingFunc(d, meta)
}

func deleteLbvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserver_rewritepolicy_bindingName := d.Id()
	err := client.DeleteResource(netscaler.Lbvserver_rewritepolicy_binding.Type(), lbvserver_rewritepolicy_bindingName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
