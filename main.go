package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/nimaposhtiban/go-ciphers/randomness"
	"github.com/nimaposhtiban/go-ciphers/symmetric"
)

func ceasar(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	action := values.Get("action")
	text := values.Get("text")
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(200)
	var ceasar symmetric.Caesar = symmetric.Caesar(text)
	if action == "cipher" {
		text, _ = ceasar.Encrypt()
	}
	if action == "decipher" {
		text, _ = ceasar.Decrypt()
	}
	fmt.Println(text + "\n" + action)

	fmt.Fprintln(w, string(text))
}

func vigenere(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	text := values.Get("text")
	key := values.Get("key")
	action := values.Get("action")
	var vigenere symmetric.Vigenere = symmetric.Vigenere(text)
	if action == "cipher" {
		text, _ = vigenere.Encrypt(key)
	}
	if action == "decipher" {
		text, _ = vigenere.Decrypt(key)
	}
	fmt.Println("Vigener !!")
	fmt.Fprintln(w, string(text))
}

func otp(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	text := values.Get("text")
	key := values.Get("key")
	action := values.Get("action")
	var otp_ symmetric.OneTimePad = symmetric.OneTimePad{}
	otp_.Key = key
	if action == "cipher" {
		otp_.PlainText = text
		text = otp_.Encrypt()
	}
	if action == "decipher" {
		otp_.Hash = text
		text = otp_.Decrypt()
	}
	fmt.Println("Vigener !!")
	fmt.Fprintln(w, string(text))
}

func NonUniformEntropy(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	text := values.Get("text")
	numbers := strings.Split(text, ",")
	nums := make([]float64, 0)
	for _, num := range numbers {
		n, _ := strconv.ParseFloat(num, 64)
		nums = append(nums, n)
	}
	result := randomness.CalculateNonUniformEntropy(nums...)
	fmt.Fprintln(w, result)
}

func UniformEntropy(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	text := values.Get("text")
	probabilityStr := values.Get("key")
	probability, _ := strconv.ParseFloat(probabilityStr, 64)
	occurrence, _ := strconv.ParseInt(text, 10, 64)
	result := randomness.CalculateUniformEntropy(probability, int(occurrence))
	fmt.Fprintln(w, result)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	page, _ := os.ReadFile("index.html")
	w.Write(page)
}

func main() {
	http.HandleFunc("/uniformentropy", UniformEntropy)
	http.HandleFunc("/nonuniformentropy", NonUniformEntropy)
	http.HandleFunc("/otp", otp)
	http.HandleFunc("/vigenere", vigenere)
	http.HandleFunc("/ceasar", ceasar)
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":12345", nil)
}
