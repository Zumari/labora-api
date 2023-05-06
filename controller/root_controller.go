package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/Zumari/labora-api/models"
)

func HandleRootResource(response http.ResponseWriter, request *http.Request) {
	msg := "Soy el recurso raiz"
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(msg))
}

func HandleOtherResource(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("respuesta fallida"))
		}
	}()
	msg := "Soy otro recurso"
	var item *models.Item = nil
	fmt.Println(item.ID) // KA BOOM
	response.Write([]byte(msg))
	response.WriteHeader(http.StatusOK)
}

func ChatgptHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("Endpointer, request from ip %s\n", request.RemoteAddr)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	//fmt.Fprintf(response, text)
	response.WriteHeader(http.StatusOK)
	_, err := response.Write([]byte(inputStr))
	if err != nil {
		fmt.Printf("error while writting bytes to response writer: %+v", err)
	}
}
