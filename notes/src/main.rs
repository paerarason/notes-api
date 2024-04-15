use rocket::{get, post, delete, put, routes};
use rocket::serde::{Serialize, Deserialize};

#[derive(Debug, Serialize, Deserialize)]
struct TodoItem {
    id: usize,
    title: String,
    completed: bool,
}

#[post("/todos", data = "<todo>")]
fn create_todo(todo: Json<TodoItem>) -> Json<TodoItem> {
    // Implement logic to create a new todo item
}

#[get("/todos")]
fn get_todos() -> Json<Vec<TodoItem>> {
    // Implement logic to fetch all todo items
}

#[put("/todos/<id>", data = "<todo>")]
fn update_todo(id: usize, todo: Json<TodoItem>) -> Json<TodoItem> {
    // Implement logic to update a todo item
}

#[delete("/todos/<id>")]
fn delete_todo(id: usize) -> Json<TodoItem> {
    // Implement logic to delete a todo item
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![create_todo, get_todos, update_todo, delete_todo])
}
