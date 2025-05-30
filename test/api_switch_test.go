/*
さくらのVPS APIドキュメント

Testing SwitchAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sakura_vps

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func Test_sakura_vps_SwitchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SwitchAPIService DeleteSwitch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var switchId int32

		httpRes, err := apiClient.SwitchAPI.DeleteSwitch(context.Background(), switchId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchAPIService GetSwitch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var switchId int32

		resp, httpRes, err := apiClient.SwitchAPI.GetSwitch(context.Background(), switchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchAPIService GetSwitchList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwitchAPI.GetSwitchList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchAPIService PostSwitch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SwitchAPI.PostSwitch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SwitchAPIService PutSwitch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var switchId int32

		resp, httpRes, err := apiClient.SwitchAPI.PutSwitch(context.Background(), switchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
