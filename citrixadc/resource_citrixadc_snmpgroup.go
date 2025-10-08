package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSnmpgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmpgroupFunc,
		ReadContext:   readSnmpgroupFunc,
		UpdateContext: updateSnmpgroupFunc,
		DeleteContext: deleteSnmpgroupFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"securitylevel": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"readviewname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createSnmpgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Get("name").(string)
	snmpgroup := snmp.Snmpgroup{
		Name:          d.Get("name").(string),
		Readviewname:  d.Get("readviewname").(string),
		Securitylevel: d.Get("securitylevel").(string),
	}

	_, err := client.AddResource(service.Snmpgroup.Type(), snmpgroupName, &snmpgroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(snmpgroupName)

	return readSnmpgroupFunc(ctx, d, meta)
}

func readSnmpgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpgroup state %s", snmpgroupName)

	dataArr, err := client.FindAllResources(service.Snmpgroup.Type())

	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: snmpgroup does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["name"].(string) == d.Get("name") {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllResources snmpgroup not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmpgroup state %s", snmpgroupName)
		d.SetId("")
		return nil
	}

	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpgroup state %s", snmpgroupName)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("name", data["name"])
	d.Set("readviewname", data["readviewname"])
	d.Set("securitylevel", data["securitylevel"])

	return nil

}

func updateSnmpgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Get("name").(string)

	snmpgroup := snmp.Snmpgroup{
		Name:          d.Get("name").(string),
		Securitylevel: d.Get("securitylevel").(string),
		Readviewname:  d.Get("readviewname").(string),
	}

	hasChange := false
	if d.HasChange("readviewname") {
		log.Printf("[DEBUG]  citrixadc-provider: Readviewname has changed for snmpgroup %s, starting update", snmpgroupName)
		//snmpgroup.Readviewname = d.Get("readviewname").(string)
		hasChange = true
	}
	if d.HasChange("securitylevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Securitylevel has changed for snmpgroup %s, starting update", snmpgroupName)
		//snmpgroup.Securitylevel = d.Get("securitylevel").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpgroup.Type(), &snmpgroup)
		if err != nil {
			return diag.Errorf("Error updating snmpgroup %s", snmpgroupName)
		}
	}
	return readSnmpgroupFunc(ctx, d, meta)
}

func deleteSnmpgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("securitylevel:%s", d.Get("securitylevel").(string)))

	err := client.DeleteResourceWithArgs(service.Snmpgroup.Type(), snmpgroupName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
