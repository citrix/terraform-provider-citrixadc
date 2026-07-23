package apiprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/api"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ApiprofileResourceModel describes the resource data model.
type ApiprofileResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Apivisibility types.String `tfsdk:"apivisibility"`
	Name          types.String `tfsdk:"name"`
}

func (r *ApiprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the apiprofile resource.",
			},
			"apivisibility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable the schema lookup for the requests/apispecs that are bounded to the API profile. The default value of this parameter is DISABLED.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the API profile to add",
			},
		},
	}
}

func apiprofileGetThePayloadFromthePlan(ctx context.Context, data *ApiprofileResourceModel) api.Apiprofile {
	tflog.Debug(ctx, "In apiprofileGetThePayloadFromthePlan Function")

	apiprofile := api.Apiprofile{}
	if !data.Apivisibility.IsNull() && !data.Apivisibility.IsUnknown() {
		apiprofile.Apivisibility = data.Apivisibility.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		apiprofile.Name = data.Name.ValueString()
	}

	return apiprofile
}

func apiprofileSetAttrFromGet(ctx context.Context, data *ApiprofileResourceModel, getResponseData map[string]interface{}) *ApiprofileResourceModel {
	tflog.Debug(ctx, "In apiprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["apivisibility"]; ok && val != nil {
		data.Apivisibility = types.StringValue(val.(string))
	} else {
		data.Apivisibility = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	return data
}
