package gslbvserver_domain_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbvserverDomainBindingResourceModel describes the resource data model.
type GslbvserverDomainBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Backupip         types.String `tfsdk:"backupip"`
	Backupipflag     types.Bool   `tfsdk:"backupipflag"`
	CookieDomain     types.String `tfsdk:"cookie_domain"`
	CookieDomainflag types.Bool   `tfsdk:"cookie_domainflag"`
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
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"backupipflag": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"cookie_domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.",
			},
			"cookie_domainflag": schema.BoolAttribute{
				Optional: true,
				Computed: true,
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
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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

func gslbvserver_domain_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbvserverDomainBindingResourceModel) gslb.Gslbvserverdomainbinding {
	tflog.Debug(ctx, "In gslbvserver_domain_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbvserver_domain_binding := gslb.Gslbvserverdomainbinding{}
	if !data.Backupip.IsNull() {
		gslbvserver_domain_binding.Backupip = data.Backupip.ValueString()
	}
	if !data.Backupipflag.IsNull() {
		gslbvserver_domain_binding.Backupipflag = data.Backupipflag.ValueBool()
	}
	if !data.CookieDomain.IsNull() {
		gslbvserver_domain_binding.Cookiedomain = data.CookieDomain.ValueString()
	}
	if !data.CookieDomainflag.IsNull() {
		gslbvserver_domain_binding.Cookiedomainflag = data.CookieDomainflag.ValueBool()
	}
	if !data.Cookietimeout.IsNull() {
		gslbvserver_domain_binding.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Domainname.IsNull() {
		gslbvserver_domain_binding.Domainname = data.Domainname.ValueString()
	}
	if !data.Name.IsNull() {
		gslbvserver_domain_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		gslbvserver_domain_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Sitedomainttl.IsNull() {
		gslbvserver_domain_binding.Sitedomainttl = utils.IntPtr(int(data.Sitedomainttl.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		gslbvserver_domain_binding.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return gslbvserver_domain_binding
}

func gslbvserver_domain_bindingSetAttrFromGet(ctx context.Context, data *GslbvserverDomainBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverDomainBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_domain_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["cookie_domain"]; ok && val != nil {
		data.CookieDomain = types.StringValue(val.(string))
	} else {
		data.CookieDomain = types.StringNull()
	}
	if val, ok := getResponseData["cookie_domainflag"]; ok && val != nil {
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("backupipflag:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Backupipflag.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("cookie_domainflag:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.CookieDomainflag.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Domainname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
