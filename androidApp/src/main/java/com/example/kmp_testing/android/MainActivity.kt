package com.example.kmp_testing.android

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.example.kmp_testing.Greeting
import com.example.kmp_testing.feature.todo.TodoUseCase
import com.example.kmp_testing.model.Todo

class MainActivity : ComponentActivity() {
    private val todoUseCase = TodoUseCase(
        listOf(
            Todo("勉強"),
            Todo("お風呂入る")
        )
    )

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            MyApplicationTheme {
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    MainView(Greeting().greet(), todoUseCase.todos)
                }
            }
        }
    }
}

@Composable
fun ListView(todos: List<Todo>) {
    Column {
        todos.forEach { todo ->
            Text(text = todo.title)
        }
    }
}

@Composable
fun MainView(text: String, todos: List<Todo>) {
    Column {
        Text(text = text)
        ListView(todos = todos)
    }
}

@Preview
@Composable
fun DefaultPreview() {
    val todoUseCase = TodoUseCase(
        listOf(
            Todo("勉強"),
            Todo("お風呂入る")
        )
    )

    MyApplicationTheme {
        MainView("Hello, Android!", todoUseCase.todos)
    }
}
