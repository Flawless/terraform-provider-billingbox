package provider

import (
   "context"
   "fmt"
   "net/http"

   "github.com/hashicorp/terraform-plugin-framework/path"
   "github.com/hashicorp/terraform-plugin-framework/resource"
   "github.com/hashicorp/terraform-plugin-framework/resource/schema"
   "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
   "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
   "github.com/hashicorp/terraform-plugin-framework/types"

   // your shared metadata helpers
   "github.com/hashicorp/billingbox/internal/provider/schema_metadata"
)

// UserResource implements the Terraform resource for User.
type UserResource struct {
   client *http.Client
}

// NewUserResource returns a UserResource.
func NewUserResource() resource.Resource {
   return &UserResource{}
}

// UserResourceModel describes the User resource data model.
type UserResourceModel struct {
   Id        types.String         `tfsdk:"id"`
   UserName  types.String         `tfsdk:"user_name"`
   Active    types.Bool           `tfsdk:"active"`
   Locale    types.String         `tfsdk:"locale"`
   TwoFactor types.Bool           `tfsdk:"two_factor"`
   // … add other top‐level attributes here …

   Metadata  schema_metadata.MetadataModel `tfsdk:"metadata"`
}

func (r *UserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
   resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
   resp.Schema = schema.Schema{
      MarkdownDescription: "User resource",
      Attributes: map[string]schema.Attribute{
         "id": schema.StringAttribute{
            Computed:            true,
            MarkdownDescription: "Unique identifier for the User",
            PlanModifiers: []planmodifier.String{
               stringplanmodifier.UseStateForUnknown(),
            },
         },
         "user_name": schema.StringAttribute{
            Required:            true,
            MarkdownDescription: "Unique username of the User",
         },
         "active": schema.BoolAttribute{
            Optional:            true,
            MarkdownDescription: "Indicates the User's administrative status",
         },
         "locale": schema.StringAttribute{
            Optional:            true,
            MarkdownDescription: "Default localization setting for the User",
         },
         "two_factor": schema.BoolAttribute{
            Optional:            true,
            MarkdownDescription: "Two‐factor settings",
         },
         // … define additional attributes as needed …

         // shared metadata nested block
         "metadata": schema.SingleNestedAttribute{
            Computed:            true,
            MarkdownDescription: "Server‐side resource metadata",
            Attributes:         schema_metadata.MetadataAttributes(),
            PlanModifiers: []planmodifier.Object{
               planmodifier.SuppressChanges(), // no diff on metadata
            },
         },
      },
   }
}

func (r *UserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
   if req.ProviderData == nil {
      return
   }
   client, ok := req.ProviderData.(*http.Client)
   if !ok {
      resp.Diagnostics.AddError("Unexpected Resource Configure Type",
         fmt.Sprintf("Expected *http.Client, got: %T", req.ProviderData))
      return
   }
   r.client = client
}

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
   var plan UserResourceModel
   resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
   if resp.Diagnostics.HasError() {
      return
   }
   // TODO: call your API to create the User, parse response
   var state UserResourceModel
   state.Id = types.StringValue("generated-id")
   state.UserName = plan.UserName
   state.Active = plan.Active
   state.Locale = plan.Locale
   state.TwoFactor = plan.TwoFactor
   state.Metadata = schema_metadata.MetadataFromResponse(/* API‐response metadata */)
   resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
   var state UserResourceModel
   resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
   if resp.Diagnostics.HasError() {
      return
   }
   // TODO: call your API to refresh the User by state.Id
   resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
   var plan UserResourceModel
   resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
   if resp.Diagnostics.HasError() {
      return
   }
   // TODO: call your API to update the User
   var state UserResourceModel
   state.Id = plan.Id
   state.UserName = plan.UserName
   // … copy other fields …
   state.Metadata = schema_metadata.MetadataFromResponse(/* API‐response metadata */)
   resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
   var state UserResourceModel
   resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
   if resp.Diagnostics.HasError() {
      return
   }
   // TODO: call your API to delete the User by state.Id
}

func (r *UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
   resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
