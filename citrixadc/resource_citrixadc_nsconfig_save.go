package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsconfigSave() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsconfigSaveFunc,
		Read:          schema.Noop,
		DeleteContext: deleteNsconfigSaveFunc,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"concurrent_save_ok": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"concurrent_save_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: true,
			},
			"concurrent_save_timeout": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "5m",
				ForceNew: true,
			},
			"concurrent_save_interval": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
				ForceNew: true,
			},
			"save_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
		},
	}
}

func createNsconfigSaveFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsconfigSaveFunc")

	timestamp := d.Get("timestamp").(string)
	log.Printf("[DEBUG]  citrixadc-provider: timestamp %s", timestamp)

	err := doSaveConfig(d, meta)

	if err != nil {
		if !strings.Contains(err.Error(), "\"errorcode\": 293") {
			return diag.FromErr(err)
		}
		// Fallthrough

		// Check concurrent save flag
		if !d.Get("concurrent_save_ok").(bool) {
			return diag.FromErr(err)
		}
		// Fallthrough

		concurrentSaveRetries := d.Get("concurrent_save_retries").(int)

		// Do retries only when it is a non zero value
		if concurrentSaveRetries > 0 {

			// Do retries
			var concurrent_save_interval time.Duration
			if concurrent_save_interval, err = time.ParseDuration(d.Get("concurrent_save_interval").(string)); err != nil {
				return diag.FromErr(err)
			}

			var concurrent_save_timeout time.Duration
			if concurrent_save_timeout, err = time.ParseDuration(d.Get("concurrent_save_timeout").(string)); err != nil {
				return diag.FromErr(err)
			}
			stateConf := &resource.StateChangeConf{
				Pending:        []string{"saving"},
				Target:         []string{"saved"},
				Refresh:        saveConfigPoll(d, meta),
				PollInterval:   concurrent_save_interval,
				Delay:          concurrent_save_interval,
				Timeout:        concurrent_save_timeout,
				NotFoundChecks: concurrentSaveRetries,
			}

			_, err = stateConf.WaitForState()
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(timestamp)

	return nil
}

func doSaveConfig(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In doSaveConfig")
	client := meta.(*NetScalerNitroClient).client

	nsconfig := ns.Nsconfig{
		All: d.Get("all").(bool),
	}

	err := client.ActOnResource(service.Nsconfig.Type(), &nsconfig, "save")
	return err
}

func saveConfigPoll(d *schema.ResourceData, meta interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] citrixadc-provider: In saveConfigPoll")
		err := doSaveConfig(d, meta)
		if err != nil {
			if strings.Contains(err.Error(), "\"errorcode\": 293") {
				return nil, "saving", nil
			} else {
				return nil, "saving", err
			}
		} else {
			return "saved", "saved", nil
		}
	}
}

func deleteNsconfigSaveFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsconfigSaveFunc")

	if !d.Get("save_on_destroy").(bool) {
		log.Printf("[DEBUG]  citrixadc-provider: No save_on_destroy")
		d.SetId("")
		return nil
	}
	// Fallthrough

	err := doSaveConfig(d, meta)

	if err != nil {
		if !strings.Contains(err.Error(), "\"errorcode\": 293") {
			return diag.FromErr(err)
		}
		// Fallthrough

		// Check concurrent save flag
		if !d.Get("concurrent_save_ok").(bool) {
			return diag.FromErr(err)
		}
		// Fallthrough

		concurrentSaveRetries := d.Get("concurrent_save_retries").(int)

		// Do retries only when it is a non zero value
		if concurrentSaveRetries > 0 {

			// Do retries
			var concurrent_save_interval time.Duration
			if concurrent_save_interval, err = time.ParseDuration(d.Get("concurrent_save_interval").(string)); err != nil {
				return diag.FromErr(err)
			}

			var concurrent_save_timeout time.Duration
			if concurrent_save_timeout, err = time.ParseDuration(d.Get("concurrent_save_timeout").(string)); err != nil {
				return diag.FromErr(err)
			}
			stateConf := &resource.StateChangeConf{
				Pending:        []string{"saving"},
				Target:         []string{"saved"},
				Refresh:        saveConfigPoll(d, meta),
				PollInterval:   concurrent_save_interval,
				Delay:          concurrent_save_interval,
				Timeout:        concurrent_save_timeout,
				NotFoundChecks: concurrentSaveRetries,
			}

			_, err = stateConf.WaitForState()
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	d.SetId("")
	return nil
}
