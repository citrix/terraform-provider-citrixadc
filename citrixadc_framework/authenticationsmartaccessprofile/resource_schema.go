package authenticationsmartaccessprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationsmartaccessprofileResourceModel describes the resource data model.
type AuthenticationsmartaccessprofileResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Tags    types.String `tfsdk:"tags"`
}

func (r *AuthenticationsmartaccessprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationsmartaccessprofile resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Optional comment for the profile.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Smartaccess profile",
			},
			"tags": schema.StringAttribute{
				Required:    true,
				Description: "The tag that is associated with Smartaccess profile.",
			},
		},
	}
}

func authenticationsmartaccessprofileGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationsmartaccessprofileResourceModel) authentication.Authenticationsmartaccessprofile {
	tflog.Debug(ctx, "In authenticationsmartaccessprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationsmartaccessprofile := authentication.Authenticationsmartaccessprofile{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		authenticationsmartaccessprofile.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationsmartaccessprofile.Name = data.Name.ValueString()
	}
	if !data.Tags.IsNull() && !data.Tags.IsUnknown() {
		authenticationsmartaccessprofile.Tags = data.Tags.ValueString()
	}

	return authenticationsmartaccessprofile
}

func authenticationsmartaccessprofileSetAttrFromGet(ctx context.Context, data *AuthenticationsmartaccessprofileResourceModel, getResponseData map[string]interface{}) *AuthenticationsmartaccessprofileResourceModel {
	tflog.Debug(ctx, "In authenticationsmartaccessprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["tags"]; ok && val != nil {
		data.Tags = types.StringValue(val.(string))
	} else {
		data.Tags = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
