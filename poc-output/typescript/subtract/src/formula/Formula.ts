import * as chalk from 'chalk'
import { api, Output } from './output'

function run(n1: string, n2: string, sub: string) {
    const cmd: string = `rit math sum --number_1=${n1} --number_2=${n2}`

    api.runFormula(cmd)

    const data: Output = api.readOutput()
    const value: number = parseInt(data.output.sum)

    console.log(
        chalk.blue(
            chalk.bold('\nResult:'),
            parseInt(sub) - value,
            '\n'
        )
    )
}

const Formula = run
export default Formula
