package citrixadc

import (
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

type Csvserverappfwpolicybinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

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
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createCsvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Get("name").(string)
	appfwPolicyName := d.Get("policyname").(string)
	bindingId := fmt.Sprintf("%s,%s", csvserverName, appfwPolicyName)

	csvserver_appfwpolicy_binding := Csvserverappfwpolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   csvserverName,
		Policyname:             appfwPolicyName,
		Priority:               d.Get("priority").(int),
		Targetlbvserver:        d.Get("targetlbvserver").(string),
	}

	_, err := client.AddResource(service.Csvserver_appfwpolicy_binding.Type(), csvserverName, &csvserver_appfwpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readCsvserver_appfwpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this csvserver_appfwpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readCsvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readCsvserver_appfwpolicy_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.SplitN(bindingId, ",", 2)
	csvserverName := idSlice[0]
	appfwPolicyName := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver_appfwpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             service.Csvserver_appfwpolicy_binding.Type(),
		ResourceName:             csvserverName,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_appfwpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == appfwPolicyName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_appfwpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

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

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	csvserverName := idSlice[0]
	appfwPolicyName := idSlice[1]

	args := make(map[string]string)
	args["policyname"] = appfwPolicyName
	err := client.DeleteResourceWithArgsMap(service.Csvserver_appfwpolicy_binding.Type(), csvserverName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
