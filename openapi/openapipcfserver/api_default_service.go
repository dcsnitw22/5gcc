/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

import (
	openapi_udr_cli "5gCore/udrClient"
	"context"
	"errors"
	"fmt"
	"strconv"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() DefaultAPIServicer {
	return &DefaultAPIService{}
}

// SmPoliciesPost -
// ! error reponse with 511 need to transfered to api or db code
// TODO send PCCRULE (5.6.2.6) and check for what to send in sessionRule
func (s *DefaultAPIService) SmPoliciesPost(ctx context.Context, smPolicyContextData SmPolicyContextData) (ImplResponse, error) {

	//TODO Put in main code
	var dbcli DBClient
	var err error

	dbcli.redisClient, err = dbcli.DBConnect()

	if err != nil {
		return Response(511, ProblemDetails{
			Cause: "INTERNAL_ERROR",
		}), err
	}

	// if dbcli.PolicyExists(ctx, smPolicyContextData.Supi) {
	//put code here
	// }

	var sessRule map[string]SessionRule = map[string]SessionRule{}
	var pccRule map[string]PccRule = map[string]PccRule{}

	cfg := openapi_udr_cli.NewConfiguration()
	udrCli := openapi_udr_cli.NewAPIClient(cfg)
	fmt.Println("-------------------------------")
	fmt.Println("supi"+smPolicyContextData.Supi, "mcc"+smPolicyContextData.ServNfId.Guami.PlmnId.Mcc+"mnc"+smPolicyContextData.ServNfId.Guami.PlmnId.Mnc, "dnn"+smPolicyContextData.Dnn, smPolicyContextData.SliceInfo.Sst, smPolicyContextData.SliceInfo.Sd)
	fmt.Println("-------------------------------")
	smSubsData, _, err := udrCli.SessionManagementSubscriptionDataAPI.QuerySmData(
		ctx,
		smPolicyContextData.Supi,
		smPolicyContextData.ServNfId.Guami.PlmnId.Mcc+smPolicyContextData.ServNfId.Guami.PlmnId.Mnc,
	).Dnn(smPolicyContextData.Dnn).SingleNssai(openapi_udr_cli.Snssai{
		Sst: smPolicyContextData.SliceInfo.Sst,
		Sd:  &smPolicyContextData.SliceInfo.Sd,
	}).Execute()
	if err != nil {
		return Response(403, ProblemDetails{
			Detail: "POLICY_CONTEXT_DENIED",
		}), err
	}

	fmt.Printf("responseBody:%+v", smSubsData.ExtendedSmSubsData.IndividualSmSubsData)

	for _, data := range smSubsData.ExtendedSmSubsData.IndividualSmSubsData {
		sst := fmt.Sprint(data.SingleNssai.Sst)
		sd := *data.SingleNssai.Sd
		// var dnn string
		if data.DnnConfigurations == nil {
			data.DnnConfigurations = &map[string]openapi_udr_cli.DnnConfiguration{}
		}
		for _, v := range *data.DnnConfigurations {
			// dnn = k
			sessRuleId := sst + sd + strconv.Itoa(int(smPolicyContextData.PduSessionId))
			pccRuleId := sst + sd + strconv.Itoa(int(smPolicyContextData.PduSessionId))

			qosData := QosData{
				QosId:  sst + sd + strconv.Itoa(int(v.Var5gQosProfile.Var5qi)), // change QosId later, use var5qi
				Var5qi: v.Var5gQosProfile.Var5qi,
				Arp: Arp{
					PriorityLevel: v.Var5gQosProfile.Arp.PriorityLevel.Get(),
				},
				PriorityLevel: v.Var5gQosProfile.PriorityLevel,
			}

			created := dbcli.QosCreate(ctx, qosData.QosId, qosData)
			if !created {
				continue
			}

			sessRule[sessRuleId] = SessionRule{
				AuthSessAmbr: Ambr{
					Uplink:   v.SessionAmbr.Uplink,
					Downlink: v.SessionAmbr.Downlink,
				},
				AuthDefQos: AuthorizedDefaultQos{
					Var5qi: v.Var5gQosProfile.Var5qi,
					Arp: Arp{
						PriorityLevel: v.Var5gQosProfile.Arp.PriorityLevel.Get(),
					},
					PriorityLevel: v.Var5gQosProfile.PriorityLevel,
				},
				SessRuleId: sessRuleId,
			}
			pccRule[pccRuleId] = PccRule{
				PccRuleId:  pccRuleId,
				RefQosData: []string{qosData.QosId},
			}
		}

	}

	//TODO sessionRule to be corrected
	smPolicyDes := SmPolicyDecision{
		SessRules: sessRule,
		PccRules:  &pccRule,
	}

	created := dbcli.PolicyCreate(
		ctx,
		smPolicyContextData.Supi+smPolicyContextData.SmfId,
		smPolicyDes,
	)

	if !created {
		return Response(403, ProblemDetails{
			Detail: "POLICY_CONTEXT_DENIED",
		}), errors.New("unable to create policy context")
	}

	created = dbcli.ContextCreate(
		ctx,
		smPolicyContextData.Supi+smPolicyContextData.SmfId,
		SmPolicyContextData{
			Supi:         smPolicyContextData.Supi,
			Dnn:          smPolicyContextData.Dnn,
			PduSessionId: smPolicyContextData.PduSessionId,
			ServNfId: ServingNfIdentity{
				Guami: Guami{
					PlmnId: smPolicyContextData.ServNfId.Guami.PlmnId,
					AmfId:  smPolicyContextData.ServNfId.Guami.AmfId,
				},
			},
			SliceInfo: Snssai{
				Sst: smPolicyContextData.SliceInfo.Sst,
				Sd:  smPolicyContextData.SliceInfo.Sd,
			},
			SmfId: smPolicyContextData.SmfId,
		},
	)

	if !created {
		//! ask if to keep this or not
		// deleted := dbcli.PolicyDelete(ctx, smPolicyContextData.Supi+smPolicyContextData.SmfId)
		// if !deleted {
		// 	fmt.Println("unable to revert changes inconsistent configuration")
		// }
		return Response(403, ProblemDetails{
			Detail: "POLICY_CONTEXT_DENIED",
		}), errors.New("unable to create policy context")
	}

	fmt.Println(smPolicyDes)
	return Response(201, smPolicyDes), nil
}

// SmPoliciesSmPolicyIdDeletePost -
func (s *DefaultAPIService) SmPoliciesSmPolicyIdDeletePost(ctx context.Context, smPolicyId string, smPolicyDeleteData SmPolicyDeleteData) (ImplResponse, error) {

	var dbcli DBClient
	var err error
	dbcli.redisClient, err = dbcli.DBConnect()
	if err != nil {
		return Response(511, ProblemDetails{
			Detail: "INTERNAL_ERROR",
		}), err
	}

	//! no reponse found for this case in spec
	//TODO recheck in spec
	if !dbcli.PolicyExists(ctx, smPolicyId) {
		return Response(404, ProblemDetails{
			Detail: "INVALID_SMPOLICYID",
		}), errors.New("invalid smpolicyid")
	}

	//! error handling to query about : do i have to revert on failure
	deleted := dbcli.PolicyDelete(ctx, smPolicyId)
	if !deleted {
		return Response(511, ProblemDetails{
			Detail: "NOT_DELETED",
		}), errors.New("not deleted")
	}

	deleted = dbcli.ContextDelete(ctx, smPolicyId)
	if !deleted {
		return Response(511, ProblemDetails{
			Detail: "NOT_DELETED",
		}), errors.New("not deleted")
	}

	return Response(204, nil), nil
}

// // SmPoliciesSmPolicyIdGet -
func (s *DefaultAPIService) SmPoliciesSmPolicyIdGet(ctx context.Context, smPolicyId string) (ImplResponse, error) {

	var dbcli DBClient
	var err error
	dbcli.redisClient, err = dbcli.DBConnect()
	if err != nil {
		return Response(511, ProblemDetails{
			Detail: "INTERNAL_ERROR",
		}), err
	}

	// 	//! no reponse found for this case in spec
	// 	//TODO recheck in spec
	if !dbcli.PolicyExists(ctx, smPolicyId) {
		return Response(404, ProblemDetails{
			Detail: "INVALID_SMPOLICYID",
		}), errors.New("invalid smpolicyid")
	}

	smPolDec, err := dbcli.PolicyGet(ctx, smPolicyId)
	// 	//TODO check spec for proper response
	if err != nil {
		return Response(511, ProblemDetails{}), err
	}

	// 	//TODO store context in redis
	smConData, err := dbcli.ContextGet(ctx, smPolicyId)
	if err != nil {
		fmt.Println(err)
		return Response(511, ProblemDetails{}), err
	}

	return Response(200, SmPolicyControl{
		Context: smConData,
		Policy:  smPolDec,
	}), nil
}

// // SmPoliciesSmPolicyIdUpdatePost -
func (s *DefaultAPIService) SmPoliciesSmPolicyIdUpdatePost(ctx context.Context, smPolicyId string, smPolicyUpdateContextData SmPolicyUpdateContextData) (ImplResponse, error) {

	var dbcli DBClient
	var err error
	dbcli.redisClient, err = dbcli.DBConnect()
	if err != nil {
		return Response(511, ProblemDetails{
			Detail: "INTERNAL_ERROR",
		}), err
	}

	// 	//! no reponse found for this case in spec
	// 	//TODO recheck in spec
	if !dbcli.PolicyExists(ctx, smPolicyId) {
		return Response(404, ProblemDetails{
			Detail: "INVALID_SMPOLICYID",
		}), errors.New("invalid smpolicyid")
	}

	smPolicyContextData, err := dbcli.ContextGet(ctx, smPolicyId)
	if err != nil {
		return Response(511, ProblemDetails{
			Detail: "INTERNAL_ERROR",
		}), err
	}

	// fmt.Println(smPolicyContextData)

	// var sessRule map[string]SessionRule = map[string]SessionRule{}

	// cfg := openapi_udr_cli.NewConfiguration()
	// udrCli := openapi_udr_cli.NewAPIClient(cfg)
	// smSubsData, _, err := udrCli.SessionManagementSubscriptionDataAPI.QuerySmData(
	// 	ctx,
	// 	smPolicyContextData.Supi,
	// 	smPolicyContextData.ServNfId.Guami.PlmnId.Mcc+smPolicyContextData.ServNfId.Guami.PlmnId.Mcc,
	// ).Dnn(smPolicyContextData.Dnn).SingleNssai(openapi_udr_cli.Snssai{
	// 	Sst: smPolicyContextData.SliceInfo.Sst,
	// 	Sd:  &smPolicyContextData.SliceInfo.Sd,
	// }).Execute()
	// // fmt.Println(err)
	// if err != nil {
	// 	return Response(403, ProblemDetails{
	// 		Detail: "POLICY_CONTEXT_DENIED",
	// 	}), err
	// }

	// // fmt.Printf("responseBody:%v", resp.Body)

	// for _, data := range smSubsData.ExtendedSmSubsData.IndividualSmSubsData {
	// 	sst := fmt.Sprint(data.SingleNssai.Sst)
	// 	sd := *data.SingleNssai.Sd
	// 	var dnn string
	// 	if data.DnnConfigurations == nil {
	// 		data.DnnConfigurations = &map[string]openapi_udr_cli.DnnConfiguration{}
	// 	}
	// 	for k, v := range *data.DnnConfigurations {
	// 		dnn = k
	// 		sessRuleId := sst + sd + dnn
	// 		sessRule[sessRuleId] = SessionRule{
	// 			AuthSessAmbr: Ambr{
	// 				Uplink:   v.SessionAmbr.Uplink,
	// 				Downlink: v.SessionAmbr.Downlink,
	// 			},
	// 			AuthDefQos: AuthorizedDefaultQos{
	// 				Var5qi: v.Var5gQosProfile.Var5qi,
	// 				Arp: Arp{
	// 					PriorityLevel: v.Var5gQosProfile.Arp.PriorityLevel.Get(),
	// 				},
	// 				PriorityLevel: v.Var5gQosProfile.PriorityLevel,
	// 			},
	// 			SessRuleId: sessRuleId,
	// 		}
	// 	}
	// }

	// //TODO correct join logic in udr and check create policy function in redis
	// smPolicyDes := SmPolicyDecision{
	// 	SessRules: sessRule,
	// }

	var sessRule map[string]SessionRule = map[string]SessionRule{}
	var pccRule map[string]PccRule = map[string]PccRule{}

	cfg := openapi_udr_cli.NewConfiguration()
	udrCli := openapi_udr_cli.NewAPIClient(cfg)
	smSubsData, _, err := udrCli.SessionManagementSubscriptionDataAPI.QuerySmData(
		ctx,
		smPolicyContextData.Supi,
		smPolicyContextData.ServNfId.Guami.PlmnId.Mcc+smPolicyContextData.ServNfId.Guami.PlmnId.Mnc,
	).Dnn(smPolicyContextData.Dnn).SingleNssai(openapi_udr_cli.Snssai{
		Sst: smPolicyContextData.SliceInfo.Sst,
		Sd:  &smPolicyContextData.SliceInfo.Sd,
	}).Execute()
	if err != nil {
		return Response(403, ProblemDetails{
			Detail: "POLICY_CONTEXT_DENIED",
		}), err
	}

	fmt.Printf("responseBody:%+v", smSubsData.ExtendedSmSubsData.IndividualSmSubsData)

	for _, data := range smSubsData.ExtendedSmSubsData.IndividualSmSubsData {
		sst := fmt.Sprint(data.SingleNssai.Sst)
		sd := *data.SingleNssai.Sd
		var dnn string
		if data.DnnConfigurations == nil {
			data.DnnConfigurations = &map[string]openapi_udr_cli.DnnConfiguration{}
		}
		for k, v := range *data.DnnConfigurations {
			dnn = k
			sessRuleId := sst + sd + dnn
			pccRuleId := sst + sd + dnn

			qosData := QosData{
				QosId:  sst + sd + dnn, // change QosId later
				Var5qi: v.Var5gQosProfile.Var5qi,
				Arp: Arp{
					PriorityLevel: v.Var5gQosProfile.Arp.PriorityLevel.Get(),
				},
				PriorityLevel: v.Var5gQosProfile.PriorityLevel,
			}

			created := dbcli.QosCreate(ctx, qosData.QosId, qosData)
			if !created {
				continue // error or continue ASk GURU
			}

			sessRule[sessRuleId] = SessionRule{
				AuthSessAmbr: Ambr{
					Uplink:   v.SessionAmbr.Uplink,
					Downlink: v.SessionAmbr.Downlink,
				},
				AuthDefQos: AuthorizedDefaultQos{
					Var5qi: v.Var5gQosProfile.Var5qi,
					Arp: Arp{
						PriorityLevel: v.Var5gQosProfile.Arp.PriorityLevel.Get(),
					},
					PriorityLevel: v.Var5gQosProfile.PriorityLevel,
				},
				SessRuleId: sessRuleId,
			}
			pccRule[pccRuleId] = PccRule{
				PccRuleId: pccRuleId,
				// Precedence: , ASK GURU ABOUT THE LOGIC
				RefQosData: []string{qosData.QosId},
			}
		}

	}

	//TODO sessionRule to be corrected
	smPolicyDes := SmPolicyDecision{
		SessRules: sessRule,
		PccRules:  &pccRule,
	}

	updated := dbcli.PolicyCreate(
		ctx,
		smPolicyContextData.Supi+smPolicyContextData.SmfId,
		smPolicyDes,
	)
	if !updated {
		return Response(403, ProblemDetails{
			Detail: "POLICY_CONTEXT_DENIED",
		}), errors.New("unable to create policy context")
	}

	return Response(201, smPolicyDes), nil
}