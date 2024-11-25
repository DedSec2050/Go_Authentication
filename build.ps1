Write-Host "[>>]: Building for production..." -ForegroundColor Cyan

$appName = "pro"

function Build-Unix {
  Write-Host "Building for Unix..." -ForegroundColor Cyan
  
  $env:GOOS = "linux"
  $env:GOARCH = "amd64"
  go build -o "$appName" -ldflags "-X 'main.Release=prod'" ./

  if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful: $appName" -ForegroundColor Green
  } else {
    Write-Host "Build failed" -ForegroundColor Red
  }
}

Build-Unix