package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estrutura para resposta do ViaCEP
type ViaCEP struct {
	Cep       string `json:"cep"`
	Localidade string `json:"localidade"`
	Estado    string `json:"uf"`
}

// Estrutura para resposta de temperatura
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"` // Temperatura em Kelvin
	} `json:"main"`
}

// Estrutura combinada para resposta final
type FullResponse struct {
Cep        string  `json:"cep"`
Localidade string  `json:"localidade"`
Estado     string  `json:"estado"`
TempC      float64 `json:"temperatura_celsius"`
TempF      float64 `json:"temperatura_fahrenheit"`
TempK      float64 `json:"temperatura_kelvin"`
}

func main() {
	http.HandleFunc("/cep", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "Parâmetro 'cep' é obrigatório", http.StatusBadRequest)
		return
	}

	// Chamada para ViaCEP
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil || res.StatusCode != http.StatusOK {
		http.Error(w, "Erro ao buscar informações do CEP", http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	var c ViaCEP
	if err := json.NewDecoder(res.Body).Decode(&c); err != nil {
		http.Error(w, "Erro ao decodificar resposta do ViaCEP", http.StatusInternalServerError)
		return
	}

	// Chamada para temperatura usando a cidade retornada pelo ViaCEP
	weather, err := getTemperature(c.Localidade)
	if err != nil {
		http.Error(w, "Erro ao buscar temperatura: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tempK := weather.Main.Temp
	tempC := tempK - 273
	tempF := tempC * 1.8 + 32

	// Resposta combinada
	response := FullResponse{
		Cep:        c.Cep,
		Localidade: c.Localidade,
		Estado:     c.Estado,
		TempC:      round(tempC, 0),
		TempF:      round(tempF, 0),
		TempK:      round(tempK, 0),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getTemperature(city string) (*WeatherResponse, error) {
	apiKey := "f8acd7aad5a44b58fafb9ae2464c8ecf"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,br&appid=%s", city, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API OpenWeather retornou status %d", res.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(res.Body).Decode(&weather); err != nil {
		return nil, err
	}
	return &weather, nil
}

func round(val float64, precision int) float64 {
	p := float64(1)
	for i := 0; i < precision; i++ {
		p *= 10
	}
	return float64(int(val*p+0.5)) / p
}