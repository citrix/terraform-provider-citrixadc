package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/cs"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCsvserver_responderpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCsvserver_responderpolicy_bindingFunc,
		Read:          readCsvserver_responderpolicy_bindingFunc,
		Update:        updateCsvserver_responderpolicy_bindingFunc,
		Delete:        deleteCsvserver_responderpolicy_bindingFunc,
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

func createCsvserver_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserver_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	var csvserver_responderpolicy_bindingName string
	if v, ok := d.GetOk("name"); ok {
		csvserver_responderpolicy_bindingName = v.(string)
	} else {
		csvserver_responderpolicy_bindingName = resource.PrefixedUniqueId("tf-csvserver_responderpolicy_binding-")
		d.Set("name", csvserver_responderpolicy_bindingName)
	}
	csvserver_responderpolicy_binding := cs.Csvserver_responderpolicy_binding{
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

	_, err := client.AddResource(netscaler.Csvserver_responderpolicy_binding.Type(), csvserver_responderpolicy_bindingName, &csvserver_responderpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(csvserver_responderpolicy_bindingName)

	err = readCsvserver_responderpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this csvserver_responderpolicy_binding but we can't read it ?? %s", csvserver_responderpolicy_bindingName)
		return nil
	}
	return nil
}

func readCsvserver_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserver_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_responderpolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver_responderpolicy_binding state %s", csvserver_responderpolicy_bindingName)
	data, err := client.FindResource(netscaler.Csvserver_responderpolicy_binding.Type(), csvserver_responderpolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_responderpolicy_binding state %s", csvserver_responderpolicy_bindingName)
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

func updateCsvserver_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCsvserver_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_responderpolicy_bindingName := d.Get("name").(string)

	csvserver_responderpolicy_binding := cs.Csvserver_responderpolicy_binding{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bindpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Bindpoint has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Bindpoint = d.Get("bindpoint").(string)
		hasChange = true
	}
	if d.HasChange("gotopriorityexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Gotopriorityexpression has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Gotopriorityexpression = d.Get("gotopriorityexpression").(string)
		hasChange = true
	}
	if d.HasChange("invoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Invoke has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Invoke = d.Get("invoke").(bool)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("labeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Labeltype has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Labeltype = d.Get("labeltype").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("targetlbvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Targetlbvserver has changed for csvserver_responderpolicy_binding %s, starting update", csvserver_responderpolicy_bindingName)
		csvserver_responderpolicy_binding.Targetlbvserver = d.Get("targetlbvserver").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Csvserver_responderpolicy_binding.Type(), csvserver_responderpolicy_bindingName, &csvserver_responderpolicy_binding)
		if err != nil {
			return fmt.Errorf("Error updating csvserver_responderpolicy_binding %s", csvserver_responderpolicy_bindingName)
		}
	}
	return readCsvserver_responderpolicy_bindingFunc(d, meta)
}

func deleteCsvserver_responderpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserver_responderpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_responderpolicy_bindingName := d.Id()
	err := client.DeleteResource(netscaler.Csvserver_responderpolicy_binding.Type(), csvserver_responderpolicy_bindingName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
