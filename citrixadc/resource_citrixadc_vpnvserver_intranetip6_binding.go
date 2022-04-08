package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcVpnvserver_intranetip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_intranetip6_bindingFunc,
		Read:          readVpnvserver_intranetip6_bindingFunc,
		Delete:        deleteVpnvserver_intranetip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"intranetip6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"numaddr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	intranetip6 := d.Get("intranetip6")
	bindingId := fmt.Sprintf("%s,%s", name, intranetip6)
	vpnvserver_intranetip6_binding := vpn.Vpnvserverintranetip6binding{
		Intranetip6: d.Get("intranetip6").(string),
		Name:        d.Get("name").(string),
		Numaddr:     d.Get("numaddr").(int),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_intranetip6_binding.Type(), &vpnvserver_intranetip6_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_intranetip6_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_intranetip6_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip6 := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_intranetip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_intranetip6_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_intranetip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip6"].(string) == intranetip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_intranetip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("intranetip6", data["intranetip6"])
	d.Set("name", data["name"])
	d.Set("numaddr", data["numaddr"])

	return nil

}

func deleteVpnvserver_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip6 := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip6:%s", intranetip6))
	if val, ok := d.GetOk("numaddr"); ok {
		args = append(args, fmt.Sprintf("numaddr:%d", (val.(int))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnvserver_intranetip6_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
