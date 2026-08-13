package nitro_info

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Backward-compatible migration of the legacy SDKv2 `citrixadc_nitro_info` data
// source: a generic NITRO query escape-hatch. Two `workflow.lifecycle` modes are
// supported, exactly as in SDKv2:
//   - "binding_list"   -> populates nitro_list (a list of {object = map})
//   - "object_by_name" -> populates nitro_object (a single map)
// The type name and all attributes (names/types/optionality) are preserved.

var _ datasource.DataSource = (*nitroInfoDataSource)(nil)
var _ datasource.DataSourceWithConfigure = (*nitroInfoDataSource)(nil)

// element type of nitro_list: object{ object = map[string]string }
var nitroListObjectType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"object": types.MapType{ElemType: types.StringType},
	},
}

func NitroInfoDataSource() datasource.DataSource {
	return &nitroInfoDataSource{}
}

type nitroInfoDataSource struct {
	client *service.NitroClient
}

type NitroInfoDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Workflow    types.Map    `tfsdk:"workflow"`
	QueryArgs   types.Map    `tfsdk:"query_args"`
	PrimaryId   types.String `tfsdk:"primary_id"`
	SecondaryId types.String `tfsdk:"secondary_id"`
	NitroList   types.List   `tfsdk:"nitro_list"`
	NitroObject types.Map    `tfsdk:"nitro_object"`
}

func (d *nitroInfoDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nitro_info"
}

func (d *nitroInfoDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *nitroInfoDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Generic NITRO query data source (binding_list / object_by_name workflows).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"workflow": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Workflow definition: lifecycle, endpoint, bound_resource_missing_errorcode.",
			},
			"query_args": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Additional NITRO query arguments.",
			},
			"primary_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Primary resource id/name to query.",
			},
			"secondary_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secondary resource id/name (reserved).",
			},
			"nitro_list": schema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Result list for the binding_list workflow.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object": schema.MapAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"nitro_object": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Result object for the object_by_name workflow.",
			},
		},
	}
}

func (d *nitroInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "In NitroInfoDataSource Read")
	var data NitroInfoDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflowMap := make(map[string]string)
	if !data.Workflow.IsNull() && !data.Workflow.IsUnknown() {
		resp.Diagnostics.Append(data.Workflow.ElementsAs(ctx, &workflowMap, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	argsMap := make(map[string]string)
	if !data.QueryArgs.IsNull() && !data.QueryArgs.IsUnknown() {
		resp.Diagnostics.Append(data.QueryArgs.ElementsAs(ctx, &argsMap, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// All Computed attributes must be known after Read: default both result
	// containers to empty, and echo the config values / defaults for the rest.
	emptyList, dl := types.ListValue(nitroListObjectType, []attr.Value{})
	resp.Diagnostics.Append(dl...)
	emptyObj, dm := types.MapValue(types.StringType, map[string]attr.Value{})
	resp.Diagnostics.Append(dm...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.NitroList = emptyList
	data.NitroObject = emptyObj
	if data.SecondaryId.IsUnknown() {
		data.SecondaryId = types.StringNull()
	}
	primaryId := data.PrimaryId.ValueString()
	data.Id = types.StringValue("nitro-info-" + workflowMap["endpoint"] + "-" + primaryId)

	lifecycle := workflowMap["lifecycle"]
	switch lifecycle {
	case "binding_list":
		d.readBindingList(ctx, &data, workflowMap, argsMap, primaryId, resp)
	case "object_by_name":
		d.readObjectByName(ctx, &data, workflowMap, argsMap, primaryId, resp)
	default:
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Lifecycle %q is not implemented", lifecycle))
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *nitroInfoDataSource) readBindingList(ctx context.Context, data *NitroInfoDataSourceModel, workflowMap, argsMap map[string]string, primaryId string, resp *datasource.ReadResponse) {
	missingCodeStr := workflowMap["bound_resource_missing_errorcode"]
	missingErrorCode, err := strconv.Atoi(missingCodeStr)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("bound_resource_missing_errorcode must be an integer: %s", err))
		return
	}
	dataArr, err := d.client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             workflowMap["endpoint"],
		ResourceName:             primaryId,
		ResourceMissingErrorCode: missingErrorCode,
		ArgsMap:                  argsMap,
	})
	if err != nil {
		if strings.Contains(err.Error(), missingCodeStr) {
			// bound resource missing -> empty list (already set)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nitro_info binding_list, got error: %s", err))
		return
	}

	elems := make([]attr.Value, 0, len(dataArr))
	for _, item := range dataArr {
		itemMap := make(map[string]string, len(item))
		for k, v := range item {
			itemMap[k] = fmt.Sprintf("%v", v)
		}
		objVal, di := types.MapValueFrom(ctx, types.StringType, itemMap)
		resp.Diagnostics.Append(di...)
		elemObj, do := types.ObjectValue(nitroListObjectType.AttrTypes, map[string]attr.Value{"object": objVal})
		resp.Diagnostics.Append(do...)
		elems = append(elems, elemObj)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	listVal, dlv := types.ListValue(nitroListObjectType, elems)
	resp.Diagnostics.Append(dlv...)
	data.NitroList = listVal
}

func (d *nitroInfoDataSource) readObjectByName(ctx context.Context, data *NitroInfoDataSourceModel, workflowMap, argsMap map[string]string, primaryId string, resp *datasource.ReadResponse) {
	missingCodeStr := workflowMap["bound_resource_missing_errorcode"]
	missingErrorCode, err := strconv.Atoi(missingCodeStr)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("bound_resource_missing_errorcode must be an integer: %s", err))
		return
	}
	dataArr, err := d.client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             workflowMap["endpoint"],
		ResourceName:             primaryId,
		ResourceMissingErrorCode: missingErrorCode,
		ArgsMap:                  argsMap,
	})
	if err != nil {
		if strings.Contains(err.Error(), missingCodeStr) {
			// missing -> empty object (already set)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nitro_info object_by_name, got error: %s", err))
		return
	}
	if len(dataArr) > 1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Too many results %d", len(dataArr)))
		return
	}
	if len(dataArr) == 0 {
		return // empty object already set
	}

	outputMap := make(map[string]string, len(dataArr[0]))
	for k, v := range dataArr[0] {
		outputMap[k] = fmt.Sprintf("%v", v)
	}
	objVal, dm := types.MapValueFrom(ctx, types.StringType, outputMap)
	resp.Diagnostics.Append(dm...)
	data.NitroObject = objVal
}
