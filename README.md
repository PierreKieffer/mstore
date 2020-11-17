# mstore 

<p align="center">
  <img src="logo.png">
</p>


![Go](https://github.com/PierreKieffer/mstore/workflows/Go/badge.svg)

In-memory storage management system for embedded data storage.

As an example of application, mstore can be used as the core of a cache system, or simply as an embedded database.

mstore uses an indexing system to optimize operations performance.
mstore can manage multiple documents on the same index, but it is not possible to have several documents with the same key (the key is unique and is used to reference a document).


						      +-------+             data memory slot
						      |       |            +--------+-----------+
					       index  |       |            |        |           |
				     +-------> memory | Index +----------->+ Key_i  | Data      |
				     |         slot   |   01  |            |        |           |
				     |                |       |            +--------+-----------+
				     |                +-------+
				     |                |       |
				     |                |       |
				     |                |       |
	+-----------+                |                |       |
	|           |                |                |       |
	| Data      +----------------+                +-------+
	|           | hash function                   |       |            +--------+-----------+             +--------+----------+
	+-----------+                                 |       |            |        |           |             |        |          |
						      | Index +----------->+ Key_j  |  Data     +------------>+ Key_k  | Data     |
						      |   ..  |            |        |           |             |        |          |
						      |       |            +--------+-----------+             +--------+----------+
						      +-------+
						      |       |
						      |       |
						      |       |
						      |       |
						      |       |
						      +-------+
						      |       |
						      |       |
						      |       |
						      +       +


## Download
```bash 
go get github.com/PierreKieffer/mstore
```
## Usage 

```go
import(
	"github.com/PierreKieffer/mstore"
)
```
### Initialize storage 

```go 
s := mstore.InitStorage()
```
By default, the storage is built with 1000 indexes,

It is possible to configure the maximum number of indexes that make up the storage

```go
s := mstore.InitStorage(10)
```

### Document type 
mstore uses a Document type : 

``` go 
type Document struct {
        Key  string
        Data interface{}
}
```
Documents are referenced by document key which must be unique.

If no custom key is passed, a key is generated in UnixNano format.

### Insert  

```go
var data interface{} 
document := mstore.Document{Key: "custom-key", Data: data}
err = mstore.Insert(s, document)
if err != nil {
	fmt.Println(err)
}
```

### Update 

```go
err = mstore.Update(s, document)
if err != nil {
	fmt.Println(err)
}
```

### Delete 

```go
err := mstore.Delete(s, key)
if err != nil {
	fmt.Println(err)
}
```

### Find 

```go
doc, err := mstore.Find(s, key)
if err != nil {
	fmt.Println(err)
}
```





