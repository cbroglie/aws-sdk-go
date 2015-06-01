// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitoidentityiface provides an interface for the Amazon Cognito Identity.
package cognitoidentityiface

import (
	"github.com/awslabs/aws-sdk-go/service/cognitoidentity"
)

// CognitoIdentityAPI is the interface type for cognitoidentity.CognitoIdentity.
type CognitoIdentityAPI interface {
	CreateIdentityPool(*cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error)

	DeleteIdentityPool(*cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error)

	DescribeIdentity(*cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error)

	DescribeIdentityPool(*cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error)

	GetCredentialsForIdentity(*cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error)

	GetID(*cognitoidentity.GetIDInput) (*cognitoidentity.GetIDOutput, error)

	GetIdentityPoolRoles(*cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error)

	GetOpenIDToken(*cognitoidentity.GetOpenIDTokenInput) (*cognitoidentity.GetOpenIDTokenOutput, error)

	GetOpenIDTokenForDeveloperIdentity(*cognitoidentity.GetOpenIDTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIDTokenForDeveloperIdentityOutput, error)

	ListIdentities(*cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error)

	ListIdentityPools(*cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error)

	LookupDeveloperIdentity(*cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error)

	MergeDeveloperIdentities(*cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)

	SetIdentityPoolRoles(*cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error)

	UnlinkDeveloperIdentity(*cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)

	UnlinkIdentity(*cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error)

	UpdateIdentityPool(*cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error)
}
