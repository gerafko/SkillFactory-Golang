package electronic

import "fmt"

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
	Phone
}

type Smartphone interface {
	OS() string
	Phone
}

type applePhone struct {
	brand     string
	model     string
	phoneType string
	os        string
}

type androidPhone struct {
	brand     string
	model     string
	phoneType string
	os        string
}

type radioPhone struct {
	brand        string
	model        string
	phoneType    string
	buttonsCount int
}

func (a applePhone) Brand() string {
	return fmt.Sprint(a.brand)
}
func (a applePhone) Model() string {
	return fmt.Sprint(a.model)
}
func (a applePhone) Type() string {
	return fmt.Sprint(a.phoneType)
}
func (a applePhone) OS() string {
	return fmt.Sprint(a.os)
}
func NewApplePhone(model string) applePhone {
	return applePhone{"apple", model, "smartphone", "ios"}
}

func (a androidPhone) Brand() string {
	return fmt.Sprint(a.brand)
}
func (a androidPhone) Model() string {
	return fmt.Sprint(a.model)
}
func (a androidPhone) Type() string {
	return fmt.Sprint(a.phoneType)
}
func (a androidPhone) OS() string {
	return fmt.Sprint(a.os)
}
func NewAndroidPhone(brand, model string) androidPhone {
	return androidPhone{brand, model, "smartphone", "android"}
}

func (a radioPhone) Brand() string {
	return fmt.Sprint(a.brand)
}
func (a radioPhone) Model() string {
	return fmt.Sprint(a.model)
}
func (a radioPhone) Type() string {
	return fmt.Sprint(a.phoneType)
}
func (a radioPhone) ButtonsCount() string {
	return fmt.Sprint(a.buttonsCount)
}
func NewRadioPhone(brand, model string, buttonsCount int) radioPhone {
	return radioPhone{brand, model, "station", buttonsCount}
}
