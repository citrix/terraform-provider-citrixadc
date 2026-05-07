package autoscaleprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AutoscaleprofileResourceModel describes the resource data model.
type AutoscaleprofileResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Apikey                types.String `tfsdk:"apikey"`
	ApikeyWo              types.String `tfsdk:"apikey_wo"`
	ApikeyWoVersion       types.Int64  `tfsdk:"apikey_wo_version"`
	Name                  types.String `tfsdk:"name"`
	Sharedsecret          types.String `tfsdk:"sharedsecret"`
	SharedsecretWo        types.String `tfsdk:"sharedsecret_wo"`
	SharedsecretWoVersion types.Int64  `tfsdk:"sharedsecret_wo_version"`
	Type                  types.String `tfsdk:"type"`
	Url                   types.String `tfsdk:"url"`
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
				Optional:    true,
				Sensitive:   true,
				Description: "api key for authentication with service",
			},
			"apikey_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "api key for authentication with service",
			},
			"apikey_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a apikey_wo update.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "AutoScale profile name.",
			},
			"sharedsecret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "shared secret for authentication with service",
			},
			"sharedsecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "shared secret for authentication with service",
			},
			"sharedsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a sharedsecret_wo update.",
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

func autoscaleprofileGetThePayloadFromthePlan(ctx context.Context, data *AutoscaleprofileResourceModel) autoscale.Autoscaleprofile {
	tflog.Debug(ctx, "In autoscaleprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	autoscaleprofile := autoscale.Autoscaleprofile{}
	if !data.Apikey.IsNull() && !data.Apikey.IsUnknown() {
		autoscaleprofile.Apikey = data.Apikey.ValueString()
	}
	// Skip write-only attribute: apikey_wo
	// Skip version tracker attribute: apikey_wo_version
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		autoscaleprofile.Name = data.Name.ValueString()
	}
	if !data.Sharedsecret.IsNull() && !data.Sharedsecret.IsUnknown() {
		autoscaleprofile.Sharedsecret = data.Sharedsecret.ValueString()
	}
	// Skip write-only attribute: sharedsecret_wo
	// Skip version tracker attribute: sharedsecret_wo_version
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		autoscaleprofile.Type = data.Type.ValueString()
	}
	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		autoscaleprofile.Url = data.Url.ValueString()
	}

	return autoscaleprofile
}

func autoscaleprofileGetThePayloadFromtheConfig(ctx context.Context, data *AutoscaleprofileResourceModel, payload *autoscale.Autoscaleprofile) {
	tflog.Debug(ctx, "In autoscaleprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: apikey_wo -> apikey
	if !data.ApikeyWo.IsNull() {
		apikeyWo := data.ApikeyWo.ValueString()
		if apikeyWo != "" {
			payload.Apikey = apikeyWo
		}
	}
	// Handle write-only secret attribute: sharedsecret_wo -> sharedsecret
	if !data.SharedsecretWo.IsNull() {
		sharedsecretWo := data.SharedsecretWo.ValueString()
		if sharedsecretWo != "" {
			payload.Sharedsecret = sharedsecretWo
		}
	}
}

func autoscaleprofileSetAttrFromGet(ctx context.Context, data *AutoscaleprofileResourceModel, getResponseData map[string]interface{}) *AutoscaleprofileResourceModel {
	tflog.Debug(ctx, "In autoscaleprofileSetAttrFromGet Function")

	// Convert API response to model
	// apikey is not returned by NITRO API (secret/ephemeral) - retain from config
	// apikey_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// apikey_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// sharedsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	// sharedsecret_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// sharedsecret_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
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
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
