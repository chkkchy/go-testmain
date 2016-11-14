package hoge

import (
	"log"
	"testing"
)

//func TestMain(m *testing.M) {
//	log.Println("This gets run BEFORE any tests get run!")
//	//exitVal := m.Run()
//	log.Println("This gets run AFTER any tests get run!")
//	//os.Exit(exitVal)
//}

func TestOne(t *testing.T) {
	Hogefunc1()
	log.Println("TestOne running")
}

func TestTwo(t *testing.T) {
	Hogefunc2()
	log.Println("TestTwo running")
	//t.Fail()
}
