// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitoidentity provides a client for Amazon Cognito Identity.
package cognitoidentity

import (
	"sync"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateIdentityPoolRequest generates a request for the CreateIdentityPool operation.
func (c *CognitoIdentity) CreateIdentityPoolRequest(input *CreateIdentityPoolInput) (req *aws.Request, output *IdentityPool) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateIdentityPool == nil {
		opCreateIdentityPool = &aws.Operation{
			Name:       "CreateIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateIdentityPoolInput{}
	}

	req = c.newRequest(opCreateIdentityPool, input, output)
	output = &IdentityPool{}
	req.Data = output
	return
}

// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your AWS account. The limit on identity pools
// is 60 per account.
func (c *CognitoIdentity) CreateIdentityPool(input *CreateIdentityPoolInput) (*IdentityPool, error) {
	req, out := c.CreateIdentityPoolRequest(input)
	err := req.Send()
	return out, err
}

var opCreateIdentityPool *aws.Operation

// DeleteIdentityPoolRequest generates a request for the DeleteIdentityPool operation.
func (c *CognitoIdentity) DeleteIdentityPoolRequest(input *DeleteIdentityPoolInput) (req *aws.Request, output *DeleteIdentityPoolOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteIdentityPool == nil {
		opDeleteIdentityPool = &aws.Operation{
			Name:       "DeleteIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteIdentityPoolInput{}
	}

	req = c.newRequest(opDeleteIdentityPool, input, output)
	output = &DeleteIdentityPoolOutput{}
	req.Data = output
	return
}

// Deletes a user pool. Once a pool is deleted, users will not be able to authenticate
// with the pool.
func (c *CognitoIdentity) DeleteIdentityPool(input *DeleteIdentityPoolInput) (*DeleteIdentityPoolOutput, error) {
	req, out := c.DeleteIdentityPoolRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteIdentityPool *aws.Operation

// DescribeIdentityRequest generates a request for the DescribeIdentity operation.
func (c *CognitoIdentity) DescribeIdentityRequest(input *DescribeIdentityInput) (req *aws.Request, output *IdentityDescription) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeIdentity == nil {
		opDescribeIdentity = &aws.Operation{
			Name:       "DescribeIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeIdentityInput{}
	}

	req = c.newRequest(opDescribeIdentity, input, output)
	output = &IdentityDescription{}
	req.Data = output
	return
}

// Returns metadata related to the given identity, including when the identity
// was created and any associated linked logins.
func (c *CognitoIdentity) DescribeIdentity(input *DescribeIdentityInput) (*IdentityDescription, error) {
	req, out := c.DescribeIdentityRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeIdentity *aws.Operation

// DescribeIdentityPoolRequest generates a request for the DescribeIdentityPool operation.
func (c *CognitoIdentity) DescribeIdentityPoolRequest(input *DescribeIdentityPoolInput) (req *aws.Request, output *IdentityPool) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeIdentityPool == nil {
		opDescribeIdentityPool = &aws.Operation{
			Name:       "DescribeIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeIdentityPoolInput{}
	}

	req = c.newRequest(opDescribeIdentityPool, input, output)
	output = &IdentityPool{}
	req.Data = output
	return
}

// Gets details about a particular identity pool, including the pool name, ID
// description, creation date, and current number of users.
func (c *CognitoIdentity) DescribeIdentityPool(input *DescribeIdentityPoolInput) (*IdentityPool, error) {
	req, out := c.DescribeIdentityPoolRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeIdentityPool *aws.Operation

// GetCredentialsForIdentityRequest generates a request for the GetCredentialsForIdentity operation.
func (c *CognitoIdentity) GetCredentialsForIdentityRequest(input *GetCredentialsForIdentityInput) (req *aws.Request, output *GetCredentialsForIdentityOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetCredentialsForIdentity == nil {
		opGetCredentialsForIdentity = &aws.Operation{
			Name:       "GetCredentialsForIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetCredentialsForIdentityInput{}
	}

	req = c.newRequest(opGetCredentialsForIdentity, input, output)
	output = &GetCredentialsForIdentityOutput{}
	req.Data = output
	return
}

// Returns credentials for the the provided identity ID. Any provided logins
// will be validated against supported login providers. If the token is for
// cognito-identity.amazonaws.com, it will be passed through to AWS Security
// Token Service with the appropriate role for the token.
func (c *CognitoIdentity) GetCredentialsForIdentity(input *GetCredentialsForIdentityInput) (*GetCredentialsForIdentityOutput, error) {
	req, out := c.GetCredentialsForIdentityRequest(input)
	err := req.Send()
	return out, err
}

var opGetCredentialsForIdentity *aws.Operation

// GetIDRequest generates a request for the GetID operation.
func (c *CognitoIdentity) GetIDRequest(input *GetIDInput) (req *aws.Request, output *GetIDOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetID == nil {
		opGetID = &aws.Operation{
			Name:       "GetId",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetIDInput{}
	}

	req = c.newRequest(opGetID, input, output)
	output = &GetIDOutput{}
	req.Data = output
	return
}

// Generates (or retrieves) a Cognito ID. Supplying multiple logins will create
// an implicit linked account.
func (c *CognitoIdentity) GetID(input *GetIDInput) (*GetIDOutput, error) {
	req, out := c.GetIDRequest(input)
	err := req.Send()
	return out, err
}

var opGetID *aws.Operation

// GetIdentityPoolRolesRequest generates a request for the GetIdentityPoolRoles operation.
func (c *CognitoIdentity) GetIdentityPoolRolesRequest(input *GetIdentityPoolRolesInput) (req *aws.Request, output *GetIdentityPoolRolesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetIdentityPoolRoles == nil {
		opGetIdentityPoolRoles = &aws.Operation{
			Name:       "GetIdentityPoolRoles",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetIdentityPoolRolesInput{}
	}

	req = c.newRequest(opGetIdentityPoolRoles, input, output)
	output = &GetIdentityPoolRolesOutput{}
	req.Data = output
	return
}

// Gets the roles for an identity pool.
func (c *CognitoIdentity) GetIdentityPoolRoles(input *GetIdentityPoolRolesInput) (*GetIdentityPoolRolesOutput, error) {
	req, out := c.GetIdentityPoolRolesRequest(input)
	err := req.Send()
	return out, err
}

var opGetIdentityPoolRoles *aws.Operation

// GetOpenIDTokenRequest generates a request for the GetOpenIDToken operation.
func (c *CognitoIdentity) GetOpenIDTokenRequest(input *GetOpenIDTokenInput) (req *aws.Request, output *GetOpenIDTokenOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetOpenIDToken == nil {
		opGetOpenIDToken = &aws.Operation{
			Name:       "GetOpenIdToken",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetOpenIDTokenInput{}
	}

	req = c.newRequest(opGetOpenIDToken, input, output)
	output = &GetOpenIDTokenOutput{}
	req.Data = output
	return
}

// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned by GetId. You can optionally add additional logins for the identity.
// Supplying multiple logins creates an implicit link.
//
// The OpenId token is valid for 15 minutes.
func (c *CognitoIdentity) GetOpenIDToken(input *GetOpenIDTokenInput) (*GetOpenIDTokenOutput, error) {
	req, out := c.GetOpenIDTokenRequest(input)
	err := req.Send()
	return out, err
}

var opGetOpenIDToken *aws.Operation

// GetOpenIDTokenForDeveloperIdentityRequest generates a request for the GetOpenIDTokenForDeveloperIdentity operation.
func (c *CognitoIdentity) GetOpenIDTokenForDeveloperIdentityRequest(input *GetOpenIDTokenForDeveloperIdentityInput) (req *aws.Request, output *GetOpenIDTokenForDeveloperIdentityOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetOpenIDTokenForDeveloperIdentity == nil {
		opGetOpenIDTokenForDeveloperIdentity = &aws.Operation{
			Name:       "GetOpenIdTokenForDeveloperIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetOpenIDTokenForDeveloperIdentityInput{}
	}

	req = c.newRequest(opGetOpenIDTokenForDeveloperIdentity, input, output)
	output = &GetOpenIDTokenForDeveloperIdentityOutput{}
	req.Data = output
	return
}

// Registers (or retrieves) a Cognito IdentityId and an OpenID Connect token
// for a user authenticated by your backend authentication process. Supplying
// multiple logins will create an implicit linked account. You can only specify
// one developer provider as part of the Logins map, which is linked to the
// identity pool. The developer provider is the "domain" by which Cognito will
// refer to your users.
//
// You can use GetOpenIdTokenForDeveloperIdentity to create a new identity
// and to link new logins (that is, user credentials issued by a public provider
// or developer provider) to an existing identity. When you want to create a
// new identity, the IdentityId should be null. When you want to associate a
// new login with an existing authenticated/unauthenticated identity, you can
// do so by providing the existing IdentityId. This API will create the identity
// in the specified IdentityPoolId.
func (c *CognitoIdentity) GetOpenIDTokenForDeveloperIdentity(input *GetOpenIDTokenForDeveloperIdentityInput) (*GetOpenIDTokenForDeveloperIdentityOutput, error) {
	req, out := c.GetOpenIDTokenForDeveloperIdentityRequest(input)
	err := req.Send()
	return out, err
}

var opGetOpenIDTokenForDeveloperIdentity *aws.Operation

// ListIdentitiesRequest generates a request for the ListIdentities operation.
func (c *CognitoIdentity) ListIdentitiesRequest(input *ListIdentitiesInput) (req *aws.Request, output *ListIdentitiesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListIdentities == nil {
		opListIdentities = &aws.Operation{
			Name:       "ListIdentities",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListIdentitiesInput{}
	}

	req = c.newRequest(opListIdentities, input, output)
	output = &ListIdentitiesOutput{}
	req.Data = output
	return
}

// Lists the identities in a pool.
func (c *CognitoIdentity) ListIdentities(input *ListIdentitiesInput) (*ListIdentitiesOutput, error) {
	req, out := c.ListIdentitiesRequest(input)
	err := req.Send()
	return out, err
}

var opListIdentities *aws.Operation

// ListIdentityPoolsRequest generates a request for the ListIdentityPools operation.
func (c *CognitoIdentity) ListIdentityPoolsRequest(input *ListIdentityPoolsInput) (req *aws.Request, output *ListIdentityPoolsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListIdentityPools == nil {
		opListIdentityPools = &aws.Operation{
			Name:       "ListIdentityPools",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListIdentityPoolsInput{}
	}

	req = c.newRequest(opListIdentityPools, input, output)
	output = &ListIdentityPoolsOutput{}
	req.Data = output
	return
}

// Lists all of the Cognito identity pools registered for your account.
func (c *CognitoIdentity) ListIdentityPools(input *ListIdentityPoolsInput) (*ListIdentityPoolsOutput, error) {
	req, out := c.ListIdentityPoolsRequest(input)
	err := req.Send()
	return out, err
}

var opListIdentityPools *aws.Operation

// LookupDeveloperIdentityRequest generates a request for the LookupDeveloperIdentity operation.
func (c *CognitoIdentity) LookupDeveloperIdentityRequest(input *LookupDeveloperIdentityInput) (req *aws.Request, output *LookupDeveloperIdentityOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opLookupDeveloperIdentity == nil {
		opLookupDeveloperIdentity = &aws.Operation{
			Name:       "LookupDeveloperIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &LookupDeveloperIdentityInput{}
	}

	req = c.newRequest(opLookupDeveloperIdentity, input, output)
	output = &LookupDeveloperIdentityOutput{}
	req.Data = output
	return
}

// Retrieves the IdentityID associated with a DeveloperUserIdentifier or the
// list of DeveloperUserIdentifiers associated with an IdentityId for an existing
// identity. Either IdentityID or DeveloperUserIdentifier must not be null.
// If you supply only one of these values, the other value will be searched
// in the database and returned as a part of the response. If you supply both,
// DeveloperUserIdentifier will be matched against IdentityID. If the values
// are verified against the database, the response returns both values and is
// the same as the request. Otherwise a ResourceConflictException is thrown.
func (c *CognitoIdentity) LookupDeveloperIdentity(input *LookupDeveloperIdentityInput) (*LookupDeveloperIdentityOutput, error) {
	req, out := c.LookupDeveloperIdentityRequest(input)
	err := req.Send()
	return out, err
}

var opLookupDeveloperIdentity *aws.Operation

// MergeDeveloperIdentitiesRequest generates a request for the MergeDeveloperIdentities operation.
func (c *CognitoIdentity) MergeDeveloperIdentitiesRequest(input *MergeDeveloperIdentitiesInput) (req *aws.Request, output *MergeDeveloperIdentitiesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opMergeDeveloperIdentities == nil {
		opMergeDeveloperIdentities = &aws.Operation{
			Name:       "MergeDeveloperIdentities",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &MergeDeveloperIdentitiesInput{}
	}

	req = c.newRequest(opMergeDeveloperIdentities, input, output)
	output = &MergeDeveloperIdentitiesOutput{}
	req.Data = output
	return
}

// Merges two users having different IdentityIds, existing in the same identity
// pool, and identified by the same developer provider. You can use this action
// to request that discrete users be merged and identified as a single user
// in the Cognito environment. Cognito associates the given source user (SourceUserIdentifier)
// with the IdentityId of the DestinationUserIdentifier. Only developer-authenticated
// users can be merged. If the users to be merged are associated with the same
// public provider, but as two different users, an exception will be thrown.
func (c *CognitoIdentity) MergeDeveloperIdentities(input *MergeDeveloperIdentitiesInput) (*MergeDeveloperIdentitiesOutput, error) {
	req, out := c.MergeDeveloperIdentitiesRequest(input)
	err := req.Send()
	return out, err
}

var opMergeDeveloperIdentities *aws.Operation

// SetIdentityPoolRolesRequest generates a request for the SetIdentityPoolRoles operation.
func (c *CognitoIdentity) SetIdentityPoolRolesRequest(input *SetIdentityPoolRolesInput) (req *aws.Request, output *SetIdentityPoolRolesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opSetIdentityPoolRoles == nil {
		opSetIdentityPoolRoles = &aws.Operation{
			Name:       "SetIdentityPoolRoles",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &SetIdentityPoolRolesInput{}
	}

	req = c.newRequest(opSetIdentityPoolRoles, input, output)
	output = &SetIdentityPoolRolesOutput{}
	req.Data = output
	return
}

// Sets the roles for an identity pool. These roles are used when making calls
// to GetCredentialsForIdentity action.
func (c *CognitoIdentity) SetIdentityPoolRoles(input *SetIdentityPoolRolesInput) (*SetIdentityPoolRolesOutput, error) {
	req, out := c.SetIdentityPoolRolesRequest(input)
	err := req.Send()
	return out, err
}

var opSetIdentityPoolRoles *aws.Operation

// UnlinkDeveloperIdentityRequest generates a request for the UnlinkDeveloperIdentity operation.
func (c *CognitoIdentity) UnlinkDeveloperIdentityRequest(input *UnlinkDeveloperIdentityInput) (req *aws.Request, output *UnlinkDeveloperIdentityOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUnlinkDeveloperIdentity == nil {
		opUnlinkDeveloperIdentity = &aws.Operation{
			Name:       "UnlinkDeveloperIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &UnlinkDeveloperIdentityInput{}
	}

	req = c.newRequest(opUnlinkDeveloperIdentity, input, output)
	output = &UnlinkDeveloperIdentityOutput{}
	req.Data = output
	return
}

// Unlinks a DeveloperUserIdentifier from an existing identity. Unlinked developer
// users will be considered new identities next time they are seen. If, for
// a given Cognito identity, you remove all federated identities as well as
// the developer user identifier, the Cognito identity becomes inaccessible.
func (c *CognitoIdentity) UnlinkDeveloperIdentity(input *UnlinkDeveloperIdentityInput) (*UnlinkDeveloperIdentityOutput, error) {
	req, out := c.UnlinkDeveloperIdentityRequest(input)
	err := req.Send()
	return out, err
}

var opUnlinkDeveloperIdentity *aws.Operation

// UnlinkIdentityRequest generates a request for the UnlinkIdentity operation.
func (c *CognitoIdentity) UnlinkIdentityRequest(input *UnlinkIdentityInput) (req *aws.Request, output *UnlinkIdentityOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUnlinkIdentity == nil {
		opUnlinkIdentity = &aws.Operation{
			Name:       "UnlinkIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &UnlinkIdentityInput{}
	}

	req = c.newRequest(opUnlinkIdentity, input, output)
	output = &UnlinkIdentityOutput{}
	req.Data = output
	return
}

// Unlinks a federated identity from an existing account. Unlinked logins will
// be considered new identities next time they are seen. Removing the last linked
// login will make this identity inaccessible.
func (c *CognitoIdentity) UnlinkIdentity(input *UnlinkIdentityInput) (*UnlinkIdentityOutput, error) {
	req, out := c.UnlinkIdentityRequest(input)
	err := req.Send()
	return out, err
}

var opUnlinkIdentity *aws.Operation

// UpdateIdentityPoolRequest generates a request for the UpdateIdentityPool operation.
func (c *CognitoIdentity) UpdateIdentityPoolRequest(input *IdentityPool) (req *aws.Request, output *IdentityPool) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUpdateIdentityPool == nil {
		opUpdateIdentityPool = &aws.Operation{
			Name:       "UpdateIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &IdentityPool{}
	}

	req = c.newRequest(opUpdateIdentityPool, input, output)
	output = &IdentityPool{}
	req.Data = output
	return
}

// Updates a user pool.
func (c *CognitoIdentity) UpdateIdentityPool(input *IdentityPool) (*IdentityPool, error) {
	req, out := c.UpdateIdentityPoolRequest(input)
	err := req.Send()
	return out, err
}

var opUpdateIdentityPool *aws.Operation

// Input to the CreateIdentityPool action.
type CreateIdentityPoolInput struct {
	// TRUE if the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities *bool `type:"boolean" required:"true"`

	// The "domain" by which Cognito will refer to your users. This name acts as
	// a placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use
	// letters as well as period (.), underscore (_), and dash (-).
	//
	// Once you have set a developer provider name, you cannot change it. Please
	// take care in setting this parameter.
	DeveloperProviderName *string `type:"string"`

	// A string that you provide.
	IdentityPoolName *string `type:"string" required:"true"`

	// A list of OpendID Connect provider ARNs.
	OpenIDConnectProviderARNs []*string `locationName:"OpenIdConnectProviderARNs" type:"list"`

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders *map[string]*string `type:"map"`

	metadataCreateIdentityPoolInput `json:"-" xml:"-"`
}

type metadataCreateIdentityPoolInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Credentials for the the provided identity ID.
type Credentials struct {
	// The Access Key portion of the credentials.
	AccessKeyID *string `locationName:"AccessKeyId" type:"string"`

	// The date at which these credentials will expire.
	Expiration *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The Secret Access Key portion of the credentials
	SecretKey *string `type:"string"`

	// The Session Token portion of the credentials
	SessionToken *string `type:"string"`

	metadataCredentials `json:"-" xml:"-"`
}

type metadataCredentials struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the DeleteIdentityPool action.
type DeleteIdentityPoolInput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDeleteIdentityPoolInput `json:"-" xml:"-"`
}

type metadataDeleteIdentityPoolInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteIdentityPoolOutput struct {
	metadataDeleteIdentityPoolOutput `json:"-" xml:"-"`
}

type metadataDeleteIdentityPoolOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the DescribeIdentity action.
type DescribeIdentityInput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string" required:"true"`

	metadataDescribeIdentityInput `json:"-" xml:"-"`
}

type metadataDescribeIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the DescribeIdentityPool action.
type DescribeIdentityPoolInput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeIdentityPoolInput `json:"-" xml:"-"`
}

type metadataDescribeIdentityPoolInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the GetCredentialsForIdentity action.
type GetCredentialsForIdentityInput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	Logins *map[string]*string `type:"map"`

	metadataGetCredentialsForIdentityInput `json:"-" xml:"-"`
}

type metadataGetCredentialsForIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a successful GetCredentialsForIdentity operation.
type GetCredentialsForIdentityOutput struct {
	// Credentials for the the provided identity ID.
	Credentials *Credentials `type:"structure"`

	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	metadataGetCredentialsForIdentityOutput `json:"-" xml:"-"`
}

type metadataGetCredentialsForIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the GetId action.
type GetIDInput struct {
	// A standard AWS account ID (9+ digits).
	AccountID *string `locationName:"AccountId" type:"string"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	//
	// The available provider names for Logins are as follows:  Facebook: graph.facebook.com
	//  Google: accounts.google.com  Amazon: www.amazon.com
	Logins *map[string]*string `type:"map"`

	metadataGetIDInput `json:"-" xml:"-"`
}

type metadataGetIDInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a GetId request.
type GetIDOutput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	metadataGetIDOutput `json:"-" xml:"-"`
}

type metadataGetIDOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the GetIdentityPoolRoles action.
type GetIdentityPoolRolesInput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	metadataGetIdentityPoolRolesInput `json:"-" xml:"-"`
}

type metadataGetIdentityPoolRolesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a successful GetIdentityPoolRoles operation.
type GetIdentityPoolRolesOutput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// The map of roles associated with this pool. Currently only authenticated
	// and unauthenticated roles are supported.
	Roles *map[string]*string `type:"map"`

	metadataGetIdentityPoolRolesOutput `json:"-" xml:"-"`
}

type metadataGetIdentityPoolRolesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the GetOpenIdTokenForDeveloperIdentity action.
type GetOpenIDTokenForDeveloperIdentityInput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	// Each name-value pair represents a user from a public provider or developer
	// provider. If the user is from a developer provider, the name-value pair will
	// follow the syntax "developer_provider_name": "developer_user_identifier".
	// The developer provider is the "domain" by which Cognito will refer to your
	// users; you provided this domain while creating/updating the identity pool.
	// The developer user identifier is an identifier from your backend that uniquely
	// identifies a user. When you create an identity pool, you can specify the
	// supported logins.
	Logins *map[string]*string `type:"map" required:"true"`

	// The expiration time of the token, in seconds. You can specify a custom expiration
	// time for the token so that you can cache it. If you don't provide an expiration
	// time, the token is valid for 15 minutes. You can exchange the token with
	// Amazon STS for temporary AWS credentials, which are valid for a maximum of
	// one hour. The maximum token duration you can set is 24 hours. You should
	// take care in setting the expiration time for a token, as there are significant
	// security implications: an attacker could use a leaked token to access your
	// AWS resources for the token's duration.
	TokenDuration *int64 `type:"long"`

	metadataGetOpenIDTokenForDeveloperIdentityInput `json:"-" xml:"-"`
}

type metadataGetOpenIDTokenForDeveloperIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a successful GetOpenIdTokenForDeveloperIdentity request.
type GetOpenIDTokenForDeveloperIdentityOutput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// An OpenID token.
	Token *string `type:"string"`

	metadataGetOpenIDTokenForDeveloperIdentityOutput `json:"-" xml:"-"`
}

type metadataGetOpenIDTokenForDeveloperIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the GetOpenIdToken action.
type GetOpenIDTokenInput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	Logins *map[string]*string `type:"map"`

	metadataGetOpenIDTokenInput `json:"-" xml:"-"`
}

type metadataGetOpenIDTokenInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a successful GetOpenIdToken request.
type GetOpenIDTokenOutput struct {
	// A unique identifier in the format REGION:GUID. Note that the IdentityId returned
	// may not match the one passed on input.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// An OpenID token, valid for 15 minutes.
	Token *string `type:"string"`

	metadataGetOpenIDTokenOutput `json:"-" xml:"-"`
}

type metadataGetOpenIDTokenOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A description of the identity.
type IdentityDescription struct {
	// Date on which the identity was created.
	CreationDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// Date on which the identity was last modified.
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	Logins []*string `type:"list"`

	metadataIdentityDescription `json:"-" xml:"-"`
}

type metadataIdentityDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

// An object representing a Cognito identity pool.
type IdentityPool struct {
	// TRUE if the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities *bool `type:"boolean" required:"true"`

	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string `type:"string"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// A string that you provide.
	IdentityPoolName *string `type:"string" required:"true"`

	// A list of OpendID Connect provider ARNs.
	OpenIDConnectProviderARNs []*string `locationName:"OpenIdConnectProviderARNs" type:"list"`

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders *map[string]*string `type:"map"`

	metadataIdentityPool `json:"-" xml:"-"`
}

type metadataIdentityPool struct {
	SDKShapeTraits bool `type:"structure"`
}

// A description of the identity pool.
type IdentityPoolShortDescription struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// A string that you provide.
	IdentityPoolName *string `type:"string"`

	metadataIdentityPoolShortDescription `json:"-" xml:"-"`
}

type metadataIdentityPoolShortDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the ListIdentities action.
type ListIdentitiesInput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// The maximum number of identities to return.
	MaxResults *int64 `type:"integer" required:"true"`

	// A pagination token.
	NextToken *string `type:"string"`

	metadataListIdentitiesInput `json:"-" xml:"-"`
}

type metadataListIdentitiesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The response to a ListIdentities request.
type ListIdentitiesOutput struct {
	// An object containing a set of identities and associated mappings.
	Identities []*IdentityDescription `type:"list"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// A pagination token.
	NextToken *string `type:"string"`

	metadataListIdentitiesOutput `json:"-" xml:"-"`
}

type metadataListIdentitiesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the ListIdentityPools action.
type ListIdentityPoolsInput struct {
	// The maximum number of identities to return.
	MaxResults *int64 `type:"integer" required:"true"`

	// A pagination token.
	NextToken *string `type:"string"`

	metadataListIdentityPoolsInput `json:"-" xml:"-"`
}

type metadataListIdentityPoolsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The result of a successful ListIdentityPools action.
type ListIdentityPoolsOutput struct {
	// The identity pools returned by the ListIdentityPools action.
	IdentityPools []*IdentityPoolShortDescription `type:"list"`

	// A pagination token.
	NextToken *string `type:"string"`

	metadataListIdentityPoolsOutput `json:"-" xml:"-"`
}

type metadataListIdentityPoolsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the LookupDeveloperIdentityInput action.
type LookupDeveloperIdentityInput struct {
	// A unique ID used by your backend authentication process to identify a user.
	// Typically, a developer identity provider would issue many developer user
	// identifiers, in keeping with the number of users.
	DeveloperUserIdentifier *string `type:"string"`

	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// The maximum number of identities to return.
	MaxResults *int64 `type:"integer"`

	// A pagination token. The first call you make will have NextToken set to null.
	// After that the service will return NextToken values as needed. For example,
	// let's say you make a request with MaxResults set to 10, and there are 20
	// matches in the database. The service will return a pagination token as a
	// part of the response. This token can be used to call the API again and get
	// results starting from the 11th match.
	NextToken *string `type:"string"`

	metadataLookupDeveloperIdentityInput `json:"-" xml:"-"`
}

type metadataLookupDeveloperIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a successful LookupDeveloperIdentity action.
type LookupDeveloperIdentityOutput struct {
	// This is the list of developer user identifiers associated with an identity
	// ID. Cognito supports the association of multiple developer user identifiers
	// with an identity ID.
	DeveloperUserIdentifierList []*string `type:"list"`

	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// A pagination token. The first call you make will have NextToken set to null.
	// After that the service will return NextToken values as needed. For example,
	// let's say you make a request with MaxResults set to 10, and there are 20
	// matches in the database. The service will return a pagination token as a
	// part of the response. This token can be used to call the API again and get
	// results starting from the 11th match.
	NextToken *string `type:"string"`

	metadataLookupDeveloperIdentityOutput `json:"-" xml:"-"`
}

type metadataLookupDeveloperIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the MergeDeveloperIdentities action.
type MergeDeveloperIdentitiesInput struct {
	// User identifier for the destination user. The value should be a DeveloperUserIdentifier.
	DestinationUserIdentifier *string `type:"string" required:"true"`

	// The "domain" by which Cognito will refer to your users. This is a (pseudo)
	// domain name that you provide while creating an identity pool. This name acts
	// as a placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use
	// letters as well as period (.), underscore (_), and dash (-).
	DeveloperProviderName *string `type:"string" required:"true"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// User identifier for the source user. The value should be a DeveloperUserIdentifier.
	SourceUserIdentifier *string `type:"string" required:"true"`

	metadataMergeDeveloperIdentitiesInput `json:"-" xml:"-"`
}

type metadataMergeDeveloperIdentitiesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned in response to a successful MergeDeveloperIdentities action.
type MergeDeveloperIdentitiesOutput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	metadataMergeDeveloperIdentitiesOutput `json:"-" xml:"-"`
}

type metadataMergeDeveloperIdentitiesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the SetIdentityPoolRoles action.
type SetIdentityPoolRolesInput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	// The map of roles associated with this pool. Currently only authenticated
	// and unauthenticated roles are supported.
	Roles *map[string]*string `type:"map" required:"true"`

	metadataSetIdentityPoolRolesInput `json:"-" xml:"-"`
}

type metadataSetIdentityPoolRolesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityPoolRolesOutput struct {
	metadataSetIdentityPoolRolesOutput `json:"-" xml:"-"`
}

type metadataSetIdentityPoolRolesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the UnlinkDeveloperIdentity action.
type UnlinkDeveloperIdentityInput struct {
	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string `type:"string" required:"true"`

	// A unique ID used by your backend authentication process to identify a user.
	DeveloperUserIdentifier *string `type:"string" required:"true"`

	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string" required:"true"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataUnlinkDeveloperIdentityInput `json:"-" xml:"-"`
}

type metadataUnlinkDeveloperIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UnlinkDeveloperIdentityOutput struct {
	metadataUnlinkDeveloperIdentityOutput `json:"-" xml:"-"`
}

type metadataUnlinkDeveloperIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Input to the UnlinkIdentity action.
type UnlinkIdentityInput struct {
	// A unique identifier in the format REGION:GUID.
	IdentityID *string `locationName:"IdentityId" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	Logins *map[string]*string `type:"map" required:"true"`

	// Provider names to unlink from this identity.
	LoginsToRemove []*string `type:"list" required:"true"`

	metadataUnlinkIdentityInput `json:"-" xml:"-"`
}

type metadataUnlinkIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UnlinkIdentityOutput struct {
	metadataUnlinkIdentityOutput `json:"-" xml:"-"`
}

type metadataUnlinkIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}
