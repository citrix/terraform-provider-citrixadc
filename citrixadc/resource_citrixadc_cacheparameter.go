package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCacheparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCacheparameterFunc,
		Read:          readCacheparameterFunc,
		Update:        updateCacheparameterFunc,
		Delete:        deleteCacheparameterFunc,
		Schema: map[string]*schema.Schema{
			"enablebypass": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablehaobjpersist": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxpostlen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prefetchmaxpending": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verifyusing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"via": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCacheparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCacheparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheparameterName := resource.PrefixedUniqueId("tf-cacheparameter-")
	cacheparameter := make(map[string]interface{})

	if v, ok := d.GetOk("enablebypass"); ok {
		cacheparameter["enablebypass"] = v.(int)
	}
	if v, ok := d.GetOkExists("enablehaobjpersist"); ok {
		cacheparameter["enablehaobjpersist"] = v.(string)
	}
	if v, ok := d.GetOkExists("maxpostlen"); ok {
		cacheparameter["maxpostlen"] = v.(int)
	}
	if v, ok := d.GetOk("memlimit"); ok {
		cacheparameter["memlimit"] = v.(int)
	}
	if v, ok := d.GetOk("prefetchmaxpending"); ok {
		cacheparameter["prefetchmaxpending"] = v.(int)
	}
	if v, ok := d.GetOk("undefaction"); ok {
		cacheparameter["undefaction"] = v.(string)
	}
	if v, ok := d.GetOk("verifyusing"); ok {
		cacheparameter["verifyusing"] = v.(string)
	}
	if v, ok := d.GetOk("via"); ok {
		cacheparameter["via"] = v.(string)
	}

	err := client.UpdateUnnamedResource(service.Cacheparameter.Type(), &cacheparameter)
	if err != nil {
		return err
	}

	d.SetId(cacheparameterName)

	err = readCacheparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cacheparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readCacheparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCacheparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading cacheparameter state")
	data, err := client.FindResource(service.Cacheparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cacheparameter state")
		d.SetId("")
		return nil
	}
	d.Set("enablebypass", data["enablebypass"])
	d.Set("enablehaobjpersist", data["enablehaobjpersist"])
	d.Set("maxpostlen", data["maxpostlen"])
	d.Set("memlimit", data["memlimit"])
	d.Set("prefetchmaxpending", data["prefetchmaxpending"])
	d.Set("undefaction", data["undefaction"])
	d.Set("verifyusing", data["verifyusing"])
	d.Set("via", data["via"])

	return nil

}

func updateCacheparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCacheparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	cacheparameter := cache.Cacheparameter{}
	hasChange := false
	if d.HasChange("enablebypass") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablebypass has changed for cacheparameter, starting update")
		cacheparameter.Enablebypass = d.Get("enablebypass").(string)
		hasChange = true
	}
	if d.HasChange("enablehaobjpersist") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablehaobjpersist has changed for cacheparameter, starting update")
		cacheparameter.Enablehaobjpersist = d.Get("enablehaobjpersist").(string)
		hasChange = true
	}
	if d.HasChange("maxpostlen") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpostlen has changed for cacheparameter, starting update")
		cacheparameter.Maxpostlen = d.Get("maxpostlen").(int)
		hasChange = true
	}
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for cacheparameter, starting update")
		cacheparameter.Memlimit = d.Get("memlimit").(int)
		hasChange = true
	}
	if d.HasChange("prefetchmaxpending") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefetchmaxpending has changed for cacheparameter, starting update")
		cacheparameter.Prefetchmaxpending = d.Get("prefetchmaxpending").(int)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for cacheparameter, starting update")
		cacheparameter.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}
	if d.HasChange("verifyusing") {
		log.Printf("[DEBUG]  citrixadc-provider: Verifyusing has changed for cacheparameter, starting update")
		cacheparameter.Verifyusing = d.Get("verifyusing").(string)
		hasChange = true
	}
	if d.HasChange("via") {
		log.Printf("[DEBUG]  citrixadc-provider: Via has changed for cacheparameter, starting update")
		cacheparameter.Via = d.Get("via").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cacheparameter.Type(), &cacheparameter)
		if err != nil {
			return fmt.Errorf("Error updating cacheparameter")
		}
	}
	return readCacheparameterFunc(d, meta)
}

func deleteCacheparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCacheparameterFunc")
	//cacheparameter does not suppor DELETE operation
	d.SetId("")

	return nil
}
