package citrixadc

import (
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func dataSourceCitrixAdcSslcipherSslvserverBindings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCitrixAdcSslcipherSslvserverBindingsRead,
		Schema: map[string]*schema.Schema{
			"ciphername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			/*
				"bound_sslvservers": &schema.Schema{
					Type:     schema.TypeList,
					Elem:     &schema.Schema{Type: schema.TypeString},
					Computed: true,
				},
			*/
			"bound_sslvservers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCitrixAdcSslcipherSslvserverBindingsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcNsversionRead")
	client := meta.(*NetScalerNitroClient).client
	sslvserverFindParams := service.FindParams{
		ResourceType: "sslvserver",
	}

	sslvserverArr, err := client.FindResourceArrayWithParams(sslvserverFindParams)
	if err != nil {
		log.Printf("[ERROR] citrixadc-provider: Error during read %s", err)
		return err
	}

	boundSslvservers := make([]string, 0)
	for _, sslvserver := range sslvserverArr {
		bindingFindParams := service.FindParams{
			ResourceType:             "sslvserver_sslciphersuite_binding",
			ResourceName:             sslvserver["vservername"].(string),
			ResourceMissingErrorCode: 461,
		}
		bindingArr, err := client.FindResourceArrayWithParams(bindingFindParams)

		// Unexpected error
		if err != nil {
			log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
			return err
		}

		// Iterate through results to find the one with the right id
		for _, v := range bindingArr {
			if v["ciphername"].(string) == d.Get("ciphername").(string) {
				boundSslvservers = append(boundSslvservers, sslvserver["vservername"].(string))
			}
		}

	}
	d.SetId(resource.PrefixedUniqueId("tf-sslcipher-sslvserver-bindings-"))
	d.Set("bound_sslvservers", strings.Join(boundSslvservers, ","))

	return nil

}
