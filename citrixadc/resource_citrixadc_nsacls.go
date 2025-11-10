package citrixadc

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsacls() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsaclsFunc,
		UpdateContext: updateNsaclsFunc,
		ReadContext:   readNsaclsFunc,
		DeleteContext: deleteNsaclsFunc,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
						"srcportdataset": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcipdataset": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destportdataset": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destipdataset": {
							Type:     schema.TypeString,
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

	// Try to use GetRawConfig to access user-specified values, fallback to d.Get if not available
	rawConfig := d.GetRawConfig()
	var aclsProcessed bool

	if !rawConfig.IsNull() && rawConfig.IsKnown() && rawConfig.Type().HasAttribute("acl") {
		aclsAttr := rawConfig.GetAttr("acl")
		if !aclsAttr.IsNull() && aclsAttr.IsKnown() {
			aclsList := aclsAttr.AsValueSet().Values()
			for _, aclValue := range aclsList {
				if !aclValue.IsNull() && aclValue.IsKnown() {
					acl := constructAclFromCtyValue(aclValue)
					log.Printf("[DEBUG] netscaler-provider: creating acl from raw config: %+v", acl)

					_ = createSingleAcl(acl, meta)
				}
			}
			aclsProcessed = true
		}
	}

	// Fallback to regular d.Get() if raw config is not available or doesn't contain acl
	if !aclsProcessed {
		acls := d.Get("acl").(*schema.Set).List()
		for _, val := range acls {
			acl := val.(map[string]interface{})
			log.Printf("[DEBUG] netscaler-provider: creating acl from d.Get: %+v", acl)

			_ = createSingleAcl(acl, meta)
		}
	}

	nsacls := ns.Nsacls{
		Type: d.Get("type").(string),
	}
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
		if v, ok := a["destportdataset"]; ok {
			acl["destportdataset"] = v.(string)
		}
		if v, ok := a["destipdataset"]; ok {
			acl["destipdataset"] = v.(string)
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
		if v, ok := a["srcportdataset"]; ok {
			acl["srcportdataset"] = v.(string)
		}
		if v, ok := a["srcipdataset"]; ok {
			acl["srcipdataset"] = v.(string)
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
	d.Set("type", d.Get("type"))

	// Reset the trigger to "No" after read operations so that subsequent plans
	// will detect a change when users set acls_apply_trigger = "Yes".
	// This creates a toggle mechanism: "No" (default) -> "Yes" (user sets) triggers update,
	// then reset back to "No" (read function) for the next potential trigger cycle.
	// So, that the "nsacls" will be applied in every run when the user value is "Yes"
	d.Set("acls_apply_trigger", "No")

	return nil
}

func updateNsaclsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider: In updateNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client

	if d.HasChange("acl") {
		// Try to get ACLs from raw config for more accurate diff calculation
		var oldAcls, newAcls []map[string]interface{}
		var rawConfigAvailable bool

		// Try to get ACLs from raw config for accurate diff calculation
		// We need to get the old state from actual Terraform state and new state from raw plan
		rawPlan := d.GetRawPlan()

		if !rawPlan.IsNull() {
			// Get new ACLs from raw plan (user-specified values only)
			newAcls = getAclsFromRawConfig(rawPlan)

			// Get old ACLs from the previous state (also user-specified values only)
			// We use d.GetChange() but filter out computed/default values
			if origData, _ := d.GetChange("acl"); origData != nil {
				for _, val := range origData.(*schema.Set).List() {
					oldAclFull := val.(map[string]interface{})
					// Filter to only user-specified values (similar to raw config)
					oldAclFiltered := filterToUserSpecifiedValues(oldAclFull)
					oldAcls = append(oldAcls, oldAclFiltered)
				}
			}
			rawConfigAvailable = true
			log.Printf("[DEBUG] netscaler-provider: Using raw config for diff calculation")
		}

		// Fallback to traditional d.GetChange if raw config is not available
		if !rawConfigAvailable {
			log.Printf("[DEBUG] netscaler-provider: Falling back to d.GetChange for diff calculation")
			orig, noo := d.GetChange("acl")
			if orig != nil {
				for _, val := range orig.(*schema.Set).List() {
					oldAcls = append(oldAcls, val.(map[string]interface{}))
				}
			}
			if noo != nil {
				for _, val := range noo.(*schema.Set).List() {
					newAcls = append(newAcls, val.(map[string]interface{}))
				}
			}
		}

		// Create name sets for efficient lookup
		oldNamesSet := getAclNamesSet(oldAcls)
		newNamesSet := getAclNamesSet(newAcls)

		// Find ACLs to remove (in old but not in new)
		var toRemove []map[string]interface{}
		for _, oldAcl := range oldAcls {
			if aclName, exists := oldAcl["aclname"]; exists {
				if name, ok := aclName.(string); ok {
					if !newNamesSet[name] {
						toRemove = append(toRemove, oldAcl)
						log.Printf("[DEBUG] netscaler-provider: ACL to remove: %s", name)
					}
				}
			}
		}

		// Find ACLs to add (in new but not in old)
		var toAdd []map[string]interface{}
		for _, newAcl := range newAcls {
			if aclName, exists := newAcl["aclname"]; exists {
				if name, ok := aclName.(string); ok {
					if !oldNamesSet[name] {
						toAdd = append(toAdd, newAcl)
						log.Printf("[DEBUG] netscaler-provider: ACL to add: %s", name)
					}
				}
			}
		}

		// Find ACLs to update (exist in both but with different values)
		var toUpdate []map[string]interface{}
		for _, newAcl := range newAcls {
			if aclName, exists := newAcl["aclname"]; exists {
				if name, ok := aclName.(string); ok {
					if oldNamesSet[name] {
						// ACL exists in both old and new, check if it has changed
						if oldAcl, found := findAclByName(oldAcls, name); found {
							if !aclMapsEqual(oldAcl, newAcl) {
								toUpdate = append(toUpdate, newAcl)
								log.Printf("[DEBUG] netscaler-provider: ACL to update: %s", name)
								log.Printf("[DEBUG] netscaler-provider: Old ACL: %+v", oldAcl)
								log.Printf("[DEBUG] netscaler-provider: New ACL: %+v", newAcl)
							}
						}
					}
				}
			}
		}

		log.Printf("[DEBUG] netscaler-provider: Need to remove %d ACLs", len(toRemove))
		log.Printf("[DEBUG] netscaler-provider: Need to add %d ACLs", len(toAdd))
		log.Printf("[DEBUG] netscaler-provider: Need to update %d ACLs", len(toUpdate))

		// Process removals
		for _, acl := range toRemove {
			if aclName, exists := acl["aclname"]; exists {
				log.Printf("[DEBUG] netscaler-provider: Deleting ACL %s", aclName)
				err := deleteSingleAcl(acl, meta)
				if err != nil {
					log.Printf("[DEBUG] netscaler-provider: Error deleting ACL %s: %v", aclName, err)
				}
			}
		}

		// Process updates
		for _, acl := range toUpdate {
			if aclName, exists := acl["aclname"]; exists {
				log.Printf("[DEBUG] netscaler-provider: Creating (updating) ACL %s", aclName)
				err := createSingleAcl(acl, meta)
				if err != nil {
					log.Printf("[DEBUG] netscaler-provider: Error updating ACL %s: %v", aclName, err)
				}
			}
		}

		// Process additions
		for _, acl := range toAdd {
			if aclName, exists := acl["aclname"]; exists {
				log.Printf("[DEBUG] netscaler-provider: Adding ACL %s", aclName)
				err := createSingleAcl(acl, meta)
				if err != nil {
					log.Printf("[DEBUG] netscaler-provider: Error adding ACL %s: %v", aclName, err)
				}
			}
		}
	}

	nsacls := ns.Nsacls{}
	err := client.ApplyResource(service.Nsacls.Type(), &nsacls)

	return diag.FromErr(err)
}

func deleteNsaclsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteNsaclsFunc")

	// Try to use GetRawConfig to access user-specified values, fallback to d.Get if not available
	rawConfig := d.GetRawConfig()
	var aclsProcessed bool

	if !rawConfig.IsNull() && rawConfig.IsKnown() && rawConfig.Type().HasAttribute("acl") {
		aclsAttr := rawConfig.GetAttr("acl")
		if !aclsAttr.IsNull() && aclsAttr.IsKnown() {
			aclsList := aclsAttr.AsValueSet().Values()
			for _, aclValue := range aclsList {
				if !aclValue.IsNull() && aclValue.IsKnown() {
					acl := constructAclFromCtyValue(aclValue)
					_ = deleteSingleAcl(acl, meta)
				}
			}
			aclsProcessed = true
		}
	}

	// Fallback to regular d.Get() if raw config is not available or doesn't contain acl
	if !aclsProcessed {
		acls := d.Get("acl").(*schema.Set).List()
		for _, val := range acls {
			acl := val.(map[string]interface{})
			_ = deleteSingleAcl(acl, meta)
		}
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
	if acl["destipval"] != nil && acl["destipval"] != "" || acl["destipdataset"] != nil && acl["destipdataset"] != "" {
		destip = true
	}
	if acl["destportval"] != nil && acl["destportval"] != "" || acl["destportdataset"] != nil && acl["destportdataset"] != "" {
		destport = true
	}
	if acl["srcipval"] != nil && acl["srcipval"] != "" || acl["srcipdataset"] != nil && acl["srcipdataset"] != "" {
		srcip = true
	}
	if acl["srcportval"] != nil && acl["srcportval"] != "" || acl["srcportdataset"] != nil && acl["srcportdataset"] != "" {
		srcport = true
	}

	if acl["destipop"] != nil && acl["destipval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if acl["destportop"] != nil && acl["destipval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if acl["srcipop"] != nil && acl["srcipval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcipop without srcipval", nsaclName)
	}
	if acl["srcportop"] != nil && acl["srcportval"] == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcportop without srcportval", nsaclName)
	}

	nsacl := ns.Nsacl{
		Aclname:  acl["aclname"].(string),
		Destip:   destip,
		Destport: destport,
		Srcip:    srcip,
		Srcport:  srcport,
	}

	// Only set string fields if they were explicitly provided by the user
	if isStringValueSet(acl, "aclaction") {
		nsacl.Aclaction = acl["aclaction"].(string)
	}
	if isStringValueSet(acl, "destipop") {
		nsacl.Destipop = acl["destipop"].(string)
	}
	if isStringValueSet(acl, "destipval") {
		nsacl.Destipval = acl["destipval"].(string)
	}
	if isStringValueSet(acl, "destipdataset") {
		nsacl.Destipdataset = acl["destipdataset"].(string)
	}
	if isStringValueSet(acl, "destportop") {
		nsacl.Destportop = acl["destportop"].(string)
	}
	if isStringValueSet(acl, "destportval") {
		nsacl.Destportval = acl["destportval"].(string)
	}
	if isStringValueSet(acl, "destportdataset") {
		nsacl.Destportdataset = acl["destportdataset"].(string)
	}
	if isBoolValueSet(acl, "established") {
		nsacl.Established = acl["established"].(bool)
	}
	if isStringValueSet(acl, "interface") {
		nsacl.Interface = acl["interface"].(string)
	}
	if isStringValueSet(acl, "logstate") {
		nsacl.Logstate = acl["logstate"].(string)
	}
	if isStringValueSet(acl, "protocol") {
		nsacl.Protocol = acl["protocol"].(string)
	}
	if isStringValueSet(acl, "srcipop") {
		nsacl.Srcipop = acl["srcipop"].(string)
	}
	if isStringValueSet(acl, "srcipval") {
		nsacl.Srcipval = acl["srcipval"].(string)
	}
	if isStringValueSet(acl, "srcipdataset") {
		nsacl.Srcipdataset = acl["srcipdataset"].(string)
	}
	if isStringValueSet(acl, "srcmac") {
		nsacl.Srcmac = acl["srcmac"].(string)
	}
	if isStringValueSet(acl, "srcportop") {
		nsacl.Srcportop = acl["srcportop"].(string)
	}
	if isStringValueSet(acl, "srcportval") {
		nsacl.Srcportval = acl["srcportval"].(string)
	}
	if isStringValueSet(acl, "srcportdataset") {
		nsacl.Srcportdataset = acl["srcportdataset"].(string)
	}
	if isStringValueSet(acl, "state") {
		nsacl.State = acl["state"].(string)
	}

	// Only set integer fields if they were explicitly provided by the user
	if isIntValueSet(acl, "icmpcode") {
		nsacl.Icmpcode = intPtr(acl["icmpcode"].(int))
	}
	if isIntValueSet(acl, "icmptype") {
		nsacl.Icmptype = intPtr(acl["icmptype"].(int))
	}
	if isIntValueSet(acl, "priority") {
		nsacl.Priority = intPtr(acl["priority"].(int))
	}
	if isIntValueSet(acl, "protocolnumber") {
		nsacl.Protocolnumber = intPtr(acl["protocolnumber"].(int))
	}
	if isIntValueSet(acl, "ratelimit") {
		nsacl.Ratelimit = intPtr(acl["ratelimit"].(int))
	}
	if isIntValueSet(acl, "td") {
		nsacl.Td = intPtr(acl["td"].(int))
	}
	if isIntValueSet(acl, "ttl") {
		nsacl.Ttl = intPtr(acl["ttl"].(int))
	}
	if isIntValueSet(acl, "vlan") {
		nsacl.Vlan = intPtr(acl["vlan"].(int))
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

// Helper function to construct ACL map from cty.Value (raw config)
func constructAclFromCtyValue(aclValue cty.Value) map[string]interface{} {
	acl := make(map[string]interface{})

	if !aclValue.IsNull() && aclValue.IsKnown() {
		aclMap := aclValue.AsValueMap()

		// Only include attributes that are explicitly set in the configuration
		if val, exists := aclMap["aclname"]; exists && !val.IsNull() && val.IsKnown() {
			acl["aclname"] = val.AsString()
		}
		if val, exists := aclMap["aclaction"]; exists && !val.IsNull() && val.IsKnown() {
			acl["aclaction"] = val.AsString()
		}
		if val, exists := aclMap["destipop"]; exists && !val.IsNull() && val.IsKnown() {
			acl["destipop"] = val.AsString()
		}
		if val, exists := aclMap["destipval"]; exists && !val.IsNull() && val.IsKnown() {
			acl["destipval"] = val.AsString()
		}
		if val, exists := aclMap["destipdataset"]; exists && !val.IsNull() && val.IsKnown() {
			acl["destipdataset"] = val.AsString()
		}
		if val, exists := aclMap["destportop"]; exists && !val.IsNull() && val.IsKnown() {
			acl["destportop"] = val.AsString()
		}
		if val, exists := aclMap["destportval"]; exists && !val.IsNull() && val.IsKnown() {
			acl["destportval"] = val.AsString()
		}
		if val, exists := aclMap["destportdataset"]; exists && !val.IsNull() && val.IsKnown() {
			acl["destportdataset"] = val.AsString()
		}
		if val, exists := aclMap["established"]; exists && !val.IsNull() && val.IsKnown() {
			acl["established"] = val.True()
		}
		if val, exists := aclMap["icmpcode"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["icmpcode"] = int(intVal)
			}
		}
		if val, exists := aclMap["icmptype"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["icmptype"] = int(intVal)
			}
		}
		if val, exists := aclMap["interface"]; exists && !val.IsNull() && val.IsKnown() {
			acl["interface"] = val.AsString()
		}
		if val, exists := aclMap["logstate"]; exists && !val.IsNull() && val.IsKnown() {
			acl["logstate"] = val.AsString()
		}
		if val, exists := aclMap["priority"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["priority"] = int(intVal)
			}
		}
		if val, exists := aclMap["protocol"]; exists && !val.IsNull() && val.IsKnown() {
			acl["protocol"] = val.AsString()
		}
		if val, exists := aclMap["protocolnumber"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["protocolnumber"] = int(intVal)
			}
		}
		if val, exists := aclMap["ratelimit"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["ratelimit"] = int(intVal)
			}
		}
		if val, exists := aclMap["srcipop"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcipop"] = val.AsString()
		}
		if val, exists := aclMap["srcipval"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcipval"] = val.AsString()
		}
		if val, exists := aclMap["srcipdataset"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcipdataset"] = val.AsString()
		}
		if val, exists := aclMap["srcmac"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcmac"] = val.AsString()
		}
		if val, exists := aclMap["srcportop"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcportop"] = val.AsString()
		}
		if val, exists := aclMap["srcportval"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcportval"] = val.AsString()
		}
		if val, exists := aclMap["srcportdataset"]; exists && !val.IsNull() && val.IsKnown() {
			acl["srcportdataset"] = val.AsString()
		}
		if val, exists := aclMap["state"]; exists && !val.IsNull() && val.IsKnown() {
			acl["state"] = val.AsString()
		}
		if val, exists := aclMap["td"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["td"] = int(intVal)
			}
		}
		if val, exists := aclMap["ttl"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["ttl"] = int(intVal)
			}
		}
		if val, exists := aclMap["vlan"]; exists && !val.IsNull() && val.IsKnown() {
			bigFloat := val.AsBigFloat()
			if intVal, accuracy := bigFloat.Int64(); accuracy == big.Exact {
				acl["vlan"] = int(intVal)
			}
		}
	}

	return acl
}

