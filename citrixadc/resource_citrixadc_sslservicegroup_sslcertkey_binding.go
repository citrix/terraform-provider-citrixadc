package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslservicegroup_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslservicegroup_sslcertkey_bindingFunc,
		ReadContext:   readSslservicegroup_sslcertkey_bindingFunc,
		DeleteContext: deleteSslservicegroup_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ca": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certkeyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crlcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocspcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"snicert": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservicegroup_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservicegroup_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupname := d.Get("servicegroupname").(string)
	certkeyname := d.Get("certkeyname").(string)
	snicert := d.Get("snicert").(bool)
	ca := d.Get("ca").(bool)
	bindingId := fmt.Sprintf("%s,%s,%t,%t", servicegroupname, certkeyname, snicert, ca)
	sslservicegroup_sslcertkey_binding := ssl.Sslservicegroupsslcertkeybinding{
		Ca:               d.Get("ca").(bool),
		Certkeyname:      d.Get("certkeyname").(string),
		Crlcheck:         d.Get("crlcheck").(string),
		Ocspcheck:        d.Get("ocspcheck").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Snicert:          d.Get("snicert").(bool),
	}

	_, err := client.AddResource(service.Sslservicegroup_sslcertkey_binding.Type(), servicegroupname, &sslservicegroup_sslcertkey_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslservicegroup_sslcertkey_bindingFunc(ctx, d, meta)
}

func readSslservicegroup_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservicegroup_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	servicegroupname := idSlice[0]
	certkeyname := idSlice[1]
	snicert := false
	ca := false
	if len(idSlice) > 2 {
		snicert = idSlice[2] == "true"
		ca = idSlice[3] == "true"
	} else {
		snicert = d.Get("snicert").(bool)
		ca = d.Get("ca").(bool)
		bindingId = fmt.Sprintf("%s,%t,%t", bindingId, snicert, ca)
		d.SetId(bindingId)
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading sslservicegroup_sslcertkey_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslservicegroup_sslcertkey_binding",
		ResourceName:             servicegroupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["certkeyname"].(string) == certkeyname && v["snicert"].(bool) == snicert && v["ca"].(bool) == ca {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams sslcertkey not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ca", data["ca"])
	d.Set("certkeyname", data["certkeyname"])
	d.Set("crlcheck", data["crlcheck"])
	d.Set("ocspcheck", data["ocspcheck"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("snicert", data["snicert"])

	return nil

}

func deleteSslservicegroup_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservicegroup_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	certkeyname := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["certkeyname"] = url.QueryEscape(certkeyname)
	if v, ok := d.GetOk("ca"); ok && v.(bool) {
		argsMap["ca"] = url.QueryEscape(fmt.Sprintf("%v", v))
	}
	if v, ok := d.GetOk("snicert"); ok && v.(bool) {
		argsMap["snicert"] = url.QueryEscape(fmt.Sprintf("%v", v))
	}

	err := client.DeleteResourceWithArgsMap(service.Sslservicegroup_sslcertkey_binding.Type(), name, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
