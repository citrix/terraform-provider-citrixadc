package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNslicenseproxyserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNslicenseproxyserverFunc,
		ReadContext:   readNslicenseproxyserverFunc,
		UpdateContext: updateNslicenseproxyserverFunc,
		DeleteContext: deleteNslicenseproxyserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createNslicenseproxyserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var nslicenseproxyserverName string
	if v, ok := d.GetOk("serverip"); ok {
		nslicenseproxyserverName = v.(string)
	} else if v, ok := d.GetOk("servername"); ok {
		nslicenseproxyserverName = v.(string)
	}
	nslicenseproxyserver := ns.Nslicenseproxyserver{
		Port:       d.Get("port").(int),
		Serverip:   d.Get("serverip").(string),
		Servername: d.Get("servername").(string),
	}

	_, err := client.AddResource(service.Nslicenseproxyserver.Type(), nslicenseproxyserverName, &nslicenseproxyserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nslicenseproxyserverName)

	return readNslicenseproxyserverFunc(ctx, d, meta)
}

func readNslicenseproxyserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseproxyserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nslicenseproxyserver state %s", nslicenseproxyserverName)
	data, err := client.FindResource(service.Nslicenseproxyserver.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslicenseproxyserver state %s", nslicenseproxyserverName)
		d.SetId("")
		return nil
	}
	setToInt("port", d, data["port"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])

	return nil

}

func updateNslicenseproxyserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseproxyserverName := d.Id()
	nslicenseproxyserver := ns.Nslicenseproxyserver{}

	if v, ok := d.GetOk("serverip"); ok {
		nslicenseproxyserver.Serverip = v.(string)
	} else if v, ok := d.GetOk("servername"); ok {
		nslicenseproxyserver.Servername = v.(string)
	}
	hasChange := false
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for nslicenseproxyserver %s, starting update", nslicenseproxyserverName)
		nslicenseproxyserver.Port = d.Get("port").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nslicenseproxyserver.Type(), &nslicenseproxyserver)
		if err != nil {
			return diag.Errorf("Error updating nslicenseproxyserver %s", nslicenseproxyserverName)
		}
	}
	return readNslicenseproxyserverFunc(ctx, d, meta)
}

func deleteNslicenseproxyserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseproxyserverId := d.Id()
	idSlice := strings.SplitN(nslicenseproxyserverId, ",", 2)
	nslicenseproxyserverName := idSlice[0]

	err := client.DeleteResource(service.Nslicenseproxyserver.Type(), nslicenseproxyserverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
