package nd6

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	// nd6 has no single-key GET endpoint; readNd6FromApi matches the enumerated
	// array on (neighbor, td, nodeid). A bare passthrough would populate only id,
	// leaving those key attributes null so Read finds nothing and drops the
	// resource. Parse the composite id "neighbor,td,nodeid" into the key
	// attributes. td/nodeid default to 0, so a legacy neighbor-only id (the SDK v2
	// format) is still accepted. neighbor is an IPv6 address (colons, no commas).
	parts := strings.Split(req.ID, ",")
	neighbor := parts[0]
	var td, nodeid int64
	if len(parts) == 3 {
		var err error
		if td, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			resp.Diagnostics.AddError("Invalid import ID for nd6", fmt.Sprintf("td component %q is not an integer in ID %q", parts[1], req.ID))
			return
		}
		if nodeid, err = strconv.ParseInt(parts[2], 10, 64); err != nil {
			resp.Diagnostics.AddError("Invalid import ID for nd6", fmt.Sprintf("nodeid component %q is not an integer in ID %q", parts[2], req.ID))
			return
		}
	} else if len(parts) != 1 {
		resp.Diagnostics.AddError(
			"Invalid import ID for nd6",
			fmt.Sprintf("Expected import ID in the format \"neighbor,td,nodeid\" (or just \"neighbor\"), got %q", req.ID),
		)
		return
	}
	// The default td/nodeid (0) is omitted from NITRO's GET, so a freshly-created
	// resource holds them as null in state. Seed the same shape here (0 -> null)
	// so the imported state matches; readNd6FromApi treats null as 0 when matching
	// and reconciles any non-zero value from the GET response afterward.
	tdAttr := types.Int64Null()
	if td != 0 {
		tdAttr = types.Int64Value(td)
	}
	nodeidAttr := types.Int64Null()
	if nodeid != 0 {
		nodeidAttr = types.Int64Value(nodeid)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("neighbor"), types.StringValue(neighbor))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("td"), tdAttr)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("nodeid"), nodeidAttr)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(fmt.Sprintf("%s,%d,%d", neighbor, td, nodeid)))...)
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
	if !r.readNd6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nd6 not found on ADC immediately after creation")
		}
		return
	}

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

	found := r.readNd6FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Self-healing: the nd6 entry no longer exists on the ADC. Remove it from
	// state so Terraform plans a recreate (SDK v2 d.SetId("") drift contract).
	if !found {
		tflog.Debug(ctx, "nd6 not found on ADC; removing from state for recreation")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Nd6ResourceModel

	// Read Terraform prior state and plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// nd6 exposes no NITRO update endpoint and every attribute is ForceNew /
	// RequiresReplace (matching the SDK v2 resource, which defined no Update
	// callback at all). Terraform therefore never routes changes through
	// Update; this body is a documented no-op that only refreshes state.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nd6; all attributes are RequiresReplace")

	// Read the current state back
	if !r.readNd6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nd6 not found on ADC")
		}
		return
	}

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

	// Build the resource name and delete args for deletion.
	// The NITRO nd6 delete endpoint identifies the entry by neighbor plus a
	// vlan (or vxlan) qualifier and, optionally, td. This mirrors the SDK v2
	// resource's deleteNd6Func exactly for backward compatibility (which used
	// d.GetOk, treating a zero value as unset — hence the != 0 guards).
	resourceId := data.Neighbor.ValueString()
	args := make([]string, 0)
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() && data.Vlan.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("vlan:%d", data.Vlan.ValueInt64()))
	} else if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() && data.Vxlan.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("vxlan:%d", data.Vxlan.ValueInt64()))
	} else {
		args = append(args, "vlan:1")
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() && data.Td.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("td:%d", data.Td.ValueInt64()))
	}
	err := r.client.DeleteResourceWithArgs(service.Nd6.Type(), resourceId, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nd6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nd6 resource")
}

// Helper function to read nd6 data from API. Returns found=false (without adding
// a diagnostic) when the specific nd6 entry no longer exists on the ADC, so
// callers can distinguish drift (Read: remove from state and recreate) from a
// post-write inconsistency (Create/Update: report an error). A genuine
// transport/API error is reported via diags and also returns found=false.
func (r *Nd6Resource) readNd6FromApi(ctx context.Context, data *Nd6ResourceModel, diags *diag.Diagnostics) bool {
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
		return false
	}

	// Resource is missing (no nd6 entries at all on the ADC)
	if len(dataArr) == 0 {
		return false
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

	// Resource is missing (this specific nd6 entry is gone)
	if foundIndex == -1 {
		return false
	}

	nd6SetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
