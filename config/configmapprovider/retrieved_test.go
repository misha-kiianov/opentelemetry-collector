// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configmapprovider

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/config"
)

func TestNewRetrieved_NilGetFunc(t *testing.T) {
	_, err := NewRetrieved(nil)
	assert.Error(t, err)
}

func TestNewRetrieved_Default(t *testing.T) {
	expectedCfg := config.NewMapFromStringMap(map[string]interface{}{"test": nil})
	ret, err := NewRetrieved(func(context.Context) (*config.Map, error) { return expectedCfg, nil })
	require.NoError(t, err)
	cfg, err := ret.Get(context.Background())
	require.NoError(t, err)
	assert.Equal(t, expectedCfg, cfg)
	assert.NoError(t, ret.Close(context.Background()))
}

func TestNewRetrieved_GetError(t *testing.T) {
	expectedErr := errors.New("test")
	_, err := NewRetrieved(func(context.Context) (*config.Map, error) { return nil, expectedErr })
	assert.Equal(t, expectedErr, err)
}

func TestNewRetrieved_WithClose(t *testing.T) {
	expectedCfg := config.NewMapFromStringMap(map[string]interface{}{"test": nil})
	expectedCloseErr := errors.New("test")
	ret, err := NewRetrieved(
		func(context.Context) (*config.Map, error) { return expectedCfg, nil },
		WithClose(func(ctx context.Context) error { return expectedCloseErr }))
	require.NoError(t, err)
	cfg, err := ret.Get(context.Background())
	require.NoError(t, err)
	assert.Equal(t, expectedCfg, cfg)
	assert.Equal(t, expectedCloseErr, ret.Close(context.Background()))
}
