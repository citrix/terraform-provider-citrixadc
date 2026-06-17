package nslimitsessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*NslimitsessionsDataSource)(nil)

func NSlimitsessionsDataSource() datasource.DataSource {
	return &NslimitsessionsDataSource{}
}

type NslimitsessionsDataSource struct {
	client *service.NitroClient
}

func (d *NslimitsessionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitsessions"
}

func (d *NslimitsessionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NslimitsessionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NslimitsessionsDataSourceSchema()
}

func (d *NslimitsessionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NslimitsessionsDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Lookup keyed on limitidentifier; detail is an optional GET filter.
	limitidentifier_Name := data.Limitidentifier
	detail_Name := data.Detail

	var dataArr []map[string]interface{}
	var err error

	// NITRO requires limitidentifier as a mandatory GET arg for nslimitsessions
	// (a GET without it returns errorcode 1095 "Required argument missing").
	// detail is an optional GET arg. These are attached as
	// args=limitidentifier:<value>[,detail:<value>] via ArgsMap; the nitro-go
	// client URL-encodes the values.
	var argsMap map[string]string = make(map[string]string)
	if !limitidentifier_Name.IsNull() && limitidentifier_Name.ValueString() != "" {
		argsMap["limitidentifier"] = limitidentifier_Name.ValueString()
	}
	if !detail_Name.IsNull() {
		if detail_Name.ValueBool() {
			argsMap["detail"] = "true"
		} else {
			argsMap["detail"] = "false"
		}
	}

	findParams := service.FindParams{
		ResourceType:             service.Nslimitsessions.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nslimitsessions, got error: %s", err))
		return
	}

	// A freshly-created rate limit identifier has no active sessions, so NITRO
	// returns errorcode 0 with no nslimitsessions key (empty array). The
	// identifier still exists with zero sessions, so resolve the datasource
	// from the lookup key rather than failing "not found".
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check limitidentifier (required lookup key)
		if val, ok := v["limitidentifier"].(string); ok {
			if limitidentifier_Name.IsNull() || val != limitidentifier_Name.ValueString() {
				match = false
				continue
			}
		} else if !limitidentifier_Name.IsNull() {
			match = false
			continue
		}

		// Check detail (optional filter)
		if !detail_Name.IsNull() {
			if val, ok := v["detail"].(bool); ok {
				if val != detail_Name.ValueBool() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		// A matching session record was returned; copy its attributes.
		nslimitsessionsSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	} else {
		// No sessions (empty/OK response) or no record matched the lookup key.
		// The identifier exists with zero active sessions; resolve the
		// datasource from the lookup key and set its ID.
		data.Limitidentifier = limitidentifier_Name
		data.Id = types.StringValue(limitidentifier_Name.ValueString())
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
