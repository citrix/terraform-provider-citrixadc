package ipsecalgsession

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ipsecalg"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IpsecalgsessionResourceModel describes the resource data model.
type IpsecalgsessionResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Destip      types.String `tfsdk:"destip"`
	DestipAlg   types.String `tfsdk:"destip_alg"`
	Natip       types.String `tfsdk:"natip"`
	NatipAlg    types.String `tfsdk:"natip_alg"`
	Sourceip    types.String `tfsdk:"sourceip"`
	SourceipAlg types.String `tfsdk:"sourceip_alg"`
}

func (r *IpsecalgsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipsecalgsession resource.",
			},
			"destip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination IP address (flush scope).",
			},
			"destip_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination IP address.",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Natted Source IP address (flush scope).",
			},
			"natip_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Natted Source IP address.",
			},
			"sourceip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Original Source IP address (flush scope).",
			},
			"sourceip_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Original Source IP address.",
			},
		},
	}
}

// ipsecalgsessionGetTheFlushPayload builds the ?action=flush payload. Per the
// NITRO doc, flush accepts only the canonical scope fields sourceip/natip/destip
// (a bare flush with no fields flushes all sessions). The *_alg twins are GET
// filter names and are NOT sent, to avoid emitting conflicting duplicate keys.
func ipsecalgsessionGetTheFlushPayload(ctx context.Context, data *IpsecalgsessionResourceModel) ipsecalg.Ipsecalgsession {
	tflog.Debug(ctx, "In ipsecalgsessionGetTheFlushPayload Function")

	ipsecalgsession := ipsecalg.Ipsecalgsession{}
	if !data.Sourceip.IsNull() && !data.Sourceip.IsUnknown() {
		ipsecalgsession.Sourceip = data.Sourceip.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		ipsecalgsession.Natip = data.Natip.ValueString()
	}
	if !data.Destip.IsNull() && !data.Destip.IsUnknown() {
		ipsecalgsession.Destip = data.Destip.ValueString()
	}

	return ipsecalgsession
}

// ipsecalgsessionSetAttrFromGetForDatasource faithfully copies every read-back
// field from the GET (get-all) response into the model. Used only by the
// datasource; the resource is action-only and does not read back.
func ipsecalgsessionSetAttrFromGetForDatasource(ctx context.Context, data *IpsecalgsessionResourceModel, getResponseData map[string]interface{}) *IpsecalgsessionResourceModel {
	tflog.Debug(ctx, "In ipsecalgsessionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["destip"]; ok && val != nil {
		data.Destip = types.StringValue(val.(string))
	} else {
		data.Destip = types.StringNull()
	}
	if val, ok := getResponseData["destip_alg"]; ok && val != nil {
		data.DestipAlg = types.StringValue(val.(string))
	} else {
		data.DestipAlg = types.StringNull()
	}
	if val, ok := getResponseData["natip"]; ok && val != nil {
		data.Natip = types.StringValue(val.(string))
	} else {
		data.Natip = types.StringNull()
	}
	if val, ok := getResponseData["natip_alg"]; ok && val != nil {
		data.NatipAlg = types.StringValue(val.(string))
	} else {
		data.NatipAlg = types.StringNull()
	}
	if val, ok := getResponseData["sourceip"]; ok && val != nil {
		data.Sourceip = types.StringValue(val.(string))
	} else {
		data.Sourceip = types.StringNull()
	}
	if val, ok := getResponseData["sourceip_alg"]; ok && val != nil {
		data.SourceipAlg = types.StringValue(val.(string))
	} else {
		data.SourceipAlg = types.StringNull()
	}

	return data
}
