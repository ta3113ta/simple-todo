use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug)]
pub struct Todo {
    pub id: String,
    pub title: String,
    pub note: String,
    pub completed: bool,
}