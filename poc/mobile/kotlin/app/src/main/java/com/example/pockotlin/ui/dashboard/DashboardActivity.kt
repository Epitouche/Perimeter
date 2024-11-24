package com.example.pockotlin.ui.dashboard

import android.os.Bundle
import android.widget.Button
import android.widget.Spinner
import androidx.appcompat.app.AppCompatActivity
import com.example.pockotlin.R

class SelectActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.fragment_dashboard)

        val dropdownMenu = findViewById<Spinner>(R.id.dropdownMenu)
        val backButton = findViewById<Button>(R.id.backButton)

        // Example: Retrieving input text from MainActivity
        val userInput = intent.getStringExtra("userInput")

        backButton.setOnClickListener {
            finish() // Go back to the previous activity
        }
    }
}