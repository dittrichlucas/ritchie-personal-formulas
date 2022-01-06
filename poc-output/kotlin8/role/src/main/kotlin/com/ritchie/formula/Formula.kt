package com.ritchie.formula

import com.github.tomaslanger.chalk.Chalk

class Formula(private val inputRole: String) {
    private val api = Output()

    fun run() {
//        println(Chalk.on(String.format("Your role is %s.", inputRole)).green())
        api.runCmd("ls")
        val key = "role"
//
        api.setOutput(key, inputRole)
//        val result = Output().getOutput(key)
//        println(result)
//        val cmd = api.runCmd("rit -v")
//        println(cmd)
//        val result = api.removeOutput("test")
//        println(result)
//        println(test.output["result"])
//        Output().removeOutput("test")
//        println(removeKey)
    }
}
