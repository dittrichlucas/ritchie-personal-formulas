import { existsSync, readFileSync, unlinkSync, writeFileSync } from 'fs'
import { execSync } from 'child_process'
import { homedir } from 'os'

const outputFile: string = homedir + '/.rit/output.json'

export type Output = {
    output: {
        sum: string
    }
}

function runFormula(cmd: string): void {
    try {
        execSync(cmd)
    } catch (error) {
        throw error
    }
}

function writeOutput(data: Output): void {
    try {
        const dataJSON: string = JSON.stringify(data)

        writeFileSync(outputFile, dataJSON)
    } catch (error) {
        throw error
    }
}

function readOutput(): Output {
    try {
        if (existsSync(outputFile)) {
            const rawData: any = readFileSync(outputFile)
            const parseData: Output = JSON.parse(rawData)

            return parseData
        }
    } catch (error) {
        throw error
    }
}

function removeFile(): void {
    try {
        unlinkSync(outputFile)
    } catch (error) {
        throw error
    }
}

export const api: {
    readOutput: () => Output
    removeFile: () => void
    runFormula: (cmd: string) => void
    writeOutput: (data: Output) => void
} = { readOutput, removeFile, runFormula, writeOutput }

