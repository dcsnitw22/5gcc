package openapi

func database() map[string]SessionManagementSubscriptionData {
	// Create a map with string keys (ueid) and SdmSubscription values
	subscriptions := make(map[string]SessionManagementSubscriptionData)

	// Define and initialize SdmSubscription objects
	createdSubscription1 := SessionManagementSubscriptionData{
		SingleNssai: Snssai{
			Sst: 1,
			Sd:  "sd1",
		},
		InternalGroupIds:               []string{"internalGroup1", "internalGroup2"},
		SharedVnGroupDataIds:           map[string]string{"groupId1": "sharedDataId1", "groupId2": "sharedDataId2"},
		SharedDnnConfigurationsId:      "sharedDnnConfigurationsId1",
		OdbPacketServices:              &OdbPacketServices{ /* Fill OdbPacketServices fields here */ },
		TraceData:                      &TraceData{ /* Fill TraceData fields here */ },
		SharedTraceDataId:              "sharedTraceDataId1",
		ExpectedUeBehavioursList:       map[string]ExpectedUeBehaviourData{"dnn1": ExpectedUeBehaviourData{ /* Fill ExpectedUeBehaviourData fields here */ }},
		SuggestedPacketNumDlList:       map[string]SuggestedPacketNumDl{"dnn1": SuggestedPacketNumDl{ /* Fill SuggestedPacketNumDl fields here */ }},
		Var3gppChargingCharacteristics: "3gppChargingCharacteristics1",
		SupportedFeatures:              "supportedFeatures1",
	}

	createdSubscription2 := SessionManagementSubscriptionData{
		SingleNssai: Snssai{
			Sst: 2,
			Sd:  "sd2",
		},
		InternalGroupIds:               []string{"internalGroup3", "internalGroup4"},
		SharedVnGroupDataIds:           map[string]string{"groupId3": "sharedDataId3", "groupId4": "sharedDataId4"},
		SharedDnnConfigurationsId:      "sharedDnnConfigurationsId2",
		OdbPacketServices:              &OdbPacketServices{ /* Fill OdbPacketServices fields here */ },
		TraceData:                      &TraceData{ /* Fill TraceData fields here */ },
		SharedTraceDataId:              "sharedTraceDataId2",
		ExpectedUeBehavioursList:       map[string]ExpectedUeBehaviourData{"dnn3": ExpectedUeBehaviourData{ /* Fill ExpectedUeBehaviourData fields here */ }},
		SuggestedPacketNumDlList:       map[string]SuggestedPacketNumDl{"dnn3": SuggestedPacketNumDl{ /* Fill SuggestedPacketNumDl fields here */ }},
		Var3gppChargingCharacteristics: "3gppChargingCharacteristics2",
		SupportedFeatures:              "supportedFeatures2",
	}

	createdSubscription3 := SessionManagementSubscriptionData{
		SingleNssai: Snssai{
			Sst: 3,
			Sd:  "sd3",
		},
		InternalGroupIds:               []string{"internalGroup5", "internalGroup6"},
		SharedVnGroupDataIds:           map[string]string{"groupId5": "sharedDataId5", "groupId6": "sharedDataId6"},
		SharedDnnConfigurationsId:      "sharedDnnConfigurationsId3",
		OdbPacketServices:              &OdbPacketServices{ /* Fill OdbPacketServices fields here */ },
		TraceData:                      &TraceData{ /* Fill TraceData fields here */ },
		SharedTraceDataId:              "sharedTraceDataId3",
		ExpectedUeBehavioursList:       map[string]ExpectedUeBehaviourData{"dnn5": ExpectedUeBehaviourData{ /* Fill ExpectedUeBehaviourData fields here */ }},
		SuggestedPacketNumDlList:       map[string]SuggestedPacketNumDl{"dnn5": SuggestedPacketNumDl{ /* Fill SuggestedPacketNumDl fields here */ }},
		Var3gppChargingCharacteristics: "3gppChargingCharacteristics3",
		SupportedFeatures:              "supportedFeatures3",
	}

	// Populate the map with the createdSubscription objects
	subscriptions["erb32"] = createdSubscription1
	subscriptions["sup12"] = createdSubscription2
	subscriptions["dsv21"] = createdSubscription3

	return subscriptions
}
