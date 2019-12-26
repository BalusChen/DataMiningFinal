package classify

import (
	"github.com/BalusChen/DataMiningFinalProject/core"
)


const (
	KindsOfWifiOptions  = 2
	KindsOfPriceOptions = 4
)

// Wifi 根据是否具有 WiFi 来进行分类
func Wifi(infos []*core.ElecCarInfo) ([4][]uint64, int, string) {
	// 由于 price 只有 0，1，2，3 这四种，所以直接用 array 替代 map
	var data [4][]uint64

	for i := 0; i < KindsOfPriceOptions; i++ {
		data[i] = make([]uint64, KindsOfWifiOptions)
	}

	for _, v := range infos {
		if v.HasWifi {
			data[v.Price][1]++
		} else {
			data[v.Price][0]++
		}
	}

	return data, KindsOfWifiOptions, "hasWifi"
}
