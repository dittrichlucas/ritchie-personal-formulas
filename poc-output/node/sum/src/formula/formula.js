const clc = require("cli-color")
const api = require("./output.js")

function Run(n1, n2) {
    const sum = parseInt(n1) + parseInt(n2)

    console.info(clc.bold("\nSum: " + sum))

    const data = {
        input: "Funfou?",
        output: String(sum)
    }

    api.writeOutput(data)
}

const formula = Run
module.exports = formula
