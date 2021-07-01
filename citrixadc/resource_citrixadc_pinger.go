package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/utility"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcPinger() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPingerFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"c": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"i": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"n": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"p": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"q": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"s": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"t": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"forcenew_id_set": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createPingerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPingerFunc")
	client := meta.(*NetScalerNitroClient).client
	pingerId := resource.PrefixedUniqueId("tf-pinger-")
	ping := utility.Ping{
		C:        uint32(d.Get("c").(int)),
		HostName: d.Get("hostname").(string),
		I:        uint32(d.Get("i").(int)),
		N:        d.Get("n").(bool),
		P:        d.Get("p").(string),
		Q:        d.Get("q").(bool),
		S:        uint32(d.Get("s").(int)),
		T:        uint32(d.Get("t").(int)),
	}

	if err := client.ActOnResource("ping", &ping, ""); err != nil {
		return err
	}

	d.SetId(pingerId)

	return nil
}
