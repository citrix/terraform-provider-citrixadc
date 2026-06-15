package linkset_channel_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LinksetChannelBindingResourceModel describes the resource data model.
type LinksetChannelBindingResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Id    types.String `tfsdk:"id"`
	Ifnum types.String `tfsdk:"ifnum"`
}

func (r *LinksetChannelBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the linkset_channel_binding resource.",
			},
			"id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ID of the linkset to which to bind the interfaces.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The interfaces to be bound to the linkset.",
			},
		},
	}
}

func linkset_channel_bindingGetThePayloadFromthePlan(ctx context.Context, data *LinksetChannelBindingResourceModel) network.Linksetchannelbinding {
	tflog.Debug(ctx, "In linkset_channel_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	linkset_channel_binding := network.Linksetchannelbinding{}
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		linkset_channel_binding.Id = data.Id.ValueString()
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		linkset_channel_binding.Ifnum = data.Ifnum.ValueString()
	}

	return linkset_channel_binding
}

func linkset_channel_bindingSetAttrFromGet(ctx context.Context, data *LinksetChannelBindingResourceModel, getResponseData map[string]interface{}) *LinksetChannelBindingResourceModel {
	tflog.Debug(ctx, "In linkset_channel_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
	} else {
		data.Id = types.StringNull()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else {
		data.Ifnum = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Id.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
