package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcNssimpleacl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNssimpleaclFunc,
		ReadContext:   readNssimpleaclFunc,
		DeleteContext: deleteNssimpleaclFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"aclname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aclaction": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"srcip": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"destport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"estsessions": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNssimpleaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNssimpleaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nssimpleaclName := d.Get("aclname").(string)
	nssimpleacl := ns.Nssimpleacl{
		Aclaction:   d.Get("aclaction").(string),
		Aclname:     d.Get("aclname").(string),
		Estsessions: d.Get("estsessions").(bool),
		Protocol:    d.Get("protocol").(string),
		Srcip:       d.Get("srcip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("destport"); !raw.IsNull() {
		nssimpleacl.Destport = intPtr(d.Get("destport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nssimpleacl.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		nssimpleacl.Ttl = intPtr(d.Get("ttl").(int))
	}

	_, err := client.AddResource(service.Nssimpleacl.Type(), nssimpleaclName, &nssimpleacl)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nssimpleaclName)

	return readNssimpleaclFunc(ctx, d, meta)
}

func readNssimpleaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNssimpleaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nssimpleaclName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nssimpleacl state %s", nssimpleaclName)
	data, err := client.FindResource(service.Nssimpleacl.Type(), nssimpleaclName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nssimpleacl state %s", nssimpleaclName)
		d.SetId("")
		return nil
	}
	d.Set("aclname", data["aclname"])
	d.Set("aclaction", data["aclaction"])
	d.Set("aclname", data["aclname"])
	setToInt("destport", d, data["destport"])
	d.Set("estsessions", data["estsessions"])
	d.Set("protocol", data["protocol"])
	d.Set("srcip", data["srcip"])
	setToInt("td", d, data["td"])
	setToInt("ttl", d, data["ttl"])

	return nil

}

func deleteNssimpleaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNssimpleaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nssimpleaclName := d.Id()
	err := client.DeleteResource(service.Nssimpleacl.Type(), nssimpleaclName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
