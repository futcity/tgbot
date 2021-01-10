///////////////////////////////////////////////////////////////////
//
// Future City Project
//
// Copyright (C) 2020-2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
///////////////////////////////////////////////////////////////////

package utils

import (
	"errors"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// HTTPGetData get data from micro controllers
func HTTPGetData(host string, port int, request string, timeout int, response interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	c := &fasthttp.HostClient{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	statusCode, body, err := c.GetTimeout(nil, fmt.Sprintf("http://%s:%d%s", host, port, request), time.Duration(timeout)*time.Second)

	if err != nil {
		return err
	}
	if statusCode != fasthttp.StatusOK {
		return errors.New("Request status fail")
	}

	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return err
		}
	}

	return nil
}
