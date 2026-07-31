package rnat

import (
	"context"
	"fmt"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RnatClearResource{}
var _ resource.ResourceWithConfigure = (*RnatClearResource)(nil)

func NewRnatClearResource() resource.Resource {
	return &RnatClearResource{}
}

// RnatClearResource defines the resource implementation.
//
// This mirrors the SDK v2 `citrixadc_rnat_clear` resource, which is NOT a plain
// action resource. It manages a *set* of RNAT rules identified by a synthetic
// handle (rnatsname):
//   - Create applies every rule in the `rnat` set with UpdateUnnamedResource
//     (NITRO PUT ?/config/rnat), then stores rnatsname as the Terraform ID.
//   - Update diffs the old and new sets, clearing removed rules with
//     ActOnResource(..., "clear") and (re)applying added rules with
//     UpdateUnnamedResource.
//   - Delete clears every rule in state with ActOnResource(..., "clear").
//
// The SDK v2 Read called FindAllResources and blindly d.Set("rnat", allRnats);
// that write silently failed for the poorly-typed schema (redirectport is a
// TypeBool here but an Integer in NITRO, and GET returns many keys absent from
// the nested schema) and the error was swallowed, so it was effectively a no-op.
// We therefore implement Read as a state-preserving no-op (same choice as the
// sibling servicegroup_servicegroupmemberlist_binding migration), which keeps the
// observable behavior of the SDK resource while remaining framework-safe.
type RnatClearResource struct {
	client *service.NitroClient
}

// RnatClearResourceModel describes the resource data model.
type RnatClearResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Rnatsname types.String `tfsdk:"rnatsname"`
	Rnat      types.Set    `tfsdk:"rnat"`
}

// RnatClearRnatModel describes a single rnat rule entry inside the `rnat` set.
type RnatClearRnatModel struct {
	Aclname      types.String `tfsdk:"aclname"`
	Natip        types.String `tfsdk:"natip"`
	Natip2       types.String `tfsdk:"natip2"`
	Netmask      types.String `tfsdk:"netmask"`
	Network      types.String `tfsdk:"network"`
	Redirectport types.Bool   `tfsdk:"redirectport"`
	Td           types.Int64  `tfsdk:"td"`
}

// rnatClearRnatObjectType returns the object type of a single `rnat` set element.
// It MUST match the tfsdk tags on RnatClearRnatModel exactly.
func rnatClearRnatObjectType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aclname":      types.StringType,
			"natip":        types.StringType,
			"natip2":       types.StringType,
			"netmask":      types.StringType,
			"network":      types.StringType,
			"redirectport": types.BoolType,
			"td":           types.Int64Type,
		},
	}
}

func (r *RnatClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat_clear"
}

func (r *RnatClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnat_clear resource (the rnatsname handle).",
			},
			// rnatsname is Optional+Computed exactly as in SDK v2: when omitted,
			// Create generates a synthetic "tf-rnat-*" handle. Create always
			// resolves it to a known value, so Computed is safe here.
			"rnatsname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name handle for this set of RNAT rules. If omitted, a unique tf-rnat-* value is generated.",
			},
			// The set of RNAT rules managed by this resource. Optional+Computed as
			// in SDK v2; Create always resolves an omitted set to a null value.
			"rnat": schema.SetNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set of RNAT rules to apply. Rules removed from this set are cleared on the appliance.",
				NestedObject: schema.NestedAttributeObject{
					// Nested attributes are Optional (not Computed): Read is a
					// no-op, so a Computed nested attribute would remain unknown
					// after apply (Pattern 13). This does not change the stored
					// state type compared with SDK v2.
					Attributes: map[string]schema.Attribute{
						"aclname": schema.StringAttribute{
							Optional:    true,
							Description: "An extended ACL defined for the RNAT entry.",
						},
						"natip": schema.StringAttribute{
							Optional:    true,
							Description: "Any NetScaler-owned IPv4 address except the NSIP address used to replace source IP addresses of server-generated packets.",
						},
						"natip2": schema.StringAttribute{
							Optional:    true,
							Description: "Secondary NAT IP address (provider-side only; not part of the NITRO rnat object).",
						},
						"netmask": schema.StringAttribute{
							Optional:    true,
							Description: "The subnet mask for the network address.",
						},
						"network": schema.StringAttribute{
							Optional:    true,
							Description: "The network address defined for the RNAT entry.",
						},
						"redirectport": schema.BoolAttribute{
							Optional:    true,
							Description: "Redirect port flag (kept for backward compatibility; not sent to NITRO).",
						},
						"td": schema.Int64Attribute{
							Optional:    true,
							Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity.",
						},
					},
				},
			},
		},
	}
}

