/*
decentro-in-collections

Testing CollectionsApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package decentro_in_collections

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // decentro_in_collections "github.com/decentro-in/decentro-in-collections-sdk-go/tree/master"
)

func Test_decentro_in_collections_CollectionsApiService(t *testing.T) {

    // configuration := decentro_in_collections.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetClientId("CLIENT_ID")
    configuration.SetClientSecret("CLIENT_SECRET")
    configuration.SetModuleSecret("MODULE_SECRET")
    configuration.SetProviderSecret("PROVIDER_SECRET")
    client := decentro_in_collections.NewAPIClient(configuration)
    */

    t.Run("Test CollectionsApiService GeneratePaymentLink", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        generatePaymentLinkRequest := *decentro_in_collections.NewGeneratePaymentLinkRequest(
            "ABCDEF12345",
            "00000000000000000",
            4,
            "Welcome to Decentro",
            0,
        )
        generatePaymentLinkRequest.SetExpiryTime(10)
        generatePaymentLinkRequest.SetCustomizedQrWithLogo(0)
        generatePaymentLinkRequest.SetGenerateUri(0)
        
        request := client.CollectionsApi.GeneratePaymentLink(
            generatePaymentLinkRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test CollectionsApiService GetTransactionStatus", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.CollectionsApi.GetTransactionStatus(
            "transactionId_example",
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test CollectionsApiService IssueCollectRequest", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        issueCollectRequestRequest := *decentro_in_collections.NewIssueCollectRequestRequest(
            "ABCDEF12345",
            "test001@abc",
            "00000000000000000",
            1,
            "Welcome to Decentro",
        )
        issueCollectRequestRequest.SetExpiryTime(30)
        
        request := client.CollectionsApi.IssueCollectRequest(
            issueCollectRequestRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test CollectionsApiService IssueUpiRefund", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        issueUpiRefundRequest := *decentro_in_collections.NewIssueUpiRefundRequest(
            "ABCDEF12345",
        )
        issueUpiRefundRequest.SetTransactionId("4CE9587AE3D143CDAC72E7D0CF14D028")
        issueUpiRefundRequest.SetBankReferenceNumber("111111111111")
        issueUpiRefundRequest.SetPurposeMessage("Welcome to Decentro")
        
        request := client.CollectionsApi.IssueUpiRefund(
            issueUpiRefundRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test CollectionsApiService ValidateUpiHandle", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        validateUpiHandleRequest := *decentro_in_collections.NewValidateUpiHandleRequest(
            "ABCDEF12345",
            "test001@abc",
        )
        validateUpiHandleRequest.SetType("BASIC")
        
        request := client.CollectionsApi.ValidateUpiHandle(
            validateUpiHandleRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}