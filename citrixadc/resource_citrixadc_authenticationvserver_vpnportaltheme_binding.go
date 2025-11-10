package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcAuthenticationvserver_vpnportaltheme_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationvserver_vpnportaltheme_bindingFunc,
		ReadContext:   readAuthenticationvserver_vpnportaltheme_bindingFunc,
		DeleteContext: deleteAuthenticationvserver_vpnportaltheme_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"portaltheme": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createAuthenticationvserver_vpnportaltheme_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationvserver_vpnportaltheme_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	portaltheme := d.Get("portaltheme")
	bindingId := fmt.Sprintf("%s,%s", name, portaltheme)
	authenticationvserver_vpnportaltheme_binding := authentication.Authenticationvservervpnportalthemebinding{
		Name:        d.Get("name").(string),
		Portaltheme: d.Get("portaltheme").(string),
	}

	err := client.UpdateUnnamedResource("authenticationvserver_vpnportaltheme_binding", &authenticationvserver_vpnportaltheme_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAuthenticationvserver_vpnportaltheme_bindingFunc(ctx, d, meta)
}

func readAuthenticationvserver_vpnportaltheme_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationvserver_vpnportaltheme_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	portaltheme := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationvserver_vpnportaltheme_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "authenticationvserver_vpnportaltheme_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationvserver_vpnportaltheme_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["portaltheme"].(string) == portaltheme {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationvserver_vpnportaltheme_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("portaltheme", data["portaltheme"])

	return nil

}

func deleteAuthenticationvserver_vpnportaltheme_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationvserver_vpnportaltheme_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	portaltheme := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("portaltheme:%s", portaltheme))

	err := client.DeleteResourceWithArgs("authenticationvserver_vpnportaltheme_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
