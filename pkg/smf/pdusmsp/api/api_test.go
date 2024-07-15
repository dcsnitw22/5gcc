package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"w5gc.io/wipro5gcore/openapi/openapiserver"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/api"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/sm"
)

var _ = Describe("PostSMContext", func() {
	//Initiate variables
	var (
		apiServerInfo *api.ApiServerInfo
		apiChan       chan *api.SessionMessage
		apiRec        chan *api.Receiver
	)

	BeforeEach(func() {
		apiChan = make(chan *api.SessionMessage, api.ApiChannelCapacity)
		apiRec = make(chan *api.Receiver)
		apiServerInfo = &api.ApiServerInfo{ApiChannel: apiChan, ApiReceiver: apiRec}
	})

	Context("PostSmContexts 200 test", func() {
		It("should successfully handle a POST request with JSON data and binary file", func() {
			// JSON data from a file
			jsonFilePath := "/home/ubuntu/test_data/smContextCreate.json"
			jsonFile, err := os.Open(jsonFilePath)
			Expect(err).To(BeNil())
			defer jsonFile.Close()

			// Binary file
			binaryFilePath := "/home/ubuntu/test_data/n1msgtest"
			binaryFile, err := os.Open(binaryFilePath)
			Expect(err).To(BeNil())
			defer binaryFile.Close()

			// Create a new buffer to write the multipart payload
			var b bytes.Buffer
			w := multipart.NewWriter(&b)

			// JSON data from the file to the buffer
			jsonField, _ := w.CreateFormField("jsonData")
			_, err = io.Copy(jsonField, jsonFile)
			Expect(err).To(BeNil())

			// Binary data to the buffer
			binaryField, _ := w.CreateFormFile("binaryDataN1SmMessage", "n1msg")
			_, err = io.Copy(binaryField, binaryFile)
			Expect(err).To(BeNil())

			w.Close()

			// HTTP request with the multipart payload
			request, _ := http.NewRequest("POST", "/nsmf-pdusession/v1/sm-contexts", &b)
			request.Header.Set("Content-Type", w.FormDataContentType())

			// Response recorder
			response := httptest.NewRecorder()

			//Add appropriate response to Receive Channel
			go func() {
				apiServerInfo.ApiReceiver <- &api.Receiver{RecievedResponse: openapiserver.Response(201, openapiserver.SmContextCreatedData{
					UpCnxState:   1,
					PduSessionId: 1257,
				}), RecievedErr: nil}
			}()

			// Create SM Context
			apiServerInfo.PostSmContexts(response, request)

			// 200 OK must be the response
			Expect(response.Code).To(Equal(201))
		})
	})
})

var _ = Describe("ReleaseSMContext", func() {
	//Initiate variables
	var (
		apiServerInfo *api.ApiServerInfo
		apiChan       chan *api.SessionMessage
		apiRec        chan *api.Receiver
	)

	BeforeEach(func() {
		apiChan = make(chan *api.SessionMessage, api.ApiChannelCapacity)
		apiRec = make(chan *api.Receiver)
		apiServerInfo = &api.ApiServerInfo{ApiChannel: apiChan, ApiReceiver: apiRec}
	})

	Context("ReleaseSmContexts 200 test", func() {
		It("should successfully handle a POST request with JSON data and binary file", func() {
			// JSON data from a file
			jsonFilePath := "/home/ubuntu/test_data/smContextRelease.json"
			jsonFile, err := os.Open(jsonFilePath)
			Expect(err).To(BeNil())
			defer jsonFile.Close()

			// Binary file
			binaryFilePath := "/home/ubuntu/test_data/n2infotest"
			binaryFile, err := os.Open(binaryFilePath)
			Expect(err).To(BeNil())
			defer binaryFile.Close()

			// Create a new buffer to write the multipart payload
			var b bytes.Buffer
			w := multipart.NewWriter(&b)

			// JSON data from the file to the buffer
			jsonField, _ := w.CreateFormField("jsonData")
			_, err = io.Copy(jsonField, jsonFile)
			Expect(err).To(BeNil())

			// Binary data to the buffer
			binaryField, _ := w.CreateFormFile("binaryDataN2SmInformation", "n2infotest")
			_, err = io.Copy(binaryField, binaryFile)
			Expect(err).To(BeNil())

			w.Close()

			// HTTP request with the multipart payload
			request, _ := http.NewRequest("POST", "/nsmf-pdusession/v1/sm-contexts/smcontext/release", &b)
			request.Header.Set("Content-Type", w.FormDataContentType())

			// Response recorder
			response := httptest.NewRecorder()

			//Add appropriate data to receive channel
			go func() {
				apiServerInfo.ApiReceiver <- &api.Receiver{RecievedResponse: openapiserver.Response(http.StatusNoContent, nil), RecievedErr: nil}
			}()

			// Create SM Context
			apiServerInfo.ReleaseSmContext(response, request)

			// 204 must be the response
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})
	})
})

