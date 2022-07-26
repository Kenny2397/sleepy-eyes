package drawflow

import (
	"context"
	"log"
	"net/http"

	"github.com/Kenny2397/visual-programming/database"
	"github.com/Kenny2397/visual-programming/server"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/go-chi/chi/v5"
)

func GetDrawflowByIdg(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if -> identifier
		ifParam := chi.URLParam(r, "if")

		query := `
    query all($a: string){
      var(func: has(identifier)){
      q as identifier
      }
      query(func: eq(val(q), $a)){
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

		req := &api.Request{
			Query: query,
			Vars:  map[string]string{"$a": ifParam},
		}

		res, err := txn.Do(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		response := res.GetJson()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

		// fmt.Printf("%s\n", res.Json)

	}
}
