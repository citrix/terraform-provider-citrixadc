package aaassoprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaassoprofileResourceModel describes the resource data model.
type AaassoprofileResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func (r *AaassoprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaassoprofile resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SSO Profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a SSO Profile is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Description: "Password with which the user logs on. Required for Single sign on to  external server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my group\" or 'my group').",
			},
		},
	}
}

func aaassoprofileGetThePayloadFromtheConfig(ctx context.Context, data *AaassoprofileResourceModel) aaa.Aaassoprofile {
	tflog.Debug(ctx, "In aaassoprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaassoprofile := aaa.Aaassoprofile{}
	if !data.Name.IsNull() {
		aaassoprofile.Name = data.Name.ValueString()
	}
	if !data.Password.IsNull() {
		aaassoprofile.Password = data.Password.ValueString()
	}
	if !data.Username.IsNull() {
		aaassoprofile.Username = data.Username.ValueString()
	}

	return aaassoprofile
}

func aaassoprofileSetAttrFromGet(ctx context.Context, data *AaassoprofileResourceModel, getResponseData map[string]interface{}) *AaassoprofileResourceModel {
	tflog.Debug(ctx, "In aaassoprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