// Helper function to get ACLs from raw config, returns empty slice if not available
func getAclsFromRawConfig(rawConfig cty.Value) []map[string]interface{} {
	var acls []map[string]interface{}

	if !rawConfig.IsNull() && rawConfig.IsKnown() && rawConfig.Type().HasAttribute("acl") {
		aclsAttr := rawConfig.GetAttr("acl")
		if !aclsAttr.IsNull() && aclsAttr.IsKnown() {
			aclsList := aclsAttr.AsValueSet().Values()
			for _, aclValue := range aclsList {
				if !aclValue.IsNull() && aclValue.IsKnown() {
					acl := constructAclFromCtyValue(aclValue)
					acls = append(acls, acl)
				}
			}
		}
	}

	return acls
}

// Helper function to find ACL by name in a slice
func findAclByName(acls []map[string]interface{}, name string) (map[string]interface{}, bool) {
	for _, acl := range acls {
		if aclName, exists := acl["aclname"]; exists && aclName == name {
			return acl, true
		}
	}
	return nil, false
}

// Helper function to create a set of ACL names from ACL slice
func getAclNamesSet(acls []map[string]interface{}) map[string]bool {
	names := make(map[string]bool)
	for _, acl := range acls {
		if aclName, exists := acl["aclname"]; exists {
			if name, ok := aclName.(string); ok {
				names[name] = true
			}
		}
	}
	return names
}

