package policypatsetfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*PolicypatsetfileDataSource)(nil)

func POlicypatsetfileDataSource() datasource.DataSource {
	return &PolicypatsetfileDataSource{}
}

type PolicypatsetfileDataSource struct {
	client *service.NitroClient
}

func (d *PolicypatsetfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatsetfile"
}

func (d *PolicypatsetfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *PolicypatsetfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = PolicypatsetfileDataSourceSchema()
}

func (d *PolicypatsetfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PolicypatsetfileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An imported patset file is only listed by the filtered GET
	// /policypatsetfile?args=imported:true (Pattern 15), not by a plain GET
	// /policypatsetfile/<name>. Fetch that list and match by name.
	name := data.Name.ValueString()

	getResponseData, err := findImportedPolicypatsetfileByName(d.client, name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read policypatsetfile, got error: %s", err))
		return
	}

	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("policypatsetfile %s not found.", name))
		return
	}

	policypatsetfileSetAttrFromGetForDatasource(ctx, &data, getResponseData)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
