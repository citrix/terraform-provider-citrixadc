package citrixadc

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsnameserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsnameserverFunc,
		ReadContext:   readDnsnameserverFunc,
		UpdateContext: updateDnsnameserverFunc,
		DeleteContext: deleteDnsnameserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"dnsprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsvservername": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply
				ForceNew: true,
			},
		},
	}
}

func createDnsnameserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnameserver := dns.Dnsnameserver{
		Dnsprofilename: d.Get("dnsprofilename").(string),
		Local:          d.Get("local").(bool),
		State:          d.Get("state").(string),
		Type:           d.Get("type").(string),
	}
	var PrimaryId string
	if Ip, ok := d.GetOk("ip"); ok {
		PrimaryId = Ip.(string)
		dnsnameserver.Ip = PrimaryId
	} else if dnsvserver, ok := d.GetOk("dnsvservername"); ok {
		PrimaryId = dnsvserver.(string)
		dnsnameserver.Dnsvservername = PrimaryId
	}

	_, err := client.AddResource(service.Dnsnameserver.Type(), PrimaryId, &dnsnameserver)
	if err != nil {
		return diag.FromErr(err)
	}
	if val, ok := d.GetOk("type"); ok {
		PrimaryId = PrimaryId + "," + val.(string)
	} else {
		// the default value of attribute type is "UDP". So, it is appended implicitly when not specified by the user.
		PrimaryId = PrimaryId + ",UDP"
	}

	d.SetId(PrimaryId)

	return readDnsnameserverFunc(ctx, d, meta)
}

func readDnsnameserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	PrimaryId := d.Id()

	// To make the resource backward compatible, in the prev state file user will have ID with 1 value, but in release v1.27.0 we have updated Id. So here we are changing the code to make it backward compatible
	// here we are checking for id, if it has 1 elements then we are appending the 2rd attribute to the old Id.
	oldIdSlice := strings.Split(PrimaryId, ",")

	if len(oldIdSlice) == 1 {
		if val, ok := d.GetOk("type"); ok {
			PrimaryId = PrimaryId + "," + val.(string)
		} else {
			PrimaryId = PrimaryId + ",UDP"
		}

		d.SetId(PrimaryId)
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading dnsnameserver state %s", PrimaryId)
	findParams := service.FindParams{
		ResourceType: service.Dnsnameserver.Type(),
	}

	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnameserver state %s", PrimaryId)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: dns nameserver does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	idSlice := strings.SplitN(PrimaryId, ",", 2)
	name := idSlice[0]
	dns_type := idSlice[1]

	// When type is UDP_TCP, the ADC creates two separate entries (UDP and TCP)
	// We need to check for both entries
	var typesToCheck []string
	if dns_type == "UDP_TCP" {
		typesToCheck = []string{"UDP", "TCP"}
	} else {
		typesToCheck = []string{dns_type}
	}

	foundIndex := -1
	for _, checkType := range typesToCheck {
		for i, dnsnameserver := range dataArray {
			match := false
			if dnsnameserver["ip"] == name || dnsnameserver["dnsvservername"] == name {
				match = true
			}
			if match == true {
				if dnsnameserver["type"] != checkType {
					match = false
				}
			}
			if match {
				foundIndex = i
				break
			}
		}
		if foundIndex != -1 {
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnsnameserver not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnameserver state %s", PrimaryId)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("dnsvservername", data["dnsvservername"])
	d.Set("ip", data["ip"])
	// attribute local is not part of GET response.
	d.Set("local", d.Get("local").(bool))
	d.Set("state", data["state"])
	// Keep the original type from config (UDP_TCP) rather than the individual entry type
	d.Set("type", dns_type)

	return nil

}

func updateDnsnameserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client

	PrimaryId := d.Id()

	idSlice := strings.SplitN(PrimaryId, ",", 2)
	name := idSlice[0]

	dnsnameserver := dns.Dnsnameserver{
		Ip:             d.Get("ip").(string),
		Dnsvservername: d.Get("dnsvservername").(string),
	}
	hasChange := false
	stateChange := false
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsprofilename has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}

	if d.HasChange("local") {
		log.Printf("[DEBUG]  citrixadc-provider: Local has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.Local = d.Get("local").(bool)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.State = d.Get("state").(string)
		stateChange = true
	}
	if stateChange {
		err := doDnsvserverStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling dnsnameserver %s", PrimaryId)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Dnsnameserver.Type(), name, &dnsnameserver)
		if err != nil {
			return diag.Errorf("Error updating dnsnameserver %s", PrimaryId)
		}
	}
	return readDnsnameserverFunc(ctx, d, meta)
}

func deleteDnsnameserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	PrimaryId := d.Id()
	idSlice := strings.SplitN(PrimaryId, ",", 2)
	Name := idSlice[0]
	dns_type := idSlice[1]

	if val, ok := d.GetOk("dnsvservername"); ok && Name == val { // if the user gives `dnsvservername`, then we need to directly call delete operation.
		err := client.DeleteResource(service.Dnsnameserver.Type(), Name)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("")
		return nil
	}

	// When type is UDP_TCP, we need to delete both UDP and TCP entries
	var typesToDelete []string
	if dns_type == "UDP_TCP" {
		typesToDelete = []string{"UDP", "TCP"}
	} else {
		typesToDelete = []string{dns_type}
	}

	for _, deleteType := range typesToDelete {
		argsMap := make(map[string]string)
		argsMap["type"] = url.QueryEscape(deleteType)
		err := client.DeleteResourceWithArgsMap(service.Dnsnameserver.Type(), Name, argsMap)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")

	return nil
}
func doDnsvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doDnsvserverStateChange")

	dnsvserver := dns.Dnsnameserver{
		Ip:             d.Get("ip").(string),
		Dnsvservername: d.Get("dnsvservername").(string),
	}

	newstate := d.Get("state")

	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Dnsnameserver.Type(), dnsvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		err := client.ActOnResource(service.Dnsnameserver.Type(), dnsvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("doDnsvserverStateChange : \"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
