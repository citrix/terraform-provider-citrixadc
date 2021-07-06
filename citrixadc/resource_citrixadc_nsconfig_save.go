package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"strings"
	"time"
)

func resourceCitrixAdcNsconfigSave() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsconfigSaveFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"all": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"concurrent_save_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"concurrent_save_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: true,
			},
			"concurrent_save_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "5m",
				ForceNew: true,
			},
			"concurrent_save_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
				ForceNew: true,
			},
		},
	}
}

func createNsconfigSaveFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsconfigSaveFunc")

	timestamp := d.Get("timestamp").(string)
	log.Printf("[DEBUG]  citrixadc-provider: timestamp %s", timestamp)

	err := doSaveConfig(d, meta)

	if err != nil {
		if !strings.Contains(err.Error(), "\"errorcode\": 293") {
			return err
		}
		// Fallthrough

		// Check concurrent save flag
		if !d.Get("concurrent_save_ok").(bool) {
			return err
		}
		// Fallthrough

		concurrentSaveRetries := d.Get("concurrent_save_retries").(int)

		// Do retries only when it is a non zero value
		if concurrentSaveRetries > 0 {

			// Do retries
			var concurrent_save_interval time.Duration
			if concurrent_save_interval, err = time.ParseDuration(d.Get("concurrent_save_interval").(string)); err != nil {
				return err
			}

			var concurrent_save_timeout time.Duration
			if concurrent_save_timeout, err = time.ParseDuration(d.Get("concurrent_save_timeout").(string)); err != nil {
				return err
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
				return err
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
