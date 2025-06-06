/*
さくらのVPS APIドキュメント

Testing ZoneAPIService

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

func Test_sakura_vps_ZoneAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ZoneAPIService GetZoneList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ZoneAPI.GetZoneList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
