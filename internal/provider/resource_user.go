// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-billingbox/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
  
// helper to get string fields with diagnostics
func getStringField(m map[string]interface{}, key string, diags *diag.Diagnostics, ctxMsg string) (string, bool) {
	raw, ok := m[key]
	if !ok {
		diags.AddError("Missing field", fmt.Sprintf("%s: %q not found in response", ctxMsg, key))
		return "", false
	}
	str, ok := raw.(string)
	if !ok {
		diags.AddError("Type error", fmt.Sprintf("%s: %q is not a string", ctxMsg, key))
		return "", false
	}
	return str, true
}

var _ resource.Resource = &UserResource{}
var _ resource.ResourceWithImportState = &UserResource{}

func NewUserResource() resource.Resource {
	return &UserResource{}
}

// UserResource defines the resource implementation.
type UserResource struct {
	client *client.Client
}

type UserName struct {
	GivenName       types.String `json:"givenName,omitempty" tfsdk:"given_name"`
	MiddleName      types.String `json:"middleName,omitempty" tfsdk:"middle_name"`
	FamilyName      types.String `json:"familyName,omitempty" tfsdk:"family_name"`
	HonorificPrefix types.String `json:"honorificPrefix,omitempty" tfsdk:"honorific_prefix"`
}

// UserResourceModel describes the resource data model.
type UserResourceModel struct {
	ResourceType types.String `json:"resourceType" tfsdk:"resource_type"`
	ID           types.String `json:"id,omitempty" tfsdk:"id"`
	Password     types.String `json:"password,omitempty" tfsdk:"password"`
	Name         *UserName    `json:"name,omitempty" tfsdk:"name"`
	Meta         types.Object `json:"meta,omitempty" tfsdk:"meta"`
}

func (r *UserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Unique identifier for the User",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Resource type of the User. Always set to 'User'.",
				Default:             stringdefault.StaticString("User"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"password": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "User cleartext password. Only set when you want to change the password.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Components of the user's real name",
				Attributes: map[string]schema.Attribute{
					"given_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The given name of the User",
					},
					"middle_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The middle name(s) of the User",
					},
					"family_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The family name of the User",
					},
					"honorific_prefix": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The honorific prefix(es) of the User",
					},
				},
			},
			// shared metadata from schema_metadata.go
			"meta": MetaAttributes(),
		},
	}
}

func (r *UserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert to API type
	user := &client.User{
		Resource: client.Resource{
			ID: data.ID.ValueString(),
		},
		Password: data.Password.ValueString(),
		Name: &client.UserName{
			GivenName:       data.Name.GivenName.ValueString(),
			MiddleName:      data.Name.MiddleName.ValueString(),
			FamilyName:      data.Name.FamilyName.ValueString(),
			HonorificPrefix: data.Name.HonorificPrefix.ValueString(),
		},
	}

	// Create the user
	result, err := r.client.CreateResource("User", user)
	if err != nil {
		resp.Diagnostics.AddError("Error creating user", err.Error())
		return
	}

	// Update the model with the response data
	data.ID = types.StringValue(result["id"].(string))
	data.ResourceType = types.StringValue("User")
	if meta, ok := result["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{
			"version_id":   types.StringValue(meta["versionId"].(string)),
			"created_at":   types.StringValue(meta["createdAt"].(string)),
			"last_updated": types.StringValue(meta["lastUpdated"].(string)),
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := r.client.GetResource("User", data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", err.Error())
		return
	}

	// Update the model with the response data
	data.ID = types.StringValue(user["id"].(string))
	data.ResourceType = types.StringValue("User")

	// Handle name if present
	if name, ok := user["name"].(map[string]interface{}); ok {
		if data.Name == nil {
			data.Name = &UserName{}
		}
		if v, ok := name["givenName"].(string); ok {
			data.Name.GivenName = types.StringValue(v)
		}
		if v, ok := name["middleName"].(string); ok {
			data.Name.MiddleName = types.StringValue(v)
		}
		if v, ok := name["familyName"].(string); ok {
			data.Name.FamilyName = types.StringValue(v)
		}
		if v, ok := name["honorificPrefix"].(string); ok {
			data.Name.HonorificPrefix = types.StringValue(v)
		}
	}

	// Handle metadata
	if meta, ok := user["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{
			"version_id":   types.StringValue(meta["versionId"].(string)),
			"created_at":   types.StringValue(meta["createdAt"].(string)),
			"last_updated": types.StringValue(meta["lastUpdated"].(string)),
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	// Don't update password from server response to prevent diff
	// The password field will keep its state value unless explicitly changed

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert to API type
	user := &client.User{
		Resource: client.Resource{
			ResourceType: "User",
			ID:           data.ID.ValueString(),
		},
	}

	// Only include password in update if it's explicitly set
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		user.Password = data.Password.ValueString()
	}

	if data.Name != nil {
		user.Name = &client.UserName{
			GivenName:       data.Name.GivenName.ValueString(),
			MiddleName:      data.Name.MiddleName.ValueString(),
			FamilyName:      data.Name.FamilyName.ValueString(),
			HonorificPrefix: data.Name.HonorificPrefix.ValueString(),
		}
	}

	// Update the user
	result, err := r.client.UpdateResource("User", data.ID.ValueString(), user)
	if err != nil {
		resp.Diagnostics.AddError("Error updating user", err.Error())
		return
	}

	// Update the model with the response data
	data.ID = types.StringValue(result["id"].(string))
	data.ResourceType = types.StringValue("User")

	// Handle name if present
	if name, ok := result["name"].(map[string]interface{}); ok {
		if data.Name == nil {
			data.Name = &UserName{}
		}
		if v, ok := name["givenName"].(string); ok {
			data.Name.GivenName = types.StringValue(v)
		}
		if v, ok := name["middleName"].(string); ok {
			data.Name.MiddleName = types.StringValue(v)
		}
		if v, ok := name["familyName"].(string); ok {
			data.Name.FamilyName = types.StringValue(v)
		}
		if v, ok := name["honorificPrefix"].(string); ok {
			data.Name.HonorificPrefix = types.StringValue(v)
		}
	}

	// Handle metadata
	if meta, ok := result["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{
			"version_id":   types.StringValue(meta["versionId"].(string)),
			"created_at":   types.StringValue(meta["createdAt"].(string)),
			"last_updated": types.StringValue(meta["lastUpdated"].(string)),
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteResource("User", data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting user", err.Error())
		return
	}
}

func (r *UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
