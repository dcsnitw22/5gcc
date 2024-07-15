package openapi

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"log"

// 	"github.com/go-sql-driver/mysql"
// )

// type DBClient struct {
// 	sqlClient *sql.DB
// }

// type dnnConfMap struct {
// 	dnnMap map[string]DnnConfiguration
// }

// func DBConnect(dbname string) *sql.DB {

// 	cfg := mysql.Config{
// 		User:   "keshav",
// 		Passwd: "Keshav@123",
// 		Net:    "tcp",
// 		Addr:   "127.0.0.1:3306",
// 		DBName: dbname,
// 	}

// 	db, err := sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Database Connected!")
// 	return db
// }

// func (dbCli DBClient) GetDataForSliceDnn(ueId string, servingPlmnId string, singleNssai Snssai, dnn string) SmSubsData {

// 	var data SmSubsData
// 	var dnnConf map[string]DnnConfiguration = map[string]DnnConfiguration{}

// 	rows, err := dbCli.sqlClient.Query("select Var5gQosProfileVar5qi,"+
// 		"Var5gQosProfilePriorityLevel,Var5gQosProfileArpPriorityLevel,"+
// 		"AmbrUplink,AmbrDownlink from DnnConfiguration join "+
// 		"(select dnnId from SmSubsData join "+
// 		" (select sNssaiId from sNssai where sst=? and sd=?) "+
// 		"as sv on sv.sNssaiId=SmSubsData.sNssaiId where dnn=? and ueId=? and servingPlmnId=?) as vd "+
// 		"on vd.dnnId=DnnConfiguration.dnnId",
// 		singleNssai.Sst, singleNssai.Sd, dnn, ueId, servingPlmnId)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		fmt.Println("Now rows found")
// 		return data
// 	}
// 	if err != nil {
// 		return data
// 	}
// 	for rows.Next() {
// 		var (
// 			v5qi  int32
// 			plvl  int32
// 			aplvl int32
// 			ulnk  string
// 			dlnk  string
// 		)
// 		rows.Scan(&v5qi,
// 			&plvl,
// 			&aplvl,
// 			&ulnk,
// 			&dlnk,
// 		)

// 		dnnConf[dnn] = DnnConfiguration{
// 			Var5gQosProfile: SubscribedDefaultQos{
// 				Var5qi:        v5qi,
// 				PriorityLevel: plvl,
// 				Arp: Arp{
// 					PriorityLevel: &aplvl,
// 				},
// 			},
// 			SessionAmbr: Ambr{
// 				Uplink:   ulnk,
// 				Downlink: dlnk,
// 			},
// 		}
// 	}
// 	data.IndividualSmSubsData = append(data.IndividualSmSubsData,
// 		SessionManagementSubscriptionData{
// 			SingleNssai: Snssai{
// 				Sst: singleNssai.Sst,
// 				Sd:  singleNssai.Sd,
// 			},
// 			DnnConfigurations: dnnConf,
// 		},
// 	)
// 	return data
// }

// func (dbCli DBClient) GetDataForSlice(ueId string, servingPlmnId string, singleNssai Snssai) SmSubsData {

// 	var data SmSubsData
// 	var dnnConf map[string]DnnConfiguration = map[string]DnnConfiguration{}
// 	rows, err := dbCli.sqlClient.Query("select Var5gQosProfileVar5qi,"+
// 		"Var5gQosProfilePriorityLevel,Var5gQosProfileArpPriorityLevel,"+
// 		"AmbrUplink,AmbrDownlink,dnn from DnnConfiguration join "+
// 		"(select dnnId,dnn from SmSubsData join "+
// 		" (select sNssaiId from sNssai where sst=? and sd=?) "+
// 		"as sv on sv.sNssaiId=SmSubsData.sNssaiId where ueId=? and servingPlmnId=?) as vd "+
// 		"on vd.dnnId=DnnConfiguration.dnnId",
// 		singleNssai.Sst, singleNssai.Sd, ueId, servingPlmnId)

// 	if errors.Is(err, sql.ErrNoRows) {
// 		fmt.Println("no rows found")
// 		return data
// 	}
// 	if err != nil {
// 		return data
// 	}

// 	for rows.Next() {
// 		var (
// 			v5qi  int32
// 			plvl  int32
// 			aplvl int32
// 			ulnk  string
// 			dlnk  string
// 			dnn   string
// 		)
// 		rows.Scan(&v5qi, &plvl, &aplvl, &ulnk, &dlnk, &dnn)

// 		dnnConf[dnn] = DnnConfiguration{
// 			Var5gQosProfile: SubscribedDefaultQos{
// 				Var5qi:        v5qi,
// 				PriorityLevel: plvl,
// 				Arp: Arp{
// 					PriorityLevel: &aplvl,
// 				},
// 			},
// 			SessionAmbr: Ambr{
// 				Uplink:   ulnk,
// 				Downlink: dlnk,
// 			},
// 		}
// 	}
// 	data.IndividualSmSubsData = append(data.IndividualSmSubsData,
// 		SessionManagementSubscriptionData{
// 			SingleNssai: Snssai{
// 				Sst: singleNssai.Sst,
// 				Sd:  singleNssai.Sd,
// 			},
// 			DnnConfigurations: dnnConf,
// 		},
// 	)
// 	return data
// }

// func (dbCli DBClient) GetDataForDnn(ueId string, servingPlmnId string, dnn string) SmSubsData {

