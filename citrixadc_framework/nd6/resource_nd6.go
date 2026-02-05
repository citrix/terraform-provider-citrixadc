package nd6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Nd6Resource{}
var _ resource.ResourceWithConfigure = (*Nd6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nd6Resource)(nil)

func NewNd6Resource() resource.Resource {
	return &Nd6Resource{}
}

// Nd6Resource defines the resource implementation.
type Nd6Resource struct {
	client *service.NitroClient
}

func (r *Nd6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nd6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6"
}

func (r *Nd6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nd6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nd6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nd6 resource")

	// Build payload as a map to properly control which fields are sent
	payload := map[string]interface{}{
		"neighbor": data.Neighbor.ValueString(),
		"mac":      data.Mac.ValueString(),
	}

	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		payload["ifnum"] = data.Ifnum.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		payload["nodeid"] = int(data.Nodeid.ValueInt64())
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		payload["td"] = int(data.Td.ValueInt64())
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		payload["vlan"] = int(data.Vlan.ValueInt64())
	}
	if !data.Vtep.IsNull() && !data.Vtep.IsUnknown() {
		payload["vtep"] = data.Vtep.ValueString()
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		payload["vxlan"] = int(data.Vxlan.ValueInt64())
	}

	// Make API call
	_, err := r.client.AddResource(service.Nd6.Type(), data.Neighbor.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nd6, got error: %s", err))
		return
	}

	// Generate ID based on neighbor address (the primary identifier)
	data.Id = data.Neighbor

	tflog.Trace(ctx, "Created nd6 resource")

	// Read the updated state back
	r.readNd6FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nd6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nd6 resource")

	r.readNd6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Nd6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nd6 resource")

	// Create API request body from the model
	// nd6 := nd6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nd6.Type(), &nd6)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nd6, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nd6 resource")

	// Read the updated state back
	r.readNd6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nd6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nd6 resource")

	// Build the resource name for deletion
	resourceId := data.Neighbor.ValueString()
	err := r.client.DeleteResource(service.Nd6.Type(), resourceId)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nd6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nd6 resource")
}

// Helper function to read nd6 data from API
func (r *Nd6Resource) readNd6FromApi(ctx context.Context, data *Nd6ResourceModel, diags *diag.Diagnostics) {
	neighbor_Name := data.Neighbor.ValueString()

	// Default to "0" for td and nodeid if not set
	td_Name := "0"
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		td_Name = fmt.Sprintf("%d", data.Td.ValueInt64())
	}

	nodeid_Name := "0"
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		nodeid_Name = fmt.Sprintf("%d", data.Nodeid.ValueInt64())
	}

	findParams := service.FindParams{
		ResourceType:             service.Nd6.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nd6, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "nd6 returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if v["neighbor"].(string) != neighbor_Name {
			match = false
		}

		// Handle td - it might be nil or a string
		tdVal := "0"
		if v["td"] != nil {
			tdVal = v["td"].(string)
		}
		if tdVal != td_Name {
			match = false
		}

		// Handle nodeid - it might be nil or a string
		nodeidVal := "0"
		if v["nodeid"] != nil {
			nodeidVal = v["nodeid"].(string)
		}
		if nodeidVal != nodeid_Name {
			match = false
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("nd6 with neighbor %s not found", neighbor_Name))
		return
	}

	nd6SetAttrFromGet(ctx, data, dataArr[foundIndex])
}
