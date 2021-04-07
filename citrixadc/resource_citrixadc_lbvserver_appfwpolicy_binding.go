package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/lb"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbvserver_appfwpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbvserver_appfwpolicy_bindingFunc,
		Read:          readLbvserver_appfwpolicy_bindingFunc,
		Update:        updateLbvserver_appfwpolicy_bindingFunc,
		Delete:        deleteLbvserver_appfwpolicy_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserver_appfwpolicy_bindingName := d.Get("name").(string)
	lbvserver_appfwpolicy_binding := lb.Lbvserverappfwpolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   d.Get("name").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	_, err := client.AddResource(netscaler.Lbvserver_appfwpolicy_binding.Type(), lbvserver_appfwpolicy_bindingName, &lbvserver_appfwpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(lbvserver_appfwpolicy_bindingName)

	err = readLbvserver_appfwpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbvserver_appfwpolicy_binding but we can't read it ?? %s", lbvserver_appfwpolicy_bindingName)
		return nil
	}
	return nil
}

func readLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserver_appfwpolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_appfwpolicy_binding state %s", lbvserver_appfwpolicy_bindingName)
	data, err := client.FindResource(netscaler.Lbvserver_appfwpolicy_binding.Type(), lbvserver_appfwpolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_appfwpolicy_binding state %s", lbvserver_appfwpolicy_bindingName)
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

func updateLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserver_appfwpolicy_bindingName := d.Get("name").(string)

	lbvserver_appfwpolicy_binding := lb.Lbvserverappfwpolicybinding{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bindpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Bindpoint has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Bindpoint = d.Get("bindpoint").(string)
		hasChange = true
	}
	if d.HasChange("gotopriorityexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Gotopriorityexpression has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Gotopriorityexpression = d.Get("gotopriorityexpression").(string)
		hasChange = true
	}
	if d.HasChange("invoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Invoke has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Invoke = d.Get("invoke").(bool)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("labeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Labeltype has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Labeltype = d.Get("labeltype").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for lbvserver_appfwpolicy_binding %s, starting update", lbvserver_appfwpolicy_bindingName)
		lbvserver_appfwpolicy_binding.Priority = d.Get("priority").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Lbvserver_appfwpolicy_binding.Type(), lbvserver_appfwpolicy_bindingName, &lbvserver_appfwpolicy_binding)
		if err != nil {
			return fmt.Errorf("Error updating lbvserver_appfwpolicy_binding %s", lbvserver_appfwpolicy_bindingName)
		}
	}
	return readLbvserver_appfwpolicy_bindingFunc(d, meta)
}

func deleteLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make(map[string]string)
	args["policyname"] = d.Get("policyname").(string)
	err := client.DeleteResourceWithArgsMap(netscaler.Lbvserver_appfwpolicy_binding.Type(), d.Get("name").(string), args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
