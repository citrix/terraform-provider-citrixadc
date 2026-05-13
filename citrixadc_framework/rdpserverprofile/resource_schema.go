package rdpserverprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/rdp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RdpserverprofileResourceModel describes the resource data model.
type RdpserverprofileResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Psk            types.String `tfsdk:"psk"`
	PskWo          types.String `tfsdk:"psk_wo"`
	PskWoVersion   types.Int64  `tfsdk:"psk_wo_version"`
	Rdpip          types.String `tfsdk:"rdpip"`
	Rdpport        types.Int64  `tfsdk:"rdpport"`
	Rdpredirection types.String `tfsdk:"rdpredirection"`
}

func (r *RdpserverprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rdpserverprofile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the rdp server profile",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Pre shared key value",
			},
			"psk_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Pre shared key value",
			},
			"psk_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a psk_wo update.",
			},
			"rdpip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.",
			},
			"rdpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP port on which the RDP connection is established.",
			},
			"rdpredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable RDP redirection support. This needs to be enabled in presence of connection broker or session directory with IP cookie(msts cookie) based redirection support",
			},
		},
	}
}

func rdpserverprofileGetThePayloadFromthePlan(ctx context.Context, data *RdpserverprofileResourceModel) rdp.Rdpserverprofile {
	tflog.Debug(ctx, "In rdpserverprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rdpserverprofile := rdp.Rdpserverprofile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rdpserverprofile.Name = data.Name.ValueString()
	}
	if !data.Psk.IsNull() && !data.Psk.IsUnknown() {
		rdpserverprofile.Psk = data.Psk.ValueString()
	}
	// Skip write-only attribute: psk_wo
	// Skip version tracker attribute: psk_wo_version
	if !data.Rdpip.IsNull() && !data.Rdpip.IsUnknown() {
		rdpserverprofile.Rdpip = data.Rdpip.ValueString()
	}
	if !data.Rdpport.IsNull() && !data.Rdpport.IsUnknown() {
		rdpserverprofile.Rdpport = utils.IntPtr(int(data.Rdpport.ValueInt64()))
	}
	if !data.Rdpredirection.IsNull() && !data.Rdpredirection.IsUnknown() {
		rdpserverprofile.Rdpredirection = data.Rdpredirection.ValueString()
	}

	return rdpserverprofile
}

func rdpserverprofileGetThePayloadFromtheConfig(ctx context.Context, data *RdpserverprofileResourceModel, payload *rdp.Rdpserverprofile) {
	tflog.Debug(ctx, "In rdpserverprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: psk_wo -> psk
	if !data.PskWo.IsNull() {
		pskWo := data.PskWo.ValueString()
		if pskWo != "" {
			payload.Psk = pskWo
		}
	}
}

func rdpserverprofileSetAttrFromGet(ctx context.Context, data *RdpserverprofileResourceModel, getResponseData map[string]interface{}) *RdpserverprofileResourceModel {
	tflog.Debug(ctx, "In rdpserverprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// psk is not returned by NITRO API (secret/ephemeral) - retain from config
	// psk_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// psk_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["rdpip"]; ok && val != nil {
		data.Rdpip = types.StringValue(val.(string))
	} else {
		data.Rdpip = types.StringNull()
	}
	if val, ok := getResponseData["rdpport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rdpport = types.Int64Value(intVal)
		}
	} else {
		data.Rdpport = types.Int64Null()
	}
	if val, ok := getResponseData["rdpredirection"]; ok && val != nil {
		data.Rdpredirection = types.StringValue(val.(string))
	} else {
		data.Rdpredirection = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