// Helper function to compare two ACL maps for equality (only user-specified fields)
func aclMapsEqual(acl1, acl2 map[string]interface{}) bool {
	// Get all keys from both maps
	allKeys := make(map[string]bool)
	for k := range acl1 {
		allKeys[k] = true
	}
	for k := range acl2 {
		allKeys[k] = true
	}

	// Compare each key
	for key := range allKeys {
		val1, exists1 := acl1[key]
		val2, exists2 := acl2[key]

		// If key exists in only one map
		if exists1 != exists2 {
			return false
		}

		// If key exists in both maps, compare values
		if exists1 && exists2 {
			// Handle nil values
			if val1 == nil && val2 == nil {
				continue
			}
			if val1 == nil || val2 == nil {
				return false
			}

			// Compare values
			if val1 != val2 {
				return false
			}
		}
	}

	return true
}

// Helper function to filter ACL map to only user-specified values (removes default/computed values)
func filterToUserSpecifiedValues(acl map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})

	// Only include fields that have meaningful (non-default) values
	if val, exists := acl["aclname"]; exists && val != nil && val.(string) != "" {
		filtered["aclname"] = val
	}
	if val, exists := acl["aclaction"]; exists && val != nil && val.(string) != "" {
		filtered["aclaction"] = val
	}
	if val, exists := acl["destipop"]; exists && val != nil && val.(string) != "" {
		filtered["destipop"] = val
	}
	if val, exists := acl["destipval"]; exists && val != nil && val.(string) != "" {
		filtered["destipval"] = val
	}
	if val, exists := acl["destipdataset"]; exists && val != nil && val.(string) != "" {
		filtered["destipdataset"] = val
	}
	if val, exists := acl["destportop"]; exists && val != nil && val.(string) != "" {
		filtered["destportop"] = val
	}
	if val, exists := acl["destportval"]; exists && val != nil && val.(string) != "" {
		filtered["destportval"] = val
	}
	if val, exists := acl["destportdataset"]; exists && val != nil && val.(string) != "" {
		filtered["destportdataset"] = val
	}
	if val, exists := acl["established"]; exists && val != nil {
		// Only include if explicitly set to true (false is default)
		if boolVal, ok := val.(bool); ok && boolVal {
			filtered["established"] = val
		}
	}
	if val, exists := acl["interface"]; exists && val != nil && val.(string) != "" {
		filtered["interface"] = val
	}
	if val, exists := acl["logstate"]; exists && val != nil {
		// Only include if not default value
		if strVal, ok := val.(string); ok && strVal != "" && strVal != "DISABLED" {
			filtered["logstate"] = val
		}
	}
	if val, exists := acl["protocol"]; exists && val != nil && val.(string) != "" {
		filtered["protocol"] = val
	}
	if val, exists := acl["srcipop"]; exists && val != nil && val.(string) != "" {
		filtered["srcipop"] = val
	}
	if val, exists := acl["srcipval"]; exists && val != nil && val.(string) != "" {
		filtered["srcipval"] = val
	}
	if val, exists := acl["srcipdataset"]; exists && val != nil && val.(string) != "" {
		filtered["srcipdataset"] = val
	}
	if val, exists := acl["srcmac"]; exists && val != nil && val.(string) != "" {
		filtered["srcmac"] = val
	}
	if val, exists := acl["srcportop"]; exists && val != nil && val.(string) != "" {
		filtered["srcportop"] = val
	}
	if val, exists := acl["srcportval"]; exists && val != nil && val.(string) != "" {
		filtered["srcportval"] = val
	}
	if val, exists := acl["srcportdataset"]; exists && val != nil && val.(string) != "" {
		filtered["srcportdataset"] = val
	}
	if val, exists := acl["state"]; exists && val != nil {
		// Only include if not default value
		if strVal, ok := val.(string); ok && strVal != "" && strVal != "ENABLED" {
			filtered["state"] = val
		}
	}

	// Handle integer fields - only include if they have meaningful values
	intFields := map[string]interface{}{
		"icmpcode":       65536, // Default value to exclude
		"icmptype":       65536, // Default value to exclude
		"priority":       -1,    // No specific default, include if > 0
		"protocolnumber": 0,     // Default value to exclude
		"ratelimit":      100,   // Default value to exclude
		"td":             0,     // Default value to exclude
		"ttl":            0,     // Default value to exclude
		"vlan":           -1,    // No specific default, include if > 0
	}

	for field, defaultVal := range intFields {
		if val, exists := acl[field]; exists && val != nil {
			if intVal, ok := val.(int); ok {
				// Include field if it's not the default value
				if field == "priority" || field == "vlan" {
					// For priority and vlan, include if > 0
					if intVal > 0 {
						filtered[field] = val
					}
				} else {
					// For other fields, include if not equal to default
					if intVal != defaultVal.(int) {
						filtered[field] = val
					}
				}
			}
		}
	}

	return filtered
}

// Helper function to check if a string value was explicitly set in the nested schema
func isStringValueSet(acl map[string]interface{}, key string) bool {
	if val, exists := acl[key]; exists && val != nil {
		if str, ok := val.(string); ok && str != "" {
			return true
		}
	}
	return false
}

// Helper function to check if an int value was explicitly set in the nested schema
func isIntValueSet(acl map[string]interface{}, key string) bool {
	if val, exists := acl[key]; exists && val != nil {
		if _, ok := val.(int); ok {
			return true
		}
	}
	return false
}

// Helper function to check if a bool value was explicitly set in the nested schema
func isBoolValueSet(acl map[string]interface{}, key string) bool {
	if val, exists := acl[key]; exists && val != nil {
		if _, ok := val.(bool); ok {
			return true
		}
	}
	return false
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