func (r *RnatClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat_clear resource")

	// Resolve the rnatsname handle: use the configured value if present,
	// otherwise generate a unique "tf-rnat-*" name (mirrors SDK v2
	// resource.PrefixedUniqueId("tf-rnat-")).
	var rnatName string
	if !data.Rnatsname.IsNull() && !data.Rnatsname.IsUnknown() && data.Rnatsname.ValueString() != "" {
		rnatName = data.Rnatsname.ValueString()
	} else {
		rnatName = fmt.Sprintf("tf-rnat-%d", time.Now().UnixNano())
	}
	data.Rnatsname = types.StringValue(rnatName)

	// Apply every rule in the set. SDK v2 ignored per-rule errors here.
	elems := rnat_clearElementsFromSet(ctx, data.Rnat, &resp.Diagnostics)
	for i := range elems {
		payload := rnat_clearGetThePayload(ctx, &elems[i])
		if err := r.client.UpdateUnnamedResource(service.Rnat.Type(), &payload); err != nil {
			tflog.Debug(ctx, fmt.Sprintf("error applying rnat rule: %s", err))
		}
	}

	// Resolve an omitted (unknown) set to a known null value.
	if data.Rnat.IsUnknown() {
		data.Rnat = types.SetNull(rnatClearRnatObjectType())
	}

	// ID scheme matches SDK v2: the rnatsname handle.
	data.Id = types.StringValue(rnatName)

	tflog.Trace(ctx, "Created rnat_clear resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// SDK v2 Read's FindAllResources + d.Set("rnat", ...) silently failed for the
	// mistyped schema and the error was discarded, making it a de-facto no-op.
	// Preserve prior state.
	tflog.Debug(ctx, "Read is a no-op for rnat_clear; preserving prior state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RnatClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rnat_clear resource")

	oldElems := rnat_clearElementsFromSet(ctx, state.Rnat, &resp.Diagnostics)
	newElems := rnat_clearElementsFromSet(ctx, data.Rnat, &resp.Diagnostics)

	// Rules present in the old set but not the new set are cleared.
	for i := range oldElems {
		if !rnat_clearContains(newElems, oldElems[i]) {
			payload := rnat_clearGetThePayload(ctx, &oldElems[i])
			if err := r.client.ActOnResource(service.Rnat.Type(), payload, "clear"); err != nil {
				tflog.Debug(ctx, fmt.Sprintf("error clearing rnat rule: %s", err))
			}
		}
	}

	// Rules present in the new set but not the old set are (re)applied.
	for i := range newElems {
		if !rnat_clearContains(oldElems, newElems[i]) {
			payload := rnat_clearGetThePayload(ctx, &newElems[i])
			if err := r.client.UpdateUnnamedResource(service.Rnat.Type(), &payload); err != nil {
				tflog.Debug(ctx, fmt.Sprintf("error applying rnat rule: %s", err))
			}
		}
	}

	// Preserve the ID/handle from prior state.
	data.Id = state.Id
	if data.Rnatsname.IsNull() || data.Rnatsname.IsUnknown() {
		data.Rnatsname = state.Rnatsname
	}
	if data.Rnat.IsUnknown() {
		data.Rnat = types.SetNull(rnatClearRnatObjectType())
	}

	tflog.Trace(ctx, "Updated rnat_clear resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat_clear resource")

	// Clear every managed rule. SDK v2 ignored per-rule errors here.
	elems := rnat_clearElementsFromSet(ctx, data.Rnat, &resp.Diagnostics)
	for i := range elems {
		payload := rnat_clearGetThePayload(ctx, &elems[i])
		if err := r.client.ActOnResource(service.Rnat.Type(), payload, "clear"); err != nil {
			tflog.Debug(ctx, fmt.Sprintf("error clearing rnat rule: %s", err))
		}
	}

	tflog.Trace(ctx, "Deleted rnat_clear resource")
}

// rnat_clearElementsFromSet decodes the `rnat` set into a slice, tolerating
// null/unknown sets (returns an empty slice).
func rnat_clearElementsFromSet(ctx context.Context, set types.Set, diags *diag.Diagnostics) []RnatClearRnatModel {
	var elems []RnatClearRnatModel
	if set.IsNull() || set.IsUnknown() {
		return elems
	}
	diags.Append(set.ElementsAs(ctx, &elems, false)...)
	return elems
}

// rnat_clearContains reports whether target appears in list (value equality on
// all fields, mirroring SDK v2 schema.Set difference semantics).
func rnat_clearContains(list []RnatClearRnatModel, target RnatClearRnatModel) bool {
	for _, e := range list {
		if e == target {
			return true
		}
	}
	return false
}

// rnat_clearGetThePayload builds the NITRO rnat object for a single rule.
// It includes ONLY the fields the NITRO rnat add/clear payload accepts and that
// exist on the network.Rnat struct. natip2 has no NITRO field and redirectport
// is a bool in this legacy schema (NITRO expects an Integer), so both are
// excluded from the payload - matching the effective behavior of the SDK v2
// mapstructure decode.
func rnat_clearGetThePayload(ctx context.Context, m *RnatClearRnatModel) network.Rnat {
	tflog.Debug(ctx, "In rnat_clearGetThePayload Function")

	rnat := network.Rnat{}
	if !m.Aclname.IsNull() && !m.Aclname.IsUnknown() {
		rnat.Aclname = m.Aclname.ValueString()
	}
	if !m.Natip.IsNull() && !m.Natip.IsUnknown() {
		rnat.Natip = m.Natip.ValueString()
	}
	if !m.Netmask.IsNull() && !m.Netmask.IsUnknown() {
		rnat.Netmask = m.Netmask.ValueString()
	}
	if !m.Network.IsNull() && !m.Network.IsUnknown() {
		rnat.Network = m.Network.ValueString()
	}
	if !m.Td.IsNull() && !m.Td.IsUnknown() {
		rnat.Td = utils.IntPtr(int(m.Td.ValueInt64()))
	}
	return rnat
}
