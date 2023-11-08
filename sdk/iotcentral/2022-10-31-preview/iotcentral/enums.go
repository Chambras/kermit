package iotcentral

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

        // CapabilityAggregateType enumerates the values for capability aggregate type.
    type CapabilityAggregateType string

    const (
            // CapabilityAggregateTypeAvg Average of the capability data
        CapabilityAggregateTypeAvg CapabilityAggregateType = "avg"
            // CapabilityAggregateTypeCount Count of the capability data
        CapabilityAggregateTypeCount CapabilityAggregateType = "count"
            // CapabilityAggregateTypeMax Maximum of the capability data
        CapabilityAggregateTypeMax CapabilityAggregateType = "max"
            // CapabilityAggregateTypeMin Minimum of the capability data
        CapabilityAggregateTypeMin CapabilityAggregateType = "min"
            // CapabilityAggregateTypeSum Sum of the capability data
        CapabilityAggregateTypeSum CapabilityAggregateType = "sum"
            )
    // PossibleCapabilityAggregateTypeValues returns an array of possible values for the CapabilityAggregateType const type.
    func PossibleCapabilityAggregateTypeValues() []CapabilityAggregateType {
        return []CapabilityAggregateType{CapabilityAggregateTypeAvg,CapabilityAggregateTypeCount,CapabilityAggregateTypeMax,CapabilityAggregateTypeMin,CapabilityAggregateTypeSum}
    }

        // CertificateEntry enumerates the values for certificate entry.
    type CertificateEntry string

    const (
            // CertificateEntryPrimary The entry of primary certificate
        CertificateEntryPrimary CertificateEntry = "primary"
            // CertificateEntrySecondary The entry of secondary certificate
        CertificateEntrySecondary CertificateEntry = "secondary"
            )
    // PossibleCertificateEntryValues returns an array of possible values for the CertificateEntry const type.
    func PossibleCertificateEntryValues() []CertificateEntry {
        return []CertificateEntry{CertificateEntryPrimary,CertificateEntrySecondary}
    }

        // DestinationSource enumerates the values for destination source.
    type DestinationSource string

    const (
            // DestinationSourceAudit Destination source from audit logs
        DestinationSourceAudit DestinationSource = "audit"
            // DestinationSourceDeviceConnectivity Destination source from device connectivity
        DestinationSourceDeviceConnectivity DestinationSource = "deviceConnectivity"
            // DestinationSourceDeviceLifecycle Destination source from device lifecycle
        DestinationSourceDeviceLifecycle DestinationSource = "deviceLifecycle"
            // DestinationSourceDeviceTemplateLifecycle Destination source from device template lifecycle
        DestinationSourceDeviceTemplateLifecycle DestinationSource = "deviceTemplateLifecycle"
            // DestinationSourceProperties Destination source from device properties
        DestinationSourceProperties DestinationSource = "properties"
            // DestinationSourceTelemetry Destination source from device telemetry
        DestinationSourceTelemetry DestinationSource = "telemetry"
            )
    // PossibleDestinationSourceValues returns an array of possible values for the DestinationSource const type.
    func PossibleDestinationSourceValues() []DestinationSource {
        return []DestinationSource{DestinationSourceAudit,DestinationSourceDeviceConnectivity,DestinationSourceDeviceLifecycle,DestinationSourceDeviceTemplateLifecycle,DestinationSourceProperties,DestinationSourceTelemetry}
    }

        // DeviceType enumerates the values for device type.
    type DeviceType string

    const (
            // DeviceTypeIotEdge The edge device type
        DeviceTypeIotEdge DeviceType = "iotEdge"
            )
    // PossibleDeviceTypeValues returns an array of possible values for the DeviceType const type.
    func PossibleDeviceTypeValues() []DeviceType {
        return []DeviceType{DeviceTypeIotEdge}
    }

        // EnrollmentGroupType enumerates the values for enrollment group type.
    type EnrollmentGroupType string

    const (
            // EnrollmentGroupTypeIoTEdgedevices IoT Edge devices
        EnrollmentGroupTypeIoTEdgedevices EnrollmentGroupType = "iotEdge"
            // EnrollmentGroupTypeIoTdevices Regular (non-Edge) IoT devices
        EnrollmentGroupTypeIoTdevices EnrollmentGroupType = "iot"
            )
    // PossibleEnrollmentGroupTypeValues returns an array of possible values for the EnrollmentGroupType const type.
    func PossibleEnrollmentGroupTypeValues() []EnrollmentGroupType {
        return []EnrollmentGroupType{EnrollmentGroupTypeIoTEdgedevices,EnrollmentGroupTypeIoTdevices}
    }

        // FileUploadState enumerates the values for file upload state.
    type FileUploadState string

    const (
                // FileUploadStateDeleting ...
        FileUploadStateDeleting FileUploadState = "deleting"
                // FileUploadStateFailed ...
        FileUploadStateFailed FileUploadState = "failed"
                // FileUploadStatePending ...
        FileUploadStatePending FileUploadState = "pending"
                // FileUploadStateSucceeded ...
        FileUploadStateSucceeded FileUploadState = "succeeded"
                // FileUploadStateUpdating ...
        FileUploadStateUpdating FileUploadState = "updating"
            )
    // PossibleFileUploadStateValues returns an array of possible values for the FileUploadState const type.
    func PossibleFileUploadStateValues() []FileUploadState {
        return []FileUploadState{FileUploadStateDeleting,FileUploadStateFailed,FileUploadStatePending,FileUploadStateSucceeded,FileUploadStateUpdating}
    }

        // ImageTileTextUnits enumerates the values for image tile text units.
    type ImageTileTextUnits string

    const (
            // ImageTileTextUnitsPixels Set the text unit to pixels
        ImageTileTextUnitsPixels ImageTileTextUnits = "px"
            )
    // PossibleImageTileTextUnitsValues returns an array of possible values for the ImageTileTextUnits const type.
    func PossibleImageTileTextUnitsValues() []ImageTileTextUnits {
        return []ImageTileTextUnits{ImageTileTextUnitsPixels}
    }

        // JobBatchType enumerates the values for job batch type.
    type JobBatchType string

    const (
            // JobBatchTypeNumber Job Batching based on number of devices.
        JobBatchTypeNumber JobBatchType = "number"
            // JobBatchTypePercentage Job Batching based percentage of total applied devices.
        JobBatchTypePercentage JobBatchType = "percentage"
            )
    // PossibleJobBatchTypeValues returns an array of possible values for the JobBatchType const type.
    func PossibleJobBatchTypeValues() []JobBatchType {
        return []JobBatchType{JobBatchTypeNumber,JobBatchTypePercentage}
    }

        // JobCancellationThresholdType enumerates the values for job cancellation threshold type.
    type JobCancellationThresholdType string

    const (
            // JobCancellationThresholdTypeNumber Job cancellation threshold based on specified number of devices.
        JobCancellationThresholdTypeNumber JobCancellationThresholdType = "number"
            // JobCancellationThresholdTypePercentage Job cancellation threshold based on percentage of total devices.
        JobCancellationThresholdTypePercentage JobCancellationThresholdType = "percentage"
            )
    // PossibleJobCancellationThresholdTypeValues returns an array of possible values for the JobCancellationThresholdType const type.
    func PossibleJobCancellationThresholdTypeValues() []JobCancellationThresholdType {
        return []JobCancellationThresholdType{JobCancellationThresholdTypeNumber,JobCancellationThresholdTypePercentage}
    }

        // JobRecurrence enumerates the values for job recurrence.
    type JobRecurrence string

    const (
            // JobRecurrenceDaily The job will run once daily
        JobRecurrenceDaily JobRecurrence = "daily"
            // JobRecurrenceMonthly The job will run once every month
        JobRecurrenceMonthly JobRecurrence = "monthly"
            // JobRecurrenceWeekly The job will run once every week
        JobRecurrenceWeekly JobRecurrence = "weekly"
            )
    // PossibleJobRecurrenceValues returns an array of possible values for the JobRecurrence const type.
    func PossibleJobRecurrenceValues() []JobRecurrence {
        return []JobRecurrence{JobRecurrenceDaily,JobRecurrenceMonthly,JobRecurrenceWeekly}
    }

        // OAuthRequestType enumerates the values for o auth request type.
    type OAuthRequestType string

    const (
            // OAuthRequestTypeAuto Use automatic Content-Type for token request
        OAuthRequestTypeAuto OAuthRequestType = "auto"
            // OAuthRequestTypeJSON Content-Type as JSON for token request
        OAuthRequestTypeJSON OAuthRequestType = "json"
            // OAuthRequestTypeURLEncoded Content-Type as UrlEncoded for token request
        OAuthRequestTypeURLEncoded OAuthRequestType = "urlencoded"
            )
    // PossibleOAuthRequestTypeValues returns an array of possible values for the OAuthRequestType const type.
    func PossibleOAuthRequestTypeValues() []OAuthRequestType {
        return []OAuthRequestType{OAuthRequestTypeAuto,OAuthRequestTypeJSON,OAuthRequestTypeURLEncoded}
    }

        // TileTextSizeUnit enumerates the values for tile text size unit.
    type TileTextSizeUnit string

    const (
            // TileTextSizeUnitPoints Set the text unit to points
        TileTextSizeUnitPoints TileTextSizeUnit = "pt"
            )
    // PossibleTileTextSizeUnitValues returns an array of possible values for the TileTextSizeUnit const type.
    func PossibleTileTextSizeUnitValues() []TileTextSizeUnit {
        return []TileTextSizeUnit{TileTextSizeUnitPoints}
    }

        // TileTimeRangeDuration enumerates the values for tile time range duration.
    type TileTimeRangeDuration string

    const (
            // TileTimeRangeDurationOneSpaceDay 1 Day
        TileTimeRangeDurationOneSpaceDay TileTimeRangeDuration = "P1D"
            // TileTimeRangeDurationOneSpaceHour 1 Hour
        TileTimeRangeDurationOneSpaceHour TileTimeRangeDuration = "PT1H"
            // TileTimeRangeDurationOneSpaceWeek 1 Week
        TileTimeRangeDurationOneSpaceWeek TileTimeRangeDuration = "P1W"
            // TileTimeRangeDurationOneTwoSpaceHours 12 Hours
        TileTimeRangeDurationOneTwoSpaceHours TileTimeRangeDuration = "PT12H"
            // TileTimeRangeDurationThreeZeroSpaceDays 30 Days
        TileTimeRangeDurationThreeZeroSpaceDays TileTimeRangeDuration = "P30D"
            // TileTimeRangeDurationThreeZeroSpaceMinutes 30 Minutes
        TileTimeRangeDurationThreeZeroSpaceMinutes TileTimeRangeDuration = "PT30M"
            )
    // PossibleTileTimeRangeDurationValues returns an array of possible values for the TileTimeRangeDuration const type.
    func PossibleTileTimeRangeDurationValues() []TileTimeRangeDuration {
        return []TileTimeRangeDuration{TileTimeRangeDurationOneSpaceDay,TileTimeRangeDurationOneSpaceHour,TileTimeRangeDurationOneSpaceWeek,TileTimeRangeDurationOneTwoSpaceHours,TileTimeRangeDurationThreeZeroSpaceDays,TileTimeRangeDurationThreeZeroSpaceMinutes}
    }

        // TileTimeRangeResolution enumerates the values for tile time range resolution.
    type TileTimeRangeResolution string

    const (
            // TileTimeRangeResolutionFiveSpaceMinutes 5 Minutes
        TileTimeRangeResolutionFiveSpaceMinutes TileTimeRangeResolution = "PT5M"
            // TileTimeRangeResolutionOneSpaceDay 1 Day
        TileTimeRangeResolutionOneSpaceDay TileTimeRangeResolution = "P1D"
            // TileTimeRangeResolutionOneSpaceHour 1 Hour
        TileTimeRangeResolutionOneSpaceHour TileTimeRangeResolution = "PT1H"
            // TileTimeRangeResolutionOneSpaceMinute 1 Minute
        TileTimeRangeResolutionOneSpaceMinute TileTimeRangeResolution = "PT1M"
            // TileTimeRangeResolutionOneSpaceWeek 1 Week
        TileTimeRangeResolutionOneSpaceWeek TileTimeRangeResolution = "P1W"
            // TileTimeRangeResolutionOneTwoSpaceHours 12 Hours
        TileTimeRangeResolutionOneTwoSpaceHours TileTimeRangeResolution = "PT12H"
            // TileTimeRangeResolutionOneZeroSpaceMinutes 10 Minutes
        TileTimeRangeResolutionOneZeroSpaceMinutes TileTimeRangeResolution = "PT10M"
            // TileTimeRangeResolutionThreeZeroSpaceMinutes 30 Minutes
        TileTimeRangeResolutionThreeZeroSpaceMinutes TileTimeRangeResolution = "PT30M"
            )
    // PossibleTileTimeRangeResolutionValues returns an array of possible values for the TileTimeRangeResolution const type.
    func PossibleTileTimeRangeResolutionValues() []TileTimeRangeResolution {
        return []TileTimeRangeResolution{TileTimeRangeResolutionFiveSpaceMinutes,TileTimeRangeResolutionOneSpaceDay,TileTimeRangeResolutionOneSpaceHour,TileTimeRangeResolutionOneSpaceMinute,TileTimeRangeResolutionOneSpaceWeek,TileTimeRangeResolutionOneTwoSpaceHours,TileTimeRangeResolutionOneZeroSpaceMinutes,TileTimeRangeResolutionThreeZeroSpaceMinutes}
    }

        // Type enumerates the values for type.
    type Type string

    const (
                // TypeAttestation ...
        TypeAttestation Type = "Attestation"
                // TypeSymmetricKey ...
        TypeSymmetricKey Type = "symmetricKey"
                // TypeTpm ...
        TypeTpm Type = "tpm"
                // TypeX509 ...
        TypeX509 Type = "x509"
            )
    // PossibleTypeValues returns an array of possible values for the Type const type.
    func PossibleTypeValues() []Type {
        return []Type{TypeAttestation,TypeSymmetricKey,TypeTpm,TypeX509}
    }

        // TypeBasicBlobStorageV1DestinationAuth enumerates the values for type basic blob storage v1 destination auth.
    type TypeBasicBlobStorageV1DestinationAuth string

    const (
                // TypeBasicBlobStorageV1DestinationAuthTypeBlobStorageV1DestinationAuth ...
        TypeBasicBlobStorageV1DestinationAuthTypeBlobStorageV1DestinationAuth TypeBasicBlobStorageV1DestinationAuth = "BlobStorageV1DestinationAuth"
                // TypeBasicBlobStorageV1DestinationAuthTypeConnectionString ...
        TypeBasicBlobStorageV1DestinationAuthTypeConnectionString TypeBasicBlobStorageV1DestinationAuth = "connectionString"
                // TypeBasicBlobStorageV1DestinationAuthTypeSystemAssignedManagedIdentity ...
        TypeBasicBlobStorageV1DestinationAuthTypeSystemAssignedManagedIdentity TypeBasicBlobStorageV1DestinationAuth = "systemAssignedManagedIdentity"
            )
    // PossibleTypeBasicBlobStorageV1DestinationAuthValues returns an array of possible values for the TypeBasicBlobStorageV1DestinationAuth const type.
    func PossibleTypeBasicBlobStorageV1DestinationAuthValues() []TypeBasicBlobStorageV1DestinationAuth {
        return []TypeBasicBlobStorageV1DestinationAuth{TypeBasicBlobStorageV1DestinationAuthTypeBlobStorageV1DestinationAuth,TypeBasicBlobStorageV1DestinationAuthTypeConnectionString,TypeBasicBlobStorageV1DestinationAuthTypeSystemAssignedManagedIdentity}
    }

        // TypeBasicDataExplorerV1DestinationAuth enumerates the values for type basic data explorer v1 destination
        // auth.
    type TypeBasicDataExplorerV1DestinationAuth string

    const (
                // TypeBasicDataExplorerV1DestinationAuthTypeDataExplorerV1DestinationAuth ...
        TypeBasicDataExplorerV1DestinationAuthTypeDataExplorerV1DestinationAuth TypeBasicDataExplorerV1DestinationAuth = "DataExplorerV1DestinationAuth"
                // TypeBasicDataExplorerV1DestinationAuthTypeServicePrincipal ...
        TypeBasicDataExplorerV1DestinationAuthTypeServicePrincipal TypeBasicDataExplorerV1DestinationAuth = "servicePrincipal"
                // TypeBasicDataExplorerV1DestinationAuthTypeSystemAssignedManagedIdentity ...
        TypeBasicDataExplorerV1DestinationAuthTypeSystemAssignedManagedIdentity TypeBasicDataExplorerV1DestinationAuth = "systemAssignedManagedIdentity"
            )
    // PossibleTypeBasicDataExplorerV1DestinationAuthValues returns an array of possible values for the TypeBasicDataExplorerV1DestinationAuth const type.
    func PossibleTypeBasicDataExplorerV1DestinationAuthValues() []TypeBasicDataExplorerV1DestinationAuth {
        return []TypeBasicDataExplorerV1DestinationAuth{TypeBasicDataExplorerV1DestinationAuthTypeDataExplorerV1DestinationAuth,TypeBasicDataExplorerV1DestinationAuthTypeServicePrincipal,TypeBasicDataExplorerV1DestinationAuthTypeSystemAssignedManagedIdentity}
    }

        // TypeBasicDestination enumerates the values for type basic destination.
    type TypeBasicDestination string

    const (
                // TypeBasicDestinationTypeBlobstorageV1 ...
        TypeBasicDestinationTypeBlobstorageV1 TypeBasicDestination = "blobstorage@v1"
                // TypeBasicDestinationTypeDataexplorerV1 ...
        TypeBasicDestinationTypeDataexplorerV1 TypeBasicDestination = "dataexplorer@v1"
                // TypeBasicDestinationTypeDestination ...
        TypeBasicDestinationTypeDestination TypeBasicDestination = "Destination"
                // TypeBasicDestinationTypeEventhubsV1 ...
        TypeBasicDestinationTypeEventhubsV1 TypeBasicDestination = "eventhubs@v1"
                // TypeBasicDestinationTypeExportDestination ...
        TypeBasicDestinationTypeExportDestination TypeBasicDestination = "ExportDestination"
                // TypeBasicDestinationTypeServicebusqueueV1 ...
        TypeBasicDestinationTypeServicebusqueueV1 TypeBasicDestination = "servicebusqueue@v1"
                // TypeBasicDestinationTypeServicebustopicV1 ...
        TypeBasicDestinationTypeServicebustopicV1 TypeBasicDestination = "servicebustopic@v1"
                // TypeBasicDestinationTypeWebhookV1 ...
        TypeBasicDestinationTypeWebhookV1 TypeBasicDestination = "webhook@v1"
            )
    // PossibleTypeBasicDestinationValues returns an array of possible values for the TypeBasicDestination const type.
    func PossibleTypeBasicDestinationValues() []TypeBasicDestination {
        return []TypeBasicDestination{TypeBasicDestinationTypeBlobstorageV1,TypeBasicDestinationTypeDataexplorerV1,TypeBasicDestinationTypeDestination,TypeBasicDestinationTypeEventhubsV1,TypeBasicDestinationTypeExportDestination,TypeBasicDestinationTypeServicebusqueueV1,TypeBasicDestinationTypeServicebustopicV1,TypeBasicDestinationTypeWebhookV1}
    }

        // TypeBasicEventHubsV1DestinationAuth enumerates the values for type basic event hubs v1 destination auth.
    type TypeBasicEventHubsV1DestinationAuth string

    const (
                // TypeBasicEventHubsV1DestinationAuthTypeConnectionString ...
        TypeBasicEventHubsV1DestinationAuthTypeConnectionString TypeBasicEventHubsV1DestinationAuth = "connectionString"
                // TypeBasicEventHubsV1DestinationAuthTypeEventHubsV1DestinationAuth ...
        TypeBasicEventHubsV1DestinationAuthTypeEventHubsV1DestinationAuth TypeBasicEventHubsV1DestinationAuth = "EventHubsV1DestinationAuth"
                // TypeBasicEventHubsV1DestinationAuthTypeSystemAssignedManagedIdentity ...
        TypeBasicEventHubsV1DestinationAuthTypeSystemAssignedManagedIdentity TypeBasicEventHubsV1DestinationAuth = "systemAssignedManagedIdentity"
            )
    // PossibleTypeBasicEventHubsV1DestinationAuthValues returns an array of possible values for the TypeBasicEventHubsV1DestinationAuth const type.
    func PossibleTypeBasicEventHubsV1DestinationAuthValues() []TypeBasicEventHubsV1DestinationAuth {
        return []TypeBasicEventHubsV1DestinationAuth{TypeBasicEventHubsV1DestinationAuthTypeConnectionString,TypeBasicEventHubsV1DestinationAuthTypeEventHubsV1DestinationAuth,TypeBasicEventHubsV1DestinationAuthTypeSystemAssignedManagedIdentity}
    }

        // TypeBasicGroupAttestation enumerates the values for type basic group attestation.
    type TypeBasicGroupAttestation string

    const (
                // TypeBasicGroupAttestationTypeGroupAttestation ...
        TypeBasicGroupAttestationTypeGroupAttestation TypeBasicGroupAttestation = "GroupAttestation"
                // TypeBasicGroupAttestationTypeSymmetricKey ...
        TypeBasicGroupAttestationTypeSymmetricKey TypeBasicGroupAttestation = "symmetricKey"
                // TypeBasicGroupAttestationTypeX509 ...
        TypeBasicGroupAttestationTypeX509 TypeBasicGroupAttestation = "x509"
            )
    // PossibleTypeBasicGroupAttestationValues returns an array of possible values for the TypeBasicGroupAttestation const type.
    func PossibleTypeBasicGroupAttestationValues() []TypeBasicGroupAttestation {
        return []TypeBasicGroupAttestation{TypeBasicGroupAttestationTypeGroupAttestation,TypeBasicGroupAttestationTypeSymmetricKey,TypeBasicGroupAttestationTypeX509}
    }

        // TypeBasicJobData enumerates the values for type basic job data.
    type TypeBasicJobData string

    const (
                // TypeBasicJobDataTypeDeviceManifestMigration ...
        TypeBasicJobDataTypeDeviceManifestMigration TypeBasicJobData = "deviceManifestMigration"
                // TypeBasicJobDataTypeDeviceTemplateMigration ...
        TypeBasicJobDataTypeDeviceTemplateMigration TypeBasicJobData = "deviceTemplateMigration"
                // TypeBasicJobDataTypeJobData ...
        TypeBasicJobDataTypeJobData TypeBasicJobData = "JobData"
            )
    // PossibleTypeBasicJobDataValues returns an array of possible values for the TypeBasicJobData const type.
    func PossibleTypeBasicJobDataValues() []TypeBasicJobData {
        return []TypeBasicJobData{TypeBasicJobDataTypeDeviceManifestMigration,TypeBasicJobDataTypeDeviceTemplateMigration,TypeBasicJobDataTypeJobData}
    }

        // TypeBasicJobScheduleEnd enumerates the values for type basic job schedule end.
    type TypeBasicJobScheduleEnd string

    const (
                // TypeBasicJobScheduleEndTypeDate ...
        TypeBasicJobScheduleEndTypeDate TypeBasicJobScheduleEnd = "date"
                // TypeBasicJobScheduleEndTypeJobScheduleEnd ...
        TypeBasicJobScheduleEndTypeJobScheduleEnd TypeBasicJobScheduleEnd = "JobScheduleEnd"
                // TypeBasicJobScheduleEndTypeOccurrences ...
        TypeBasicJobScheduleEndTypeOccurrences TypeBasicJobScheduleEnd = "occurrences"
            )
    // PossibleTypeBasicJobScheduleEndValues returns an array of possible values for the TypeBasicJobScheduleEnd const type.
    func PossibleTypeBasicJobScheduleEndValues() []TypeBasicJobScheduleEnd {
        return []TypeBasicJobScheduleEnd{TypeBasicJobScheduleEndTypeDate,TypeBasicJobScheduleEndTypeJobScheduleEnd,TypeBasicJobScheduleEndTypeOccurrences}
    }

        // TypeBasicQueryRangeConfiguration enumerates the values for type basic query range configuration.
    type TypeBasicQueryRangeConfiguration string

    const (
                // TypeBasicQueryRangeConfigurationTypeCount ...
        TypeBasicQueryRangeConfigurationTypeCount TypeBasicQueryRangeConfiguration = "count"
                // TypeBasicQueryRangeConfigurationTypeQueryRangeConfiguration ...
        TypeBasicQueryRangeConfigurationTypeQueryRangeConfiguration TypeBasicQueryRangeConfiguration = "QueryRangeConfiguration"
                // TypeBasicQueryRangeConfigurationTypeTime ...
        TypeBasicQueryRangeConfigurationTypeTime TypeBasicQueryRangeConfiguration = "time"
            )
    // PossibleTypeBasicQueryRangeConfigurationValues returns an array of possible values for the TypeBasicQueryRangeConfiguration const type.
    func PossibleTypeBasicQueryRangeConfigurationValues() []TypeBasicQueryRangeConfiguration {
        return []TypeBasicQueryRangeConfiguration{TypeBasicQueryRangeConfigurationTypeCount,TypeBasicQueryRangeConfigurationTypeQueryRangeConfiguration,TypeBasicQueryRangeConfigurationTypeTime}
    }

        // TypeBasicServiceBusQueueV1DestinationAuth enumerates the values for type basic service bus queue v1
        // destination auth.
    type TypeBasicServiceBusQueueV1DestinationAuth string

    const (
                // TypeBasicServiceBusQueueV1DestinationAuthTypeConnectionString ...
        TypeBasicServiceBusQueueV1DestinationAuthTypeConnectionString TypeBasicServiceBusQueueV1DestinationAuth = "connectionString"
                // TypeBasicServiceBusQueueV1DestinationAuthTypeServiceBusQueueV1DestinationAuth ...
        TypeBasicServiceBusQueueV1DestinationAuthTypeServiceBusQueueV1DestinationAuth TypeBasicServiceBusQueueV1DestinationAuth = "ServiceBusQueueV1DestinationAuth"
                // TypeBasicServiceBusQueueV1DestinationAuthTypeSystemAssignedManagedIdentity ...
        TypeBasicServiceBusQueueV1DestinationAuthTypeSystemAssignedManagedIdentity TypeBasicServiceBusQueueV1DestinationAuth = "systemAssignedManagedIdentity"
            )
    // PossibleTypeBasicServiceBusQueueV1DestinationAuthValues returns an array of possible values for the TypeBasicServiceBusQueueV1DestinationAuth const type.
    func PossibleTypeBasicServiceBusQueueV1DestinationAuthValues() []TypeBasicServiceBusQueueV1DestinationAuth {
        return []TypeBasicServiceBusQueueV1DestinationAuth{TypeBasicServiceBusQueueV1DestinationAuthTypeConnectionString,TypeBasicServiceBusQueueV1DestinationAuthTypeServiceBusQueueV1DestinationAuth,TypeBasicServiceBusQueueV1DestinationAuthTypeSystemAssignedManagedIdentity}
    }

        // TypeBasicServiceBusTopicV1DestinationAuth enumerates the values for type basic service bus topic v1
        // destination auth.
    type TypeBasicServiceBusTopicV1DestinationAuth string

    const (
                // TypeBasicServiceBusTopicV1DestinationAuthTypeConnectionString ...
        TypeBasicServiceBusTopicV1DestinationAuthTypeConnectionString TypeBasicServiceBusTopicV1DestinationAuth = "connectionString"
                // TypeBasicServiceBusTopicV1DestinationAuthTypeServiceBusTopicV1DestinationAuth ...
        TypeBasicServiceBusTopicV1DestinationAuthTypeServiceBusTopicV1DestinationAuth TypeBasicServiceBusTopicV1DestinationAuth = "ServiceBusTopicV1DestinationAuth"
                // TypeBasicServiceBusTopicV1DestinationAuthTypeSystemAssignedManagedIdentity ...
        TypeBasicServiceBusTopicV1DestinationAuthTypeSystemAssignedManagedIdentity TypeBasicServiceBusTopicV1DestinationAuth = "systemAssignedManagedIdentity"
            )
    // PossibleTypeBasicServiceBusTopicV1DestinationAuthValues returns an array of possible values for the TypeBasicServiceBusTopicV1DestinationAuth const type.
    func PossibleTypeBasicServiceBusTopicV1DestinationAuthValues() []TypeBasicServiceBusTopicV1DestinationAuth {
        return []TypeBasicServiceBusTopicV1DestinationAuth{TypeBasicServiceBusTopicV1DestinationAuthTypeConnectionString,TypeBasicServiceBusTopicV1DestinationAuthTypeServiceBusTopicV1DestinationAuth,TypeBasicServiceBusTopicV1DestinationAuthTypeSystemAssignedManagedIdentity}
    }

        // TypeBasicTileConfiguration enumerates the values for type basic tile configuration.
    type TypeBasicTileConfiguration string

    const (
                // TypeBasicTileConfigurationTypeCommand ...
        TypeBasicTileConfigurationTypeCommand TypeBasicTileConfiguration = "command"
                // TypeBasicTileConfigurationTypeDataExplorer ...
        TypeBasicTileConfigurationTypeDataExplorer TypeBasicTileConfiguration = "dataExplorer"
                // TypeBasicTileConfigurationTypeDeviceCount ...
        TypeBasicTileConfigurationTypeDeviceCount TypeBasicTileConfiguration = "deviceCount"
                // TypeBasicTileConfigurationTypeExternalContent ...
        TypeBasicTileConfigurationTypeExternalContent TypeBasicTileConfiguration = "externalContent"
                // TypeBasicTileConfigurationTypeImage ...
        TypeBasicTileConfigurationTypeImage TypeBasicTileConfiguration = "image"
                // TypeBasicTileConfigurationTypeLabel ...
        TypeBasicTileConfigurationTypeLabel TypeBasicTileConfiguration = "label"
                // TypeBasicTileConfigurationTypeMarkdown ...
        TypeBasicTileConfigurationTypeMarkdown TypeBasicTileConfiguration = "markdown"
                // TypeBasicTileConfigurationTypeTileConfiguration ...
        TypeBasicTileConfigurationTypeTileConfiguration TypeBasicTileConfiguration = "TileConfiguration"
            )
    // PossibleTypeBasicTileConfigurationValues returns an array of possible values for the TypeBasicTileConfiguration const type.
    func PossibleTypeBasicTileConfigurationValues() []TypeBasicTileConfiguration {
        return []TypeBasicTileConfiguration{TypeBasicTileConfigurationTypeCommand,TypeBasicTileConfigurationTypeDataExplorer,TypeBasicTileConfigurationTypeDeviceCount,TypeBasicTileConfigurationTypeExternalContent,TypeBasicTileConfigurationTypeImage,TypeBasicTileConfigurationTypeLabel,TypeBasicTileConfigurationTypeMarkdown,TypeBasicTileConfigurationTypeTileConfiguration}
    }

        // TypeBasicUser enumerates the values for type basic user.
    type TypeBasicUser string

    const (
                // TypeBasicUserTypeAdGroup ...
        TypeBasicUserTypeAdGroup TypeBasicUser = "adGroup"
                // TypeBasicUserTypeEmail ...
        TypeBasicUserTypeEmail TypeBasicUser = "email"
                // TypeBasicUserTypeServicePrincipal ...
        TypeBasicUserTypeServicePrincipal TypeBasicUser = "servicePrincipal"
                // TypeBasicUserTypeUser ...
        TypeBasicUserTypeUser TypeBasicUser = "User"
            )
    // PossibleTypeBasicUserValues returns an array of possible values for the TypeBasicUser const type.
    func PossibleTypeBasicUserValues() []TypeBasicUser {
        return []TypeBasicUser{TypeBasicUserTypeAdGroup,TypeBasicUserTypeEmail,TypeBasicUserTypeServicePrincipal,TypeBasicUserTypeUser}
    }

        // TypeBasicWebhookV1DestinationAuth enumerates the values for type basic webhook v1 destination auth.
    type TypeBasicWebhookV1DestinationAuth string

    const (
                // TypeBasicWebhookV1DestinationAuthTypeHeader ...
        TypeBasicWebhookV1DestinationAuthTypeHeader TypeBasicWebhookV1DestinationAuth = "header"
                // TypeBasicWebhookV1DestinationAuthTypeOauth ...
        TypeBasicWebhookV1DestinationAuthTypeOauth TypeBasicWebhookV1DestinationAuth = "oauth"
                // TypeBasicWebhookV1DestinationAuthTypeWebhookV1DestinationAuth ...
        TypeBasicWebhookV1DestinationAuthTypeWebhookV1DestinationAuth TypeBasicWebhookV1DestinationAuth = "WebhookV1DestinationAuth"
            )
    // PossibleTypeBasicWebhookV1DestinationAuthValues returns an array of possible values for the TypeBasicWebhookV1DestinationAuth const type.
    func PossibleTypeBasicWebhookV1DestinationAuthValues() []TypeBasicWebhookV1DestinationAuth {
        return []TypeBasicWebhookV1DestinationAuth{TypeBasicWebhookV1DestinationAuthTypeHeader,TypeBasicWebhookV1DestinationAuthTypeOauth,TypeBasicWebhookV1DestinationAuthTypeWebhookV1DestinationAuth}
    }

