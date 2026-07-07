package appfwlearningdata

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwlearningdataDataSourceModel describes the datasource data model.
//
// appfwlearningdata get(all) returns a list of learned-data entries. This
// BEST-EFFORT datasource exposes profilename/securitycheck as optional lookup
// filters and surfaces the readonly (learned) fields of the FIRST matching entry.
// (NITRO can return many rows; this datasource reports a single representative
// row — verify against the live table if you need the full set.)
type AppfwlearningdataDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Profilename            types.String `tfsdk:"profilename"`
	Securitycheck          types.String `tfsdk:"securitycheck"`
	Url                    types.String `tfsdk:"url"`
	Name                   types.String `tfsdk:"name"`
	Fieldtype              types.String `tfsdk:"fieldtype"`
	Fieldformatminlength   types.Int64  `tfsdk:"fieldformatminlength"`
	Fieldformatmaxlength   types.Int64  `tfsdk:"fieldformatmaxlength"`
	Fieldformatcharmappcre types.String `tfsdk:"fieldformatcharmappcre"`
	ValueType              types.String `tfsdk:"value_type"`
	Value                  types.String `tfsdk:"value"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Data                   types.String `tfsdk:"data"`
}

func AppfwlearningdataDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the profile to look up learned data for.",
			},
			"securitycheck": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the security check to look up learned data for.",
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: "Learnt url.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Learnt field name.",
			},
			"fieldtype": schema.StringAttribute{
				Computed:    true,
				Description: "Learnt field type.",
			},
			"fieldformatminlength": schema.Int64Attribute{
				Computed:    true,
				Description: "The minimum allowed length for data in this form field.",
			},
			"fieldformatmaxlength": schema.Int64Attribute{
				Computed:    true,
				Description: "The maximum allowed length for data in this form field.",
			},
			"fieldformatcharmappcre": schema.StringAttribute{
				Computed:    true,
				Description: "Form field value allowed character map.",
			},
			"value_type": schema.StringAttribute{
				Computed:    true,
				Description: "Learnt field value type.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "Learnt field value.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Learnt entity hit count.",
			},
			"data": schema.StringAttribute{
				Computed:    true,
				Description: "Learned data.",
			},
		},
	}
}

// appfwlearningdataSetAttrFromGet populates the datasource model from a single
// learned-data entry returned by NITRO get(all).
func appfwlearningdataSetAttrFromGet(ctx context.Context, data *AppfwlearningdataDataSourceModel, getResponseData map[string]interface{}) *AppfwlearningdataDataSourceModel {
	tflog.Debug(ctx, "In appfwlearningdataSetAttrFromGet Function")

	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["fieldtype"]; ok && val != nil {
		data.Fieldtype = types.StringValue(val.(string))
	} else {
		data.Fieldtype = types.StringNull()
	}
	if val, ok := getResponseData["fieldformatminlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatminlength = types.Int64Value(intVal)
		} else {
			data.Fieldformatminlength = types.Int64Null()
		}
	} else {
		data.Fieldformatminlength = types.Int64Null()
	}
	if val, ok := getResponseData["fieldformatmaxlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatmaxlength = types.Int64Value(intVal)
		} else {
			data.Fieldformatmaxlength = types.Int64Null()
		}
	} else {
		data.Fieldformatmaxlength = types.Int64Null()
	}
	if val, ok := getResponseData["fieldformatcharmappcre"]; ok && val != nil {
		data.Fieldformatcharmappcre = types.StringValue(val.(string))
	} else {
		data.Fieldformatcharmappcre = types.StringNull()
	}
	if val, ok := getResponseData["value_type"]; ok && val != nil {
		data.ValueType = types.StringValue(val.(string))
	} else {
		data.ValueType = types.StringNull()
	}
	if val, ok := getResponseData["value"]; ok && val != nil {
		data.Value = types.StringValue(val.(string))
	} else {
		data.Value = types.StringNull()
	}
	if val, ok := getResponseData["hits"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hits = types.Int64Value(intVal)
		} else {
			data.Hits = types.Int64Null()
		}
	} else {
		data.Hits = types.Int64Null()
	}
	if val, ok := getResponseData["data"]; ok && val != nil {
		data.Data = types.StringValue(val.(string))
	} else {
		data.Data = types.StringNull()
	}

	return data
}
