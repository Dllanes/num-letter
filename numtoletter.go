package numletter

import (
	"fmt"
	"strings"
	"strconv"
)

var unidades = []string{
	"",
	"UN ",
	"DOS ",
	"TRES ",
	"CUATRO ",
	"CINCO ",
	"SEIS ",
	"SIETE ",
	"OCHO ",
	"NUEVE ",
	"DIEZ ",
	"ONCE ",
	"DOCE ",
	"TRECE ",
	"CATORCE ",
	"QUINCE ",
	"DIECISEIS ",
	"DIECISIETE ",
	"DIECIOCHO ",
	"DIECINUEVE ",
	"VEINTE ",
}

var decenas = []string{
    "VEINTI",
    "TREINTA ",
    "CUARENTA ",
    "CINCUENTA ",
    "SESENTA ",
    "SETENTA ",
    "OCHENTA ",
    "NOVENTA ",
}

var centenas = []string{
    "CIENTO ",
    "DOSCIENTOS ",
    "TRESCIENTOS ",
    "CUATROCIENTOS ",
    "QUINIENTOS ",
    "SEISCIENTOS ",
    "SETECIENTOS ",
    "OCHOCIENTOS ",
    "NOVECIENTOS ",
}

func Num2Letter(numero, moneda, centimos string, ForzarCentimos bool) string {

	convertido := ""
	decimales := ""

	div_decimales := strings.Split(numero,".")

	if len(div_decimales) > 1 {
		numero = div_decimales[0]
		decNumberStr := div_decimales[1]
		if len(decNumberStr) == 2 {
			decNumberStrFill := strings.Repeat("0", 7) + decNumberStr
			decCientos := decNumberStrFill[6 : ]
			decimales = ConvertGroup1(decCientos)
		}else{
			decNumberStrFill := strings.Repeat("0", 7) + decNumberStr + strings.Repeat("0", 2)
			decCientos := decNumberStrFill[6 : ]
			decimales = ConvertGroup1(decCientos)
		}
	}else{
		if len(div_decimales) == 1 && ForzarCentimos {
			decimales = "CER0 "
		}
	}

	var valor_convertido string

	if l, _ := strconv.Atoi(numero); l%1000000 != 0 {
		numeroFill := strings.Repeat("0", 9 - len(numero)) + numero

		millones := numeroFill[ : 3]
		miles := numeroFill[3 : 6]
		cientos := numeroFill[6 : ]

		if f, _ := strconv.Atoi(millones); f > 0 {
			if millones == "001" {
				convertido += "UN MILLON "
			}else {
				convertido += fmt.Sprintf("%sMILLONES ", ConvertGroup1(millones) )
			}
		}

		if f, _ := strconv.Atoi(miles); f > 0 {

			convertido += fmt.Sprintf("%sMIL ", ConvertGroup1(miles) )

		}

		if f, _ := strconv.Atoi(cientos); f > 0 {

			convertido += ConvertGroup1(cientos)

		}

		valor_convertido = convertido +  strings.ToUpper(moneda)
	}else{
		if l == 0 {
			valor_convertido = "CERO " +  strings.ToUpper(moneda)
		}else{
			numeroFill := strings.Repeat("0", 9 - len(numero)) + numero
			millones := numeroFill[ : 3]

			if f, _ := strconv.Atoi(millones); f > 0 {
				if millones == "001" {
					convertido += "UN MILLON DE "
				}else {
					convertido += fmt.Sprintf("%sMILLONES DE ", ConvertGroup1(millones) )
				}
			}
			valor_convertido = convertido + strings.ToUpper(moneda)
		}
	}

	if decimales != "" {
		valor_convertido += " CON " + decimales + strings.ToUpper(centimos)
	}

	return valor_convertido

}

func ConvertGroup1( n string ) string {
	final := ""

	if n[0] != '0' {
		if n == "100" {
			return "CIEN "
		}
		final += centenas[n[0]-49]
	}

	k, _ := strconv.Atoi(n[1 : ])

	if k <= 20 {
		final += unidades[k]
	}else{
		if k > 30 && n[2] != '0' {
			final += decenas[n[1] - 50] + "Y " + unidades[n[2] - 48]
		}else{
			final += decenas[n[1] - 50] + unidades[n[2] - 48]
		}
	}

	return final

}
