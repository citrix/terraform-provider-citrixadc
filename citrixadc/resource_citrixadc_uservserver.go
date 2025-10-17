package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/user"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcUservserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createUservserverFunc,
		ReadContext:   readUservserverFunc,
		UpdateContext: updateUservserverFunc,
		DeleteContext: deleteUservserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"defaultlb": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"userprotocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"params": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createUservserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createUservserverFunc")
	client := meta.(*NetScalerNitroClient).client
	uservserverName := d.Get("name").(string)
	uservserver := user.Uservserver{
		Comment:      d.Get("comment").(string),
		Defaultlb:    d.Get("defaultlb").(string),
		Ipaddress:    d.Get("ipaddress").(string),
		Name:         d.Get("name").(string),
		Params:       d.Get("params").(string),
		State:        d.Get("state").(string),
		Userprotocol: d.Get("userprotocol").(string),
	}

	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		uservserver.Port = intPtr(d.Get("port").(int))
	}

	_, err := client.AddResource("uservserver", uservserverName, &uservserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(uservserverName)

	return readUservserverFunc(ctx, d, meta)
}

func readUservserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readUservserverFunc")
	client := meta.(*NetScalerNitroClient).client
	uservserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading uservserver state %s", uservserverName)
	data, err := client.FindResource("uservserver", uservserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing uservserver state %s", uservserverName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("defaultlb", data["defaultlb"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("params", data["params"])
	setToInt("port", d, data["port"])
	d.Set("state", data["state"])
	d.Set("userprotocol", data["userprotocol"])

	return nil

}

func updateUservserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateUservserverFunc")
	client := meta.(*NetScalerNitroClient).client
	uservserverName := d.Get("name").(string)

	uservserver := user.Uservserver{
		Name: d.Get("name").(string),
	}

	hasChange := false
	stateChange := false

	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for uservserver %s, starting update", uservserverName)
		uservserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("defaultlb") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultlb has changed for uservserver %s, starting update", uservserverName)
		uservserver.Defaultlb = d.Get("defaultlb").(string)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for uservserver %s, starting update", uservserverName)
		uservserver.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("params") {
		log.Printf("[DEBUG]  citrixadc-provider: Params has changed for uservserver %s, starting update", uservserverName)
		uservserver.Params = d.Get("params").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for uservserver %s, starting update", uservserverName)
		//uservserver.State = d.Get("state").(string)
		stateChange = true
	}

	if stateChange {
		err := doUservserverStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling user vserver %s", uservserverName)
		}
	}
	if hasChange {
		_, err := client.UpdateResource("uservserver", uservserverName, &uservserver)
		if err != nil {
			return diag.Errorf("Error updating uservserver %s", uservserverName)
		}
	}
	return readUservserverFunc(ctx, d, meta)
}

func deleteUservserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteUservserverFunc")
	client := meta.(*NetScalerNitroClient).client
	uservserverName := d.Id()
	err := client.DeleteResource("uservserver", uservserverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func doUservserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doUservserverStateChange")

	uservserver := user.Uservserver{
		Name: d.Get("name").(string),
	}
	newstate := d.Get("state").(string)

	if strings.ToLower(newstate) == "enabled" {
		err := client.ActOnResource("uservserver", uservserver, "enable")
		if err != nil {
			return err
		}
	} else if strings.ToLower(newstate) == "disabled" {
		err := client.ActOnResource("uservserver", uservserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
