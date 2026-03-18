package nsservicepath_nsservicefunction_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsservicepathNsservicefunctionBindingDataSource)(nil)

func NSservicepathNsservicefunctionBindingDataSource() datasource.DataSource {
	return &NsservicepathNsservicefunctionBindingDataSource{}
}

type NsservicepathNsservicefunctionBindingDataSource struct {
	client *service.NitroClient
}

func (d *NsservicepathNsservicefunctionBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsservicepath_nsservicefunction_binding"
}

func (d *NsservicepathNsservicefunctionBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsservicepathNsservicefunctionBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsservicepathNsservicefunctionBindingDataSourceSchema()
}

func (d *NsservicepathNsservicefunctionBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicepathname_Name := data.Servicepathname.ValueString()
	servicefunction_Name := data.Servicefunction

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Nsservicepath_nsservicefunction_binding.Type(),
		ResourceName:             servicepathname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsservicepath_nsservicefunction_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nsservicepath_nsservicefunction_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check servicefunction
		if val, ok := v["servicefunction"].(string); ok {
			if servicefunction_Name.IsNull() || val != servicefunction_Name.ValueString() {
				match = false
				continue
			}
		} else if !servicefunction_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nsservicepath_nsservicefunction_binding with servicefunction %s not found", servicefunction_Name))
		return
	}

	nsservicepath_nsservicefunction_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
