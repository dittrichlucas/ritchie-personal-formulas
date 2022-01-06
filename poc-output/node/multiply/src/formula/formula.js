const clc = require('cli-color')
const api = require('./output.js')

function Run(n1, n2, multiplier) {
    api.runFormula(`rit math sum --number_1=${n1} --number_2=${n2}`)
    const data = api.readOutput()
    // const value = parseInt(data.output.sum)

    console.log(data)
    api.writeOutput("vai pocaria")

    // console.log(
    //     clc.blue(
    //         clc.bold('\nResult:'),
    //         value * parseInt(multiplier),
    //         '\n'
    //     )
    // )

    // api.remove()
}

const formula = Run
module.exports = formula
