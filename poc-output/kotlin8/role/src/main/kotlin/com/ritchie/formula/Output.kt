package com.ritchie.formula

import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import java.io.File
import java.io.FileWriter
import java.lang.Exception

// TODO: documentar os métodos

class Output {
    private val homeDir: String = System.getProperty("user.home")
    private val outputFile: String = "$homeDir/.rit/output.json"

    private val fileNotFound = NoSuchFileException(File(outputFile), reason = "no such file")
    private val gson = Gson()
    private val type = object : TypeToken<Map<String, Any>>() {}.type

    fun runCmd(cmd: String) {
        val cmd = Runtime.getRuntime().exec(cmd)
        cmd.waitFor()

        val exitCode = cmd.exitValue()

        if (exitCode > 0) { throw RuntimeException("command not executed") }
    }

    fun setOutput(key: String, value: Any) {
        val checkFileExist: Boolean = isFileExist()
        if (checkFileExist) {
            val data = parseData()
            val newData = data + mapOf(key to value)

            FileWriter(File(outputFile)).use { writer -> writer.write(gson.toJson(newData)) }
        } else {
            val data = mapOf(key to value)
            val jsonData = gson.toJson(data)
            File(outputFile).writeText(jsonData)
        }
    }

    fun getOutput(key: String): Any? {
        val checkFileExist: Boolean = isFileExist()
        if (checkFileExist) {
            val data = parseData()

            return data[key]

        } else throw fileNotFound
    }

    fun removeOutput(key: String) {
        val checkFileExist: Boolean = isFileExist()
        if (checkFileExist) {
            val data = parseData()

            val removed = data.remove(key)
            if (removed === null || removed === "") {
                throw Exception("key not found")
            } else {
                FileWriter(File(outputFile)).use { writer ->
                    writer.write(gson.toJson(data))
                }
            }
        } else throw fileNotFound
    }

    private fun isFileExist(): Boolean {
        return File(outputFile).exists()
    }

    private fun parseData(): MutableMap<String, Any> {
        val jsonData = File(outputFile).readText()

        return gson.fromJson(jsonData, type)
    }
}
