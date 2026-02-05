package nsacls

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsaclsResourceModel describes the resource data model.
type NsaclsResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (r *NsaclsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsacls resource.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("CLASSIC"),
				Description: "Type of the acl ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
		},
	}
}

func nsaclsGetThePayloadFromtheConfig(ctx context.Context, data *NsaclsResourceModel) ns.Nsacls {
	tflog.Debug(ctx, "In nsaclsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsacls := ns.Nsacls{}
	if !data.Type.IsNull() {
		nsacls.Type = data.Type.ValueString()
	}

	return nsacls
}

func nsaclsSetAttrFromGet(ctx context.Context, data *NsaclsResourceModel, getResponseData map[string]interface{}) *NsaclsResourceModel {
	tflog.Debug(ctx, "In nsaclsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsacls-config")

	return data
}
