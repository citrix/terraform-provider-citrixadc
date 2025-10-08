package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsservicefunction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsservicefunctionFunc,
		ReadContext:   readNsservicefunctionFunc,
		UpdateContext: updateNsservicefunctionFunc,
		DeleteContext: deleteNsservicefunctionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"servicefunctionname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ingressvlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createNsservicefunctionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Get("servicefunctionname").(string)

	nsservicefunction := ns.Nsservicefunction{
		Ingressvlan:         d.Get("ingressvlan").(int),
		Servicefunctionname: d.Get("servicefunctionname").(string),
	}

	_, err := client.AddResource(service.Nsservicefunction.Type(), nsservicefunctionName, &nsservicefunction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsservicefunctionName)

	return readNsservicefunctionFunc(ctx, d, meta)
}

func readNsservicefunctionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsservicefunction state %s", nsservicefunctionName)
	data, err := client.FindResource(service.Nsservicefunction.Type(), nsservicefunctionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicefunction state %s", nsservicefunctionName)
		d.SetId("")
		return nil
	}
	setToInt("ingressvlan", d, data["ingressvlan"])
	d.Set("servicefunctionname", data["servicefunctionname"])

	return nil

}

func updateNsservicefunctionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Get("servicefunctionname").(string)

	nsservicefunction := ns.Nsservicefunction{
		Servicefunctionname: d.Get("servicefunctionname").(string),
	}
	hasChange := false
	if d.HasChange("ingressvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Ingressvlan has changed for nsservicefunction %s, starting update", nsservicefunctionName)
		nsservicefunction.Ingressvlan = d.Get("ingressvlan").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsservicefunction.Type(), nsservicefunctionName, &nsservicefunction)
		if err != nil {
			return diag.Errorf("Error updating nsservicefunction %s", nsservicefunctionName)
		}
	}
	return readNsservicefunctionFunc(ctx, d, meta)
}

func deleteNsservicefunctionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Id()
	err := client.DeleteResource(service.Nsservicefunction.Type(), nsservicefunctionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
