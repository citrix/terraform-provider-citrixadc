package linkset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*LinksetDataSource)(nil)

func LInksetDataSource() datasource.DataSource {
	return &LinksetDataSource{}
}

type LinksetDataSource struct {
	client *service.NitroClient
}

func (d *LinksetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkset"
}

func (d *LinksetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LinksetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LinksetDataSourceSchema()
}

func (d *LinksetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LinksetDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	id_Name := data.Linksetid.ValueString()

	getResponseData, err := d.client.FindResource(service.Linkset.Type(), id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read linkset, got error: %s", err))
		return
	}

	linksetDataSourceSetAttrFromGet(ctx, &data, getResponseData)

	// Populate the interfacebinding convenience block from the ADC.
	if err := linksetDataSourceReadInterfaceBindings(ctx, d.client, &data, id_Name); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read linkset interface bindings, got error: %s", err))
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// linksetDataSourceReadInterfaceBindings reads the interface bindings for the
// linkset and populates the interfacebinding set on the data-source model.
// Matching the SDK v2 behavior, a "not found" / error while listing bindings is
// treated as "no bindings" (empty set) rather than a hard failure.
func linksetDataSourceReadInterfaceBindings(ctx context.Context, client *service.NitroClient, data *LinksetDataSourceModel, linksetName string) error {
	bindings, err := client.FindResourceArray(service.Linkset_interface_binding.Type(), linksetName)
	if err != nil {
		bindings = []map[string]interface{}{}
	}

	processedBindings := make([]string, 0, len(bindings))
	for _, val := range bindings {
		if ifnum, ok := val["ifnum"].(string); ok {
			processedBindings = append(processedBindings, ifnum)
		}
	}

	interfaceSet, diags := types.SetValueFrom(ctx, types.StringType, processedBindings)
	if diags.HasError() {
		return fmt.Errorf("error converting interface bindings to set")
	}
	data.Interfacebinding = interfaceSet

	return nil
}
