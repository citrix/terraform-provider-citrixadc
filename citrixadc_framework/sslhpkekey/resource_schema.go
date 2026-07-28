package sslhpkekey

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslhpkekeyResourceModel describes the resource data model.
type SslhpkekeyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Dhkem       types.String `tfsdk:"dhkem"`
	File        types.String `tfsdk:"file"`
	Hpkekeyname types.String `tfsdk:"hpkekeyname"`
}

func (r *SslhpkekeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslhpkekey resource.",
			},
			"dhkem": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of curve used for HPKE",
			},
			"file": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the HPKE key file",
			},
			"hpkekeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the hpke key configured that is used to decrypt ECH",
			},
		},
	}
}

func sslhpkekeyGetThePayloadFromthePlan(ctx context.Context, data *SslhpkekeyResourceModel) ssl.Sslhpkekey {
	tflog.Debug(ctx, "In sslhpkekeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslhpkekey := ssl.Sslhpkekey{}
	if !data.Dhkem.IsNull() && !data.Dhkem.IsUnknown() {
		sslhpkekey.Dhkem = data.Dhkem.ValueString()
	}
	if !data.File.IsNull() && !data.File.IsUnknown() {
		sslhpkekey.File = data.File.ValueString()
	}
	if !data.Hpkekeyname.IsNull() && !data.Hpkekeyname.IsUnknown() {
		sslhpkekey.Hpkekeyname = data.Hpkekeyname.ValueString()
	}

	return sslhpkekey
}

func sslhpkekeySetAttrFromGet(ctx context.Context, data *SslhpkekeyResourceModel, getResponseData map[string]interface{}) *SslhpkekeyResourceModel {
	tflog.Debug(ctx, "In sslhpkekeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dhkem"]; ok && val != nil {
		data.Dhkem = types.StringValue(val.(string))
	} else {
		data.Dhkem = types.StringNull()
	}
	if val, ok := getResponseData["file"]; ok && val != nil {
		data.File = types.StringValue(val.(string))
	} else {
		data.File = types.StringNull()
	}
	if val, ok := getResponseData["hpkekeyname"]; ok && val != nil {
		data.Hpkekeyname = types.StringValue(val.(string))
	} else {
		data.Hpkekeyname = types.StringNull()
	}

	// ID is set once in Create (resource) and in the datasource Read; do not
	// recompute it here so a sparse GET response cannot wipe it (Pattern 6).

	return data
}
