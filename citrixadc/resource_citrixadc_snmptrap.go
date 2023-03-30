package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSnmptrap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmptrapFunc,
		Read:          readSnmptrapFunc,
		Update:        updateSnmptrapFunc,
		Delete:        deleteSnmptrapFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"trapdestination": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"trapclass": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allpartitions": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"communityname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "V2",
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSnmptrapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmptrapFunc")
	client := meta.(*NetScalerNitroClient).client
	snmptrapId := d.Get("trapclass").(string) + "," + d.Get("trapdestination").(string) + "," + d.Get("version").(string)

	snmptrap := snmp.Snmptrap{
		Allpartitions:   d.Get("allpartitions").(string),
		Communityname:   d.Get("communityname").(string),
		Destport:        d.Get("destport").(int),
		Severity:        d.Get("severity").(string),
		Srcip:           d.Get("srcip").(string),
		Td:              d.Get("td").(int),
		Trapclass:       d.Get("trapclass").(string),
		Trapdestination: d.Get("trapdestination").(string),
		Version:         d.Get("version").(string),
	}

	_, err := client.AddResource(service.Snmptrap.Type(), snmptrapId, &snmptrap)
	if err != nil {
		return err
	}

	d.SetId(snmptrapId)

	err = readSnmptrapFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmptrap but we can't read it ?? %s", snmptrapId)
		return nil
	}
	return nil
}

func readSnmptrapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmptrapFunc")
	client := meta.(*NetScalerNitroClient).client
	snmptrapId := d.Id()

	idSlice := strings.SplitN(snmptrapId, ",", 3)
	trapclass := idSlice[0]
	trapdestination := idSlice[1]
	version := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading snmptrap state %s", snmptrapId)
	findParams := service.FindParams{
		ResourceType: service.Snmptrap.Type(),
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmptrap state %s", snmptrapId)
		d.SetId("")
		return nil
	}

	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: snmptrap does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["trapclass"].(string) == trapclass && v["trapdestination"].(string) == trapdestination && v["version"].(string) == version {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams snmptrap not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmptrap state %s", snmptrapId)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("allpartitions", data["allpartitions"])
	d.Set("communityname", data["communityname"])
	d.Set("destport", data["destport"])
	d.Set("severity", data["severity"])
	d.Set("srcip", data["srcip"])
	d.Set("td", data["td"])
	d.Set("trapclass", data["trapclass"])
	d.Set("trapdestination", data["trapdestination"])
	d.Set("version", data["version"])

	return nil

}

func updateSnmptrapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmptrapFunc")
	client := meta.(*NetScalerNitroClient).client
	snmptrapId := d.Id()

	idSlice := strings.SplitN(snmptrapId, ",", 3)
	trapclass := idSlice[0]
	trapdestination := idSlice[1]
	version := idSlice[2]

	snmptrap := snmp.Snmptrap{
		Trapclass:       trapclass,
		Trapdestination: trapdestination,
		Version:         version,
	}
	hasChange := false
	if d.HasChange("allpartitions") {
		log.Printf("[DEBUG]  citrixadc-provider: Allpartitions has changed for snmptrap %s, starting update", snmptrapId)
		snmptrap.Allpartitions = d.Get("allpartitions").(string)
		hasChange = true
	}
	if d.HasChange("communityname") {
		log.Printf("[DEBUG]  citrixadc-provider: Communityname has changed for snmptrap %s, starting update", snmptrapId)
		snmptrap.Communityname = d.Get("communityname").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG]  citrixadc-provider: Destport has changed for snmptrap %s, starting update", snmptrapId)
		snmptrap.Destport = d.Get("destport").(int)
		hasChange = true
	}
	if d.HasChange("severity") {
		log.Printf("[DEBUG]  citrixadc-provider: Severity has changed for snmptrap %s, starting update", snmptrapId)
		snmptrap.Severity = d.Get("severity").(string)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for snmptrap %s, starting update", snmptrapId)
		snmptrap.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for snmptrap %s, starting update", snmptrapId)
		snmptrap.Td = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmptrap.Type(), &snmptrap)
		if err != nil {
			return fmt.Errorf("Error updating snmptrap %s", snmptrapId)
		}
	}
	return readSnmptrapFunc(d, meta)
}

func deleteSnmptrapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmptrapFunc")
	client := meta.(*NetScalerNitroClient).client
	snmptrapId := d.Id()
	idSlice := strings.SplitN(snmptrapId, ",", 3)

	trapclass := idSlice[0]
	trapdestination := idSlice[1]
	version := idSlice[2]

	args := make([]string, 0)

	args = append(args, fmt.Sprintf("trapdestination:%s", trapdestination))
	args = append(args, fmt.Sprintf("version:%s", version))
	if val, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%d", val.(int)))
	}

	err := client.DeleteResourceWithArgs(service.Snmptrap.Type(), trapclass, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
