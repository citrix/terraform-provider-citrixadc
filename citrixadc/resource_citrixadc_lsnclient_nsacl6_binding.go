package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsnclient_nsacl6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnclient_nsacl6_bindingFunc,
		Read:          readLsnclient_nsacl6_bindingFunc,
		Delete:        deleteLsnclient_nsacl6_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"acl6name": &schema.Schema{
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

func createLsnclient_nsacl6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclient_nsacl6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	clientname := d.Get("clientname")
	acl6name := d.Get("acl6name")
	bindingId := fmt.Sprintf("%s,%s", clientname, acl6name)
	lsnclient_nsacl6_binding := lsn.Lsnclientnsacl6binding{
		Acl6name:   d.Get("acl6name").(string),
		Clientname: d.Get("clientname").(string),
		Td:         d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource("lsnclient_nsacl6_binding", &lsnclient_nsacl6_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsnclient_nsacl6_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnclient_nsacl6_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsnclient_nsacl6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnclient_nsacl6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	clientname := idSlice[0]
	acl6name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsnclient_nsacl6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsnclient_nsacl6_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient_nsacl6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["acl6name"].(string) == acl6name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams acl6name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient_nsacl6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("acl6name", data["acl6name"])
	d.Set("clientname", data["clientname"])
	d.Set("td", data["td"])

	return nil

}

func deleteLsnclient_nsacl6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnclient_nsacl6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	acl6name := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("acl6name:%s", acl6name))
	if v, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%v", v.(int)))

	}

	err := client.DeleteResourceWithArgs("lsnclient_nsacl6_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
