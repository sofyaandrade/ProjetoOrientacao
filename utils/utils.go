package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	VOLUME  = "VOLUME"
	PESO    = "PESO"
	UMIDADE = "UMIDADE"
)

func Float32ToString(valor float64) string {
	return strconv.FormatFloat(valor, 'f', 0, 32)
}

func Float64ToString(valor float64) string {
	return strconv.FormatFloat(valor, 'f', 0, 64)
}

func StringToInt(valor string) int {
	n, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func StringToUint(valor string) uint {
	n, err := strconv.ParseUint(valor, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return uint(n)
}

func StringToInt32(valor string) int64 {
	n, err := strconv.ParseInt(valor, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func StringToInt64(valor string) int64 {
	n, err := strconv.ParseInt(valor, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func StringToFloat64(valor string) float64 {
	n, err := strconv.ParseFloat(valor, 10)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func UintToString(valor uint) string {
	return strconv.Itoa(int(valor))
}

func IntToString(valor int) string {
	return strconv.Itoa(int(valor))
}

func ConvertIconeToBytes(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

func JsonToString(body []byte, ponteiro interface{}) (interface{}, error) {
	err := json.Unmarshal([]byte(body), ponteiro)
	if err != nil {
		fmt.Println(err)
	}
	return ponteiro, err
}

func BaseCemParaBaseUm(valor float64) float64 {
	return ArredondarFloat(valor/100, 4)
}

func BaseUmParaBaseCem(valor float64) float64 {
	return ArredondarFloat(valor*100, 2)
}

func ArredondarFloat(valor float64, precisao uint) float64 {
	ratio := math.Pow(10, float64(precisao))
	return math.Round(valor*ratio) / ratio
}

func IsNumeroMaiorQue(numero, outroNumero float64) bool {
	return numero > outroNumero
}

func IsNumeroMenorQue(numero, outroNumero float64) bool {
	return numero < outroNumero
}

func SubstituirZeroPorMaisProximo(m map[int]float32) {
	for chave, valor := range m {
		if valor == 0 {
			var valorMaisProximo float32
			distanciaMaisProxima := float32(math.MaxFloat32)

			for _, outroValor := range m {
				if outroValor != 0 {
					distancia := float32(math.Abs(float64(outroValor - valor)))
					if distancia < distanciaMaisProxima {
						distanciaMaisProxima = distancia
						valorMaisProximo = outroValor
					}
				}
			}

			m[chave] = valorMaisProximo
		}
	}
}
func MilimetroParaCentimetro(valorEmMM float32) float32 {
	return valorEmMM / 100
}

func IsEquipamentoVolume(equipamentoTipo string) bool {
	return strings.Contains(equipamentoTipo, VOLUME)
}

func IsEquipamentoPeso(equipamentoTipo string) bool {
	return strings.Contains(equipamentoTipo, PESO)
}

func IsEquipamentoUmidade(equipamentoTipo string) bool {
	return strings.Contains(equipamentoTipo, UMIDADE)
}

func SolicitarNovoTempoEmMinutos(tempo float64) time.Time {
	conversaoTempo := tempo * 60
	return time.Now().Add(time.Minute * time.Duration(conversaoTempo))
}
