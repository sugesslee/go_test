package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/syndtr/goleveldb/leveldb"
)

func main(){
	fmt.Println("hello")

	// 创建或打开一个数据库
	db,err := leveldb.OpenFile("./data/block1", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 存入数据
	_ = db.Put([]byte("1"), []byte("6"), nil)
	_ = db.Put([]byte("2"), []byte("7"), nil)
	_ = db.Put([]byte("3"), []byte("8"), nil)
	_ = db.Put([]byte("foo-4"), []byte("9"), nil)
	_ = db.Put([]byte("5"), []byte("10"), nil)
	_ = db.Put([]byte("6"), []byte("11"), nil)
	_ = db.Put([]byte("moo-7"), []byte("12"), nil)
	_ = db.Put([]byte("8"), []byte("13"), nil)

	// 遍历数据库内容
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		panic(err)

	}

	fmt.Println("***************************************************")

	// 删除某条数据
	err = db.Delete([]byte("2"), nil)

	// 读取某条数据
	data, err := db.Get([]byte("2"), nil)
	fmt.Printf("[2]:%s:%s\n",data,err)

	// 根据前缀遍历数据库内容
	fmt.Println("***************************************************")
	iter = db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err = iter.Error()

	// 遍历从指定 key
	fmt.Println("***************************************************")
	iter = db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte("5")); ok; ok = iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err = iter.Error()

	// 遍历子集范围
	fmt.Println("***************************************************")
	iter = db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("loo")}, nil)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err = iter.Error()

	// 批量操作
	fmt.Println("***************************************************")
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	err = db.Write(batch, nil)

	// 遍历数据库内容
	iter = db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err = iter.Error()

}