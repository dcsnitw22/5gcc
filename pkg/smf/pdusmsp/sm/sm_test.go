package sm_test

// import (
// 	"errors"
// 	"os"

// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// 	openapiserver "w5gc.io/wipro5gcore/openapi/openapiserver"
// 	db "w5gc.io/wipro5gcore/pkg/smf/pdusmsp/database"
// 	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/sm"
// )

// var _ = DescribeTable("Create SmContext",
// 	func(jsonData openapiserver.SmContextCreateData,
// 		binaryDataN1SmMessage *os.File,
// 		expectedResponse openapiserver.ImplResponse,
// 		expectedError error) {
// 		client := db.NewDBManager()
// 		a := sm.NewSessionManager(client)
// 		resp, err := a.ProcessNsmfCreateSmContextRequest(jsonData, binaryDataN1SmMessage)
// 		if err != nil {
// 			Expect(err).To(BeNil())
// 		} else {
// 			Expect(err).Should(Equal(expectedError))
// 		}
// 		Expect(resp).Should(Equal(expectedResponse))
// 	},
// 	Entry(
// 		"Response 200 OK",
// 		openapiserver.SmContextCreateData{
// 			PduSessionId:       1233345,
// 			Supi:               "airtelgprs.com",
// 			ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 			SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 			ServingNetwork: openapiserver.PlmnId{
// 				Mcc: "404",
// 				Mnc: "10",
// 			},
// 			AnType: "3GPP_ACCESS",
// 			Pei:    "280129383791827",
// 			Guami: openapiserver.Guami{
// 				PlmnId: openapiserver.PlmnId{
// 					Mcc: "404",
// 					Mnc: "10",
// 				},
// 				AmfId: "218A9E",
// 			},
// 		},
// 		func() *os.File {
// 			file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 			return file
// 		}(),
// 		openapiserver.Response(201,
// 			openapiserver.SmContextCreatedData{
// 				UpCnxState:   openapiserver.ACTIVATING,
// 				PduSessionId: 1233345,
// 			}),
// 		nil,
// 	),

// 	Entry(
// 		"Response 403 invalid request data",
// 		openapiserver.SmContextCreateData{
// 			PduSessionId:       0, //Cause of err
// 			Supi:               "airtelgprs.com",
// 			ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 			SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 			ServingNetwork: openapiserver.PlmnId{
// 				Mcc: "404",
// 				Mnc: "10",
// 			},
// 			AnType: "3GPP_ACCESS",
// 			Pei:    "280129383791827",
// 			Guami: openapiserver.Guami{
// 				PlmnId: openapiserver.PlmnId{
// 					Mcc: "404",
// 					Mnc: "10",
// 				},
// 				AmfId: "218A9E",
// 			},
// 		},
// 		func() *os.File {
// 			file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 			return file
// 		}(),
// 		openapiserver.Response(
// 			403,
// 			openapiserver.SmContextCreateError{
// 				Error: openapiserver.ProblemDetails{
// 					Title:  "invalid request data",
// 					Type:   "validityErr",
// 					Status: 403,
// 					Detail: "invalid PduSessionId "},
// 				N1SmMsg: openapiserver.RefToBinaryData{
// 					ContentId: "",
// 				},
// 			},
// 		),
// 		errors.New("invalid PduSessionId"),
// 	),
// )

// var _ = DescribeTable("Release SmContext",
// 	func(smContextRef string,
// 		smContextReleaseData openapiserver.SmContextReleaseData,
// 		expectedResponse openapiserver.ImplResponse,
// 		expectedError error) {
// 		client := db.NewDBManager()
// 		a := sm.NewSessionManager(client)
// 		// Expect(a.ProcessNsmfCreateSmContextRequest(
// 		//  jsonData,
// 		//  binaryDataN1SmMessage)).Error().Should(Equal(expectedError))
// 		Expect(a.ProcessNsmfReleaseSmContextRequest(
// 			smContextRef,
// 			smContextReleaseData)).Should(Equal(expectedResponse))
// 	},

