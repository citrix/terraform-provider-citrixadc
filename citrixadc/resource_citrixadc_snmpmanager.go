package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcSnmpmanager() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmpmanagerFunc,
		ReadContext:   readSnmpmanagerFunc,
		UpdateContext: updateSnmpmanagerFunc,
		DeleteContext: deleteSnmpmanagerFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domainresolveretry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func createSnmpmanagerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Get("ipaddress").(string)

	snmpmanager := snmp.Snmpmanager{
		Domainresolveretry: d.Get("domainresolveretry").(int),
		Ipaddress:          d.Get("ipaddress").(string),
		Netmask:            d.Get("netmask").(string),
	}

	_, err := client.AddResource(service.Snmpmanager.Type(), snmpmanagerName, &snmpmanager)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(snmpmanagerName)

	return readSnmpmanagerFunc(ctx, d, meta)
}

func readSnmpmanagerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpmanager state %s", snmpmanagerName)

	dataArr, err := client.FindAllResources(service.Snmpmanager.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpmanager state %s", snmpmanagerName)
		d.SetId("")
		return nil
	}

	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: snmpmanager does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == snmpmanagerName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllResources snmpgroup not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmpmanager state %s", snmpmanagerName)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	setToInt("domainresolveretry", d, data["domainresolveretry"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("netmask", data["netmask"])

	return nil

}

func updateSnmpmanagerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Id()

	snmpmanager := snmp.Snmpmanager{
		Ipaddress: snmpmanagerName,
	}
	hasChange := false
	if d.HasChange("domainresolveretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainresolveretry has changed for snmpmanager %s, starting update", snmpmanagerName)
		snmpmanager.Domainresolveretry = d.Get("domainresolveretry").(int)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for snmpmanager %s, starting update", snmpmanagerName)
		snmpmanager.Netmask = d.Get("netmask").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpmanager.Type(), &snmpmanager)
		if err != nil {
			return diag.Errorf("Error updating snmpmanager %s", snmpmanagerName)
		}
	}
	return readSnmpmanagerFunc(ctx, d, meta)
}

func deleteSnmpmanagerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Id()

	args := make([]string, 0)

	args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(d.Get("netmask").(string))))

	err := client.DeleteResourceWithArgs(service.Snmpmanager.Type(), snmpmanagerName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
