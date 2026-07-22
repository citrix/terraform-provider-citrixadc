package citrixadc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsvpxparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsvpxparamFunc,
		ReadContext:   readNsvpxparamFunc,
		DeleteContext: deleteNsvpxparamFunc,
		Schema: map[string]*schema.Schema{
			"kvmvirtiomultiqueue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cpuyield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"masterclockcpu1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsvpxparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsvpxparamFunc")
	client := meta.(*NetScalerNitroClient).client

	nsvpxparam := ns.Nsvpxparam{
		Cpuyield:            d.Get("cpuyield").(string),
		Masterclockcpu1:     d.Get("masterclockcpu1").(string),
		Kvmvirtiomultiqueue: d.Get("kvmvirtiomultiqueue").(string),
	}

	// On a cluster, ownernode selects the node the settings apply to and is the
	// resource's identity. On a standalone VPX ownernode is not configured and
	// there is a single implicit entry. Encode this decision in the ID so Read
	// can select the correct node even during a refresh, when the raw config is
	// not available.
	nsvpxparamName := resource.PrefixedUniqueId("tf-nsvpxparam-")
	if raw := d.GetRawConfig().GetAttr("ownernode"); !raw.IsNull() {
		ownernode := d.Get("ownernode").(int)
		nsvpxparam.Ownernode = intPtr(ownernode)
		nsvpxparamName = strconv.Itoa(ownernode)
	}

	err := client.UpdateUnnamedResource("nsvpxparam", &nsvpxparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsvpxparamName)

	return readNsvpxparamFunc(ctx, d, meta)
}

func readNsvpxparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsvpxparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvpxparamName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsvpxparam state %s", nsvpxparamName)
	findParams := service.FindParams{
		ResourceType: "nsvpxparam",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return diag.FromErr(err)
	}

	// The ID encodes the cluster node this resource represents (see
	// createNsvpxparamFunc). A numeric ID means a specific ownernode was
	// configured; anything else is a standalone VPX with a single entry. This is
	// derived from the ID rather than the raw config because the config is not
	// available during a refresh.
	foundIndex := -1
	if node, convErr := strconv.Atoi(nsvpxparamName); convErr == nil {
		// Cluster mode: match by ownernode. NITRO returns ownernode as a string
		// (e.g. "0"); compare numerically.
		target := strconv.Itoa(node)
		for index, value := range dataArr {
			if fmt.Sprintf("%v", value["ownernode"]) == target {
				foundIndex = index
				break
			}
		}
	} else {
		// In standalone VPX there is only one entry for nsvpxparam
		foundIndex = 0
	}

	if foundIndex == -1 {
		// Clear state for resource
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]

	d.Set("cpuyield", data["cpuyield"])
	d.Set("kvmvirtiomultiqueue", data["kvmvirtiomultiqueue"])
	// d.Set("masterclockcpu1", data["masterclockcpu1"])
	setToInt("ownernode", d, data["ownernode"])

	return nil

}

func deleteNsvpxparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsvpxparamFunc")
	// Just delete the reference
	// Actual configuration cannot be deleted
	d.SetId("")

	return nil
}
