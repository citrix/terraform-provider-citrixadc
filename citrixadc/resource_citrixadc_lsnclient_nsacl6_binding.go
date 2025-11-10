package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcLsnclient_nsacl6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnclient_nsacl6_bindingFunc,
		ReadContext:   readLsnclient_nsacl6_bindingFunc,
		DeleteContext: deleteLsnclient_nsacl6_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"acl6name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"clientname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnclient_nsacl6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclient_nsacl6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	clientname := d.Get("clientname")
	acl6name := d.Get("acl6name")
	bindingId := fmt.Sprintf("%s,%s", clientname, acl6name)
	lsnclient_nsacl6_binding := lsn.Lsnclientnsacl6binding{
		Acl6name:   d.Get("acl6name").(string),
		Clientname: d.Get("clientname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		lsnclient_nsacl6_binding.Td = intPtr(d.Get("td").(int))
	}

	err := client.UpdateUnnamedResource("lsnclient_nsacl6_binding", &lsnclient_nsacl6_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLsnclient_nsacl6_bindingFunc(ctx, d, meta)
}

func readLsnclient_nsacl6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
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
	setToInt("td", d, data["td"])

	return nil

}

func deleteLsnclient_nsacl6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