// 	//TODO Edit the release data below
// 	// Entry(

// 	// 	"Response 204 NO CONTENT",
// 	// 	openapiserver.SmContextCreateData{
// 	// 		PduSessionId:       1233345,
// 	// 		Supi:               "airtelgprs.com",
// 	// 		ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 	// 		SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 	// 		ServingNetwork: openapiserver.PlmnId{
// 	// 			Mcc: "404",
// 	// 			Mnc: "10",
// 	// 		},
// 	// 		AnType: "3GPP_ACCESS",
// 	// 		Pei:    "280129383791827",
// 	// 		Guami: openapiserver.Guami{
// 	// 			PlmnId: openapiserver.PlmnId{
// 	// 				Mcc: "404",
// 	// 				Mnc: "10",
// 	// 			},
// 	// 			AmfId: "218A9E",
// 	// 		},
// 	// 	},
// 	// 	func() *os.File {
// 	// 		file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 	// 		return file
// 	// 	}(),
// 	// 	openapiserver.Response(201,
// 	// 		openapiserver.SmContextCreatedData{
// 	// 			UpCnxState:   openapiserver.ACTIVATING,
// 	// 			PduSessionId: 1233345,
// 	// 		}),
// 	// 	nil,
// 	// ),

// 	// Entry(
// 	// 	"Response 403 invalid request data",
// 	// 	openapiserver.SmContextCreateData{
// 	// 		PduSessionId:       0,
// 	// 		Supi:               "airtelgprs.com",
// 	// 		ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 	// 		SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 	// 		ServingNetwork: openapiserver.PlmnId{
// 	// 			Mcc: "404",
// 	// 			Mnc: "10",
// 	// 		},
// 	// 		AnType: "3GPP_ACCESS",
// 	// 		Pei:    "280129383791827",
// 	// 		Guami: openapiserver.Guami{
// 	// 			PlmnId: openapiserver.PlmnId{
// 	// 				Mcc: "404",
// 	// 				Mnc: "10",
// 	// 			},
// 	// 			AmfId: "218A9E",
// 	// 		},
// 	// 	},
// 	// 	func() *os.File {
// 	// 		file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 	// 		return file
// 	// 	}(),
// 	// 	openapiserver.Response(
// 	// 		403,
// 	// 		openapiserver.SmContextCreateError{
// 	// 			Error: openapiserver.ProblemDetails{
// 	// 				Title:  "invalid request data",
// 	// 				Type:   "validityErr",
// 	// 				Status: 403,
// 	// 				Detail: "invalid PduSessionId "},
// 	// 			N1SmMsg: openapiserver.RefToBinaryData{
// 	// 				ContentId: "",
// 	// 			},
// 	// 		},
// 	// 	),
// 	// 	errors.New("invalid PduSessionId"),
// 	// ),

// )
// var _ = DescribeTable("Retrieve SmContext",
// 	func(smContextRef string,
// 		smContextRetrieveData openapiserver.SmContextRetrieveData,
// 		expectedResponse openapiserver.ImplResponse,
// 		expectedError error) {
// 		client := db.NewDBManager()
// 		a := sm.NewSessionManager(client)
// 		// Expect(a.ProcessNsmfCreateSmContextRequest(
// 		//  jsonData,
// 		//  binaryDataN1SmMessage)).Error().Should(Equal(expectedError))
// 		Expect(a.ProcessNsmfRetrieveSmContextRequest(
// 			smContextRef,
// 			smContextRetrieveData)).Should(Equal(expectedResponse))
// 	},

// 	//TODO Edit the reelase data below
// 	// Entry(

