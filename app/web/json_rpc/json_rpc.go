// Copyright (c) 2025 The bel2 developers

package jsonrpc

import "github.com/BeL2Labs/Arbiter_Signer/utility/events"

// resp :
//
//	{
//	    "code": 0,
//	    "message": "",
//	    "data": {
//	        "events": [
//	            {
//	                "EventName": "ArbitrationRequested",
//	                "Block": 21210570,
//	                "EventID": "0x00ace622d30c58a7d709754546e232c42c93d45cc399e0bfcc58a5b481526a8e",
//	                "QueryID": "0xe53721cc731dac8bb14a02cdfd264033a2805380f0e6af5a53f87971f7b525f7",
//	                "ArbitratorAddress": "0x0262aB0ED65373cC855C34529fDdeAa0e686D913",
//	                "DappAddress": "0x0000000000000000000000000262ab0ed65373cc855c34529fddeaa0e686d913",
//	                "Status": "failed"
//	            }
//	        ]
//	    }
//	}

// Error resp:
//
//	{
//	    "code": 58,
//	    "message": "Not Implemented",
//	    "data": null
//	}

type EventInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Events []events.EventInfo `json:"Events"`
	} `json:"data"`
}
