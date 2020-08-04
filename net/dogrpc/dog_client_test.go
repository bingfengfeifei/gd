/**
 * Copyright 2018 godog Author. All Rights Reserved.
 * Author: Chuck1024
 */

package dogrpc_test

import (
	"encoding/json"
	"github.com/chuck1024/godog"
	"github.com/chuck1024/godog/utls/network"
	"testing"
	"time"
)

func TestDogClient(t *testing.T) {
	d := godog.Default()
	c := d.NewRpcClient(time.Duration(500*time.Millisecond), 0)
	c.AddAddr(network.GetLocalIP() + ":10241")

	b := &struct {
		Data string
	}{
		Data: "How are you?",
	}
	body, _ := json.Marshal(b)

	// use godog protocol
	code, rsp, err := c.DogInvoke(1024, body)
	if err != nil {
		t.Logf("Error when sending request to server: %s", err)
	}

	t.Logf("code=%d, resp=%s", code, string(rsp))
}
