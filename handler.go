package main

import (
//	"context"
//	"encoding/json"
//	"fmt"
	"net/http"
	"strconv"
//	"time"

	"github.com/lni/dragonboat/v4"
)

type handler struct {
	nh *dragonboat.NodeHost
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer w.Write([]byte("\n"))
//	var err error
//	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	if r.Method == "GET" {
//		query := Query{
//			Key: r.URL.Path,
//		}
//		res, err := h.nh.SyncRead(ctx, shardID, query)
//		if err != nil {
//			w.WriteHeader(200)
//			w.Write([]byte("TEST HERE"))

			leaderNodeID, _, valid, err := h.nh.GetLeaderID(exampleShardID)
			if !valid || err != nil {
				w.WriteHeader(404)
				w.Write([]byte("Not Found"))
				return
			}

			w.WriteHeader(200)

//cannot use leaderNodeID (variable of type uint64) as int64 value in argument to strconv.FormatInt

//			b := []byte(strconv.FormatInt(num, 10))
			b := []byte(strconv.FormatUint(leaderNodeID, 10))

			w.Write([]byte(b))

			return
//		}
//		if _, ok := res.(Entry); !ok {
//			w.WriteHeader(404)
//			w.Write([]byte("Not Found"))
//			return
//		}
//		b, _ := json.Marshal(res.(Entry))
//		w.WriteHeader(200)
//		w.Write(b)
//	} else if r.Method == "PUT" {
//		var ver int
//		if len(r.FormValue("ver")) > 0 {
//			ver, err = strconv.Atoi(r.FormValue("ver"))
//			if err != nil {
//				w.WriteHeader(400)
//				w.Write([]byte("Version must be uint64"))
//				return
//			}
//		}
//		var entry = Entry{
//			Key: r.URL.Path,
//			Ver: uint64(ver),
//			Val: r.FormValue("val"),
//		}
//		b, err := json.Marshal(entry)
//		if err != nil {
//			w.WriteHeader(400)
//			w.Write([]byte(err.Error()))
//			return
//		}
//		res, err := h.nh.SyncPropose(ctx, h.nh.GetNoOPSession(shardID), b)
//		if err != nil {
//			w.WriteHeader(500)
//			w.Write([]byte(err.Error()))
//			return
//		}
//		if res.Value == ResultCodeFailure {
//			w.WriteHeader(400)
//			w.Write(res.Data)
//			return
//		}
//		if res.Value == ResultCodeVersionMismatch {
//			var result Entry
//			json.Unmarshal(res.Data, &result)
//			w.WriteHeader(409)
//			w.Write([]byte(fmt.Sprintf("Version mismatch (%d != %d)", entry.Ver, result.Ver)))
//			return
//		}
//		w.WriteHeader(200)
//		w.Write(res.Data)
	} else {
		w.WriteHeader(405)
		w.Write([]byte("Method not supported"))
	}
}
