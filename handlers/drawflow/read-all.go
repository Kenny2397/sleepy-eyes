package drawflow

import (
	"context"
	"log"
	"net/http"

	"github.com/Kenny2397/visual-programming/database"
	"github.com/Kenny2397/visual-programming/server"
)

func GetAllDrawflows(s server.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Query drawflow data
		const query = `{
			getAll(func: has(data))@filter(has(identifier)){
				uid
				identifier
				data{
					id
					name
					data
					class
					html
					typenode
					inputs{
						input_1{
							connections{
								node
								input
							}          
						}        
					}
					outputs{
						output_1{
							connections{
								node
								input
							}          
						}
					}			
					pos_x
					pos_y
				}
			}
		}
		`
		// dGRAPH
		dg := database.NewClient()
		txn := dg.NewTxn()
		defer txn.Discard(context.Background())

		res, err := txn.Query(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
		response := res.GetJson()
		// response, err := json.Marshal(response)
		// if err != nil {
		// 	fmt.Fprintln(w, err)
		// }
		// var r := string(response)

		w.Header().Set("Content-Type", "application/json")
		// status 200 ok!
		w.WriteHeader(http.StatusOK)
		// fmt.Println(dg)
		// json.NewEncoder(w).Encode(response)
		// fmt.Fprintln(w, string(response))
		w.Write(response)

	}
}
