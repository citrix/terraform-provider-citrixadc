package subscribersessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

var _ datasource.DataSource = (*SubscribersessionsDataSource)(nil)

func SUbscribersessionsDataSource() datasource.DataSource {
	return &SubscribersessionsDataSource{}
}

type SubscribersessionsDataSource struct {
	client *service.NitroClient
}

func (d *SubscribersessionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscribersessions"
}

func (d *SubscribersessionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SubscribersessionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SubscribersessionsDataSourceSchema()
}

// Read backs the datasource with get(all) and filters in memory.
// subscribersessions is a transient diagnostics table: the session list may
// legitimately be empty (Subscriber/Gx/PCRF Telco feature not licensed/enabled,
// or no active sessions), so an empty result MUST NOT be treated as an error.
// ip/vlan/nodeid are optional filters mirroring the get(all) args.
func (d *SubscribersessionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SubscribersessionsDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading subscribersessions datasource via get(all)")

	all, err := d.client.FindAllResources(service.Subscribersessions.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read subscribersessions, got error: %s", err))
		return
	}

	// Optional in-memory filtering by ip/vlan/nodeid. Empty result is tolerated.
	var match map[string]interface{}
	for _, rec := range all {
		if !data.Ip.IsNull() && !data.Ip.IsUnknown() && data.Ip.ValueString() != "" {
			if v, ok := rec["ip"].(string); !ok || v != data.Ip.ValueString() {
				continue
			}
		}
		if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
			if v, ok := rec["vlan"]; ok {
				if iv, cerr := utils.ConvertToInt64(v); cerr != nil || iv != data.Vlan.ValueInt64() {
					continue
				}
			} else {
				continue
			}
		}
		if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
			if v, ok := rec["nodeid"]; ok {
				if iv, cerr := utils.ConvertToInt64(v); cerr != nil || iv != data.Nodeid.ValueInt64() {
					continue
				}
			} else {
				continue
			}
		}
		match = rec
		break
	}

	if match != nil {
		subscribersessionsSetAttrFromGetForDatasourceModel(ctx, &data, match)
	} else {
		tflog.Debug(ctx, "No subscribersessions records found; returning empty result")
	}

	// Datasource sets its own synthetic ID (there is no Create).
	data.Id = types.StringValue("subscribersessions")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// subscribersessionsSetAttrFromGetForDatasourceModel faithfully copies the
// get(all) response record into the datasource model. flags/ttl/idlettl (Double)
// arrive as float64 from the JSON decoder, so they go through
// utils.ConvertToInt64. String telemetry maps directly. Missing fields are left
// as their zero-value (null) so callers can distinguish absent telemetry.
func subscribersessionsSetAttrFromGetForDatasourceModel(ctx context.Context, data *SubscribersessionsDataSourceModel, getResponseData map[string]interface{}) *SubscribersessionsDataSourceModel {
	tflog.Debug(ctx, "In subscribersessionsSetAttrFromGetForDatasourceModel Function")

	setString := func(key string, dst *types.String) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if s, ok := val.(string); ok {
				*dst = types.StringValue(s)
			}
		}
	}
	setInt := func(key string, dst *types.Int64) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				*dst = types.Int64Value(intVal)
			}
		}
	}

	// Echo back the filter fields that appear in the record.
	setString("ip", &data.Ip)
	setInt("vlan", &data.Vlan)
	setInt("nodeid", &data.Nodeid)

	// Read-only telemetry.
	setString("subscriptionidtype", &data.Subscriptionidtype)
	setString("subscriptionidvalue", &data.Subscriptionidvalue)
	setString("subscriberrules", &data.Subscriberrules)
	setInt("flags", &data.Flags)
	setInt("ttl", &data.Ttl)
	setInt("idlettl", &data.Idlettl)
	setString("avpdisplaybuffer", &data.Avpdisplaybuffer)
	setString("servicepath", &data.Servicepath)

	return data
}
