package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func smain() {
	db, err := bolt.Open("test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	// 找到表，不存在就去创建
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				return err
			}
		}

		err = bucket.Put([]byte("name"), []byte("chang"))
		err = bucket.Put([]byte("age"), []byte("20"))
		err = bucket.Put([]byte("address"), []byte("anhui"))

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err)
		return
	}
}
