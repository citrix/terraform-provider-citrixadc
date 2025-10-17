package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCacheforwardproxy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCacheforwardproxyFunc,
		ReadContext:   readCacheforwardproxyFunc,
		DeleteContext: deleteCacheforwardproxyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCacheforwardproxyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCacheforwardproxyFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheforwardproxyName := d.Get("ipaddress").(string)
	cacheforwardproxyid := fmt.Sprintf("%s,%s", cacheforwardproxyName, strconv.Itoa(d.Get("port").(int)))
	cacheforwardproxy := cache.Cacheforwardproxy{
		Ipaddress: d.Get("ipaddress").(string),
	}

	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		cacheforwardproxy.Port = intPtr(d.Get("port").(int))
	}

	_, err := client.AddResource(service.Cacheforwardproxy.Type(), cacheforwardproxyName, &cacheforwardproxy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cacheforwardproxyid)

	return readCacheforwardproxyFunc(ctx, d, meta)
}

func readCacheforwardproxyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCacheforwardproxyFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheforwardproxyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cacheforwardproxy state %s", cacheforwardproxyName)
	dataArr, err := client.FindAllResources(service.Cacheforwardproxy.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cacheforwardproxy state %s", cacheforwardproxyName)
		d.SetId("")
		return nil
	}
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing cacheforwardproxy state %s", cacheforwardproxyName)
		d.SetId("")
		return nil
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == d.Get("ipaddress").(string) &&
			int(v["port"].(float64)) == d.Get("port").(int) {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources ipaddress or port not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing cacheforwardproxy state %s", cacheforwardproxyName)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]
	d.Set("ipaddress", data["ipaddress"])
	setToInt("port", d, data["port"])

	return nil

}

func deleteCacheforwardproxyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCacheforwardproxyFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheforwardproxyName := d.Get("ipaddress").(string)
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("port:%v", d.Get("port").(int)))
	err := client.DeleteResourceWithArgs(service.Cacheforwardproxy.Type(), cacheforwardproxyName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
