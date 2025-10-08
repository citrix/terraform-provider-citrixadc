package citrixadc

import (
	"context"
	"fmt"
	"log"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceCitrixAdcNsacls() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsaclsFunc,
		UpdateContext: updateNsaclsFunc,
		ReadContext:   readNsaclsFunc,
		DeleteContext: deleteNsaclsFunc,
		Schema: map[string]*schema.Schema{
			"aclsname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acls_apply_trigger": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAclAction,
			},
			"acl": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aclaction": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aclname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"destipop": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destipval": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destportop": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destportval": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"established": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"icmpcode": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmptype": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"logstate": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocolnumber": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ratelimit": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"srcipop": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcipval": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcmac": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcportop": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcportval": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"td": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vlan": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createNsaclsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In createNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client

	var nsaclsName string
	if v, ok := d.GetOk("aclsname"); ok {
		nsaclsName = v.(string)
	} else {
		nsaclsName = resource.PrefixedUniqueId("tf-nsacl-")
		d.Set("aclsname", nsaclsName)
	}
	acls := d.Get("acl").(*schema.Set).List()
	for _, val := range acls {
		acl := val.(map[string]interface{})
		_ = createSingleAcl(acl, meta)
	}

	nsacls := ns.Nsacls{}
	err := client.ApplyResource(service.Nsacls.Type(), &nsacls)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsaclsName)

	return nil
}

func readNsaclsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	data, _ := client.FindAllResources(service.Nsacl.Type())
	acls := make([]map[string]interface{}, len(data))
	for i, a := range data {
		acl := make(map[string]interface{})
		// Only set fields defined in the schema, with proper type casting
		if v, ok := a["aclaction"]; ok {
			acl["aclaction"] = v.(string)
		}
		if v, ok := a["aclname"]; ok {
			acl["aclname"] = v.(string)
		}
		if v, ok := a["destipop"]; ok {
			acl["destipop"] = v.(string)
		}
		if v, ok := a["destipval"]; ok {
			acl["destipval"] = v.(string)
		}
		if v, ok := a["destportop"]; ok {
			acl["destportop"] = v.(string)
		}
		if v, ok := a["destportval"]; ok {
			acl["destportval"] = v.(string)
		}
		if v, ok := a["established"]; ok {
			switch val := v.(type) {
			case bool:
				acl["established"] = val
			case string:
				acl["established"] = val == "true"
			}
		}
		if v, ok := a["icmpcode"]; ok {
			switch val := v.(type) {
			case int:
				acl["icmpcode"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["icmpcode"] = intVal
				}
			}
		}
		if v, ok := a["icmptype"]; ok {
			switch val := v.(type) {
			case int:
				acl["icmptype"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["icmptype"] = intVal
				}
			}
		}
		if v, ok := a["interface"]; ok {
			acl["interface"] = v.(string)
		}
		if v, ok := a["logstate"]; ok {
			acl["logstate"] = v.(string)
		}
		if v, ok := a["priority"]; ok {
			switch val := v.(type) {
			case int:
				acl["priority"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["priority"] = intVal
				}
			}
		}
		if v, ok := a["protocol"]; ok {
			acl["protocol"] = v.(string)
		}
		if v, ok := a["protocolnumber"]; ok {
			switch val := v.(type) {
			case int:
				acl["protocolnumber"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["protocolnumber"] = intVal
				}
			}
		}
		if v, ok := a["ratelimit"]; ok {
			switch val := v.(type) {
			case int:
				acl["ratelimit"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["ratelimit"] = intVal
				}
			}
		}
		if v, ok := a["srcipop"]; ok {
			acl["srcipop"] = v.(string)
		}
		if v, ok := a["srcipval"]; ok {
			acl["srcipval"] = v.(string)
		}
		if v, ok := a["srcmac"]; ok {
			acl["srcmac"] = v.(string)
		}
		if v, ok := a["srcportop"]; ok {
			acl["srcportop"] = v.(string)
		}
		if v, ok := a["srcportval"]; ok {
			acl["srcportval"] = v.(string)
		}
		if v, ok := a["state"]; ok {
			acl["state"] = v.(string)
		}
		if v, ok := a["td"]; ok {
			switch val := v.(type) {
			case int:
				acl["td"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["td"] = intVal
				}
			}
		}
		if v, ok := a["ttl"]; ok {
			switch val := v.(type) {
			case int:
				acl["ttl"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["ttl"] = intVal
				}
			}
		}
		if v, ok := a["vlan"]; ok {
			switch val := v.(type) {
			case int:
				acl["vlan"] = val
			case string:
				var intVal int
				if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
					acl["vlan"] = intVal
				}
			}
		}
		acls[i] = acl
	}
	d.Set("acl", acls)

	// Reset the trigger to "No" after read operations so that subsequent plans
	// will detect a change when users set acls_apply_trigger = "Yes".
	// This creates a toggle mechanism: "No" (default) -> "Yes" (user sets) triggers update,
	// then reset back to "No" (read function) for the next potential trigger cycle.
	// So, that the "nsacls" will be applied in every run when the user value is "Yes"
	d.Set("acls_apply_trigger", "No")

	return nil
}

func updateNsaclsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In updateNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	if d.HasChange("acl") {
		orig, noo := d.GetChange("acl")
		if orig == nil {
			orig = new(schema.Set)
		}
		if noo == nil {
			noo = new(schema.Set)
		}
		oset := orig.(*schema.Set)
		nset := noo.(*schema.Set)

		remove := oset.Difference(nset)
		add := nset.Difference(oset)
		log.Printf("[DEBUG]  netscaler-provider: need to remove %d acls", remove.Len())
		log.Printf("[DEBUG]  netscaler-provider: need to add %d acls", add.Len())
		//if the same acl is in add as well as remove, then use update

		update := make([]ns.Nsacl, remove.Len())
		rr := make([]map[string]interface{}, remove.Len())
		ar := make([]map[string]interface{}, remove.Len())
		u := 0
		for _, r := range remove.List() {
			rname := r.(map[string]interface{})["aclname"].(string)
			for _, a := range add.List() {
				aname := a.(map[string]interface{})["aclname"].(string)
				if rname == aname {
					ar[u] = a.(map[string]interface{})
					rr[u] = r.(map[string]interface{})
					log.Printf("[DEBUG] netscaler-provider: updateNsAcls: old acl is %v", rr[u])
					log.Printf("[DEBUG] netscaler-provider: updateNsAcls: new acl is %v", ar[u])
					nsacl := ns.Nsacl{Aclname: aname}
					diffMap := make(map[string]interface{}, len(ar[u]))
					for k, v := range ar[u] {
						if v != rr[u][k] {
							diffMap[k] = v
						}
					}
					mapstructure.Decode(diffMap, &nsacl)
					nsacl.Aclname = aname
					update[u] = nsacl
					u++
				}
			}
		}
		for _, r := range rr {
			remove.Remove(r)
		}
		for _, upd := range ar {
			add.Remove(upd)
		}

		for _, val := range update {
			acl := val
			log.Printf("[DEBUG]  netscaler-provider: going to update acl %s", acl.Aclname)
			//_, err := client.UpdateResource(service.Nsacl.Type(), acl["aclname"].(string), &acl)
			err := updateSingleAcl(acl, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error updating acl %s", acl.Aclname)
			}
		}

		for _, val := range remove.List() {
			acl := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to delete acl %s", acl["aclname"].(string))
			err := deleteSingleAcl(acl, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error deleting acl %s", acl["aclname"].(string))
			}
		}

		for _, val := range add.List() {
			acl := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to add acl %s", acl["aclname"].(string))
			err := createSingleAcl(acl, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error adding acl %s", acl["aclname"].(string))
			}
		}
	}
	nsacls := ns.Nsacls{}
	err := client.ApplyResource(service.Nsacls.Type(), &nsacls)

	return diag.FromErr(err)
}

func deleteNsaclsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteNsaclsFunc")
	acls := d.Get("acl").(*schema.Set).List()

	for _, val := range acls {
		acl := val.(map[string]interface{})
		_ = deleteSingleAcl(acl, meta)
	}
	client := meta.(*NetScalerNitroClient).client
	nsacls := ns.Nsacls{}
	err := client.ApplyResource(service.Nsacls.Type(), &nsacls)
	d.SetId("")
	return diag.FromErr(err)
}

