use actix_web::{web, App, HttpServer};

mod db;
mod model;
mod service;
mod dto;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let client = db::TodoModel::init().await;

    println!("Starting http server at port 5000");

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(client.clone()))
            .service(service::find_all)
            .service(service::find_todo)
            .service(service::create)
            .service(service::update)
            .service(service::delete)
    })
    .bind("127.0.0.1:5000")?
    .run()
    .await
}
