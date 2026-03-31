package normaladapter

import "fmt"

// 手机充电器 稳定输出5v电压
type PhoneCharger interface {
	Output5V()
}

// 华为手机充电器
type HuaWeiCharger struct{}

func NewHuaWeiCharger() *HuaWeiCharger {
	return &HuaWeiCharger{}
}

func (h *HuaWeiCharger) Output5V() {
	fmt.Println("华为手机充电器稳定输出5v电压...")
}

// 小米手机充电器
type XiaoMiCharger struct{}

func NewXiaoMiCharger() *XiaoMiCharger {
	return &XiaoMiCharger{}
}

func (x *XiaoMiCharger) Output5V() {
	fmt.Println("小米手机充电器稳定输出5v电压...")
}

// macbook充电器
type MacBookCharger struct{}

func NewMacBookCharger() *MacBookCharger {
	return &MacBookCharger{}
}

func (m *MacBookCharger) Output28V() {
	fmt.Println("苹果笔记本充电器稳定输出28v电压...")
}

// 手机充电器的适配器
type MacBookChargerAdapter struct {
	core *MacBookCharger
}

func NewMacBookChargerAdapter(m *MacBookCharger) *MacBookChargerAdapter {
	return &MacBookChargerAdapter{
		core: m,
	}
}

func (m *MacBookChargerAdapter) Output5V() {
	m.core.Output28V()
	fmt.Println("适配器将电压调整为稳定输出5v...")
}

// 手机
type Phone interface {
	Charge(phoneCharger PhoneCharger)
}

// 华为手机
type HuaWeiPhone struct{}

func NewHuaweiPhone() Phone {
	return &HuaWeiPhone{}
}

func (h *HuaWeiPhone) Charge(phonecharger PhoneCharger) {
	fmt.Println("华为手机开始充电...")
	phonecharger.Output5V()
}
