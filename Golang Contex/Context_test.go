package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	result := Print()
	assert.NotEqual(t, "hellos", result)
	assert.Equal(t, "hello", result)
}

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("Background context:", background)

	todo := context.TODO()
	fmt.Println("TODO context:", todo)

}

func TestCreateChildContext(t *testing.T) {

	// ------------------------------- !IMPORTANT ------------------------------
	// ctxA adalah parrent dari child => ctxB,ctxC,ctxD,ctxE,ctxF,ctxG
	// context bersifat immutabel jadi artinya dia gabisa di ubah
	// kalau kita maksa buat di ubah maka sebenarnyadia passbyvalue atau membuat context yang baru
	// ctx child dapat mewarisi data dari parrentnya tetapi parrentnya tidak dapat mengambil data dari childnya
	// ctxB dan ctxC gabisa berkomunikasi artinya meski 1 parrent mereka tidab bisa mengambil data dari parrent sebelahnya

	// ini adalah entity dari sebuah context, ctxA
	ctxA := context.Background()

	// kita buat sebuah chid baru yang mengambil entity atau mewarisi data dari parrentya
	// jadi ctxB dan ctxC akan menjadi child dari ctxA
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	// kemudian kita memuat chid baru dalam sub child dari ctxB yang memiliki parrent ctxA
	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")

	// ini juga sama kita membuat child dalam sub cub child dari ctcC yang memiliki parrent ctxA
	ctxF := context.WithValue(ctxC, "f", "F")
	ctxG := context.WithValue(ctxC, "g", "G")

	// ini adalah contoh hirarki yang akan terjadi
	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)
	fmt.Println(ctxG)

	// kita coba lakukan test dengan mengambil key dari sebuah context yang sudah di buat

	// contexD kita nanya yang key-nya => d :: ad atau ngga => result -> ada
	// karena key => d tersimpan secara langsung ke contex D // merupakan key dari contexD
	fmt.Println("value =>", ctxD.Value("d"))

	// contexD kita nanya  key => b :: ada atau ngga => result -> ada
	// karena  contexD merupakan pewarisan dari contexB, atau child dari parrent contexB
	fmt.Println("value =>", ctxD.Value("b"))

	// contexD kita nanya  dari keluarga sebelah key => f :: ada atau ngga => result -> gada
	// karena  contexD walau memiliki parrent yang sama mereka tidak bisa berkomunikasi
	fmt.Println("value =>", ctxD.Value("f"))

	//contexB yang merupakan parrent dari contexD kita coba tanyain key => d :: ada atau ngga => result -> gada
	// karena walau child dapat mewarisi data dari parrentnya tetapi parrentnya gakan bisa ngakses data dari childnya
	fmt.Println("value =>", ctxB.Value("d"))
}
