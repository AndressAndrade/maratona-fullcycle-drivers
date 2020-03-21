package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type Driver struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Drivers struct {
	Drivers []Driver
}

// Carrega arquivo
func loadDrivers(file string) []byte {

	jsonFile, err := os.Open(file)
	// Se o erro for diferente de nulo
	if err != nil {
		panic(err.Error())
	}

	// Quando for o momento feche o arquivo
	defer jsonFile.Close()

	// Ler tudo do arquivo
	data, err := ioutil.ReadAll(jsonFile)
	// Se o erro for diferente de nulo
	if err != nil {
		panic(err.Error())
	}
	// Retorne os dados
	return data

}

// Função para listar todos os motoristas
func ListDrivers(w http.ResponseWriter, r * http.Request) {
	drivers := loadDrivers("drivers.json")
	w.Write([]byte(drivers))
}

// Função para obter motorista por ID
func GetDriverById(w http.ResponseWriter, r * http.Request) {

	// Obtem parametro da requisicao
	vars := mux.Vars(r)
	// Carrega motoristas
	data := loadDrivers("drivers.json")

	// Inicializa variável
	var drivers Drivers
	// Transforma json em objeto
	json.Unmarshal(data, &drivers)

	for _, v := range drivers.Drivers {
		if v.Uuid == vars["id"] {
			// Transforma objeto em json
			driver, _ := json.Marshal(v)
			// Retorna objeto
			w.Write([]byte(driver))
		}
	}

}

func main() {
	// Inicializa o servidor
	r := mux.NewRouter()
	// Cria o endpoint /drivers e chama função
	r.HandleFunc("/drivers", ListDrivers)
	r.HandleFunc("/drivers/{id}", GetDriverById)
	// Expoe servidor na porta 8081
	http.ListenAndServe(":8081", r)
}