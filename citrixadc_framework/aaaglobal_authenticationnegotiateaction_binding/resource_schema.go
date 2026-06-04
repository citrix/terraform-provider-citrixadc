package aaaglobal_authenticationnegotiateaction_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

)

// AaaglobalAuthenticationnegotiateactionBindingResourceModel describes the resource data model.
type AaaglobalAuthenticationnegotiateactionBindingResourceModel struct {
	Id types.String `tfsdk:"id"`
	Windowsprofile types.String `tfsdk:"windowsprofile"`
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaglobal_authenticationnegotiateaction_binding resource.",
			},
			"windowsprofile": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the negotiate profile to be bound.",
			},
		},
	}
}

func aaaglobal_authenticationnegotiateaction_bindingGetThePayloadFromthePlan(ctx context.Context, data *AaaglobalAuthenticationnegotiateactionBindingResourceModel) aaa.Aaaglobalauthenticationnegotiateactionbinding {
	tflog.Debug(ctx, "In aaaglobal_authenticationnegotiateaction_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaaglobal_authenticationnegotiateaction_binding := aaa.Aaaglobalauthenticationnegotiateactionbinding{}
	if !data.Windowsprofile.IsNull() && !data.Windowsprofile.IsUnknown() {
		aaaglobal_authenticationnegotiateaction_binding.Windowsprofile = data.Windowsprofile.ValueString()
	}

	return aaaglobal_authenticationnegotiateaction_binding
}

func aaaglobal_authenticationnegotiateaction_bindingSetAttrFromGet(ctx context.Context, data *AaaglobalAuthenticationnegotiateactionBindingResourceModel, getResponseData map[string]interface{}) *AaaglobalAuthenticationnegotiateactionBindingResourceModel {
	tflog.Debug(ctx, "In aaaglobal_authenticationnegotiateaction_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["windowsprofile"]; ok && val != nil {
		data.Windowsprofile = types.StringValue(val.(string))
	} else {
		data.Windowsprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Windowsprofile.ValueString()))

	return data
}