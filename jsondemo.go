package main

import "encoding/json"

type JsonStruct struct {
	Name	string		`json:"name"`
	Id 		string		`json:"id"`
	Avatar	string		`json:"avatar"`
}


type IntentInfo struct {
	Name					string					`json:"name"`
	Uuid					string					`json:"uuid"`
	Action					string					`json:"action"`
	Confidence				float32					`json:"confidence"`
	Parameters				map[string]interface{}	`json:"parameters"`
	AllRequiredParamsFilled bool 			        `json:"allRequiredParamsFilled"`
	FulfillmentMessages		[]interface{}			`json:"fulfillmentMessages"`
}

type OriginQuery struct {
	Type		string					`json:"type"`
	Value		string					`json:"value"`
	Parameters	map[string]interface{}	`json:"parameters"`
}

type QueryResult struct {
	ResponseID              string      `json:"responseId"`

	Language                string      `json:"language"`
	IntentInfo              IntentInfo  `json:"intentInfo"`
	OriginQuery             OriginQuery `json:"originQuery"`
}

func main()  {
	/*
	在将json字符串拆解成结构体时，json key与目标结构体不必一一对应，可以多一些或者少一些
	 */
	str := `{"name": "xxxxx", "id": "aaaaa", "email": "abcd@efg.com"}`
	var jsonStruct JsonStruct
	err := json.Unmarshal([]byte(str), &jsonStruct)
	if err != nil {
		panic(err)
	}
	println(jsonStruct.Id, jsonStruct.Name, jsonStruct.Avatar)

	val := `{
  "responseId": "7ce0fcb1-2a3b-43c8-b1a6-248ca25ff7e3-712767ed",
  "queryResult": {
    "queryText": "tell me about you",
    "action": "smalltalk.agent.acquaintance_abc",
    "parameters": {
      "AcqParam": []
    },
    "allRequiredParamsPresent": true,
    "fulfillmentMessages": [
      {
        "platform": "ACTIONS_ON_GOOGLE",
        "listSelect": {
          "title": "This is List Title",
          "items": [
            {
              "info": {
                "key": "This is the option key"
              },
              "title": "This is the Item Title",
              "description": "This is the item descrption",
              "image": {
                "imageUri": "http://www.gsfdcy.com/data/img/29/1448954-endless-legend-wallpaper.jpg",
                "accessibilityText": "beautiful image"
              }
            },
            {
              "info": {
                "key": "Option key"
              },
              "title": "Another Item Title",
              "description": "Another Item descrption",
              "image": {
                "imageUri": "http://www.gsfdcy.com/data/img/29/1448806-endless-legend-wallpaper.jpg",
                "accessibilityText": "another beautiful"
              }
            }
          ]
        }
      },
      {
        "text": {
          "text": [
            ""
          ]
        }
      }
    ],
    "outputContexts": [
      {
        "name": "projects/small-talk-1-xnlfux/agent/sessions/ad0ec261-ff02-7778-264f-3905883f34ec/contexts/smalltalkagentacquaintance-followup",
        "lifespanCount": 2,
        "parameters": {
          "AcqParam.original": [],
          "AcqParam": []
        }
      }
    ],
    "intent": {
      "name": "projects/small-talk-1-xnlfux/agent/intents/6fcd9648-9a31-4ae7-bcb5-7c3029b91bf0",
      "displayName": "smalltalk.agent.acquaintance"
    },
    "intentDetectionConfidence": 1,
    "languageCode": "en"
  }
}`

	queryResult := QueryResult{ResponseID:"xxxxxx", Language:"zh",
		IntentInfo:  IntentInfo{Name:"Intent Name", Uuid:"xxxxxxxxx", Action:"action.intent.name", Confidence:0.58,
			Parameters: map[string]interface{}{"User Name": "Jax", "Department Name": "Go Dev"},},
		OriginQuery: OriginQuery{Type: "text", Value:"This is the query value"},
	}
	result, err := json.Marshal(&queryResult)
	if err != nil {
		panic(err.Error())
	}
	println(string(result))
}
