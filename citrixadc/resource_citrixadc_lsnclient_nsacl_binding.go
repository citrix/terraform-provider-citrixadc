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

func resourceCitrixAdcLsnclient_nsacl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnclient_nsacl_bindingFunc,
		ReadContext:   readLsnclient_nsacl_bindingFunc,
		DeleteContext: deleteLsnclient_nsacl_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"aclname": {
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

func createLsnclient_nsacl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclient_nsacl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	clientname := d.Get("clientname")
	aclname := d.Get("aclname")
	bindingId := fmt.Sprintf("%s,%s", clientname, aclname)
	lsnclient_nsacl_binding := lsn.Lsnclientnsaclbinding{
		Aclname:    d.Get("aclname").(string),
		Clientname: d.Get("clientname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		lsnclient_nsacl_binding.Td = intPtr(d.Get("td").(int))
	}

	err := client.UpdateUnnamedResource("lsnclient_nsacl_binding", &lsnclient_nsacl_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLsnclient_nsacl_bindingFunc(ctx, d, meta)
}

func readLsnclient_nsacl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
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
	setToInt("td", d, data["td"])

	return nil

}

func deleteLsnclient_nsacl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
