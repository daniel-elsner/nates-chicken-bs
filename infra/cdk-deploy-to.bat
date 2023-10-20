@findstr /B /V @ %~dpnx0 > %~dpn0.ps1 && powershell -ExecutionPolicy Bypass %~dpn0.ps1 %*
@exit /B %ERRORLEVEL%
if ($args.length -ge 3) {
    $env:CDK_DEPLOY_ACCOUNT, $args = $args
    $env:CDK_DEPLOY_REGION,  $args = $args
    $profile, $args = $args
    npx cdk deploy $args --profile=$profile
    exit $lastExitCode
} else {
    [console]::error.writeline("Provide account, region, and profile as the first three args.")
    [console]::error.writeline("Additional args are passed through to cdk deploy.")
    exit 1
}