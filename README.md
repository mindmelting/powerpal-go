# Powerpal Go Library

## Usage

```go
import (
    "fmt"
    pp "github.com/mindmelting/powerpal_go"
)

func main() {
    var powerpal = pp.New("auth_key", "device_id")

    data, err := p.getData()
    
    fmt.print(data.TotalWattHours)
}

```