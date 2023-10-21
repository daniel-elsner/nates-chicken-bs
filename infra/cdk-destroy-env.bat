@findstr /B /V @ %~dpnx0 > %~dpn0.ps1 && powershell -ExecutionPolicy Bypass %~dpn0.ps1 %*
@del %~dpn0.ps1
@exit /B %ERRORLEVEL%

if ($args.length -ge 4) {
    $env:CDK_ENV_NAME, $args = $args
    $env:CDK_DEPLOY_ACCOUNT, $args = $args
    $env:CDK_DEPLOY_REGION,  $args = $args
    $profile, $args = $args
    npx cdk destroy $args --profile=$profile
    exit $lastExitCode
} else {
    [console]::error.writeline("Provide environment name, account, region, and profile as the first four args.")
    [console]::error.writeline("Additional args are passed through to cdk deploy.")
    exit 1
}