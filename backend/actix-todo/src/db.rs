use futures_util::StreamExt;
use mongodb::{
    bson::{self, doc, oid::ObjectId},
    Client, Collection,
};

use crate::{dto::CreateTodoDto, model::Todo};

const DB_NAME: &str = "projects";
const COLL_NAME: &str = "todos";

#[derive(Clone)]
pub struct TodoModel {
    pub client: Client,
}

impl TodoModel {
    pub async fn init() -> Self {
        let mongo_url: &str = "mongodb://localhost:2717";
        let client = Client::with_uri_str(mongo_url)
            .await
            .expect("failed to connect");

        Self { client }
    }

    pub async fn find(&self) -> std::result::Result<Vec<Todo>, mongodb::error::Error> {
        let collection: Collection<Todo> = self.client.database(DB_NAME).collection(COLL_NAME);
        let mut todos: Vec<Todo> = Vec::new();
        let mut cursor = collection.find(None, None).await?;

        while let Some(todo) = cursor.next().await {
            todos.push(todo.expect("can not get todo"))
        }

        Ok(todos)
    }

    pub async fn find_one(&self, id: String) -> Result<Option<Todo>, mongodb::error::Error> {
        let collection: Collection<Todo> = self.client.database(DB_NAME).collection(COLL_NAME);
        let obj_id = ObjectId::parse_str(id).unwrap();
        match collection.find_one(doc! { "_id": obj_id }, None).await {
            Ok(todo) => Ok(todo),
            Err(err) => Err(err),
        }
    }

    pub async fn create(&self, todo_insert: CreateTodoDto) -> Result<(), mongodb::error::Error> {
        let collection: Collection<Todo> = self.client.database(DB_NAME).collection(COLL_NAME);
        let todo_insert: Todo = todo_insert.into();
        let result = collection.insert_one(todo_insert, None).await;
        match result {
            Ok(_) => Ok(()),
            Err(err) => Err(err),
        }
    }

    pub async fn update(&self, id: String, update_todo: Todo) -> Result<(), mongodb::error::Error> {
        let collection: Collection<Todo> = self.client.database(DB_NAME).collection(COLL_NAME);

        let obj_id = ObjectId::parse_str(id).unwrap();
        let filter = doc! {"_id": obj_id};
        let update = doc! {"$set" : bson::to_bson(&update_todo).unwrap()};

        let result = collection.find_one_and_update(filter, update, None).await;
        match result {
            Ok(_) => Ok(()),
            Err(err) => Err(err),
        }
    }

    pub async fn delete(&self, id: String) -> Result<(), mongodb::error::Error> {
        let collection: Collection<Todo> = self.client.database(DB_NAME).collection(COLL_NAME);

        let obj_id = ObjectId::parse_str(id).unwrap();
        let filter = doc! {"_id": obj_id};

        let result = collection.find_one_and_delete(filter, None).await;
        match result {
            Ok(_) => Ok(()),
            Err(err) => Err(err),
        }
    }
}
