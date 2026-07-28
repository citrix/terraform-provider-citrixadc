package appfwprofile_grpcvalidation_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileGrpcvalidationBindingDataSource)(nil)

func APpfwprofileGrpcvalidationBindingDataSource() datasource.DataSource {
	return &AppfwprofileGrpcvalidationBindingDataSource{}
}

type AppfwprofileGrpcvalidationBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileGrpcvalidationBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_grpcvalidation_binding"
}

func (d *AppfwprofileGrpcvalidationBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileGrpcvalidationBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileGrpcvalidationBindingDataSourceSchema()
}

func (d *AppfwprofileGrpcvalidationBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileGrpcvalidationBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	grpcrelaxvalidationaction_Name := data.GrpcRelaxValidationAction
	grpcvalidation_Name := data.Grpcvalidation

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_grpcvalidation_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_grpcvalidation_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_grpcvalidation_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check grpc_relax_validation_action
		if val, ok := v["grpc_relax_validation_action"].(string); ok {
			if grpcrelaxvalidationaction_Name.IsNull() || val != grpcrelaxvalidationaction_Name.ValueString() {
				match = false
				continue
			}
		} else if !grpcrelaxvalidationaction_Name.IsNull() {
			match = false
			continue
		}

		// Check grpcvalidation
		if val, ok := v["grpcvalidation"].(string); ok {
			if grpcvalidation_Name.IsNull() || val != grpcvalidation_Name.ValueString() {
				match = false
				continue
			}
		} else if !grpcvalidation_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_grpcvalidation_binding with grpc_relax_validation_action %s not found", grpcrelaxvalidationaction_Name))
		return
	}

	appfwprofile_grpcvalidation_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
