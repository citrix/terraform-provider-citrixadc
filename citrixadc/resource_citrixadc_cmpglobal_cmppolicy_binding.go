package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"net/url"
)

func resourceCitrixAdcCmpglobal_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCmpglobal_bindingFunc,
		Read:          readCmpglobal_bindingFunc,
		Delete:        deleteCmpglobal_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createCmpglobal_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmpglobal_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)

	Cmpglobal_binding := cmp.Cmpglobalpolicybinding{
		Policyname: d.Get("policyname").(string),
		Priority:   uint32(d.Get("priority").(int)),
		State:      d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource("cmpglobal_binding", &Cmpglobal_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readCmpglobal_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this Cmpglobal_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readCmpglobal_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmpglobal_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading Cmpglobal_binding %s", policyname)

	data, err := client.FindResource("cmppolicy_cmpglobal_binding", policyname)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	d.Set("policyname", data["name"])
	d.Set("priority", data["priority"])
	d.Set("state", data["state"])

	return nil

}

func deleteCmpglobal_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmpglobal_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	argsMap := make(map[string]string)
	argsMap["policyname"] = url.QueryEscape(policyname)

	err := client.DeleteResourceWithArgsMap("cmpglobal_cmppolicy_binding", "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
