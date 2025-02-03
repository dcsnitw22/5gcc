package sm_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog"
	openapiserver "w5gc.io/wipro5gcore/openapi/openapiserver"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/apiclient"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/config"
	db "w5gc.io/wipro5gcore/pkg/smf/pdusmsp/database"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/grpc"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/sm"
)

var defaultPdusmspConfig = []byte(`
{
        "Version": "1.0",
        "NodeInfo":     {
                        "NodeId": "127.0.0.1",
			"ApiPort": ":8080"
                },
        "N11AmfNodes":[
                {
                        "NodeId": "127.0.0.1",
						"Port": "8083"
                },
				{
					"NodeId": "10.250.108.35",
					"Port": "8080"
				},
				{
					"NodeId": "10.250.110.37",
					"Port": "8080"
				}	
        ],
		"GrpcServerInfo":{
			"ServerIP": "127.0.0.1",
			"ServerPort": "50051"
		},
		"GrpcClientInfo":{
			"ClientIP": "127.0.0.1",
			"ClientPort": "50052"
		}
}`)

var _ = Describe("Test for SmContext", Ordered,
	func() {
		var dbClient *db.DBInfo
		var grpcClient grpc.Grpc
		var apiClient apiclient.ApiClient
		var sessionClient *sm.SmInfo
		BeforeAll(func() {

			dbClient = db.NewDBManager()
			dbClient.Start()

			var cfg config.PdusmspConfig
			json.Unmarshal(defaultPdusmspConfig, &cfg)

			grpcClient = grpc.NewGrpc(cfg.GrpcServerInfo, cfg.GrpcClientInfo)
			go grpcClient.Start()

			apiClient = apiclient.NewApiClient(&cfg)
			go apiClient.Start()

			sessionClient = sm.NewSessionManager(dbClient, &grpcClient, apiClient)
			go sessionClient.Start()

		})

		// Context("Test SmContext",
		// 	func()  {
		// 		It()
		// })

		var _ = DescribeTable("Create SmContext",
			func(jsonData openapiserver.SmContextCreateData,
				binaryDataN1SmMessage *os.File,
				expectedResponse openapiserver.ImplResponse,
				expectedError error) {

				// var dbClient *db.DBInfo
				// var grpcClient grpc.Grpc
				// var apiClient apiclient.ApiClient
				// var sessionClient *sm.SmInfo
				// BeforeAll(func() {

				// 	dbClient = db.NewDBManager()
				// 	dbClient.Start()

				// 	var cfg config.PdusmspConfig
				// 	json.Unmarshal(defaultPdusmspConfig, &cfg)

				// 	grpcClient = grpc.NewGrpc(cfg.GrpcServerInfo, cfg.GrpcClientInfo)
				// 	grpcClient.Start()

				// 	apiClient = apiclient.NewApiClient(&cfg)
				// 	apiClient.Start()

				// 	sessionClient = sm.NewSessionManager(dbClient, &grpcClient, apiClient)
				// 	sessionClient.Start()

				// })

				// grpc.Start()

				// a := sm.NewSessionManager(dbClient, &grpcClient, apiClient)
				resp, err := sessionClient.ProcessNsmfCreateSmContextRequest(jsonData, binaryDataN1SmMessage)
				if err == nil {
					Expect(err).To(BeNil())
				} else {
					Expect(err).Should(Equal(expectedError))
				}
				Expect(resp).Should(Equal(expectedResponse))
			},
			Entry(
				"Response 201 OK: Create SmContext",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_201_0.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
					return file
				}(),
				openapiserver.Response(201,
					openapiserver.SmContextCreatedData{
						UpCnxState:   openapiserver.ACTIVATING,
						PduSessionId: 51,
					}),
				nil,
			),

			Entry(
				"Response 403 Error: Invalid PduSessionId",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_0.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_0.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: "invalid PduSessionId ",
							},
							N1SmMsg: openapiserver.RefToBinaryData{
								ContentId: "n1msgtest",
							},
						})
				}(),
				errors.New("invalid PduSessionId"),
			),
			Entry(
				"Response 403 Error: Invalid Supi",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_1.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_1.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid supi").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid supi"),
			),
			Entry(
				"Response 403 Error: Invalid Pei",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_2.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_2.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid pei").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid pei"),
			),
			Entry(
				"Response 403 Error: Invalid guami: more than 6 characters in AmfId",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_3.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_3.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid guami").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid guami"),
			),
			Entry(
				"Response 403 Error: Invalid guami: having Alphabet outside A-F range",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_4.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_4.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid guami").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid guami"),
			),
			Entry(
				"Response 403 Error: Invalid guami:Invalid MCC: having 4 character which is a not a number",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_7.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_7.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid guami").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid guami"),
			),
			Entry(
				"Response 403 Error: Invalid guami :Invalid mcc having more than 3 character that is not a number",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_8.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_8.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid guami").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid guami"),
			),
			Entry(
				"Response 403 Error: Invalid serving network:Invalid mcc: having more than 3 characters",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_5.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_5.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					klog.Error(jsonData.PduSessionId)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid servingNetwork").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid servingNetwork"),
			),
			Entry(
				"Response 403 Error:Invalid serving network :Invalid Mnc:having more than 3 characters",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_6.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/smContextCreate_403_6.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(
						403,
						openapiserver.SmContextCreateError{
							Error: openapiserver.ProblemDetails{
								Title:  "invalid request data",
								Type:   "validityErr",
								Status: 403,
								Detail: errors.New("invalid servingNetwork").Error(),
							},
							N1SmMsg: jsonData.N1SmMsg})
				}(),
				errors.New("invalid servingNetwork"),
			),
			Entry(
				"Response 403 Error: request already in progress",
				func() openapiserver.SmContextCreateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/SmContextCreate_403_9.json")
					if err != nil {
						klog.Error("File not found")
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				func() *os.File {
					file, err := os.Open("/home/ubuntu/wipro5gc/TestCases/PostmanTestCases/n1msgtest")
					if err != nil {
						klog.Error(err.Error())
					}
					return file
				}(),
				func() openapiserver.ImplResponse {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/CreateSmContextRequest/SmContextCreate_403_9.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextCreateData
					json.Unmarshal(file, &jsonData)
					return openapiserver.Response(403, openapiserver.SmContextCreateError{
						Error: openapiserver.ProblemDetails{
							Title:  "Already in Progress",
							Type:   "ALreadyInProgressErr",
							Status: 403,
							Detail: "request already in progress",
						},
						N1SmMsg: jsonData.N1SmMsg,
					})
				}(),
				errors.New("request already in progress"),
			),
		)

		var _ = DescribeTable("Update SmContext",
			func(smcontextRef string,
				jsonData openapiserver.SmContextUpdateData,
				expectedResponse openapiserver.ImplResponse,
				expectedError error) {
				// client := db.NewDBManager()
				// var cfg config.PdusmspConfig
				// json.Unmarshal(defaultPdusmspConfig, &cfg)
				// grpc := grpc.NewGrpc(cfg.GrpcServerInfo, cfg.GrpcClientInfo)
				// apiclient := apiclient.NewApiClient(&cfg)
				// a := sm.NewSessionManager(client, &grpc, apiclient)
				resp, err := sessionClient.ProcessNsmfUpdateSmContextRequest(smcontextRef, jsonData)
				if err == nil {
					Expect(err).To(BeNil())
				} else {
					Expect(err).Should(Equal(expectedError))
				}
				Expect(resp).Should(Equal(expectedResponse))
			},
			Entry(
				"Response 200 OK",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_200.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusOK, openapiserver.SmContextUpdatedData{
					// UpCnxState: openapiserver.ACTIVATED,
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), nil,
			),
			Entry(
				"Response 400 SmContextRef Empty",
				"",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_200.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Context Reference is empty",
						Type:   "ValidityErr",
						Detail: "smContextRef must not be empty",
						Status: 400,
						Cause:  "MANDATORY_QUERY_PARMS_MISSING",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("smContextRef must not be empty"),
			),
			Entry(
				"Response 400 SmContextRef not valid",
				"invalid",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_200.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusNotFound, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Session Context Not found",
						Type:   "NotFoundErr",
						Detail: "Session context corresponding to smContextRef not found",
						Status: 404,
						Cause:  "CONTEXT_NOT_FOUND",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("session context corresponding to smContextRef not found"),
			),
			Entry(
				"Response 400 Validation- invalid pei",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_400_1.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Invalid data sent",
						Type:   "ValidityErr",
						Detail: "invalid pei",
						Status: 400,
						Cause:  "",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("invalid pei"),
			),
			Entry(
				"Response 400 Validation - invalid guami",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_400_2.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Invalid data sent",
						Type:   "ValidityErr",
						Detail: "invalid guami",
						Status: 400,
						Cause:  "",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("invalid guami"),
			),
			Entry(
				"Response 400 Validation - invalid serving network",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_400_3.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Invalid data sent",
						Type:   "ValidityErr",
						Detail: "invalid servingNetwork",
						Status: 400,
						Cause:  "",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("invalid servingNetwork"),
			),
			Entry(
				"Response 400 Validation - invalid servingnetwork1",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_400_4.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Invalid data sent",
						Type:   "ValidityErr",
						Detail: "invalid servingNetwork",
						Status: 400,
						Cause:  "",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("invalid servingNetwork"),
			),
			Entry(
				"Response 400 Validation- invalid guami1",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_400_5.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Invalid data sent",
						Type:   "ValidityErr",
						Detail: "invalid guami",
						Status: 400,
						Cause:  "",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("invalid guami"),
			),
			Entry(
				"Response 400 Validation- invalid guami2",
				"51218A9E",
				func() openapiserver.SmContextUpdateData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/UpdateSmContextRequest/smContextUpdate_400_6.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextUpdateData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, openapiserver.SmContextUpdateError{
					Error: openapiserver.ProblemDetails{
						Title:  "Invalid data sent",
						Type:   "ValidityErr",
						Detail: "invalid guami",
						Status: 400,
						Cause:  "",
					},
					N1SmMsg: openapiserver.RefToBinaryData{
						ContentId: "n1smmsg27183",
					},
					N2SmInfo: openapiserver.RefToBinaryData{
						ContentId: "n2sminfo10296",
					},
				}), errors.New("invalid guami"),
			),
		)

		var _ = DescribeTable("Retrieve SmContext",
			func(smcontextRef string,
				jsonData openapiserver.SmContextRetrieveData,
				expectedResponse openapiserver.ImplResponse,
				expectedError error) {
				client := db.NewDBManager()
				var cfg config.PdusmspConfig
				json.Unmarshal(defaultPdusmspConfig, &cfg)
				grpc := grpc.NewGrpc(cfg.GrpcServerInfo, cfg.GrpcClientInfo)
				apiclient := apiclient.NewApiClient(&cfg)
				a := sm.NewSessionManager(client, &grpc, apiclient)
				resp, err := a.ProcessNsmfRetrieveSmContextRequest(smcontextRef, jsonData)
				if err == nil {
					Expect(err).To(BeNil())
				} else {
					Expect(err).Should(Equal(expectedError))
				}
				Expect(resp).Should(Equal(expectedResponse))
			},
			Entry(
				"Response 200 OK Retrieve",
				"51218A9E",
				func() openapiserver.SmContextRetrieveData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/RetrieveSmContextRequest/smContextRetrieveData.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextRetrieveData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusOK, openapiserver.SmContextRetrievedData{
					UeEpsPdnConnection: "",
				}), nil,
			),
			Entry(
				"Response 400 bad request smcontext is empty",
				"",
				func() openapiserver.SmContextRetrieveData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/RetrieveSmContextRequest/smContextRetrieveData.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextRetrieveData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, nil), errors.New("smContextRef must not be empty"),
			),
			Entry(
				"Response 404 bad request smcontext not valid",
				"invalid",
				func() openapiserver.SmContextRetrieveData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/RetrieveSmContextRequest/smContextRetrieveData.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextRetrieveData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusNotFound, nil), errors.New("smContextRef not a valid input"),
			),
		)

		var _ = DescribeTable("Release SmContext",
			func(smcontextRef string,
				jsonData openapiserver.SmContextReleaseData,
				expectedResponse openapiserver.ImplResponse,
				expectedError error) {
				// client := db.NewDBManager()
				// var cfg config.PdusmspConfig
				// json.Unmarshal(defaultPdusmspConfig, &cfg)
				// grpc := grpc.NewGrpc(cfg.GrpcServerInfo, cfg.GrpcClientInfo)
				// apiclient := apiclient.NewApiClient(&cfg)
				// a := sm.NewSessionManager(client, &grpc, apiclient)
				resp, err := sessionClient.ProcessNsmfReleaseSmContextRequest(smcontextRef, jsonData)
				if err == nil {
					Expect(err).To(BeNil())
				} else {
					Expect(err).Should(Equal(expectedError))
				}
				Expect(resp).Should(Equal(expectedResponse))
			},
			Entry(
				"Response 204 No Content",
				"51218A9E",
				func() openapiserver.SmContextReleaseData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/ReleaseSmContextRequest/smContextRelease.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextReleaseData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusNoContent, nil), nil,
			),
			Entry(
				"Response 400 SmContextRef Empty",
				"",
				func() openapiserver.SmContextReleaseData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/ReleaseSmContextRequest/smContextRelease.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextReleaseData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(http.StatusBadRequest, nil), errors.New("smContextRef must not be empty"),
			),
			Entry(
				"Response 400 SmContextRef Not Valid",
				"invalid",
				func() openapiserver.SmContextReleaseData {
					file, err := os.ReadFile("/home/ubuntu/wipro5gc/TestCases/UnitTestCases/pkg/smf/pdusmsp/sm/ReleaseSmContextRequest/smContextRelease.json")
					if err != nil {
						klog.Error(err.Error())
					}
					var jsonData openapiserver.SmContextReleaseData
					json.Unmarshal(file, &jsonData)
					return jsonData
				}(),
				openapiserver.Response(
					http.StatusNotFound,
					openapiserver.ProblemDetails{},
				), errors.New("smContextRef not a valid input"),
			),
		)

	})
