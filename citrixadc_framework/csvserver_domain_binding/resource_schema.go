package csvserver_domain_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CsvserverDomainBindingResourceModel describes the resource data model.
type CsvserverDomainBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Backupip      types.String `tfsdk:"backupip"`
	Cookiedomain  types.String `tfsdk:"cookiedomain"`
	Cookietimeout types.Int64  `tfsdk:"cookietimeout"`
	Domainname    types.String `tfsdk:"domainname"`
	Name          types.String `tfsdk:"name"`
	Sitedomainttl types.Int64  `tfsdk:"sitedomainttl"`
	Ttl           types.Int64  `tfsdk:"ttl"`
}

func (r *CsvserverDomainBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_domain_binding resource.",
			},
			"backupip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
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
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"sitedomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
		},
	}
}

func csvserver_domain_bindingGetThePayloadFromthePlan(ctx context.Context, data *CsvserverDomainBindingResourceModel) cs.Csvserverdomainbinding {
	tflog.Debug(ctx, "In csvserver_domain_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	csvserver_domain_binding := cs.Csvserverdomainbinding{}
	if !data.Backupip.IsNull() && !data.Backupip.IsUnknown() {
		csvserver_domain_binding.Backupip = data.Backupip.ValueString()
	}
	if !data.Cookiedomain.IsNull() && !data.Cookiedomain.IsUnknown() {
		csvserver_domain_binding.Cookiedomain = data.Cookiedomain.ValueString()
	}
	if !data.Cookietimeout.IsNull() && !data.Cookietimeout.IsUnknown() {
		csvserver_domain_binding.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Domainname.IsNull() && !data.Domainname.IsUnknown() {
		csvserver_domain_binding.Domainname = data.Domainname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		csvserver_domain_binding.Name = data.Name.ValueString()
	}
	if !data.Sitedomainttl.IsNull() && !data.Sitedomainttl.IsUnknown() {
		csvserver_domain_binding.Sitedomainttl = utils.IntPtr(int(data.Sitedomainttl.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		csvserver_domain_binding.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return csvserver_domain_binding
}

func csvserver_domain_bindingSetAttrFromGet(ctx context.Context, data *CsvserverDomainBindingResourceModel, getResponseData map[string]interface{}) *CsvserverDomainBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_domain_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	} else {
		data.Backupip = types.StringNull()
	}
	if val, ok := getResponseData["cookiedomain"]; ok && val != nil {
		data.Cookiedomain = types.StringValue(val.(string))
	} else {
		data.Cookiedomain = types.StringNull()
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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Domainname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