// 	// 	"Response 204 NO CONTENT",
// 	// 	openapiserver.SmContextCreateData{
// 	// 		PduSessionId:       1233345,
// 	// 		Supi:               "airtelgprs.com",
// 	// 		ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 	// 		SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 	// 		ServingNetwork: openapiserver.PlmnId{
// 	// 			Mcc: "404",
// 	// 			Mnc: "10",
// 	// 		},
// 	// 		AnType: "3GPP_ACCESS",
// 	// 		Pei:    "280129383791827",
// 	// 		Guami: openapiserver.Guami{
// 	// 			PlmnId: openapiserver.PlmnId{
// 	// 				Mcc: "404",
// 	// 				Mnc: "10",
// 	// 			},
// 	// 			AmfId: "218A9E",
// 	// 		},
// 	// 	},
// 	// 	func() *os.File {
// 	// 		file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 	// 		return file
// 	// 	}(),
// 	// 	openapiserver.Response(201,
// 	// 		openapiserver.SmContextCreatedData{
// 	// 			UpCnxState:   openapiserver.ACTIVATING,
// 	// 			PduSessionId: 1233345,
// 	// 		}),
// 	// 	nil,
// 	// ),

// 	// Entry(
// 	// 	"Response 403 invalid request data",
// 	// 	openapiserver.SmContextCreateData{
// 	// 		PduSessionId:       0,
// 	// 		Supi:               "airtelgprs.com",
// 	// 		ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 	// 		SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 	// 		ServingNetwork: openapiserver.PlmnId{
// 	// 			Mcc: "404",
// 	// 			Mnc: "10",
// 	// 		},
// 	// 		AnType: "3GPP_ACCESS",
// 	// 		Pei:    "280129383791827",
// 	// 		Guami: openapiserver.Guami{
// 	// 			PlmnId: openapiserver.PlmnId{
// 	// 				Mcc: "404",
// 	// 				Mnc: "10",
// 	// 			},
// 	// 			AmfId: "218A9E",
// 	// 		},
// 	// 	},
// 	// 	func() *os.File {
// 	// 		file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 	// 		return file
// 	// 	}(),
// 	// 	openapiserver.Response(
// 	// 		403,
// 	// 		openapiserver.SmContextCreateError{
// 	// 			Error: openapiserver.ProblemDetails{
// 	// 				Title:  "invalid request data",
// 	// 				Type:   "validityErr",
// 	// 				Status: 403,
// 	// 				Detail: "invalid PduSessionId "},
// 	// 			N1SmMsg: openapiserver.RefToBinaryData{
// 	// 				ContentId: "",
// 	// 			},
// 	// 		},
// 	// 	),
// 	// 	errors.New("invalid PduSessionId"),
// 	// ),
// )

// var _ = DescribeTable("Update SmContext",
// 	func(smContextRef string,
// 		smContextUpdateData openapiserver.SmContextUpdateData,
// 		expectedResponse openapiserver.ImplResponse,
// 		expectedError error) {
// 		client := db.NewDBManager()
// 		a := sm.NewSessionManager(client)
// 		// Expect(a.ProcessNsmfCreateSmContextRequest(
// 		//  jsonData,
// 		//  binaryDataN1SmMessage)).Error().Should(Equal(expectedError))
// 		Expect(a.ProcessNsmfUpdateSmContextRequest(
// 			smContextRef,
// 			smContextUpdateData)).Should(Equal(expectedResponse))
// 	},

// 	//TODO Edit the update data below
// 	// Entry(

// 	// 	"Response 204 NO CONTENT",
// 	// 	openapiserver.SmContextCreateData{
// 	// 		PduSessionId:       1233345,
// 	// 		Supi:               "airtelgprs.com",
// 	// 		ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 	// 		SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 	// 		ServingNetwork: openapiserver.PlmnId{
// 	// 			Mcc: "404",
// 	// 			Mnc: "10",
// 	// 		},
// 	// 		AnType: "3GPP_ACCESS",
// 	// 		Pei:    "280129383791827",
// 	// 		Guami: openapiserver.Guami{
// 	// 			PlmnId: openapiserver.PlmnId{
// 	// 				Mcc: "404",
// 	// 				Mnc: "10",
// 	// 			},
// 	// 			AmfId: "218A9E",
// 	// 		},
// 	// 	},
// 	// 	func() *os.File {
// 	// 		file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 	// 		return file
// 	// 	}(),
// 	// 	openapiserver.Response(201,
// 	// 		openapiserver.SmContextCreatedData{
// 	// 			UpCnxState:   openapiserver.ACTIVATING,
// 	// 			PduSessionId: 1233345,
// 	// 		}),
// 	// 	nil,
// 	// ),

