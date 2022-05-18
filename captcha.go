package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"syscall/js"

	captcha "github.com/dchest/captcha"
)

var currentCaptchaId string

func generateCaptcha(this js.Value, inputs []js.Value) interface{} {
	currentCaptchaId = captcha.New()
	img_buf := bytes.NewBufferString("")
	captcha.WriteImage(img_buf, currentCaptchaId, 192, 96)

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(img_buf.Bytes())
}

func verifyCaptcha(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("Checking captcha:", currentCaptchaId, inputs[0].String())

	return captcha.VerifyString(currentCaptchaId, inputs[0].String())
}

func main() {
	fmt.Println("WebAssembly starting...")

	js.Global().Set("generateCaptcha", js.FuncOf(generateCaptcha))
	js.Global().Set("verifyCaptcha", js.FuncOf(verifyCaptcha))

	<-make(chan bool) // Waiting...
}
