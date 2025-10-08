package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/pcp"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPcpserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPcpserverFunc,
		ReadContext:   readPcpserverFunc,
		UpdateContext: updatePcpserverFunc,
		DeleteContext: deletePcpserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pcpprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPcpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Get("name").(string)
	pcpserver := pcp.Pcpserver{
		Ipaddress:  d.Get("ipaddress").(string),
		Name:       d.Get("name").(string),
		Pcpprofile: d.Get("pcpprofile").(string),
		Port:       d.Get("port").(int),
	}

	_, err := client.AddResource("pcpserver", pcpserverName, &pcpserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(pcpserverName)

	return readPcpserverFunc(ctx, d, meta)
}

func readPcpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading pcpserver state %s", pcpserverName)
	data, err := client.FindResource("pcpserver", pcpserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing pcpserver state %s", pcpserverName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("pcpprofile", data["pcpprofile"])
	setToInt("port", d, data["port"])

	return nil

}

func updatePcpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Get("name").(string)

	pcpserver := pcp.Pcpserver{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("pcpprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Pcpprofile has changed for pcpserver %s, starting update", pcpserverName)
		pcpserver.Pcpprofile = d.Get("pcpprofile").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for pcpserver %s, starting update", pcpserverName)
		pcpserver.Port = d.Get("port").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("pcpserver", &pcpserver)
		if err != nil {
			return diag.Errorf("Error updating pcpserver %s", pcpserverName)
		}
	}
	return readPcpserverFunc(ctx, d, meta)
}

func deletePcpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Id()
	err := client.DeleteResource("pcpserver", pcpserverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
