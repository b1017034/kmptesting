package com.example.kmp_testing.model

data class Todo(
  val title: String,
  val content: String,
  var status: TodoStatus
)

class Todo2(
  val title: String,
  val content: String,
  var status: TodoStatus
)
enum class TodoStatus {
  Todo, Doing, Done
}