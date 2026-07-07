package nssourceroutecachetable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NssourceroutecachetableDataSourceModel describes the datasource data model. It
// exposes the synthetic id and the read-only get(all) attributes.
//
// NITRO keys "_nextgenapiresource", "__count" and "Interface" cannot be used as
// tfsdk tags verbatim (leading underscore / uppercase), so the tags are
// lowercased/stripped while the exact NITRO keys are read in SetAttrFromGet.
type NssourceroutecachetableDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Sourceip           types.String `tfsdk:"sourceip"`
	Sourcemac          types.String `tfsdk:"sourcemac"`
	Vlan               types.Int64  `tfsdk:"vlan"`
	Interface          types.String `tfsdk:"interface"`
	Nextgenapiresource types.String `tfsdk:"nextgenapiresource"`
	Count              types.Int64  `tfsdk:"record_count"`
}

func NssourceroutecachetableDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssourceroutecachetable datasource.",
			},
			"sourceip": schema.StringAttribute{
				Computed:    true,
				Description: "Source ip of the connection.",
			},
			"sourcemac": schema.StringAttribute{
				Computed:    true,
				Description: "Source MAC address of an incoming IPv6 packet.",
			},
			"vlan": schema.Int64Attribute{
				Computed:    true,
				Description: "ID of the VLAN.",
			},
			"interface": schema.StringAttribute{
				Computed:    true,
				Description: "ID of an interface (NITRO key: Interface).",
			},
			"nextgenapiresource": schema.StringAttribute{
				Computed:    true,
				Description: "Read-only attribute (_nextgenapiresource).",
			},
			"record_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Count parameter (NITRO key: __count). Renamed from 'count' which is a reserved Terraform root attribute name.",
			},
		},
	}
}

// nssourceroutecachetableSetAttrFromGet maps the get(all) response (exact NITRO
// keys) into the datasource model.
func nssourceroutecachetableSetAttrFromGet(ctx context.Context, data *NssourceroutecachetableDataSourceModel, getResponseData map[string]interface{}) *NssourceroutecachetableDataSourceModel {
	tflog.Debug(ctx, "In nssourceroutecachetableSetAttrFromGet Function")

	if val, ok := getResponseData["sourceip"]; ok && val != nil {
		data.Sourceip = types.StringValue(val.(string))
	} else {
		data.Sourceip = types.StringNull()
	}
	if val, ok := getResponseData["sourcemac"]; ok && val != nil {
		data.Sourcemac = types.StringValue(val.(string))
	} else {
		data.Sourcemac = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		} else {
			data.Vlan = types.Int64Null()
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["Interface"]; ok && val != nil {
		data.Interface = types.StringValue(val.(string))
	} else {
		data.Interface = types.StringNull()
	}
	if val, ok := getResponseData["_nextgenapiresource"]; ok && val != nil {
		data.Nextgenapiresource = types.StringValue(val.(string))
	} else {
		data.Nextgenapiresource = types.StringNull()
	}
	if val, ok := getResponseData["__count"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Count = types.Int64Value(intVal)
		} else {
			data.Count = types.Int64Null()
		}
	} else {
		data.Count = types.Int64Null()
	}

	data.Id = types.StringValue("nssourceroutecachetable-config")

	return data
}
