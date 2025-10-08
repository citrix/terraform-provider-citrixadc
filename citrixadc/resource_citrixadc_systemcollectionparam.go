package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSystemcollectionparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystemcollectionparamFunc,
		ReadContext:   readSystemcollectionparamFunc,
		UpdateContext: updateSystemcollectionparamFunc,
		DeleteContext: deleteSystemcollectionparamFunc,
		Schema: map[string]*schema.Schema{
			"communityname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datapath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loglevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSystemcollectionparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemcollectionparamFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcollectionparamName := resource.PrefixedUniqueId("tf-systemcollectionparam-")

	systemcollectionparam := system.Systemcollectionparam{
		Communityname: d.Get("communityname").(string),
		Datapath:      d.Get("datapath").(string),
		Loglevel:      d.Get("loglevel").(string),
	}

	err := client.UpdateUnnamedResource(service.Systemcollectionparam.Type(), &systemcollectionparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(systemcollectionparamName)

	return readSystemcollectionparamFunc(ctx, d, meta)
}

func readSystemcollectionparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemcollectionparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading systemcollectionparam state")
	data, err := client.FindResource(service.Systemcollectionparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systemcollectionparam state")
		d.SetId("")
		return nil
	}
	//d.Set("communityname", data["communityname"])
	d.Set("datapath", data["datapath"])
	d.Set("loglevel", data["loglevel"])

	return nil

}

func updateSystemcollectionparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemcollectionparamFunc")
	client := meta.(*NetScalerNitroClient).client

	systemcollectionparam := system.Systemcollectionparam{}
	hasChange := false
	if d.HasChange("communityname") {
		log.Printf("[DEBUG]  citrixadc-provider: Communityname has changed for systemcollectionparam, starting update")
		systemcollectionparam.Communityname = d.Get("communityname").(string)
		hasChange = true
	}
	if d.HasChange("datapath") {
		log.Printf("[DEBUG]  citrixadc-provider: Datapath has changed for systemcollectionparam, starting update")
		systemcollectionparam.Datapath = d.Get("datapath").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for systemcollectionparam, starting update")
		systemcollectionparam.Loglevel = d.Get("loglevel").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Systemcollectionparam.Type(), &systemcollectionparam)
		if err != nil {
			return diag.Errorf("Error updating systemcollectionparam")
		}
	}
	return readSystemcollectionparamFunc(ctx, d, meta)
}

func deleteSystemcollectionparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemcollectionparamFunc")
	//systemcollecitonparam does not support delete operation
	d.SetId("")

	return nil
}
