package drawflow

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kenny2397/visual-programming/database"
	"github.com/Kenny2397/visual-programming/server"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func InsertDrawflow(s server.Server) http.HandlerFunc {

	// p := Person{
	// 	Uid:   "_:pedro",
	// 	Name:  "Pedro",
	// 	DType: []string{"Person"},
	// }

	// op := &api.Operation{}
	// op.Schema = `
	// 	name: string @index(exact) .
	// 	age: int .
	// 	married: bool.

	// 	type Person {
	// 		name: string
	// 		age: int
	// 		married: bool
	// 	}
	// `

	// pb, err := json.Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(pb, err)
	// dg := database.NewClient()

	// txn := dg.NewTxn()
	// defer txn.Discard(context.Background())

	// mu := &api.Mutation{
	// 	SetJson:   pb,
	// 	CommitNow: true,
	// }

	// res, err := txn.Mutate(context.Background(), mu)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	return func(w http.ResponseWriter, r *http.Request) {
		// type Ndata struct {
		// 	Numbervalue string `json:"numbervalue,omitempty"`
		// }

		// type prueba struct {
		// 	Name string `json:"name"`
		// 	Data Ndata  `json:"data"`
		// }
		// var p prueba

		// fmt.Println(r.Body)

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		// reqBody ya esta en []Byte
		// rBody := string(reqBody)
		// fmt.Println(rBody)

		// json.Unmarshal()
		// json.Unmarshal(reqBody, &p)

		// fmt.Println(p.Name)
		// fmt.Println(p.Data)

		// json.Marshall
		// jsonData, err := json.Marshal(p)
		// if err != nil {
		// 	fmt.Fprintln(w, jsonData, err)
		// }
		// fmt.Println(string(jsonData))

		// dGRAPH
		dg := database.NewClient()

		txn := dg.NewTxn()
		defer txn.Discard(context.Background())

		mu := &api.Mutation{
			SetJson:   reqBody,
			CommitNow: true,
		}

		res, err := txn.Mutate(context.Background(), mu)
		if err != nil {
			log.Fatal(err)
		}

		// response := res.GetJson()
		w.Header().Set("Content-Type", "application/json")
		// status 200 ok!
		w.WriteHeader(http.StatusOK)
		// fmt.Println(dg)
		json.NewEncoder(w).Encode(res)
		// w.Write(response)
	}

}
