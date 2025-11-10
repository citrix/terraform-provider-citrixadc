package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcClusterinstance() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createClusterinstanceFunc,
		ReadContext:   readClusterinstanceFunc,
		UpdateContext: updateClusterinstanceFunc,
		DeleteContext: deleteClusterinstanceFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"secureheartbeats": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfdretainl2params": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clusterproxyarp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"backplanebasedview": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deadinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hellointerval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodegroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preemption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quorumtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retainconnectionsoncluster": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncstatusstrictmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createClusterinstanceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Get("clid").(int)

	clusterinstance := cluster.Clusterinstance{
		Backplanebasedview:         d.Get("backplanebasedview").(string),
		Inc:                        d.Get("inc").(string),
		Nodegroup:                  d.Get("nodegroup").(string),
		Preemption:                 d.Get("preemption").(string),
		Processlocal:               d.Get("processlocal").(string),
		Quorumtype:                 d.Get("quorumtype").(string),
		Retainconnectionsoncluster: d.Get("retainconnectionsoncluster").(string),
		Syncstatusstrictmode:       d.Get("syncstatusstrictmode").(string),
		Clusterproxyarp:            d.Get("clusterproxyarp").(string),
		Dfdretainl2params:          d.Get("dfdretainl2params").(string),
		Secureheartbeats:           d.Get("secureheartbeats").(string),
	}

	if raw := d.GetRawConfig().GetAttr("clid"); !raw.IsNull() {
		clusterinstance.Clid = intPtr(d.Get("clid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("deadinterval"); !raw.IsNull() {
		clusterinstance.Deadinterval = intPtr(d.Get("deadinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("hellointerval"); !raw.IsNull() {
		clusterinstance.Hellointerval = intPtr(d.Get("hellointerval").(int))
	}

	_, err := client.AddResource(service.Clusterinstance.Type(), strconv.Itoa(clusterinstanceName), &clusterinstance)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(clusterinstanceName))

	return readClusterinstanceFunc(ctx, d, meta)
}

func readClusterinstanceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading clusterinstance state %s", clusterinstanceName)
	data, err := client.FindResource(service.Clusterinstance.Type(), clusterinstanceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing clusterinstance state %s", clusterinstanceName)
		d.SetId("")
		return nil
	}
	d.Set("backplanebasedview", data["backplanebasedview"])
	d.Set("secureheartbeats", data["secureheartbeats"])
	d.Set("dfdretainl2params", data["dfdretainl2params"])
	d.Set("clusterproxyarp", data["clusterproxyarp"])
	setToInt("clid", d, data["clid"])
	setToInt("deadinterval", d, data["deadinterval"])
	setToInt("hellointerval", d, data["hellointerval"])
	d.Set("inc", data["inc"])
	d.Set("nodegroup", data["nodegroup"])
	d.Set("preemption", data["preemption"])
	d.Set("processlocal", data["processlocal"])
	d.Set("quorumtype", data["quorumtype"])
	d.Set("retainconnectionsoncluster", data["retainconnectionsoncluster"])
	d.Set("syncstatusstrictmode", data["syncstatusstrictmode"])

	return nil

}

func updateClusterinstanceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Get("clid").(int)

	clusterinstance := cluster.Clusterinstance{}

	if raw := d.GetRawConfig().GetAttr("clid"); !raw.IsNull() {
		clusterinstance.Clid = intPtr(d.Get("clid").(int))
	}
	hasChange := false
	if d.HasChange("secureheartbeats") {
		log.Printf("[DEBUG]  citrixadc-provider: Secureheartbeats has changed for clusterinstance, starting update")
		clusterinstance.Secureheartbeats = d.Get("secureheartbeats").(string)
		hasChange = true
	}
	if d.HasChange("dfdretainl2params") {
		log.Printf("[DEBUG]  citrixadc-provider: Dfdretainl2params has changed for clusterinstance, starting update")
		clusterinstance.Dfdretainl2params = d.Get("dfdretainl2params").(string)
		hasChange = true
	}
	if d.HasChange("clusterproxyarp") {
		log.Printf("[DEBUG]  citrixadc-provider: Clusterproxyarp has changed for clusterinstance, starting update")
		clusterinstance.Clusterproxyarp = d.Get("clusterproxyarp").(string)
		hasChange = true
	}
	if d.HasChange("backplanebasedview") {
		log.Printf("[DEBUG]  citrixadc-provider: Backplanebasedview has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Backplanebasedview = d.Get("backplanebasedview").(string)
		hasChange = true
	}
	if d.HasChange("deadinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Deadinterval has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Deadinterval = intPtr(d.Get("deadinterval").(int))
		hasChange = true
	}
	if d.HasChange("hellointerval") {
		log.Printf("[DEBUG]  citrixadc-provider: Hellointerval has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Hellointerval = intPtr(d.Get("hellointerval").(int))
		hasChange = true
	}
	if d.HasChange("inc") {
		log.Printf("[DEBUG]  citrixadc-provider: Inc has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Inc = d.Get("inc").(string)
		hasChange = true
	}
	if d.HasChange("nodegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodegroup has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Nodegroup = d.Get("nodegroup").(string)
		hasChange = true
	}
	if d.HasChange("preemption") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemption has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Preemption = d.Get("preemption").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  citrixadc-provider: Processlocal has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("quorumtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Quorumtype has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Quorumtype = d.Get("quorumtype").(string)
		hasChange = true
	}
	if d.HasChange("retainconnectionsoncluster") {
		log.Printf("[DEBUG]  citrixadc-provider: Retainconnectionsoncluster has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Retainconnectionsoncluster = d.Get("retainconnectionsoncluster").(string)
		hasChange = true
	}
	if d.HasChange("syncstatusstrictmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Syncstatusstrictmode has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Syncstatusstrictmode = d.Get("syncstatusstrictmode").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Clusterinstance.Type(), &clusterinstance)
		if err != nil {
			return diag.Errorf("Error updating clusterinstance %s", strconv.Itoa(clusterinstanceName))
		}
	}
	return readClusterinstanceFunc(ctx, d, meta)
}

func deleteClusterinstanceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Id()
	err := client.DeleteResource(service.Clusterinstance.Type(), clusterinstanceName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
