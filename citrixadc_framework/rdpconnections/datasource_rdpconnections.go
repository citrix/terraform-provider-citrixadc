package rdpconnections

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

var _ datasource.DataSource = (*RdpconnectionsDataSource)(nil)

func RDpconnectionsDataSource() datasource.DataSource {
	return &RdpconnectionsDataSource{}
}

type RdpconnectionsDataSource struct {
	client *service.NitroClient
}

func (d *RdpconnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rdpconnections"
}

func (d *RdpconnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *RdpconnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RdpconnectionsDataSourceSchema()
}

// Read backs the datasource with get(all) and filters in memory. rdpconnections
// is a transient diagnostics table: the connection list may legitimately be
// empty (no RDP proxy traffic), so an empty result MUST NOT be treated as an
// error. username is an optional filter mirroring "show rdp connections
// [-userName]".
func (d *RdpconnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RdpconnectionsDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rdpconnections datasource via get(all)")

	all, err := d.client.FindAllResources(service.Rdpconnections.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read rdpconnections, got error: %s", err))
		return
	}

	// Optional in-memory filtering by username. Empty result is tolerated.
	usernameFilter := data.Username.ValueString()
	var match map[string]interface{}
	for _, rec := range all {
		if !data.Username.IsNull() && !data.Username.IsUnknown() && usernameFilter != "" {
			if v, ok := rec["username"].(string); ok && v == usernameFilter {
				match = rec
				break
			}
			continue
		}
		// No username filter supplied: take the first record if present.
		match = rec
		break
	}

	if match != nil {
		rdpconnectionsSetAttrFromGetForDatasourceModel(ctx, &data, match)
	} else {
		tflog.Debug(ctx, "No rdpconnections records found; returning empty result")
	}

	// Datasource sets its own synthetic ID (there is no Create).
	data.Id = types.StringValue("rdpconnections")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// rdpconnectionsSetAttrFromGetForDatasourceModel faithfully copies the get(all)
// response record into the datasource model. endpointport/targetport (Integer)
// and peid (Double) arrive as float64 from the JSON decoder, so they go through
// utils.ConvertToInt64. String telemetry maps directly. Missing fields are left
// as their zero-value (null) so callers can distinguish absent telemetry.
func rdpconnectionsSetAttrFromGetForDatasourceModel(ctx context.Context, data *RdpconnectionsDataSourceModel, getResponseData map[string]interface{}) *RdpconnectionsDataSourceModel {
	tflog.Debug(ctx, "In rdpconnectionsSetAttrFromGetForDatasourceModel Function")

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

	setString("username", &data.Username)
	setString("endpointip", &data.Endpointip)
	setInt("endpointport", &data.Endpointport)
	setString("targetip", &data.Targetip)
	setInt("targetport", &data.Targetport)
	setInt("peid", &data.Peid)

	return data
}
