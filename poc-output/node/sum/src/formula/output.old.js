const { readFileSync, unlinkSync, writeFileSync } = require("fs")
const { ppid } = require("process")
const exec = require("child_process").exec
const homeDir = require('os').homedir

const ritchieDir = (pidValue) => {
    return homeDir + `/.rit/${pidValue}-output.json`
}

// Validar se o comando é válido
function runFormula(cmd) {
    proc = exec(cmd)
    childPID = proc.pid

    // console.log(`Child PID: ${childPID}`)
    return childPID

}

// Validar se é o arquivo certo (salvar na temp com o PID do processo ou algo do gênero)
function writeOutput(processPid, data) {
    const dataJSON = JSON.stringify(data)

    writeFileSync(ritchieDir(processPid), dataJSON)
}

// Validar se o arquivo existe!
function readOutput(myFilePID) {
    const rawData = readFileSync(ritchieDir(myFilePID))
    const parseData = JSON.parse(rawData)

    return parseData
}

// Validar se o value está no objeto data
function getValue(proPID, value) {
    // getFilename()
    const data = readOutput(proPID)

    return data[value]
}

function removeFile() {
    unlinkSync(ritchieDir)
}

// function getFilename() {
//     pid23 = process.pid
//     ppid24 = process.ppid

//     console.log(`PPID: ${ppid24}`)
//     console.log(`PID: ${pid23}`)
// }

const api = { getValue, removeFile, runFormula, writeOutput }
module.exports = api
