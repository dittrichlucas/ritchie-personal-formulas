import run from './formula/Formula'

const n1: string = process.env.NUMBER_1
const n2: string = process.env.NUMBER_2
const subtract: string = process.env.SUBTRACT

run(n1, n2, subtract)
