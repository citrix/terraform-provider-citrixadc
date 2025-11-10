package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsrpcnode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsrpcnodeFunc,
		ReadContext:   readNsrpcnodeFunc,
		UpdateContext: updateNsrpcnodeFunc,
		DeleteContext: deleteNsrpcnodeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"secure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validatecert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsrpcnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsrpcnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	nsrpcnodeIpaddress := d.Get("ipaddress").(string)

	nsrpcnode := ns.Nsrpcnode{
		Ipaddress:    d.Get("ipaddress").(string),
		Password:     d.Get("password").(string),
		Secure:       d.Get("secure").(string),
		Srcip:        d.Get("srcip").(string),
		Validatecert: d.Get("validatecert").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsrpcnode.Type(), &nsrpcnode)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsrpcnodeIpaddress)

	return readNsrpcnodeFunc(ctx, d, meta)
}

func readNsrpcnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsrpcnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	nsrpcnodeIpaddress := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsrpcnode state %s", nsrpcnodeIpaddress)
	findParams := service.FindParams{
		ResourceType: "nsrpcnode",
		ResourceName: nsrpcnodeIpaddress,
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Read error %s", err.Error())
		log.Printf("[WARN] citrixadc-provider: Clearing nsrpcnode state %s", nsrpcnodeIpaddress)
		d.SetId("")
		return nil
	}

	if len(dataArray) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: Failed to find nsrpcnode %s", nsrpcnodeIpaddress)
		log.Printf("[WARN] citrixadc-provider: Clearing nsrpcnode state %s", nsrpcnodeIpaddress)
		d.SetId("")
		return nil
	}

	if len(dataArray) != 1 {
		return diag.Errorf("[ERROR] Read multiple nsprcnode instances %v", dataArray)
	}
	data := dataArray[0]

	d.Set("ipaddress", data["ipaddress"])
	// Password read is a random string that changes contantly
	//d.Set("password", data["password"])
	d.Set("secure", data["secure"])
	d.Set("srcip", data["srcip"])
	d.Set("validatecert", data["validatecert"])

	return nil

}

func updateNsrpcnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsrpcnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	nsrpcnodeIpaddress := d.Get("ipaddress").(string)

	nsrpcnode := ns.Nsrpcnode{
		Ipaddress: nsrpcnodeIpaddress,
	}
	hasChange := false
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for nsrpcnode %s, starting update", nsrpcnodeIpaddress)
		nsrpcnode.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("secure") {
		log.Printf("[DEBUG]  citrixadc-provider: Secure has changed for nsrpcnode %s, starting update", nsrpcnodeIpaddress)
		nsrpcnode.Secure = d.Get("secure").(string)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for nsrpcnode %s, starting update", nsrpcnodeIpaddress)
		nsrpcnode.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("validatecert") {
		log.Printf("[DEBUG]  citrixadc-provider: Validatecert has changed for nsrpcnode %s, starting update", nsrpcnodeIpaddress)
		nsrpcnode.Validatecert = d.Get("validatecert").(string)
		nsrpcnode.Secure = d.Get("secure").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsrpcnode.Type(), &nsrpcnode)
		if err != nil {
			return diag.Errorf("Error updating nsrpcnode %s. %s", nsrpcnodeIpaddress, err.Error())
		}
	}
	return readNsrpcnodeFunc(ctx, d, meta)
}

func deleteNsrpcnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsrpcnodeFunc")
	// Rpc node always exists in ADC
	// Just remove the reference from local state

	d.SetId("")

	return nil
}
