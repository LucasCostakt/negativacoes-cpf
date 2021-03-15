package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)


type Negativacoes struct {
	Data []*Data 
}

type Data struct{
	CompanyDocument string `json:"companyDocument"`
	CompanyName string `json:"companyName"`
	CustomerDocument string `json:"customerDocument"`
	Value float64 `json:"value"`
	Contract string `json:"contract"`
	DebtDate time.Time `json:"debtDate"`
	InclusionDate time.Time `json:"inclusionDate"`
}


func NewNegativacoes(dt []*Data) Storage {
	return &Negativacoes{
		Data: dt,
	}
}

func (ng *Negativacoes) ConvertFileToStruct(namePath string) ([]*Data, error) {


	file, err := os.Open(namePath)
	if err != nil {
		log.Println("fileOpen error in ConvertFileToStruct(), ", err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var datas interface{}

	json.Unmarshal(byteValue, &datas)

	data, err := parseToStruct(ng, datas)
	if err != nil {
		log.Println("fileOpen error in ConvertFileToStruct(), ", err)
		return nil, fmt.Errorf("row.Scan(): %w", err)
	}
	
	return data, nil
} 


func parseToStruct(ng *Negativacoes, datas interface{}) ([]*Data, error){
	v := reflect.ValueOf(datas)

	for i := 0; i < v.Len(); i++ {
		strct := v.Index(i).Interface()
		myMap, _ := strct.(map[string]interface{})

		num := fmt.Sprintf("%v", myMap["value"])
		f, err := strconv.ParseFloat(num, 64)
		if err != nil {
			log.Println("Error convert float parseToStruct(), ", err)
		}
		debtDate, err := time.Parse(time.RFC3339 , fmt.Sprintf("%v", myMap["debtDate"]))	
		inclusionDate, err := time.Parse(time.RFC3339 , fmt.Sprintf("%v", myMap["inclusionDate"]))
		
		dt := &Data{
			CompanyDocument: fmt.Sprintf("%v", myMap["companyDocument"]),
			CompanyName: fmt.Sprintf("%v", myMap["companyName"]),
			CustomerDocument: fmt.Sprintf("%v", myMap["customerDocument"]),
			Value: f,
			Contract: fmt.Sprintf("%v", myMap["contract"]),
			DebtDate: debtDate,
			InclusionDate: inclusionDate,
		}
		
		ng.Data = append(ng.Data,dt)
	}


return ng.Data, nil
}