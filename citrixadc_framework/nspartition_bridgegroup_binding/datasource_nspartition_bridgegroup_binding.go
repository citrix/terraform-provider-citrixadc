package nspartition_bridgegroup_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NspartitionBridgegroupBindingDataSource)(nil)

func NSpartitionBridgegroupBindingDataSource() datasource.DataSource {
	return &NspartitionBridgegroupBindingDataSource{}
}

type NspartitionBridgegroupBindingDataSource struct {
	client *service.NitroClient
}

func (d *NspartitionBridgegroupBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspartition_bridgegroup_binding"
}

func (d *NspartitionBridgegroupBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NspartitionBridgegroupBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NspartitionBridgegroupBindingDataSourceSchema()
}

func (d *NspartitionBridgegroupBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NspartitionBridgegroupBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	partitionname_Name := data.Partitionname.ValueString()
	bridgegroup_Name := data.Bridgegroup

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Nspartition_bridgegroup_binding.Type(),
		ResourceName:             partitionname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nspartition_bridgegroup_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nspartition_bridgegroup_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bridgegroup
		if val, ok := v["bridgegroup"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if bridgegroup_Name.IsNull() || val != bridgegroup_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !bridgegroup_Name.IsNull() {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nspartition_bridgegroup_binding with bridgegroup %s not found", bridgegroup_Name))
		return
	}

	nspartition_bridgegroup_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
