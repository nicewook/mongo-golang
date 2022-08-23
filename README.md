# MongoDB Golang REST API Sample

Golang Server를 만들어서 REST API를 통해 MongoDB에 대해 CRUD(Create, Read, Update, Delete)를 수행해본다. 
- 참고 GitHub: https://github.com/DonghanKimAAI/toy-code/tree/training/mongodb/training/java 

## Getting Started

## Prerequisites

- Golang 
- MongoDB: Running server with connect information

## Execution

- Run MongoDB
- Run Golang Server
- Test REST API 

### API LIST
RequestMethod|URL|Parameters
-----|-----|-----
POST|​/api​/v1​/{db}​/{collection}​/insertone|db:database name<br>collection:colletion name<br>parameter:Dynamic Parameter<br>body:RAW JSON
GET|​/api​/v1​/test​/test_col​/find|album_id
GET|​/api​/v1​/test​/test_col​/findone|album_id
GET|​/api​/v1​/test​/vist_count​/findandupdate|userid
GET|​/api​/v1​/test​/vist_count​/total|None
POST|​/api​/v1​/product​/register|JSON
PUT|​/api​/v1​/product​/reviewput|name<br>type<br>tags
PUT|​/api​/v1​/product​/tagput|name<br>type<br>userid<br>comment

### Entity Relationshop Diagram

#### Simple Controller / Service / Dao Model

##### POJO Class

- com.project.sample.SampleEntity

##### Simple Sequence Diagram

```mermaid
sequenceDiagram
  participant SampleController
  participant SampleService
  participant SampleDao
  participant MongoDB
  SampleController->>SampleService: write(Namespace,Documents)
  SampleService->>SampleDao: write(Namespace,Documents)
  SampleDao->>MongoDB: Mongotemplate.write(Namespace,Documents)
  MongoDB->>MongoDB: CRUD Operation
  MongoDB-->>SampleDao: results
  SampleDao-->>SampleService: results
  SampleService-->>SampleController: HTTP Status
```

```mermaid
classDiagram
  class SampleEntity
  SampleEntity : ObjectId id
  SampleEntity : String ablum_id
  SampleEntity : String track_id
  SampleEntity : Integer a
```


#### Spring Data Model

##### Simple Sequence Diagram

```mermaid
sequenceDiagram
  participant StandardController
  participant ProductRepository
  participant SpringData
  participant MongoDB
  StandardController->>ProductRepository: ProductRepository<Product>
  ProductRepository->>SpringData: MongoRepository<Product>
  SpringData->>MongoDB: MongoRepository
  MongoDB->>MongoDB: CRUD Operation
  MongoDB-->>SpringData: results
  SpringData-->>ProductRepository: results
  ProductRepository-->>StandardController: HTTP Status
```

##### POJO Class

- com.project.std.dto.Product
- com.project.std.dto.Release
- com.project.std.dto.Review

```mermaid
classDiagram
  class Product
  class Release
  class Review
  Product : Object _id
  Product : String name
  Product : String type
  Product : ArrayList tags
  Product : Release release
  Product : ArrayList<Review> review
  Release : Integer version
  Release : String date
  Review : String name
  Review : String type
  Review : String userid
  Review : String comment
  Review : Date createDate
```

