package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/cs"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcCsvserver_appfwpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCsvserver_appfwpolicy_bindingFunc,
		Read:          readCsvserver_appfwpolicy_bindingFunc,
		Delete:        deleteCsvserver_appfwpolicy_bindingFunc,
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
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCsvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_appfwpolicy_bindingName := d.Get("name").(string)
	csvserver_appfwpolicy_binding := cs.Csvserverappfwpolicybinding{
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

	_, err := client.AddResource(netscaler.Csvserver_appfwpolicy_binding.Type(), csvserver_appfwpolicy_bindingName, &csvserver_appfwpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(csvserver_appfwpolicy_bindingName)

	err = readCsvserver_appfwpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this csvserver_appfwpolicy_binding but we can't read it ?? %s", csvserver_appfwpolicy_bindingName)
		return nil
	}
	return nil
}

func readCsvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserver_appfwpolicy_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver_appfwpolicy_binding state %s", csvserver_appfwpolicy_bindingName)
	data, err := client.FindResource(netscaler.Csvserver_appfwpolicy_binding.Type(), csvserver_appfwpolicy_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_appfwpolicy_binding state %s", csvserver_appfwpolicy_bindingName)
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

func deleteCsvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make(map[string]string)
	args["policyname"] = d.Get("policyname").(string)
	err := client.DeleteResourceWithArgsMap(netscaler.Csvserver_appfwpolicy_binding.Type(), d.Get("name").(string), args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
