//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/semi-technologies/weaviate/test/docker"
	"github.com/semi-technologies/weaviate/test/helper"
	"github.com/semi-technologies/weaviate/test/helper/journey"
	"github.com/stretchr/testify/require"
)

const (
	envGCSEndpoint            = "GCS_ENDPOINT"
	envGCSStorageEmulatorHost = "STORAGE_EMULATOR_HOST"
	envGCSCredentials         = "GOOGLE_APPLICATION_CREDENTIALS"
	envGCSProjectID           = "GOOGLE_CLOUD_PROJECT"
	envGCSBucket              = "BACKUP_GCS_BUCKET"

	gcsBackupJourneyClassName  = "GcsBackup"
	gcsBackupJourneyBackupID   = "gcs-backup"
	gcsBackupJourneyProjectID  = "gcs-backup-journey"
	gcsBackupJourneyBucketName = "backups"
)

func Test_BackupJourney(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	t.Run("pre-instance env setup", func(t *testing.T) {
		require.Nil(t, os.Setenv(envGCSCredentials, ""))
		require.Nil(t, os.Setenv(envGCSProjectID, gcsBackupJourneyProjectID))
		require.Nil(t, os.Setenv(envGCSBucket, gcsBackupJourneyBucketName))
	})

	compose, err := docker.New().
		WithBackendGCS(gcsBackupJourneyBucketName).
		WithText2VecContextionary().
		WithWeaviate().
		Start(ctx)
	require.Nil(t, err)
	defer func() {
		if err := compose.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminte test containers: %s", err.Error())
		}
	}()

	t.Run("post-instance env setup", func(t *testing.T) {
		require.Nil(t, os.Setenv(envGCSEndpoint, compose.GetGCS().URI()))
		require.Nil(t, os.Setenv(envGCSStorageEmulatorHost, compose.GetGCS().URI()))

		createBucket(ctx, t, gcsBackupJourneyProjectID, gcsBackupJourneyBucketName)
		helper.SetupClient(compose.GetWeaviate().URI())
	})

	// journey tests
	t.Run("backup-gcs", func(t *testing.T) {
		journey.BackupJourneyTests(t, compose.GetWeaviate().URI(),
			"gcs", gcsBackupJourneyClassName, gcsBackupJourneyBackupID)
	})
}
