package com.ritchie.formula
import com.github.tomaslanger.chalk.Chalk;

class Formula() {
    private val api = Output()
    fun run() {
        api.runCmd("rit poc kotlin8 role --rit_input_role=QA")
        val role = api.getOutput("role")
        println(role)
    }
}
