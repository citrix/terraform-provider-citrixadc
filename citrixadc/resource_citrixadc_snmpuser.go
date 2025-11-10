package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSnmpuser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmpuserFunc,
		ReadContext:   readSnmpuserFunc,
		UpdateContext: updateSnmpuserFunc,
		DeleteContext: deleteSnmpuserFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"authpasswd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privpasswd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Get("name").(string)

	snmpuser := snmp.Snmpuser{
		Authpasswd: d.Get("authpasswd").(string),
		Authtype:   d.Get("authtype").(string),
		Group:      d.Get("group").(string),
		Name:       d.Get("name").(string),
		Privpasswd: d.Get("privpasswd").(string),
		Privtype:   d.Get("privtype").(string),
	}

	_, err := client.AddResource(service.Snmpuser.Type(), snmpuserName, &snmpuser)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(snmpuserName)

	return readSnmpuserFunc(ctx, d, meta)
}

func readSnmpuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpuser state %s", snmpuserName)
	data, err := client.FindResource(service.Snmpuser.Type(), snmpuserName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpuser state %s", snmpuserName)
		d.SetId("")
		return nil
	}
	log.Printf("DATA: %v", data)
	d.Set("name", data["name"])
	//d.Set("authpasswd", data["authpasswd"])
	d.Set("authtype", data["authtype"])
	d.Set("group", data["group"])
	//d.Set("privpasswd", data["privpasswd"])
	d.Set("privtype", data["privtype"])

	return nil

}

func updateSnmpuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Get("name").(string)

	snmpuser := snmp.Snmpuser{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("authpasswd") {
		log.Printf("[DEBUG]  citrixadc-provider: Authpasswd has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Authpasswd = d.Get("authpasswd").(string)
		snmpuser.Authtype = d.Get("authtype").(string)
		hasChange = true
	}
	if d.HasChange("authtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtype has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Authtype = d.Get("authtype").(string)
		snmpuser.Authpasswd = d.Get("authpasswd").(string)
		hasChange = true
	}
	if d.HasChange("group") {
		log.Printf("[DEBUG]  citrixadc-provider: Group has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Group = d.Get("group").(string)
		hasChange = true
	}
	if d.HasChange("privpasswd") {
		log.Printf("[DEBUG]  citrixadc-provider: Privpasswd has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Privpasswd = d.Get("privpasswd").(string)
		snmpuser.Privtype = d.Get("privtype").(string)
		hasChange = true
	}
	if d.HasChange("privtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Privtype has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Privtype = d.Get("privtype").(string)
		snmpuser.Privpasswd = d.Get("privpasswd").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Snmpuser.Type(), snmpuserName, &snmpuser)
		if err != nil {
			return diag.Errorf("Error updating snmpuser %s", snmpuserName)
		}
	}
	return readSnmpuserFunc(ctx, d, meta)
}

func deleteSnmpuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Id()
	err := client.DeleteResource(service.Snmpuser.Type(), snmpuserName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
