package com.ritchie

import com.ritchie.formula.Formula

object Main {
    @JvmStatic
    fun main(args: Array<String?>?) {
        val inputRole: String = System.getenv("RIT_INPUT_ROLE")
        val formula = Formula(inputRole)
        formula.run()
    }
}
// https://www.hostinger.com.br/tutoriais/install-maven-ubuntu
// https://kotlin-code.com/ide/visual-studio-code/setup/
