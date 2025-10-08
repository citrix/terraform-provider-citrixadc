package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func resourceCitrixAdcSubscriberprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSubscriberprofileFunc,
		ReadContext:   readSubscriberprofileFunc,
		UpdateContext: updateSubscriberprofileFunc,
		DeleteContext: deleteSubscriberprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"servicepath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriberrules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subscriptionidtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriptionidvalue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscriberprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscriberprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberprofileName := d.Get("ip").(string)
	subscriberprofile := subscriber.Subscriberprofile{
		Ip:                  d.Get("ip").(string),
		Servicepath:         d.Get("servicepath").(string),
		Subscriberrules:     toStringList(d.Get("subscriberrules").([]interface{})),
		Subscriptionidtype:  d.Get("subscriptionidtype").(string),
		Subscriptionidvalue: d.Get("subscriptionidvalue").(string),
		Vlan:                d.Get("vlan").(int),
	}

	_, err := client.AddResource("subscriberprofile", subscriberprofileName, &subscriberprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(subscriberprofileName)

	return readSubscriberprofileFunc(ctx, d, meta)
}

func readSubscriberprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscriberprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading subscriberprofile state %s", subscriberprofileName)
	findParams := service.FindParams{
		ResourceType:             "subscriberprofile",
		ArgsMap:                  map[string]string{"vlan": strconv.Itoa(d.Get("vlan").(int))},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberprofile state %s", subscriberprofileName)
		d.SetId("")
		return nil
	}

	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberprofile state %s", subscriberprofileName)
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["ip"].(string) == subscriberprofileName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources ip not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberprofile state %s", subscriberprofileName)
		d.SetId("")
		return nil
	}
	data := dataArr[foundIndex]
	d.Set("ip", data["ip"])
	d.Set("servicepath", data["servicepath"])
	d.Set("subscriberrules", data["subscriberrules"])
	d.Set("subscriptionidtype", data["subscriptionidtype"])
	d.Set("subscriptionidvalue", data["subscriptionidvalue"])
	setToInt("vlan", d, data["vlan"])

	return nil

}

func updateSubscriberprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscriberprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberprofileName := d.Get("ip").(string)

	subscriberprofile := subscriber.Subscriberprofile{
		Ip: d.Get("ip").(string),
	}
	hasChange := false
	if d.HasChange("servicepath") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicepath has changed for subscriberprofile %s, starting update", subscriberprofileName)
		subscriberprofile.Servicepath = d.Get("servicepath").(string)
		hasChange = true
	}
	if d.HasChange("subscriberrules") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberrules has changed for subscriberprofile %s, starting update", subscriberprofileName)
		subscriberprofile.Subscriberrules = toStringList(d.Get("subscriberrules").([]interface{}))
		hasChange = true
	}
	if d.HasChange("subscriptionidtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriptionidtype has changed for subscriberprofile %s, starting update", subscriberprofileName)
		subscriberprofile.Subscriptionidtype = d.Get("subscriptionidtype").(string)
		hasChange = true
	}
	if d.HasChange("subscriptionidvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriptionidvalue has changed for subscriberprofile %s, starting update", subscriberprofileName)
		subscriberprofile.Subscriptionidvalue = d.Get("subscriptionidvalue").(string)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for subscriberprofile %s, starting update", subscriberprofileName)
		subscriberprofile.Vlan = d.Get("vlan").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("subscriberprofile", &subscriberprofile)
		if err != nil {
			return diag.Errorf("Error updating subscriberprofile %s", subscriberprofileName)
		}
	}
	return readSubscriberprofileFunc(ctx, d, meta)
}

func deleteSubscriberprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscriberprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberprofileName := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("vlan"); ok {
		args = append(args, fmt.Sprintf("vlan:%v", v.(int)))
	} else {
		args = append(args, "vlan:0")
	}
	err := client.DeleteResourceWithArgs("subscriberprofile", subscriberprofileName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
