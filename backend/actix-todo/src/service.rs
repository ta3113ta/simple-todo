use actix_web::{delete, get, patch, post, web, HttpResponse};

use crate::{db::TodoModel, model::Todo};

#[get("/todo")]
pub async fn find_all(model: web::Data<TodoModel>) -> HttpResponse {
    match model.find().await {
        Ok(todos) => HttpResponse::Ok().json(todos),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

#[get("/todo/{id}")]
pub async fn find_todo(model: web::Data<TodoModel>, id: web::Path<String>) -> HttpResponse {
    let id = id.into_inner();
    println!("id is: {}", id);
    match model.find_one(id).await {
        Ok(Some(todo)) => HttpResponse::Ok().json(todo),
        Ok(None) => HttpResponse::NotFound().body("No todo found"),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

#[post("/todo")]
pub async fn create(model: web::Data<TodoModel>, todo_insert: web::Json<Todo>) -> HttpResponse {
    match model.create(todo_insert.into_inner()).await {
        Ok(_) => HttpResponse::Created().body("created todo"),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

#[patch("/todo/{id}")]
pub async fn update(
    model: web::Data<TodoModel>,
    id: web::Path<String>,
    todo_update: web::Json<Todo>,
) -> HttpResponse {
    match model
        .update(id.into_inner(), todo_update.into_inner())
        .await
    {
        Ok(_) => HttpResponse::Ok().body("updated"),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

#[delete("/todo/{id}")]
pub async fn delete(model: web::Data<TodoModel>, id: web::Path<String>) -> HttpResponse {
    match model.delete(id.into_inner()).await {
        Ok(_) => HttpResponse::Ok().body("deleted"),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}