// 	var data SmSubsData
// 	rows, err := dbCli.sqlClient.Query("select v.sNssaiId,v.dnnId,Dnn,sst,sd,Var5gQosProfileVar5qi,Var5gQosProfilePriorityLevel,Var5gQosProfileArpPriorityLevel,AmbrUplink,AmbrDownlink from DnnConfiguration join (select sNssai.sNssaiId,sNssai.sst,sNssai.sd,Dnn,dnnId from sNssai join SmSubsData on sNssai.sNssaiId=SmSubsData.sNssaiId where ueId=? and servingPlmnId=?) as v on v.dnnId=DnnConfiguration.dnnId where Dnn=?", dnn, ueId, servingPlmnId)

// 	if errors.Is(err, sql.ErrNoRows) {
// 		fmt.Println("no rows found")
// 		return data
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 		return data
// 	}
// 	var sNssaidnnConfMap map[int]dnnConfMap = map[int]dnnConfMap{}
// 	var sdMap map[int]string = map[int]string{}
// 	var sstMap map[int]int32 = map[int]int32{}
// 	for rows.Next() {
// 		var (
// 			sNssaiID int
// 			sst      int32
// 			sd       string
// 			dnnId    int
// 			dnn      string
// 			v5qi     int32
// 			plvl     int32
// 			aplvl    int32
// 			aul      string
// 			adl      string
// 		)
// 		rows.Scan(&sNssaiID, &dnnId, &dnn, &sst, &sd, &v5qi, &plvl, &aplvl, &aul, &adl)
// 		var dnnConf DnnConfiguration = DnnConfiguration{
// 			Var5gQosProfile: SubscribedDefaultQos{
// 				Var5qi: v5qi,
// 				Arp: Arp{
// 					PriorityLevel: &aplvl,
// 				},
// 				PriorityLevel: plvl,
// 			},
// 			SessionAmbr: Ambr{
// 				Uplink:   aul,
// 				Downlink: adl,
// 			},
// 		}
// 		if _, ok := sNssaidnnConfMap[sNssaiID]; !ok {
// 			sNssaidnnConfMap[sNssaiID] = dnnConfMap{
// 				dnnMap: map[string]DnnConfiguration{},
// 			}
// 		}
// 		sNssaidnnConfMap[sNssaiID].dnnMap[dnn] = dnnConf
// 		sstMap[sNssaiID] = sst
// 		sdMap[sNssaiID] = sd

// 	}
// 	for id, dnnMaps := range sNssaidnnConfMap {
// 		data.IndividualSmSubsData = append(data.IndividualSmSubsData,
// 			SessionManagementSubscriptionData{
// 				DnnConfigurations: dnnMaps.dnnMap,
// 				SingleNssai: Snssai{
// 					Sst: sstMap[id],
// 					Sd:  sdMap[id],
// 				},
// 			},
// 		)
// 	}
// 	return data
// }

// func (dbCli DBClient) GetData(ueId string, servingPlmnId string) SmSubsData {

// 	var data SmSubsData
// 	rows, err := dbCli.sqlClient.Query("select v.sNssaiId,v.dnnId,Dnn,sst,sd,Var5gQosProfileVar5qi,Var5gQosProfilePriorityLevel,Var5gQosProfileArpPriorityLevel,AmbrUplink,AmbrDownlink from DnnConfiguration join (select sNssai.sNssaiId,sNssai.sst,sNssai.sd,Dnn,dnnId from sNssai join SmSubsData on sNssai.sNssaiId=SmSubsData.sNssaiId where ueId=? and servingPlmnId=?) as v on v.dnnId=DnnConfiguration.dnnId;")

// 	if errors.Is(err, sql.ErrNoRows) {
// 		fmt.Println("no rows found")
// 		return data
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 		return data
// 	}
// 	var sNssaidnnConfMap map[int]dnnConfMap = map[int]dnnConfMap{}
// 	var sdMap map[int]string = map[int]string{}
// 	var sstMap map[int]int32 = map[int]int32{}
// 	for rows.Next() {
// 		var (
// 			sNssaiID int
// 			sst      int32
// 			sd       string
// 			dnnId    int
// 			dnn      string
// 			v5qi     int32
// 			plvl     int32
// 			aplvl    int32
// 			aul      string
// 			adl      string
// 		)
// 		rows.Scan(&sNssaiID, &dnnId, &dnn, &sst, &sd, &v5qi, &plvl, &aplvl, &aul, &adl)
// 		var dnnConf DnnConfiguration = DnnConfiguration{
// 			Var5gQosProfile: SubscribedDefaultQos{
// 				Var5qi: v5qi,
// 				Arp: Arp{
// 					PriorityLevel: &aplvl,
// 				},
// 				PriorityLevel: plvl,
// 			},
// 			SessionAmbr: Ambr{
// 				Uplink:   aul,
// 				Downlink: adl,
// 			},
// 		}
// 		if _, ok := sNssaidnnConfMap[sNssaiID]; !ok {
// 			sNssaidnnConfMap[sNssaiID] = dnnConfMap{
// 				dnnMap: map[string]DnnConfiguration{},
// 			}
// 		}
// 		sNssaidnnConfMap[sNssaiID].dnnMap[dnn] = dnnConf
// 		sstMap[sNssaiID] = sst
// 		sdMap[sNssaiID] = sd

// 	}
// 	for id, dnnMaps := range sNssaidnnConfMap {
// 		data.IndividualSmSubsData = append(data.IndividualSmSubsData,
// 			SessionManagementSubscriptionData{
// 				DnnConfigurations: dnnMaps.dnnMap,
// 				SingleNssai: Snssai{
// 					Sst: sstMap[id],
// 					Sd:  sdMap[id],
// 				},
// 			},
// 		)
// 	}
// 	return data
// }

// //select * from DnnConfiguration  join (select sst,sd,ueId,servingPlmnId,Dnn,dnnId from sNssai join SmSubsData on sNssai.sNssaiId=SmSubsData.sNssaiId) as v
// // on v.dnnId=DnnConfiguration.dnnId;
