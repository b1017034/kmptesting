package com.example.kmp_testing

interface Platform {
    val name: String
}

expect fun getPlatform(): Platform