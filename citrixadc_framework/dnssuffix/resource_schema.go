package dnssuffix

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnssuffixResourceModel describes the resource data model.
type DnssuffixResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Dnssuffix types.String `tfsdk:"dnssuffix"`
}

func (r *DnssuffixResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnssuffix resource.",
			},
			"dnssuffix": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Suffix to be appended when resolving domain names that are not fully qualified.",
			},
		},
	}
}

func dnssuffixGetThePayloadFromtheConfig(ctx context.Context, data *DnssuffixResourceModel) dns.Dnssuffix {
	tflog.Debug(ctx, "In dnssuffixGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnssuffix := dns.Dnssuffix{}
	if !data.Dnssuffix.IsNull() {
		dnssuffix.Dnssuffix = data.Dnssuffix.ValueString()
	}

	return dnssuffix
}

func dnssuffixSetAttrFromGet(ctx context.Context, data *DnssuffixResourceModel, getResponseData map[string]interface{}) *DnssuffixResourceModel {
	tflog.Debug(ctx, "In dnssuffixSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Dnssuffix"]; ok && val != nil {
		data.Dnssuffix = types.StringValue(val.(string))
	} else {
		data.Dnssuffix = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("dnssuffix-config")

	return data
}
