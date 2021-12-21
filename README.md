# Powerpal Go Library

## Usage

```go
import (
    "fmt"
    pp "github.com/mindmelting/powerpalgo"
)

func main() {
    var powerpal = pp.New("auth_key", "device_id")

    data, err := powerpal.getData()
    
    fmt.println(data.TotalWattHours)
}

```