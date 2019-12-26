package classify

import (
	"github.com/BalusChen/DataMiningFinalProject/core"
)

const (
	KindsOfWifiOptions      = 2
	KindsOfBluetoothOptions = 2
	KindsOfDualSIMOptions   = 2
	KindsOfPriceOptions     = 4
)

// Wifi 根据是否具有 WiFi 来进行分类
func Wifi(infos []*core.ElecCarInfo) ([KindsOfPriceOptions][]uint64, int, string) {
	// 由于 price 只有 0，1，2，3 这四种，所以直接用 array 替代 map
	var data [KindsOfPriceOptions][]uint64

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

func Bluetooth(infos []*core.ElecCarInfo) ([KindsOfPriceOptions][]uint64, int, string) {
	var data [KindsOfPriceOptions][]uint64

	for i := 0; i < KindsOfPriceOptions; i++ {
		data[i] = make([]uint64, KindsOfBluetoothOptions)
	}

	for _, v := range infos {
		if v.HasBlueotooth {
			data[v.Price][1]++
		} else {
			data[v.Price][0]++
		}
	}

	return data, KindsOfBluetoothOptions, "hasBluetooth"
}

func DualSIM(infos []*core.ElecCarInfo) ([KindsOfPriceOptions][]uint64, int, string) {
	var data [KindsOfPriceOptions][]uint64

	for i := 0; i < KindsOfPriceOptions; i++ {
		data[i] = make([]uint64, KindsOfDualSIMOptions)
	}

	for _, v := range infos {
		if v.HasDualSIM {
			data[v.Price][1]++
		} else {
			data[v.Price][0]++
		}
	}

	return data, KindsOfDualSIMOptions, "hasDualSIM"
}