// 	// Entry(
// 	// 	"Response 403 invalid request data",
// 	// 	openapiserver.SmContextCreateData{
// 	// 		PduSessionId:       0,
// 	// 		Supi:               "airtelgprs.com",
// 	// 		ServingNfId:        "3fa85f64-5717-4562-nwi3-2c963f66afa6",
// 	// 		SmContextStatusUri: "https://example.com/notify/smcontextstatus",
// 	// 		ServingNetwork: openapiserver.PlmnId{
// 	// 			Mcc: "404",
// 	// 			Mnc: "10",
// 	// 		},
// 	// 		AnType: "3GPP_ACCESS",
// 	// 		Pei:    "280129383791827",
// 	// 		Guami: openapiserver.Guami{
// 	// 			PlmnId: openapiserver.PlmnId{
// 	// 				Mcc: "404",
// 	// 				Mnc: "10",
// 	// 			},
// 	// 			AmfId: "218A9E",
// 	// 		},
// 	// 	},
// 	// 	func() *os.File {
// 	// 		file, _ := os.Open("/home/ubuntu/wipro5gc/TestData/n1msgtest")
// 	// 		return file
// 	// 	}(),
// 	// 	openapiserver.Response(
// 	// 		403,
// 	// 		openapiserver.SmContextCreateError{
// 	// 			Error: openapiserver.ProblemDetails{
// 	// 				Title:  "invalid request data",
// 	// 				Type:   "validityErr",
// 	// 				Status: 403,
// 	// 				Detail: "invalid PduSessionId "},
// 	// 			N1SmMsg: openapiserver.RefToBinaryData{
// 	// 				ContentId: "",
// 	// 			},
// 	// 		},
// 	// 	),
// 	// 	errors.New("invalid PduSessionId"),
// 	// ),

// )

// // var _ = Describe("SM", func() {
// // 	var (
// // 	// Define any test-specific variables here
// // 	)

// // 	BeforeEach(func() {
// // 		// Setup code that runs before each test
// // 	})

// // 	AfterEach(func() {
// // 		// Teardown code that runs after each test
// // 	})

// // 	It("should validate data", func() {
// // 		// Test the validateData function
// // 		jsonData := openapiserver.SmContextCreateData{
// // 			// Populate with valid or invalid data
// // 		}
// // 		err := sm.validateData(jsonData)
// // 		Expect(err).To(BeNil()) // or Expect(err).NotTo(BeNil())
// // 	})

// // 	It("should create a new session manager", func() {
// // 		// Test the NewSessionManager function
// // 		sm := sm.NewSessionManager()
// // 		Expect(sm).NotTo(BeNil())
// // 		// You can add more expectations here
// // 	})

// // 	It("should add a new session", func() {
// // 		// Test the AddSession function
// // 		sm := sm.NewSessionManager()
// // 		sessionID := "session123"
// // 		sessionData := openapiserver.SmContextCreateData{
// // 			// Populate session data here
// // 		}
// // 		err := AddSession(context.Background(), sessionID, sessionData)
// // 		Expect(err).To(BeNil())
// // 	})

// // 	It("should get an existing session", func() {
// // 		// Test the GetSession function
// // 		sm := sm.NewSessionManager()
// // 		sessionID := "session123"
// // 		sessionData := openapiserver.SmContextCreateData{
// // 			// Populate session data here
// // 		}
// // 		err := sm.AddSession(context.Background(), sessionID, sessionData)
// // 		Expect(err).To(BeNil())

// // 		retrievedSession, err := sm.GetSession(sessionID)
// // 		Expect(err).To(BeNil())
// // 		Expect(retrievedSession).NotTo(BeNil())
// // 	})

// // 	// Add more test cases for other functions in sm.go
// // })
