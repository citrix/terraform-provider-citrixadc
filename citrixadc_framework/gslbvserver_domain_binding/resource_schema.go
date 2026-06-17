package gslbvserver_domain_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"backupipflag": schema.BoolAttribute{
				// GET-response-only flag; not accepted by the NITRO bind payload (errorcode 278).
				// Optional only (no Computed): GET does not echo it back, so a Computed value
				// would stay unknown after apply ("inconsistent result"). (Pattern 13 / Pattern 15)
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"cookiedomain": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.",
			},
			"cookiedomainflag": schema.BoolAttribute{
				// GET-response-only flag; not accepted by the NITRO bind payload (errorcode 278).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
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
// It preserves the existing plan/state values for write-only / server-non-echoed inputs
// (backupipflag, cookiedomainflag, order) and never recomputes the ID — the ID is set
// exactly once in Create. (Pattern 6 / Pattern 7 / Pattern 13)
func gslbvserver_domain_bindingSetAttrFromGet(ctx context.Context, data *GslbvserverDomainBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverDomainBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_domain_bindingSetAttrFromGet Function")

	// Convert API response to model.
	// backupipflag / cookiedomainflag / order are write-only inputs the GET does not
	// echo back reliably; preserve the configured plan/state value to avoid
	// "inconsistent result after apply" churn.
	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["cookiedomain"]; ok && val != nil {
		data.CookieDomain = types.StringValue(val.(string))
	}
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

	// ID is set once in Create; do not recompute it here.
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
