package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func resourceCitrixAdcClusternode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createClusternodeFunc,
		ReadContext:   readClusternodeFunc,
		UpdateContext: updateClusternodeFunc,
		DeleteContext: deleteClusternodeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"nodeid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backplane": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clearnodegroupconfig": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nodegroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnelmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createClusternodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := strconv.Itoa(d.Get("nodeid").(int))

	clusternode := cluster.Clusternode{
		Backplane:  d.Get("backplane").(string),
		Ipaddress:  d.Get("ipaddress").(string),
		Nodegroup:  d.Get("nodegroup").(string),
		State:      d.Get("state").(string),
		Tunnelmode: d.Get("tunnelmode").(string),
	}

	if raw := d.GetRawConfig().GetAttr("delay"); !raw.IsNull() {
		clusternode.Delay = intPtr(d.Get("delay").(int))
	}
	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		clusternode.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		clusternode.Priority = intPtr(d.Get("priority").(int))
	}

	_, err := client.AddResource(service.Clusternode.Type(), clusternodeId, &clusternode)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(clusternodeId)

	return readClusternodeFunc(ctx, d, meta)
}

func readClusternodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading clusternode state %s", clusternodeId)
	data, err := client.FindResource(service.Clusternode.Type(), clusternodeId)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing clusternode state %s", clusternodeId)
		d.SetId("")
		return nil
	}
	setToInt("nodeid", d, data["nodeid"])
	d.Set("backplane", data["backplane"])
	setToInt("delay", d, data["delay"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("nodegroup", data["nodegroup"])
	setToInt("priority", d, data["priority"])
	d.Set("state", data["state"])
	d.Set("tunnelmode", data["tunnelmode"])

	return nil

}

func updateClusternodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := strconv.Itoa(d.Get("nodeid").(int))

	clusternode := cluster.Clusternode{}

	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		clusternode.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	hasChange := false
	if d.HasChange("backplane") {
		log.Printf("[DEBUG]  citrixadc-provider: Backplane has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Backplane = d.Get("backplane").(string)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  citrixadc-provider: Delay has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Delay = intPtr(d.Get("delay").(int))
		hasChange = true
	}
	if d.HasChange("nodegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodegroup has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Nodegroup = d.Get("nodegroup").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Nodeid = intPtr(d.Get("nodeid").(int))
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Priority = intPtr(d.Get("priority").(int))
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for clusternode %s, starting update", clusternodeId)
		clusternode.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("tunnelmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Tunnelmode has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Tunnelmode = d.Get("tunnelmode").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Clusternode.Type(), clusternodeId, &clusternode)
		if err != nil {
			return diag.Errorf("Error updating clusternode %s", clusternodeId)
		}
	}
	return readClusternodeFunc(ctx, d, meta)
}

func deleteClusternodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("clearnodegroupconfig"); ok {
		args = append(args, fmt.Sprintf("clearnodegroupconfig:%s", v.(string)))
	} else {
		args = append(args, fmt.Sprintf("clearnodegroupconfig:YES"))
	}
	err := client.DeleteResourceWithArgs(service.Clusternode.Type(), clusternodeId, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
