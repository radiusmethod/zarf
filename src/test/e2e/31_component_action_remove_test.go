// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2021-Present The Zarf Authors

// Package test provides e2e tests for Zarf.
package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComponentActionRemove(t *testing.T) {
	t.Log("E2E: Component action remove")
	e2e.setupWithCluster(t)
	defer e2e.teardown(t)

	path := fmt.Sprintf("build/zarf-package-component-actions-%s.tar.zst", e2e.arch)

	stdOut, stdErr, err := e2e.execZarfCommand("package", "deploy", path, "--confirm", "--components=on-remove")
	require.NoError(t, err, stdOut, stdErr)

	stdOut, stdErr, err = e2e.execZarfCommand("package", "remove", path, "--confirm", "--components=on-remove")
	require.NoError(t, err, stdOut, stdErr)
	require.Contains(t, stdErr, "NAME")
	require.Contains(t, stdErr, "DATA")
	require.Contains(t, stdErr, "remove-test-configmap")
	require.Contains(t, stdErr, "Not Found")
}
