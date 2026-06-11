package policytracing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicytracingResourceModel describes the RESOURCE data model.
//
// policytracing is an action-only resource (clear via ?action=clear). The clear
// action takes no arguments, so the RESOURCE schema exposes only the synthetic
// id. The model therefore declares ONLY id -- the Plugin Framework requires the
// model struct's tfsdk fields to exactly match the schema attributes.
//
// The DATASOURCE keeps its own model (PolicytracingDataSourceModel) with the
// show/get filters (transactionid, detail, nodeid) and read-only output fields
// (capturesslhandshakepolicies, filterexpr, protocoltype).
type PolicytracingResourceModel struct {
	Id types.String `tfsdk:"id"`
}

// PolicytracingDataSourceModel describes the DATASOURCE data model. It exposes
// the show/get filters (transactionid, detail, nodeid) and the read-only output
// fields returned by the get(all) endpoint (capturesslhandshakepolicies,
// filterexpr, protocoltype).
type PolicytracingDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Capturesslhandshakepolicies types.String `tfsdk:"capturesslhandshakepolicies"`
	Detail                      types.String `tfsdk:"detail"`
	Filterexpr                  types.String `tfsdk:"filterexpr"`
	Nodeid                      types.Int64  `tfsdk:"nodeid"`
	Protocoltype                types.String `tfsdk:"protocoltype"`
	Transactionid               types.String `tfsdk:"transactionid"`
}

func (r *PolicytracingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	// The clear action takes no arguments, so the resource schema is minimal:
	// just the synthetic id (Pattern 13: only id is Computed, no other Computed
	// attrs since Read is a no-op). The show-filter fields (transactionid, detail,
	// nodeid) and output-only fields (capturesslhandshakepolicies, filterexpr,
	// protocoltype) are NOT writable inputs to clear and belong on the datasource.
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policytracing resource.",
			},
		},
	}
}

// policytracingGetThePayloadFromthePlan builds the body for the clear action.
// clear takes no arguments (NITRO doc shows an empty body {"policytracing":{}}),
// so the body is always empty.
func policytracingGetThePayloadFromthePlan(ctx context.Context, data *PolicytracingResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In policytracingGetThePayloadFromthePlan Function")

	policytracing := map[string]interface{}{}

	return policytracing
}

// policytracingSetAttrFromGetForDatasource faithfully copies the GET (get all)
// response into the model for the read-only datasource. It exposes the show
// filters (transactionid, detail, nodeid) and the read-only output fields
// modelled in tfdata (capturesslhandshakepolicies, filterexpr, protocoltype).
// The resource itself never calls this (its Read is a no-op).
func policytracingSetAttrFromGetForDatasource(ctx context.Context, data *PolicytracingDataSourceModel, getResponseData map[string]interface{}) *PolicytracingDataSourceModel {
	tflog.Debug(ctx, "In policytracingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["capturesslhandshakepolicies"]; ok && val != nil {
		data.Capturesslhandshakepolicies = types.StringValue(val.(string))
	} else {
		data.Capturesslhandshakepolicies = types.StringNull()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		data.Detail = types.StringValue(val.(string))
	} else {
		data.Detail = types.StringNull()
	}
	if val, ok := getResponseData["filterexpr"]; ok && val != nil {
		data.Filterexpr = types.StringValue(val.(string))
	} else {
		data.Filterexpr = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["protocoltype"]; ok && val != nil {
		data.Protocoltype = types.StringValue(val.(string))
	} else {
		data.Protocoltype = types.StringNull()
	}
	if val, ok := getResponseData["transactionid"]; ok && val != nil {
		data.Transactionid = types.StringValue(val.(string))
	} else {
		data.Transactionid = types.StringNull()
	}

	return data
}
