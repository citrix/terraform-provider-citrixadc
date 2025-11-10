package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnpool() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnpoolFunc,
		ReadContext:   readLsnpoolFunc,
		UpdateContext: updateLsnpoolFunc,
		DeleteContext: deleteLsnpoolFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"poolname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nattype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"maxportrealloctmq": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portblockallocation": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"portrealloctimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnpoolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Get("poolname").(string)

	lsnpool := make(map[string]interface{})
	if v, ok := d.GetOkExists("portrealloctimeout"); ok {
		lsnpool["portrealloctimeout"] = v.(int)
	}
	if v, ok := d.GetOk("portblockallocation"); ok {
		lsnpool["portblockallocation"] = v.(string)
	}
	if v, ok := d.GetOkExists("poolname"); ok {
		lsnpool["poolname"] = v.(string)
	}
	if v, ok := d.GetOk("nattype"); ok {
		lsnpool["nattype"] = v.(string)
	}
	if v, ok := d.GetOkExists("maxportrealloctmq"); ok {
		val, _ := strconv.Atoi(v.(string))
		lsnpool["maxportrealloctmq"] = val
	}

	_, err := client.AddResource("lsnpool", lsnpoolName, &lsnpool)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnpoolName)

	return readLsnpoolFunc(ctx, d, meta)
}

func readLsnpoolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnpool state %s", lsnpoolName)
	data, err := client.FindResource("lsnpool", lsnpoolName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnpool state %s", lsnpoolName)
		d.SetId("")
		return nil
	}
	d.Set("maxportrealloctmq", data["maxportrealloctmq"])
	d.Set("nattype", data["nattype"])
	d.Set("poolname", data["poolname"])
	d.Set("portblockallocation", data["portblockallocation"])
	setToInt("portrealloctimeout", d, data["portrealloctimeout"])

	return nil

}

func updateLsnpoolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Get("poolname").(string)

	lsnpool := lsn.Lsnpool{
		Poolname: d.Get("poolname").(string),
	}
	hasChange := false
	if d.HasChange("maxportrealloctmq") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxportrealloctmq has changed for lsnpool %s, starting update", lsnpoolName)
		val, _ := strconv.Atoi(d.Get("maxportrealloctmq").(string))
		lsnpool.Maxportrealloctmq = intPtr(val)
		hasChange = true
	}
	if d.HasChange("portrealloctimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Portrealloctimeout has changed for lsnpool %s, starting update", lsnpoolName)
		lsnpool.Portrealloctimeout = intPtr(d.Get("portrealloctimeout").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnpool", &lsnpool)
		if err != nil {
			return diag.Errorf("Error updating lsnpool %s", lsnpoolName)
		}
	}
	return readLsnpoolFunc(ctx, d, meta)
}

func deleteLsnpoolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Id()
	err := client.DeleteResource("lsnpool", lsnpoolName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
