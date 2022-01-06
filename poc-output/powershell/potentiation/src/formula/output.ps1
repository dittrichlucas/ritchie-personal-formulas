$homedir = $HOME + "/.rit/output.json"
$homedirnew = $HOME + "/.rit/output-new.json"

Function runCommand ($cmd) {
    Invoke-Expression $cmd
}

Function writeOutput ($output) {
    $output | ConvertTo-Json -Depth 100 | Out-File $homedirnew
}

Function readOutput () {
    $json = Get-Content -Raw -Path "$homedir"
    ConvertFrom-Json -InputObject $json
}

Function removeFile () {
    Remove-Item -Path $homedir
}
