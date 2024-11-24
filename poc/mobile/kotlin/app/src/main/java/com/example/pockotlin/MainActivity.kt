package com.example.pockotlin

import android.content.Intent
import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import androidx.appcompat.app.AppCompatActivity
import com.example.pockotlin.databinding.ActivityMainBinding
import com.example.pockotlin.ui.dashboard.DashboardFragment

class MainActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)

        // Initialize EditText and Button
        val inputText = findViewById<EditText>(R.id.inputText)
        val navigateButton = findViewById<Button>(R.id.navigateButton)

        // Set up button click listener
        navigateButton.setOnClickListener {
            val intent = Intent(this, DashboardFragment::class.java)
            intent.putExtra("userInput", inputText.text.toString())
            startActivity(intent)
        }
    }
}
