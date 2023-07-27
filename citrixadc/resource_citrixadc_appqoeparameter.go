package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppqoeparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppqoeparameterFunc,
		Read:          readAppqoeparameterFunc,
		Update:        updateAppqoeparameterFunc,
		Delete:        deleteAppqoeparameterFunc,
		Schema: map[string]*schema.Schema{
			"avgwaitingclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dosattackthresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxaltrespbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessionlife": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppqoeparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppqoeparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoeparameterName := resource.PrefixedUniqueId("tf-appqoeparameter-")
	appqoeparameter := make(map[string]interface{})

	if v, ok := d.GetOkExists("avgwaitingclient"); ok {
		appqoeparameter["avgwaitingclient"] = v.(int)
	}
	if v, ok := d.GetOkExists("dosattackthresh"); ok {
		appqoeparameter["dosattackthresh"] = v.(int)
	}
	if v, ok := d.GetOk("maxaltrespbandwidth"); ok {
		appqoeparameter["maxaltrespbandwidth"] = v.(int)
	}
	if v, ok := d.GetOk("sessionlife"); ok {
		appqoeparameter["sessionlife"] = v.(int)
	}

	err := client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
	if err != nil {
		return err
	}

	d.SetId(appqoeparameterName)

	err = readAppqoeparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appqoeparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readAppqoeparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppqoeparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading appqoeparameter state")
	data, err := client.FindResource(service.Appqoeparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appqoeparameter state")
		d.SetId("")
		return nil
	}
	d.Set("avgwaitingclient", data["avgwaitingclient"])
	d.Set("dosattackthresh", data["dosattackthresh"])
	d.Set("maxaltrespbandwidth", data["maxaltrespbandwidth"])
	d.Set("sessionlife", data["sessionlife"])

	return nil

}

func updateAppqoeparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppqoeparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	appqoeparameter := appqoe.Appqoeparameter{}
	hasChange := false
	if d.HasChange("avgwaitingclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Avgwaitingclient has changed for appqoeparameter, starting update")
		appqoeparameter.Avgwaitingclient = d.Get("avgwaitingclient").(int)
		hasChange = true
	}
	if d.HasChange("dosattackthresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Dosattackthresh has changed for appqoeparameter, starting update")
		appqoeparameter.Dosattackthresh = d.Get("dosattackthresh").(int)
		hasChange = true
	}
	if d.HasChange("maxaltrespbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxaltrespbandwidth has changed for appqoeparameter, starting update")
		appqoeparameter.Maxaltrespbandwidth = d.Get("maxaltrespbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("sessionlife") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlife has changed for appqoeparameter, starting update")
		appqoeparameter.Sessionlife = d.Get("sessionlife").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
		if err != nil {
			return fmt.Errorf("Error updating appqoeparameter")
		}
	}
	return readAppqoeparameterFunc(d, meta)
}

func deleteAppqoeparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppqoeparameterFunc")
	//appqoeparameter does not suppor DELETE operation
	d.SetId("")

	return nil
}
