package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnintranetapplication() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnintranetapplicationFunc,
		ReadContext:   readVpnintranetapplicationFunc,
		DeleteContext: deleteVpnintranetapplicationFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"intranetapplication": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"clientapplication": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"destip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"destport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"interception": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iprange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
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
			"spoofiip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"srcip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"srcport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnintranetapplicationFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnintranetapplicationFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnintranetapplicationName := d.Get("intranetapplication").(string)
	vpnintranetapplication := vpn.Vpnintranetapplication{
		Clientapplication:   toStringList(d.Get("clientapplication").([]interface{})),
		Destip:              d.Get("destip").(string),
		Destport:            d.Get("destport").(string),
		Hostname:            d.Get("hostname").(string),
		Interception:        d.Get("interception").(string),
		Intranetapplication: d.Get("intranetapplication").(string),
		Iprange:             d.Get("iprange").(string),
		Netmask:             d.Get("netmask").(string),
		Protocol:            d.Get("protocol").(string),
		Spoofiip:            d.Get("spoofiip").(string),
		Srcip:               d.Get("srcip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("srcport"); !raw.IsNull() {
		vpnintranetapplication.Srcport = intPtr(d.Get("srcport").(int))
	}

	_, err := client.AddResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName, &vpnintranetapplication)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnintranetapplicationName)

	return readVpnintranetapplicationFunc(ctx, d, meta)
}

func readVpnintranetapplicationFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnintranetapplicationFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnintranetapplicationName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnintranetapplication state %s", vpnintranetapplicationName)
	data, err := client.FindResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnintranetapplication state %s", vpnintranetapplicationName)
		d.SetId("")
		return nil
	}
	d.Set("intranetapplication", data["intranetapplication"])
	d.Set("clientapplication", data["clientapplication"])
	d.Set("destip", data["destip"])
	d.Set("destport", data["destport"])
	d.Set("hostname", data["hostname"])
	d.Set("interception", data["interception"])
	d.Set("intranetapplication", data["intranetapplication"])
	d.Set("iprange", data["iprange"])
	d.Set("netmask", data["netmask"])
	d.Set("protocol", data["protocol"])
	d.Set("spoofiip", data["spoofiip"])
	d.Set("srcip", data["srcip"])
	setToInt("srcport", d, data["srcport"])

	return nil

}

func deleteVpnintranetapplicationFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnintranetapplicationFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnintranetapplicationName := d.Id()
	err := client.DeleteResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
