package userprotocol

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/user"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UserprotocolResourceModel describes the resource data model.
type UserprotocolResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Extension types.String `tfsdk:"extension"`
	Name      types.String `tfsdk:"name"`
	Transport types.String `tfsdk:"transport"`
}

func (r *UserprotocolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the userprotocol resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the protocol.",
			},
			"extension": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the extension to add parsing and runtime handling of the protocol packets.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name for the user protocol. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters.",
			},
			"transport": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Transport layer's protocol.",
			},
		},
	}
}

func userprotocolGetThePayloadFromtheConfig(ctx context.Context, data *UserprotocolResourceModel) user.Userprotocol {
	tflog.Debug(ctx, "In userprotocolGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	userprotocol := user.Userprotocol{}
	if !data.Comment.IsNull() {
		userprotocol.Comment = data.Comment.ValueString()
	}
	if !data.Extension.IsNull() {
		userprotocol.Extension = data.Extension.ValueString()
	}
	if !data.Name.IsNull() {
		userprotocol.Name = data.Name.ValueString()
	}
	if !data.Transport.IsNull() {
		userprotocol.Transport = data.Transport.ValueString()
	}

	return userprotocol
}

func userprotocolSetAttrFromGet(ctx context.Context, data *UserprotocolResourceModel, getResponseData map[string]interface{}) *UserprotocolResourceModel {
	tflog.Debug(ctx, "In userprotocolSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["extension"]; ok && val != nil {
		data.Extension = types.StringValue(val.(string))
	} else {
		data.Extension = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["transport"]; ok && val != nil {
		data.Transport = types.StringValue(val.(string))
	} else {
		data.Transport = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
