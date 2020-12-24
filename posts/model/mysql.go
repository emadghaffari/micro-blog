package model

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-sql-driver/mysql"
	log "github.com/micro/micro/v3/service/logger"

	posts "github.com/emadghaffari/micro-blog/posts/proto"
	"github.com/emadghaffari/res_errors/logger"
)

type model struct {
	store *sql.DB
	// helps logically separate keys in a model where
	// multiple `Model`s share the same underlying
	// physical database.
	namespace string
}

// Model represents a place where data can be saved to and
// queried from.
type Model interface {
	// Save any object. Maintains indexes set up.
	Save(query interface{}) error
	// List objects by a query. Each query requires an appropriate index
	// to exist. List throws an error if a matching index can't be found.
	List(query interface{}, resultSlicePointer interface{}) error
	// Same as list, but accepts pointer to non slices and
	// expects to find only one element. Throws error if not found
	// or if more than two elements are found.
	Read(query string, resultPointer interface{}) error
	// Deletes a record. Delete only support Equals("id", value) for now.
	// @todo Delete only supports string keys for now.
	Delete(query string) error
}

// New db
func New(namespace string) Model {
	var err error
	username := "root"
	password := "password"
	host := "db"
	schema := "golang"
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal("ERRPR in DB: ", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal("ERRPR in DB: ", err.Error())
	}
	mysql.SetLogger(logger.GetLogger())
	log.Info("database successfuly confiqured.")

	return &model{
		store:     db,
		namespace: namespace,
	}
}

func (m *model) Save(req interface{}) error {
	ms, err := m.marshaller(req)
	if err != nil {
		return err
	}
	query := ""
	if m.namespace == "posts" {
		query = "INSERT INTO posts(title,slug,description,user_id,image) VALUES(?,?,?,?,?);"
	}
	if query == "" {
		return fmt.Errorf("invalid database namespace")
	}
	tx, err := m.store.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	if err := m.store.QueryRow(query, ms["title"], ms["slug"], ms["description"], ms["user_id"], ms["image"]).Err(); err != nil {
		return err
	}

	log.Info("Stored into DB successfully")
	return nil
}

func (m *model) List(req interface{}, resultSlicePointer interface{}) error {
	ms, err := m.marshaller(req)
	if err != nil {
		return err
	}
	query := ""
	if m.namespace == "posts" {
		query = "SELECT id,title,slug,description,user_id,image FROM posts "
		if len(ms) > 0 {
			query += "WHERE "
			for k, v := range ms {

				query += k + "=" + fmt.Sprintf("%s", v) + "AND "

			}
			// remove last AND
			query = query[:len(query)-4]
		}
	}
	if query == "" {
		return fmt.Errorf("invalid database namespace")
	}
	tx, err := m.store.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	row, err := m.store.Query(query)
	if err != nil {
		return err
	}
	defer row.Close()
	result := make([]*posts.Post, 0)
	var img *string
	for row.Next() {
		var item posts.Post
		if err := row.Scan(&item.Id, &item.Title, &item.Slug, &item.Description, &item.Author, &img); err != nil {
			return err
		}
		if img != nil {
			item.Image = *img
		}

		result = append(result, &item)
	}
	bt, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return json.Unmarshal(bt, resultSlicePointer)
}

func (m *model) Read(query string, resultPointer interface{}) error {
	return nil
}

func (m *model) Delete(query string) error {
	return nil
}

func (m *model) marshaller(req interface{}) (map[string]interface{}, error) {
	js, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	ms := map[string]interface{}{}
	de := json.NewDecoder(bytes.NewReader(js))
	de.UseNumber()
	err = de.Decode(&ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
