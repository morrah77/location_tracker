package utils

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var errorWrngOrderId = "Wrong Order ID!"
var errorWrngMax = "Wrong max value!"

func GetOrderId(path string, prefix string) (orderId string, err error) {
	orderId = strings.Trim(strings.Replace(path, prefix, ``, 1), `/`);
	if strings.ContainsAny(orderId, `/?&`) {
		return orderId, errors.New(errorWrngOrderId)
	}
	return orderId, err
}

func GetDepth(url *url.URL, key string)  (depth int, err error) {
	ds := url.Query().Get(key)
	if ds == `` {
		return -1, nil
	}
	depth, err = strconv.Atoi(ds);
	if err != nil {
		depth = -1
	}
	return depth, err
}
