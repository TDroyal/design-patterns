package normaladapter

import "testing"

func TestAdapter(t *testing.T) {
	huaweiPhone := NewHuaweiPhone()

	// 使用华为手机充电器进行充电
	huaweiCharger := NewHuaWeiCharger()
	huaweiPhone.Charge(huaweiCharger)

	// 使用适配器转换后的 macbook 充电器进行充电
	macBookCharger := NewMacBookCharger()
	macBookChargerAdapter := NewMacBookChargerAdapter(macBookCharger)
	huaweiPhone.Charge(macBookChargerAdapter)
}

/*
=== RUN   TestAdapter
华为手机开始充电...
华为手机充电器稳定输出5v电压...
华为手机开始充电...
苹果笔记本充电器稳定输出28v电压...
适配器将电压调整为稳定输出5v...
--- PASS: TestAdapter (0.00s)
PASS
*/
