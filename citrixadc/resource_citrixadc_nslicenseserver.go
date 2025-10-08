package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcNslicenseserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNslicenseserverFunc,
		ReadContext:   readNslicenseserverFunc,
		DeleteContext: deleteNslicenseserverFunc,
		Schema: map[string]*schema.Schema{
			"forceupdateip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createNslicenseserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslicenseserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseserverId := d.Get("servername").(string)

	nslicenseserver := ns.Nslicenseserver{
		Forceupdateip: d.Get("forceupdateip").(bool),
		Nodeid:        d.Get("nodeid").(int),
		Port:          d.Get("port").(int),
		Servername:    d.Get("servername").(string),
	}

	_, err := client.AddResource("nslicenseserver", "", &nslicenseserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nslicenseserverId)

	return readNslicenseserverFunc(ctx, d, meta)
}

func readNslicenseserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslicenseserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseserverId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nslicenseserver state %s", nslicenseserverId)

	findParams := service.FindParams{
		ResourceType: "nslicenseserver",
	}

	licenseServers, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslicenseserver state %s", nslicenseserverId)
		d.SetId("")
		return nil
	}
	if len(licenseServers) == 0 {
		// There is no license server configured
		d.SetId("")
	} else {
		// License server will return at most 1 element
		data := licenseServers[0]

		d.Set("forceupdateip", data["forceupdateip"])
		setToInt("nodeid", d, data["nodeid"])
		setToInt("port", d, data["port"])
		d.Set("servername", data["servername"])
	}

	return nil

}

func deleteNslicenseserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslicenseserverFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make([]string, 0, 1)
	args = append(args, fmt.Sprintf("servername:%s", d.Get("servername").(string)))
	err := client.DeleteResourceWithArgs("nslicenseserver", "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
