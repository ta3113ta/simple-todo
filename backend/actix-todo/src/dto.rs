use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::model::Todo;

#[derive(Serialize, Deserialize, Clone, Debug)]
pub struct CreateTodoDto {
    pub title: String,
    pub note: String,
    pub completed: bool,
}

impl Into<Todo> for CreateTodoDto {
    fn into(self) -> Todo {
        Todo {
            id: Uuid::new_v4().to_string(),
            title: self.title,
            note: self.note,
            completed: self.completed,
        }
    }
}
