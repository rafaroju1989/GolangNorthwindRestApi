/*package main

import (
	"GoPruebaDep/database"
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection *sql.DB

type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	//NOTA: Todo request, tiene un respond
	r.Get("/products", AllProducts)

	http.ListenAndServe(":3000", r)

}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description, '') FROM products`
	results, err := databaseConnection.Query(sql)
	println(results)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.Description)

		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}*/

package main

import (
	"GolangNorthwindRestApi/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func catch(err error) {
	if err != nil {
		println("hay un error")
		panic(err)
	}
}

var databaseConnection *sql.DB

func main() {
	databaseConnection = database.InitDB()
	// Logica

	defer databaseConnection.Close()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products", AllProductos)
	r.Post("/products", CreateProducto)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProducto)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	http.ListenAndServe(":3000", r)

}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)

	_, err = query.Exec(id)
	catch(err)
	query.Close()

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "succesfully deleted"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product

	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("Update products set product_code=?, description=? where id=?")
	catch(err)

	_, err = query.Exec(producto.Product_Code, producto.Description, id)
	catch(err)
	defer query.Close()

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "succesfully edited"})
}

func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product

	json.NewDecoder(r.Body).Decode(&producto)
	query, err := databaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)

	_, err = query.Exec(producto.Product_Code, producto.Description)
	catch(err)
	defer query.Close()

	responseWithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully create"})
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'') FROM products`

	results, err := databaseConnection.Query(sql)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	responseWithJSON(w, http.StatusOK, products)
}
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
