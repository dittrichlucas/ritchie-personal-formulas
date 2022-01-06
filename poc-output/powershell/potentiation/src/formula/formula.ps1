. $PSScriptRoot\output.ps1

Function runFormula () {
    $number_1 = $env:RIT_NUMBER_1
    $number_2 = $env:RIT_NUMBER_2
    $exponent = $env:RIT_EXPONENT

    $cmd = "rit math sum --number_1=$number_1 --number_2=$number_2"

    runCommand $cmd

    $data = readOutput
    $sum = $data.output.sum
    $pow = [Math]::Pow($sum, $exponent)
    Write-Output $pow

    $output = @{
        output = @{
            pow = $pow
        }
    }

    writeOutput $output
}
