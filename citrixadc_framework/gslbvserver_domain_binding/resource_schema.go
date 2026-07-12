package gslbvserver_domain_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbvserverDomainBindingResourceModel describes the resource data model.
type GslbvserverDomainBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Backupip         types.String `tfsdk:"backupip"`
	Backupipflag     types.Bool   `tfsdk:"backupipflag"`
	CookieDomain     types.String `tfsdk:"cookiedomain"`
	CookieDomainflag types.Bool   `tfsdk:"cookiedomainflag"`
	Cookietimeout    types.Int64  `tfsdk:"cookietimeout"`
	Domainname       types.String `tfsdk:"domainname"`
	Name             types.String `tfsdk:"name"`
	Order            types.Int64  `tfsdk:"order"`
	Sitedomainttl    types.Int64  `tfsdk:"sitedomainttl"`
	Ttl              types.Int64  `tfsdk:"ttl"`
}

func (r *GslbvserverDomainBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbvserver_domain_binding resource.",
			},
			"backupip": schema.StringAttribute{
				// backupip is ForceNew in SDK v2 and is accepted by the NITRO bind payload and
				// echoed by GET (key "backupip") when set. Class-B upgrade fix (Row-4): add
				// Computed so a value stored by SDK v2 2.2.0 (which stored "" when unset) is
				// absorbed instead of diffing a null config against the stored value on a
				// RequiresReplace attribute (which would destroy+recreate the binding on upgrade).
				// UseStateForUnknown MUST be listed BEFORE RequiresReplace: on update the
				// framework marks a Computed null-config attribute unknown, USFU then restores
				// the prior-state value so RequiresReplace sees plan==state and does NOT replace;
				// a genuine user change to a new value still trips RequiresReplace afterwards.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"backupipflag": schema.BoolAttribute{
				// Variant (b) of the Class-B upgrade fix. Verified empirically against the ADC:
				// the NITRO bind payload REJECTS backupipflag with errorcode 278 and GET NEVER
				// echoes it back. It therefore cannot be Computed (a Computed value would stay
				// unknown after apply -> "inconsistent result after apply ... still unknown").
				// Row-4 (variant a) is impossible here, so we keep it Optional-only and DROP
				// RequiresReplace: that prevents an SDK v2 upgrade (stored bool vs a null config)
				// from force-replacing the binding. It is deliberately NOT read in the resource
				// setter (GET has nothing to echo), so planned==final on every apply.
				Optional:    true,
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"cookiedomain": schema.StringAttribute{
				// cookiedomain is ForceNew in SDK v2 and is accepted by the NITRO bind payload
				// (JSON key "cookie_domain") and echoed by GET (key "cookie_domain") when set.
				// Class-B upgrade fix (Row-4): see the backupip comment above for the ordering
				// rationale of UseStateForUnknown before RequiresReplace.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.",
			},
			"cookiedomainflag": schema.BoolAttribute{
				// Variant (b), same rationale as backupipflag: the NITRO bind payload rejects it
				// (errorcode 278) and GET never echoes it, so it cannot be Computed. Kept
				// Optional-only with RequiresReplace dropped. Unlike backupipflag (which the
				// acctest configures to false), this one is left unset in config, so the resource
				// setter explicitly NULLs it to keep the SDK v2 upgrade plan a no-op (see setter).
				Optional:    true,
				Description: "The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in minutes, for the GSLB site cookie.",
			},
			"domainname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				// order requires an accompanying serviceName (errorcode 1093) and is not
				// applicable to a domain binding; GET does not echo it. Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"sitedomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to live (TTL) for the domain.",
			},
		},
	}
}

