package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslvserver_sslpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslvserver_sslpolicy_bindingFunc,
		Read:          readSslvserver_sslpolicy_bindingFunc,
		Delete:        deleteSslvserver_sslpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
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
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vservername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslvserver_sslpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslvserver_sslpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vservername := d.Get("vservername")
	policyname := d.Get("policyname")

	// Use `,` as the separator since it is invalid character for adc entity strings
	bindingId := fmt.Sprintf("%s,%s", vservername, policyname)

	sslvserver_sslpolicy_binding := ssl.Sslvserverpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               uint32(d.Get("priority").(int)),
		Type:                   d.Get("type").(string),
		Vservername:            d.Get("vservername").(string),
	}

	err := client.UpdateUnnamedResource(service.Sslvserver_sslpolicy_binding.Type(), &sslvserver_sslpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslvserver_sslpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslvserver_sslpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslvserver_sslpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslvserver_sslpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	if len(idSlice) < 2 {
		return fmt.Errorf("Cannot deduce policyname from id string")
	}

	if len(idSlice) > 2 {
		return fmt.Errorf("Too many separators \",\" in id string")
	}

	vservername := idSlice[0]
	policyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslvserver_sslpolicy_binding state %s", bindingId)
	findParams := service.FindParams{
		ResourceType:             "sslvserver_sslpolicy_binding",
		ResourceName:             vservername,
		ResourceMissingErrorCode: 461,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_sslpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_sslpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])
	d.Set("vservername", data["vservername"])

	return nil

}

func deleteSslvserver_sslpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslvserver_sslpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	vservername := idSlice[0]
	policyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))

	if v, ok := d.GetOk("priority"); ok {
		args = append(args, fmt.Sprintf("priority:%v", v))
	}

	if v, ok := d.GetOk("type"); ok {
		args = append(args, fmt.Sprintf("type:%v", v))
	}

	err := client.DeleteResourceWithArgs(service.Sslvserver_sslpolicy_binding.Type(), vservername, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