func createSingleAcl(acl map[string]interface{}, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createSingleFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsaclName string
	if v, ok := acl["aclname"]; ok {
		nsaclName = v.(string)
	} else {
		nsaclName = resource.PrefixedUniqueId("tf-nsacl-")
		acl["aclname"] = nsaclName
	}
	destip := false
	destport := false
	srcip := false
	srcport := false
	if acl["destipval"] != nil && acl["destipval"] != "" {
		destip = true
	}
	if acl["destportval"] != nil && acl["destportval"] != "" {
		destport = true
	}
	if acl["srcipval"] != nil && acl["srcipval"] != "" {
		srcip = true
	}
	if acl["srcportval"] != nil && acl["srcportval"] != "" {
		srcport = true
	}

	if acl["destipop"] != nil && acl["destipval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if acl["destportop"] != nil && acl["destportval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have destportop without destportval", nsaclName)
	}
	if acl["srcipop"] != nil && acl["srcipval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcipop without srcipval", nsaclName)
	}
	if acl["srcportop"] != nil && acl["srcportval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcportop without srcportval", nsaclName)
	}

	nsacl := ns.Nsacl{
		Aclaction:      acl["aclaction"].(string),
		Aclname:        acl["aclname"].(string),
		Destip:         destip,
		Destipop:       acl["destipop"].(string),
		Destipval:      acl["destipval"].(string),
		Destport:       destport,
		Destportop:     acl["destportop"].(string),
		Destportval:    acl["destportval"].(string),
		Established:    acl["established"].(bool),
		Icmpcode:       acl["icmpcode"].(int),
		Icmptype:       acl["icmptype"].(int),
		Interface:      acl["interface"].(string),
		Logstate:       acl["logstate"].(string),
		Priority:       acl["priority"].(int),
		Protocol:       acl["protocol"].(string),
		Protocolnumber: acl["protocolnumber"].(int),
		Ratelimit:      acl["ratelimit"].(int),
		Srcip:          srcip,
		Srcipop:        acl["srcipop"].(string),
		Srcipval:       acl["srcipval"].(string),
		Srcmac:         acl["srcmac"].(string),
		Srcport:        srcport,
		Srcportop:      acl["srcportop"].(string),
		Srcportval:     acl["srcportval"].(string),
		State:          acl["state"].(string),
		Td:             acl["td"].(int),
		Ttl:            acl["ttl"].(int),
		Vlan:           acl["vlan"].(int),
	}

	_, err := client.AddResource(service.Nsacl.Type(), nsaclName, &nsacl)
	if err != nil {
		return err
	}

	return nil
}

func deleteSingleAcl(acl map[string]interface{}, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteSingleAcl")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := acl["aclname"].(string)
	err := client.DeleteResource(service.Nsacl.Type(), nsaclName)
	if err != nil {
		return err
	}

	return nil
}

func updateSingleAcl(acl ns.Nsacl, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateSingleAcl")
	client := meta.(*NetScalerNitroClient).client

	acl.Destip = acl.Destipval != ""
	acl.Srcip = acl.Srcipval != ""
	acl.Destport = acl.Destportval != ""
	acl.Srcport = acl.Srcportval != ""

	nsaclName := acl.Aclname

	if acl.Destipop != "" && acl.Destipval == "" {
		return fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if acl.Srcipop != "" && acl.Srcipval == "" {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcipop without srcipval", nsaclName)
	}
	if acl.Destportop != "" && acl.Destportval == "" {
		return fmt.Errorf("Error in nsacl spec %s cannot have destportop without destportval", nsaclName)
	}
	if acl.Srcportop != "" && acl.Srcportval == "" {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcportop without srcportval", nsaclName)
	}

	_, err := client.UpdateResource(service.Nsacl.Type(), nsaclName, &acl)

	return err
}

func validateAclAction(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	validActions := []string{"Yes", "No"}
	for _, validAction := range validActions {
		if value == validAction {
			return nil, nil
		}
	}
	errors = append(errors, fmt.Errorf(
		"%q must be one of %v (case-sensitive). Received: %q", k, validActions, value))
	return warnings, errors
}
