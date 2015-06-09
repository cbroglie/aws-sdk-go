// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package directoryserviceiface provides an interface for the AWS Directory Service.
package directoryserviceiface

import (
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

// DirectoryServiceAPI is the interface type for directoryservice.DirectoryService.
type DirectoryServiceAPI interface {
	ConnectDirectory(*directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)

	CreateAlias(*directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)

	CreateComputer(*directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)

	CreateDirectory(*directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)

	CreateSnapshot(*directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)

	DeleteDirectory(*directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)

	DeleteSnapshot(*directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)

	DescribeDirectories(*directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)

	DescribeSnapshots(*directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)

	DisableRadius(*directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)

	DisableSSO(*directoryservice.DisableSSOInput) (*directoryservice.DisableSSOOutput, error)

	EnableRadius(*directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)

	EnableSSO(*directoryservice.EnableSSOInput) (*directoryservice.EnableSSOOutput, error)

	GetDirectoryLimits(*directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)

	GetSnapshotLimits(*directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)

	RestoreFromSnapshot(*directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)

	UpdateRadius(*directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)
}
