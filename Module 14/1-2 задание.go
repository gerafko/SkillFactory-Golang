package main

import (
	"fmt"
	"math"
)

func hashint64(val int64) uint64 {
	return uint64(val % 1000)
}

func hashstr(val string) uint64 {
	valByte := []byte(val)
	var valFloat uint64
	for i := 0; i < len(valByte); i++ {
		valFloat += uint64(math.Pow(10, float64(i))) * uint64(valByte[i])
		//fmt.Println(uint64(valByte[i]))
	}
	return valFloat % 1000
}

type hashmap struct {
	value []string
}

func (h *hashmap) Set(key, val string) {
	h.value[int(hashstr(key))] = val
}

func (h *hashmap) Get(key string) (value string, ok bool) {
	if h.value[int(hashstr(key))] != "" {
		return h.value[int(hashstr(key))], true
	}
	return "", false
}

func (h *hashmap) Delete(key string) {
	h.value[int(hashstr(key))] = ""
}

func main() {
	fmt.Println(hashint64(771))
	fmt.Println(hashstr("abc"))
	fmt.Println(hashstr("bac"))

	fmt.Println("Реализация типа хэш-мапы")
	h := &hashmap{make([]string, 1000)}
	h.Set("cat", "кот")
	h.Set("dog", "собака")
	cat, catOk := h.Get("cat")
	fmt.Println(cat, catOk)
	h.Delete("dog")
	dog, dogOk := h.Get("dog")
	fmt.Println(dog, dogOk)
}