func gslbvserver_domain_bindingGetThePayloadFromthePlan(ctx context.Context, data *GslbvserverDomainBindingResourceModel) gslb.Gslbvserverdomainbinding {
	tflog.Debug(ctx, "In gslbvserver_domain_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	gslbvserver_domain_binding := gslb.Gslbvserverdomainbinding{}
	if !data.Backupip.IsNull() && !data.Backupip.IsUnknown() {
		gslbvserver_domain_binding.Backupip = data.Backupip.ValueString()
	}
	// backupipflag and cookiedomainflag are GET-response-only flags; the NITRO add/bind
	// endpoint rejects them with errorcode 278 ("Invalid argument"). They are never sent
	// in the payload. (Pattern 15)
	if !data.CookieDomain.IsNull() && !data.CookieDomain.IsUnknown() {
		gslbvserver_domain_binding.Cookiedomain = data.CookieDomain.ValueString()
	}
	if !data.Cookietimeout.IsNull() && !data.Cookietimeout.IsUnknown() {
		gslbvserver_domain_binding.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Domainname.IsNull() && !data.Domainname.IsUnknown() {
		gslbvserver_domain_binding.Domainname = data.Domainname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		gslbvserver_domain_binding.Name = data.Name.ValueString()
	}
	// order is not applicable to a domain binding (NITRO requires an accompanying
	// serviceName, errorcode 1093); never sent in the payload.
	if !data.Sitedomainttl.IsNull() && !data.Sitedomainttl.IsUnknown() {
		gslbvserver_domain_binding.Sitedomainttl = utils.IntPtr(int(data.Sitedomainttl.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		gslbvserver_domain_binding.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return gslbvserver_domain_binding
}

// gslbvserver_domain_bindingSetAttrFromGet is used by the resource Read/Create/Update.
// backupip/cookiedomain are resolved from the GET response (defaulting to "" when absent);
// backupipflag is preserved from the plan/state (config-supplied); cookiedomainflag is nulled;
// order is left as-is (server-non-echoed). It re-derives the canonical new-format ID so a legacy
// SDK v2 id is upgraded on Read. (Pattern 6 / Pattern 7 / Pattern 13)
func gslbvserver_domain_bindingSetAttrFromGet(ctx context.Context, data *GslbvserverDomainBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverDomainBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_domain_bindingSetAttrFromGet Function")

	// Convert API response to model.
	// backupip / cookiedomain are Optional+Computed with RequiresReplace+UseStateForUnknown.
	// NITRO GET echoes them (keys "backupip" / "cookie_domain") ONLY when they hold a value;
	// when unset they are absent from the response. We must still resolve the Computed value
	// to a concrete, NON-null value ("") on the not-echoed path, otherwise the Create plan's
	// unknown would remain unknown ("inconsistent result after apply") and UseStateForUnknown
	// (which ignores a null prior state) could not keep subsequent plans empty.
	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	} else {
		data.Backupip = types.StringValue("")
	}
	if val, ok := getResponseData["cookie_domain"]; ok && val != nil {
		data.CookieDomain = types.StringValue(val.(string))
	} else {
		data.CookieDomain = types.StringValue("")
	}
	// backupipflag / cookiedomainflag: the NITRO bind payload rejects them (errorcode 278) and
	// GET never echoes them, so there is nothing to read. Variant (b) of the Class-B upgrade fix
	// (Optional-only, RequiresReplace dropped) is asymmetric between the two flags because the
	// acctest config supplies backupipflag=false but leaves cookiedomainflag unset:
	//   - backupipflag is left UNTOUCHED. Its planned value comes from config (false); nulling it
	//     would make planned(false) != final(null) ("inconsistent result after apply"). On the
	//     SDK v2 upgrade the stored false already equals the configured false, so no diff.
	//   - cookiedomainflag is explicitly NULLED. Config leaves it null, so on the SDK v2 upgrade
	//     the value stored by 2.2.0 (false) would otherwise diff false->null. Because the domain
	//     binding cannot be updated in place (a re-bind of an already-bound domain fails with
	//     errorcode 1842), that phantom diff must be eliminated: nulling it here makes the
	//     refreshed state (null) match the config (null) so the upgrade plan is a clean no-op,
	//     and on create planned(null)==final(null).
	data.CookieDomainflag = types.BoolNull()
	if val, ok := getResponseData["cookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookietimeout = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["domainname"]; ok && val != nil {
		data.Domainname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["sitedomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sitedomainttl = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.UrlEncode(data.Domainname.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// gslbvserver_domain_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response (the datasource has no prior plan/state to preserve) and sets the
// composite ID (legacy SDK v2 order: name,domainname) since the datasource has no Create.
// (Pattern 7)
func gslbvserver_domain_bindingSetAttrFromGetForDatasource(ctx context.Context, data *GslbvserverDomainBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverDomainBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_domain_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	} else {
		data.Backupip = types.StringNull()
	}
	if val, ok := getResponseData["backupipflag"]; ok && val != nil {
		data.Backupipflag = types.BoolValue(val.(bool))
	} else {
		data.Backupipflag = types.BoolNull()
	}
	if val, ok := getResponseData["cookiedomain"]; ok && val != nil {
		data.CookieDomain = types.StringValue(val.(string))
	} else {
		data.CookieDomain = types.StringNull()
	}
	if val, ok := getResponseData["cookiedomainflag"]; ok && val != nil {
		data.CookieDomainflag = types.BoolValue(val.(bool))
	} else {
		data.CookieDomainflag = types.BoolNull()
	}
	if val, ok := getResponseData["cookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookietimeout = types.Int64Value(intVal)
		}
	} else {
		data.Cookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["domainname"]; ok && val != nil {
		data.Domainname = types.StringValue(val.(string))
	} else {
		data.Domainname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["sitedomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sitedomainttl = types.Int64Value(intVal)
		}
	} else {
		data.Sitedomainttl = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}

	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.UrlEncode(data.Domainname.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
