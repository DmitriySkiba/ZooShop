package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Goods struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	Category    string  `json:"category"`
	Subcategory string  `json:"subcategory"`
	Brand       string  `json:"brand"`
	Vcode       int     `json:"vcode"`
	Status      string  `json:"status"`
}

type Delete_good struct {
	Status string `json:"status"`
	Vcode  int    `json:"vcode"`
}

type BasketAction struct {
	Vcode    int `json:"vcode"`
	IdPerson int `json:"id"`
}

type Vcode_good struct {
	Vcode int `json:"vcode"`
}

type Id_person struct {
	Id int `json:"id"`
}

type Lenta struct {
	Sort    string `json:"sort"`
	Column  string `json:"column"`
	Meaning string `json:"meaning"`
}

const (
	host     = "bubble.db.elephantsql.com"
	port     = 5432
	user     = "wzpwshep"
	password = "Lk6rq_iCBbYLk-LxtTI8uIpM7Myee4Rt"
	dbname   = "wzpwshep"
)

// Подключение к базе=====================================================================================================================================
type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := sql.Open("postgres", connStr)

	return &Database{
		db: db,
	}
}

func (d *Database) Close() {
	d.db.Close()
}

// Новый товар ==========================================================================================================================================================
func (db *Database) register_new_goods(w http.ResponseWriter, r *http.Request) {
	var newGoods Goods
	err1 := json.NewDecoder(r.Body).Decode(&newGoods)

	if newGoods.Status != "admin" {
		if newGoods.Status == "user" {
			fmt.Fprintf(w, "User have not got enough laws")
		} else {
			fmt.Fprintf(w, "Unknown status")
		}
	} else {
		if newGoods.Name == "" || newGoods.Price == 0 || newGoods.Brand == "" || err1 != nil {
			fmt.Fprintf(w, "Some fields are null")
		} else {
			//Рандомное число для артикула
			rand.Seed(time.Now().UnixNano())
			min := 1000000000
			max := 9999999999
			vcode_random := rand.Intn(max-min) + min
			fmt.Println(vcode_random)

			sqlStatement := `
				INSERT INTO goods (name, price, description, photo, category, subcategory, brand, vcode)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			db.db.Query(sqlStatement, newGoods.Name, newGoods.Price, newGoods.Description, newGoods.Photo, newGoods.Category, newGoods.Subcategory, newGoods.Brand, vcode_random)

			fmt.Fprintf(w, "New goods added successfully")
		}
	}
}

// Удалить товар ========================================================================================================================================================
func (db *Database) delete_good(w http.ResponseWriter, r *http.Request) {
	var trash Delete_good
	err := json.NewDecoder(r.Body).Decode(&trash)

	if trash.Status != "admin" {
		if trash.Status == "user" {
			fmt.Fprintf(w, "User have not enough laws")
		} else {
			fmt.Fprintf(w, "Unknown status of person")
		}
	} else {
		if err == nil {
			sqlStatement := `DELETE FROM goods WHERE vcode = $1`
			db.db.Exec(sqlStatement, trash.Vcode)
			fmt.Fprintf(w, "Good was deleted")
		}
	}
}

// Редактирование товара ============================================================================================================================================
func (db *Database) update_good(w http.ResponseWriter, r *http.Request) {

	var good_update Goods
	err1 := json.NewDecoder(r.Body).Decode(&good_update)

	if good_update.Status != "admin" {
		fmt.Fprintf(w, "User have not enough laws")
	} else {
		if good_update.Name == "" || good_update.Brand == "" || good_update.Price == 0 {
			fmt.Fprintf(w, "Some fields filled incorrect")
		} else {
			sqlStatement := `
				UPDATE goods
				SET name = $1, price = $2, description = $3, photo = $4, category = $5, subcategory = $6, brand = $7
				WHERE vcode = $8`

			res, err2 := db.db.Exec(sqlStatement, good_update.Name, good_update.Price, good_update.Description, good_update.Photo, good_update.Category, good_update.Subcategory, good_update.Brand, good_update.Vcode)

			rowsAffected, err3 := res.RowsAffected()

			if rowsAffected == 0 || err1 != nil || err2 != nil || err3 != nil {
				fmt.Fprintf(w, "No rows were updated")
				return
			}
			fmt.Fprintf(w, "Goods was updated successfully")
		}
	}
}

// Выборка товара ================================================================================================================================================================================
func (db *Database) select_good(w http.ResponseWriter, r *http.Request) {
	var good_select Vcode_good
	var jresult string

	err := json.NewDecoder(r.Body).Decode(&good_select)
	if err == nil {
		sqlStatement := ` SELECT json_build_object(
			'name', name
		  , 'price', price 
		  , 'description', description
		  , 'photo', photo
		  , 'category', category
		  , 'subcategory', subcategory
		  , 'brand' , brand
		  , 'vcode' , vcode
		  ) as jresult
	 	FROM goods
		WHERE vcode = $1`

		rows, _ := db.db.Query(sqlStatement, good_select.Vcode)

		defer rows.Close()

		for rows.Next() {
			rows.Scan(&jresult)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jresult)
	}
}

// Добавить в корзину ==============================================================================================================================================================================
func (db *Database) add2basket(w http.ResponseWriter, r *http.Request) {
	var add_good BasketAction
	err := json.NewDecoder(r.Body).Decode(&add_good)

	if err == nil {
		sqlStatement := `
			INSERT INTO basket (id, vcode)
			VALUES ($1, $2)`
		db.db.Query(sqlStatement, add_good.IdPerson, add_good.Vcode)

		fmt.Fprintf(w, "New goods added successfully into basket")
	}
}

// Удалить из корзины ============================================================================================================================================================================
func (db *Database) delete_in_basket(w http.ResponseWriter, r *http.Request) {
	var delete_good BasketAction
	err := json.NewDecoder(r.Body).Decode(&delete_good)

	if err == nil {
		sqlStatement := `
		DELETE FROM basket WHERE vcode = $1 AND id = $2 AND purchase IS NULL`

		db.db.Exec(sqlStatement, delete_good.Vcode, delete_good.IdPerson)
		fmt.Fprintf(w, "Good was deleted from basket successfully")
	}
}

// Корзина пользователя ===========================================================================================================================================================
func (db *Database) personal_basket(w http.ResponseWriter, r *http.Request) {
	var request_id Id_person
	var basket_view *string
	err := json.NewDecoder(r.Body).Decode(&request_id)

	fmt.Print(request_id.Id)

	if err == nil {
		sqlStatement := `
		select json_agg(
			json_build_object(
			  'name', name
			, 'price', price
			, 'photo', photo
			, 'vcode', goods.vcode
			)
		) as basket_view
		 from basket, goods
		 where basket.vcode = goods.vcode
		 and basket.purchase is null
		 and basket.id = $1
		 ;`

		rows, _ := db.db.Query(sqlStatement, request_id.Id)

		defer rows.Close()

		for rows.Next() {
			rows.Scan(&basket_view)
		}

		// Check if the basket is empty
		if basket_view == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Basket is empty")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(basket_view)

	}
}

// Купить товар ====================================================================================================================================================================
func (db *Database) buy(w http.ResponseWriter, r *http.Request) {
	var buy_object BasketAction
	err := json.NewDecoder(r.Body).Decode(&buy_object)

	if err == nil {
		sqlStatement := `
		UPDATE basket
		SET purchase = CURRENT_TIMESTAMP
		WHERE id = $1 AND vcode = $2
		`
		db.db.Exec(sqlStatement, buy_object.IdPerson, buy_object.Vcode)

		fmt.Fprintf(w, "Good was bought")
	}
}

// История покупок ===================================================================================================================================================================
func (db *Database) story_purchased(w http.ResponseWriter, r *http.Request) {
	var id_bought Id_person
	var story_bought *string
	err := json.NewDecoder(r.Body).Decode(&id_bought)

	fmt.Print(id_bought.Id)

	if err == nil {
		sqlStatement := `
			select json_agg(
				json_build_object(
		 		 'name', name
				, 'price', price
				, 'photo', photo
				, 'vcode', goods.vcode
				)
			) as story_bought
 			from basket, goods
 			where basket.vcode = goods.vcode
 			and basket.purchase is not null
 			and basket.id = $1
 			;`

		rows, _ := db.db.Query(sqlStatement, id_bought.Id)

		defer rows.Close()

		for rows.Next() {
			rows.Scan(&story_bought)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(story_bought)
	}
}

// Лента общая ====================================================================================================================================================================
func (db *Database) lenta(w http.ResponseWriter, r *http.Request) {
	var all_object Lenta
	var lenta *string
	err := json.NewDecoder(r.Body).Decode(&all_object)

	//Если нет фильтра - выводим по времени добавления товара в базу -> Обычная лента
	if all_object.Column == "" && all_object.Meaning == "" && err == nil {
		if all_object.Sort == "date" {
			sqlStatement := `
			SELECT json_agg(
     			json_build_object(
			      'name', name
      			, 'price', price
				, 'category', category
				, 'subcategory', subcategory 
	    		, 'photo', photo
			    , 'vcode', goods.vcode
    			)) as lenta
			FROM goods`
			rows, _ := db.db.Query(sqlStatement)

			defer rows.Close()

			for rows.Next() {
				rows.Scan(&lenta)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(lenta)
		} else {
			fmt.Fprintf(w, "It can be type of sort in common input")
		}
	}

	//Если есть типы фильтров c подзначением, но нет порядка - сортируем по колонкам и выводим по времени добавления товара в базу
	if (all_object.Column == "category" || all_object.Column == "subcategory") && all_object.Sort == "date" && err == nil {
		if all_object.Meaning != "" {
			if all_object.Column == "category" {
				sqlStatement := `
				SELECT json_agg(
					json_build_object(
			  			'name', name
						, 'price', price
						, 'category', category
						, 'subcategory', subcategory 
						, 'photo', photo
						, 'vcode', goods.vcode
				)) as lenta
				FROM goods
				WHERE category = $1`

				rows, _ := db.db.Query(sqlStatement, all_object.Meaning)

				defer rows.Close()

				for rows.Next() {
					rows.Scan(&lenta)
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(lenta)
			} else {
				sqlStatement := `
				SELECT json_agg(
					json_build_object(
						'name', name
						, 'price', price
						, 'category', category
						, 'subcategory', subcategory 
						, 'photo', photo
						, 'vcode', goods.vcode
				)) as lenta
				FROM goods
				WHERE subcategory = $1`

				rows, _ := db.db.Query(sqlStatement, all_object.Meaning)

				defer rows.Close()

				for rows.Next() {
					rows.Scan(&lenta)
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(lenta)
			}
		} else {
			fmt.Fprintf(w, "Incorrect struct of sorting json")
		}

	}

	//Если есть и фильтр колонок, и подзначение, и порядок - сортируем по колонкам и выводим в соответствующем порядке (ASC. DESC)
	if (all_object.Column == "category" || all_object.Column == "subcategory") && all_object.Sort != "date" && err == nil {
		if all_object.Meaning != "" {
			if all_object.Column == "category" {
				if all_object.Sort == "DESC" {
					sqlStatement := `
					SELECT json_agg(
						json_build_object(
							  'name', name
							, 'price', price
							, 'category', category
							, 'subcategory', subcategory 
							, 'photo', photo
							, 'vcode', goods.vcode
					)) as lenta
					FROM goods
					WHERE category = $2
					GROUP BY name
					ORDER BY name DESC`

					rows, _ := db.db.Query(sqlStatement, all_object.Column, all_object.Meaning)

					defer rows.Close()

					for rows.Next() {
						rows.Scan(&lenta)
					}

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(lenta)

				} else {
					if all_object.Sort == "ASC" {
						sqlStatement := `
							SELECT json_agg(
								json_build_object(
									  'name', name
									, 'price', price
									, 'category', category
									, 'subcategory', subcategory 
									, 'photo', photo
									, 'vcode', goods.vcode
							)) as lenta
							FROM goods
							WHERE category = $2
							GROUP BY name
							ORDER BY name ASC`

						rows, _ := db.db.Query(sqlStatement, all_object.Column, all_object.Meaning)

						defer rows.Close()

						for rows.Next() {
							rows.Scan(&lenta)
						}
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(lenta)

					} else {
						fmt.Fprintf(w, "Incorrect type of sort")
					}
				}
			} else {
				if all_object.Sort == "DESC" {
					sqlStatement := `
					SELECT json_agg(
						json_build_object(
							  'name', name
							, 'price', price
							, 'category', category
							, 'subcategory', subcategory 
							, 'photo', photo
							, 'vcode', goods.vcode
					)) as lenta
					FROM goods
					WHERE subcategory = $2
					GROUP BY name
					ORDER BY name DESC`

					rows, _ := db.db.Query(sqlStatement, all_object.Column, all_object.Meaning)

					defer rows.Close()

					for rows.Next() {
						rows.Scan(&lenta)
					}

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(lenta)

				} else {
					if all_object.Sort == "ASC" {
						sqlStatement := `
							SELECT json_agg(
								json_build_object(
									  'name', name
									, 'price', price
									, 'category', category
									, 'subcategory', subcategory 
									, 'photo', photo
									, 'vcode', goods.vcode
							)) as lenta
							FROM goods
							WHERE subcategory = $2
							GROUP BY name
							ORDER BY name ASC`

						rows, _ := db.db.Query(sqlStatement, all_object.Column, all_object.Meaning)

						defer rows.Close()

						for rows.Next() {
							rows.Scan(&lenta)
						}
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(lenta)

					} else {
						fmt.Fprintf(w, "Incorrect type of sort")
					}
				}
			}

		} else {
			fmt.Fprintf(w, "Incorrect struct of sorting json")
		}
	}
	//Если есть фильтр с параметрами без подзначение -> требуеться указать тип сортировки (ASC, DESC)
	if (all_object.Column == "name" || all_object.Column == "price") && all_object.Sort != "date" && err == nil {
		if all_object.Meaning == "" {
			if all_object.Sort == "ASC" {
				sqlStatement := `
				SELECT json_agg(
					json_build_object(
			  			'name', name
						, 'price', price
						, 'category', category
						, 'subcategory', subcategory 
						, 'photo', photo
						, 'vcode', goods.vcode
				)) as lenta
				FROM goods
				GROUP BY $1
				ORDER BY $1 ASC`

				rows, _ := db.db.Query(sqlStatement, all_object.Column)

				defer rows.Close()

				for rows.Next() {
					rows.Scan(&lenta)
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(lenta)
			} else {
				if all_object.Sort == "DESC" {
					sqlStatement := `
					SELECT json_agg(
						json_build_object(
			  				'name', name
							, 'price', price
							, 'category', category
							, 'subcategory', subcategory 
							, 'photo', photo
							, 'vcode', goods.vcode
					)) as lenta
					FROM goods
					GROUP BY $1
					ORDER BY $1 DESC`

					rows, _ := db.db.Query(sqlStatement, all_object.Column)

					defer rows.Close()

					for rows.Next() {
						rows.Scan(&lenta)
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(lenta)

				} else {
					fmt.Fprintf(w, "Invalid type of sort")
				}
			}
		} else {
			fmt.Fprintf(w, "Incorrect, there is no MEANING")
		}
	}
}

// ===============================================================MAIN================================================================================
func SetupServer(db *Database) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", db.register_new_goods)
	mux.HandleFunc("/delete", db.delete_good)
	mux.HandleFunc("/update", db.update_good)
	mux.HandleFunc("/select", db.select_good)
	mux.HandleFunc("/addTobasket", db.add2basket)
	mux.HandleFunc("/deleteInBasket", db.delete_in_basket)
	mux.HandleFunc("/viewBasket", db.personal_basket)
	mux.HandleFunc("/buy", db.buy)
	mux.HandleFunc("/storyBought", db.story_purchased)
	mux.HandleFunc("/allAndSort", db.lenta)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func main() {
	db := NewDatabase()
	defer db.Close()

	server := SetupServer(db)

	log.Fatal(server.ListenAndServe())
}
