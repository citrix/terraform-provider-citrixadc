package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcForwardingsession() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createForwardingsessionFunc,
		ReadContext:   readForwardingsessionFunc,
		UpdateContext: updateForwardingsessionFunc,
		DeleteContext: deleteForwardingsessionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"acl6name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aclname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connfailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sourceroutecache": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createForwardingsessionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Get("name").(string)

	forwardingsession := network.Forwardingsession{
		Acl6name:         d.Get("acl6name").(string),
		Aclname:          d.Get("aclname").(string),
		Connfailover:     d.Get("connfailover").(string),
		Name:             d.Get("name").(string),
		Netmask:          d.Get("netmask").(string),
		Network:          d.Get("network").(string),
		Processlocal:     d.Get("processlocal").(string),
		Sourceroutecache: d.Get("sourceroutecache").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		forwardingsession.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Forwardingsession.Type(), forwardingsessionName, &forwardingsession)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(forwardingsessionName)

	return readForwardingsessionFunc(ctx, d, meta)
}

func readForwardingsessionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading forwardingsession state %s", forwardingsessionName)
	data, err := client.FindResource(service.Forwardingsession.Type(), forwardingsessionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing forwardingsession state %s", forwardingsessionName)
		d.SetId("")
		return nil
	}
	d.Set("acl6name", data["acl6name"])
	d.Set("aclname", data["aclname"])
	d.Set("connfailover", data["connfailover"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("processlocal", data["processlocal"])
	d.Set("sourceroutecache", data["sourceroutecache"])
	setToInt("td", d, data["td"])

	return nil

}

func updateForwardingsessionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Get("name").(string)

	forwardingsession := network.Forwardingsession{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acl6name") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl6name has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Acl6name = d.Get("acl6name").(string)
		hasChange = true
	}
	if d.HasChange("aclname") {
		log.Printf("[DEBUG]  citrixadc-provider: Aclname has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Aclname = d.Get("aclname").(string)
		hasChange = true
	}
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Connfailover has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  citrixadc-provider: Processlocal has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("sourceroutecache") {
		log.Printf("[DEBUG]  citrixadc-provider: Sourceroutecache has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Sourceroutecache = d.Get("sourceroutecache").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Forwardingsession.Type(), forwardingsessionName, &forwardingsession)
		if err != nil {
			return diag.Errorf("Error updating forwardingsession %s", forwardingsessionName)
		}
	}
	return readForwardingsessionFunc(ctx, d, meta)
}

func deleteForwardingsessionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Id()
	err := client.DeleteResource(service.Forwardingsession.Type(), forwardingsessionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
