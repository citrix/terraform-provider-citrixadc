package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/utility"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPinger() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPingerFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"c": {
				Type:     schema.TypeInt,
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
			"i": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"n": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"p": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"q": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"s": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"t": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"forcenew_id_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createPingerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPingerFunc")
	client := meta.(*NetScalerNitroClient).client
	pingerId := resource.PrefixedUniqueId("tf-pinger-")
	ping := utility.Ping{
		C:        d.Get("c").(int),
		HostName: d.Get("hostname").(string),
		I:        d.Get("i").(int),
		N:        d.Get("n").(bool),
		P:        d.Get("p").(string),
		Q:        d.Get("q").(bool),
		S:        d.Get("s").(int),
		T:        d.Get("t").(int),
	}

	if err := client.ActOnResource("ping", &ping, ""); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(pingerId)

	return nil
}
