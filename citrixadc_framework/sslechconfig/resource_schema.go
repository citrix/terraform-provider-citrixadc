package sslechconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslechconfigResourceModel describes the resource data model.
type SslechconfigResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Echcipher     types.String `tfsdk:"echcipher"`
	Echconfigid   types.Int64  `tfsdk:"echconfigid"`
	Echconfigname types.String `tfsdk:"echconfigname"`
	Echpublicname types.String `tfsdk:"echpublicname"`
	Hpkekeyname   types.String `tfsdk:"hpkekeyname"`
	Version       types.Int64  `tfsdk:"version"`
}

func (r *SslechconfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslechconfig resource.",
			},
			"echcipher": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The supported cipher suite that encrypts the client Hello Message.",
			},
			"echconfigid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The config id of the ech config.",
			},
			"echconfigname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The ECH config name configured.",
			},
			"echpublicname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The public name of ech config means FQDN or any string",
			},
			"hpkekeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the configured HPKE key",
			},
			"version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(65037),
				Description: "The version of ECH for which this configuration is used.",
			},
		},
	}
}

func sslechconfigGetThePayloadFromthePlan(ctx context.Context, data *SslechconfigResourceModel) ssl.Sslechconfig {
	tflog.Debug(ctx, "In sslechconfigGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslechconfig := ssl.Sslechconfig{}
	if !data.Echcipher.IsNull() && !data.Echcipher.IsUnknown() {
		sslechconfig.Echcipher = data.Echcipher.ValueString()
	}
	if !data.Echconfigid.IsNull() && !data.Echconfigid.IsUnknown() {
		sslechconfig.Echconfigid = utils.IntPtr(int(data.Echconfigid.ValueInt64()))
	}
	if !data.Echconfigname.IsNull() && !data.Echconfigname.IsUnknown() {
		sslechconfig.Echconfigname = data.Echconfigname.ValueString()
	}
	if !data.Echpublicname.IsNull() && !data.Echpublicname.IsUnknown() {
		sslechconfig.Echpublicname = data.Echpublicname.ValueString()
	}
	if !data.Hpkekeyname.IsNull() && !data.Hpkekeyname.IsUnknown() {
		sslechconfig.Hpkekeyname = data.Hpkekeyname.ValueString()
	}
	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		sslechconfig.Version = utils.IntPtr(int(data.Version.ValueInt64()))
	}

	return sslechconfig
}

func sslechconfigSetAttrFromGet(ctx context.Context, data *SslechconfigResourceModel, getResponseData map[string]interface{}) *SslechconfigResourceModel {
	tflog.Debug(ctx, "In sslechconfigSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["echcipher"]; ok && val != nil {
		data.Echcipher = types.StringValue(val.(string))
	} else {
		data.Echcipher = types.StringNull()
	}
	if val, ok := getResponseData["echconfigid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Echconfigid = types.Int64Value(intVal)
		}
	} else {
		data.Echconfigid = types.Int64Null()
	}
	if val, ok := getResponseData["echconfigname"]; ok && val != nil {
		data.Echconfigname = types.StringValue(val.(string))
	} else {
		data.Echconfigname = types.StringNull()
	}
	if val, ok := getResponseData["echpublicname"]; ok && val != nil {
		data.Echpublicname = types.StringValue(val.(string))
	} else {
		data.Echpublicname = types.StringNull()
	}
	if val, ok := getResponseData["hpkekeyname"]; ok && val != nil {
		data.Hpkekeyname = types.StringValue(val.(string))
	} else {
		data.Hpkekeyname = types.StringNull()
	}
	if val, ok := getResponseData["version"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Version = types.Int64Value(intVal)
		}
	} else {
		data.Version = types.Int64Null()
	}

	// ID is set once in Create and preserved across Read; do not recompute here.

	return data
}

func sslechconfigSetAttrFromGetForDatasource(ctx context.Context, data *SslechconfigResourceModel, getResponseData map[string]interface{}) *SslechconfigResourceModel {
	tflog.Debug(ctx, "In sslechconfigSetAttrFromGetForDatasource Function")

	sslechconfigSetAttrFromGet(ctx, data, getResponseData)

	// Datasource has no Create; set the ID from the key (plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Echconfigname.ValueString()))

	return data
}
