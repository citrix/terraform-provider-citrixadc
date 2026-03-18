package autoscaleprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AutoscaleprofileResourceModel describes the resource data model.
type AutoscaleprofileResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Apikey       types.String `tfsdk:"apikey"`
	Name         types.String `tfsdk:"name"`
	Sharedsecret types.String `tfsdk:"sharedsecret"`
	Type         types.String `tfsdk:"type"`
	Url          types.String `tfsdk:"url"`
}

func (r *AutoscaleprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the autoscaleprofile resource.",
			},
			"apikey": schema.StringAttribute{
				Required:    true,
				Description: "api key for authentication with service",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "AutoScale profile name.",
			},
			"sharedsecret": schema.StringAttribute{
				Required:    true,
				Description: "shared secret for authentication with service",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of profile.",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "URL providing the service",
			},
		},
	}
}

func autoscaleprofileGetThePayloadFromtheConfig(ctx context.Context, data *AutoscaleprofileResourceModel) autoscale.Autoscaleprofile {
	tflog.Debug(ctx, "In autoscaleprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	autoscaleprofile := autoscale.Autoscaleprofile{}
	if !data.Apikey.IsNull() {
		autoscaleprofile.Apikey = data.Apikey.ValueString()
	}
	if !data.Name.IsNull() {
		autoscaleprofile.Name = data.Name.ValueString()
	}
	if !data.Sharedsecret.IsNull() {
		autoscaleprofile.Sharedsecret = data.Sharedsecret.ValueString()
	}
	if !data.Type.IsNull() {
		autoscaleprofile.Type = data.Type.ValueString()
	}
	if !data.Url.IsNull() {
		autoscaleprofile.Url = data.Url.ValueString()
	}

	return autoscaleprofile
}

func autoscaleprofileSetAttrFromGet(ctx context.Context, data *AutoscaleprofileResourceModel, getResponseData map[string]interface{}) *AutoscaleprofileResourceModel {
	tflog.Debug(ctx, "In autoscaleprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["apikey"]; ok && val != nil {
		data.Apikey = types.StringValue(val.(string))
	} else {
		data.Apikey = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["sharedsecret"]; ok && val != nil {
		data.Sharedsecret = types.StringValue(val.(string))
	} else {
		data.Sharedsecret = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
