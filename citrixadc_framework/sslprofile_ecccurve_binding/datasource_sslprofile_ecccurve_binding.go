package sslprofile_ecccurve_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

var _ datasource.DataSource = (*SslprofileEcccurveBindingDataSource)(nil)

func SSlprofileEcccurveBindingDataSource() datasource.DataSource {
	return &SslprofileEcccurveBindingDataSource{}
}

type SslprofileEcccurveBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslprofileEcccurveBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_ecccurve_binding"
}

func (d *SslprofileEcccurveBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslprofileEcccurveBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslprofileEcccurveBindingDataSourceSchema()
}

func (d *SslprofileEcccurveBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslprofileEcccurveBindingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name_Name := data.Name.ValueString()
	ecccurvename_Name := data.Ecccurvename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslprofile_ecccurve_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_ecccurve_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslprofile_ecccurve_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ecccurvename
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ecccurvename"].(string); ok {
			if !ecccurvename_Name.IsNull() && val == ecccurvename_Name.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslprofile_ecccurve_binding with ecccurvename %s not found", ecccurvename_Name.ValueString()))
		return
	}

	sslprofile_ecccurve_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// sslprofile_ecccurve_bindingSetAttrFromGetForDatasource faithfully copies the GET
// response into the datasource model and computes the composite ID.
func sslprofile_ecccurve_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslprofileEcccurveBindingDataSourceModel, getResponseData map[string]interface{}) {
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["ecccurvename"]; ok && val != nil {
		data.Ecccurvename = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// Composite ID: name:UrlEncode(name),ecccurvename:UrlEncode(ecccurvename)
	idParts := []string{
		fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())),
		fmt.Sprintf("ecccurvename:%s", utils.UrlEncode(data.Ecccurvename.ValueString())),
	}
	data.Id = types.StringValue(idParts[0] + "," + idParts[1])
}