var _ = Describe("RetrieveSMContext", func() {
	//Initiate variables
	var (
		apiServerInfo *api.ApiServerInfo
		apiChan       chan *api.SessionMessage
		apiRec        chan *api.Receiver
	)

	BeforeEach(func() {
		apiChan = make(chan *api.SessionMessage, api.ApiChannelCapacity)
		apiRec = make(chan *api.Receiver)
		apiServerInfo = &api.ApiServerInfo{ApiChannel: apiChan, ApiReceiver: apiRec}
	})

	Context("RetrieveSmContext 200 test", func() {
		It("should successfully handle a POST request with JSON data", func() {
			// JSON data from a file
			jsonFilePath := "/home/ubuntu/test_data/smContextRetrieve.json"
			jsonFile, err := os.Open(jsonFilePath)
			Expect(err).To(BeNil())
			defer jsonFile.Close()

			// Read JSON data from the file
			jsonData, err := ioutil.ReadAll(jsonFile)
			Expect(err).To(BeNil())

			// HTTP request with JSON data
			request, _ := http.NewRequest("POST", "/nsmf-pdusession/v1/sm-contexts/smcontext/retrieve", bytes.NewBuffer(jsonData))
			request.Header.Set("Content-Type", "application/json")

			// Response recorder
			response := httptest.NewRecorder()

			//Add appropriate data to receive channel
			go func() {
				apiServerInfo.ApiReceiver <- &api.Receiver{RecievedResponse: openapiserver.Response(http.StatusOK, openapiserver.SmContextRetrievedData{
					UeEpsPdnConnection: "",
				}), RecievedErr: nil}
			}()

			// Retrieve SM Context
			apiServerInfo.RetrieveSmContext(response, request)

			// 200 OK must be the response
			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})
})

var _ = Describe("UpdateSMContext", func() {
	//Initiate variables
	var (
		apiServerInfo *api.ApiServerInfo
		apiChan       chan *api.SessionMessage
		apiRec        chan *api.Receiver
	)

	BeforeEach(func() {
		apiChan = make(chan *api.SessionMessage, api.ApiChannelCapacity)
		apiRec = make(chan *api.Receiver)
		apiServerInfo = &api.ApiServerInfo{ApiChannel: apiChan, ApiReceiver: apiRec}
	})

	Context("UpdateSmContexts 200 test", func() {
		It("should successfully handle a POST request with JSON data and binary file", func() {
			// JSON data from a file
			jsonFilePath := "/home/ubuntu/test_data/smContextUpdate.json"
			jsonFile, err := os.Open(jsonFilePath)
			Expect(err).To(BeNil())
			defer jsonFile.Close()

			// Binary file
			n1binaryFilePath := "/home/ubuntu/test_data/n1msgtest"
			n1binaryFile, err := os.Open(n1binaryFilePath)
			Expect(err).To(BeNil())
			defer n1binaryFile.Close()

			n2binaryFilePath := "/home/ubuntu/test_data/n2infotest"
			n2binaryFile, err := os.Open(n2binaryFilePath)
			Expect(err).To(BeNil())
			defer n2binaryFile.Close()

			n2extbinaryFilePath := "/home/ubuntu/test_data/n2infoext1test"
			n2extbinaryFile, err := os.Open(n2extbinaryFilePath)
			Expect(err).To(BeNil())
			defer n2extbinaryFile.Close()

			// Create a new buffer to write the multipart payload
			var b bytes.Buffer
			w := multipart.NewWriter(&b)

			// JSON data from the file to the buffer
			jsonField, _ := w.CreateFormField("jsonData")
			_, err = io.Copy(jsonField, jsonFile)
			Expect(err).To(BeNil())

			// Binary data to the buffer
			n1binaryField, _ := w.CreateFormFile("binaryDataN1SmMessage", "n1msgtest")
			_, err = io.Copy(n1binaryField, n1binaryFile)
			Expect(err).To(BeNil())

			n2binaryField, _ := w.CreateFormFile("binaryDataN2SmInformation", "n2infotest")
			_, err = io.Copy(n2binaryField, n2binaryFile)
			Expect(err).To(BeNil())

			n2extbinaryField, _ := w.CreateFormFile("binaryDataN2SmInformationExt1", "n2infoext1test")
			_, err = io.Copy(n2extbinaryField, n2extbinaryFile)
			Expect(err).To(BeNil())

			w.Close()

			// HTTP request with the multipart payload
			request, _ := http.NewRequest("POST", "/nsmf-pdusession/v1/sm-contexts/smContextRef/modify", &b)
			request.Header.Set("Content-Type", w.FormDataContentType())

			// Response recorder
			response := httptest.NewRecorder()

			//Load updated data file
			updateData := openapiserver.SmContextUpdatedData{}
			file, err := os.Open("/home/ubuntu/test_data/smContextUpdate.json")
			Expect(err).To(BeNil())
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}

			//Decode to UpdatedSmContext data
			decoder := json.NewDecoder(file)
			decoder.Decode(&updateData)

			//Add appropriate data to receive channel
			go func() {
				apiServerInfo.ApiReceiver <- &api.Receiver{RecievedResponse: openapiserver.Response(http.StatusOK, openapiserver.SmContextUpdatedData{
					UpCnxState:   1,
					N1SmMsg:      updateData.N1SmMsg,
					N2SmInfo:     updateData.N2SmInfo,
					N2SmInfoType: updateData.N2SmInfoType,
				}), RecievedErr: nil}
			}()

			// Create SM Context
			apiServerInfo.UpdateSmContext(response, request)

			// 200 OK must be the response
			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})
})

var _ = Describe("WatchChannel", func() {
	//Initiate variables
	var (
		apiServerInfo *api.ApiServerInfo
		apiChan       chan *api.SessionMessage
		apiRec        chan *api.Receiver
	)

	BeforeEach(func() {
		apiChan = make(chan *api.SessionMessage, api.ApiChannelCapacity)
		apiRec = make(chan *api.Receiver)
		apiServerInfo = &api.ApiServerInfo{ApiChannel: apiChan, ApiReceiver: apiRec}
	})

	Context("WatchApiChannel test", func() {
		It("should successfully return channel data", func() {
			//Add appropriate data to api channel
			go func() {
				apiServerInfo.ApiChannel <- &api.SessionMessage{MsgType: sm.NSMF_CREATE_SM_CONTEXT_RESPONSE, SmContextRefID: "test"}
			}()

			//Check if channel is received
			data := apiServerInfo.WatchApiChannel()
			Expect(data).To(Equal(apiServerInfo.ApiChannel))

		})
	})

	Context("WatchRecChannel test", func() {
		It("should successfully return channel data", func() {
			//Add appropriate data to receive channel
			go func() {
				apiServerInfo.ApiReceiver <- &api.Receiver{RecievedResponse: openapiserver.Response(http.StatusNoContent, nil), RecievedErr: nil}
			}()

			//Check if channel is received
			data := apiServerInfo.WatchRecChannel()
			Expect(data).To(Equal(apiServerInfo.ApiReceiver))

		})
	})
})
