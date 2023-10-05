package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/google/data/iam/google_service_account.json
var dataGoogleServiceAccount []byte

//go:embed mapping/google/data/compute/google_compute_network.json
var dataGoogleComputeNetwork []byte

//go:embed mapping/google/data/compute/google_compute_subnetwork.json
var dataGoogleComputeSubnetwork []byte

//go:embed mapping/google/data/compute/google_compute_zones.json
var dataGoogleComputeZones []byte

//go:embed mapping/google/data/resourcemanager/google_project.json
var dataGoogleProject []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_ring.json
var dataGoogleKmsKeyRing []byte

//go:embed  mapping/google/data/cloudkms/google_kms_crypto_key.json
var dataGoogleKmsCryptoKey []byte

//go:embed  mapping/google/data/dns/google_dns_keys.json
var dataGoogleDnsKeys []byte

//go:embed  mapping/google/data/dns/google_dns_managed_zone.json
var dataGoogleDnsManagedZone []byte

//go:embed  mapping/google/data/dns/google_dns_managed_zone_iam_policy.json
var dataGoogleDnsManagedZoneIamPolicy []byte

//go:embed  mapping/google/data/dns/google_dns_record_set.json
var dataGoogleDnsRecordSet []byte

//go:embed  mapping/google/data/artifactregistry/google_artifact_registry_repository.json
var dataGoogleArtifactRegistryRepository []byte

//go:embed  mapping/google/data/artifactregistry/google_artifact_registry_repository_iam_policy.json
var dataGoogleArtifactRegistryRepositoryIamPolicy []byte

//go:embed  mapping/google/data/bigquery/google_app_engine_default_service_account.json
var dataGoogleAppEngineDefaultServiceAccount []byte

//go:embed  mapping/google/data/bigquery/google_bigquery_datapolicy_data_policy_iam_policy.json
var dataGoogleBigqueryDatapolicyDataPolicyIamPolicy []byte

//go:embed  mapping/google/data/bigquery/google_app_engine_default_service_account.json
var dataGoogleBigqueryDefaultServiceAccount []byte

//go:embed  mapping/google/data/bigtable/google_bigtable_instance_iam_policy.json
var dataGoogleBigtableInstanceIamPolicy []byte

//go:embed  mapping/google/data/analyticshub/google_bigquery_analytics_hub_data_exchange_iam_policy.json
var dataGoogleBigqueryHubDataExchangeIamPolicy []byte

//go:embed  mapping/google/data/analyticshub/google_bigquery_analytics_hub_listing_iam_policy.json
var dataGoogleBigqueryAnalyticsHubListingIamPolicy []byte

//go:embed  mapping/google/data/cloudkms/google_kms_key_ring_iam_policy.json
var dataGoogleKmsKeyRingIamPolicy []byte

//go:embed  mapping/google/data/cloudkms/google_kms_secret.json
var dataGoogleKmsSecret []byte

//go:embed  mapping/google/data/cloudkms/google_kms_secret_asymmetric.json
var dataGoogleKmsSecretAsymnetric []byte

//go:embed  mapping/google/data/pubsub/google_pubsub_subscription.json
var dataGooglePubsubSubscription []byte

//go:embed  mapping/google/data/pubsub/google_pubsub_subscription_iam_policy.json
var dataGooglePubsubSubscriptionIamPolicy []byte

//go:embed  mapping/google/data/pubsub/google_pubsub_topic.json
var dataGooglePubsubTopic []byte

//go:embed  mapping/google/data/pubsub/google_pubsub_topic_iam_policy.json
var dataGooglePubsubTopicIamPolicy []byte

//go:embed  mapping/google/data/spanner/google_spanner_database_iam_policy.json
var dataGoogleSpannerDatabaseIamPolicy []byte

//go:embed  mapping/google/data/spanner/google_spanner_instance.json
var dataGoogleSpannerInstance []byte

//go:embed  mapping/google/data/spanner/google_spanner_instance_iam_policy.json
var dataGoogleSpannerInstanceIamPolicy []byte

//go:embed  mapping/google/data/storage/google_storage_bucket.json
var dataGoogleStorageBucket []byte

//go:embed  mapping/google/data/storage/google_storage_bucket_iam_policy.json
var dataGoogleStorageBucketIamPolicy []byte

//go:embed  mapping/google/data/storage/google_storage_bucket_object.json
var dataGoogleStorageBucketObject []byte

//go:embed  mapping/google/data/storage/google_storage_bucket_object_content.json
var dataGoogleStorageBucketObjectContent []byte

//go:embed  mapping/google/data/resourcemanager/google_storage_project_service_account.json
var dataGoogleStorageProjectServiceAccount []byte

//go:embed  mapping/google/data/storagetransfer/google_storage_transfer_project_service_account.json
var dataGoogleStorageTransferProjectServiceAccount []byte

//go:embed  mapping/google/data/aiplatform/google_vertex_ai_featurestore_entitytype_iam_policy.json
var dataGoogleVertexAiFeaturestoreEntitytypeIamPolicy []byte

//go:embed  mapping/google/data/aiplatform/google_vertex_ai_featurestore_iam_policy.json
var dataGoogleVertexAiFeaturestoreIamPolicy []byte
