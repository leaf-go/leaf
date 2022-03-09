package utils
//
//import (
//	"fmt"
//	"github.com/gookit/color"
//)
//
//type output_kind int
//type output_method int
//
//const (
//	output_success output_kind = 1
//	output_warning output_kind = 2
//	output_info    output_kind = 3
//	output_notice  output_kind = 4
//	output_normal  output_kind = 5
//	output_error   output_kind = 6
//
//	output_printf  output_method = 1
//	output_sprintf output_method = 2
//)
//
//var (
//	Output = &DefaultOutput{usePrefix: true}
//)
//
//type DefaultOutput struct {
//	usePrefix bool
//}
//
//func NewOutput(usePrefix bool) *DefaultOutput {
//	return &DefaultOutput{usePrefix: usePrefix}
//}
//
//func (o *DefaultOutput) params(name, message string) (string, string, string) {
//	if !o.usePrefix {
//		return "%s%s", message, ""
//	}
//
//	return "[" + name + "] %s: %s\n", Time.DateTime(), message
//}
//
//func (o *DefaultOutput) Success(message string) {
//	ft, t, msg := o.params("success", message)
//	color.Success.Printf(ft, t, msg)
//}
//
//func (o *DefaultOutput) Warn(message string) {
//	ft, t, msg := o.params("warning", message)
//	color.Warn.Printf(ft, t, msg)
//}
//
//func (o *DefaultOutput) Error(message string) {
//	ft, t, msg := o.params("error", message)
//	color.Error.Printf(ft, t, msg)
//}
//
//func (o *DefaultOutput) Info(message string) {
//	ft, t, msg := o.params("info", message)
//	color.Info.Printf(ft, t, msg)
//}
//
//func (o *DefaultOutput) Notice(message string) {
//	ft, t, msg := o.params("notice", message)
//	color.Yellow.Printf(ft, t, msg)
//}
//
//func (o *DefaultOutput) Normal(message string) {
//	ft, t, msg := o.params("normal", message)
//	color.LightBlue.Printf(ft, t, msg)
//}
//
//func (o *DefaultOutput) Successf(format string, a ...interface{}) {
//	o.Success(fmt.Sprintf(format, a...))
//}
//
//func (o *DefaultOutput) Warnf(format string, a ...interface{}) {
//	o.Warn(fmt.Sprintf(format, a...))
//}
//
//func (o *DefaultOutput) Errorf(format string, a ...interface{}) {
//	o.Error(fmt.Sprintf(format, a...))
//}
//
//func (o *DefaultOutput) Infof(format string, a ...interface{}) {
//	o.Info(fmt.Sprintf(format, a...))
//}
//
//func (o *DefaultOutput) Noticef(format string, a ...interface{}) {
//	o.Notice(fmt.Sprintf(format, a...))
//}
//
//func (o *DefaultOutput) Normalf(format string, a ...interface{}) {
//	o.Normal(fmt.Sprintf(format, a...))
//}
//
//func (o *DefaultOutput) ToSuccessf(format string, a ...interface{}) string {
//	ft, t, msg := o.params("success", fmt.Sprintf(format, a))
//	return color.Success.Sprintf(ft, t, msg)
//}
//
//func (o *DefaultOutput) ToWarnf(format string, a ...interface{}) string {
//	ft, t, msg := o.params("waring", fmt.Sprintf(format, a))
//	return color.Warn.Sprintf(ft, t, msg)
//}
//
//func (o *DefaultOutput) ToErrorf(format string, a ...interface{}) string {
//	ft, t, msg := o.params("error", fmt.Sprintf(format, a))
//	return color.Error.Sprintf(ft, t, msg)
//}
//
//func (o *DefaultOutput) ToInfof(format string, a ...interface{}) string {
//	ft, t, msg := o.params("normal", fmt.Sprintf(format, a))
//	return color.Info.Sprintf(ft, t, msg)
//}
//
//func (o *DefaultOutput) ToNoticef(format string, a ...interface{}) string {
//	ft, t, msg := o.params("notice", fmt.Sprintf(format, a))
//	return color.Yellow.Sprintf(ft, t, msg)
//}
//
//func (o *DefaultOutput) ToNormalf(format string, a ...interface{}) string {
//	ft, t, msg := o.params("normal", fmt.Sprintf(format, a))
//	return color.LightBlue.Sprintf(ft, t, msg)
//}
//
//func (o *DefaultOutput) PanicToWarnf(format string, a ...interface{}) string {
//	panic(o.ToWarnf(format, a...))
//}
//
//func (o *DefaultOutput) PanicToErrorf(format string, a ...interface{}) string {
//	panic(o.ToErrorf(format, a...))
//}
//
//func (o *DefaultOutput) PanicToInfof(format string, a ...interface{}) string {
//	panic(o.ToInfof(format, a...))
//}
//
//func (o *DefaultOutput) PanicToNoticef(format string, a ...interface{}) string {
//	panic(o.ToNoticef(format, a...))
//}
//
//func (o *DefaultOutput) PanicToNormalf(format string, a ...interface{}) string {
//	panic(o.ToNormalf(format, a...))
//}
