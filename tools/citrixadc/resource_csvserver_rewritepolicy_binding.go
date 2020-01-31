package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/cs"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCsvserver_rewritepolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCsvserver_rewritepolicy_bindingFunc,
		Read:          readCsvserver_rewritepolicy_bindingFunc,
		Update:        updateCsvserver_rewritepolicy_bindingFunc,
		Delete:        deleteCsvserver_rewritepolicy_bindingFunc,
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
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCsvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	var csvserver_rewritepolicy_bindingName string
	if v, ok := d.GetOk("name"); ok {
		csvserver_rewritepolicy_bindingName = v.(string)
	} else {
		csvserver_rewritepolicy_bindingName = resource.PrefixedUniqueId("tf-csvserver_rewritepolicy_binding-")
		d.Set("name", csvserver_rewritepolicy_bindingName)
	}
	csvserver_rewritepolicy_binding := cs.Csvserver_rewritepolicy_binding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   d.Get("name").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Targetlbvserver:        d.Get("targetlbvserver").(string),
	}

	_, err := client.AddResource(netscaler.Csvserver_rewritepolicy_binding.Type(), csvserver_rewritepolicy_bindingName, &csvserver_rewritepolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(csvserver_rewritepolicy_bindingName)

	err = readCsvserver_rewritepolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this csvserver_rewritepolicy_binding but we can't read it ?? %s", csvserver_rewritepolicy_bindingName)
		return nil
	}
	return nil
}

func readCsvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_rewritepolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver_rewritepolicy_binding state %s", csvserver_rewritepolicy_bindingName)
	data, err := client.FindResource(netscaler.Csvserver_rewritepolicy_binding.Type(), csvserver_rewritepolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_rewritepolicy_binding state %s", csvserver_rewritepolicy_bindingName)
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
	d.Set("targetlbvserver", data["targetlbvserver"])

	return nil

}

func updateCsvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCsvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_rewritepolicy_bindingName := d.Get("name").(string)

	csvserver_rewritepolicy_binding := cs.Csvserver_rewritepolicy_binding{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bindpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Bindpoint has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Bindpoint = d.Get("bindpoint").(string)
		hasChange = true
	}
	if d.HasChange("gotopriorityexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Gotopriorityexpression has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Gotopriorityexpression = d.Get("gotopriorityexpression").(string)
		hasChange = true
	}
	if d.HasChange("invoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Invoke has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Invoke = d.Get("invoke").(bool)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("labeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Labeltype has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Labeltype = d.Get("labeltype").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("targetlbvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Targetlbvserver has changed for csvserver_rewritepolicy_binding %s, starting update", csvserver_rewritepolicy_bindingName)
		csvserver_rewritepolicy_binding.Targetlbvserver = d.Get("targetlbvserver").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Csvserver_rewritepolicy_binding.Type(), csvserver_rewritepolicy_bindingName, &csvserver_rewritepolicy_binding)
		if err != nil {
			return fmt.Errorf("Error updating csvserver_rewritepolicy_binding %s", csvserver_rewritepolicy_bindingName)
		}
	}
	return readCsvserver_rewritepolicy_bindingFunc(d, meta)
}

func deleteCsvserver_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_rewritepolicy_bindingName := d.Id()
	err := client.DeleteResource(netscaler.Csvserver_rewritepolicy_binding.Type(), csvserver_rewritepolicy_bindingName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
