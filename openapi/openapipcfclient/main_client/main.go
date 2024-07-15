package main

//TODO smf should store policy in redis
// folowing 4.16.4 flow
import (
	"context"
	"fmt"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"

	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DBClient struct {
	sqlClient *sql.DB
}

func (d DBClient) DBConnect(dbname string) *sql.DB {

	cfg := mysql.Config{
		User:   "lovekush",
		Passwd: "password123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbname,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

func main() {
	cfg := openapi.NewConfiguration()
	cli := openapi.NewAPIClient(cfg)

	//Post request
	supi := "121"
	smfId := "123"
	mnc := "15"
	mcc := "2"
	s := "192.168.1.3"
	s1 := "internet"
	t := true
	ueTimeZone := "-08:00+1"
	pei := "imeisv-3660093949053296"
	var prtlvl int32 = 127
	var f int32 = 4
	sd := "7C7e8b"
	i := "f9C91F30FB92fb206dDa9Ee9175F3D6BaF6CABFcaA9c7A4595F7bDe2ef9f2F"
	servNfInstId := "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	suppFeat := "b78f3f7B6e3BfDd7BA006042D7dBb818cA0E65A3de0e"
	var anCharIpAddr openapi.AccNetChargingAddress = openapi.AccNetChargingAddress{
		AnChargIpv4Addr: &s,
	}
	fmt.Println("=======Post request=======")

	r := openapi.SmPolicyContextData{
		AccNetChId: &openapi.AccNetChId{
			AccNetChaIdValue: 123,
			RefPccRuleIds:    []string{"123"},
			SessionChScope:   &t,
		},
		ChargEntityAddr: *openapi.NewNullableAccNetChargingAddress(&anCharIpAddr),
		// Gpsi: ,get from existing values
		Supi:        supi,
		InvalidSupi: &t,
		InterGrpIds: []string{
			"F6dA0854-825-47-bbA06a40aAD7",
		},
		PduSessionId: 225,
		// PduSessionType: "IPV4",
		Chargingcharacteristics: &s1,
		Dnn:                     s1,
		// AccessType: "NR",
		ServingNetwork: &openapi.NetworkId{
			Mcc: &mcc,
			Mnc: &mnc,
		},
		// UserLocationInfo: ,Put later
		UeTimeZone:  &ueTimeZone,
		Pei:         &pei,
		Ipv4Address: &s,
		SubsSessAmbr: &openapi.Ambr{
			Uplink:   "2.75983454480949148159094299358505246856784447265331989387573764936227804192872466959178625581 Kbps",
			Downlink: "732896568792618632691539.861064764823368462241753975269144937 Tbps",
		},
		SubsDefQos: &openapi.SubscribedDefaultQos{
			Var5qi: 255,
			Arp: openapi.Arp{
				PriorityLevel: *openapi.NewNullableInt32(&prtlvl),
				//preemptivecap and premptivevuln correct based on the existing code
			},
			PriorityLevel: &prtlvl,
		},
		NumOfPackFilter:        &f,
		Online:                 &t,
		Offline:                &t,
		Var3gppPsDataOffStatus: &t,
		RefQosIndication:       &t,
		TraceReq: *openapi.NewNullableTraceData(&openapi.TraceData{
			TraceRef: "55466-d94aA7",
			// TraceDepth: "",correct based on current code
			NeTypeList:               "536FE7F15B8b3b1ce6fD5bca7CD24ebD1a9c00FDcdcba4f8E4E2ebaf8Dc15a1Dec9C5be9dB6A9fee84dF5Ba0E",
			EventList:                "3aBDb41F36fb982744ACc789f",
			CollectionEntityIpv4Addr: &s,
			InterfaceList:            &i,
		}),
		SliceInfo: openapi.Snssai{
			Sst: 25,
			Sd:  &sd,
		},
		// QosFlowUsage: , correct based on current code
		ServNfId: &openapi.ServingNfIdentity{
			ServNfInstId: &servNfInstId,
			Guami: openapi.NewGuami(openapi.PlmnId{
				Mcc: mcc,
				Mnc: mnc,
			}, "12311"),
			AnGwAddr: *openapi.NewNullableAnGwAddress(&openapi.AnGwAddress{
				AnGwIpv4Addr: &s,
			}),
		},
		SuppFeat: &suppFeat,
		SmfId:    &smfId,
		// RecoveryTime: ,check if to put or not
	}
	a := openapi.ApiSmPoliciesPostRequest{
		ApiService: cli.DefaultAPI,
	}
	sessdes, res, err := cli.DefaultAPI.SmPoliciesPostExecute(a.SmPolicyContextData(r))
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("sessdes:%+v,\nresponse:%+v", sessdes, res)
	// var dbcli DBClient
	// dbcli.sqlClient = dbcli.DBConnect("pcf_client")
	// // var smPid string
	// for _, v := range *sessdes.SessRules {
	// 	// smPid := v.SessRuleId
	// 	_, err = dbcli.sqlClient.Query("Insert into smPolicyIdList (smPolicyId,supi)"+
	// 		" values(?,?)", v.SessRuleId, supi)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	//Get request
	fmt.Println("=======Get request=======")
	smPid := supi + smfId
	a1 := openapi.ApiSmPoliciesSmPolicyIdGetRequest{
		ApiService: cli.DefaultAPI,
	}
	fmt.Println(a1.ApiService.SmPoliciesSmPolicyIdGet(context.Background(), smPid).Execute())

	// //Update request
	fmt.Println("=======Update request=======")
	a2 := openapi.ApiSmPoliciesSmPolicyIdUpdatePostRequest{
		ApiService: cli.DefaultAPI,
	}

	// dbcli.sqlClient = dbcli.DBConnect("pcf")
	// dbcli.sqlClient.Query("UPDATE subscription SET polid=?,start_date=?"+
	// 	" where supi=?", 2, time.Now(), supi)

	r2 := openapi.SmPolicyUpdateContextData{
		AccNetChIds: []openapi.AccNetChId{
			{
				AccNetChaIdValue: 123,
				RefPccRuleIds:    []string{"123"},
				SessionChScope:   &t,
			},
		},
		ServingNetwork: &openapi.NetworkId{
			Mcc: &mcc,
			Mnc: &mnc,
		},
		Ipv4Address: &s,
		SubsDefQos: openapi.NewSubscribedDefaultQos(255,
			openapi.Arp{PriorityLevel: *openapi.NewNullableInt32(&prtlvl)}),
		NumOfPackFilter: &f,
		SubsSessAmbr: &openapi.Ambr{
			Uplink:   "12 Gbps",
			Downlink: "924974446506444091447729843 bps",
		},
	}
	fmt.Println(a2.ApiService.SmPoliciesSmPolicyIdUpdatePost(context.Background(), smPid).SmPolicyUpdateContextData(r2).Execute())

	// // //Delete request
	fmt.Println("=======Delete request=======")
	// var mnc string = "123"
	// var mcc string = "21"
	r3 := openapi.SmPolicyDeleteData{
		ServingNetwork: &openapi.NetworkId{
			Mnc: &mnc,
			Mcc: &mcc,
		},
	}

	fmt.Println(cli.DefaultAPI.SmPoliciesSmPolicyIdDeletePost(context.Background(), smPid).SmPolicyDeleteData(r3).Execute())

}
