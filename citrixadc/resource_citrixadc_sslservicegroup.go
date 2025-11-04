package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslservicegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslservicegroupFunc,
		ReadContext:   readSslservicegroupFunc,
		UpdateContext: updateSslservicegroupFunc,
		DeleteContext: deleteSslservicegroupFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"sslclientlogs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commonname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocspstapling": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendclosenotify": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverauth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sessreuse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"snienable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strictsigdigestcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls11": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls12": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls13": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client

	var sslservicegroupName string
	if v, ok := d.GetOk("servicegroupname"); ok {
		sslservicegroupName = v.(string)
	} else {
		sslservicegroupName = resource.PrefixedUniqueId("tf-servicegroup-")
		d.Set("servicegroupname", sslservicegroupName)
	}

	sslservicegroup := ssl.Sslservicegroup{
		Commonname:           d.Get("commonname").(string),
		Ocspstapling:         d.Get("ocspstapling").(string),
		Sendclosenotify:      d.Get("sendclosenotify").(string),
		Serverauth:           d.Get("serverauth").(string),
		Servicegroupname:     d.Get("servicegroupname").(string),
		Sessreuse:            d.Get("sessreuse").(string),
		Snienable:            d.Get("snienable").(string),
		Ssl3:                 d.Get("ssl3").(string),
		Sslprofile:           d.Get("sslprofile").(string),
		Strictsigdigestcheck: d.Get("strictsigdigestcheck").(string),
		Tls1:                 d.Get("tls1").(string),
		Tls11:                d.Get("tls11").(string),
		Tls12:                d.Get("tls12").(string),
		Tls13:                d.Get("tls13").(string),
		Sslclientlogs:        d.Get("sslclientlogs").(string),
	}

	if raw := d.GetRawConfig().GetAttr("sesstimeout"); !raw.IsNull() {
		sslservicegroup.Sesstimeout = intPtr(d.Get("sesstimeout").(int))
	}

	_, err := client.UpdateResource(service.Sslservicegroup.Type(), sslservicegroupName, &sslservicegroup)
	if err != nil {
		return diag.Errorf("Error updating sslservicegroup")
	}

	d.SetId(sslservicegroupName)

	return readSslservicegroupFunc(ctx, d, meta)
}

func readSslservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading sslservicegroup state")
	servicegroupname := d.Id()
	data, err := client.FindResource(service.Sslservicegroup.Type(), servicegroupname)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup state")
		d.SetId("")
		return nil
	}
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("sslclientlogs", data["sslclientlogs"])
	d.Set("commonname", data["commonname"])
	d.Set("ocspstapling", data["ocspstapling"])
	d.Set("sendclosenotify", data["sendclosenotify"])
	d.Set("serverauth", data["serverauth"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("sessreuse", data["sessreuse"])
	setToInt("sesstimeout", d, data["sesstimeout"])
	d.Set("snienable", data["snienable"])
	d.Set("ssl3", data["ssl3"])
	d.Set("sslprofile", data["sslprofile"])
	d.Set("strictsigdigestcheck", data["strictsigdigestcheck"])
	d.Set("tls1", data["tls1"])
	d.Set("tls11", data["tls11"])
	d.Set("tls12", data["tls12"])
	d.Set("tls13", data["tls13"])

	return nil

}

func updateSslservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupname := d.Get("servicegroupname").(string)

	sslservicegroup := ssl.Sslservicegroup{
		Servicegroupname: d.Get("servicegroupname").(string),
	}
	hasChange := false
	if d.HasChange("sslclientlogs") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslclientlogs has changed for sslservicegroup, starting update")
		sslservicegroup.Sslclientlogs = d.Get("sslclientlogs").(string)
		hasChange = true
	}
	if d.HasChange("commonname") {
		log.Printf("[DEBUG]  citrixadc-provider: Commonname has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Commonname = d.Get("commonname").(string)
		hasChange = true
	}
	if d.HasChange("ocspstapling") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspstapling has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Ocspstapling = d.Get("ocspstapling").(string)
		hasChange = true
	}
	if d.HasChange("sendclosenotify") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendclosenotify has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Sendclosenotify = d.Get("sendclosenotify").(string)
		hasChange = true
	}
	if d.HasChange("serverauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverauth has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Serverauth = d.Get("serverauth").(string)
		hasChange = true
	}
	if d.HasChange("servicegroupname") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicegroupname has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Servicegroupname = d.Get("servicegroupname").(string)
		hasChange = true
	}
	if d.HasChange("sessreuse") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessreuse has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Sessreuse = d.Get("sessreuse").(string)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Sesstimeout = intPtr(d.Get("sesstimeout").(int))
		hasChange = true
	}
	if d.HasChange("snienable") {
		log.Printf("[DEBUG]  citrixadc-provider: Snienable has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Snienable = d.Get("snienable").(string)
		hasChange = true
	}
	if d.HasChange("ssl3") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssl3 has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Ssl3 = d.Get("ssl3").(string)
		hasChange = true
	}
	if d.HasChange("sslprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslprofile has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Sslprofile = d.Get("sslprofile").(string)
		hasChange = true
	}
	if d.HasChange("strictsigdigestcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Strictsigdigestcheck has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Strictsigdigestcheck = d.Get("strictsigdigestcheck").(string)
		hasChange = true
	}
	if d.HasChange("tls1") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls1 has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Tls1 = d.Get("tls1").(string)
		hasChange = true
	}
	if d.HasChange("tls11") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls11 has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Tls11 = d.Get("tls11").(string)
		hasChange = true
	}
	if d.HasChange("tls12") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls12 has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Tls12 = d.Get("tls12").(string)
		hasChange = true
	}
	if d.HasChange("tls13") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls13 has changed for sslservicegroup  %s, starting update", servicegroupname)
		sslservicegroup.Tls13 = d.Get("tls13").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Sslservicegroup.Type(), servicegroupname, &sslservicegroup)
		if err != nil {
			return diag.Errorf("Error updating sslservicegroup %s", servicegroupname)
		}
	}
	return readSslservicegroupFunc(ctx, d, meta)
}

func deleteSslservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservicegroupFunc")
	// sslservicegroup does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
