package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsnclient_nsacl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnclient_nsacl_bindingFunc,
		Read:          readLsnclient_nsacl_bindingFunc,
		Delete:        deleteLsnclient_nsacl_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aclname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"clientname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnclient_nsacl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclient_nsacl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	clientname := d.Get("clientname")
	aclname := d.Get("aclname")
	bindingId := fmt.Sprintf("%s,%s", clientname, aclname)
	lsnclient_nsacl_binding := lsn.Lsnclientnsaclbinding{
		Aclname:    d.Get("aclname").(string),
		Clientname: d.Get("clientname").(string),
		Td:         d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource("lsnclient_nsacl_binding", &lsnclient_nsacl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsnclient_nsacl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnclient_nsacl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsnclient_nsacl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnclient_nsacl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	clientname := idSlice[0]
	aclname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsnclient_nsacl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsnclient_nsacl_binding",
		ResourceName:             clientname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient_nsacl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["aclname"].(string) == aclname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams aclname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient_nsacl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("aclname", data["aclname"])
	d.Set("clientname", data["clientname"])
	d.Set("td", data["td"])

	return nil

}

func deleteLsnclient_nsacl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnclient_nsacl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	aclname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("aclname:%s", aclname))
	if v, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%v", v.(int)))

	}

	err := client.DeleteResourceWithArgs("lsnclient_nsacl_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
