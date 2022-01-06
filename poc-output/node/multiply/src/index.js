const run = require("./formula/formula")

const number1 = process.env.NUMBER_1
const number2 = process.env.NUMBER_2
const multiplier = process.env.MULTIPLIER

run(number1, number2, multiplier)
