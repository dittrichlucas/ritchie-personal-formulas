const { existsSync, readFileSync, unlinkSync, writeFileSync } = require("fs")
const { execSync } = require("child_process")
const { homedir } = require('os')

const outputPath = homedir + '/.rit/output.json'

function runFormula(cmd) {
    try {
        execSync(cmd)
    } catch (error) {
        throw error
    }
}

function writeOutput(data) {
    try {
        const dataJSON = JSON.stringify(data)

        writeFileSync(outputPath, dataJSON)
    } catch (error) {
        throw error
    }
}

function readOutput() {
    try {
        if(existsSync(outputPath)) {
            const rawData = readFileSync(outputPath)
            const parseData = JSON.parse(rawData)

            return parseData
        }
    } catch (error) {
        throw error
    }
}

function removeFile() {
    try {
        unlinkSync(outputPath)
    } catch (error) {
        throw error
    }
}

const api = { readOutput, removeFile, runFormula, writeOutput }
module.exports = api
